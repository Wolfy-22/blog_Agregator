-- +goose Up
Create table users (
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name TEXT UNIQUE not null
);

-- +goose Down
drop table users;