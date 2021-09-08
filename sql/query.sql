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

-- name: DeleteItemById :exec
delete from items where id = $1;