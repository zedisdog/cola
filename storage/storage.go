package storage

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/zedisdog/cola/errx"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

var drivers = make(map[string]IDriver)

func SetDriver(name string, driver IDriver) {
	drivers[name] = driver
}

func RemoveDriver(name string) {
	delete(drivers, name)
}

func Driver(name string) *Storage {
	if driver, ok := drivers[name]; ok {
		return New(driver)
	} else {
		panic(errx.New("driver not found."))
	}
}

func defaultDriver() *Storage {
	if len(drivers) < 1 {
		panic(errx.New("no driver found."))
	}
	for _, v := range drivers {
		return New(v)
	}
	return nil
}

func Put(path string, data []byte) error {
	return defaultDriver().Put(path, data)
}

func Get(path string) ([]byte, error) {
	return defaultDriver().Get(path)
}

func Remove(path string) error {
	return defaultDriver().Remove(path)
}

func PutString(path string, data string) error {
	return defaultDriver().PutString(path, data)
}

func GetString(path string) (data string, err error) {
	return defaultDriver().GetString(path)
}

func PutFile(path string, file *multipart.FileHeader) (err error) {
	return defaultDriver().PutFile(path, file)
}

//PutFileQuick is similar than PutFile, but don't set filename.
func PutFileQuick(file *multipart.FileHeader, directory string) (path string, err error) {
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	//path = directory/xxxx.jpg
	path = fmt.Sprintf(
		"%s%s%s.%s",
		strings.Trim(directory, "\\/"),
		string(os.PathSeparator),
		id.String(),
		filepath.Ext(file.Filename),
	)
	err = defaultDriver().PutFile(path, file)
	return
}

func GetMime(path string) string {
	return defaultDriver().GetMime(path)
}

func Path(path string) string {
	return defaultDriver().Path(path)
}

func Base64(path string) (string, error) {
	return defaultDriver().Base64(path)
}

func New(driver IDriver) *Storage {
	return &Storage{
		driver: driver,
	}
}

type Storage struct {
	driver IDriver
}

func (s *Storage) Put(path string, data []byte) error {
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

type IDriver interface {
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
