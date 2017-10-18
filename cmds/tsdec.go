package main

//  tox save data decrypt/encrypt

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/kitech/go-toxcore"
)

func init() {
	log.SetFlags(log.Flags() ^ log.Ldate ^ log.Ltime)
}

var tsfile string = "tox_save.tox"
var pass string
var tofile string = "./tsdec.bin"

var crypt_mode string // enc/dec
var decrypt_mode = false
var encrypt_mode = false

func printHelp() {
	log.Println("For help: /path/to/tsdec [options] <tsfile> -h")
}

func main() {
	// flag.StringVar(&tsfile, "tsfile", "", "tox save data file")
	flag.StringVar(&pass, "pass", pass, "tox save data password")
	flag.StringVar(&tofile, "tofile", tofile, "result file")

	flag.Parse()
	// log.Println(flag.Args())
	if len(flag.Args()) < 1 {
		printHelp()
		flag.Usage()
		return
	}
	tsfile = flag.Arg(0)

	data, err := ioutil.ReadFile(tsfile)
	if err != nil {
		log.Println(err)
		return
	}

	isencrypt := tox.IsDataEncrypted(data)
	log.Println("Is encrypt: ", isencrypt)
	if isencrypt && pass == "" {
		log.Println("Need -pass argument.")
		flag.Usage()
		return
	}

	if isencrypt {
		ok, err, salt := tox.GetSalt(data)
		if err != nil {
			log.Println(ok, err, len(salt), salt)
		}
		pkey := tox.NewToxPassKey()
		ok, err = pkey.DeriveWithSalt([]byte(pass), salt)
		if err != nil {
			log.Println(ok, err)
		}
		ok, err, datad := pkey.Decrypt(data)
		pkey.Free()
		if err != nil {
			// log.Println(ok, err, len(datad), datad[0:32])
			log.Println("Decrypt error, check your -pass:", err)
			return
		}

		data = datad
	}

	opts := tox.NewToxOptions()
	opts.Savedata_type = tox.SAVEDATA_TYPE_TOX_SAVE
	opts.Savedata_data = data
	t := tox.NewTox(opts)
	fnums := t.SelfGetFriendList()
	log.Println(fnums, t)
	log.Println("Self Name:", t.SelfGetName())
	log.Println("Self ID:", t.SelfGetAddress())
	mystmsg, err := t.SelfGetStatusMessage()
	log.Println("Status:", mystmsg)
	log.Println("------------------------------------------")
	log.Println("Friend Count:", len(fnums))

	if tox.FileExist(tofile) {
		log.Println(os.ErrExist, tofile)
		// return
	}

	if isencrypt { // do decrypt
		log.Println("Decrypting...")
		err := t.WriteSavedata(tofile)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Save decrypted data OK: ", tofile)
		}
	} else { // do encrypt
		log.Println("Encrypting...")

		pakey := tox.NewToxPassKey()
		// first time, there is no salt
		salt := make([]byte, tox.PASS_KEY_LENGTH)
		_, err, salt := tox.GetSalt(data)
		if err != nil {
			log.Println(err, "GetSalt")
			// return
		}
		_ = salt

		// _, err = pakey.DeriveWithSalt([]byte(pass), salt)
		_, err = pakey.Derive([]byte(pass))
		if err != nil {
			log.Println(err)
			return
		}
		_, err, encdata := pakey.Encrypt(data)
		if err != nil {
			log.Println(err)
			return
		}

		err = ioutil.WriteFile(tofile, encdata, 0755)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Save encrypted data OK: ", tofile)
		}
	}
}

func encrypt() {

}

// @param pt plain tox
func decrypt(pt *tox.Tox) {

}
