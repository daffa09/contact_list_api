CREATE TABLE album
(
    id         VARCHAR PRIMARY KEY,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


CREATE TABLE contact (
    id INTEGER PRIMARY KEY,
    name VARCHAR NOT NULL,
    age INT NOT NULL,
    email VARCHAR NOT NULL,
    phone VARCHAR NOT NULL
);
