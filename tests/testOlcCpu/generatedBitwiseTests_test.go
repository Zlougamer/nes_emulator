// This file is generated! Do not edit it!!!!
package testOlcCpu

import (
	"testing"
	"github.com/Zlougamer/nes_emulator/olcCpu"
)



func TestAndAbsoluteAllZerosSettingZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2D, 0xCF, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndAbsoluteZerosAndOnesSettingNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xAA), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndZpAllZerosSettingZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x25, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndZpZerosAndOnesSettingNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x25, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xAA), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndImmediateAllZerosSettingZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x29, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndImmediateZerosAndOnesSettingNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x29, 0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xAA), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndAbsXAllZerosSettingZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x3d, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndAbsXZerosAndOnesSettingNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x3D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xAA), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndAbsYAllZerosSettingZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x39, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndAbsYZerosAndOnesSettingNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x39, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xAA), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndIndIndexedXAllZerosSettingZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x21, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndIndIndexedXZerosAndOnesSettingNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x21, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xAA), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndIndexedIndYAllZerosSettingZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x31, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndIndexedIndYZerosAndOnesSettingNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x31, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xAA), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndZpXAllZerosSettingZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x35, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestAndZpXAllZerosAndOnesSettingNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x35, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xAA), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestEorAbsoluteFlipsBitsOverSettingZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x4D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
    
}


func TestEorAbsoluteFlipsBitsOverSettingNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x4D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
    
}


func TestEorZpFlipsBitsOverSettingZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x45, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
    
}


func TestEorZpFlipsBitsOverSettingNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x45, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
    
}


func TestEorImmediateFlipsBitsOverSettingZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x49, 0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestEorImmediateFlipsBitsOverSettingNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x49, 0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestEorAbsXIndexedFlipsBitsOverSettingZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x5D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestEorAbsXIndexedFlipsBitsOverSettingNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x5D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestEorAbsYIndexedFlipsBitsOverSettingZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x59, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(regSet.Y)))
    
}


func TestEorAbsYIndexedFlipsBitsOverSettingNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x59, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(regSet.Y)))
    
}


func TestEorIndIndexedXFlipsBitsOverSettingZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x41, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
    
}


func TestEorIndIndexedXFlipsBitsOverSettingNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x41, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
    
}


func TestEorIndexedIndYFlipsBitsOverSettingZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x51, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(regSet.Y)))
    
}


func TestEorIndexedIndYFlipsBitsOverSettingNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x51, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(regSet.Y)))
    
}


func TestEorZpXIndexedFlipsBitsOverSettingZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x55, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestEorZpXIndexedFlipsBitsOverSettingNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x55, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestOraAbsoluteZeroesOrZerosSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x0D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraAbsoluteTurnsBitsOnSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x03
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x0D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x82})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x83), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraZpZeroesOrZerosSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x05, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraZpTurnsBitsOnSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x03
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x05, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x82})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x83), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraImmediateZeroesOrZerosSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x09, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraImmediateTurnsBitsOnSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x03
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x09, 0x82})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x83), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraAbsXIndexedZeroesOrZerosSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x1D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraAbsXIndexedTurnsBitsOnSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x03
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x1D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x82})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x83), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraAbsYIndexedZeroesOrZerosSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x19, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraAbsYIndexedTurnsBitsOnSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x03
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x19, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x83})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x83), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraIndIndexedXZeroesOrZerosSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x01, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraIndIndexedXTurnsBitsOnSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x03
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x01, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x82})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x83), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraIndexedIndYZeroesOrZerosSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x11, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraIndexedIndYTurnsBitsOnSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x03
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x11, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x82})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x83), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraZpXIndexedZeroesOrZerosSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x55, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestOraZpXIndexedTurnsBitsOnSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x03
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x15, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x82})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x83), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBitAbsCopiesBit7OfMemoryToNFlagWhen0(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBitAbsCopiesBit7OfMemoryToNFlagWhen1(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBitAbsCopiesBit_6OfMemoryToVFlagWhen_0(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBitAbsCopiesBit_6OfMemoryToVFlagWhen_1(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBitAbsStoresResultOfAndInZPreservesAWhen_1(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xFEED)))
    
}


func TestBitAbsStoresResultOfAndWhenNonzeroInZPreservesA(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0x01), mpu.Read(uint16(0xFEED)))
    
}


func TestBitAbsStoresResultOfAndWhenZeroInZPreservesA(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xFEED)))
    
}


func TestBitZpCopiesBit_7OfMemoryToNFlagWhen_0(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBitZpCopiesBit_7OfMemoryToNFlagWhen_1(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBitZpCopiesBit_6OfMemoryToVFlagWhen_0(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBitZpCopiesBit_6OfMemoryToVFlagWhen_1(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.V)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestBitZpStoresResultOfAndInZPreservesAWhen_1(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
    
}


func TestBitZpStoresResultOfAndWhenNonzeroInZPreservesA(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x01
    
    
    regSet.Status = regSet.Status | uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x01})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x01), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
	assertEqual(t, uint8(0x01), mpu.Read(uint16(0x0010)))
    
}


func TestBitZpStoresResultOfAndWhenZeroInZPreservesA(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
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



