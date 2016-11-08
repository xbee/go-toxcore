package tox

/*
#cgo CFLAGS: -g -O2 -Wall
#include <stdlib.h>
#include <stdint.h>

static uint8_t *char2uint8(char *s) { return (uint8_t*)s; }
static unsigned char *char2uchar(char *s) { return (unsigned char*)s; }
static int8_t *short2char(int16_t *s) { return (int8_t*)s; }
*/
import "C"
import "unsafe"
import "errors"
import "fmt"
import "io/ioutil"
import "os"
import "bytes"

// *C.char ==> *C.uint8_t
func char2uint8(s *C.char) *C.uint8_t {
	return C.char2uint8(s)
}

func pointer2uint8(b unsafe.Pointer) *C.uint8_t {
	return C.char2uint8((*C.char)(b))
}

func pointer2uchar(b unsafe.Pointer) *C.uchar {
	return C.char2uchar((*C.char)(b))
}

func bytes2uint8(ba []byte) *C.uint8_t {
	return C.char2uint8((*C.char)((unsafe.Pointer)(&ba[0])))
}

func bytes2uchar(ba []byte) *C.uchar {
	return C.char2uchar((*C.char)((unsafe.Pointer)(&ba[0])))
}

func short2char(b *C.int16_t) *C.int8_t {
	return C.short2char(b)
}

func toxerr(errno interface{}) error {
	return errors.New(fmt.Sprintf("toxcore error: %v", errno))
}

var toxdebug = false

func SetDebug(debug bool) {
	toxdebug = debug
}

var loglevel = 0

func SetLogLevel(level int) {
	loglevel = level
}

func FileExist(fname string) bool {
	_, err := os.Stat(fname)
	if err != nil {
		return false
	}
	return true
}

func (this *Tox) WriteSavedata(fname string) error {
	if !FileExist(fname) {
		err := ioutil.WriteFile(fname, this.GetSavedata(), 0755)
		if err != nil {
			return err
		}
	} else {
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			return err
		}
		liveData := this.GetSavedata()
		if bytes.Compare(data, liveData) != 0 {
			err := ioutil.WriteFile(fname, this.GetSavedata(), 0755)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (this *Tox) LoadSavedata(fname string) ([]byte, error) {
	return ioutil.ReadFile(fname)
}

func LoadSavedata(fname string) ([]byte, error) {
	return ioutil.ReadFile(fname)
}
