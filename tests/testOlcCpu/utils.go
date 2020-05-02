package testOlcCpu

import (
	"fmt"
	"github.com/Zlougamer/nes_emulator/olcCpu"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	message := fmt.Sprintf("%x != %x", a, b)
	//t.Fatal(message)
	panic(message)
}

func write(mpu olcCpu.Olc6502, startAddress uint16, bytes []uint8) {
	length := len(bytes)
	for i := 0; i < length; i++ {
		mpu.Write(startAddress + uint16(i), bytes[i])
	}
}

