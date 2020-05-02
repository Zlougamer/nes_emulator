// This file is generated! Do not edit it!!!!
package testOlcCpu

import (
	"testing"
	"github.com/Zlougamer/nes_emulator/olcCpu"
)



func TestRolAccumulatorZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
}


func TestRolAccumulator80AndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x80
    regSet.Status = regSet.Status & ^uint8(olcCpu.C | olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
}


func TestRolAccumulatorZeroAndCarryOneClearsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
}


func TestRolAccumulatorSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x40
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x81), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
}


func TestRolAccumulatorShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x7F
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0xFE), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
}


func TestRolAccumulatorShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    regSet.Status = regSet.Status | ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0xFE), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
}


func TestRolAbsoluteZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
    
}


func TestRolAbsolute80AndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C | olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
    
}


func TestRolAbsoluteZeroAndCarryOneClearsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x01), mpu.Read(uint16(0xABCD)))
    
}


func TestRolAbsoluteSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0xABCD)))
    
}


func TestRolAbsoluteShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x7F})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD)))
    
}


func TestRolAbsoluteShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD)))
    
}


func TestRolZpZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x26, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
    
}


func TestRolZp80AndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C | olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x26, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
    
}


func TestRolZpZeroAndCarryOneClearsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x26, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x01), mpu.Read(uint16(0x0010)))
    
}


func TestRolZpSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x26, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0x0010)))
    
}


func TestRolZpShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x26, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x7F})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010)))
    
}


func TestRolZpShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x26, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010)))
    
}


func TestRolAbsXIndexedZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x3E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRolAbsXIndexed80AndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C | olcCpu.Z)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x3E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRolAbsXIndexedZeroAndCarryOneClearsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x3E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x01), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRolAbsXIndexedSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x3E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRolAbsXIndexedShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x3E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x7F})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRolAbsXIndexedShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x3E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRolZpXIndexedZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x36, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestRolZpXIndexed80AndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C | olcCpu.Z)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x36, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestRolZpXIndexedZeroAndCarryOneClearsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x36, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x01), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestRolZpXIndexedSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x36, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestRolZpXIndexedShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x36, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x7F})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestRolZpXIndexedShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x36, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestRorAccumulatorZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
}


func TestRorAccumulatorZeroAndCarryOneRotatesInSetsNFlags(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
}


func TestRorAccumulatorShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x02
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x81), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
}


func TestRorAccumulatorShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x03
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x81), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
}


func TestRorAbsoluteZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
    
}


func TestRorAbsoluteZeroAndCarryOneRotatesInSetsNFlags(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD)))
    
}


func TestRorAbsoluteShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0xABCD)))
    
}


func TestRorAbsoluteShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x03})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0xABCD)))
    
}


func TestRorZpZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x66, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
    
}


func TestRorZpZeroAndCarryOneRotatesInSetsNFlags(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x66, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010)))
    
}


func TestRorZpZeroAbsoluteShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x66, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0x0010)))
    
}


func TestRorZpShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x66, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x03})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0x0010)))
    
}


func TestRorAbsXIndexedZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRorAbsXIndexedZAndC1RotatesInSetsNFlags(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRorAbsXIndexedShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRorAbsXIndexedShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x03})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestRorZpXIndexedZeroAndCarryZeroSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x76, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestRorZpXIndexedZeroAndCarryOneRotatesInSetsNFlags(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x76, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestRorZpXIndexedZeroAbsoluteShiftsOutZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x76, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestRorZpXIndexedShiftsOutOne(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x76, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x03})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
	assertEqual(t, uint8(0x81), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}




