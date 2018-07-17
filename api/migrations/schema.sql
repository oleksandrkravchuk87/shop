--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.12
-- Dumped by pg_dump version 9.5.12

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: alert_severity; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alert_severity (
    pcd character varying(50) NOT NULL,
    name character varying(150),
    "order" integer,
    icon character varying(150)
);


ALTER TABLE public.alert_severity OWNER TO postgres;

--
-- Name: cars; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cars (
    id uuid NOT NULL,
    name text NOT NULL,
    color text NOT NULL,
    wheels integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.cars OWNER TO postgres;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id uuid NOT NULL,
    alias character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    "descr" character varying(255) NOT NULL,
    logo character varying(255) NOT NULL,
    parent_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: dropdown; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dropdown (
    pcd character varying(50) NOT NULL,
    behaviour integer,
    scorecard_id integer,
    value_pcd character varying(50) NOT NULL,
    value_cd character varying(150),
    value_engine_id integer,
    value_name character varying(150),
    value_is_default character varying(1) DEFAULT 'N'::character varying,
    value_format text,
    value_order integer
);


ALTER TABLE public.dropdown OWNER TO postgres;

--
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id uuid NOT NULL,
    alias character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    "descr" character varying(255) NOT NULL,
    pictures character varying(255) NOT NULL,
    price integer NOT NULL,
    count integer NOT NULL,
    category_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.items OWNER TO postgres;

--
-- Name: migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.migrations (
    id character varying(255) NOT NULL
);


ALTER TABLE public.migrations OWNER TO postgres;

--
-- Name: mongodbs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mongodbs (
    domain text NOT NULL,
    mongodb boolean,
    host text,
    port text
);


ALTER TABLE public.mongodbs OWNER TO postgres;

--
-- Name: ordereds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ordereds (
    id integer NOT NULL,
    order_id integer NOT NULL,
    item_id uuid NOT NULL,
    item_cnt integer NOT NULL,
    item_sum integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.ordereds OWNER TO postgres;

--
-- Name: ordereds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ordereds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ordereds_id_seq OWNER TO postgres;

--
-- Name: ordereds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ordereds_id_seq OWNED BY public.ordereds.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id integer NOT NULL,
    status character varying(255) NOT NULL,
    sum integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_id_seq OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: rule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.rule (
    pcd character varying(50) NOT NULL,
    name character varying(150),
    "order" integer,
    type_id integer,
    arche_type_id integer,
    interval_id integer,
    interval_generic_metric character varying(50),
    condition_reference_id integer,
    condition_reference character varying(50),
    condition_range_source_id integer,
    condition_from character varying(50),
    condition_to character varying(50),
    condition_transformation_id integer,
    condition_transformation character varying(50),
    threshold_source_id integer,
    threshold_type_id integer,
    threshold_min character varying(50),
    threshold_max character varying(50),
    threshold_transformation_id integer,
    rule_interval_fallback integer,
    rule_default_process_order integer DEFAULT 999
);


ALTER TABLE public.rule OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: tempconfigs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tempconfigs (
    rest_api_root text NOT NULL,
    host text,
    port text,
    remoting text,
    legasy_explorer boolean
);


ALTER TABLE public.tempconfigs OWNER TO postgres;

--
-- Name: tsconfigs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tsconfigs (
    module text NOT NULL,
    target text,
    source_map boolean,
    excluding integer
);


ALTER TABLE public.tsconfigs OWNER TO postgres;

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ordereds ALTER COLUMN id SET DEFAULT nextval('public.ordereds_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: alert_severity_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_severity
    ADD CONSTRAINT alert_severity_pkey PRIMARY KEY (pcd);


--
-- Name: cars_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cars
    ADD CONSTRAINT cars_pkey PRIMARY KEY (id);


--
-- Name: categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: dropdown_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dropdown
    ADD CONSTRAINT dropdown_pkey PRIMARY KEY (pcd, value_pcd);


--
-- Name: items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);


--
-- Name: mongodbs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mongodbs
    ADD CONSTRAINT mongodbs_pkey PRIMARY KEY (domain);


--
-- Name: ordereds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ordereds
    ADD CONSTRAINT ordereds_pkey PRIMARY KEY (id);


--
-- Name: orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: rule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rule
    ADD CONSTRAINT rule_pkey PRIMARY KEY (pcd);


--
-- Name: tempconfigs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tempconfigs
    ADD CONSTRAINT tempconfigs_pkey PRIMARY KEY (rest_api_root);


--
-- Name: tsconfigs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tsconfigs
    ADD CONSTRAINT tsconfigs_pkey PRIMARY KEY (module);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

