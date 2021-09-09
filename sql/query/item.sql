-- name: ListItems :many
select *
from items
order by code
LIMIT $1 OFFSET $2;
-- name: GetItemById :one
select *
from items
where id = $1
limit 1;
-- name: CreateItem :one
insert into items (name, code, unit, cost)
VALUES ($1, $2, $3, $4)
returning *;
-- name: UpdateItem :one
update items
set name = $2,
  code = $3,
  unit = $4,
  cost = $5
where id = $1
returning *;
-- name: DeleteItem :exec
delete from items
where id = $1;