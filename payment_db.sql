--
-- PostgreSQL database dump
--

\restrict bn0WuARXAXmYSUaT9FcwC5f5Y2Dc1HUWEgqx0euXXc0xahiHbMtgZBe1pLl6Vhx

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: list_payment_waiting; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.list_payment_waiting (
    id bigint NOT NULL,
    payment_id bigint NOT NULL,
    order_id bigint NOT NULL,
    user_id bigint NOT NULL,
    amount numeric(15,2) NOT NULL,
    order_code text NOT NULL,
    icon_method_payment text NOT NULL,
    number_payment text,
    code_qris text,
    checkout_at timestamp with time zone,
    expired_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.list_payment_waiting OWNER TO postgres;

--
-- Name: list_payment_waiting_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.list_payment_waiting_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.list_payment_waiting_id_seq OWNER TO postgres;

--
-- Name: list_payment_waiting_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.list_payment_waiting_id_seq OWNED BY public.list_payment_waiting.id;


--
-- Name: payment_codes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_codes (
    id bigint NOT NULL,
    code text NOT NULL,
    payment_id bigint NOT NULL,
    expired_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.payment_codes OWNER TO postgres;

--
-- Name: payment_codes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payment_codes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_codes_id_seq OWNER TO postgres;

--
-- Name: payment_codes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payment_codes_id_seq OWNED BY public.payment_codes.id;


--
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_methods (
    id bigint NOT NULL,
    payment_method character varying(50) NOT NULL,
    number_payment text,
    url_code text,
    url_icon text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.payment_methods OWNER TO postgres;

--
-- Name: payment_methods_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payment_methods_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_methods_id_seq OWNER TO postgres;

--
-- Name: payment_methods_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payment_methods_id_seq OWNED BY public.payment_methods.id;


--
-- Name: payment_proofs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_proofs (
    id bigint CONSTRAINT proof_payments_id_not_null NOT NULL,
    payment_id bigint CONSTRAINT proof_payments_payment_id_not_null NOT NULL,
    proof_url text CONSTRAINT proof_payments_proof_url_not_null NOT NULL,
    note text,
    uploaded_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    verified_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.payment_proofs OWNER TO postgres;

--
-- Name: payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payments (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    order_id bigint NOT NULL,
    amount numeric(15,2) NOT NULL,
    payment_method character varying(50),
    status character varying(50) DEFAULT 'Pending'::character varying,
    paid_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    approved_at timestamp with time zone
);


ALTER TABLE public.payments OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payments_id_seq OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


--
-- Name: proof_payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.proof_payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.proof_payments_id_seq OWNER TO postgres;

--
-- Name: proof_payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.proof_payments_id_seq OWNED BY public.payment_proofs.id;


--
-- Name: refund_proofs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.refund_proofs (
    id bigint NOT NULL,
    refund_id bigint NOT NULL,
    file_url text NOT NULL,
    note text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.refund_proofs OWNER TO postgres;

--
-- Name: refund_proofs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.refund_proofs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.refund_proofs_id_seq OWNER TO postgres;

--
-- Name: refund_proofs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.refund_proofs_id_seq OWNED BY public.refund_proofs.id;


--
-- Name: refunds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.refunds (
    id bigint NOT NULL,
    rejected_id bigint NOT NULL,
    bank_name character varying(100) NOT NULL,
    account_number character varying(100) NOT NULL,
    account_name character varying(150) NOT NULL,
    status character varying(50) DEFAULT 'requested'::character varying NOT NULL,
    transferred_at timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.refunds OWNER TO postgres;

--
-- Name: refunds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.refunds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.refunds_id_seq OWNER TO postgres;

--
-- Name: refunds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.refunds_id_seq OWNED BY public.refunds.id;


--
-- Name: rejected_payment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.rejected_payment (
    id bigint NOT NULL,
    payment_id bigint NOT NULL,
    user_id bigint NOT NULL,
    amount bigint NOT NULL,
    order_code character varying(50) NOT NULL,
    payment_code character varying(50) NOT NULL,
    order_name character varying(200) NOT NULL,
    admin_note text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.rejected_payment OWNER TO postgres;

--
-- Name: rejected_payment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.rejected_payment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.rejected_payment_id_seq OWNER TO postgres;

--
-- Name: rejected_payment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.rejected_payment_id_seq OWNED BY public.rejected_payment.id;


--
-- Name: list_payment_waiting id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.list_payment_waiting ALTER COLUMN id SET DEFAULT nextval('public.list_payment_waiting_id_seq'::regclass);


--
-- Name: payment_codes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_codes ALTER COLUMN id SET DEFAULT nextval('public.payment_codes_id_seq'::regclass);


--
-- Name: payment_methods id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN id SET DEFAULT nextval('public.payment_methods_id_seq'::regclass);


--
-- Name: payment_proofs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_proofs ALTER COLUMN id SET DEFAULT nextval('public.proof_payments_id_seq'::regclass);


--
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- Name: refund_proofs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refund_proofs ALTER COLUMN id SET DEFAULT nextval('public.refund_proofs_id_seq'::regclass);


--
-- Name: refunds id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refunds ALTER COLUMN id SET DEFAULT nextval('public.refunds_id_seq'::regclass);


--
-- Name: rejected_payment id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rejected_payment ALTER COLUMN id SET DEFAULT nextval('public.rejected_payment_id_seq'::regclass);


--
-- Data for Name: list_payment_waiting; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.list_payment_waiting (id, payment_id, order_id, user_id, amount, order_code, icon_method_payment, number_payment, code_qris, checkout_at, expired_at, created_at) FROM stdin;
1	8	9	6	2600.00	payment_w5uW5z	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-02-12 17:35:34.392565+07	2026-02-12 19:35:34.324923+07	2026-02-12 17:35:34.392565+07
2	9	10	1	2600.00	payment_XlLc18	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-02-15 23:43:25.089108+07	2026-02-16 01:43:25.075782+07	2026-02-15 23:43:25.089108+07
3	10	11	1	2600.00	payment_9cWznC	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-02-16 22:02:57.743718+07	2026-02-17 00:02:57.722682+07	2026-02-16 22:02:57.743718+07
4	11	13	1	6000.00	payment_VevACb	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-02-24 06:56:02.712892+07	2026-02-24 08:56:02.691909+07	2026-02-24 06:56:02.712892+07
5	12	15	1	6000.00	payment_LEbxxP	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-07 23:32:47.119718+07	2026-03-08 01:32:47.109182+07	2026-03-07 23:32:47.119718+07
6	13	16	1	6000.00	payment_LG2zPN	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-08 02:46:56.471868+07	2026-03-08 04:46:56.467533+07	2026-03-08 02:46:56.471868+07
7	14	17	1	6000.00	payment_iYP2JC	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-08 03:27:27.095372+07	2026-03-08 05:27:27.089987+07	2026-03-08 03:27:27.095372+07
8	15	18	1	6000.00	payment_eQ8DU2	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-08 03:46:07.634853+07	2026-03-08 05:46:07.619277+07	2026-03-08 03:46:07.634853+07
9	16	19	1	6000.00	payment_Re3qMQ	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-08 06:33:42.453319+07	2026-03-08 08:33:42.423627+07	2026-03-08 06:33:42.453319+07
10	17	20	1	6000.00	payment_YM4tp0	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-08 07:43:11.201045+07	2026-03-08 09:43:11.171705+07	2026-03-08 07:43:11.201045+07
11	18	21	1	6000.00	payment_hFk6fv	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-08 18:29:37.038648+07	2026-03-08 20:29:37.020376+07	2026-03-08 18:29:37.038648+07
12	19	22	1	6000.00	payment_LlPyQO	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-08 18:48:27.107368+07	2026-03-08 20:48:27.091797+07	2026-03-08 18:48:27.107368+07
13	20	23	1	6000.00	payment_NVN7V2	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-09 03:09:35.104848+07	2026-03-09 05:09:35.096152+07	2026-03-09 03:09:35.104848+07
14	21	24	1	6000.00	payment_VNqMek	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-09 03:58:56.464781+07	2026-03-09 05:58:56.457154+07	2026-03-09 03:58:56.464781+07
15	22	29	1	5000.00	payment_6zeuiP	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-16 01:35:05.518638+07	2026-03-16 03:35:05.504759+07	2026-03-16 01:35:05.518638+07
16	23	30	1	6000.00	payment_1hOkFu	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-03-16 03:20:39.871696+07	2026-03-16 05:20:39.856456+07	2026-03-16 03:20:39.871696+07
17	24	31	1	5200.00	payment_RzPGQx	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-16 17:14:14.481687+07	2026-03-16 19:14:14.448918+07	2026-03-16 17:14:14.481687+07
18	25	33	6	2600.00	payment_rDhCMG	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-17 01:28:51.369896+07	2026-03-17 03:28:51.344125+07	2026-03-17 01:28:51.369896+07
19	26	34	8	1250.00	payment_jnypz6	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-17 01:37:52.208067+07	2026-03-17 03:37:52.204069+07	2026-03-17 01:37:52.208067+07
20	27	35	8	3200.00	payment_TUtn9l	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-17 02:08:18.022634+07	2026-03-17 04:08:18.007139+07	2026-03-17 02:08:18.022634+07
21	28	37	8	6600.00	payment_r25Dz4	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-17 08:59:21.72147+07	2026-03-17 10:59:21.710242+07	2026-03-17 08:59:21.72147+07
22	29	38	6	5000.00	payment_JbwbDW	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-24 10:15:46.987469+07	2026-03-24 12:15:46.97421+07	2026-03-24 10:15:46.987469+07
23	30	39	6	6000.00	payment_v7E6md	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-24 11:25:50.815934+07	2026-03-24 13:25:50.802012+07	2026-03-24 11:25:50.815934+07
24	31	43	8	6000.00	payment_QcfQHU	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-03-26 08:03:44.394348+07	2026-03-26 10:03:44.381897+07	2026-03-26 08:03:44.394348+07
25	32	46	8	5000.00	payment_Ru24yk	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-27 06:05:49.721655+07	2026-03-27 08:05:49.713177+07	2026-03-27 06:05:49.721655+07
26	33	47	8	5000.00	payment_XjCnzc	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-27 07:16:02.269325+07	2026-03-27 09:16:02.262705+07	2026-03-27 07:16:02.269325+07
27	34	48	1	10000.00	payment_PdWB3G	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-03-27 10:05:43.374526+07	2026-03-27 12:05:43.365715+07	2026-03-27 10:05:43.374526+07
28	35	49	8	2200.00	payment_2Flsqv	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-03-28 09:40:15.988145+07	2026-03-28 11:40:15.98237+07	2026-03-28 09:40:15.988145+07
29	36	50	1	48000.00	payment_6bStzp	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-03-31 16:59:48.885942+07	2026-03-31 18:59:48.870226+07	2026-03-31 16:59:48.885942+07
30	37	51	1	2500.00	payment_hya6Ir	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-04-01 11:08:27.963941+07	2026-04-01 13:08:27.944354+07	2026-04-01 11:08:27.963941+07
31	38	52	8	5000.00	payment_ngzJqm	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-04-08 10:08:58.876119+07	2026-04-08 12:08:58.864038+07	2026-04-08 10:08:58.876119+07
32	39	53	1	7800.00	payment_mDD5gT	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-04-11 17:43:59.270524+07	2026-04-11 19:43:59.235844+07	2026-04-11 17:43:59.270524+07
33	40	54	1	5000.00	payment_2LBXAZ	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-04-16 08:05:27.766787+07	2026-04-16 10:05:27.746035+07	2026-04-16 08:05:27.766787+07
34	41	55	1	2500.00	payment_uWmhON	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-04-22 23:00:55.113992+07	2026-04-23 01:00:55.103672+07	2026-04-22 23:00:55.113992+07
35	42	56	1	6000.00	payment_4c4GHb	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-04-23 20:03:11.705362+07	2026-04-23 22:03:11.696825+07	2026-04-23 20:03:11.705362+07
36	43	57	1	5000.00	payment_njIvjU	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-04-24 20:26:17.780531+07	2026-04-24 22:26:17.767936+07	2026-04-24 20:26:17.780531+07
37	44	58	1	3200.00	payment_oa4G2k	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-04-25 19:18:32.735909+07	2026-04-25 21:18:32.725819+07	2026-04-25 19:18:32.735909+07
38	45	59	8	1600.00	payment_SRC4sY	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-04-26 14:35:28.124256+07	2026-04-26 16:35:28.11334+07	2026-04-26 14:35:28.124256+07
39	46	60	8	10000.00	payment_LUJPPq	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-04-26 14:56:43.137873+07	2026-04-26 16:56:43.130194+07	2026-04-26 14:56:43.137873+07
40	47	62	15	10000.00	payment_Pz757i	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-04-26 15:03:29.32265+07	2026-04-26 17:03:29.317055+07	2026-04-26 15:03:29.32265+07
41	48	61	8	2000.00	payment_iHRjZo	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-04-26 15:03:57.710799+07	2026-04-26 17:03:57.70621+07	2026-04-26 15:03:57.710799+07
42	49	66	15	5000.00	payment_Z9ApNZ	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-04-26 16:07:05.431111+07	2026-04-26 18:07:05.422836+07	2026-04-26 16:07:05.431111+07
43	50	67	1	2600.00	payment_ovDjOV	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-04-26 21:17:39.119596+07	2026-04-26 23:17:39.089275+07	2026-04-26 21:17:39.119596+07
44	51	68	8	11000.00	payment_BEtDCX	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-04-26 21:36:22.205414+07	2026-04-26 23:36:22.19542+07	2026-04-26 21:36:22.205414+07
45	52	69	8	5000.00	payment_ImbL7i	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-04-27 14:29:24.612026+07	2026-04-27 16:29:24.602724+07	2026-04-27 14:29:24.612026+07
46	53	70	15	30000.00	payment_AKi64K	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-04-29 15:58:51.650432+07	2026-04-29 17:58:51.633921+07	2026-04-29 15:58:51.650432+07
47	54	71	1	2100.00	payment_1S9G6D	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-04-29 21:27:43.859056+07	2026-04-29 23:27:43.835736+07	2026-04-29 21:27:43.859056+07
48	55	72	1	2200.00	payment_tB1cGX	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-04-30 23:41:37.33887+07	2026-05-01 01:41:37.328234+07	2026-04-30 23:41:37.33887+07
49	56	78	15	180000.00	payment_Het2dr	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-05-04 09:06:30.098699+07	2026-05-04 11:06:30.053888+07	2026-05-04 09:06:30.098699+07
50	57	79	8	200000.00	payment_TYOmbO	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-05-04 15:21:29.846628+07	2026-05-04 17:21:29.839052+07	2026-05-04 15:21:29.846628+07
51	58	80	8	200000.00	payment_53RrrS	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-05-04 23:01:16.475877+07	2026-05-05 01:01:16.459193+07	2026-05-04 23:01:16.475877+07
52	59	81	15	150000.00	payment_8HcJ70	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-05-05 12:49:52.576047+07	2026-05-05 14:49:52.536766+07	2026-05-05 12:49:52.576047+07
53	60	83	15	5200.00	payment_2JXSCM	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-05-06 20:33:50.039031+07	2026-05-06 22:33:50.030152+07	2026-05-06 20:33:50.039031+07
54	61	85	16	2000.00	payment_IFtRiT	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-05-06 23:20:38.029999+07	2026-05-07 01:20:37.924086+07	2026-05-06 23:20:38.029999+07
55	62	84	15	6000.00	payment_ANWdvi	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-05-06 23:21:09.742722+07	2026-05-07 01:21:09.72621+07	2026-05-06 23:21:09.742722+07
56	63	86	16	13200.00	payment_ixoCOO	localhost:8083/static/icon-method/1768797275722968700.svg		localhost:8083/static/code-qris/1768797275724963100.png	2026-05-06 23:23:59.002122+07	2026-05-07 01:23:58.901649+07	2026-05-06 23:23:59.002122+07
57	64	87	15	200000.00	payment_t6e3oP	localhost:8083/static/icon-method/1768669366805942600.svg	1238192748334234		2026-05-11 21:52:52.961447+07	2026-05-11 23:52:52.943788+07	2026-05-11 21:52:52.961447+07
58	65	88	16	22000.00	payment_9XoBHk	localhost:8083/static/icon-method/1768760144410657100.svg	089692612004		2026-05-12 11:45:07.892219+07	2026-05-12 13:45:07.873448+07	2026-05-12 11:45:07.892219+07
59	66	89	16	5000.00	payment_aKx636	localhost:8083/static/icon-method/1778918317928483000.svg	089692612004		2026-05-16 15:23:35.325963+07	2026-05-16 17:23:35.311593+07	2026-05-16 15:23:35.325963+07
60	67	90	1	5000.00	payment_QPx6e4	localhost:8083/static/icon-method/1778918457997000500.png	089692612004		2026-05-19 15:15:24.482193+07	2026-05-19 17:15:24.469939+07	2026-05-19 15:15:24.482193+07
61	68	92	16	5000.00	payment_ZaxE4J	localhost:8083/static/icon-method/1778918317928483000.svg	089692612004		2026-05-19 21:49:37.791762+07	2026-05-19 23:49:37.779862+07	2026-05-19 21:49:37.791762+07
62	69	93	1	25000.00	payment_1gTaMp	localhost:8083/static/icon-method/1778918457997000500.png	089692612004		2026-05-21 14:33:48.273298+07	2026-05-21 16:33:48.262348+07	2026-05-21 14:33:48.273298+07
63	70	94	8	6000.00	payment_yOOrgP	localhost:8083/static/icon-method/1778918317928483000.svg	089692612004		2026-05-23 12:26:29.498476+07	2026-05-23 14:26:29.470964+07	2026-05-23 12:26:29.498476+07
64	71	95	1	5000.00	payment_haq45q	localhost:8083/static/icon-method/1778918457997000500.png	089692612004		2026-05-23 12:29:36.28572+07	2026-05-23 14:29:36.27904+07	2026-05-23 12:29:36.28572+07
65	72	96	1	2000.00	payment_Wlu7St	localhost:8083/static/icon-method/1778918457997000500.png	089692612004		2026-05-23 18:56:36.033143+07	2026-05-23 20:56:36.02634+07	2026-05-23 18:56:36.033143+07
66	73	97	6	1600.00	payment_oVOeGV	localhost:8083/static/icon-method/1778918317928483000.svg	089692612004		2026-06-11 10:54:09.613823+07	2026-06-11 12:54:09.575354+07	2026-06-11 10:54:09.613823+07
67	74	98	15	5000.00	payment_8zTIqd	localhost:8083/static/icon-method/1778918317928483000.svg	089692612004		2026-06-20 19:56:18.457112+07	2026-06-20 21:56:18.438855+07	2026-06-20 19:56:18.457112+07
68	75	99	15	170000.00	payment_xLoGFB	localhost:8083/static/icon-method/1778918317928483000.svg	089692612004		2026-06-20 20:07:30.138603+07	2026-06-20 22:07:30.13075+07	2026-06-20 20:07:30.138603+07
\.


--
-- Data for Name: payment_codes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_codes (id, code, payment_id, expired_at, created_at) FROM stdin;
2	payment_DXVjNW	2	2026-02-11 02:06:18.089084+07	2026-02-11 00:06:18.09608+07
3	payment_uGzPhT	3	2026-02-11 16:23:04.489421+07	2026-02-11 14:23:04.489421+07
5	payment_h7QFvg	5	2026-02-12 03:01:29.838022+07	2026-02-12 01:01:29.83879+07
6	payment_h2v372	6	2026-02-12 03:24:42.993728+07	2026-02-12 01:24:42.993728+07
7	payment_MzB5RE	7	2026-02-12 17:16:09.920863+07	2026-02-12 15:16:09.921374+07
8	payment_w5uW5z	8	2026-02-12 19:35:34.324923+07	2026-02-12 17:35:34.324923+07
9	payment_XlLc18	9	2026-02-16 01:43:25.075782+07	2026-02-15 23:43:25.076707+07
10	payment_9cWznC	10	2026-02-17 00:02:57.722682+07	2026-02-16 22:02:57.725199+07
11	payment_VevACb	11	2026-02-24 08:56:02.691909+07	2026-02-24 06:56:02.691909+07
12	payment_LEbxxP	12	2026-03-08 01:32:47.109182+07	2026-03-07 23:32:47.109182+07
13	payment_LG2zPN	13	2026-03-08 04:46:56.467533+07	2026-03-08 02:46:56.467533+07
14	payment_iYP2JC	14	2026-03-08 05:27:27.089987+07	2026-03-08 03:27:27.089987+07
15	payment_eQ8DU2	15	2026-03-08 05:46:07.619277+07	2026-03-08 03:46:07.619787+07
16	payment_Re3qMQ	16	2026-03-08 08:33:42.423627+07	2026-03-08 06:33:42.424577+07
17	payment_YM4tp0	17	2026-03-08 09:43:11.171705+07	2026-03-08 07:43:11.171705+07
18	payment_hFk6fv	18	2026-03-08 20:29:37.020376+07	2026-03-08 18:29:37.020885+07
19	payment_LlPyQO	19	2026-03-08 20:48:27.091797+07	2026-03-08 18:48:27.092797+07
20	payment_NVN7V2	20	2026-03-09 05:09:35.096152+07	2026-03-09 03:09:35.096842+07
21	payment_VNqMek	21	2026-03-09 05:58:56.457154+07	2026-03-09 03:58:56.458028+07
22	payment_6zeuiP	22	2026-03-16 03:35:05.504759+07	2026-03-16 01:35:05.504759+07
23	payment_1hOkFu	23	2026-03-16 05:20:39.856456+07	2026-03-16 03:20:39.856456+07
24	payment_RzPGQx	24	2026-03-16 19:14:14.448918+07	2026-03-16 17:14:14.44943+07
25	payment_rDhCMG	25	2026-03-17 03:28:51.344125+07	2026-03-17 01:28:51.345982+07
26	payment_jnypz6	26	2026-03-17 03:37:52.204069+07	2026-03-17 01:37:52.204069+07
27	payment_TUtn9l	27	2026-03-17 04:08:18.007139+07	2026-03-17 02:08:18.012744+07
28	payment_r25Dz4	28	2026-03-17 10:59:21.710242+07	2026-03-17 08:59:21.710242+07
29	payment_JbwbDW	29	2026-03-24 12:15:46.97421+07	2026-03-24 10:15:46.97472+07
30	payment_v7E6md	30	2026-03-24 13:25:50.802012+07	2026-03-24 11:25:50.802012+07
31	payment_QcfQHU	31	2026-03-26 10:03:44.381897+07	2026-03-26 08:03:44.381897+07
32	payment_Ru24yk	32	2026-03-27 08:05:49.713177+07	2026-03-27 06:05:49.713862+07
33	payment_XjCnzc	33	2026-03-27 09:16:02.262705+07	2026-03-27 07:16:02.262705+07
34	payment_PdWB3G	34	2026-03-27 12:05:43.365715+07	2026-03-27 10:05:43.365715+07
35	payment_2Flsqv	35	2026-03-28 11:40:15.98237+07	2026-03-28 09:40:15.98298+07
36	payment_6bStzp	36	2026-03-31 18:59:48.870226+07	2026-03-31 16:59:48.870812+07
37	payment_hya6Ir	37	2026-04-01 13:08:27.944354+07	2026-04-01 11:08:27.944354+07
38	payment_ngzJqm	38	2026-04-08 12:08:58.864038+07	2026-04-08 10:08:58.865099+07
39	payment_mDD5gT	39	2026-04-11 19:43:59.235844+07	2026-04-11 17:43:59.235844+07
40	payment_2LBXAZ	40	2026-04-16 10:05:27.746035+07	2026-04-16 08:05:27.746551+07
41	payment_uWmhON	41	2026-04-23 01:00:55.103672+07	2026-04-22 23:00:55.10568+07
42	payment_4c4GHb	42	2026-04-23 22:03:11.696825+07	2026-04-23 20:03:11.696825+07
43	payment_njIvjU	43	2026-04-24 22:26:17.767936+07	2026-04-24 20:26:17.768688+07
44	payment_oa4G2k	44	2026-04-25 21:18:32.725819+07	2026-04-25 19:18:32.726404+07
45	payment_SRC4sY	45	2026-04-26 16:35:28.11334+07	2026-04-26 14:35:28.11334+07
46	payment_LUJPPq	46	2026-04-26 16:56:43.130194+07	2026-04-26 14:56:43.130194+07
47	payment_Pz757i	47	2026-04-26 17:03:29.317055+07	2026-04-26 15:03:29.317055+07
48	payment_iHRjZo	48	2026-04-26 17:03:57.70621+07	2026-04-26 15:03:57.707247+07
49	payment_Z9ApNZ	49	2026-04-26 18:07:05.422836+07	2026-04-26 16:07:05.422836+07
50	payment_ovDjOV	50	2026-04-26 23:17:39.089275+07	2026-04-26 21:17:39.089275+07
51	payment_BEtDCX	51	2026-04-26 23:36:22.19542+07	2026-04-26 21:36:22.19542+07
52	payment_ImbL7i	52	2026-04-27 16:29:24.602724+07	2026-04-27 14:29:24.603235+07
53	payment_AKi64K	53	2026-04-29 17:58:51.633921+07	2026-04-29 15:58:51.63464+07
54	payment_1S9G6D	54	2026-04-29 23:27:43.835736+07	2026-04-29 21:27:43.835736+07
55	payment_tB1cGX	55	2026-05-01 01:41:37.328234+07	2026-04-30 23:41:37.32875+07
56	payment_Het2dr	56	2026-05-04 11:06:30.053888+07	2026-05-04 09:06:30.053888+07
57	payment_TYOmbO	57	2026-05-04 17:21:29.839052+07	2026-05-04 15:21:29.839569+07
58	payment_53RrrS	58	2026-05-05 01:01:16.459193+07	2026-05-04 23:01:16.459703+07
59	payment_8HcJ70	59	2026-05-05 14:49:52.536766+07	2026-05-05 12:49:52.537278+07
60	payment_2JXSCM	60	2026-05-06 22:33:50.030152+07	2026-05-06 20:33:50.030721+07
61	payment_IFtRiT	61	2026-05-07 01:20:37.924086+07	2026-05-06 23:20:37.92477+07
62	payment_ANWdvi	62	2026-05-07 01:21:09.72621+07	2026-05-06 23:21:09.72621+07
63	payment_ixoCOO	63	2026-05-07 01:23:58.901649+07	2026-05-06 23:23:58.901649+07
64	payment_t6e3oP	64	2026-05-11 23:52:52.943788+07	2026-05-11 21:52:52.944389+07
65	payment_9XoBHk	65	2026-05-12 13:45:07.873448+07	2026-05-12 11:45:07.874449+07
66	payment_aKx636	66	2026-05-16 17:23:35.311593+07	2026-05-16 15:23:35.311593+07
67	payment_QPx6e4	67	2026-05-19 17:15:24.469939+07	2026-05-19 15:15:24.470937+07
68	payment_ZaxE4J	68	2026-05-19 23:49:37.779862+07	2026-05-19 21:49:37.779862+07
69	payment_1gTaMp	69	2026-05-21 16:33:48.262348+07	2026-05-21 14:33:48.262874+07
70	payment_yOOrgP	70	2026-05-23 14:26:29.470964+07	2026-05-23 12:26:29.470964+07
71	payment_haq45q	71	2026-05-23 14:29:36.27904+07	2026-05-23 12:29:36.27904+07
72	payment_Wlu7St	72	2026-05-23 20:56:36.02634+07	2026-05-23 18:56:36.026914+07
73	payment_oVOeGV	73	2026-06-11 12:54:09.575354+07	2026-06-11 10:54:09.575354+07
74	payment_8zTIqd	74	2026-06-20 21:56:18.438855+07	2026-06-20 19:56:18.438855+07
75	payment_xLoGFB	75	2026-06-20 22:07:30.13075+07	2026-06-20 20:07:30.131258+07
\.


--
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_methods (id, payment_method, number_payment, url_code, url_icon, created_at, updated_at) FROM stdin;
7	BRI	1238192748334234		localhost:8083/static/icon-method/1768669366805942600.svg	2026-01-18 00:02:46.809518+07	2026-01-18 00:02:46.809518+07
9	Dana	089692612004		localhost:8083/static/icon-method/1768760144410657100.svg	2026-01-19 01:15:44.418089+07	2026-01-19 01:15:44.418089+07
10	Qris		localhost:8083/static/code-qris/1768797275724963100.png	localhost:8083/static/icon-method/1768797275722968700.svg	2026-01-19 11:34:35.72906+07	2026-01-19 11:34:35.72906+07
12	Gopay	089692612004		localhost:8083/static/icon-method/1778918317928483000.svg	2026-05-16 14:58:37.932775+07	2026-05-16 14:58:37.932775+07
13	Shopeepay	089692612004		localhost:8083/static/icon-method/1778918457997000500.png	2026-05-16 15:00:57.999392+07	2026-05-16 15:00:57.999392+07
\.


--
-- Data for Name: payment_proofs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_proofs (id, payment_id, proof_url, note, uploaded_at, verified_at, created_at) FROM stdin;
2	5	localhost:8083/static/bukti/1770832953926174500.png	pembayaran cetak dokumen sertifikat	2026-02-12 01:02:33.932276+07	\N	2026-02-12 01:02:33.932276+07
3	6	localhost:8083/static/bukti/1770834299580797100.png	pembayaran cetak dokumen sertifikat	2026-02-12 01:24:59.585778+07	\N	2026-02-12 01:24:59.585778+07
4	9	localhost:8083/static/bukti/1771173955605011600.png	pembayaran cetak dokumen sertifikat	2026-02-15 23:45:55.609386+07	\N	2026-02-15 23:45:55.609386+07
5	20	localhost:8083/static/bukti/1773000932113672700.png	pembayaran cetak foto	2026-03-09 03:15:32.118343+07	\N	2026-03-09 03:15:32.118343+07
6	21	localhost:8083/static/bukti/1773003558070033500.png	pembayaran cetak foto	2026-03-09 03:59:18.077662+07	\N	2026-03-09 03:59:18.077662+07
7	33	localhost:8083/static/bukti/1774571887345695200.png	bayar cetak	2026-03-27 07:38:07.354495+07	\N	2026-03-27 07:38:07.354495+07
8	34	localhost:8083/static/bukti/1774581352710278200.webp	Sudah bayar orderan ini	2026-03-27 10:15:52.715069+07	\N	2026-03-27 10:15:52.715069+07
9	35	localhost:8083/static/bukti/1774667680442423800.png	Saya lebihin dikit	2026-03-28 10:14:40.448792+07	\N	2026-03-28 10:14:40.448792+07
10	36	localhost:8083/static/bukti/1774951346332362600.jpg	bayar	2026-03-31 17:02:26.337533+07	\N	2026-03-31 17:02:26.337533+07
11	37	localhost:8083/static/bukti/1775016541723930700.jpg	bayar dokumen	2026-04-01 11:09:01.728562+07	\N	2026-04-01 11:09:01.728562+07
12	38	localhost:8083/static/bukti/1775618876708514000.png	bayar stickernya	2026-04-08 10:27:56.712687+07	\N	2026-04-08 10:27:56.712687+07
13	39	localhost:8083/static/bukti/1775904267520902800.JPG	bayar	2026-04-11 17:44:27.526356+07	\N	2026-04-11 17:44:27.526356+07
14	40	localhost:8083/static/bukti/1776302961608775900.JPG	bayarnya	2026-04-16 08:29:21.61477+07	\N	2026-04-16 08:29:21.61477+07
15	41	localhost:8083/static/bukti/1776873677485827000.JPG	byarnaya	2026-04-22 23:01:17.492219+07	\N	2026-04-22 23:01:17.492219+07
16	42	localhost:8083/static/bukti/1776949503895868600.jpg	Bayar cetak mya	2026-04-23 20:05:03.900599+07	\N	2026-04-23 20:05:03.900599+07
17	43	localhost:8083/static/bukti/1777037470305800900.jpg	Bayar	2026-04-24 20:31:10.313851+07	\N	2026-04-24 20:31:10.313851+07
18	45	localhost:8083/static/bukti/1777188973025453200.png	Bayar	2026-04-26 14:36:13.029326+07	\N	2026-04-26 14:36:13.029326+07
19	48	localhost:8083/static/bukti/1777190792626491100.jpg	Bayar cstak	2026-04-26 15:06:32.630769+07	\N	2026-04-26 15:06:32.630769+07
20	47	localhost:8083/static/bukti/1777190811579312600.jpg	batar foto	2026-04-26 15:06:51.581953+07	\N	2026-04-26 15:06:51.581953+07
21	49	localhost:8083/static/bukti/1777195284157525700.jpg	bayarnya	2026-04-26 16:21:24.169986+07	\N	2026-04-26 16:21:24.169986+07
22	50	localhost:8083/static/bukti/1777213087039491700.JPG	bukti	2026-04-26 21:18:07.044917+07	\N	2026-04-26 21:18:07.044917+07
23	51	localhost:8083/static/bukti/1777214268393197800.jpg	Rani	2026-04-26 21:37:48.398875+07	\N	2026-04-26 21:37:48.398875+07
24	53	localhost:8083/static/bukti/1777453323453980500.jpg	Bayar pake carmen	2026-04-29 16:02:03.461515+07	\N	2026-04-29 16:02:03.461515+07
25	54	localhost:8083/static/bukti/1777473290304594900.jpg	bayar cetak ijazahnya	2026-04-29 21:34:50.309631+07	\N	2026-04-29 21:34:50.309631+07
26	55	localhost:8083/static/bukti/1777567372482866200.jpg	Bayarnya ini	2026-04-30 23:42:52.488755+07	\N	2026-04-30 23:42:52.488755+07
27	56	localhost:8083/static/bukti/1777860444330843000.jpg	ini pembayarannya	2026-05-04 09:07:24.334522+07	\N	2026-05-04 09:07:24.334522+07
28	58	localhost:8083/static/bukti/1777910499597499200.png	Bayarnya	2026-05-04 23:01:39.629726+07	\N	2026-05-04 23:01:39.629726+07
29	59	localhost:8083/static/bukti/1777960231000098900.jpg	Bayar pake aqua	2026-05-05 12:50:31.007897+07	\N	2026-05-05 12:50:31.007897+07
30	60	localhost:8083/static/bukti/1778074564061350000.jpg	Ini bayarnya	2026-05-06 20:36:04.07418+07	\N	2026-05-06 20:36:04.07418+07
31	61	localhost:8083/static/bukti/1778084491542164400.jpg	byar 	2026-05-06 23:21:31.55027+07	\N	2026-05-06 23:21:31.55027+07
32	62	localhost:8083/static/bukti/1778084506548726800.jpg	Byr	2026-05-06 23:21:46.552146+07	\N	2026-05-06 23:21:46.552146+07
33	63	localhost:8083/static/bukti/1778085034799860500.JPG	byr\r\n	2026-05-06 23:30:34.802888+07	\N	2026-05-06 23:30:34.802888+07
34	64	localhost:8083/static/bukti/1778511213344521300.jpg	Pembayaran	2026-05-11 21:53:33.348477+07	\N	2026-05-11 21:53:33.348477+07
35	65	localhost:8083/static/bukti/1778561142728812000.png	Bayarnya 	2026-05-12 11:45:42.738011+07	\N	2026-05-12 11:45:42.738011+07
36	67	localhost:8083/static/bukti/1779178578945307300.webp	Ini bayarnya	2026-05-19 15:16:18.952649+07	\N	2026-05-19 15:16:18.952649+07
37	68	localhost:8083/static/bukti/1779202251521833200.webp	bayar ini	2026-05-19 21:50:51.526907+07	\N	2026-05-19 21:50:51.526907+07
38	69	localhost:8083/static/bukti/1779348908328539800.jpg	Bayar	2026-05-21 14:35:08.335603+07	\N	2026-05-21 14:35:08.335603+07
39	70	localhost:8083/static/bukti/1779514000904342800.JPG	bayar cetak foto	2026-05-23 12:26:40.908339+07	\N	2026-05-23 12:26:40.908339+07
40	71	localhost:8083/static/bukti/1779514232702681500.jpg	Bayar nya ini mas	2026-05-23 12:30:32.710863+07	\N	2026-05-23 12:30:32.710863+07
41	72	localhost:8083/static/bukti/1779537423650584200.png	bayar	2026-05-23 18:57:03.655047+07	\N	2026-05-23 18:57:03.655047+07
42	73	localhost:8083/static/bukti/1781150075279031400.JPG	ini filenya	2026-06-11 10:54:35.363162+07	\N	2026-06-11 10:54:35.363162+07
43	74	localhost:8083/static/bukti/1781960287161568000.JPG	bayar	2026-06-20 19:58:07.166383+07	\N	2026-06-20 19:58:07.166383+07
\.


--
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payments (id, user_id, order_id, amount, payment_method, status, paid_at, created_at, updated_at, approved_at) FROM stdin;
61	16	85	2000.00	BRI	Success	2026-05-06 23:21:31.554335+07	2026-05-06 23:20:37.90235+07	2026-05-06 23:32:17.113954+07	2026-05-06 23:32:17.113954+07
2	6	2	15600.00	BRI	Expired	\N	2026-02-11 00:06:18.081154+07	2026-02-11 02:06:35.719286+07	\N
3	6	4	2600.00	BRI	Cancelled	\N	2026-02-11 14:23:04.477699+07	2026-02-11 14:27:17.340454+07	\N
55	1	72	2200.00	Qris	Success	2026-04-30 23:42:52.490876+07	2026-04-30 23:41:37.317226+07	2026-04-30 23:44:11.400019+07	2026-04-30 23:44:11.400019+07
35	8	49	2200.00	Qris	Success	2026-03-28 10:14:40.454187+07	2026-03-28 09:40:15.973338+07	2026-03-28 10:17:05.206024+07	2026-03-28 10:17:05.206024+07
5	6	6	2600.00	Dana	Cancelled	2026-02-12 01:02:33.935504+07	2026-02-12 01:01:29.821928+07	2026-02-12 01:08:06.66061+07	\N
6	6	7	2600.00	Dana	Cancelled	2026-02-12 01:24:59.59415+07	2026-02-12 01:24:42.98931+07	2026-02-12 01:25:59.089423+07	\N
7	6	8	2600.00	Dana	Expired	\N	2026-02-12 15:16:09.897954+07	2026-02-12 17:22:40.974625+07	\N
8	6	9	2600.00	Dana	Expired	\N	2026-02-12 17:35:34.319352+07	2026-02-13 18:06:29.098017+07	\N
46	8	60	10000.00	BRI	Expired	\N	2026-04-26 14:56:43.126249+07	2026-04-26 17:00:29.781819+07	\N
9	1	10	2600.00	Dana	Success	2026-02-15 23:45:55.612214+07	2026-02-15 23:43:25.071586+07	2026-02-15 23:49:31.639472+07	2026-02-15 23:49:31.639472+07
10	1	11	2600.00	Dana	Expired	\N	2026-02-16 22:02:57.689607+07	2026-02-17 00:04:49.020895+07	\N
11	1	13	6000.00	Dana	Expired	\N	2026-02-24 06:56:02.658485+07	2026-02-24 09:00:00.038592+07	\N
12	1	15	6000.00	Dana	Expired	\N	2026-03-07 23:32:47.097036+07	2026-03-08 01:36:41.53031+07	\N
13	1	16	6000.00	Qris	Expired	\N	2026-03-08 02:46:56.464566+07	2026-03-08 04:49:15.323277+07	\N
14	1	17	6000.00	Qris	Expired	\N	2026-03-08 03:27:27.083593+07	2026-03-08 06:26:37.345243+07	\N
15	1	18	6000.00	Dana	Expired	\N	2026-03-08 03:46:07.609102+07	2026-03-08 06:26:37.345243+07	\N
16	1	19	6000.00	Dana	Cancelled	\N	2026-03-08 06:33:42.407712+07	2026-03-08 06:56:47.91001+07	\N
17	1	20	6000.00	Dana	Expired	\N	2026-03-08 07:43:11.161822+07	2026-03-08 09:47:26.062293+07	\N
18	1	21	6000.00	Dana	Expired	\N	2026-03-08 18:29:36.989724+07	2026-03-08 21:53:27.207627+07	\N
19	1	22	6000.00	Dana	Expired	\N	2026-03-08 18:48:27.075526+07	2026-03-08 21:53:27.207627+07	\N
36	1	50	48000.00	Dana	Success	2026-03-31 17:02:26.34136+07	2026-03-31 16:59:48.853784+07	2026-03-31 17:08:05.107759+07	2026-03-31 17:08:05.107759+07
20	1	23	6000.00	Dana	Cancelled	2026-03-09 03:15:32.179697+07	2026-03-09 03:09:35.087892+07	2026-03-09 03:29:38.397774+07	\N
21	1	24	6000.00	Qris	Success	2026-03-09 03:59:18.08089+07	2026-03-09 03:58:56.452584+07	2026-03-09 03:59:35.102542+07	2026-03-09 03:59:35.102542+07
22	1	29	5000.00	Qris	Expired	\N	2026-03-16 01:35:05.483875+07	2026-03-16 03:40:04.014818+07	\N
23	1	30	6000.00	BRI	Expired	\N	2026-03-16 03:20:39.83928+07	2026-03-16 15:03:45.031111+07	\N
24	1	31	5200.00	Dana	Expired	\N	2026-03-16 17:14:14.433681+07	2026-03-16 20:34:56.250146+07	\N
25	6	33	2600.00	Qris	Expired	\N	2026-03-17 01:28:51.331192+07	2026-03-17 08:08:36.452819+07	\N
26	8	34	1250.00	Dana	Expired	\N	2026-03-17 01:37:52.201035+07	2026-03-17 08:08:36.452819+07	\N
27	8	35	3200.00	Qris	Expired	\N	2026-03-17 02:08:18.00041+07	2026-03-17 08:08:36.452819+07	\N
28	8	37	6600.00	Qris	Expired	\N	2026-03-17 08:59:21.695001+07	2026-03-24 10:02:10.144169+07	\N
29	6	38	5000.00	Dana	Expired	\N	2026-03-24 10:15:46.962593+07	2026-03-24 12:16:19.096229+07	\N
30	6	39	6000.00	Qris	Expired	\N	2026-03-24 11:25:50.786194+07	2026-03-25 10:07:11.036907+07	\N
31	8	43	6000.00	BRI	Cancelled	\N	2026-03-26 08:03:44.361545+07	2026-03-26 08:20:46.91417+07	\N
32	8	46	5000.00	Dana	Cancelled	\N	2026-03-27 06:05:49.701642+07	2026-03-27 06:40:27.419564+07	\N
49	15	66	5000.00	Dana	Success	2026-04-26 16:21:24.175102+07	2026-04-26 16:07:05.415251+07	2026-04-26 19:32:46.82069+07	2026-04-26 19:32:46.82069+07
37	1	51	2500.00	BRI	Cancelled	2026-04-01 11:09:01.731775+07	2026-04-01 11:08:27.867068+07	2026-04-01 11:18:05.59729+07	\N
60	15	83	5200.00	Dana	Success	2026-05-06 20:36:04.078836+07	2026-05-06 20:33:50.010987+07	2026-05-06 20:42:50.735804+07	2026-05-06 20:42:50.735804+07
39	1	53	7800.00	BRI	Cancelled	2026-04-11 17:44:27.53701+07	2026-04-11 17:43:59.17399+07	2026-04-14 06:55:24.265845+07	\N
40	1	54	5000.00	Dana	Cancelled	2026-04-16 08:29:21.626319+07	2026-04-16 08:05:27.713047+07	2026-04-16 08:29:53.813474+07	\N
33	8	47	5000.00	Qris	Success	2026-03-27 07:38:07.367677+07	2026-03-27 07:16:02.250082+07	2026-04-22 23:55:46.406597+07	2026-04-22 23:55:46.406597+07
56	15	78	180000.00	Dana	Success	2026-05-04 09:07:24.344471+07	2026-05-04 09:06:30.019908+07	2026-05-04 09:57:52.318824+07	2026-05-04 09:57:52.318824+07
42	1	56	6000.00	Dana	Success	2026-04-23 20:05:03.905835+07	2026-04-23 20:03:11.683931+07	2026-04-23 20:07:24.938791+07	2026-04-23 20:07:24.938791+07
51	8	68	11000.00	Dana	Success	2026-04-26 21:37:48.400955+07	2026-04-26 21:36:22.149862+07	2026-04-26 21:39:33.577064+07	2026-04-26 21:39:33.577064+07
43	1	57	5000.00	BRI	Success	2026-04-24 20:31:10.320656+07	2026-04-24 20:26:17.75795+07	2026-04-24 20:56:34.904401+07	2026-04-24 20:56:34.904401+07
44	1	58	3200.00	Dana	Expired	\N	2026-04-25 19:18:32.708992+07	2026-04-25 21:20:17.103326+07	\N
50	1	67	2600.00	BRI	Success	2026-04-26 21:18:07.048643+07	2026-04-26 21:17:39.057952+07	2026-04-26 22:13:18.058172+07	2026-04-26 22:13:18.058172+07
45	8	59	1600.00	Qris	Success	2026-04-26 14:36:13.035328+07	2026-04-26 14:35:28.10063+07	2026-04-26 14:37:06.634884+07	2026-04-26 14:37:06.634884+07
52	8	69	5000.00	BRI	Expired	\N	2026-04-27 14:29:24.579602+07	2026-04-27 16:30:57.667807+07	\N
57	8	79	200000.00	Dana	Expired	\N	2026-05-04 15:21:29.821884+07	2026-05-04 22:54:09.424803+07	\N
53	15	70	30000.00	BRI	Cancelled	2026-04-29 16:02:03.467106+07	2026-04-29 15:58:51.611038+07	2026-04-29 16:05:17.330949+07	\N
34	1	48	10000.00	BRI	Success	2026-03-27 10:15:52.720155+07	2026-03-27 10:05:43.361928+07	2026-04-29 21:09:27.0649+07	2026-04-29 21:09:27.0649+07
54	1	71	2100.00	Dana	Success	2026-04-29 21:34:50.319515+07	2026-04-29 21:27:43.824292+07	2026-04-29 21:35:46.711536+07	2026-04-29 21:35:46.711536+07
41	1	55	2500.00	BRI	Success	2026-04-22 23:01:17.497+07	2026-04-22 23:00:55.092814+07	2026-04-29 21:38:03.410167+07	2026-04-29 21:38:03.410167+07
47	15	62	10000.00	BRI	Success	2026-04-26 15:06:51.583702+07	2026-04-26 15:03:29.314687+07	2026-04-30 22:41:59.526484+07	2026-04-30 22:41:59.526484+07
58	8	80	200000.00	BRI	Success	2026-05-04 23:01:39.670073+07	2026-05-04 23:01:16.391117+07	2026-05-04 23:03:14.158724+07	2026-05-04 23:03:14.158724+07
65	16	88	22000.00	Dana	Success	2026-05-12 11:45:42.743566+07	2026-05-12 11:45:07.86046+07	2026-05-12 11:46:25.182696+07	2026-05-12 11:46:25.182696+07
59	15	81	150000.00	Dana	Success	2026-05-05 12:50:31.031447+07	2026-05-05 12:49:52.488+07	2026-05-05 12:52:55.034859+07	2026-05-05 12:52:55.034859+07
38	8	52	5000.00	Dana	Success	2026-04-08 10:27:56.719843+07	2026-04-08 10:08:58.848418+07	2026-05-05 22:07:59.209054+07	2026-05-05 22:07:59.209054+07
63	16	86	13200.00	Qris	Cancelled	2026-05-06 23:30:34.806138+07	2026-05-06 23:23:58.884506+07	2026-05-28 16:25:49.231085+07	\N
48	8	61	2000.00	Dana	Success	2026-04-26 15:06:32.633682+07	2026-04-26 15:03:57.70517+07	2026-05-06 23:31:12.034104+07	2026-05-06 23:31:12.034104+07
64	15	87	200000.00	BRI	Success	2026-05-11 21:53:33.352548+07	2026-05-11 21:52:52.925864+07	2026-05-11 21:54:12.1916+07	2026-05-11 21:54:12.1916+07
66	16	89	5000.00	Gopay	Expired	\N	2026-05-16 15:23:35.30101+07	2026-05-16 20:36:41.649115+07	\N
67	1	90	5000.00	Shopeepay	Success	2026-05-19 15:16:18.9632+07	2026-05-19 15:15:24.452438+07	2026-05-19 15:20:17.335516+07	2026-05-19 15:20:17.335516+07
68	16	92	5000.00	Gopay	Success	2026-05-19 21:50:51.534509+07	2026-05-19 21:49:37.771017+07	2026-05-19 21:57:19.353199+07	2026-05-19 21:57:19.353199+07
69	1	93	25000.00	Shopeepay	Success	2026-05-21 14:35:08.347535+07	2026-05-21 14:33:48.249105+07	2026-05-21 14:36:01.404646+07	2026-05-21 14:36:01.404646+07
71	1	95	5000.00	Shopeepay	Success	2026-05-23 12:30:32.711891+07	2026-05-23 12:29:36.275797+07	2026-05-23 12:30:32.711891+07	\N
62	15	84	6000.00	Dana	Success	2026-05-06 23:21:46.5547+07	2026-05-06 23:21:09.724585+07	2026-05-28 16:11:31.605242+07	2026-05-28 16:11:31.605242+07
72	1	96	2000.00	Shopeepay	Success	2026-05-23 18:57:03.6576+07	2026-05-23 18:56:36.018419+07	2026-05-23 18:59:21.534101+07	2026-05-23 18:59:21.534101+07
70	8	94	6000.00	Gopay	Success	2026-05-23 12:26:40.912338+07	2026-05-23 12:26:29.454482+07	2026-06-20 21:18:39.069033+07	2026-06-20 21:18:39.069033+07
73	6	97	1600.00	Gopay	Cancelled	2026-06-11 10:54:35.368771+07	2026-06-11 10:54:09.554898+07	2026-06-20 20:45:44.678334+07	\N
74	15	98	5000.00	Gopay	Success	2026-06-20 19:58:07.172632+07	2026-06-20 19:56:18.422795+07	2026-06-20 20:34:33.016987+07	2026-06-20 20:34:33.016987+07
75	15	99	170000.00	Gopay	Expired	\N	2026-06-20 20:07:30.125446+07	2026-06-23 11:16:34.098793+07	\N
\.


--
-- Data for Name: refund_proofs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.refund_proofs (id, refund_id, file_url, note, created_at) FROM stdin;
1	2	localhost:8083/static/refund/1775018626578245900.jpg	bukti tranfernya	2026-04-01 11:43:46.582663
2	3	localhost:8083/static/refund/1776176404108739400.JPG	Tranfer refund ke shopeepay	2026-04-14 21:20:04.124871
3	4	localhost:8083/static/refund/1776667434650330700.JPG	tranfer refund ke rekening tertuju, mohon maaf atas tidak kenyamanannya	2026-04-20 13:43:54.765746
4	5	localhost:8083/static/refund/1777456461649353800.jpg	bukti refund	2026-04-29 16:54:21.65509
5	6	localhost:8083/static/refund/1779960525299530300.JPG	bukti baar ke tujuan	2026-05-28 16:28:45.303745
6	7	localhost:8083/static/refund/1781963612352209100.jpg	bukti	2026-06-20 20:53:32.355543
\.


--
-- Data for Name: refunds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.refunds (id, rejected_id, bank_name, account_number, account_name, status, transferred_at, created_at) FROM stdin;
1	1	dana	0897128923	wati	requested	\N	2026-03-30 09:28:26.576356
2	4	dana	0897128923	Ara	accepted	2026-04-06 12:17:07.313687	2026-04-01 11:20:49.864264
3	5	Shopeepay	089692612004	Annas	accepted	2026-04-16 07:58:07.84571	2026-04-14 08:23:16.37038
4	6	Dana	089692612004	Annas	accepted	2026-04-20 13:46:35.322517	2026-04-16 08:34:07.664891
5	7	Shopee pay	089692612004	Annas	accepted	2026-04-29 16:55:17.78801	2026-04-29 16:10:28.84277
6	8	Gopay	089692612004	Annas	accepted	2026-05-28 16:29:41.90965	2026-05-28 16:26:47.02078
7	9	ShopeePay	0896292612004	Annas Aulia Rahman	accepted	2026-06-20 20:55:36.559042	2026-06-20 20:49:08.515171
\.


--
-- Data for Name: rejected_payment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.rejected_payment (id, payment_id, user_id, amount, order_code, payment_code, order_name, admin_note, created_at) FROM stdin;
1	35	8	2200	cek_12123	payment_2Flsqv	cetak cek	alesan	2026-03-29 15:22:29.916156
2	35	8	2200	cek_12123	payment_2Flsqv	cetak cek	alesan	2026-03-29 15:33:04.301547
3	34	1	10000	cek_12123	payment_PdWB3G	cetak cek	alesan	2026-03-31 11:50:31.536909
4	37	1	2500	order_hxhINGzm	payment_hya6Ir	Ijazah	alesan bahan kosong, dan layanan lupa di nonaktifkan	2026-04-01 11:18:53.745036
5	39	1	7800	order_eJgYMcCo	payment_mDD5gT	Sertifikat	Bahan sedang kosong	2026-04-14 06:55:24.510079
6	40	1	5000	order_9cWAfZAI	payment_2LBXAZ	Cetak foto	bahan kosong	2026-04-16 08:29:53.979547
7	53	15	30000	order_2hQHDhh0	payment_AKi64K	Print Dokumen	kertas telah habis, dari admin akan mengajukan refund	2026-04-29 16:05:17.678605
8	63	16	13200	order_QcBKc4fZ	payment_ixoCOO	Print Dokumen	kertas bahan baru saja habis	2026-05-28 16:25:49.583923
9	73	6	1600	order_QcUwwlC0	payment_oVOeGV	Curiculum Vitae	kesalahan harga	2026-06-20 20:45:45.399629
\.


--
-- Name: list_payment_waiting_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.list_payment_waiting_id_seq', 68, true);


--
-- Name: payment_codes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payment_codes_id_seq', 75, true);


--
-- Name: payment_methods_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payment_methods_id_seq', 13, true);


--
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payments_id_seq', 75, true);


--
-- Name: proof_payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.proof_payments_id_seq', 43, true);


--
-- Name: refund_proofs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.refund_proofs_id_seq', 6, true);


--
-- Name: refunds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.refunds_id_seq', 7, true);


--
-- Name: rejected_payment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.rejected_payment_id_seq', 9, true);


--
-- Name: list_payment_waiting list_payment_waiting_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.list_payment_waiting
    ADD CONSTRAINT list_payment_waiting_pkey PRIMARY KEY (id);


--
-- Name: payment_codes payment_codes_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_codes
    ADD CONSTRAINT payment_codes_code_key UNIQUE (code);


--
-- Name: payment_codes payment_codes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_codes
    ADD CONSTRAINT payment_codes_pkey PRIMARY KEY (id);


--
-- Name: payment_methods payment_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_pkey PRIMARY KEY (id);


--
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: payment_proofs proof_payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_proofs
    ADD CONSTRAINT proof_payments_pkey PRIMARY KEY (id);


--
-- Name: refund_proofs refund_proofs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refund_proofs
    ADD CONSTRAINT refund_proofs_pkey PRIMARY KEY (id);


--
-- Name: refunds refunds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refunds
    ADD CONSTRAINT refunds_pkey PRIMARY KEY (id);


--
-- Name: rejected_payment rejected_payment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rejected_payment
    ADD CONSTRAINT rejected_payment_pkey PRIMARY KEY (id);


--
-- Name: list_payment_waiting fk_payment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.list_payment_waiting
    ADD CONSTRAINT fk_payment FOREIGN KEY (payment_id) REFERENCES public.payments(id) ON DELETE CASCADE;


--
-- Name: payment_proofs fk_payment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_proofs
    ADD CONSTRAINT fk_payment FOREIGN KEY (payment_id) REFERENCES public.payments(id) ON DELETE CASCADE;


--
-- Name: rejected_payment fk_payment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rejected_payment
    ADD CONSTRAINT fk_payment FOREIGN KEY (payment_id) REFERENCES public.payments(id) ON DELETE CASCADE;


--
-- Name: payment_codes fk_payment_codes_payment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_codes
    ADD CONSTRAINT fk_payment_codes_payment FOREIGN KEY (payment_id) REFERENCES public.payments(id) ON DELETE CASCADE;


--
-- Name: refund_proofs fk_refund; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refund_proofs
    ADD CONSTRAINT fk_refund FOREIGN KEY (refund_id) REFERENCES public.refunds(id) ON DELETE CASCADE;


--
-- Name: refunds fk_rejected; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refunds
    ADD CONSTRAINT fk_rejected FOREIGN KEY (rejected_id) REFERENCES public.rejected_payment(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict bn0WuARXAXmYSUaT9FcwC5f5Y2Dc1HUWEgqx0euXXc0xahiHbMtgZBe1pLl6Vhx

