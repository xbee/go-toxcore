package toxin

/*
// #include "DHT.h"
// #include "tox.h"
#include "onion.h"

*/
import "C"

// import "unsafe"

type Onion struct {
	o *C.Onion
}

type OnionPath struct {
	op *C.Onion_Path
}

func NewOnion(dht *DHT) *Onion {
	o := C.new_onion(dht.dht)
	this := &Onion{o: o}
	return this
}

func (this *Onion) Kill() {
	C.kill_onion(this.o)
}
