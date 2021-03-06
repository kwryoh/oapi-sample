// Code generated by sqlc. DO NOT EDIT.
// source: item.sql

package db

import (
	"context"
)

const createItem = `-- name: CreateItem :one
insert into items (name, code, unit, cost)
VALUES ($1, $2, $3, $4)
returning id, name, code, unit, cost, created_at, updated_at
`

type CreateItemParams struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Unit string `json:"unit"`
	Cost string `json:"cost"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.Name,
		arg.Code,
		arg.Unit,
		arg.Cost,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Unit,
		&i.Cost,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :exec
delete from items
where id = $1
`

func (q *Queries) DeleteItem(ctx context.Context, id uint64) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const findItemById = `-- name: FindItemById :one
select id, name, code, unit, cost, created_at, updated_at
from items
where id = $1
limit 1
`

func (q *Queries) FindItemById(ctx context.Context, id uint64) (Item, error) {
	row := q.db.QueryRowContext(ctx, findItemById, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Unit,
		&i.Cost,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listItems = `-- name: ListItems :many
select id, name, code, unit, cost, created_at, updated_at
from items
order by code
LIMIT $1 OFFSET $2
`

type ListItemsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListItems(ctx context.Context, arg ListItemsParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItems, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Unit,
			&i.Cost,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateItem = `-- name: UpdateItem :one
update items
set name = $2,
  code = $3,
  unit = $4,
  cost = $5
where id = $1
returning id, name, code, unit, cost, created_at, updated_at
`

type UpdateItemParams struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Unit string `json:"unit"`
	Cost string `json:"cost"`
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, updateItem,
		arg.ID,
		arg.Name,
		arg.Code,
		arg.Unit,
		arg.Cost,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Unit,
		&i.Cost,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
