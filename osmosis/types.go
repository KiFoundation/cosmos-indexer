package osmosis

import (
	"net/http"
	"time"
)

const (
	ChainID = "osmosis-1"
	Name    = "osmosis"
)

type Result struct {
	Data Data
}

type Data struct {
	Value EventDataNewBlockHeader
}

type EventDataNewBlockHeader struct {
	Header Header `json:"header"`
	NumTxs string `json:"num_txs"` // Number of txs in a block
}

type Header struct {
	// basic block info
	ChainID string    `json:"chain_id"`
	Height  string    `json:"height"`
	Time    time.Time `json:"time"`
}

type TendermintNewBlockHeader struct {
	Result Result
}

type URIClient struct {
	Address    string
	Client     *http.Client
	AuthHeader string
}
