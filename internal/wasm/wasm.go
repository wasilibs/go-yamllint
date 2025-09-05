package wasm

import (
	_ "embed"
)

//go:embed python.wasm
var Python []byte

//go:embed pysite.zip.wasm
var Site []byte
