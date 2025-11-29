-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, name, url, user_id)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
returning *;

-- name: DeleteFeeds :exec
delete from feeds;

-- name: GetFeeds :many
select *
from feeds;

-- name: GetFeedByURL :one
select * from feeds where url = $1;

-- name: UnfollowFeed :one
delete * from feeds where url = $1;