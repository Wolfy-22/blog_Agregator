-- +goose Up
create table feeds (
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name text not null,
    url  text not null UNIQUE,

    user_id UUID not null references users(id)
    on delete cascade
);

-- +goose Down
drop table feeds;