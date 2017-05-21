package tox

/*
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>
*/
import "C"
import "unsafe"

const (
	SAVEDATA_TYPE_NONE       = C.TOX_SAVEDATA_TYPE_NONE
	SAVEDATA_TYPE_TOX_SAVE   = C.TOX_SAVEDATA_TYPE_TOX_SAVE
	SAVEDATA_TYPE_SECRET_KEY = C.TOX_SAVEDATA_TYPE_SECRET_KEY
)

const (
	PROXY_TYPE_NONE   = C.TOX_PROXY_TYPE_NONE
	PROXY_TYPE_HTTP   = C.TOX_PROXY_TYPE_HTTP
	PROXY_TYPE_SOCKS5 = C.TOX_PROXY_TYPE_SOCKS5
)

type ToxOptions struct {
	Ipv6_enabled  bool
	Udp_enabled   bool
	Proxy_type    int32
	Proxy_host    string
	Proxy_port    uint16
	Savedata_type int
	Savedata_data []byte
	Tcp_port      uint16
	ThreadSafe    bool
}

func NewToxOptions() *ToxOptions {
	toxopts := new(C.struct_Tox_Options)
	C.tox_options_default(toxopts)

	opts := new(ToxOptions)
	opts.Ipv6_enabled = bool(toxopts.ipv6_enabled)
	opts.Udp_enabled = bool(toxopts.udp_enabled)
	opts.Proxy_type = int32(toxopts.proxy_type)
	opts.Proxy_port = uint16(toxopts.proxy_port)
	opts.Tcp_port = uint16(toxopts.tcp_port)

	return opts
}

func (this *ToxOptions) toCToxOptions() *C.struct_Tox_Options {
	toxopts := new(C.struct_Tox_Options)
	C.tox_options_default(toxopts)
	toxopts.ipv6_enabled = (C._Bool)(this.Ipv6_enabled)
	toxopts.udp_enabled = (C._Bool)(this.Udp_enabled)

	if this.Savedata_data != nil {
		toxopts.savedata_data = pointer2uint8(C.malloc(C.size_t(len(this.Savedata_data))))
		C.memcpy(unsafe.Pointer(toxopts.savedata_data),
			unsafe.Pointer(&this.Savedata_data[0]), C.size_t(len(this.Savedata_data)))
		toxopts.savedata_length = C.size_t(len(this.Savedata_data))
		toxopts.savedata_type = C.TOX_SAVEDATA_TYPE(this.Savedata_type)
	}
	toxopts.tcp_port = (C.uint16_t)(this.Tcp_port)

	toxopts.proxy_type = C.TOX_PROXY_TYPE(this.Proxy_type)
	toxopts.proxy_port = C.uint16_t(this.Proxy_port)
	if len(this.Proxy_host) > 0 {
		toxopts.proxy_host = C.CString(this.Proxy_host)
	}

	return toxopts
}

type BootNode struct {
	Addr   string
	Port   int
	Pubkey string
}
