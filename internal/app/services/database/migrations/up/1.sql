CREATE TABLE IF NOT EXISTS users
(
    id bigserial not null,
    guid varchar(255) not null,
    PRIMARY KEY(id)
);
