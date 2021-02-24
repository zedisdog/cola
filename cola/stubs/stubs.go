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
