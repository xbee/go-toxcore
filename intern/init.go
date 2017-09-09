package toxin

import (
	"log"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}
