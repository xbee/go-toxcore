package tox

/*
#include <stdlib.h>
#include <tox/tox.h>
*/
import "C"

const (
	SAVEDATA_TYPE_NONE       = C.TOX_SAVEDATA_TYPE_NONE
	SAVEDATA_TYPE_TOX_SAVE   = C.TOX_SAVEDATA_TYPE_TOX_SAVE
	SAVEDATA_TYPE_SECRET_KEY = C.TOX_SAVEDATA_TYPE_SECRET_KEY
)

type ToxOptions struct {
	Ipv6_enabled  bool
	Udp_enabled   bool
	Proxy_type    int32
	Proxy_address string
	Proxy_port    uint16
	Savedata_type int
	Savedata_data []byte
}

func NewToxOptions() *ToxOptions {
	toxopts := new(C.struct_Tox_Options)
	C.tox_options_default(toxopts)

	opts := new(ToxOptions)
	opts.Ipv6_enabled = bool(toxopts.ipv6_enabled)
	opts.Udp_enabled = bool(toxopts.udp_enabled)
	opts.Proxy_type = int32(toxopts.proxy_type)
	opts.Proxy_port = uint16(toxopts.proxy_port)

	return opts
}

type BootNode struct {
	Addr   string
	Port   int
	Pubkey string
}
