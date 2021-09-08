// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
)

const deleteItemById = `-- name: DeleteItemById :exec
delete from items where id = $1
`

func (q *Queries) DeleteItemById(ctx context.Context, id uint64) error {
	_, err := q.db.ExecContext(ctx, deleteItemById, id)
	return err
}

const insertItem = `-- name: InsertItem :one
insert into items (
    name, code, unit, cost
) VALUES (
    $1, $2, $3, $4
)
returning id, name, code, unit, cost, created_at, updated_at
`

type InsertItemParams struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Unit string `json:"unit"`
	Cost string `json:"cost"`
}

func (q *Queries) InsertItem(ctx context.Context, arg InsertItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, insertItem,
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

const selectItemById = `-- name: SelectItemById :one
select id, name, code, unit, cost, created_at, updated_at from items where id = $1 limit 1
`

func (q *Queries) SelectItemById(ctx context.Context, id uint64) (Item, error) {
	row := q.db.QueryRowContext(ctx, selectItemById, id)
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

const selectItems = `-- name: SelectItems :many
select id, name, code, unit, cost, created_at, updated_at from items order by code
`

func (q *Queries) SelectItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, selectItems)
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
