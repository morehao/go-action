package channel

import (
	"testing"
)

func Test_InitChannel(t *testing.T) {
	InitChannel()
}

func Test_NoReceiverChannel(t *testing.T) {
	NoReceiverChannel()
}

func Test_WithReceiverChannel(t *testing.T) {
	WithReceiverChannel()
}

func Test_Close(t *testing.T) {
	Close()
}

func Test_Panic(t *testing.T) {
	Panic()
}

func Test_GetValGoodExample(t *testing.T) {
	GetValGoodExample()
}
