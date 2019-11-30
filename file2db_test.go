package file2db

import (
	"testing"

	"bytes"
	"io/ioutil"
	"os"
)

func TestLoadFile(t *testing.T) {
	f, e1 := os.Open("./file2db_test.go")
	defer f.Close()
	switch e1 {
	case nil:
		break
	default:
		t.Fatalf("Unable to open file: %v\n", e1)
	}

	m, e3 := f.Stat()
	switch e3 {
	case nil:
		break
	default:
		t.Fatalf("Unable to get meta info: %v\n", e3)
	}

	b, e4 := ioutil.ReadAll(f)
	switch e4 {
	case nil:
		break
	default:
		t.Fatalf("Unable to load: %v\n", e4)
	}

	t.Run("file exists", func(t *testing.T) {
		fi := FileInfo{}
		e5 := LoadFile("file2db_test.go", &fi)
		switch e5 {
		case nil:
			break
		default:
			t.Fatalf("Unable to load: %v\n", e5)
		}

		switch bytes.Equal(b, fi.Data) {
		case false:
			t.Errorf("bytes differ.")
		}

		switch fi.Meta.Name == m.Name() {
		case false:
			t.Errorf("filename differ; %v != %v\n", fi.Meta.Name, m.Name())
		}
		switch fi.Meta.Stat.Mode == uint32(m.Mode()) {
		case false:
			t.Errorf("file mode differ; %v != %v\n", fi.Meta.Stat.Mode, m.Mode())
		}
		switch fi.Meta.Stat.Size == m.Size() {
		case false:
			t.Errorf("file size differ; %v != %v\n", fi.Meta.Stat.Size, m.Size())
		}
		switch fi.Meta.Stat.Modified.Seconds == m.ModTime().Unix() {
		case false:
			t.Errorf("time in seconds differ; %v != %v\n", fi.Meta.Stat.Modified.Seconds, m.ModTime().Unix())
		}
		switch fi.Meta.Stat.Modified.Nanos == m.ModTime().UnixNano() {
		case false:
			t.Errorf("time in nanos differ; %v != %v\n", fi.Meta.Stat.Modified.Nanos, m.ModTime().UnixNano())
		}
	})

	t.Run("file not exists", func(t *testing.T) {
		e6 := LoadFile("assuming-this-file-does-not-exist", nil)
		switch e6 {
		case nil:
			t.Errorf("should return error.")
		default:
			break
		}
	})
}
