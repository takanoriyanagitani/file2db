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

	_, e3 := d.Exec(`DROP TABLE IF EXISTS file2pg.filestore`)
	switch e3 {
	case nil:
		break
	default:
		t.Fatalf("Unable to drop test table file2pg.filestore: %v", e3)
	}

	_, e4 := d.Exec(`
	  CREATE TABLE IF NOT EXISTS file2pg.filestore(
		  name TEXT   NOT NULL,
			size BIGINT NOT NULL,
			unix BIGINT NOT NULL,
			sqno BIGSERIAL NOT NULL,
			meta JSONB,
			created_time TIMESTAMP WITH TIME ZONE NOT NULL,
			updated_time TIMESTAMP WITH TIME ZONE,
			CONSTRAINT filestore_pkc PRIMARY KEY(name, size, unix, sqno)
		)
	`)
	switch e4 {
	case nil:
		break
	default:
		t.Fatalf("Unable to drop test table file2pg.filestore: %v", e4)
	}

	s5, e5 := d.Prepare(`
	  INSERT INTO file2pg.filestore(
		  name,
			size,
			unix,
			meta,
			created_time
		)
		VALUES (
		  $1::TEXT,
			$2::BIGINT,
			$3::BIGINT,
			$4::JSONB,
			CLOCK_TIMESTAMP()
		)
	`)
	defer s5.Close()

	switch e5 {
	case nil:
		break
	default:
		t.Fatalf("Unable to get prepared statement: %v\n", e5)
	}
}
