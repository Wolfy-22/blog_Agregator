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

-- name: MarkFeedFetched :one
update feeds
set updated_at = $2, last_fetched_at = $2
where id = $1
returning *;

-- name: GetNextFeedToFetch :one
select * from feeds
order by last_fetched_at asc nulls first
limit 1;