package toxin

import (
	"testing"
	"time"
)

func TestNCN(t *testing.T) {
	ncn := new_NetworkCoreNative()
	t.Log(ncn)

	ncn.sendpacket()
	ncn.poll()
	time.Sleep(5 * time.Second)

}
