package migrate

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4/source"
	"io"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"
)

var embedDriver = NewEmbedDriver()

func NewEmbedDriver() *eDriver {
	e := &eDriver{
		sorts: make([]uint, 0, 20),
		files: make(map[string]file),
	}
	return e
}

type eDriver struct {
	sorts []uint
	files map[string]file
}

func Add(f fs.FS) {
	embedDriver.add(f)
}

func (e *eDriver) add(f fs.FS) {
	dirEntries, _ := fs.ReadDir(f, ".")
	for _, entry := range dirEntries {
		// 取version
		tmp, err := strconv.ParseUint(strings.Split(entry.Name(), "_")[0], 10, 64)
		if err != nil {
			panic(fmt.Errorf("file name invalid: %w", err))
		}
		version := uint(tmp)
		has := false
		for _, v := range e.sorts {
			if v == version {
				has = true
				break
			}
		}
		if !has {
			e.sorts = append(e.sorts, version)
		}

		//取file
		f := file{
			fs:   f,
			name: entry.Name(),
		}
		key := fmt.Sprintf("%d_%s", version, strings.Split(entry.Name(), ".")[1])
		e.files[key] = f
	}
	sort.Slice(e.sorts, func(i, j int) bool {
		return e.sorts[i] < e.sorts[j]
	})
	return
}

func (e *eDriver) Open(url string) (source.Driver, error) {
	return e, nil
}

func (e eDriver) Close() error {
	return nil
}

func (e eDriver) First() (version uint, err error) {
	if len(e.sorts) < 1 {
		return 0, os.ErrNotExist
	}
	return e.sorts[0], nil
}

func (e eDriver) find(version uint) (index int, err error) {
	var ver uint
	for index, ver = range e.sorts {
		if ver == version {
			return
		}
	}
	return 0, os.ErrNotExist
}

func (e eDriver) Prev(version uint) (prevVersion uint, err error) {
	index, err := e.find(version)
	if err != nil {
		return
	}
	if index-1 >= 0 {
		return e.sorts[index-1], nil
	}
	return 0, os.ErrNotExist
}

func (e eDriver) Next(version uint) (nextVersion uint, err error) {
	index, err := e.find(version)
	if err != nil {
		return
	}
	if index+1 < len(e.sorts) {
		return e.sorts[index+1], nil
	}
	return 0, os.ErrNotExist
}

func (e eDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	file, ok := e.files[fmt.Sprintf("%d_%s", version, "up")]
	if !ok {
		return nil, "", os.ErrNotExist
	}
	identifier = file.name
	r, err = file.Open()
	return
}

func (e eDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	file, ok := e.files[fmt.Sprintf("%d_%s", version, "down")]
	if !ok {
		return nil, "", os.ErrNotExist
	}
	identifier = file.name
	r, err = file.Open()
	return
}

type file struct {
	fs   fs.FS
	name string
}

func (f file) Open() (io.ReadCloser, error) {
	return f.fs.Open(f.name)
}
