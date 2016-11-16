package tox

import (
	"encoding/hex"
	"strconv"
	"strings"
	"testing"
	"time"
)

var bsnodes = []string{
	"biribiri.org", "33445", "F404ABAA1C99A9D37D61AB54898F56793E1DEF8BD46B1038B9D822E8460FAB67",
	"178.62.250.138", "33445", "788236D34978D1D5BD822F0A5BEBD2C53C64CC31CD3149350EE27D4D9A2F9B6B",
	"198.98.51.198", "33445", "1D5A5F2F5D6233058BF0259B09622FB40B482E4FA0931EB8FD3AB8E7BF7DAF6F",
}

func TestCreate(t *testing.T) {
	t.Run("no options", func(t *testing.T) {
		_t := NewTox(nil)
		if _t == nil {
			t.Error("nil")
		}
		_t.Kill()
	})
	t.Run("default options", func(t *testing.T) {
		opts := NewToxOptions()
		_t := NewTox(opts)
		if _t == nil {
			t.Error("nil")
		}
		_t.Kill()
	})
	t.Run("tcp options", func(t *testing.T) {
		opts := NewToxOptions()
		opts.Tcp_port = 34567
		_t := NewTox(opts)
		if _t == nil {
			t.Error("nil")
		}
		_t.Kill()
	})
	t.Run("tcp conflict", func(t *testing.T) {
		opts := NewToxOptions()
		opts.Tcp_port = 34567
		_t, _t2 := NewTox(opts), NewTox(opts)
		if _t == nil || _t2 != nil {
			t.Error("should non-nil/nil", _t, _t2)
		}
		_t.Kill()
		_t2.Kill()
	})
	t.Run("save profile", func(t *testing.T) {
		_t := NewTox(nil)
		sz := _t.GetSavedataSize()
		dat := _t.GetSavedata()
		if sz <= 0 || dat == nil || len(dat) != int(sz) {
			t.Error("cannot zero")
		}
		_t.Kill()
	})
	t.Run("load profile", func(t *testing.T) {
		_t := NewTox(nil)
		dat := _t.GetSavedata()
		_t.Kill()

		opts := NewToxOptions()
		opts.Savedata_data = dat
		opts.Savedata_type = SAVEDATA_TYPE_TOX_SAVE
		_t2 := NewTox(opts)
		dat2 := _t2.GetSavedata()
		if len(dat2) != len(dat) || string(dat2) != string(dat) {
			t.Error("must ==")
		}
		_t2.Kill()
	})
	t.Run("load error profile", func(t *testing.T) {
		_t := NewTox(nil)
		dat := _t.GetSavedata()
		addr := _t.SelfGetAddress()
		_t.Kill()

		opts := NewToxOptions()
		opts.Savedata_data = append([]byte("set-broken"), dat...)
		opts.Savedata_type = SAVEDATA_TYPE_TOX_SAVE
		_t2 := NewTox(opts)
		if _t2 == nil {
			t.Error("must non-nil")
		}
		if addr == _t2.SelfGetAddress() {
			t.Error("must !=", addr, _t2.SelfGetAddress())
		}
		_t2.Kill()
	})
	t.Run("load seckey", func(t *testing.T) {
		_t := NewTox(nil)
		addr := _t.SelfGetAddress()
		seckey := _t.SelfGetSecretKey()
		_t.Kill()

		opts := NewToxOptions()
		opts.Savedata_type = SAVEDATA_TYPE_SECRET_KEY
		binsk, _ := hex.DecodeString(seckey)
		opts.Savedata_data = binsk
		_t2 := NewTox(opts)
		if _t2.SelfGetSecretKey() != seckey {
			t.Error("must =")
		}
		if _t2.SelfGetAddress()[0:PUBLIC_KEY_SIZE*2] != addr[0:PUBLIC_KEY_SIZE*2] {
			t.Error("must =", _t2.SelfGetAddress(), addr)
		}
	})
	t.Run("destroy", func(t *testing.T) {
		_t := NewTox(nil)
		_t.Kill()
		if _t.toxcore != nil {
			t.Error("must nil")
		}
	})
}

func TestBase(t *testing.T) {
	_t := NewTox(nil)
	defer _t.Kill()

	t.Run("name", func(t *testing.T) {
		if _t.SelfGetName() != "" {
			t.Error("must empty")
		}
		if _t.SelfGetNameSize() != 0 {
			t.Error("must zero")
		}
		tname := "test name"
		if err := _t.SelfSetName(tname); err != nil {
			t.Error(err)
		}
		if size := _t.SelfGetNameSize(); size != len(tname) {
			t.Error("must =", size, len(tname))
		}
		tname = strings.Repeat("n", MAX_NAME_LENGTH)
		if err := _t.SelfSetName(tname); err != nil {
			t.Error(err)
		}
		tname = strings.Repeat("n", MAX_NAME_LENGTH+1)
		if err := _t.SelfSetName(tname); err == nil {
			t.Error("must failed", err)
		}
	})
	t.Run("local status", func(t *testing.T) {
		if _t.SelfGetStatusMessageSize() != 0 {
			t.Error("must zero")
		}
		if stm, err := _t.SelfGetStatusMessage(); err != nil || len(stm) != 0 {
			t.Error("must empty", stm, err)
		}
		tmsg := "test status msg"
		if ok, err := _t.SelfSetStatusMessage(tmsg); !ok || err != nil {
			t.Error("must ok", err)
		}
		if stm, err := _t.SelfGetStatusMessage(); err != nil || stm != tmsg {
			t.Error("must =", stm, err)
		}
		tmsg = strings.Repeat("s", MAX_STATUS_MESSAGE_LENGTH)
		if ok, err := _t.SelfSetStatusMessage(tmsg); !ok || err != nil {
			t.Error("must ok", err)
		}
		tmsg = strings.Repeat("s", MAX_STATUS_MESSAGE_LENGTH+1)
		if ok, err := _t.SelfSetStatusMessage(tmsg); ok || err == nil {
			t.Error("must failed", err)
		}
		if _t.SelfGetConnectionStatus() != CONNECTION_NONE {
			t.Error("must none")
		}
	})
	t.Run("address/pubkey", func(t *testing.T) {
		addr := _t.SelfGetAddress()
		if len(addr) != ADDRESS_SIZE*2 {
			t.Error("size")
		}
		pubkey := _t.SelfGetPublicKey()
		if len(pubkey) != PUBLIC_KEY_SIZE*2 {
			t.Error("size")
		}
		if addr[0:len(pubkey)] != pubkey {
			t.Error(addr)
		}
	})
	t.Run("seckey", func(t *testing.T) {
		seckey := _t.SelfGetSecretKey()
		if len(seckey) != SECRET_KEY_SIZE*2 {
			t.Error("size")
		}
	})
	t.Run("nospam", func(t *testing.T) {
	})
}

func TestBootstrap(t *testing.T) {
	_t := NewTox(nil)
	defer _t.Kill()
	port, _ := strconv.Atoi(bsnodes[1])

	t.Run("success", func(t *testing.T) {
		if ok, err := _t.Bootstrap(bsnodes[0], uint16(port), bsnodes[2]); !ok || err != nil {
			t.Error("must ok", ok, err)
		}
	})
	t.Run("failed", func(t *testing.T) {
		brkey := bsnodes[2]
		brkey = "XYZAB" + bsnodes[2][3:]
		if ok, err := _t.Bootstrap(bsnodes[0], uint16(port), brkey); ok || err == nil {
			t.Error("must failed", ok, err)
		}
		if ok, err := _t.Bootstrap("a.b.c.d", uint16(port), bsnodes[2]); ok || err == nil {
			t.Error("must failed", ok, err)
		}
	})
	t.Run("relay", func(t *testing.T) {
		if ok, err := _t.AddTcpRelay(bsnodes[0], uint16(port), bsnodes[2]); !ok || err != nil {
			t.Error("must ok", ok, err)
		}
		if ok, err := _t.AddTcpRelay("a.b.c.d", uint16(port), bsnodes[2]); ok || err == nil {
			t.Error("must failed", ok, err)
		}
	})
}

// login udp / login tcp
func TestLogin(t *testing.T) {
	_t := NewTox(nil)
	defer _t.Kill()

	t.Run("default", func(t *testing.T) {

	})
	stopch := make(chan struct{}, 0)
	_t.CallbackSelfConnectionStatus(func(_t *Tox, status uint32, d interface{}) {
		t.Log(status)
		stopch <- struct{}{}
	}, nil)

	go func() {
		interval := _t.IterationInterval()
		tickch := time.Tick(time.Duration(interval) * time.Millisecond)
		for {
			select {
			case <-tickch:
				_t.Iterate()
			}
		}
	}()

	t.Log("waiting...")
	<-stopch
}

type MiniTox struct {
	t *Tox
}

func NewMiniTox() *MiniTox {
	this := &MiniTox{}
	this.t = NewTox(nil)
	return this
}

func NewMiniTox2() *MiniTox {
	this := &MiniTox{}
	opts := NewToxOptions()
	this.t = NewTox(opts)
	return this
}

func (this *MiniTox) Iterate() {
	tickch := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-tickch:
			this.t.Iterate()
		}
	}
}

func (this *MiniTox) bootstrap() {
	for idx := 0; idx < len(bsnodes)/3; idx++ {
		port, err := strconv.Atoi(bsnodes[1+idx*3])
		_, err = this.t.Bootstrap(bsnodes[0+idx*3], uint16(port), bsnodes[2+idx*3])
		if err != nil {
		}
		_, err = this.t.AddTcpRelay(bsnodes[0+idx*3], uint16(port), bsnodes[2+idx*3])
		if err != nil {
		}
	}
}

var err error

func waitcond(cond func() bool) {
	// TODO might infinite loop
	cnter := 0
	for {
		if cond() {
			print("\n")
			return
		}

		if cnter%15 == 0 {
			// print(".")
		}
		cnter += 1
		time.Sleep(51 * time.Millisecond)
	}
}

func _TestFriend(t *testing.T) {
	stopch := make(chan struct{}, 0)
	t1 := NewMiniTox()
	t2 := NewMiniTox()
	// t1.bootstrap()
	// t2.bootstrap()

	t1.t.CallbackFriendRequest(func(_ *Tox, friendId, msg string, d interface{}) {
		println("got request", friendId, msg)
		_, err := t1.t.FriendAddNorequest(friendId)
		if err != nil {
			t.Fail()
		}
	}, nil)
	t1.t.CallbackFriendMessage(func(_ *Tox, friendNumber uint32, msg string, d interface{}) {
		println(friendNumber, msg)
	}, nil)
	t1.t.CallbackFriendConnectionStatus(func(_ *Tox, friendNumber uint32, status uint32,
		d interface{}) {
		println(friendNumber, status)
	}, nil)
	t1nameChanged := false
	t2.t.CallbackFriendName(func(_ *Tox, friendNumber uint32, name string, d interface{}) {
		t.Log(friendNumber, name)
		if len(name) > 0 {
			t1nameChanged = true
		}
	}, nil)
	t1statusMessageChanged := false
	t2.t.CallbackFriendStatusMessage(func(_ *Tox, friendNumber uint32, stmsg string, d interface{}) {
		t.Log(friendNumber, stmsg)
		if len(stmsg) > 0 {
			t1statusMessageChanged = true
		}
	}, nil)

	go t1.Iterate()
	go t2.Iterate()

	go func() {
		waitcond(func() bool {
			return t1.t.SelfGetConnectionStatus() == 2 && t2.t.SelfGetConnectionStatus() == 2
		})
		friendNumber, err := t2.t.FriendAdd(t1.t.SelfGetAddress(), "hoho")
		if err != nil {
			println(err)
			t.Fail()
		}
		_, err = t2.t.FriendAdd(t1.t.SelfGetAddress(), "hehe")
		if err == nil {
			t.Fail()
		}
		if t2.t.SelfGetFriendListSize() != 1 {
			t.Fail()
		}
		lst := t2.t.SelfGetFriendList()
		if len(lst) != 1 {
			t.Fail()
		}

		friendNumber2, err := t2.t.FriendByPublicKey(t1.t.SelfGetAddress())
		if err != nil {
			t.Fail()
		}
		if friendNumber2 != friendNumber {
			t.Fail()
		}
		pubkey, err := t2.t.FriendGetPublicKey(friendNumber)
		if err != nil {
			t.Fail()
		}
		if pubkey != t1.t.SelfGetPublicKey() {
			t.Fail()
		}
		if !t2.t.FriendExists(friendNumber) {
			t.Fail()
		}

		println("wait frined count", friendNumber, t2.t.SelfGetFriendListSize())
		waitcond(func() bool {
			return 1 == t1.t.SelfGetFriendListSize()
		})
		waitcond(func() bool {
			status, err := t2.t.FriendGetConnectionStatus(friendNumber)
			if err != nil {
				return false
			}
			return status == 2
		})
		_, err = t2.t.FriendSendMessage(friendNumber, "foo")
		if err != nil {
			t.Fail()
		}
		_, err = t2.t.FriendSendAction(friendNumber, "actfoo")
		if err != nil {
			t.Fail()
		}

		err = t1.t.SelfSetName("t1")
		if err != nil {
			t.Fail()
		}
		waitcond(func() bool { return t1nameChanged })
		t1name, err := t2.t.FriendGetName(friendNumber)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if t1name != "t1" {
			t.Error(t1name)
			t.Fail()
		}
		_, err = t1.t.SelfSetStatusMessage("t1status")
		if err != nil {
			t.Error(err)
		}
		waitcond(func() bool { return t1statusMessageChanged })
		t1stmsg, err := t2.t.FriendGetStatusMessage(friendNumber)
		if err != nil {
			t.Error(err)
		}
		if t1stmsg != "t1status" {
			t.Error(t1stmsg, t1stmsg != "t1status")
		}
		t1stmsgsz, err := t2.t.FriendGetStatusMessageSize(friendNumber)
		if err != nil {
			t.Error(err)
		}
		if t1stmsgsz != len("t1status") {
			t.Error(t1stmsgsz, len("t1status"))
		}

		t1st, err := t2.t.FriendGetStatus(friendNumber)
		if err != nil {
			t.Error(err)
		}
		if t1st != uint8(USER_STATUS_NONE) {
			t.Error(t1st)
		}

		_, err = t2.t.FriendDelete(friendNumber)
		if err != nil {
			t.Fail()
		}
		if t2.t.FriendExists(friendNumber) {
			t.Fail()
		}

		stopch <- struct{}{}
	}()

	t.Log("waiting...")
	<-stopch
}
