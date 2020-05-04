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
	mpu := olcCpu.CreateOlc6502ByParams(regSet)
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
	mpu := olcCpu.CreateOlc6502ByParams(regSet)

	write(mpu, uint16(0x0000), []uint8{0xEA})
	mpu.Clock()

	assertEqual(t,uint16(0x0001), regSet.Pc)
}

// PHP

func TestPhpPushesProcessorStatusAndUpdatesSp(t *testing.T) {
	for flags := 0; flags < 0x100; flags++ {
		regSet := olcCpu.CreateRegisterSet()
		mpu := olcCpu.CreateOlc6502ByParams(regSet)
		regSet.Status = uint8(flags) | olcCpu.B | olcCpu.U
		write(mpu, uint16(0x0000), []uint8{0x08})

		mpu.Clock()

		assertEqual(t, uint16(0x0001), regSet.Pc)
		expectedFlags := uint8(flags) | olcCpu.B | olcCpu.U
		assertEqual(t, expectedFlags, mpu.Read(uint16(0x01FF)))
		assertEqual(t, uint8(0xFE), regSet.Stkp)
	}
}

// BRK

func TestBrkPushesPcPlus2AndStatusThenSetsPcToIrqVector(t *testing.T) {
	regSet := olcCpu.CreateRegisterSet()
	mpu := olcCpu.CreateOlc6502ByParams(regSet)

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

func TestBrkInterrupt(t *testing.T) {
	regSet := olcCpu.CreateRegisterSet()
	mpu := olcCpu.CreateOlc6502ByParams(regSet)
	regSet.Status = 0x00

	write(mpu, uint16(0xFFFE), []uint8{0x00, 0x04})
	write(mpu, uint16(0x0000), []uint8{
		0xA9, 0x01,  // LDA #$01
		0x00, 0xEA,  // BRK + skepped byte
		0xEA, 0xEA,  // NOP, NOP
		0xA9, 0x03,  // LDA #$03
	})
	write(mpu, uint16(0x0400), []uint8{
		0xA9, 0x02,  // LDA #$02
		0x40,		 // RTI
	})

	mpu.Clock()  // LDA #$01
	assertEqual(t, uint8(0x01), regSet.A)
	assertEqual(t, uint16(0x0002), regSet.Pc)
	mpu.Clock()  // LDA #$01

	mpu.Clock()  // BRK
	assertEqual(t, uint16(0x0400), regSet.Pc)
	mpu.Clock()  // BRK
	mpu.Clock()  // BRK
	mpu.Clock()  // BRK
	mpu.Clock()  // BRK
	mpu.Clock()  // BRK
	mpu.Clock()  // BRK

	mpu.Clock()  // LDA #$02
	assertEqual(t, uint8(0x02), regSet.A)
	assertEqual(t, uint16(0x0402), regSet.Pc)
	mpu.Clock()  // LDA #$02

	mpu.Clock()  // RTI
	assertEqual(t, uint16(0x0004), regSet.Pc)

	mpu.Clock()  // NOP
	mpu.Clock()  // NOP
	mpu.Clock()  // NOP
	mpu.Clock()  // NOP
	mpu.Clock()  // NOP
	mpu.Clock()  // NOP
	mpu.Clock()  // NOP
	mpu.Clock()  // NOP
	mpu.Clock()  // NOP

	mpu.Clock()  // LDA #$03
	assertEqual(t, uint16(0x0008), regSet.Pc)
	assertEqual(t, uint8(0x03), regSet.A)
}


func BrkPreservesDecimalFlagWhenItIsSet(t *testing.T) {
	regSet := olcCpu.CreateRegisterSet()
	mpu := olcCpu.CreateOlc6502ByParams(regSet)

	regSet.Status = uint8(0x00)
	regSet.Pc = uint16(0xC000)

	write(mpu, uint16(0xC000), []uint8{0x00})

	mpu.Clock()

	assertEqual(t, uint8(olcCpu.B), regSet.Status)
}

// LDA ZP, X-Indexed

func LdaZpXIndexedPageWraps(t *testing.T) {
	regSet := olcCpu.CreateRegisterSet()
	mpu := olcCpu.CreateOlc6502ByParams(regSet)

	regSet.Status = uint8(0x00)
	regSet.A = uint8(0x00)
	regSet.X = uint8(0xFF)

	write(mpu, uint16(0x0000), []uint8{0xB5, 0x80})
	write(mpu, uint16(0x007F), []uint8{0x42})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x42), regSet.A)
}

// AND Indexed, Indirect (Y)

func AndIndexedIndYHasPageWrapBug(t *testing.T) {
	regSet := olcCpu.CreateRegisterSet()
	mpu := olcCpu.CreateOlc6502ByParams(regSet)

	regSet.Status = uint8(0x00)
	regSet.Pc = uint16(0x1000)
	regSet.A = uint8(0x42)
	regSet.Y = uint8(0x02)

	write(mpu, uint16(0x1000), []uint8{0x31, 0xFF})
	write(mpu, uint16(0x00FF), []uint8{0x10})
	write(mpu, uint16(0x0100), []uint8{0x20})
	write(mpu, uint16(0x0000), []uint8{0x00})

	write(mpu, uint16(0x2012), []uint8{0x00})
	write(mpu, uint16(0x0012), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint8(0x42), regSet.A)
}

// JMP

func JmpJumpsToAddressWithPageWrapBug(t *testing.T) {
	regSet := olcCpu.CreateRegisterSet()
	mpu := olcCpu.CreateOlc6502ByParams(regSet)

	write(mpu, uint16(0x00FF), []uint8{0x00})
	write(mpu, uint16(0x0000), []uint8{0x6C, 0xFF, 0x00})

	mpu.Clock()

	assertEqual(t, uint16(0x6C00), regSet.Pc)
}
