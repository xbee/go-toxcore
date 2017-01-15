package toxin

/*
#include "DHT.h"
*/
import "C"
import "unsafe"

type IP4 [4]byte
type IP6 [16]byte
type IP [17]byte
type IPS struct {
	family uint8
	ipx    IP6
}
type IP_Port struct {
	ip   IP
	port uint16
}

const SIZE_IP4 = len(IP4{})
const SIZE_IP6 = len(IP6{})
const SIZE_IP = len(IP{})
const SIZE_PORT = 2
const SIZE_IPPORT = SIZE_IP + SIZE_PORT

const TOX_ENABLE_IPV6_DEFAULT = 1

/* addr_resolve return values */
const TOX_ADDR_RESOLVE_INET = 1
const TOX_ADDR_RESOLVE_INET6 = 2

func (ip IP) ToC() C.IP {
	var cip C.IP
	C.memcpy((unsafe.Pointer)(&cip), ((unsafe.Pointer)(&ip)), C.size_t(SIZE_IP))
	return cip
}

func (ip *IP) Init() {
	C.ip_init((*C.IP)((unsafe.Pointer)(ip)), 0)
}

type NetworkCore struct {
	net *C.Networking_Core
}

func NetworkingAtStartup() int {
	r := C.networking_at_startup()
	return int(r)
}

func SockValid(sock int) int {
	r := C.sock_valid((C.sock_t)(sock))
	return int(r)
}

func SockKill(sock int) {
	C.kill_sock((C.sock_t)(sock))
}

/* Set socket as nonblocking
 *
 * return 1 on success
 * return 0 on failure
 */
func SetSocketNonblock(sock int) int {
	r := C.set_socket_nonblock((C.sock_t)(sock))
	return int(r)
}

/* Set socket to not emit SIGPIPE
 *
 * return 1 on success
 * return 0 on failure
 */
func SetSocketNosigpipe(sock int) int {
	r := C.set_socket_nosigpipe((C.sock_t)(sock))
	return int(r)
}

/* Enable SO_REUSEADDR on socket.
 *
 * return 1 on success
 * return 0 on failure
 */
func SetSocketReuseaddr(sock int) int {
	r := C.set_socket_reuseaddr((C.sock_t)(sock))
	return int(r)
}

/* Set socket to dual (IPv4 + IPv6 socket)
 *
 * return 1 on success
 * return 0 on failure
 */
func SetSocketDualstack(sock int) int {
	r := C.set_socket_dualstack((C.sock_t)(sock))
	return int(r)
}

/* return current monotonic time in milliseconds (ms). */
func CurrentTimeMonotonic() uint64 {
	return uint64(C.current_time_monotonic())
}

/* Basic network functions: */

/* Function to send packet(data) of length length to ip_port. */
// int sendpacket(Networking_Core *net, IP_Port ip_port, const uint8_t *data, uint16_t length);
func (this *NetworkCore) SendPacket() {
}

/* Function to call when packet beginning with byte is received. */
// void networking_registerhandler(Networking_Core *net, uint8_t byte, packet_handler_callback cb, void *object);

/* Call this several times a second. */
func (this *NetworkCore) Poll() { C.networking_poll(this.net) }

/* Initialize networking.
 * bind to ip and port.
 * ip must be in network order EX: 127.0.0.1 = (7F000001).
 * port is in host byte order (this means don't worry about it).
 *
 * return Networking_Core object if no problems
 * return NULL if there are problems.
 *
 * If error is non NULL it is set to 0 if no issues, 1 if socket related error, 2 if other.
 */
func NewNetworkCore(ip IP, port uint16) *NetworkCore {
	// Networking_Core *new_networking(IP ip, uint16_t port);
	this := &NetworkCore{}
	this.net = C.new_networking(ip.ToC(), C.uint16_t(port))

	return this
}
func NewNetworkCoreEx(ip IP, port_from, port_to uint16) (*NetworkCore, uint) {
	// Networking_Core *new_networking_ex(IP ip, uint16_t port_from, uint16_t port_to, unsigned int *error);
	this := &NetworkCore{}
	var error C.uint
	this.net = C.new_networking_ex(ip.ToC(), C.uint16_t(port_from), C.uint16_t(port_to), &error)

	return this, uint(error)
}

/* Function to cleanup networking stuff (doesn't do much right now). */
func (this *NetworkCore) Kill() { C.kill_networking(this.net) }
