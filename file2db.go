package file2db

import (
	"database/sql"
	"fmt"
	"os"
)

type Saver interface {
	Save(meta *os.FileInfo, content []byte) error
}

type sqlSaver struct{ s *sql.Stmt }

func (ss *sqlSaver) CheckResultCount(i int64) error {
	switch i {
	default:
		return fmt.Errorf("Unexpected error; rows affected: %v\n", i)
	case 0:
		return nil
	case 1:
		return nil
	}
}

func (ss *sqlSaver) CheckResult(r sql.Result) error {
	i, e := r.RowsAffected()
	switch e {
	default:
		return ss.CheckResultCount(i)
	case nil:
		return nil
	}
}

func (ss *sqlSaver) Save(meta *os.FileInfo, content []byte) error {
	r, e := ss.s.Exec()
	switch e {
	default:
		return e
	case nil:
		return ss.CheckResult(r)
	}
}

func save(meta *os.FileInfo, content []byte, s *sql.Stmt) {
}

func file2db(f *os.File, s *sql.Stmt) {
}
