package pgrpc

import (
	"log"
	"os"
)

const MIN_IDLE = 1
const MAX_IDLE = 5
const MAX_ID_LEN = 64

var Log Logger = log.New(os.Stdout, "pgrpc: ", log.Lshortfile|log.Ltime)

type Logger interface {
	Println(a ...interface{})
}
