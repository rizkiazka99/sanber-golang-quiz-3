-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
    id BIGINT PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    created_by varchar(255) NOT NULL,
    modified_at timestamp NOT NULL,
    modified_by varchar(255) NOT NULL
);

CREATE TABLE books (
    id BIGINT PRIMARY KEY NOT NULL,
    title varchar(255) NOT NULL,
    description varchar(500) NOT NULL,
    image_url varchar(255) NOT NULL,
    release_year int NOT NULL,
    price int NOT NULL,
    total_page int NOT NULL,
    thickness varchar(128) NOT NULL,
    category_id BIGINT NOT NULL,
    created_at timestamp NOT NULL,
    created_by varchar(255) NOT NULL,
    modified_at timestamp NOT NULL,
    modified_by varchar(255) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE users (
    id BIGINT PRIMARY KEY NOT NULL,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    created_by varchar(255) NOT NULL,
    modified_at timestamp NOT NULL,
    modified_by varchar(255) NOT NULL
);

-- +migrate StatementEnd