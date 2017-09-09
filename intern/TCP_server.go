package toxin

/*
#include "TCP_server.h"
*/
import "C"

// import "unsafe"

type TCPServer struct {
	s *C.TCP_Server
}

func NewTCPServer(ipv6_enabled bool, ports []uint16, secret_key []byte, onion *Onion) *TCPServer {
	this := &TCPServer{}
	ports_ := (*C.uint16_t)(&ports[0])
	secret_key_ := (*C.uint8_t)(&secret_key[0])
	this.s = C.new_TCP_server(1, C.uint16_t(len(ports)), ports_, secret_key_, onion.o)
	return this
}

func (this *TCPServer) Do()   { C.do_TCP_server(this.s) }
func (this *TCPServer) Kill() { C.kill_TCP_server(this.s) }
