package toxin

/*
#include "LAN_discovery.h"
*/
import "C"

func (this *DHT) LANdiscoverySend(port uint16) {
	C.send_LANdiscovery(C.uint16_t(port), this.dht)
}
func (this *DHT) LANdiscoveryInit() { C.LANdiscovery_init(this.dht) }
func (this *DHT) LANdiscoveryKill() { C.LANdiscovery_kill(this.dht) }
