package toxin

import (
	"testing"
)

func TestNewNetwork(t *testing.T) {
	var ip IP
	rc := NetworkingAtStartup()
	t.Log(rc)
	(&ip).Init()

	net := NewNetworkCore(ip, 12345)
	t.Log(net)
}
