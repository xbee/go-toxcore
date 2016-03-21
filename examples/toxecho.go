package main

import (
	"fmt"
	"log"
	"tox"
)


func main() {

	t := tox.NewTox()
	r, err := t.BootstrapFromAddress("0.0.0.0", 34567, "jjjjjjjjjjj")
	log.Println(r, err);
	
	s, _ := t.Size()
	fmt.Println("aaaaaaa", t, s)
	t.Do()
	t.Kill()

	tox.TestCCallGo()
}


func _dirty_init() {
	fmt.Println("ddddddddd")
	tox.KeepPkg()
}
