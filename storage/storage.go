package storage

import (
	"errors"
	"io"
	"mime/multipart"
)

func New(driver Driver) *Storage {
	return &Storage{
		driver: driver,
	}
}

type Storage struct {
	driver Driver
}

func (s Storage) Put(path string, data []byte) error {
	return s.driver.Put(path, data)
}

func (s Storage) Get(path string) ([]byte, error) {
	return s.driver.Get(path)
}

func (s Storage) Remove(path string) error {
	return s.driver.Remove(path)
}

func (s Storage) PutString(path string, data string) error {
	return s.driver.Put(path, []byte(data))
}

func (s Storage) GetString(path string) (data string, err error) {
	tmp, err := s.driver.Get(path)
	if err != nil {
		return
	}
	data = string(tmp)
	return
}

func (s Storage) PutFile(path string, file *multipart.FileHeader) (err error) {
	fp, err := file.Open()
	if err != nil {
		return
	}
	defer fp.Close()
	data, err := io.ReadAll(fp)
	if err != nil {
		return
	}
	return s.Put(path, data)
}

func (s Storage) GetMime(path string) string {
	if ss, ok := interface{}(s.driver).(DriverHasMime); ok {
		return ss.GetMime(path)
	}
	panic(errors.New("driver is not implement interface <DriverHasMime>"))
}

func (s Storage) Path(path string) string {
	if ss, ok := interface{}(s.driver).(DriverHasPath); ok {
		return ss.Path(path)
	}
	panic(errors.New("driver is not implement interface <DriverHasPath>"))
}

func (s Storage) Base64(path string) (string, error) {
	if ss, ok := interface{}(s.driver).(DriverHasBase64); ok {
		return ss.Base64(path)
	}
	panic(errors.New("driver is not implement interface <DriverHasPath>"))
}

type Driver interface {
	Put(path string, data []byte) error
	Get(path string) ([]byte, error)
	Remove(path string) error
}

type DriverHasMime interface {
	GetMime(path string) string
}

type DriverHasPath interface {
	Path(path string) string
}

type DriverHasBase64 interface {
	Base64(path string) (string, error)
}
