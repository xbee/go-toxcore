package toxin

import (
	"log"
	"net"
	"syscall"
)

type packet_handler_callbacks func(object interface{}, ip_port IP_Port, data []byte, userdata interface{})

type PacketHandlers struct {
	fn     packet_handler_callbacks
	object interface{}
}

type Socket int
type NetworkCoreNative struct {
	/*
	   Logger *log;
	   Packet_Handles packethandlers[256];

	   sa_family_t family;
	   uint16_t port;
	   //* Our UDP socket.
	   Socket sock;
	*/

	packetHandlers [256]PacketHandlers
	family         uint16
	port           int
	sock           int
}

// bind port
func new_NetworkCoreNative() *NetworkCoreNative {
	port := 33556
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		log.Println(err)
	}
	this := &NetworkCoreNative{}
	this.sock = sock

	err = syscall.SetNonblock(sock, true)
	if err != nil {
		log.Println(err)
	}

	var sa syscall.SockaddrInet4
	sa.Port = port
	sa.Addr = [4]byte{}
	err = syscall.Bind(this.sock, &sa)
	if err != nil {
		log.Println(err)
	}

	this.port = port
	return this
}

func (this *NetworkCoreNative) sendpacket() {
	var tosa syscall.SockaddrInet4
	tosa.Port = 33445
	ip := net.ParseIP("127.0.0.1")
	tosa.Addr = [4]byte{ip[0], ip[1], ip[2], ip[3]}

	err := syscall.Sendto(this.sock, []byte{0, 1, 2}, 0, &tosa)
	if err != nil {
		log.Println(err)
	}
}

func (this *NetworkCoreNative) poll() {
	p := make([]byte, 2000)
	log.Println("recvfrom...")
	n, from, err := syscall.Recvfrom(this.sock, p, 0)
	log.Println("recvfrom done")
	if err != nil {
		log.Println(err, n, from)
		if err == syscall.EAGAIN {
			log.Println(syscall.EAGAIN, syscall.EAGAIN.Temporary())
		}
	} else {
		var ip_port IP_Port
		this.packetHandlers[0].fn(this.packetHandlers[0].object, ip_port, p, nil)
	}
}
