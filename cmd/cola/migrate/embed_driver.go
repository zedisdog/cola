package migrate

import (
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4/source"
	"io"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"
)

type fileName struct {
	dir  string
	name string
	ext  string
}

func (f fileName) toString(action string) string {
	return fmt.Sprintf("%s/%s.%s.%s", f.dir, f.name, action, f.ext)
}

func NewEmbed(f embed.FS) *Embed {
	e := &Embed{
		sorts: make([]uint, 0, 20),
		files: make(map[uint]fileName),
		fs:    f,
	}
	dirEntries, _ := fs.ReadDir(f, "migrations")
	for _, entry := range dirEntries {
		sFileName := strings.Split(entry.Name(), ".")

		// 保存版本slice
		u64, err := strconv.ParseUint(strings.Split(sFileName[0], "_")[0], 10, 64)
		if err != nil {
			panic(err)
		}
		version := uint(u64)
		_, err = e.find(version)
		if err != nil {
			e.sorts = append(e.sorts, version)
		}

		// 保存文件名
		if _, ok := e.files[version]; !ok {
			e.files[version] = fileName{
				dir:  "migrations",
				name: strings.Join(sFileName[:len(sFileName)-2], "."),
				ext:  sFileName[len(sFileName)-1],
			}
		}
	}
	sort.Slice(e.sorts, func(i, j int) bool {
		return e.sorts[i] < e.sorts[j]
	})
	return e
}

type Embed struct {
	sorts []uint
	fs    embed.FS
	files map[uint]fileName
}

func (e *Embed) Open(url string) (source.Driver, error) {
	return e, nil
}

func (e Embed) Close() error {
	return nil
}

func (e Embed) First() (version uint, err error) {
	if len(e.sorts) < 1 {
		return 0, os.ErrNotExist
	}
	return e.sorts[0], nil
}

func (e Embed) find(version uint) (index int, err error) {
	var ver uint
	for index, ver = range e.sorts {
		if ver == version {
			return
		}
	}
	return 0, os.ErrNotExist
}

func (e Embed) Prev(version uint) (prevVersion uint, err error) {
	index, err := e.find(version)
	if err != nil {
		return
	}
	if index-1 >= 0 {
		return e.sorts[index-1], nil
	}
	return 0, os.ErrNotExist
}

func (e Embed) Next(version uint) (nextVersion uint, err error) {
	index, err := e.find(version)
	if err != nil {
		return
	}
	if index+1 < len(e.sorts) {
		return e.sorts[index+1], nil
	}
	return 0, os.ErrNotExist
}

func (e Embed) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	file, ok := e.files[version]
	if !ok {
		return nil, "", os.ErrNotExist
	}
	identifier = file.toString("up")
	r, err = e.fs.Open(identifier)
	return
}

func (e Embed) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	file, ok := e.files[version]
	if !ok {
		return nil, "", os.ErrNotExist
	}
	identifier = file.toString("down")
	r, err = e.fs.Open(identifier)
	return
}
