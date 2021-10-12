package stubs

import (
	"embed"
	_ "embed"
)

//go:embed controller.go.stub
var ControllerTemp string

//go:embed model.go.stub
var ModelTemp string

//go:embed dao.go.stub
var DaoTemp string

//go:embed service.go.stub
var ServiceTemp string

//go:embed template/*
var Template embed.FS
