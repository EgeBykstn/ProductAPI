CREATE TABLE products
(
    id bigint NOT NULL,
    name text COLLATE pg_catalog."default",
    CONSTRAINT student_pkey PRIMARY KEY (id)

);

INSERT INTO students(id, code, Name, Category,  Price, Color,Size  ) VALUES
    (1, "yönetici", "ege", "person", 135, "blue", 180),
    (2, "çalışan", "ege", "person", 55, "blue", 180),
    (3, "head", "ege",  "person",  65, "blue", 180);

