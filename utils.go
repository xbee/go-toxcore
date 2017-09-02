package tox

import "errors"
import "fmt"
import "io/ioutil"
import "os"
import "bytes"

func toxerr(errno interface{}) error {
	return errors.New(fmt.Sprintf("toxcore error: %v", errno))
}

func toxerrf(f string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(f, args...))
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

func ConnStatusString(status int) (s string) {
	switch status {
	case CONNECTION_NONE:
		s = "CONNECTION_NONE"
	case CONNECTION_TCP:
		s = "CONNECTION_TCP"
	case CONNECTION_UDP:
		s = "CONNECTION_UDP"
	}
	return
}
