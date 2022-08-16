CREATE TABLE products
(
    id smallint NOT NULL,
    code text  COLLATE pg_catalog."default",
    name text COLLATE pg_catalog."default",
    Category text COLLATE pg_catalog."default",
    Price text COLLATE pg_catalog."default",
    Color text COLLATE pg_catalog."default",
    Size bigint COLLATE pg_catalog."default",
    CONSTRAINT products._pkey PRIMARY KEY (id)

);

INSERT INTO products(id, Code, Name, Category,  Price, Color,Size  ) VALUES
    (1, "yönetici", "ahmet", "person", 135, "blue", 180),
    (2, "çalışan", "ege", "person", 55, "blue", 180),
    (3, "çalışan", "ozan",  "person",  65, "blue", 180);

