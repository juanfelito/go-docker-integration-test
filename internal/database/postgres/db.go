package postgres

import (
	"context"
	"fmt"

	"docker-example/internal/database"

	"github.com/jackc/pgx/v4/pgxpool"
)

type db struct {
	pool *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *db {
	d := &db{
		pool: pool,
	}
	return d
}

func (d *db) ReadMessage(id string) string {
	query := `SELECT id, content FROM message WHERE id = $1`

	row := d.pool.QueryRow(context.Background(), query, id)

	var message database.Message
	if err := row.Scan(&message.ID, &message.Message); err != nil {
		fmt.Println("error scanning result into the struct: ", err)
		return ""
	}

	updateQuery := `UPDATE message SET seen = true WHERE id = $1`

	if _, err := d.pool.Exec(context.Background(), updateQuery, id); err != nil {
		fmt.Println("error updating last seen field: ", err)
	}

	return message.Message
}
