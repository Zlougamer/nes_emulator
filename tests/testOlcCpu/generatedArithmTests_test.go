// This file is generated! Do not edit it!!!!
package testOlcCpu

import (
	"testing"
	"github.com/Zlougamer/nes_emulator/olcCpu"
)



func TestAdcBcdOffAbsoluteCarryClearInAccumulatorZeroes(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6D, 0x00, 0xC0})
    write(mpu, uint16(0xC000), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsoluteCarrySetInAccumulatorZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = olcCpu.C
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6D, 0x00, 0xC0})
    write(mpu, uint16(0xC000), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsoluteCarryClearInNoCarryClearOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6D, 0x00, 0xC0})
    write(mpu, uint16(0xC000), []uint8{0xFE})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsoluteCarryClearInCarrySetOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x02
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6D, 0x00, 0xC0})
    write(mpu, uint16(0xC000), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsoluteOverflowSetNoCarry01Plus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6D, 0x00, 0xC0})
    write(mpu, uint16(0xC000), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x02), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsoluteOverflowSetNoCarry01PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6D, 0x00, 0xC0})
    write(mpu, uint16(0xC000), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsoluteOverflowSetNoCarry7fPlus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x7f
    
    
    regSet.Status = ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6D, 0x00, 0xC0})
    write(mpu, uint16(0xC000), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsoluteOverflowSetNoCarry80PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6D, 0x00, 0xC0})
    write(mpu, uint16(0xC000), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x7F), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsoluteOverflowSetOn40Plus40(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x40
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6D, 0x00, 0xC0})
    write(mpu, uint16(0xC000), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpCarryClearInAccumulatorZeroes(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})
    write(mpu, uint16(0x00B0), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpCarrySetInAccumulatorZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = olcCpu.C
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})
    write(mpu, uint16(0x00B0), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpCarryClearInNoCarryClearOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})
    write(mpu, uint16(0x00B0), []uint8{0xFE})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpCarryClearInCarrySetOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x02
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})
    write(mpu, uint16(0x00B0), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpOverflowSetNoCarry01Plus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})
    write(mpu, uint16(0x00B0), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x02), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpOverflowSetNoCarry01PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})
    write(mpu, uint16(0x00B0), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpOverflowSetNoCarry7fPlus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x7f
    
    
    regSet.Status = ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})
    write(mpu, uint16(0x00B0), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpOverflowSetNoCarry80PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})
    write(mpu, uint16(0x00B0), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x7F), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpOverflowSetOn40Plus40(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x40
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})
    write(mpu, uint16(0x00B0), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffImmediateCarryClearInAccumulatorZeroes(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x69, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffImmediateCarrySetInAccumulatorZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = olcCpu.C
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x69, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffImmediateCarryClearInNoCarryClearOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x69, 0xFE})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffImmediateCarryClearInCarrySetOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x02
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x69, 0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffImmediateOverflowSetNoCarry01Plus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x69, 0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x02), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffImmediateOverflowSetNoCarry01PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x69, 0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffImmediateOverflowSetNoCarry7fPlus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x7f
    
    
    regSet.Status = ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x69, 0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffImmediateOverflowSetNoCarry80PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x69, 0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x7F), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffImmediateOverflowSetOn40Plus40(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x40
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x69, 0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsXCarryClearInAccumulatorZeroes(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsXCarrySetInAccumulatorZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = olcCpu.C
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsXCarryClearInNoCarryClearOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.X), []uint8{0xFE})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsXCarryClearInCarrySetOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x02
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsXOverflowSetNoCarry01Plus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.X), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x02), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsXOverflowSetNoCarry01PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsXOverflowSetNoCarry7fPlus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x7f
    
    
    regSet.Status = ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.X), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsXOverflowSetNoCarry80PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x7F), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsXOverflowSetOn40Plus40(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x40
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.X), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsYCarryClearInAccumulatorZeroes(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsYCarrySetInAccumulatorZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = olcCpu.C
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsYCarryClearInNoCarryClearOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.Y), []uint8{0xFE})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsYCarryClearInCarrySetOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x02
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsYOverflowSetNoCarry01Plus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.Y), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x02), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsYOverflowSetNoCarry01PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsYOverflowSetNoCarry7fPlus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x7f
    
    
    regSet.Status = ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.Y), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsYOverflowSetNoCarry80PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x7F), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffAbsYOverflowSetOn40Plus40(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x40
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})
    write(mpu, uint16(0xC000) + uint16(regSet.Y), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpXCarryClearInAccumulatorZeroes(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x75, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpXCarrySetInAccumulatorZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = olcCpu.C
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x75, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpXCarryClearInNoCarryClearOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x75, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFE})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpXCarryClearInCarrySetOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x02
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x75, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpXOverflowSetNoCarry01Plus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x75, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x02), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpXOverflowSetNoCarry01PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x75, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpXOverflowSetNoCarry7fPlus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x7f
    
    
    regSet.Status = ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x75, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpXOverflowSetNoCarry80PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x75, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x7F), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpXOverflowSetOn40Plus40(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x40
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x75, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpYCarryClearInAccumulatorZeroes(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpYCarrySetInAccumulatorZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = olcCpu.C
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpYCarryClearInNoCarryClearOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xFE})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpYCarryClearInCarrySetOut(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x02
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpYOverflowSetNoCarry01Plus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x02), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpYOverflowSetNoCarry01PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpYOverflowSetNoCarry7fPlus01(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x7f
    
    
    regSet.Status = ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpYOverflowSetNoCarry80PlusFF(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x7F), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAdcBcdOffZpYOverflowSetOn40Plus40(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x40
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x40})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsAllZerosAndNoBorrowIsZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xED, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsDowntoZeroNoBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xED, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsDowntoZeroWithBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xED, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsDowntoFourWithBorrowClearsZN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x07
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xED, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x04), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcZpAllZerosAndNoBorrowIsZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE5, 0x10})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcZpDowntoZeroNoBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE5, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcZpDowntoZeroWithBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE5, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcZpDowntoFourWithBorrowClearsZN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x07
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE5, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x04), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcImmAllZerosAndNoBorrowIsZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE9, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcImmDowntoZeroNoBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE9, 0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcImmDowntoZeroWithBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE9, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcBcdOnImmediate0aMinus00CarrySet(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x0a
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE9, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x0a), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcBcdOnImmediate9aMinus00CarrySet(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x9a
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE9, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x9a), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcBcdOnImmediate00Minus01CarrySet(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | ^uint8(0)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE9, 0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcBcdOnImmediate_20Minus_0aCarryUnset(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x20
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE9, 0x0a})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x15), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsXAllZerosAndNoBorrowIsZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x0D
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xFD, 0xE0, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsXDowntoZeroNoBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x0D
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xFD, 0xE0, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsXDowntoZeroWithBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x0D
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xFD, 0xE0, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsXDowntoFourWithBorrowClearsZN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x07
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x0D
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xFD, 0xE0, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x04), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsYAllZerosAndNoBorrowIsZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x0D
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF9, 0xE0, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsYDowntoZeroNoBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x0D

    write(mpu, uint16(0x0000), []uint8{0xF9, 0xE0, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsYDowntoZeroWithBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x0D

    write(mpu, uint16(0x0000), []uint8{0xF9, 0xE0, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcAbsYDowntoFourWithBorrowClearsZN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x07
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x0D

    write(mpu, uint16(0x0000), []uint8{0xF9, 0xE0, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x04), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcIndXAllZerosAndNoBorrowIsZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE1, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcIndXDowntoZeroNoBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE1, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcIndXDowntoZeroWithBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE1, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcIndXDowntoFourWithBorrowClearsZN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x07
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE1, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x04), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcIndYAllZerosAndNoBorrowIsZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xF1, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcIndYDowntoZeroNoBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xF1, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED) + uint16(regSet.Y), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcIndYDowntoZeroWithBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xF1, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcIndYDowntoFourWithBorrowClearsZN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x07
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xF1, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED) + uint16(regSet.Y), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x04), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcZpXAllZerosAndNoBorrowIsZero(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x0D
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF5, 0x10})
    write(mpu, uint16(0x001D), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcZpXDowntoZeroNoBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x0D
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF5, 0x10})
    write(mpu, uint16(0x001D), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcZpXDowntoZeroWithBorrowSetsZClearsN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x0D
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF5, 0x10})
    write(mpu, uint16(0x001D), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSbcZpXDowntoFourWithBorrowClearsZN(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x07
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x0D
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF5, 0x10})
    write(mpu, uint16(0x001D), []uint8{0x02})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x04), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}



