//go:build integration

package app

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestPostgresAcceptsAWriteAndARead(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	t.Cleanup(cancel)

	container, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("app"),
		postgres.WithUsername("app"),
		postgres.WithPassword("app"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(time.Minute),
		),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = container.Terminate(context.Background())
	})

	connString, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = conn.Close(context.Background())
	})

	if _, err := conn.Exec(ctx, "CREATE TABLE notes (id serial PRIMARY KEY, body text NOT NULL)"); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Exec(ctx, "INSERT INTO notes (body) VALUES ($1)", "hello"); err != nil {
		t.Fatal(err)
	}

	var body string
	if err := conn.QueryRow(ctx, "SELECT body FROM notes").Scan(&body); err != nil {
		t.Fatal(err)
	}
	if body != "hello" {
		t.Fatalf("body = %q", body)
	}
}
