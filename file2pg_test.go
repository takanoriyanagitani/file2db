package file2db

import (
	"testing"

	"database/sql"
	_ "github.com/lib/pq"
)

func TestFile2pg(t *testing.T) {
	d, e1 := sql.Open("postgres", "")
	defer d.Close()
	switch e1 {
	case nil:
		break
	default:
		t.Fatalf("Unable to connect to test database: %v\n", e1)
	}

	_, e2 := d.Exec(`CREATE SCHEMA IF NOT EXISTS file2pg`)
	switch e2 {
	case nil:
		break
	default:
		t.Fatalf("Unable to create test schema file2pg: %v", e2)
	}
}
