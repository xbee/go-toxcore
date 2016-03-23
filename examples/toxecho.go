package main

import (
	"fmt"
	"io/ioutil"
	"log"
	// "os"
	"time"
	"unsafe"

	"tox"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

var server = []interface{}{
	"205.185.116.116", uint16(33445), "A179B09749AC826FF01F37A9613F6B57118AE014D4196A0E1105A98F93A54702",
}
var fname = "./toxecho.data"
var debug = false

func main() {
	opt := tox.NewToxOptions()
	if tox.FileExist(fname) {
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Println(err)
		} else {
			opt.Savedata_data = data
			opt.Savedata_type = tox.SAVEDATA_TYPE_TOX_SAVE
		}
	}
	t := tox.NewTox(opt)

	r, err := t.Bootstrap(server[0].(string), server[1].(uint16), server[2].(string))
	r2, err := t.AddTcpRelay(server[0].(string), server[1].(uint16), server[2].(string))
	if debug {
		log.Println("bootstrap:", r, err, r2)
	}

	sz := t.GetSavedataSize()
	sd := t.GetSavedata()
	if debug {
		log.Println("savedata:", sz, t)
		log.Println("savedata", len(sd), t)
	}
	t.WriteSavedata(fname)

	pubkey := t.SelfGetPublicKey()
	seckey := t.SelfGetSecretKey()
	toxid := t.SelfGetAddress()
	if debug {
		log.Println("keys:", pubkey, seckey, len(pubkey), len(seckey))
	}
	log.Println("toxid:", toxid)

	// callbacks
	t.CallbackSelfConnectionStatus(func(t *tox.Tox, status uint32, userData unsafe.Pointer) {
		if debug {
			log.Println(status, userData)
		}
	}, nil)
	t.CallbackFriendRequest(func(t *tox.Tox, friendId string, message string, userData unsafe.Pointer) {
		log.Println(friendId, message)
		num, err := t.FriendAddNorequest(friendId)
		if debug {
			log.Println(num, err)
		}
		if num < 100000 {
			t.WriteSavedata(fname)
		}
	}, nil)

	// loops
	shutdown := false
	loopc := 0
	for !shutdown {
		t.Iterate()
		status := t.SelfGetConnectionStatus()
		if loopc%5500 == 0 {
			if status == 0 {
				if debug {
					fmt.Print(".")
				}
			} else {
				if debug {
					fmt.Print(status, ",")
				}
			}
		}
		loopc += 1
		time.Sleep(50 * time.Microsecond)
	}

	t.Kill()
}

func _dirty_init() {
	log.Println("ddddddddd")
	tox.KeepPkg()
}
