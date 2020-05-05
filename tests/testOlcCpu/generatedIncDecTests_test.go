// This file is generated! Do not edit it!!!!
package testOlcCpu

import (
	"testing"
	"github.com/Zlougamer/nes_emulator/olcCpu"
    "github.com/stretchr/testify/assert"
)



func TestIncAbsIncrementsMemory(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xEE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x09})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x0A), mpu.Read(uint16(0xABCD)))
    
}


func TestIncAbsIncrementsMemoryRollsOverAndSetsZeroFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xEE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xFF})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
    
}


func TestIncAbsSetsNegativeFlagWhenIncrementingAbove7F(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xEE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x7F})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD)))
    
}


func TestIncZpIncrementsMemory(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE6, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x09})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x0A), mpu.Read(uint16(0x0010)))
    
}


func TestIncZpIncrementsMemoryRollsOverAndSetsZeroFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE6, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xFF})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
    
}


func TestIncZpSetsNegativeFlagWhenIncrementingAbove7F(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE6, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x7F})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010)))
    
}


func TestIncAbsXIncrementsMemory(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xFE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x09})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x0A), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestIncAbsXIncrementsMemoryRollsOverAndSetsZeroFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xFE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestIncAbsXSetsNegativeFlagWhenIncrementingAbove7F(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xFE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x7F})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestIncZpXIncrementsMemory(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF6, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x09})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x0A), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestIncZpXIncrementsMemoryRollsOverAndSetsZeroFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF6, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestIncZpXSetsNegativeFlagWhenIncrementingAbove7F(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xF6, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x7F})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestInXIncrementsX(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x09
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE8})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x0A), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestInxAboveFFRollsOverAndSetsZeroFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0xFF
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE8})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestInxSetsNegativeFlagWhenIncrementingAbove7F(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x7F
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xE8})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestInyIncrementsY(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x09

    write(mpu, uint16(0x0000), []uint8{0xC8})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x0A), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestInyAboveFFRollsOverAndSetsZeroFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0xC8})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x00), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestInySetsNegativeFlagWhenIncrementingAbove7F(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x7F

    write(mpu, uint16(0x0000), []uint8{0xC8})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x80), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestDecAbsDecrementsMemory(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xCE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x10})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x0F), mpu.Read(uint16(0xABCD)))
    
}


func TestDecAbsBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xCE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
    
}


func TestDecAbsSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xCE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x01})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
    
}


func TestDecZpDecrementsMemory(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC6, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x10})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x0F), mpu.Read(uint16(0x0010)))
    
}


func TestDecZpBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC6, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
    
}


func TestDecZpSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xC6, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x01})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
    
}


func TestDecAbsXDecrementsMemory(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xDE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x10})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x0F), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestDecAbsXBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xDE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestDecAbsXSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xDE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x01})
    
    mpu.Clock()

	assert.Equal(uint16(0x0003), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestDecZpXDecrementsMemory(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xD6, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x10})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x0F), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestDecZpXBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xD6, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestDecZpXSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xD6, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x01})
    
    mpu.Clock()

	assert.Equal(uint16(0x0002), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestDexDecrementsX(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x10
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xCA})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x0F), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestDexBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xCA})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0xFF), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestDexSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x01
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xCA})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestDeyDecrementsY(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x10

    write(mpu, uint16(0x0000), []uint8{0x88})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x0F), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestDeyBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x88})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0xFF), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestDeySetsZeroFlagWhenDecrementingToZero(t *testing.T) {
    assert := assert.New(t)
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x01

    write(mpu, uint16(0x0000), []uint8{0x88})
    
    mpu.Clock()

	assert.Equal(uint16(0x0001), regSet.Pc, "should be equal")
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x00), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}



