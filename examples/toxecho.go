package main

import (
	"fmt"
	"io/ioutil"
	"log"
	// "os"
	"math/rand"
	"strconv"
	"strings"
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
var nickPrefix = "EchoBot."
var statusText = "Send me text, file, audio, video."

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
	opt.Tcp_port = 33445
	t := tox.NewTox(opt)

	r, err := t.Bootstrap(server[0].(string), server[1].(uint16), server[2].(string))
	r2, err := t.AddTcpRelay(server[0].(string), server[1].(uint16), server[2].(string))
	if debug {
		log.Println("bootstrap:", r, err, r2)
	}

	pubkey := t.SelfGetPublicKey()
	seckey := t.SelfGetSecretKey()
	toxid := t.SelfGetAddress()
	if debug {
		log.Println("keys:", pubkey, seckey, len(pubkey), len(seckey))
	}
	log.Println("toxid:", toxid)

	defaultName, err := t.SelfGetName()
	humanName := nickPrefix + toxid[0:5]
	if humanName != defaultName {
		t.SelfSetName(humanName)
	}
	humanName, err = t.SelfGetName()
	if debug {
		log.Println(humanName, defaultName, err)
	}

	defaultStatusText, err := t.SelfGetStatusMessage()
	if defaultStatusText != statusText {
		t.SelfSetStatusMessage(statusText)
	}
	if debug {
		log.Println(statusText, defaultStatusText, err)
	}

	sz := t.GetSavedataSize()
	sd := t.GetSavedata()
	if debug {
		log.Println("savedata:", sz, t)
		log.Println("savedata", len(sd), t)
	}
	err = t.WriteSavedata(fname)
	if debug {
		log.Println("savedata write:", err)
	}

	// callbacks
	t.CallbackSelfConnectionStatus(func(t *tox.Tox, status uint32, userData unsafe.Pointer) {
		if debug {
			log.Println("on self conn status:", status, userData)
		}
	}, nil)
	t.CallbackFriendRequest(func(t *tox.Tox, friendId string, message string, userData unsafe.Pointer) {
		log.Println(friendId, message)
		num, err := t.FriendAddNorequest(friendId)
		if debug {
			log.Println("on friend request:", num, err)
		}
		if num < 100000 {
			t.WriteSavedata(fname)
		}
	}, nil)
	t.CallbackFriendMessage(func(t *tox.Tox, friendNumber uint32, message string, userData unsafe.Pointer) {
		if debug {
			log.Println("on friend message:", friendNumber, message)
		}
		n, err := t.FriendSendMessage(friendNumber, "Re: "+message)
		if err != nil {
			log.Println(n, err)
		}
	}, nil)
	t.CallbackFriendConnectionStatus(func(t *tox.Tox, friendNumber uint32, status uint32, userData unsafe.Pointer) {
		if debug {
			friendId, err := t.FriendGetPublicKey(friendNumber)
			log.Println("on friend connection status:", friendNumber, status, friendId, err)
		}
	}, nil)
	t.CallbackFriendStatus(func(t *tox.Tox, friendNumber uint32, status uint8, userData unsafe.Pointer) {
		if debug {
			friendId, err := t.FriendGetPublicKey(friendNumber)
			log.Println("on friend status:", friendNumber, status, friendId, err)
		}
	}, nil)
	t.CallbackFriendStatusMessage(func(t *tox.Tox, friendNumber uint32, statusText string, userData unsafe.Pointer) {
		if debug {
			friendId, err := t.FriendGetPublicKey(friendNumber)
			log.Println("on friend status text:", friendNumber, statusText, friendId, err)
		}
	}, nil)

	// some vars for file echo
	var recvFiles = make(map[uint64]uint32, 0)
	var sendFiles = make(map[uint64]uint32, 0)
	var sendDatas = make(map[string][]byte, 0)
	var chunkReqs = make([]string, 0)
	trySendChunk := func(friendNumber uint32, fileNumber uint32, position uint64) {
		sentKeys := make(map[string]bool, 0)
		for _, reqkey := range chunkReqs {
			lst := strings.Split(reqkey, "_")
			pos, err := strconv.ParseUint(lst[2], 10, 64)
			if err != nil {
			}
			if data, ok := sendDatas[reqkey]; ok {
				r, err := t.FileSendChunk(friendNumber, fileNumber, pos, data)
				if err != nil {
					if err.Error() == "toxcore error: 7" || err.Error() == "toxcore error: 8" {
					} else {
						log.Println("file send chunk error:", err, r, reqkey)
					}
					break
				} else {
					delete(sendDatas, reqkey)
					sentKeys[reqkey] = true
				}
			}
		}
		leftChunkReqs := make([]string, 0)
		for _, reqkey := range chunkReqs {
			if _, ok := sentKeys[reqkey]; !ok {
				leftChunkReqs = append(leftChunkReqs, reqkey)
			}
		}
		chunkReqs = leftChunkReqs
	}
	if trySendChunk != nil {
	}

	t.CallbackFileRecvControl(func(t *tox.Tox, friendNumber uint32, fileNumber uint32,
		control int, userData unsafe.Pointer) {
		if debug {
			friendId, err := t.FriendGetPublicKey(friendNumber)
			log.Println("on recv file control:", friendNumber, fileNumber, control, friendId, err)
		}
		key := uint64(uint64(friendNumber)<<32 | uint64(fileNumber))
		if control == tox.FILE_CONTROL_RESUME {
			if fno, ok := sendFiles[key]; ok {
				t.FileControl(friendNumber, fno, tox.FILE_CONTROL_RESUME)
			}
		} else if control == tox.FILE_CONTROL_PAUSE {
			if fno, ok := sendFiles[key]; ok {
				t.FileControl(friendNumber, fno, tox.FILE_CONTROL_PAUSE)
			}
		} else if control == tox.FILE_CONTROL_CANCEL {
			if fno, ok := sendFiles[key]; ok {
				t.FileControl(friendNumber, fno, tox.FILE_CONTROL_CANCEL)
			}
		}
	}, nil)
	t.CallbackFileRecv(func(t *tox.Tox, friendNumber uint32, fileNumber uint32, kind uint32,
		fileSize uint64, fileName string, userData unsafe.Pointer) {
		if debug {
			friendId, err := t.FriendGetPublicKey(friendNumber)
			log.Println("on recv file:", friendNumber, fileNumber, kind, fileSize, fileName, friendId, err)
		}
		if fileSize > 1024*1024*1024 {
			// good guy
		}

		var reFileName = "Re_" + fileName
		reFileNumber, err := t.FileSend(friendNumber, kind, fileSize, reFileName, reFileName)
		if err != nil {
		}
		recvFiles[uint64(uint64(friendNumber)<<32|uint64(fileNumber))] = reFileNumber
		sendFiles[uint64(uint64(friendNumber)<<32|uint64(reFileNumber))] = fileNumber
	}, nil)
	t.CallbackFileRecvChunk(func(t *tox.Tox, friendNumber uint32, fileNumber uint32,
		position uint64, data []byte, userData unsafe.Pointer) {
		friendId, err := t.FriendGetPublicKey(friendNumber)
		if debug {
			// log.Println("on recv chunk:", friendNumber, fileNumber, position, len(data), friendId, err)
		}

		if len(data) == 0 {
			if debug {
				log.Println("recv file finished:", friendNumber, fileNumber, friendId, err)
			}
		} else {
			reFileNumber := recvFiles[uint64(uint64(fileNumber)<<32|uint64(fileNumber))]
			key := makekey(friendNumber, reFileNumber, position)
			sendDatas[key] = data
			trySendChunk(friendNumber, reFileNumber, position)
		}
	}, nil)
	t.CallbackFileChunkRequest(func(t *tox.Tox, friendNumber uint32, fileNumber uint32, position uint64,
		length int, userData unsafe.Pointer) {
		friendId, err := t.FriendGetPublicKey(friendNumber)
		if length == 0 {
			if debug {
				log.Println("send file finished:", friendNumber, fileNumber, friendId, err)
			}
			origFileNumber := sendFiles[uint64(uint64(fileNumber)<<32|uint64(fileNumber))]
			delete(sendFiles, uint64(uint64(fileNumber)<<32|uint64(fileNumber)))
			delete(recvFiles, uint64(uint64(fileNumber)<<32|uint64(origFileNumber)))
		} else {
			key := makekey(friendNumber, fileNumber, position)
			chunkReqs = append(chunkReqs, key)
			trySendChunk(friendNumber, fileNumber, position)
		}
	}, nil)

	// audio/video
	av := tox.NewToxAV(t)
	if av == nil {
	}
	av.CallbackCall(func(av *tox.ToxAV, friendNumber uint32, audioEnabled bool,
		videoEnabled bool, userData unsafe.Pointer) {
		if debug {
			log.Println("oncall:", friendNumber, audioEnabled, videoEnabled)
		}
		var audioBitRate uint32 = 48
		var videoBitRate uint32 = 64
		r, err := av.Answer(friendNumber, audioBitRate, videoBitRate)
		if err != nil {
			log.Println(err, r)
		}
	}, nil)
	av.CallbackCallState(func(av *tox.ToxAV, friendNumber uint32, state uint32, userData unsafe.Pointer) {
		if debug {
			log.Println("on call state:", friendNumber, state)
		}
	}, nil)
	av.CallbackAudioReceiveFrame(func(av *tox.ToxAV, friendNumber uint32, pcm []byte,
		sampleCount int, channels int, samplingRate int, userData unsafe.Pointer) {
		if debug {
			if rand.Int()%23 == 3 {
				log.Println("on recv audio frame:", friendNumber, len(pcm), sampleCount, channels, samplingRate)
			}
		}
		r, err := av.AudioSendFrame(friendNumber, pcm, sampleCount, channels, samplingRate)
		if err != nil {
			log.Println(err, r)
		}
	}, nil)

	// toxav loops
	go func() {
		shutdown := false
		loopc := 0
		itval := uint32(0)
		for !shutdown {
			iv := av.IterationInterval()
			if iv != itval {
				log.Println("av itval changed:", itval, iv)
				itval = iv
			}

			av.Iterate()
			loopc += 1
			time.Sleep(50 * time.Microsecond)
		}

		av.Kill()
	}()

	// toxcore loops
	shutdown := false
	loopc := 0
	itval := uint32(0)
	for !shutdown {
		iv := t.IterationInterval()
		if iv != itval {
			if debug {
				if itval-iv > 10 || iv-itval > 10 {
					log.Println("tox itval changed:", itval, iv)
				}
			}
			itval = iv
		}

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

func makekey(no uint32, a0 interface{}, a1 interface{}) string {
	return fmt.Sprintf("%d_%v_%v", no, a0, a1)
}

func _dirty_init() {
	log.Println("ddddddddd")
	tox.KeepPkg()
}
