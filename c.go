package tox

/*
#cgo CFLAGS: -g -O2 -std=c99 -Wall
// #cgo LDFLAGS: -ltoxcore -ltoxdns -ltoxav -ltoxencryptsave -lvpx -lopus -lsodium -lm
#cgo pkg-config: libtoxcore libtoxav
*/
import "C"

// TODO what about Windows/MacOS?
