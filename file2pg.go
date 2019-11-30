package file2db

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func file2pg(data []byte, s *sql.Stmt, j []byte) error {
	r, e1 := s.Exec(j, data)
	switch e1 {
	case nil:
		break
	default:
		return e1
	}

	switch r {
	case nil:
		return fmt.Errorf("result nil.")
	}

	count, e2 := r.RowsAffected()
	switch e2 {
	case nil:
		break
	default:
		return e2
	}

	switch count {
	case 1:
		return nil
	default:
		return fmt.Errorf("count != 1")
	}
}

// File2pg inserts a row using specified statement.
func File2pg(f *FileInfo, s *sql.Stmt) error {
	b, e := json.Marshal(&f.Meta)
	switch e {
	case nil:
		return file2pg(f.Data, s, b)
	default:
		return e
	}
}
