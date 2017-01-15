package toxin

/*
#cgo CFLAGS: -g -O2 -std=c99 -Wall
#cgo CFLAGS: -I/path/to/toxcore/toxcore
#cgo pkg-config: libtoxcore
*/
import "C"

// install
// CGO_CFLAGS="-I/path/to/toxcore/toxcore" go install
// TODO what about Windows/MacOS?
