CREATE TABLE IF NOT EXISTS urls (
    id bigserial not null,
    user_id integer not null,
    url varchar(255) not null,
    short_url varchar(255) not null unique,
    is_deleted boolean not null DEFAULT false,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
);