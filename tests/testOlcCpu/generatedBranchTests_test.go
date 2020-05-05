// This file is generated! Do not edit it!!!!
package testOlcCpu

import (
	"testing"
	"github.com/Zlougamer/nes_emulator/olcCpu"
    "github.com/stretchr/testify/assert"
)



func TestBccCarryClearBranchesRelativeForward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x90, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002 + 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBccCarryClearBranchesRelativeBackward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    regSet.Pc = 0x0050
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0050), []uint8{0x90, 0x06 ^ 0xFF + 1})
    
    mpu.Clock()

	assert.Equal(uint16(0x0052 - 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBccCarrySetDoesNotBranch(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x90, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBcsCarrySetBranchesRelativeForward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xB0, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002 + 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBcsCarrySetBranchesRelativeBackward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    regSet.Pc = 0x0050
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0050), []uint8{0xB0, 0x06 ^ 0xFF + 1})
    
    mpu.Clock()

	assert.Equal(uint16(0x0052 - 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBcsCarryClearDoesNotBranch(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xB0, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBeqZeroSetBranchesRelativeForward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF0, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002 + 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBeqZeroSetBranchesRelativeBackward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    regSet.Pc = 0x0050
    
    regSet.Status = regSet.Status | uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0050), []uint8{0xF0, 0x06 ^ 0xFF + 1})
    
    mpu.Clock()

	assert.Equal(uint16(0x0052 - 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBeqZeroClearDoesNotBranch(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF0, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBmiNegativeSetBranchesRelativeForward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x30, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002 + 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBmiNegativeSetBranchesRelativeBackward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    regSet.Pc = 0x0050
    
    regSet.Status = regSet.Status | uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0050), []uint8{0x30, 0x06 ^ 0xFF + 1})
    
    mpu.Clock()

	assert.Equal(uint16(0x0052 - 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBmiNegativeClearDoesNotBranch(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x30, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBneZeroClearBranchesRelativeForward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xD0, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002 + 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBneZeroClearBranchesRelativeBackward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    regSet.Pc = 0x0050
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0050), []uint8{0xD0, 0x06 ^ 0xFF + 1})
    
    mpu.Clock()

	assert.Equal(uint16(0x0052 - 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBneZeroSetDoesNotBranch(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xD0, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBplNegativeClearBranchesRelativeForward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x10, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002 + 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBplNegativeClearBranchesRelativeBackward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    regSet.Pc = 0x0050
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0050), []uint8{0x10, 0x06 ^ 0xFF + 1})
    
    mpu.Clock()

	assert.Equal(uint16(0x0052 - 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBplNegativeSetDoesNotBranch(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x10, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBvcOverflowClearBranchesRelativeForward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x50, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002 + 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBvcOverflowClearBranchesRelativeBackward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    regSet.Pc = 0x0050
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0050), []uint8{0x50, 0x06 ^ 0xFF + 1})
    
    mpu.Clock()

	assert.Equal(uint16(0x0052 - 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBvcOverflowSetDoesNotBranch(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x50, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBvsOverflowSetBranchesRelativeForward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x70, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002 + 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBvsOverflowSetBranchesRelativeBackward(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    regSet.Pc = 0x0050
    
    regSet.Status = regSet.Status | uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0050), []uint8{0x70, 0x06 ^ 0xFF + 1})
    
    mpu.Clock()

	assert.Equal(uint16(0x0052 - 0x0006), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBvsOverflowClearDoesNotBranch(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x70, 0x06})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestJmpAbsJumpsToAbsoluteAddress(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x4C, 0xCD, 0xAB})
    
    mpu.Clock()

	assert.Equal(uint16(0xABCD), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestJmpIndJumpsToIndirectAddress(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x6C, 0x00, 0x02})
    write(mpu, uint16(0x0200), []uint8{0xCD, 0xAB})
    
    mpu.Clock()

	assert.Equal(uint16(0xABCD), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestJsrPushesPcPlus2AndSetsPc(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    regSet.Pc = 0xC000
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0xC000), []uint8{0x20, 0xD2, 0xFF})
    
    mpu.Clock()

	assert.Equal(uint16(0xFFD2), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    assertEqual(t, uint8(0xFD), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xC0), mpu.Read(uint16(0x01FF)))
    assertEqual(t, uint8(0x02), mpu.Read(uint16(0x01FE)))
    
}


func TestClcClearsCarryFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x18})
    
    mpu.Clock()

	assert.Equal(uint16(0x01), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestCliClearsInterruptMaskFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.I)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x58})
    
    mpu.Clock()

	assert.Equal(uint16(0x01), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    assertEqual(t, false, regSet.Status & olcCpu.I != 0)
    
    
    
    
}


func TestClvClearsOverflowFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xB8})
    
    mpu.Clock()

	assert.Equal(uint16(0x01), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSecSetsCarryFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.C)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x38})
    
    mpu.Clock()

	assert.Equal(uint16(0x01), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestSedSetsDecimalModeFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.I)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x78})
    
    mpu.Clock()

	assert.Equal(uint16(0x01), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    assertEqual(t, true, regSet.Status & olcCpu.I != 0)
    
    
    
    
}


func TestCmpImmSetsZeroCarryClearsNegFlagsIfEqual(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x10
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC9, 0x10})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x10), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestCmpImmClearsZeroCarryTakesNegIfLessUnsigned(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x01
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC9, 0x10})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestCmpImmClearsZeroSetsCarryTakesNegIfLessSigned(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC9, 0x10})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestCmpImmClearsZeroCarryTakesNegIfLessSignedNega(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFE
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC9, 0xFF})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0xFE), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestCmpImmClearsZeroSetsCarryTakesNegIfMoreUnsigned(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x10
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC9, 0x01})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x10), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestCmpImmClearsZeroCarryTakesNegIfMoreSigned(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x02
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC9, 0xFF})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x02), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestCmpImmClearsZeroCarryTakesNegIfMoreSignedNega(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC9, 0xFE})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestCpxImmSetsZeroCarryClearsNegFlagsIfEqual(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x20
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE0, 0x20})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x20), regSet.X)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestCpyImmSetsZeroCarryClearsNegFlagsIfEqual(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x30

    write(mpu, uint16(0x0000), []uint8{0xC0, 0x30})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x30), regSet.Y)
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}



