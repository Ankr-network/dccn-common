package pgrpc

const MIN_IDLE = 1
const MAX_IDLE = 5
const MAX_ID_LEN = 64

var idPlaceholder []byte

func init() {
	idPlaceholder = make([]byte, MAX_ID_LEN)
	for i := range idPlaceholder {
		idPlaceholder[i] = 0x20 //SPACE
	}
}
