CREATE TABLE IF NOT EXISTS urls (
    id bigserial not null,
    user_id integer not null,
    url varchar(255) not null unique,
    short_url varchar(255) not null unique,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
);