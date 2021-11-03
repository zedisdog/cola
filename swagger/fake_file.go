package swagger

import (
	"io"
	"io/fs"
	"time"
)

// fakeFile implements FileLike and also fs.FileInfo.
type fakeFile struct {
	name     string
	contents string
	mode     fs.FileMode
	offset   int
}

// Reset prepares a fakeFile for reuse.
func (f *fakeFile) Reset() *fakeFile {
	f.offset = 0
	return f
}

// FileLike methods.

func (f *fakeFile) Name() string {
	// A bit of a cheat: we only have a basename, so that's also ok for FileInfo.
	return f.name
}

func (f *fakeFile) Stat() (fs.FileInfo, error) {
	return f, nil
}

func (f *fakeFile) Read(p []byte) (int, error) {
	if f.offset >= len(f.contents) {
		return 0, io.EOF
	}
	n := copy(p, f.contents[f.offset:])
	f.offset += n
	return n, nil
}

func (f *fakeFile) Close() error {
	return nil
}

// fs.FileInfo methods.

func (f *fakeFile) Size() int64 {
	return int64(len(f.contents))
}

func (f *fakeFile) Mode() fs.FileMode {
	return f.mode
}

func (f *fakeFile) ModTime() time.Time {
	return time.Time{}
}

func (f *fakeFile) IsDir() bool {
	return false
}

func (f *fakeFile) Sys() interface{} {
	return nil
}
