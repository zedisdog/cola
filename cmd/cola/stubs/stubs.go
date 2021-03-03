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

//go:embed config.yaml.stub
var ConfigTemp string

//go:embed routes.go.stub
var RoutesTemp string

//go:embed controller.go.stub
var ControllerTemp string

//go:embed model.go.stub
var ModelTemp string

//go:embed docker-compose.yml.stub
var DockerComposeTemp string

//go:embed config.go.stub
var ConfigGoTemp string
