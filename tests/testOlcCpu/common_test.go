package testOlcCpu

import (
	"github.com/Zlougamer/nes_emulator/olcCpu"
	"testing"
)

//func TestMain(m *testing.M) {
//	log.SetLevel(log.InfoLevel)
//	code := m.Run()
//	os.Exit(code)
//}

func TestOlcReadWrite(t *testing.T) {
	mpu := olcCpu.CreateOlc6502()
	var addr uint16 = 0xdead
	var data uint8 = 0xbe

	result := mpu.Read(addr)

	assertEqual(t, result, uint8(0x00))

	mpu.Write(addr, data)
	result = mpu.Read(addr)

	assertEqual(t, result, data)
}

func TestOlcReadWriteArray(t *testing.T) {
	mpu := olcCpu.CreateOlc6502()
	var baseAddr uint16 = 0xfff0
	var initData uint8 = 0xf0
	var maxSize uint16 = 0x0100
	var i uint16

	for i = 0x0000; i < maxSize; i++ {
		mpu.Write(baseAddr + i, initData + uint8(i))
	}

	for i = 0x0000; i < maxSize; i++ {
		result := mpu.Read(baseAddr + i)
		expected := initData + uint8(i)
		assertEqual(t, result, expected)
	}
}

//// Reset

func TestResetSetsRegisterToInitialStates(t *testing.T) {
	regSet := olcCpu.CreateRegisterSet()
	mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)
	mpu.Reset()

	assertEqual(t,uint8(0xFD), regSet.Stkp)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, uint8(0x00), regSet.X)
	assertEqual(t, uint8(0x00), regSet.Y)
	assertEqual(t, uint8(0x00 | olcCpu.U), regSet.Status)
}

// NOP

func TestNopDoesNothing(t *testing.T) {
	regSet := olcCpu.CreateRegisterSet()
	mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

	write(mpu, uint16(0x0000), []uint8{0xEA})
	mpu.Clock()

	assertEqual(t,uint16(0x0001), regSet.Pc)
}

// BRK

func TestBrkPushesPcPlus2AndStatusThenSetsPcToIrqVector(t *testing.T) {
	regSet := olcCpu.CreateRegisterSet()
	mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

	regSet.Status = uint8(olcCpu.U)
	write(mpu, uint16(0xFFFE), []uint8{0xCD, 0xAB})

	// $C000 brk
	write(mpu, uint16(0xC000), []uint8{0x00})
	regSet.Pc = 0xC000
	mpu.Clock()

	assertEqual(t, uint16(0xABCD), regSet.Pc)

	assertEqual(t, uint8(0xC0), mpu.Read(uint16(0x01FF)))  // Is it correct?
	// Should be 0x01FD ?
	assertEqual(t, uint8(0x02), mpu.Read(uint16(0x01FE)))
	assertEqual(t, uint8(olcCpu.B | olcCpu.U | olcCpu.I), mpu.Read(uint16(0x01FD))) // strange flags
	// Should be B | U ?
	assertEqual(t, uint8(0xFC), regSet.Stkp)

	assertEqual(t, uint8(olcCpu.U | olcCpu.I), regSet.Status)
	// Should be B | U | I ?
}
