package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DefiantLabs/cosmos-indexer/config"
	dbTypes "github.com/DefiantLabs/cosmos-indexer/db"
	"github.com/DefiantLabs/cosmos-indexer/juno"
	"github.com/DefiantLabs/cosmos-indexer/osmosis"
	"github.com/DefiantLabs/cosmos-indexer/rest"

	"gorm.io/gorm"
)

type AssetList struct {
	Assets []Asset
}

type Asset struct {
	Denoms []DenomUnit `json:"denom_units"`
	Symbol string
	Base   string
	Name   string
}

type DenomUnit struct {
	Denom    string
	Exponent int
	Aliases  []string
}

var ChainSpecificDenomUpsertFunctions = map[string]func(db *gorm.DB){
	osmosis.ChainID: UpsertOsmosisDenoms,
	juno.ChainID:    UpsertJunoDenoms,
}

func DoChainSpecificUpsertDenoms(db *gorm.DB, chain string) {
	if chain == osmosis.ChainID {
		UpsertOsmosisDenoms(db)
	}

	if chain == juno.ChainID {
		UpsertJunoDenoms(db)
	}
	// may want to move this elsewhere, or eliminate entirely
	// I would prefer we just grab the denoms when needed always
	// Current problem: we use the denom cache in various blocks later on
	dbTypes.CacheDenoms(db)
	dbTypes.CacheIBCDenoms(db)
}

func UpsertOsmosisDenoms(db *gorm.DB) {
	config.Log.Info("Updating Omsosis specific denoms")
	url := "https://raw.githubusercontent.com/osmosis-labs/assetlists/main/osmosis-1/osmosis-1.assetlist.json"

	denomAssets, err := getAssetsList(url)
	if err != nil {
		config.Log.Fatal("Download Osmosis Denom Metadata", err)
	} else {
		denoms := assetListToDenoms(denomAssets)
		err = dbTypes.UpsertDenoms(db, denoms)
		if err != nil {
			config.Log.Fatal("Upsert Osmosis Denom Metadata", err)
		}
	}
}

func UpsertJunoDenoms(db *gorm.DB) {
	config.Log.Info("Updating Juno specific denoms")
	url := "https://raw.githubusercontent.com/cosmos/chain-registry/master/juno/assetlist.json"

	denomAssets, err := getAssetsList(url)
	if err != nil {
		config.Log.Fatal("Error downloading Juno Denom Metadata", err)
	} else {
		denoms := assetListToDenoms(denomAssets)
		err = dbTypes.UpsertDenoms(db, denoms)
		if err != nil {
			config.Log.Fatal("Error upserting Juno Denom Metadata", err)
		}
	}
}

func assetListToDenoms(assets *AssetList) []dbTypes.DenomDBWrapper {
	denoms := make([]dbTypes.DenomDBWrapper, len(assets.Assets))
	for i, asset := range assets.Assets {
		denoms[i].Denom = dbTypes.Denom{Base: asset.Base, Name: asset.Name, Symbol: asset.Symbol}
		denoms[i].DenomUnits = make([]dbTypes.DenomUnitDBWrapper, len(asset.Denoms))

		for ii, denomUnit := range asset.Denoms {
			denoms[i].DenomUnits[ii].DenomUnit = dbTypes.DenomUnit{Exponent: uint(denomUnit.Exponent), Name: denomUnit.Denom}
		}
	}

	return denoms
}

func getAssetsList(assetsURL string) (*AssetList, error) {
	assets := &AssetList{}
	err := getJSON(assetsURL, assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func getJSON(url string, target interface{}) error {
	myClient := &http.Client{Timeout: 10 * time.Second}

	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("got status code: %v from url: %v", r.Status, url)
	}

	return json.NewDecoder(r.Body).Decode(target)
}

func IBCDenomUpsertTask(apiHost string, db *gorm.DB) {
	config.Log.Info(fmt.Sprintf("Updating IBC Denom Metadata from %s", apiHost))
	ibcDenomsMetadata, err := rest.GetIBCDenomTraces(apiHost)
	if err != nil {
		config.Log.Error(fmt.Sprintf("Error in IBC Denom Metadata Update task when reaching out to the API at %s ", apiHost), err)
		return
	}

	denoms := make([]dbTypes.IBCDenom, len(ibcDenomsMetadata.DenomTraces))
	for i, t := range ibcDenomsMetadata.DenomTraces {
		denoms[i] = dbTypes.IBCDenom{
			Hash:      t.IBCDenom(),
			Path:      t.Path,
			BaseDenom: t.BaseDenom,
		}
	}

	err = dbTypes.UpsertIBCDenoms(db, denoms)
	if err != nil {
		config.Log.Error("Error updating database in IBC Denom Metadata Update task", err)
		return
	}
	config.Log.Info("IBC Denom Metadata Update Complete")
}

func ValidateDenoms(db *gorm.DB) error {
	config.Log.Info("Running post-update denom validations")
	var denoms []dbTypes.Denom

	// Find all denoms which are missing an entry in the denom_units table
	// This is currently required due to a bug that was introduced by dropping the denom_units table without thinking fully through the consequences
	// We may want to remove this at some point, since UNKNOWN denoms get a single denom_unit added during indexing already
	res := db.Joins("LEFT JOIN denom_units ON denoms.id = denom_units.denom_id").Where("denom_units.denom_id IS NULL").Find(&denoms)

	if res.Error != nil {
		config.Log.Error("Error getting denoms in denom validator", res.Error)
		return res.Error
	}

	if len(denoms) > 0 {
		config.Log.Infof("Adding missing denom units for %d denoms", len(denoms))
		err := db.Transaction(func(dbTransaction *gorm.DB) error {
			for _, denom := range denoms {
				missingBaseDenomUnit := dbTypes.DenomUnit{DenomID: denom.ID, Name: denom.Base, Exponent: 0}
				txRes := db.Create(&missingBaseDenomUnit)
				if txRes.Error != nil {
					return txRes.Error
				}
			}
			return nil
		})
		if err != nil {
			config.Log.Error("Error backfilling missing denom_units in validator", res.Error)
			return err
		}
	} else {
		config.Log.Info("All denoms have at least 1 corresponding denom_unit")
	}

	return nil
}
