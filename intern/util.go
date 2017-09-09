package toxin

/*
#include <util.h>
*/
import "C"

func host_to_net(n uint16) uint16 {
	return (n << 8) | (n >> 8)
}
func net_to_host(n uint16) uint16 { return host_to_net(n) }

func lendian_to_host16(lendian uint16) uint16 {
	return uint16(C.lendian_to_host16(C.uint16_t(lendian)))
}

func host_tolendian16(lhost uint16) uint16 {
	return lendian_to_host16(lhost)
}

func is_timeout(timestamp uint64, timeout uint64) bool {
	r := C.is_timeout(C.uint64_t(timestamp), C.uint64_t(timeout))
	if r == 0 {
		return false
	}
	return true
}
