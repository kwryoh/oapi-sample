-- name: SelectItems :many
select * from items order by code;

-- name: SelectItemById :one
select * from items where id = $1 limit 1;

-- name: InsertItem :one
insert into items (
    name, code, unit, cost
) VALUES (
    $1, $2, $3, $4
)
returning *;

-- name: UpdateItemById :one
update items
set name = $2
  , code = $3
  , unit = $4
  , cost = $5
where id = $1
returning *;

-- name: DeleteItemById :exec
delete from items where id = $1;