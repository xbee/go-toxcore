package tox

/*
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>
#include <tox/toxav.h>
*/
import "C"

// import "unsafe"

type ToxAV struct {
	tox   *Tox
	toxav *C.ToxAV
}

func NewToxAV(tox *Tox) *ToxAV {
	tav := new(ToxAV)
	tav.tox = tox

	var cerr C.TOXAV_ERR_NEW
	tav.toxav = C.toxav_new(tox.toxcore, &cerr)
	if cerr != 0 {
	}

	return tav
}
