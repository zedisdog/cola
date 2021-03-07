package storage

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/zedisdog/cola/storage/drivers"
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

func (s Storage) GetMime(path string) string {
	if ss, ok := interface{}(s.driver).(DriverHasMime); ok {
		return ss.GetMime(path)
	}
	panic(errors.New("driver is not implement interface <DriverHasMime>"))
}

func (s Storage) GetUrl(path string) string {
	if ss, ok := interface{}(s.driver).(DriverHasUrl); ok {
		return ss.GetUrl(path)
	}
	panic(errors.New("driver is not implement interface <DriverHasUrl>"))
}

func (s Storage) Path(path string) string {
	if ss, ok := interface{}(s.driver).(DriverHasPath); ok {
		return ss.Path(path)
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

type DriverHasUrl interface {
	GetUrl(path string) string
}

type DriverHasPath interface {
	Path(path string) string
}

// NewByViper 从viper中获取配置
//   storage:
//     path: ./storage
//     driver: local
func NewByViper(v *viper.Viper) *Storage {
	var driver Driver
	switch v.GetString("storage.driver") {
	case "local":
		driver = drivers.NewLocal(v.GetString("storage.path"), v.GetString("storage.url"))
	}
	return &Storage{
		driver: driver,
	}
}
