-- Table: public.Forum

-- DROP TABLE IF EXISTS public."Forum";

CREATE TABLE IF NOT EXISTS public."Forum"
(
    id bigint NOT NULL,
    name character varying COLLATE pg_catalog."default",
    "topicKeyword" character varying COLLATE pg_catalog."default",
    users character varying[] COLLATE pg_catalog."default",
    CONSTRAINT "Forum_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Forum"
    OWNER to hcidzrejqbzxrj;


-- Table: public.User

-- DROP TABLE IF EXISTS public."User";

CREATE TABLE IF NOT EXISTS public."User"
(
    name character varying COLLATE pg_catalog."default",
    topics character varying[] COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."User"
    OWNER to hcidzrejqbzxrj;