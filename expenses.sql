--
-- PostgreSQL database dump
--

-- Dumped from database version 14.7 (Ubuntu 14.7-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.7 (Ubuntu 14.7-0ubuntu0.22.04.1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cat (
    id integer NOT NULL,
    name character varying(15) NOT NULL
);


ALTER TABLE public.cat OWNER TO postgres;

--
-- Name: cat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.cat ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.cat_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: city; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.city (
    id integer NOT NULL,
    city character varying(15) NOT NULL
);


ALTER TABLE public.city OWNER TO postgres;

--
-- Name: city_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.city ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.city_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: expense; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.expense (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    subcat_id integer NOT NULL,
    nds integer
);


ALTER TABLE public.expense OWNER TO postgres;

--
-- Name: expense_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.expense ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.expense_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: income; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.income (
    year integer NOT NULL,
    month integer NOT NULL,
    income_amount integer NOT NULL
);


ALTER TABLE public.income OWNER TO postgres;

--
-- Name: market_or_supplier; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.market_or_supplier (
    name character varying(50) NOT NULL,
    address character varying(100) DEFAULT ''::character varying NOT NULL,
    id integer NOT NULL
);


ALTER TABLE public.market_or_supplier OWNER TO postgres;

--
-- Name: market_or_supplier_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.market_or_supplier_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.market_or_supplier_id_seq OWNER TO postgres;

--
-- Name: market_or_supplier_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.market_or_supplier_id_seq OWNED BY public.market_or_supplier.id;


--
-- Name: purchase; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.purchase (
    id integer NOT NULL,
    purchase_date date NOT NULL,
    city_id integer NOT NULL,
    online boolean NOT NULL,
    description text,
    mos_id integer
);


ALTER TABLE public.purchase OWNER TO postgres;

--
-- Name: purchase_check; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.purchase_check (
    id integer NOT NULL,
    expense_id integer NOT NULL,
    purchase_id integer NOT NULL,
    count real NOT NULL,
    price integer NOT NULL
);


ALTER TABLE public.purchase_check OWNER TO postgres;

--
-- Name: purchase_check_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.purchase_check ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.purchase_check_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: purchase_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.purchase ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.purchase_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: reply; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reply (
    id integer NOT NULL,
    rate integer NOT NULL,
    description text,
    mos_id integer,
    expense_id integer,
    CONSTRAINT reply_rate_check CHECK (((rate > 0) AND (rate <= 10)))
);


ALTER TABLE public.reply OWNER TO postgres;

--
-- Name: reply_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.reply ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.reply_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: subcat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.subcat (
    id integer NOT NULL,
    name character varying(35) NOT NULL,
    cat_id integer NOT NULL
);


ALTER TABLE public.subcat OWNER TO postgres;

--
-- Name: subcat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.subcat ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.subcat_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: market_or_supplier id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.market_or_supplier ALTER COLUMN id SET DEFAULT nextval('public.market_or_supplier_id_seq'::regclass);


--
-- Name: cat cat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cat
    ADD CONSTRAINT cat_pkey PRIMARY KEY (id);


--
-- Name: city city_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.city
    ADD CONSTRAINT city_pkey PRIMARY KEY (id);


--
-- Name: expense expense_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.expense
    ADD CONSTRAINT expense_pkey PRIMARY KEY (id);


--
-- Name: income income_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.income
    ADD CONSTRAINT income_pkey PRIMARY KEY (year, month);


--
-- Name: market_or_supplier market_or_supplier_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.market_or_supplier
    ADD CONSTRAINT market_or_supplier_pk UNIQUE (name, address);


--
-- Name: purchase_check purchase_check_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchase_check
    ADD CONSTRAINT purchase_check_pkey PRIMARY KEY (id);


--
-- Name: purchase purchase_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchase
    ADD CONSTRAINT purchase_pkey PRIMARY KEY (id);


--
-- Name: reply reply_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reply
    ADD CONSTRAINT reply_pkey PRIMARY KEY (id);


--
-- Name: subcat subcat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subcat
    ADD CONSTRAINT subcat_pkey PRIMARY KEY (id);


--
-- Name: market_or_supplier_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX market_or_supplier_id_uindex ON public.market_or_supplier USING btree (id);


--
-- Name: expense expense_subcat_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.expense
    ADD CONSTRAINT expense_subcat_id_fkey FOREIGN KEY (subcat_id) REFERENCES public.subcat(id);


--
-- Name: purchase_check purchase_check_expense_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchase_check
    ADD CONSTRAINT purchase_check_expense_id_fkey FOREIGN KEY (expense_id) REFERENCES public.expense(id);


--
-- Name: purchase_check purchase_check_purchase_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchase_check
    ADD CONSTRAINT purchase_check_purchase_id_fkey FOREIGN KEY (purchase_id) REFERENCES public.purchase(id);


--
-- Name: purchase purchase_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchase
    ADD CONSTRAINT purchase_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.city(id);


--
-- Name: purchase purchase_market_or_supplier_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.purchase
    ADD CONSTRAINT purchase_market_or_supplier_id_fk FOREIGN KEY (mos_id) REFERENCES public.market_or_supplier(id);


--
-- Name: reply reply_expense_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reply
    ADD CONSTRAINT reply_expense_id_fk FOREIGN KEY (expense_id) REFERENCES public.expense(id);


--
-- Name: reply reply_market_or_supplier_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reply
    ADD CONSTRAINT reply_market_or_supplier_id_fk FOREIGN KEY (mos_id) REFERENCES public.market_or_supplier(id);


--
-- Name: subcat subcat_cat_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subcat
    ADD CONSTRAINT subcat_cat_id_fkey FOREIGN KEY (cat_id) REFERENCES public.cat(id);


--
-- PostgreSQL database dump complete
--

