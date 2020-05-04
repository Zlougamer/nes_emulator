// This file is generated! Do not edit it!!!!
package testOlcCpu

import (
	"testing"
	"github.com/Zlougamer/nes_emulator/olcCpu"
)



func TestStaAbsoluteStoresALeavesAAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x8D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
    
}


func TestStaAbsoluteStoresALeavesAAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x8D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
    
}


func TestStaZpStoresALeavesAAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x85, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
    
}


func TestStaZpStoresALeavesAAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x85, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
    
}


func TestStaAbsXIndexedStoresALeavesAAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x9D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestStaAbsXIndexedStoresALeavesAAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x9D, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(regSet.X)))
    
}


func TestStaAbsYIndexedStoresALeavesAAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x99, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(regSet.Y)))
    
}


func TestStaAbsYIndexedStoresALeavesAAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x99, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(regSet.Y)))
    
}


func TestStaIndIndexedXStoresALeavesAAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x81, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xFEED)))
    
}


func TestStaIndIndexedXStoresALeavesAAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x81, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xFEED)))
    
}


func TestStaIndexedIndYStoresALeavesAAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x91, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xFEED) + uint16(regSet.Y)))
    
}


func TestStaIndexedIndYStoresALeavesAAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x91, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xED, 0xFE})
    write(mpu, uint16(0xFEED) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xFEED) + uint16(regSet.Y)))
    
}


func TestStaZpXIndexedStoresALeavesAAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x95, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0xFF), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestStaZpXIndexedStoresALeavesAAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x95, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestStxAbsoluteStoresXLeavesXAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x0
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0xFF
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x8E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0xFF), regSet.X)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
    
}


func TestStxAbsoluteStoresXLeavesXAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x8E, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
    
}


func TestStxZpStoresXLeavesXAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0xFF
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x86, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0xFF), regSet.X)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
    
}


func TestStxZpStoresXLeavesXAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x86, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
    
}


func TestStxZpYIndexedStoresXLeavesXAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0xFF
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x96, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0xFF), regSet.X)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(regSet.Y)))
    
}


func TestStxZpYIndexedStoresXLeavesXAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0x96, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.Y), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(regSet.Y)))
    
}


func TestStyAbsoluteStoresYLeavesYAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0x8C, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0xFF), regSet.Y)
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
    
}


func TestStyAbsoluteStoresYLeavesYAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x8C, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x00), regSet.Y)
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
    
}


func TestStyZpStoresYLeavesYAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0x84, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0xFF), regSet.Y)
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
    
}


func TestStyZpStoresYLeavesYAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x84, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x00), regSet.Y)
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
    
}


func TestStyZpXIndexedStoresYLeavesYAndNFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.N)
    regSet.X = 0x03
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0x94, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0xFF), regSet.Y)
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestStyZpXIndexedStoresYLeavesYAndZFlagUnchanged(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0xFF & ^uint8(olcCpu.Z)
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x94, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, true, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(regSet.X)))
    
}


func TestLdaAbsoluteLoadsASetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xAD, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaAbsoluteLoadsASetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xAD, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaZpLoadsASetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA5, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaZpLoadsASetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA5, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaImmediateLoadsASetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA9, 0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaImmediateLoadsASetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA9, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaAbsXIndexedLoadsASetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xBD, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaAbsXIndexedLoadsASetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xBD, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaAbsXIndexedDoesNotPageWrap(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0xFF
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xBD, 0x80, 0x00})
    write(mpu, uint16(0x0080) + uint16(regSet.X), []uint8{0x42})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x42), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaAbsYIndexedLoadsASetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xB9, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaAbsYIndexedLoadsASetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xB9, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaAbsYIndexedDoesNotPageWrap(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0xB9, 0x80, 0x00})
    write(mpu, uint16(0x0080) + uint16(regSet.Y), []uint8{0x42})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x42), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaIndIndexedXLoadsASetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA1, 0x10})
    write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaIndIndexedXLoadsASetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA1, 0x10})
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


func TestLdaIndexedIndYLoadsASetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xB1, 0x10})
    write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaIndexedIndYLoadsASetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xB1, 0x10})
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


func TestLdaZpXIndexedLoadsASetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xB5, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdaZpXIndexedLoadsASetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xB5, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxAbsoluteLoadsXSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xAE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxAbsoluteLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0xFF
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xAE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxZpLoadsXSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA6, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxZpLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0xFF
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA6, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxImmediateLoadsXSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA2, 0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxImmediateLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0xFF
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA2, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxAbsYIndexedLoadsXSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xBE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxAbsYIndexedLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0xFF
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xBE, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxZpYIndexedLoadsXSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xB6, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.Y), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdxZpYIndexedLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0xFF
    regSet.Y = 0x03

    write(mpu, uint16(0x0000), []uint8{0xB6, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.Y), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyAbsoluteLoadsYSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xAC, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x80), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyAbsoluteLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0xAC, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyZpLoadsXSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA4, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x80), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyZpLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0xA4, 0x10})
    write(mpu, uint16(0x0010), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x00), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyImmediateLoadsXSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA0, 0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x80), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyImmediateLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0xA0, 0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyAbsXIndexedLoadsXSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xBC, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x80), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyAbsXIndexedLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0xBC, 0xCD, 0xAB})
    write(mpu, uint16(0xABCD) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x00), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyZpXIndexedLoadsXSetsNFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xB4, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x80})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x80), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestLdyZpXIndexedLoadsXSetsZFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0x03
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0xB4, 0x10})
    write(mpu, uint16(0x0010) + uint16(regSet.X), []uint8{0x00})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0002), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x00), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTaxTransfersAccumulatorIntoX(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xAB
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0xAB), regSet.A)
    assertEqual(t, uint8(0xAB), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTaxSetsNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTaxSetsZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0xFF
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xAA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTayTransfersAccumulatorIntoY(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xAB
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA8})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0xAB), regSet.A)
    
    assertEqual(t, uint8(0xAB), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTaySetsNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xA8})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    
    assertEqual(t, uint8(0x80), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTaySetsZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0xA8})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    assertEqual(t, uint8(0x00), regSet.Y)
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTsxTransfersStackPointerIntoX(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xAB
    regSet.Stkp = 0xAB
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xBA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0xAB), regSet.A)
    assertEqual(t, uint8(0xAB), regSet.X)
    
    
    assertEqual(t, uint8(0xAB), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTsxSetsNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x80
    regSet.Stkp = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0xBA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    assertEqual(t, uint8(0x80), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTsxSetsZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    regSet.Stkp = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0xFF

    write(mpu, uint16(0x0000), []uint8{0xBA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    assertEqual(t, uint8(0x00), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTxaTransfersXIntoA(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0xAB
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x8A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0xAB), regSet.A)
    assertEqual(t, uint8(0xAB), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTxaSetsNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    regSet.Stkp = 0x80
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x80
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x8A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x80), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTxaSetsZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xFF
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x8A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTxsTransfersXIntoStackPointer(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = 0x00
    regSet.X = 0xAB
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x9A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0xAB), regSet.X)
    
    
    assertEqual(t, uint8(0xAB), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTxsDoesNotSetNegativeFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.N)
    regSet.X = 0x80
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x9A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x80), regSet.X)
    
    
    assertEqual(t, uint8(0x80), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestTxsDoesNotSetZeroFlag(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x9A})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    assertEqual(t, uint8(0x00), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    
}


func TestPhaPushesAAndUpdatesSp(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0xAB
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x48})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0xAB), regSet.A)
    assertEqual(t, uint8(0x00), regSet.X)
    
    
    assertEqual(t, uint8(0xFE), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    
    
    assertEqual(t, uint8(0xAB), mpu.Read(uint16(0x01FF)))
    
}


func TestPlaPullsTopByteFromStackIntoAAndUpdatesSp(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    regSet.Stkp = 0xFE
    
    
    regSet.Status = regSet.Status & ^uint8(olcCpu.Z)
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x68})
    write(mpu, uint16(0x01FF), []uint8{0xAB})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0xAB), regSet.A)
    
    
    assertEqual(t, uint8(0xFF), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.B != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.U != 0)
    
    
}


func TestPlpPullsTopByteFromStackIntoFlagsAndUpdatesSp(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    regSet.Stkp = 0xFE
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x28})
    write(mpu, uint16(0x01FF), []uint8{0xBA})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0001), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    assertEqual(t, uint8(0xFF), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, true, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    assertEqual(t, false, regSet.Status & olcCpu.I != 0)
    
    assertEqual(t, true, regSet.Status & olcCpu.B != 0)
    
    assertEqual(t, true, regSet.Status & olcCpu.U != 0)
    
    
}


func TestRtiRestoresStatusAndPcAndUpdatesSp(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    regSet.Stkp = 0xFC
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x40})
    write(mpu, uint16(0x01FD), []uint8{0xFC, 0x03, 0xC0})
    
    mpu.Clock()

	assertEqual(t, uint16(0xC003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    assertEqual(t, uint8(0xFF), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, true, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, true, regSet.Status & olcCpu.V != 0)
    assertEqual(t, true, regSet.Status & olcCpu.I != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.B != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.U != 0)
    
    
}


func TestRtiForcesBreakAndUnusedFlagsHigh(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    regSet.Stkp = 0xFC
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x40})
    write(mpu, uint16(0x01FD), []uint8{0x00, 0x03, 0xC0})
    
    mpu.Clock()

	assertEqual(t, uint16(0xC003), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    assertEqual(t, uint8(0xFF), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    assertEqual(t, false, regSet.Status & olcCpu.I != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.B != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.U != 0)
    
    
}


func TestRtsRestoresPcAndIncrementsThenUpdatesSp(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    regSet.Stkp = 0xFD
    
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x0000), []uint8{0x60})
    write(mpu, uint16(0x01FE), []uint8{0x03, 0xC0})
    
    mpu.Clock()

	assertEqual(t, uint16(0xC004), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    assertEqual(t, uint8(0xFF), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    assertEqual(t, false, regSet.Status & olcCpu.I != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.B != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.U != 0)
    
    
}


func TestRtsWrapsAroundTopOfMemory(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = 0x00
    regSet.Stkp = 0xFD
    
    regSet.Pc = 0x1000
    
    regSet.Status = 0x00
    regSet.X = 0x00
    regSet.Y = 0x00

    write(mpu, uint16(0x1000), []uint8{0x60})
    write(mpu, uint16(0x01FE), []uint8{0xFF, 0xFF})
    
    mpu.Clock()

	assertEqual(t, uint16(0x0000), regSet.Pc)
	assertEqual(t, uint8(0x00), regSet.A)
    
    
    assertEqual(t, uint8(0xFF), regSet.Stkp)
    
	assertEqual(t, false, regSet.Status & olcCpu.C != 0)
	assertEqual(t, false, regSet.Status & olcCpu.N != 0)
	assertEqual(t, false, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, false, regSet.Status & olcCpu.V != 0)
    assertEqual(t, false, regSet.Status & olcCpu.I != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.B != 0)
    
    assertEqual(t, false, regSet.Status & olcCpu.U != 0)
    
    
}



