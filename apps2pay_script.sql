--
-- PostgreSQL database dump
--

-- Dumped from database version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)

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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cinemas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cinemas (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text,
    location text,
    city_id bigint NOT NULL
);


ALTER TABLE public.cinemas OWNER TO postgres;

--
-- Name: cinemas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cinemas_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cinemas_id_seq OWNER TO postgres;

--
-- Name: cinemas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cinemas_id_seq OWNED BY public.cinemas.id;


--
-- Name: cities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cities (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text
);


ALTER TABLE public.cities OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cities_id_seq OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- Name: movies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.movies (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    genre text,
    duration bigint,
    rating numeric,
    release_date timestamp with time zone,
    description text
);


ALTER TABLE public.movies OWNER TO postgres;

--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.movies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.movies_id_seq OWNER TO postgres;

--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: showtimes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.showtimes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    date timestamp with time zone,
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    movie_id bigint,
    studio_id bigint,
    cinema_id bigint
);


ALTER TABLE public.showtimes OWNER TO postgres;

--
-- Name: showtimes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.showtimes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.showtimes_id_seq OWNER TO postgres;

--
-- Name: showtimes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.showtimes_id_seq OWNED BY public.showtimes.id;


--
-- Name: studios; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.studios (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    total_seat bigint,
    cinema_id bigint NOT NULL
);


ALTER TABLE public.studios OWNER TO postgres;

--
-- Name: studios_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.studios_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.studios_id_seq OWNER TO postgres;

--
-- Name: studios_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.studios_id_seq OWNED BY public.studios.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text,
    password text,
    email text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: cinemas id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas ALTER COLUMN id SET DEFAULT nextval('public.cinemas_id_seq'::regclass);


--
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: showtimes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes ALTER COLUMN id SET DEFAULT nextval('public.showtimes_id_seq'::regclass);


--
-- Name: studios id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.studios ALTER COLUMN id SET DEFAULT nextval('public.studios_id_seq'::regclass);


--
-- Data for Name: cinemas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cinemas (id, created_at, updated_at, name, location, city_id) FROM stdin;
1	2024-01-01 10:00:00+07	2024-01-01 10:00:00+07	Cinema One	123 Main St	1
2	2024-01-02 11:00:00+07	2024-01-02 11:00:00+07	Cinema Two	456 Oak Ave	2
3	2024-01-03 12:00:00+07	2024-01-03 12:00:00+07	Cinema Three	789 Pine Rd	3
4	2024-01-04 13:00:00+07	2024-01-04 13:00:00+07	Cinema Four	321 Birch Blvd	4
5	2024-01-05 14:00:00+07	2024-01-05 14:00:00+07	Cinema Five	654 Cedar St	5
6	2024-01-06 15:00:00+07	2024-01-06 15:00:00+07	Cinema Six	987 Maple Dr	6
7	2024-01-07 16:00:00+07	2024-01-07 16:00:00+07	Cinema Seven	213 Spruce Ln	7
8	2024-01-08 17:00:00+07	2024-01-08 17:00:00+07	Cinema Eight	546 Elm St	8
9	2024-01-09 18:00:00+07	2024-01-09 18:00:00+07	Cinema Nine	879 Fir Ave	9
10	2024-01-10 19:00:00+07	2024-01-10 19:00:00+07	Cinema Ten	112 Willow Rd	10
\.


--
-- Data for Name: cities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cities (id, created_at, updated_at, name) FROM stdin;
1	2024-01-01 10:00:00+07	2024-01-01 10:00:00+07	Jakarta
2	2024-01-02 11:00:00+07	2024-01-02 11:00:00+07	Surabaya
3	2024-01-03 12:00:00+07	2024-01-03 12:00:00+07	Bandung
4	2024-01-04 13:00:00+07	2024-01-04 13:00:00+07	Medan
5	2024-01-05 14:00:00+07	2024-01-05 14:00:00+07	Semarang
6	2024-01-06 15:00:00+07	2024-01-06 15:00:00+07	Makassar
7	2024-01-07 16:00:00+07	2024-01-07 16:00:00+07	Palembang
8	2024-01-08 17:00:00+07	2024-01-08 17:00:00+07	Bogor
9	2024-01-09 18:00:00+07	2024-01-09 18:00:00+07	Malang
10	2024-01-10 19:00:00+07	2024-01-10 19:00:00+07	Yogyakarta
\.


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.movies (id, created_at, updated_at, deleted_at, title, genre, duration, rating, release_date, description) FROM stdin;
1	2024-01-01 10:00:00+07	2024-01-01 10:00:00+07	\N	Inception	Sci-Fi	148	8.8	2010-07-16 00:00:00+07	A thief who steals corporate secrets through the use of dream-sharing technology.
2	2024-01-02 11:00:00+07	2024-01-02 11:00:00+07	\N	The Dark Knight	Action	152	9.0	2008-07-18 00:00:00+07	When the menace known as the Joker emerges from his mysterious past, he wreaks havoc and chaos on the people of Gotham.
3	2024-01-03 12:00:00+07	2024-01-03 12:00:00+07	\N	Interstellar	Adventure	169	8.6	2014-11-07 00:00:00+07	A team of explorers travel through a wormhole in space in an attempt to ensure humanity survival.
4	2024-01-04 13:00:00+07	2024-01-04 13:00:00+07	\N	The Matrix	Sci-Fi	136	8.7	1999-03-31 00:00:00+07	A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.
5	2024-01-05 14:00:00+07	2024-01-05 14:00:00+07	\N	Forrest Gump	Drama	142	8.8	1994-07-06 00:00:00+07	The presidencies of Kennedy and Johnson, the Vietnam War, the Watergate scandal and other historical events unfold through the perspective of an Alabama man.
6	2024-01-06 15:00:00+07	2024-01-06 15:00:00+07	\N	The Shawshank Redemption	Drama	142	9.3	1994-09-23 00:00:00+07	Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.
7	2024-01-07 16:00:00+07	2024-01-07 16:00:00+07	\N	Pulp Fiction	Crime	154	8.9	1994-10-14 00:00:00+07	The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.
8	2024-01-08 17:00:00+07	2024-01-08 17:00:00+07	\N	Fight Club	Drama	139	8.8	1999-10-15 00:00:00+07	An insomniac office worker and a devil-may-care soap maker form an underground fight club that evolves into much more.
9	2024-01-09 18:00:00+07	2024-01-09 18:00:00+07	\N	The Godfather	Crime	175	9.2	1972-03-24 00:00:00+07	The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.
10	2024-01-10 19:00:00+07	2024-01-10 19:00:00+07	\N	The Lord of the Rings: The Return of the King	Fantasy	201	8.9	2003-12-17 00:00:00+07	Gandalf and Aragorn lead the World of Men against Sauron army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.
\.


--
-- Data for Name: showtimes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.showtimes (id, created_at, updated_at, deleted_at, date, start_time, end_time, movie_id, studio_id, cinema_id) FROM stdin;
1	2024-01-01 10:00:00+07	2024-01-01 10:00:00+07	\N	2024-01-15 00:00:00+07	2024-01-15 10:00:00+07	2024-01-15 12:30:00+07	1	1	1
2	2024-01-02 11:00:00+07	2024-01-02 11:00:00+07	\N	2024-01-16 00:00:00+07	2024-01-16 13:00:00+07	2024-01-16 15:30:00+07	2	3	1
3	2024-01-03 12:00:00+07	2024-01-03 12:00:00+07	\N	2024-01-17 00:00:00+07	2024-01-17 16:00:00+07	2024-01-17 18:30:00+07	2	3	1
4	2024-01-04 13:00:00+07	2024-01-04 13:00:00+07	\N	2024-01-18 00:00:00+07	2024-01-18 19:00:00+07	2024-01-18 21:30:00+07	2	3	1
5	2024-01-05 14:00:00+07	2024-01-05 14:00:00+07	\N	2024-01-19 00:00:00+07	2024-01-19 22:00:00+07	2024-01-19 00:30:00+07	5	5	1
6	2024-01-06 15:00:00+07	2024-01-06 15:00:00+07	\N	2024-01-20 00:00:00+07	2024-01-20 10:00:00+07	2024-01-20 12:30:00+07	6	2	1
7	2024-01-07 16:00:00+07	2024-01-07 16:00:00+07	\N	2024-01-21 00:00:00+07	2024-01-21 13:00:00+07	2024-01-21 15:30:00+07	7	8	2
8	2024-01-08 17:00:00+07	2024-01-08 17:00:00+07	\N	2024-01-22 00:00:00+07	2024-01-22 16:00:00+07	2024-01-22 18:30:00+07	8	6	3
9	2024-01-09 18:00:00+07	2024-01-09 18:00:00+07	\N	2024-01-23 00:00:00+07	2024-01-23 19:00:00+07	2024-01-23 21:30:00+07	9	5	2
10	2024-01-10 19:00:00+07	2024-01-10 19:00:00+07	\N	2024-01-24 00:00:00+07	2024-01-24 22:00:00+07	2024-01-24 00:30:00+07	10	10	4
11	2024-05-22 10:17:13.701982+07	2024-05-22 10:17:13.701982+07	\N	2024-05-22 07:00:00+07	2006-01-02 22:04:05+07	2006-01-02 22:04:05+07	1	1	1
\.


--
-- Data for Name: studios; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.studios (id, created_at, updated_at, deleted_at, name, total_seat, cinema_id) FROM stdin;
1	2024-01-01 10:00:00+07	2024-01-01 10:00:00+07	\N	Studio One	100	1
2	2024-01-02 11:00:00+07	2024-01-02 11:00:00+07	\N	Studio Two	120	1
3	2024-01-03 12:00:00+07	2024-01-03 12:00:00+07	\N	Studio Three	150	1
4	2024-01-04 13:00:00+07	2024-01-04 13:00:00+07	\N	Studio Four	200	1
5	2024-01-05 14:00:00+07	2024-01-05 14:00:00+07	\N	Studio Five	80	2
6	2024-01-06 15:00:00+07	2024-01-06 15:00:00+07	\N	Studio Six	60	3
7	2024-01-07 16:00:00+07	2024-01-07 16:00:00+07	\N	Studio Seven	90	2
8	2024-01-08 17:00:00+07	2024-01-08 17:00:00+07	\N	Studio Eight	110	2
9	2024-01-09 18:00:00+07	2024-01-09 18:00:00+07	\N	Studio Nine	130	3
10	2024-01-10 19:00:00+07	2024-01-10 19:00:00+07	\N	Studio Ten	140	4
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, created_at, updated_at, deleted_at, username, password, email) FROM stdin;
6f84d2e8-e9c5-423d-bdeb-cf16b153d499	2024-05-22 10:14:32.773329+07	2024-05-22 10:14:32.773329+07	\N	123123	$2a$10$F6b8vZ7BnFcMgHuLa7WZ1e3qFyws.DN3UxbOKBojQk1ApO85zmHxy	adminq12211@gmaillll.com
\.


--
-- Name: cinemas_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cinemas_id_seq', 10, true);


--
-- Name: cities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cities_id_seq', 10, true);


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_id_seq', 10, true);


--
-- Name: showtimes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.showtimes_id_seq', 11, true);


--
-- Name: studios_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.studios_id_seq', 10, true);


--
-- Name: cinemas cinemas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT cinemas_pkey PRIMARY KEY (id);


--
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: showtimes showtimes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_pkey PRIMARY KEY (id);


--
-- Name: studios studios_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.studios
    ADD CONSTRAINT studios_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: cinemas fk_cinemas_city; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT fk_cinemas_city FOREIGN KEY (city_id) REFERENCES public.cities(id);


--
-- Name: showtimes fk_showtimes_cinema; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT fk_showtimes_cinema FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id);


--
-- Name: showtimes fk_showtimes_movie; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT fk_showtimes_movie FOREIGN KEY (studio_id) REFERENCES public.movies(id);


--
-- Name: showtimes fk_showtimes_studio; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT fk_showtimes_studio FOREIGN KEY (studio_id) REFERENCES public.studios(id);


--
-- Name: studios fk_studios_cinema; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.studios
    ADD CONSTRAINT fk_studios_cinema FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id);


--
-- PostgreSQL database dump complete
--

