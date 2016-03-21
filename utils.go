package tox

/*
#cgo CFLAGS: -g -O2 -Wall
#include <stdlib.h>
#include <stdint.h>

static uint8_t *char2uint8(char *s) { return (uint8_t*)s; }

*/
import "C"

// *C.char ==> *C.uint8_t
func char2uint8(s *C.char) *C.uint8_t {
	return C.char2uint8(s)
}

