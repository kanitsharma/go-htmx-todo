// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package todos

import (
	"context"
)

const listTodos = `-- name: ListTodos :many
SELECT id, name, status FROM todos 
ORDER BY name
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Name, &i.Status); err != nil {
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
