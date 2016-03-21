package tox

type ToxOptions struct {
	Ipv6_enabled  bool
	Udp_enabled   bool
	Proxy_type    int32
	Proxy_address string
	Proxy_port    uint16
}

func NewToxOptions() *ToxOptions {
	return &ToxOptions{}
}
