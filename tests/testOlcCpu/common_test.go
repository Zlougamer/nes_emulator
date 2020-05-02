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

//// eor Absolute
//
//func TestEorAbsoluteFlipsBitsOverSettingZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xFF
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x4D, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//}
//
//func TestEorAbsoluteFlipsBitsOverSettingNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x00
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x4D, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//// eor Zero Page
//
//func TestEorZpFlipsBitsOverSettingZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xFF
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x45, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//}
//
//func TestEorZpFlipsBitsOverSettingNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x00
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x45, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//// eor Immediate
//
//func TestEorImmediateFlipsBitsOverSettingZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xFF
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x49, 0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//}
//
//func TestEorImmediateFlipsBitsOverSettingNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x00
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x49, 0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.A)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//// eor Absolute, X-Indexed
//
//func TestEorAbsXIndexedFlipsBitsOverSettingZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xFF
//	mpu.regSet.X = 0x03
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x5D, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.X), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//}
//
//func TestEorAbsXIndexedFlipsBitsOverSettingNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x00
//	mpu.regSet.X = 0x03
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x5D, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.X), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//// eor Absolute, Y-Indexed
//
//func TestEorAbsYIndexedFlipsBitsOverSettingZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xFF
//	mpu.regSet.Y = 0x03
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x59, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.Y), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.Y)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//}
//
//func TestEorAbsYIndexedFlipsBitsOverSettingNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x00
//	mpu.regSet.Y = 0x03
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x59, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.Y), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.Y)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//// eor Indirect, X-Indexed
//
//func TestEorIndIndexedXFlipsBitsOverSettingZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xFF
//	mpu.regSet.X = 0x03
//
//	write(mpu, uint16(0x0000), []uint8{0x41, 0x10})
//	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//}
//
//func TestEorIndIndexedXFlipsBitsOverSettingNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x00
//	mpu.regSet.X = 0x03
//
//	write(mpu, uint16(0x0000), []uint8{0x41, 0x10})
//	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//// eor Indirect, Y-Indexed
//
//func TestEorIndIndexedYFlipsBitsOverSettingZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xFF
//	mpu.regSet.Y = 0x03
//
//	write(mpu, uint16(0x0000), []uint8{0x51, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.Y), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.Y)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//}
//
//func TestEorIndIndexedYFlipsBitsOverSettingNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x00
//	mpu.regSet.Y = 0x03
//
//	write(mpu, uint16(0x0000), []uint8{0x51, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.Y), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.Y)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//// eor Zero Page, X-Indexed
//
//func TestEorZpXFlipsBitsOverSettingZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xFF
//	mpu.regSet.X = 0x03
//
//	write(mpu, uint16(0x0000), []uint8{0x55, 0x10})
//	write(mpu, uint16(0x0010) + uint16(mpu.regSet.X), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//}
//
//func TestEorZpXFlipsBitsOverSettingNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x00
//	mpu.regSet.X = 0x03
//
//	write(mpu, uint16(0x0000), []uint8{0x55, 0x10})
//	write(mpu, uint16(0x0010) + uint16(mpu.regSet.X), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.A)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//









//
//// asl Accumulator
//
//func TestAslAccumulatorSetsZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x00
//
//	// $0000 asl A
//	mpu.Write(uint16(0x0000), uint8(0x0A))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//func TestAslAccumulatorSetsNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x40
//
//	// $0000 asl A
//	mpu.Write(uint16(0x0000), uint8(0x0A))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x80), mpu.regSet.A)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestAslAccumulatorShiftsOutZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x7F
//
//	// $0000 asl A
//	mpu.Write(uint16(0x0000), uint8(0x0A))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFE), mpu.regSet.A)
//	assertEqual(t, uint8(0), mpu.regSet.Status & C)
//}
//
//func TestAslAccumulatorShiftsOutOne(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xFF
//
//	// $0000 asl A
//	mpu.Write(uint16(0x0000), uint8(0x0A))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFE), mpu.regSet.A)
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//}
//
//func TestAslAccumulator80SetsZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0x80
//	mpu.regSet.Status &= ^uint8(Z)
//
//	// $0000 asl A
//	mpu.Write(uint16(0x0000), uint8(0x0A))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//}
//
//// asl Absolute
//
//func TestAslAbsoluteSetsZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	// $0000 asl A
//	write(mpu, uint16(0x0000), []uint8{0x0E, 0xCD, 0xAB})
//	mpu.Write(uint16(0xABCD), uint8(0x00))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(0x00), mpu.regSet.A)
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//func TestAslAbsoluteSetsNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	// $0000 asl A
//	write(mpu, uint16(0x0000), []uint8{0x0E, 0xCD, 0xAB})
//	mpu.Write(uint16(0xABCD), uint8(0x40))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestAslAbsoluteShiftsOutZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xAA
//
//	// $0000 asl A
//	write(mpu, uint16(0x0000), []uint8{0x0E, 0xCD, 0xAB})
//	mpu.Write(uint16(0xABCD), uint8(0x7F))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xAA), mpu.regSet.A)
//	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & C)
//}
//
//func TestAslAbsoluteShiftsOutOne(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xAA
//
//	// $0000 asl A
//	write(mpu, uint16(0x0000), []uint8{0x0E, 0xCD, 0xAB})
//	mpu.Write(uint16(0xABCD), uint8(0xFF))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xAA), mpu.regSet.A)
//	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//}
//
//// asl Zero Page
//
//func TestAslZpSetsZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	// $0000 asl A
//	write(mpu, uint16(0x0000), []uint8{0x06, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0x00))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//func TestAslZpSetsNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	// $0000 asl A
//	write(mpu, uint16(0x0000), []uint8{0x06, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0x40))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestAslZpShiftsOutZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xAA
//
//	// $0000 asl A
//	write(mpu, uint16(0x0000), []uint8{0x06, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0x7F))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xAA), mpu.regSet.A)
//	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & C)
//}
//
//func TestAslZpShiftsOutOne(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xAA
//
//	// $0000 asl A
//	write(mpu, uint16(0x0000), []uint8{0x06, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0xFF))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xAA), mpu.regSet.A)
//	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//}
//
//// asl Absolute, X-Indexed
//
//func TestAslAbsXIndexedSetsZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.X = 0x03
//
//	// $0000 asl $ABCD,X
//	write(mpu, uint16(0x0000), []uint8{0x1E, 0xCD, 0xAB})
//	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.X), uint8(0x00))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//func TestAslAbsXIndexedSetsNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.X = 0x03
//
//	// $0000 asl $ABCD,X
//	write(mpu, uint16(0x0000), []uint8{0x1E, 0xCD, 0xAB})
//	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.X), uint8(0x40))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestAslAbsXIndexedShiftsOutZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xAA
//	mpu.regSet.X = 0x03
//
//	// $0000 asl $ABCD,X
//	write(mpu, uint16(0x0000), []uint8{0x1E, 0xCD, 0xAB})
//	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.X), uint8(0x7F))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xAA), mpu.regSet.A)
//	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & C)
//}
//
//func TestAslAbsXIndexedShiftsOutOne(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xAA
//	mpu.regSet.X = 0x03
//
//	// $0000 asl $ABCD,X
//	write(mpu, uint16(0x0000), []uint8{0x1E, 0xCD, 0xAB})
//	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.X), uint8(0xFF))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xAA), mpu.regSet.A)
//	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//}
//
//// asl Zero Page, X-Indexed
//
//func TestAslZpXIndexedSetsZFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.X = 0x03
//
//	// $0000 asl $0010,X
//	write(mpu, uint16(0x0000), []uint8{0x16, 0x10})
//	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.X), uint8(0x00))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//func TestAslZpXIndexedSetsNFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.X = 0x03
//
//	// $0000 asl $0010,X
//	write(mpu, uint16(0x0000), []uint8{0x16, 0x10})
//	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.X), uint8(0x40))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestAslZpXIndexedShiftsOutZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xAA
//	mpu.regSet.X = 0x03
//
//	// $0000 asl $0010,X
//	write(mpu, uint16(0x0000), []uint8{0x16, 0x10})
//	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.X), uint8(0x7F))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xAA), mpu.regSet.A)
//	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & C)
//}
//
//func TestAslZpXIndexedShiftsOutOne(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.A = 0xAA
//	mpu.regSet.X = 0x03
//
//	// $0000 asl $0010,X
//	write(mpu, uint16(0x0000), []uint8{0x16, 0x10})
//	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.X), uint8(0xFF))
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xAA), mpu.regSet.A)
//	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//}
//
//// bcc
//
//func TestBccCarryClearBranchesRelativeForward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status = ^uint8(C)
//
//	// $0000 bcc +6
//	write(mpu, uint16(0x0000), []uint8{0x90, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBccCarryClearBranchesRelativeBackward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status = ^uint8(C)
//	mpu.regSet.Pc = uint16(0x0050)
//
//	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
//	// $0000 bcc -6
//	write(mpu, uint16(0x0050), []uint8{0x90, rel})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBccCarrySetDoesNotBranch(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(C)
//
//	// $0000 bcc +6
//	write(mpu, uint16(0x0000), []uint8{0x90, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//}
//
//// bcs
//
//func TestBcsCarrySetBranchesRelativeForward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(C)
//
//	// $0000 bcc +6
//	write(mpu, uint16(0x0000), []uint8{0xB0, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBcsCarrySetBranchesRelativeBackward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(C)
//	mpu.regSet.Pc = uint16(0x0050)
//
//	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
//	// $0000 bcc -6
//	write(mpu, uint16(0x0050), []uint8{0xB0, rel})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBcsCarryClearDoesNotBranch(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(C)
//
//	// $0000 bcc +6
//	write(mpu, uint16(0x0000), []uint8{0xB0, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//}
//
//// beq
//
//func TestBeqZetoSetBranchesRelativeForward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(Z)
//
//	// $0000 beq +6
//	write(mpu, uint16(0x0000), []uint8{0xF0, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBeqZeroSetBranchesRelativeBackward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(Z)
//	mpu.regSet.Pc = uint16(0x0050)
//
//	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
//	// $0000 beq -6
//	write(mpu, uint16(0x0050), []uint8{0xF0, rel})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBeqZeroClearDoesNotBranch(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(Z)
//
//	// $0000 beq +6
//	write(mpu, uint16(0x0000), []uint8{0xF0, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//}
//
//// bit (Absolute)
//
//func TestBitAbsCopiesBit7OfMemoryToNFlagWhen0(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(N)
//
//	// $0000 bit $FEED
//	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
//	mpu.Write(uint16(0xFEED), uint8(0xFF))
//	mpu.regSet.A = 0xFF
//	mpu.Clock()
//
//	assertEqual(t, uint8(N), uint8(mpu.regSet.Status & N))
//}
//
//func TestBitAbsCopiesBit7OfMemoryToNFlagWhen1(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(N)
//
//	// $0000 bit $FEED
//	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
//	mpu.Write(uint16(0xFEED), uint8(0x00))
//	mpu.regSet.A = 0xFF
//	mpu.Clock()
//
//	assertEqual(t, uint8(0), uint8(mpu.regSet.Status & N))
//}
//
//func TestBitAbsCopiesBit6OfMemoryToVFlagWhen0(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(V)
//
//	// $0000 bit $FEED
//	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
//	mpu.Write(uint16(0xFEED), uint8(0xFF))
//	mpu.regSet.A = 0xFF
//	mpu.Clock()
//
//	assertEqual(t, uint8(V), uint8(mpu.regSet.Status & V))
//}
//
//func TestBitAbsCopiesBit6OfMemoryToVFlagWhen1(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(V)
//
//	// $0000 bit $FEED
//	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
//	mpu.Write(uint16(0xFEED), uint8(0x00))
//	mpu.regSet.A = 0xFF
//	mpu.Clock()
//
//	assertEqual(t, uint8(0), uint8(mpu.regSet.Status & V))
//}
//
//func TestBitAbsStoresResultOfAndInZPreservesAWhen1(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(Z)
//
//	// $0000 bit $FEED
//	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
//	mpu.Write(uint16(0xFEED), uint8(0x00))
//	mpu.regSet.A = 0x01
//	mpu.Clock()
//
//	assertEqual(t, uint8(Z), uint8(mpu.regSet.Status & Z))
//	assertEqual(t, uint8(0x01), mpu.regSet.A)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xFEED)))
//}
//
//func TestBitAbsStoresResultOfAndWhenNonzeroInZPreservesA(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(Z)
//
//	// $0000 bit $FEED
//	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
//	mpu.Write(uint16(0xFEED), uint8(0x01))
//	mpu.regSet.A = 0x01
//	mpu.Clock()
//
//	assertEqual(t, uint8(0), uint8(mpu.regSet.Status & Z))
//	assertEqual(t, uint8(0x01), mpu.regSet.A)
//	assertEqual(t, uint8(0x01), mpu.Read(uint16(0xFEED)))
//}
//
//func TestBitAbsStoresResultOfAndWhenZeroInZPreservesA(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(Z)
//
//	// $0000 bit $FEED
//	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
//	mpu.Write(uint16(0xFEED), uint8(0x00))
//	mpu.regSet.A = 0x01
//	mpu.Clock()
//
//	assertEqual(t, uint8(Z), uint8(mpu.regSet.Status & Z))
//	assertEqual(t, uint8(0x01), mpu.regSet.A)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xFEED)))
//}
//
//// bit (Zero Page)
//
//func TestBitZpCopiesBit7OfMemoryToNFlagWhen0(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(N)
//
//	// $0000 bit $0010
//	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0xFF))
//	mpu.regSet.A = 0xFF
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(2), mpu.cycles)
//	assertEqual(t, uint8(N), uint8(mpu.regSet.Status & N))
//}
//
//func TestBitZpCopiesBit7OfMemoryToNFlagWhen1(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(N)
//
//	// $0000 bit $0010
//	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0x00))
//	mpu.regSet.A = 0xFF
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(2), mpu.cycles)
//	assertEqual(t, uint8(0), uint8(mpu.regSet.Status & N))
//}
//
//func TestBitZpCopiesBit6OfMemoryToVFlagWhen0(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(V)
//
//	// $0000 bit $0010
//	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0xFF))
//	mpu.regSet.A = 0xFF
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(2), mpu.cycles)
//	assertEqual(t, uint8(V), uint8(mpu.regSet.Status & V))
//}
//
//func TestBitZpCopiesBit6OfMemoryToVFlagWhen1(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(V)
//
//	// $0000 bit $0010
//	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0x00))
//	mpu.regSet.A = 0xFF
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(2), mpu.cycles)
//	assertEqual(t, uint8(0), uint8(mpu.regSet.Status & V))
//}
//
//func TestBitZpStoresResultOfAndInZPreservesAWhen1(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(Z)
//
//	// $0000 bit $0010
//	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0x00))
//	mpu.regSet.A = 0x01
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(2), mpu.cycles)
//	assertEqual(t, uint8(Z), uint8(mpu.regSet.Status & Z))
//	assertEqual(t, uint8(0x01), mpu.regSet.A)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
//}
//
//func TestBitZpStoresResultOfAndWhenNonzeroInZPreservesA(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(Z)
//
//	// $0000 bit $0010
//	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0x01))
//	mpu.regSet.A = 0x01
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(2), mpu.cycles)
//	assertEqual(t, uint8(0), uint8(mpu.regSet.Status & Z))
//	assertEqual(t, uint8(0x01), mpu.regSet.A)
//	assertEqual(t, uint8(0x01), mpu.Read(uint16(0x0010)))
//}
//
//func TestBitZpStoresResultOfAndWhenZeroInZPreservesA(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(Z)
//
//	// $0000 bit $0010
//	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
//	mpu.Write(uint16(0x0010), uint8(0x00))
//	mpu.regSet.A = 0x01
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(2), mpu.cycles)
//	assertEqual(t, uint8(Z), uint8(mpu.regSet.Status & Z))
//	assertEqual(t, uint8(0x01), mpu.regSet.A)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
//}
//
//// bmi
//
//func TestBmiNegativeSetBranchesRelativeForward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(N)
//
//	// $0000 bmi +6
//	write(mpu, uint16(0x0000), []uint8{0x30, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBmiNegativeSetBranchesRelativeBackward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(N)
//	mpu.regSet.Pc = uint16(0x0050)
//
//	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
//	// $0000 bmi -6
//	write(mpu, uint16(0x0050), []uint8{0x30, rel})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBmiNegativeClearDoesNotBranch(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(N)
//
//	// $0000 bmi +6
//	write(mpu, uint16(0x0000), []uint8{0x30, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//}
//
//// bne
//
//func TestBneZeroSetBranchesRelativeForward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= uint8(Z)
//
//	// $0000 bne +6
//	write(mpu, uint16(0x0000), []uint8{0xD0, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBneZeroSetBranchesRelativeBackward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(Z)
//	mpu.regSet.Pc = uint16(0x0050)
//
//	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
//	// $0000 bne -6
//	write(mpu, uint16(0x0050), []uint8{0xD0, rel})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBneZeroClearDoesNotBranch(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(Z)
//
//	// $0000 bmi +6
//	write(mpu, uint16(0x0000), []uint8{0xD0, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//}
//
//// bpl
//
//func TestBplNegativeSetBranchesRelativeForward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(N)
//
//	// $0000 bpl +6
//	write(mpu, uint16(0x0000), []uint8{0x10, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBplNegativeSetBranchesRelativeBackward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(N)
//	mpu.regSet.Pc = uint16(0x0050)
//
//	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
//	// $0000 bne -6
//	write(mpu, uint16(0x0050), []uint8{0x10, rel})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBplNegativeClearDoesNotBranch(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(N)
//
//	// $0000 bmi +6
//	write(mpu, uint16(0x0000), []uint8{0x10, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//}
//
//// bvc
//
//func TestBvcOverflowClearBranchesRelativeForward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(V)
//
//	// $0000 bvc +6
//	write(mpu, uint16(0x0000), []uint8{0x50, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBvcOverflowSetBranchesRelativeBackward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(V)
//	mpu.regSet.Pc = uint16(0x0050)
//
//	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
//	// $0000 bvc -6
//	write(mpu, uint16(0x0050), []uint8{0x50, rel})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBvcOverflowClearDoesNotBranch(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(V)
//
//	// $0000 bvc +6
//	write(mpu, uint16(0x0000), []uint8{0x50, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//}
//
//
//// bvs
//
//func TestBvsOverflowSetBranchesRelativeForward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(V)
//
//	// $0000 bvc +6
//	write(mpu, uint16(0x0000), []uint8{0x70, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBvsOverflowSetBranchesRelativeBackward(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(V)
//	mpu.regSet.Pc = uint16(0x0050)
//
//	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
//	// $0000 bvs -6
//	write(mpu, uint16(0x0050), []uint8{0x70, rel})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.Pc)
//}
//
//func TestBvsOverflowSetDoesNotBranch(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status &= ^uint8(V)
//
//	// $0000 bvs +6
//	write(mpu, uint16(0x0000), []uint8{0x70, 0x06})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//}
//
//// clc
//
//func TestClcClearsCarryFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(C)
//
//	// $0000 clc
//	write(mpu, uint16(0x0000), []uint8{0x18})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.Status & C)
//}
//
//// cld
//
//// NOWAY
//
//// cli
//
//func TestCliClearsInterruptMaskFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(I)
//
//	// $0000 cli
//	write(mpu, uint16(0x0000), []uint8{0x58})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.Status & I)
//}
//
//// clv
//
//func TestClvClearsOverflowFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Status |= uint8(V)
//
//	// $0000 clv
//	write(mpu, uint16(0x0000), []uint8{0xB8})
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.Status & V)
//}
//
////    Compare instructions
////
////    See http://6502.org/tutorials/compare_instructions.html
////    and http://www.6502.org/tutorials/compare_beyond.html
////    Cheat sheet:
////
////    - Comparison is actually subtraction "register - memory"
////    - Z contains equality result (1 equal, 0 not equal)
////    - C contains result of unsigned comparison (0 if A<m, 1 if A>=m)
////    - N holds MSB of subtraction result (*NOT* of signed subtraction)
////    - V is not affected by comparison
////    - D has no effect on comparison
//
//// cmp Immediate
//
//func TestCmpImmSetsZeroCarryClearsNegFlagsIfEqual(t *testing.T) {
//	// Comparison: A == m
//	mpu := CreateOlc6502()
//	// $0000 cmp #10, A will be 10
//
//	write(mpu, uint16(0x0000), []uint8{0xC9, 10})
//	mpu.regSet.A = 10
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.Status & N)
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//}
//
//func TestCmpImmClearsZeroCarryTakesNegIfLessUnsigned(t *testing.T) {
//	// Comparison: A < m (unsigned)
//	mpu := CreateOlc6502()
//	// $0000 cmp #10, A will be 1
//
//	write(mpu, uint16(0x0000), []uint8{0xC9, 10})
//	mpu.regSet.A = 1
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & C)
//}
//
//func TestCmpImmClearsZeroSetsCarryTakesNegIfLessSigned(t *testing.T) {
//	// Comparison: A < #nn (signed), A negative
//	mpu := CreateOlc6502()
//	// $0000 cmp #10, A will be -1
//
//	write(mpu, uint16(0x0000), []uint8{0xC9, 1})
//	mpu.regSet.A = 0xFF
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)  // 0XFF - 0x01 == 0xFE
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//}
//
//func TestCmpImmClearsZeroCarryTakesNegIfLessSignedNeg(t *testing.T) {
//	// Comparison: A < m (signed), A and m both negative
//	mpu := CreateOlc6502()
//	// $0000 cmp #0xFF (-1), A will be -2 (0xFE)
//
//	write(mpu, uint16(0x0000), []uint8{0xC9, 0xFF})
//	mpu.regSet.A = 0xFE
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)  // 0XFF - 0xFF == 0xFF
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & C)
//}
//
//func TestCmpImmClearsZeroSetsCarryTakesNegIfMoreUnsigned(t *testing.T) {
//	// Comparison: A > m (unsigned)
//	mpu := CreateOlc6502()
//	// $0000 cmp #1, A will be 10
//
//	write(mpu, uint16(0x0000), []uint8{0xC9, 1})
//	mpu.regSet.A = 10
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)  // 0XFF - 0xFF == 0xFF
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//}
//
//func TestCmpImmClearsZeroCarryTakesNegIfMoreSigned(t *testing.T) {
//	// Comparison: A > m (signed), memory negative
//	mpu := CreateOlc6502()
//	// $0000 cmp #$FF (-1), A will be 2
//
//	write(mpu, uint16(0x0000), []uint8{0xC9, 0xFF})
//	mpu.regSet.A = 2
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)  // 0XFF - 0xFF == 0xFF
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & C)
//}
//
//func TestCmpImmClearsZeroCarryTakesNegIfMoreSignedNeg(t *testing.T) {
//	// Comparison: A > m (signed), A and memory both negative
//	mpu := CreateOlc6502()
//	// $0000 cmp #$FF (-2), A will be -1 (0xFF)
//
//	write(mpu, uint16(0x0000), []uint8{0xC9, 0xFE})
//	mpu.regSet.A = 0xFF
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)  // 0XFF - 0xFE == 0x01
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//}
//
//// cpx Immediate
//
//func TestCpxImmSetsZeroCarryClearsNegFlagsIfEqual(t *testing.T) {
//	// Comparison: X == m
//	mpu := CreateOlc6502()
//	// $0000 cpx #$20
//
//	write(mpu, uint16(0x0000), []uint8{0xE0, 0x20})
//	mpu.regSet.X = 0x20
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//// cpy Immediate
//
//func TestCpyImmSetsZeroCarryClearsNegFlagsIfEqual(t *testing.T) {
//	// Comparison: Y == m
//	mpu := CreateOlc6502()
//	// $0000 cpy #$30
//
//	write(mpu, uint16(0x0000), []uint8{0xC0, 0x30})
//	mpu.regSet.Y = 0x30
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(C), mpu.regSet.Status & C)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//// dec Absolute
//
//func TestDecAbsDecrementsMemory(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0xABCD
//
//	write(mpu, uint16(0x0000), []uint8{0xCE, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0x10})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x0F), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestDecAbsBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0xABCD
//
//	write(mpu, uint16(0x0000), []uint8{0xCE, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0x00})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//}
//
//func TestDecAbsSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0xABCD
//
//	write(mpu, uint16(0x0000), []uint8{0xCE, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0x01})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//// dec Zero Page
//
//func TestDecZpDecremensMemory(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010
//
//	write(mpu, uint16(0x0000), []uint8{0xC6, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0x10})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x0F), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestDecZpBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010
//
//	write(mpu, uint16(0x0000), []uint8{0xC6, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0x00})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//}
//
//func TestDecZpSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010
//
//	write(mpu, uint16(0x0000), []uint8{0xC6, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0x01})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//// dec Absolute, X-Indexed
//
//func TestDecAbsXDecrementsMemory(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0xABCD,X
//
//	write(mpu, uint16(0x0000), []uint8{0xDE, 0xCD, 0xAB})
//	mpu.regSet.X = 0x03
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.X), []uint8{0x10})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x0F), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestDecAbsXBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010
//
//	write(mpu, uint16(0x0000), []uint8{0xDE, 0xCD, 0xAB})
//	mpu.regSet.X = 0x03
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.X), []uint8{0x00})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//}
//
//func TestDecAbsXSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010
//
//	write(mpu, uint16(0x0000), []uint8{0xDE, 0xCD, 0xAB})
//	mpu.regSet.X = 0x03
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.X), []uint8{0x01})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//// dec Zero Page, X-Indexed
//
//func TestDecZpXDecrementsMemory(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010,X
//
//	write(mpu, uint16(0x0000), []uint8{0xD6, 0x10})
//	mpu.regSet.X = 0x03
//	write(mpu, uint16(0x0010) + uint16(mpu.regSet.X), []uint8{0x10})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x0F), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestDecZpXBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010,X
//
//	write(mpu, uint16(0x0000), []uint8{0xD6, 0x10})
//	mpu.regSet.X = 0x03
//	write(mpu, uint16(0x0010) + uint16(mpu.regSet.X), []uint8{0x00})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//}
//
//func TestDecZpXSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010,X
//
//	write(mpu, uint16(0x0000), []uint8{0xD6, 0x10})
//	mpu.regSet.X = 0x03
//	write(mpu, uint16(0x0010) + uint16(mpu.regSet.X), []uint8{0x01})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//// dex
//
//func TestDexDecrementsX(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010,X
//	mpu.regSet.X = 0x10
//
//	write(mpu, uint16(0x0000), []uint8{0xCA})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x0F), mpu.regSet.X)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestDexBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.X = 0x00
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0xCA})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.X)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestDexSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.X = 0x01
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0xCA})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.X)
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//// dey
//
//func TestDeyDecrementsY(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010,Y
//	mpu.regSet.Y = 0x10
//
//	write(mpu, uint16(0x0000), []uint8{0x88})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x0F), mpu.regSet.Y)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestDeyBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Y = 0x00
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x88})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0xFF), mpu.regSet.Y)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestDeySetsZeroFlagWhenDecrementingToZero(t *testing.T) {
//	mpu := CreateOlc6502()
//	mpu.regSet.Y = 0x01
//	// $0000 dex
//
//	write(mpu, uint16(0x0000), []uint8{0x88})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0001), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.regSet.Y)
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//// inc Absolute
//
//func TestIncAbsIncrementsMemory(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010,X
//
//	write(mpu, uint16(0x0000), []uint8{0xEE, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0x09})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x0A), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestIncAbsBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//	// $0000 dec 0x0010,X
//
//	write(mpu, uint16(0x0000), []uint8{0xEE, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//func TestIncAbsSetsZeroFlagWhenIncrementingAbove7F(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	write(mpu, uint16(0x0000), []uint8{0xEE, 0xCD, 0xAB})
//	write(mpu, uint16(0xABCD), []uint8{0x7F})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//}
//
//// inc Zero Page
//
//func TestIncZpIncrementsMemory(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	write(mpu, uint16(0x0000), []uint8{0xE6, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0x09})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x0A), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestIncZpBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	write(mpu, uint16(0x0000), []uint8{0xE6, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//func TestIncZpSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	write(mpu, uint16(0x0000), []uint8{0xE6, 0x10})
//	write(mpu, uint16(0x0010), []uint8{0x7F})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0002), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//}
//
//// inc Absolute, X-Indexed
//
//func TestIncAbsXIncrementsMemory(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	write(mpu, uint16(0x0000), []uint8{0xFE, 0xCD, 0xAB})
//	mpu.regSet.X = 0x03
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.X), []uint8{0x09})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x0A), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//}
//
//func TestIncAbsXBelow00RollsOverAndSetsZeroFlag(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	write(mpu, uint16(0x0000), []uint8{0xFE, 0xCD, 0xAB})
//	mpu.regSet.X = 0x03
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.X), []uint8{0xFF})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(Z), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(0), mpu.regSet.Status & N)
//}
//
//func TestIncAbsXSetsNegativeFlagWhenIncrementsAbove7F(t *testing.T) {
//	mpu := CreateOlc6502()
//
//	write(mpu, uint16(0x0000), []uint8{0xFE, 0xCD, 0xAB})
//	mpu.regSet.X = 0x03
//	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.X), []uint8{0x7F})
//
//	mpu.Clock()
//
//	assertEqual(t, uint16(0x0003), mpu.regSet.Pc)
//	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.X)))
//	assertEqual(t, uint8(0), mpu.regSet.Status & Z)
//	assertEqual(t, uint8(N), mpu.regSet.Status & N)
//}
