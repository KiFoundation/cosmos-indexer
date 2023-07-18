--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: taxuser
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO taxuser;

\connect postgres

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: taxuser
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: addresses; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.addresses (
    id bigint NOT NULL,
    address text
);


ALTER TABLE public.addresses OWNER TO taxuser;

--
-- Name: addresses_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.addresses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.addresses_id_seq OWNER TO taxuser;

--
-- Name: addresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.addresses_id_seq OWNED BY public.addresses.id;


--
-- Name: blocks; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.blocks (
    id bigint NOT NULL,
    time_stamp timestamp with time zone,
    height bigint,
    blockchain_id bigint,
    indexed boolean
);


ALTER TABLE public.blocks OWNER TO taxuser;

--
-- Name: blocks_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.blocks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.blocks_id_seq OWNER TO taxuser;

--
-- Name: blocks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.blocks_id_seq OWNED BY public.blocks.id;


--
-- Name: chains; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.chains (
    id bigint NOT NULL,
    chain_id text,
    name text
);


ALTER TABLE public.chains OWNER TO taxuser;

--
-- Name: chains_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.chains_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.chains_id_seq OWNER TO taxuser;

--
-- Name: chains_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.chains_id_seq OWNED BY public.chains.id;


--
-- Name: denom_units; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.denom_units (
    id bigint NOT NULL,
    denom_id bigint,
    exponent bigint,
    name text
);


ALTER TABLE public.denom_units OWNER TO taxuser;

--
-- Name: denom_units_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.denom_units_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.denom_units_id_seq OWNER TO taxuser;

--
-- Name: denom_units_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.denom_units_id_seq OWNED BY public.denom_units.id;


--
-- Name: denoms; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.denoms (
    id bigint NOT NULL,
    base text,
    name text,
    symbol text
);


ALTER TABLE public.denoms OWNER TO taxuser;

--
-- Name: denoms_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.denoms_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.denoms_id_seq OWNER TO taxuser;

--
-- Name: denoms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.denoms_id_seq OWNED BY public.denoms.id;


--
-- Name: epoches; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.epoches (
    id bigint NOT NULL,
    blockchain_id bigint,
    start_height bigint,
    identifier text,
    epoch_number bigint
);


ALTER TABLE public.epoches OWNER TO taxuser;

--
-- Name: epoches_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.epoches_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.epoches_id_seq OWNER TO taxuser;

--
-- Name: epoches_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.epoches_id_seq OWNED BY public.epoches.id;


--
-- Name: failed_blocks; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.failed_blocks (
    id bigint NOT NULL,
    height bigint,
    blockchain_id bigint
);


ALTER TABLE public.failed_blocks OWNER TO taxuser;

--
-- Name: failed_blocks_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.failed_blocks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.failed_blocks_id_seq OWNER TO taxuser;

--
-- Name: failed_blocks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.failed_blocks_id_seq OWNED BY public.failed_blocks.id;


--
-- Name: failed_event_blocks; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.failed_event_blocks (
    id bigint NOT NULL,
    height bigint,
    blockchain_id bigint
);


ALTER TABLE public.failed_event_blocks OWNER TO taxuser;

--
-- Name: failed_event_blocks_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.failed_event_blocks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.failed_event_blocks_id_seq OWNER TO taxuser;

--
-- Name: failed_event_blocks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.failed_event_blocks_id_seq OWNED BY public.failed_event_blocks.id;


--
-- Name: fees; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.fees (
    id bigint NOT NULL,
    tx_id bigint,
    amount numeric(78,0),
    denomination_id bigint,
    payer_address_id bigint
);


ALTER TABLE public.fees OWNER TO taxuser;

--
-- Name: fees_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.fees_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fees_id_seq OWNER TO taxuser;

--
-- Name: fees_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.fees_id_seq OWNED BY public.fees.id;


--
-- Name: ibc_denoms; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.ibc_denoms (
    id bigint NOT NULL,
    hash text,
    path text,
    base_denom text
);


ALTER TABLE public.ibc_denoms OWNER TO taxuser;

--
-- Name: ibc_denoms_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.ibc_denoms_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ibc_denoms_id_seq OWNER TO taxuser;

--
-- Name: ibc_denoms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.ibc_denoms_id_seq OWNED BY public.ibc_denoms.id;


--
-- Name: message_types; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.message_types (
    id bigint NOT NULL,
    message_type text NOT NULL
);


ALTER TABLE public.message_types OWNER TO taxuser;

--
-- Name: message_types_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.message_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.message_types_id_seq OWNER TO taxuser;

--
-- Name: message_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.message_types_id_seq OWNED BY public.message_types.id;


--
-- Name: messages; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.messages (
    id bigint NOT NULL,
    tx_id bigint,
    message_type_id bigint,
    message_index bigint
);


ALTER TABLE public.messages OWNER TO taxuser;

--
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.messages_id_seq OWNER TO taxuser;

--
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.messages_id_seq OWNED BY public.messages.id;


--
-- Name: taxable_event; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.taxable_event (
    id bigint NOT NULL,
    source bigint,
    amount numeric(78,0),
    denomination_id bigint,
    address_id bigint,
    event_hash text,
    block_id bigint
);


ALTER TABLE public.taxable_event OWNER TO taxuser;

--
-- Name: taxable_event_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.taxable_event_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.taxable_event_id_seq OWNER TO taxuser;

--
-- Name: taxable_event_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.taxable_event_id_seq OWNED BY public.taxable_event.id;


--
-- Name: taxable_tx; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.taxable_tx (
    id bigint NOT NULL,
    message_id bigint,
    amount_sent numeric(78,0),
    amount_received numeric(78,0),
    denomination_sent_id bigint,
    denomination_received_id bigint,
    sender_address_id bigint,
    receiver_address_id bigint
);


ALTER TABLE public.taxable_tx OWNER TO taxuser;

--
-- Name: taxable_tx_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.taxable_tx_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.taxable_tx_id_seq OWNER TO taxuser;

--
-- Name: taxable_tx_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.taxable_tx_id_seq OWNED BY public.taxable_tx.id;


--
-- Name: txes; Type: TABLE; Schema: public; Owner: taxuser
--

CREATE TABLE public.txes (
    id bigint NOT NULL,
    hash text,
    code bigint,
    block_id bigint,
    signer_address_id bigint
);


ALTER TABLE public.txes OWNER TO taxuser;

--
-- Name: txes_id_seq; Type: SEQUENCE; Schema: public; Owner: taxuser
--

CREATE SEQUENCE public.txes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.txes_id_seq OWNER TO taxuser;

--
-- Name: txes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: taxuser
--

ALTER SEQUENCE public.txes_id_seq OWNED BY public.txes.id;


--
-- Name: addresses id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.addresses ALTER COLUMN id SET DEFAULT nextval('public.addresses_id_seq'::regclass);


--
-- Name: blocks id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.blocks ALTER COLUMN id SET DEFAULT nextval('public.blocks_id_seq'::regclass);


--
-- Name: chains id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.chains ALTER COLUMN id SET DEFAULT nextval('public.chains_id_seq'::regclass);


--
-- Name: denom_units id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.denom_units ALTER COLUMN id SET DEFAULT nextval('public.denom_units_id_seq'::regclass);


--
-- Name: denoms id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.denoms ALTER COLUMN id SET DEFAULT nextval('public.denoms_id_seq'::regclass);


--
-- Name: epoches id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.epoches ALTER COLUMN id SET DEFAULT nextval('public.epoches_id_seq'::regclass);


--
-- Name: failed_blocks id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.failed_blocks ALTER COLUMN id SET DEFAULT nextval('public.failed_blocks_id_seq'::regclass);


--
-- Name: failed_event_blocks id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.failed_event_blocks ALTER COLUMN id SET DEFAULT nextval('public.failed_event_blocks_id_seq'::regclass);


--
-- Name: fees id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.fees ALTER COLUMN id SET DEFAULT nextval('public.fees_id_seq'::regclass);


--
-- Name: ibc_denoms id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.ibc_denoms ALTER COLUMN id SET DEFAULT nextval('public.ibc_denoms_id_seq'::regclass);


--
-- Name: message_types id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.message_types ALTER COLUMN id SET DEFAULT nextval('public.message_types_id_seq'::regclass);


--
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.messages ALTER COLUMN id SET DEFAULT nextval('public.messages_id_seq'::regclass);


--
-- Name: taxable_event id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_event ALTER COLUMN id SET DEFAULT nextval('public.taxable_event_id_seq'::regclass);


--
-- Name: taxable_tx id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_tx ALTER COLUMN id SET DEFAULT nextval('public.taxable_tx_id_seq'::regclass);


--
-- Name: txes id; Type: DEFAULT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.txes ALTER COLUMN id SET DEFAULT nextval('public.txes_id_seq'::regclass);


--
-- Name: addresses addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (id);


--
-- Name: blocks blocks_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.blocks
    ADD CONSTRAINT blocks_pkey PRIMARY KEY (id);


--
-- Name: chains chains_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.chains
    ADD CONSTRAINT chains_pkey PRIMARY KEY (id);


--
-- Name: denom_units denom_units_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.denom_units
    ADD CONSTRAINT denom_units_pkey PRIMARY KEY (id);


--
-- Name: denoms denoms_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.denoms
    ADD CONSTRAINT denoms_pkey PRIMARY KEY (id);


--
-- Name: epoches epoches_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.epoches
    ADD CONSTRAINT epoches_pkey PRIMARY KEY (id);


--
-- Name: failed_blocks failed_blocks_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.failed_blocks
    ADD CONSTRAINT failed_blocks_pkey PRIMARY KEY (id);


--
-- Name: failed_event_blocks failed_event_blocks_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.failed_event_blocks
    ADD CONSTRAINT failed_event_blocks_pkey PRIMARY KEY (id);


--
-- Name: fees fees_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.fees
    ADD CONSTRAINT fees_pkey PRIMARY KEY (id);


--
-- Name: ibc_denoms ibc_denoms_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.ibc_denoms
    ADD CONSTRAINT ibc_denoms_pkey PRIMARY KEY (id);


--
-- Name: message_types message_types_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.message_types
    ADD CONSTRAINT message_types_pkey PRIMARY KEY (id);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- Name: taxable_event taxable_event_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_event
    ADD CONSTRAINT taxable_event_pkey PRIMARY KEY (id);


--
-- Name: taxable_tx taxable_tx_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_tx
    ADD CONSTRAINT taxable_tx_pkey PRIMARY KEY (id);


--
-- Name: txes txes_pkey; Type: CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.txes
    ADD CONSTRAINT txes_pkey PRIMARY KEY (id);


--
-- Name: chainepochidentifierheight; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX chainepochidentifierheight ON public.epoches USING btree (blockchain_id, start_height, identifier);


--
-- Name: chainheight; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX chainheight ON public.blocks USING btree (height, blockchain_id);


--
-- Name: failedchaineventheight; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX failedchaineventheight ON public.failed_event_blocks USING btree (height, blockchain_id);


--
-- Name: failedchainheight; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX failedchainheight ON public.failed_blocks USING btree (height, blockchain_id);


--
-- Name: idx_addr; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE INDEX idx_addr ON public.taxable_event USING btree (address_id);


--
-- Name: idx_addresses_address; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX idx_addresses_address ON public.addresses USING btree (address);


--
-- Name: idx_chains_chain_id; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX idx_chains_chain_id ON public.chains USING btree (chain_id);


--
-- Name: idx_denom_units_denom_id_name; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX idx_denom_units_denom_id_name ON public.denom_units USING btree (denom_id, name);


--
-- Name: idx_denoms_base; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX idx_denoms_base ON public.denoms USING btree (base);


--
-- Name: idx_ibc_denoms_hash; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX idx_ibc_denoms_hash ON public.ibc_denoms USING btree (hash);


--
-- Name: idx_message_types_message_type; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX idx_message_types_message_type ON public.message_types USING btree (message_type);


--
-- Name: idx_msg; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE INDEX idx_msg ON public.taxable_tx USING btree (message_id);


--
-- Name: idx_payer_addr; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE INDEX idx_payer_addr ON public.fees USING btree (payer_address_id);


--
-- Name: idx_receiver; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE INDEX idx_receiver ON public.taxable_tx USING btree (receiver_address_id);


--
-- Name: idx_sender; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE INDEX idx_sender ON public.taxable_tx USING btree (sender_address_id);


--
-- Name: idx_teblkid; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE INDEX idx_teblkid ON public.taxable_event USING btree (block_id);


--
-- Name: idx_teevthash; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX idx_teevthash ON public.taxable_event USING btree (event_hash);


--
-- Name: idx_txes_hash; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX idx_txes_hash ON public.txes USING btree (hash);


--
-- Name: idx_txid_typeid; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE INDEX idx_txid_typeid ON public.messages USING btree (tx_id);


--
-- Name: txDenomFee; Type: INDEX; Schema: public; Owner: taxuser
--

CREATE UNIQUE INDEX "txDenomFee" ON public.fees USING btree (tx_id, denomination_id);


--
-- Name: blocks fk_blocks_chain; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.blocks
    ADD CONSTRAINT fk_blocks_chain FOREIGN KEY (blockchain_id) REFERENCES public.chains(id);


--
-- Name: denom_units fk_denom_units_denom; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.denom_units
    ADD CONSTRAINT fk_denom_units_denom FOREIGN KEY (denom_id) REFERENCES public.denoms(id);


--
-- Name: epoches fk_epoches_chain; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.epoches
    ADD CONSTRAINT fk_epoches_chain FOREIGN KEY (blockchain_id) REFERENCES public.chains(id);


--
-- Name: failed_blocks fk_failed_blocks_chain; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.failed_blocks
    ADD CONSTRAINT fk_failed_blocks_chain FOREIGN KEY (blockchain_id) REFERENCES public.chains(id);


--
-- Name: failed_event_blocks fk_failed_event_blocks_chain; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.failed_event_blocks
    ADD CONSTRAINT fk_failed_event_blocks_chain FOREIGN KEY (blockchain_id) REFERENCES public.chains(id);


--
-- Name: fees fk_fees_denomination; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.fees
    ADD CONSTRAINT fk_fees_denomination FOREIGN KEY (denomination_id) REFERENCES public.denoms(id);


--
-- Name: fees fk_fees_payer_address; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.fees
    ADD CONSTRAINT fk_fees_payer_address FOREIGN KEY (payer_address_id) REFERENCES public.addresses(id);


--
-- Name: messages fk_messages_message_type; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT fk_messages_message_type FOREIGN KEY (message_type_id) REFERENCES public.message_types(id);


--
-- Name: messages fk_messages_tx; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT fk_messages_tx FOREIGN KEY (tx_id) REFERENCES public.txes(id);


--
-- Name: taxable_event fk_taxable_event_block; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_event
    ADD CONSTRAINT fk_taxable_event_block FOREIGN KEY (block_id) REFERENCES public.blocks(id);


--
-- Name: taxable_event fk_taxable_event_denomination; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_event
    ADD CONSTRAINT fk_taxable_event_denomination FOREIGN KEY (denomination_id) REFERENCES public.denoms(id);


--
-- Name: taxable_event fk_taxable_event_event_address; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_event
    ADD CONSTRAINT fk_taxable_event_event_address FOREIGN KEY (address_id) REFERENCES public.addresses(id);


--
-- Name: taxable_tx fk_taxable_tx_denomination_received; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_tx
    ADD CONSTRAINT fk_taxable_tx_denomination_received FOREIGN KEY (denomination_received_id) REFERENCES public.denoms(id);


--
-- Name: taxable_tx fk_taxable_tx_denomination_sent; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_tx
    ADD CONSTRAINT fk_taxable_tx_denomination_sent FOREIGN KEY (denomination_sent_id) REFERENCES public.denoms(id);


--
-- Name: taxable_tx fk_taxable_tx_message; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_tx
    ADD CONSTRAINT fk_taxable_tx_message FOREIGN KEY (message_id) REFERENCES public.messages(id);


--
-- Name: taxable_tx fk_taxable_tx_receiver_address; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_tx
    ADD CONSTRAINT fk_taxable_tx_receiver_address FOREIGN KEY (receiver_address_id) REFERENCES public.addresses(id);


--
-- Name: taxable_tx fk_taxable_tx_sender_address; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.taxable_tx
    ADD CONSTRAINT fk_taxable_tx_sender_address FOREIGN KEY (sender_address_id) REFERENCES public.addresses(id);


--
-- Name: txes fk_txes_block; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.txes
    ADD CONSTRAINT fk_txes_block FOREIGN KEY (block_id) REFERENCES public.blocks(id);


--
-- Name: fees fk_txes_fees; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.fees
    ADD CONSTRAINT fk_txes_fees FOREIGN KEY (tx_id) REFERENCES public.txes(id);


--
-- Name: txes fk_txes_signer_address; Type: FK CONSTRAINT; Schema: public; Owner: taxuser
--

ALTER TABLE ONLY public.txes
    ADD CONSTRAINT fk_txes_signer_address FOREIGN KEY (signer_address_id) REFERENCES public.addresses(id);


--
-- PostgreSQL database dump complete
--
