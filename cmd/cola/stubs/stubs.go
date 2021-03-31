package stubs

import (
	_ "embed"
)

//go:embed main.go.stub
var MainTemp string

//go:embed log.go.stub
var LogTemp string

//go:embed db.go.stub
var DbTemp string

//go:embed config.yaml
var ConfigTemp string

//go:embed routes.go.stub
var RoutesTemp string

//go:embed controller.go.stub
var ControllerTemp string

//go:embed model.go.stub
var ModelTemp string

//go:embed docker-compose.yml
var DockerComposeTemp string

//go:embed config.go.stub
var ConfigGoTemp string

//go:embed dao.go.stub
var DaoTemp string

//go:embed service.go.stub
var ServiceTemp string

//go:embed storage.go.stub
var StorageTemp string

//go:embed air.conf
var AirConf string
