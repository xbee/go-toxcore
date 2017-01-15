package toxin

/*
// #include "DHT.h"
// #include "tox.h"
#include "onion_announce.h"

*/
import "C"

type OnionAnnounce struct {
	oa *C.Onion_Announce
}

func NewOnionAnnounce(dht *DHT) *OnionAnnounce {
	oa := C.new_onion_announce(dht.dht)
	this := &OnionAnnounce{oa: oa}
	return this
}

func (this *OnionAnnounce) Kill() {
	C.kill_onion_announce(this.oa)
}
