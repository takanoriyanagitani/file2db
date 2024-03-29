package file2db

import (
	"io/ioutil"
	"os"
	"time"
)

// TimeSpecJ time.Time.Unix() and time.Time.UnixNano() as json
type TimeSpecJ struct {
	Seconds int64 `json:"seconds"`
	Nanos   int64 `json:"nanos"`
}

// FileStatJ contains os.File.Stat() as json
type FileStatJ struct {
	Mode     uint32    `json:"mode"`
	Size     int64     `json:"size"`
	Modified TimeSpecJ `json:"modified"`
}

// FileMetaJ contains file meta info
type FileMetaJ struct {
	Name string    `json:"name"`
	Stat FileStatJ `json:"stat"`
}

// FileInfo contains file meta info and its contents
type FileInfo struct {
	Meta FileMetaJ
	Data []byte
}

func time2j(t time.Time) TimeSpecJ {
	return TimeSpecJ{
		Seconds: t.Unix(),
		Nanos:   t.UnixNano(),
	}
}

func file2info(filename string, info *FileInfo, data []byte, meta os.FileInfo) error {
	info.Data = data
	info.Meta = FileMetaJ{
		Name: filename,
		Stat: FileStatJ{
			Mode:     uint32(meta.Mode()),
			Size:     meta.Size(),
			Modified: time2j(meta.ModTime()),
		},
	}
	return nil
}
func loadMeta(filename string, info *FileInfo, f *os.File, data []byte) error {
	m, e := f.Stat()
	switch e {
	default:
		return e
	case nil:
		return file2info(filename, info, data, m)
	}
}
func loadFile(filename string, info *FileInfo, f *os.File) error {
	b, e := ioutil.ReadAll(f)
	switch e {
	default:
		return e
	case nil:
		return loadMeta(filename, info, f, b)
	}
}

// LoadFile loads file into go object. assuming file is not nil.
func LoadFile(filename string, file *FileInfo) error {
	f, e := os.Open(filename)
	defer f.Close()
	switch e {
	default:
		return e
	case nil:
		return loadFile(filename, file, f)
	}
}
