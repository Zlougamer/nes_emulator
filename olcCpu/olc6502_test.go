package olcCpu

import (
	"fmt"
	"os"
	"testing"
	log "github.com/sirupsen/logrus"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	message := fmt.Sprintf("%x != %x", a, b)
	t.Fatal(message)
}

func write(mpu *olc6502, startAddress uint16, bytes []uint8) {
	length := len(bytes)
	for i := 0; i < length; i++ {
		mpu.Write(startAddress + uint16(i), bytes[i])
	}
}

func TestMain(m *testing.M) {
	log.SetLevel(log.InfoLevel)
	code := m.Run()
	os.Exit(code)
}

func TestOlcReadWrite(t *testing.T) {
	olc := CreateOlc6502()
	var addr uint16 = 0xdead
	var data uint8 = 0xbe

	result := olc.Read(addr)

	assertEqual(t, result, uint8(0x00))

	olc.Write(addr, data)
	result = olc.Read(addr)

	assertEqual(t, result, data)
}

func TestOlcReadWriteArray(t *testing.T) {
	olc := CreateOlc6502()
	var baseAddr uint16 = 0xfff0
	var initData uint8 = 0xf0
	var maxSize uint16 = 0x0100
	var i uint16

	for i = 0x0000; i < maxSize; i++ {
		olc.Write(baseAddr + i, initData + uint8(i))
	}

	for i = 0x0000; i < maxSize; i++ {
		result := olc.Read(baseAddr + i)
		expected := initData + uint8(i)
		assertEqual(t, result, expected)
	}
}

// Reset

func TestResetSetsRegisterToInitialStates(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.Reset()

	assertEqual(t,uint8(0xFD), mpu.regSet.stkp)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0x00), mpu.regSet.x)
	assertEqual(t, uint8(0x00), mpu.regSet.y)
	assertEqual(t, uint8(0x00 | U), mpu.regSet.status)
}

// ADC Absolute

func TestAdcBcdOffAbsoluteCarryClearInAccumulatorZeroes(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00

	// $0000 ADC $C000
	write(mpu, uint16(0x0000), []uint8{0x6d, 0x00, 0xC0})

	mpu.Write(uint16(0xC000), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestAdcBcdOffAbsoluteCarrySetInAccumulatorZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.status |= C

	// $0000 ADC $C000
	write(mpu, uint16(0x0000), []uint8{0x6d, 0x00, 0xC0})

	mpu.Write(uint16(0xC000), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}

func TestAdcBcdOffAbsoluteCarryClearInNoCarryClearOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x01

	// $0000 ADC $C000
	write(mpu, uint16(0x0000), []uint8{0x6d, 0x00, 0xC0})

	mpu.Write(uint16(0xC000), uint8(0xFE))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffAbsoluteCarryClearInCarrySetOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x02

	// $0000 ADC $C000
	write(mpu, uint16(0x0000), []uint8{0x6d, 0x00, 0xC0})

	mpu.Write(uint16(0xC000), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffAbsoluteOverflowClrNoCarry01Plus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01

	// $0000 ADC $C000
	write(mpu, uint16(0x0000), []uint8{0x6d, 0x00, 0xC0})

	mpu.Write(uint16(0xC000), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x02), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffAbsoluteOverflowSetNoCarry7fPlus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x7f

	// $0000 ADC $C000
	write(mpu, uint16(0x0000), []uint8{0x6d, 0x00, 0xC0})

	mpu.Write(uint16(0xC000), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestAdcBcdOffAbsoluteOverflowSetNoCarry80PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x80

	// $0000 ADC $C000
	write(mpu, uint16(0x0000), []uint8{0x6d, 0x00, 0xC0})

	mpu.Write(uint16(0xC000), uint8(0xff))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x7f), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAdcBcdOffAbsoluteOverflowSetOn40Plus40(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)
	mpu.regSet.a = 0x40

	// $0000 ADC $C000
	write(mpu, uint16(0x0000), []uint8{0x6d, 0x00, 0xC0})

	mpu.Write(uint16(0xC000), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// ADC ZeroPage

func TestAdcBcdOffZpCarryClearInAccumulatorZeroes(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00

	// $0000 ADC $00B0
	write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})

	mpu.Write(uint16(0x00B0), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestAdcBcdOffZpCarrySetInAccumulatorZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.status |= C

	// $0000 ADC $00B0
	write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})

	mpu.Write(uint16(0x00B0), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}


func TestAdcBcdOffZpCarryClearInNoCarryClearOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x01

	// $0000 ADC $00B0
	write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})

	mpu.Write(uint16(0x00B0), uint8(0xFE))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffZpCarryClearInCarrySetOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x02

	// $0000 ADC $00B0
	write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})

	mpu.Write(uint16(0x00B0), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffZpOverflowClrNoCarry01Plus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01

	// $0000 ADC $00B0
	write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})

	mpu.Write(uint16(0x00B0), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x02), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffZpOverflowClrNoCarry01PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01

	// $0000 ADC $00B0
	write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})

	mpu.Write(uint16(0x00B0), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffZpOverflowSetNoCarry7fPlus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x7f

	// $0000 ADC $00B0
	write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})

	mpu.Write(uint16(0x00B0), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestAdcBcdOffZpOverflowSetNoCarry80PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x80

	// $0000 ADC $00B0
	write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})

	mpu.Write(uint16(0x00B0), uint8(0xff))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x7f), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAdcBcdOffZpOverflowSetOn40Plus40(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)
	mpu.regSet.a = 0x40

	// $0000 ADC $00B0
	write(mpu, uint16(0x0000), []uint8{0x65, 0xB0})

	mpu.Write(uint16(0x00B0), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// ADC Immediate

func TestAdcBcdOffImmediateCarryClearInAccumulatorZeroes(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00

	// $0000 ADC $00
	write(mpu, uint16(0x0000), []uint8{0x69, 0x00})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestAdcBcdOffImmediateCarrySetInAccumulatorZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.status |= C

	// $0000 ADC $00
	write(mpu, uint16(0x0000), []uint8{0x69, 0x00})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}


func TestAdcBcdOffImmediateCarryClearInNoCarryClearOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x01

	// $0000 ADC $00
	write(mpu, uint16(0x0000), []uint8{0x69, 0xFE})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffImmediateCarryClearInCarrySetOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x02

	// $0000 ADC $FF
	write(mpu, uint16(0x0000), []uint8{0x69, 0xFF})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffImmediateOverflowClrNoCarry01Plus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01

	// $0000 ADC $00
	write(mpu, uint16(0x0000), []uint8{0x69, 0x01})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x02), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffImmediateOverflowClrNoCarry01PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01

	// $0000 ADC $00
	write(mpu, uint16(0x0000), []uint8{0x69, 0xFF})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffImmediateOverflowSetNoCarry7fPlus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x7f

	// $0000 ADC $01
	write(mpu, uint16(0x0000), []uint8{0x69, 0x01})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestAdcBcdOffImmediateOverflowSetNoCarry80PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x80

	// $0000 ADC $FF
	write(mpu, uint16(0x0000), []uint8{0x69, 0xFF})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x7f), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
}

func TestAdcBcdOffImmediateOverflowSetOn40Plus40(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)
	mpu.regSet.a = 0x40

	// $0000 ADC $40
	write(mpu, uint16(0x0000), []uint8{0x69, 0x40})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// ADC Absolute, X-Indexed

func TestAdcBcdOffAbsXCarryClearInAccumulatorZeroes(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.x = 0x03

	// $0000 ADC $C000,X
	write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.x), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestAdcBcdOffAbsXCarrySetInAccumulatorZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.status |= C
	mpu.regSet.x = 0x03

	// $0000 ADC $C000,X
	write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.x), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}


func TestAdcBcdOffAbsXCarryClearInNoCarryClearOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x01
	mpu.regSet.x = 0x03

	// $0000 ADC $C000,X
	write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.x), uint8(0xFE))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffAbsXCarryClearInCarrySetOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x02
	mpu.regSet.x = 0x03

	// $0000 ADC $C000,X
	write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.x), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffAbsXOverflowClrNoCarry01Plus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.x = 0x00

	// $0000 ADC $C000,X
	write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.x), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x02), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffAbsXOverflowClrNoCarry01PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.x = 0x00

	// $0000 ADC $C000,X
	write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.x), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffAbsXOverflowSetNoCarry7fPlus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x7f
	mpu.regSet.x = 0x00

	// $0000 ADC $C000,X
	write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.x), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestAdcBcdOffAbsXOverflowSetNoCarry80PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x80
	mpu.regSet.x = 0x00

	// $0000 ADC $C000,X
	write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.x), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x7f), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
}

func TestAdcBcdOffAbsXOverflowSetOn40Plus40(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)
	mpu.regSet.a = 0x40
	mpu.regSet.x = 0x03

	// $0000 ADC $C000,X
	write(mpu, uint16(0x0000), []uint8{0x7D, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.x), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// ADC Absolute, Y-Indexed

func TestAdcBcdOffAbsYCarryClearInAccumulatorZeroes(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.y = 0x03

	// $0000 ADC $C000,Y
	write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.y), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestAdcBcdOffAbsYCarrySetInAccumulatorZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.status |= C
	mpu.regSet.y = 0x03

	// $0000 ADC $C000,Y
	write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.y), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}


func TestAdcBcdOffAbsYCarryClearInNoCarryClearOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x01
	mpu.regSet.y = 0x03

	// $0000 ADC $C000,Y
	write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.y), uint8(0xFE))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffAbsYCarryClearInCarrySetOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x02
	mpu.regSet.y = 0x03

	// $0000 ADC $C000,Y
	write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.y), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffAbsYOverflowClrNoCarry01Plus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.y = 0x00

	// $0000 ADC $C000,Y
	write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.y), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x02), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffAbsYOverflowClrNoCarry01PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.y = 0x00

	// $0000 ADC $C000,Y
	write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.y), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffAbsYOverflowSetNoCarry7fPlus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x7f
	mpu.regSet.y = 0x00

	// $0000 ADC $C000,Y
	write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.y), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestAdcBcdOffAbsYOverflowSetNoCarry80PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x80
	mpu.regSet.y = 0x00

	// $0000 ADC $C000,Y
	write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.y), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x7f), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
}

func TestAdcBcdOffAbsYOverflowSetOn40Plus40(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)
	mpu.regSet.a = 0x40
	mpu.regSet.y = 0x03

	// $0000 ADC $C000,Y
	write(mpu, uint16(0x0000), []uint8{0x79, 0x00, 0xC0})

	mpu.Write(uint16(0xC000) + uint16(mpu.regSet.y), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// ADC Absolute, Y-Indexed

func TestAdcBcdOffZpXCarryClearInAccumulatorZeroes(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.x = 0x03

	// $0000 ADC $0010,X
	write(mpu, uint16(0x0000), []uint8{0x75, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestAdcBcdOffZpXCarrySetInAccumulatorZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.status |= C
	mpu.regSet.x = 0x03

	// $0000 ADC $0010,X
	write(mpu, uint16(0x0000), []uint8{0x75, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}


func TestAdcBcdOffZpXCarryClearInNoCarryClearOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x01
	mpu.regSet.x = 0x03

	// $0000 ADC $0010,X
	write(mpu, uint16(0x0000), []uint8{0x75, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0xFE))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffZpXCarryClearInCarrySetOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x02
	mpu.regSet.x = 0x03

	// $0000 ADC $0010,X
	write(mpu, uint16(0x0000), []uint8{0x75, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffZpXOverflowClrNoCarry01Plus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.x = 0x03

	// $0000 ADC $0010,X
	write(mpu, uint16(0x0000), []uint8{0x75, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x02), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffZpXOverflowClrNoCarry01PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.x = 0x03

	// $0000 ADC $0010,X
	write(mpu, uint16(0x0000), []uint8{0x75, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffZpXOverflowSetNoCarry7fPlus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x7f
	mpu.regSet.x = 0x03

	// $0000 ADC $0010,X
	write(mpu, uint16(0x0000), []uint8{0x75, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestAdcBcdOffZpXOverflowSetNoCarry80PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x80
	mpu.regSet.x = 0x01

	// $0000 ADC $0010,X
	write(mpu, uint16(0x0000), []uint8{0x75, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x7f), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
}

func TestAdcBcdOffZpXOverflowSetOn40Plus40(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)
	mpu.regSet.a = 0x40
	mpu.regSet.x = 0x03

	// $0000 ADC $0010,X
	write(mpu, uint16(0x0000), []uint8{0x75, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// ADC Indirect, Indexed (X)

func TestAdcBcdOffIndIndexedCarryClearInAccumulatorZeroes(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.x = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x61, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestAdcBcdOffIndIndexedCarrySetInAccumulatorZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.status |= C
	mpu.regSet.x = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x61, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}


func TestAdcBcdOffIndIndexedCarryClearInNoCarryClearOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x01
	mpu.regSet.x = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x61, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0xFE))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffIndIndexedCarryClearInCarrySetOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x02
	mpu.regSet.x = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x61, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffIndIndexedOverflowClrNoCarry01Plus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.x = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x61, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x02), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffIndIndexedOverflowClrNoCarry01PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.x = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x61, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffIndIndexedOverflowSetNoCarry7fPlus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x7f
	mpu.regSet.x = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x61, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestAdcBcdOffIndIndexedOverflowSetNoCarry80PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x80
	mpu.regSet.x = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x61, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x7f), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
}

func TestAdcBcdOffIndIndexedXOverflowSetOn40Plus40(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)
	mpu.regSet.a = 0x40
	mpu.regSet.x = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x61, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// ADC Indexed, Indirect (Y)

func TestAdcBcdOffIndIndexedYCarryClearInAccumulatorZeroes(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.y = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestAdcBcdOffIndIndexedYCarrySetInAccumulatorZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.status |= C
	mpu.regSet.y = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}


func TestAdcBcdOffIndIndexedYCarryClearInNoCarryClearOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x01
	mpu.regSet.y = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0xFE))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffIndIndexedYCarryClearInCarrySetOut(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x02
	mpu.regSet.y = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & uint8(N))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAdcBcdOffIndIndexedYOverflowClrNoCarry01Plus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.y = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x02), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffIndIndexedYOverflowClrNoCarry01PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x01
	mpu.regSet.y = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & V)
}

func TestAdcBcdOffIndIndexedYOverflowSetNoCarry7fPlus01(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x7f
	mpu.regSet.y = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0x01))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestAdcBcdOffIndIndexedYOverflowSetNoCarry80PlusFF(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)
	mpu.regSet.a = 0x80
	mpu.regSet.y = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x7f), mpu.regSet.a)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
}

func TestAdcBcdOffIndIndexedYOverflowSetOn40Plus40(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)
	mpu.regSet.a = 0x40
	mpu.regSet.y = 0x03

	// $0000 ADC ($0010,X)
	// $0013 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x71, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(V), mpu.regSet.status & V)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// AND (Absolute)

func TestAndAbsoluteAllZerosSettingZeroFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF

	// $0000 AND $ABCD
	write(mpu, uint16(0x0000), []uint8{0x2D, 0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAndAbsoluteZerosAndOnesSettingNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF

	// $0000 AND $ABCD
	write(mpu, uint16(0x0000), []uint8{0x2D, 0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0xAA))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// AND (Zero Page)

func TestAndZpAllZerosSettingZeroFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF

	// $0000 AND $0010
	write(mpu, uint16(0x0000), []uint8{0x25, 0x10})

	mpu.Write(uint16(0x0010), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAndZpZerosAndOnesSettingNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF

	// $0000 AND $0010
	write(mpu, uint16(0x0000), []uint8{0x25, 0x10})

	mpu.Write(uint16(0x0010), uint8(0xAA))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// AND (Immediate)

func TestAndImmediateAllZerosSettingZeroFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF

	// $0000 AND $0010
	write(mpu, uint16(0x0000), []uint8{0x29, 0x00})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAndImmediateZerosAndOnesSettingNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF

	// $0000 AND #$AA
	write(mpu, uint16(0x0000), []uint8{0x29, 0xAA})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// AND (Absolute, X-Indexed)

func TestAndAbsXAllZerosSettingZeroFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.x = 0x03

	// $0000 AND $ABCD,X
	write(mpu, uint16(0x0000), []uint8{0x3D, 0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.x), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAndAbsXZerosAndOnesSettingNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.x = 0x03

	// $0000 AND $ABCD,X
	write(mpu, uint16(0x0000), []uint8{0x3D, 0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.x), uint8(0xAA))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// AND (Absolute, Y-Indexed)

func TestAndAbsYAllZerosSettingZeroFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.y = 0x03

	// $0000 AND $ABCD,Y
	write(mpu, uint16(0x0000), []uint8{0x39, 0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAndAbsYZerosAndOnesSettingNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.x = 0x03

	// $0000 AND $ABCD,Y
	write(mpu, uint16(0x0000), []uint8{0x39, 0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0xAA))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// AND (Indirect, X-Indexed)

func TestAndIndIndexedXAllZerosSettingZeroFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.x = 0x03

	// $0000 AND $ABCD,Y
	write(mpu, uint16(0x0000), []uint8{0x21, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAndIndIndexedXZerosAndOnesSettingNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.x = 0x03

	// $0000 AND $ABCD,Y
	write(mpu, uint16(0x0000), []uint8{0x21, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD), uint8(0xAA))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// AND (Indirect, Y-Indexed)

func TestAndIndIndexedYAllZerosSettingZeroFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.y = 0x03

	// $0000 AND ($0010),Y
	// $0010 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x31, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAndIndIndexedYZerosAndOnesSettingNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.y = 0x03

	// $0000 AND ($0010),Y
	// $0010 Vector to $ABCD
	write(mpu, uint16(0x0000), []uint8{0x31, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})

	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.y), uint8(0xAA))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// AND (Zero Page, X-Indexed)

func TestAndZpXAllZerosSettingZeroFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.y = 0x03

	// $0000 AND $0010,X
	write(mpu, uint16(0x0000), []uint8{0x35, 0x10})
	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.x), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAndZpXZerosAndOnesSettingNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.x = 0x03

	// $0000 AND $0010,X
	write(mpu, uint16(0x0000), []uint8{0x35, 0x10})

	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0xAA))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// ASL Accumulator

func TestAslAccumulatorSetsZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00

	// $0000 ASL A
	mpu.Write(uint16(0x0000), uint8(0x0A))
	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAslAccumulatorSetsNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x40

	// $0000 ASL A
	mpu.Write(uint16(0x0000), uint8(0x0A))
	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAslAccumulatorShiftsOutZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x7F

	// $0000 ASL A
	mpu.Write(uint16(0x0000), uint8(0x0A))
	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0xFE), mpu.regSet.a)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}

func TestAslAccumulatorShiftsOutOne(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF

	// $0000 ASL A
	mpu.Write(uint16(0x0000), uint8(0x0A))
	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0xFE), mpu.regSet.a)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
}

func TestAslAccumulator80SetsZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x80
	mpu.regSet.status &= ^uint8(Z)

	// $0000 ASL A
	mpu.Write(uint16(0x0000), uint8(0x0A))
	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

// ASL Absolute

func TestAslAbsoluteSetsZFlag(t *testing.T) {
	mpu := CreateOlc6502()

	// $0000 ASL A
	write(mpu, uint16(0x0000), []uint8{0x0E, 0xCD, 0xAB})
	mpu.Write(uint16(0xABCD), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAslAbsoluteSetsNFlag(t *testing.T) {
	mpu := CreateOlc6502()

	// $0000 ASL A
	write(mpu, uint16(0x0000), []uint8{0x0E, 0xCD, 0xAB})
	mpu.Write(uint16(0xABCD), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAslAbsoluteShiftsOutZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xAA

	// $0000 ASL A
	write(mpu, uint16(0x0000), []uint8{0x0E, 0xCD, 0xAB})
	mpu.Write(uint16(0xABCD), uint8(0x7F))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}

func TestAslAbsoluteShiftsOutOne(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xAA

	// $0000 ASL A
	write(mpu, uint16(0x0000), []uint8{0x0E, 0xCD, 0xAB})
	mpu.Write(uint16(0xABCD), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(C), mpu.regSet.status & C)
}

// ASL Zero Page

func TestAslZpSetsZFlag(t *testing.T) {
	mpu := CreateOlc6502()

	// $0000 ASL A
	write(mpu, uint16(0x0000), []uint8{0x06, 0x10})
	mpu.Write(uint16(0x0010), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAslZpSetsNFlag(t *testing.T) {
	mpu := CreateOlc6502()

	// $0000 ASL A
	write(mpu, uint16(0x0000), []uint8{0x06, 0x10})
	mpu.Write(uint16(0x0010), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAslZpShiftsOutZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xAA

	// $0000 ASL A
	write(mpu, uint16(0x0000), []uint8{0x06, 0x10})
	mpu.Write(uint16(0x0010), uint8(0x7F))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}

func TestAslZpShiftsOutOne(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xAA

	// $0000 ASL A
	write(mpu, uint16(0x0000), []uint8{0x06, 0x10})
	mpu.Write(uint16(0x0010), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(C), mpu.regSet.status & C)
}

// ASL Absolute, X-Indexed

func TestAslAbsXIndexedSetsZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.x = 0x03

	// $0000 ASL $ABCD,X
	write(mpu, uint16(0x0000), []uint8{0x1E, 0xCD, 0xAB})
	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.x), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAslAbsXIndexedSetsNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.x = 0x03

	// $0000 ASL $ABCD,X
	write(mpu, uint16(0x0000), []uint8{0x1E, 0xCD, 0xAB})
	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.x), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAslAbsXIndexedShiftsOutZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xAA
	mpu.regSet.x = 0x03

	// $0000 ASL $ABCD,X
	write(mpu, uint16(0x0000), []uint8{0x1E, 0xCD, 0xAB})
	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.x), uint8(0x7F))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}

func TestAslAbsXIndexedShiftsOutOne(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xAA
	mpu.regSet.x = 0x03

	// $0000 ASL $ABCD,X
	write(mpu, uint16(0x0000), []uint8{0x1E, 0xCD, 0xAB})
	mpu.Write(uint16(0xABCD) + uint16(mpu.regSet.x), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(C), mpu.regSet.status & C)
}

// ASL Zero Page, X-Indexed

func TestAslZpXIndexedSetsZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.x = 0x03

	// $0000 ASL $0010,X
	write(mpu, uint16(0x0000), []uint8{0x16, 0x10})
	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0x00))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestAslZpXIndexedSetsNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.x = 0x03

	// $0000 ASL $0010,X
	write(mpu, uint16(0x0000), []uint8{0x16, 0x10})
	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0x40))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestAslZpXIndexedShiftsOutZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xAA
	mpu.regSet.x = 0x03

	// $0000 ASL $0010,X
	write(mpu, uint16(0x0000), []uint8{0x16, 0x10})
	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0x7F))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}

func TestAslZpXIndexedShiftsOutOne(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xAA
	mpu.regSet.x = 0x03

	// $0000 ASL $0010,X
	write(mpu, uint16(0x0000), []uint8{0x16, 0x10})
	mpu.Write(uint16(0x0010) + uint16(mpu.regSet.x), uint8(0xFF))
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xAA), mpu.regSet.a)
	assertEqual(t, uint8(0xFE), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(C), mpu.regSet.status & C)
}

// BCC

func TestBccCarryClearBranchesRelativeForward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status = ^uint8(C)

	// $0000 BCC +6
	write(mpu, uint16(0x0000), []uint8{0x90, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.pc)
}

func TestBccCarryClearBranchesRelativeBackward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status = ^uint8(C)
	mpu.regSet.pc = uint16(0x0050)

	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
	// $0000 BCC -6
	write(mpu, uint16(0x0050), []uint8{0x90, rel})
	mpu.Clock()

	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.pc)
}

func TestBccCarrySetDoesNotBranch(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(C)

	// $0000 BCC +6
	write(mpu, uint16(0x0000), []uint8{0x90, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
}

// BCS

func TestBcsCarrySetBranchesRelativeForward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(C)

	// $0000 BCC +6
	write(mpu, uint16(0x0000), []uint8{0xB0, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.pc)
}

func TestBcsCarrySetBranchesRelativeBackward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(C)
	mpu.regSet.pc = uint16(0x0050)

	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
	// $0000 BCC -6
	write(mpu, uint16(0x0050), []uint8{0xB0, rel})
	mpu.Clock()

	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.pc)
}

func TestBcsCarryClearDoesNotBranch(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(C)

	// $0000 BCC +6
	write(mpu, uint16(0x0000), []uint8{0xB0, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
}

// BEQ

func TestBeqZetoSetBranchesRelativeForward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(Z)

	// $0000 BEQ +6
	write(mpu, uint16(0x0000), []uint8{0xF0, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.pc)
}

func TestBeqZeroSetBranchesRelativeBackward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(Z)
	mpu.regSet.pc = uint16(0x0050)

	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
	// $0000 BEQ -6
	write(mpu, uint16(0x0050), []uint8{0xF0, rel})
	mpu.Clock()

	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.pc)
}

func TestBeqZeroClearDoesNotBranch(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(Z)

	// $0000 BEQ +6
	write(mpu, uint16(0x0000), []uint8{0xF0, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
}

// BIT (Absolute)

func TestBitAbsCopiesBit7OfMemoryToNFlagWhen0(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(N)

	// $0000 BIT $FEED
	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
	mpu.Write(uint16(0xFEED), uint8(0xFF))
	mpu.regSet.a = 0xFF
	mpu.Clock()

	assertEqual(t, uint8(N), uint8(mpu.regSet.status & N))
}

func TestBitAbsCopiesBit7OfMemoryToNFlagWhen1(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(N)

	// $0000 BIT $FEED
	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
	mpu.Write(uint16(0xFEED), uint8(0x00))
	mpu.regSet.a = 0xFF
	mpu.Clock()

	assertEqual(t, uint8(0), uint8(mpu.regSet.status & N))
}

func TestBitAbsCopiesBit6OfMemoryToVFlagWhen0(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)

	// $0000 BIT $FEED
	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
	mpu.Write(uint16(0xFEED), uint8(0xFF))
	mpu.regSet.a = 0xFF
	mpu.Clock()

	assertEqual(t, uint8(V), uint8(mpu.regSet.status & V))
}

func TestBitAbsCopiesBit6OfMemoryToVFlagWhen1(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(V)

	// $0000 BIT $FEED
	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
	mpu.Write(uint16(0xFEED), uint8(0x00))
	mpu.regSet.a = 0xFF
	mpu.Clock()

	assertEqual(t, uint8(0), uint8(mpu.regSet.status & V))
}

func TestBitAbsStoresResultOfAndInZPreservesAWhen1(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(Z)

	// $0000 BIT $FEED
	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
	mpu.Write(uint16(0xFEED), uint8(0x00))
	mpu.regSet.a = 0x01
	mpu.Clock()

	assertEqual(t, uint8(Z), uint8(mpu.regSet.status & Z))
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xFEED)))
}

func TestBitAbsStoresResultOfAndWhenNonzeroInZPreservesA(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(Z)

	// $0000 BIT $FEED
	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
	mpu.Write(uint16(0xFEED), uint8(0x01))
	mpu.regSet.a = 0x01
	mpu.Clock()

	assertEqual(t, uint8(0), uint8(mpu.regSet.status & Z))
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0x01), mpu.Read(uint16(0xFEED)))
}

func TestBitAbsStoresResultOfAndWhenZeroInZPreservesA(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(Z)

	// $0000 BIT $FEED
	write(mpu, uint16(0x0000), []uint8{0x2C, 0xED, 0xFE})
	mpu.Write(uint16(0xFEED), uint8(0x00))
	mpu.regSet.a = 0x01
	mpu.Clock()

	assertEqual(t, uint8(Z), uint8(mpu.regSet.status & Z))
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xFEED)))
}

// BIT (Zero Page)

func TestBitZpCopiesBit7OfMemoryToNFlagWhen0(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(N)

	// $0000 BIT $0010
	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
	mpu.Write(uint16(0x0010), uint8(0xFF))
	mpu.regSet.a = 0xFF
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(2), mpu.cycles)
	assertEqual(t, uint8(N), uint8(mpu.regSet.status & N))
}

func TestBitZpCopiesBit7OfMemoryToNFlagWhen1(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(N)

	// $0000 BIT $0010
	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
	mpu.Write(uint16(0x0010), uint8(0x00))
	mpu.regSet.a = 0xFF
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(2), mpu.cycles)
	assertEqual(t, uint8(0), uint8(mpu.regSet.status & N))
}

func TestBitZpCopiesBit6OfMemoryToVFlagWhen0(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)

	// $0000 BIT $0010
	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
	mpu.Write(uint16(0x0010), uint8(0xFF))
	mpu.regSet.a = 0xFF
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(2), mpu.cycles)
	assertEqual(t, uint8(V), uint8(mpu.regSet.status & V))
}

func TestBitZpCopiesBit6OfMemoryToVFlagWhen1(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(V)

	// $0000 BIT $0010
	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
	mpu.Write(uint16(0x0010), uint8(0x00))
	mpu.regSet.a = 0xFF
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(2), mpu.cycles)
	assertEqual(t, uint8(0), uint8(mpu.regSet.status & V))
}

func TestBitZpStoresResultOfAndInZPreservesAWhen1(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(Z)

	// $0000 BIT $0010
	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
	mpu.Write(uint16(0x0010), uint8(0x00))
	mpu.regSet.a = 0x01
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(2), mpu.cycles)
	assertEqual(t, uint8(Z), uint8(mpu.regSet.status & Z))
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
}

func TestBitZpStoresResultOfAndWhenNonzeroInZPreservesA(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(Z)

	// $0000 BIT $0010
	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
	mpu.Write(uint16(0x0010), uint8(0x01))
	mpu.regSet.a = 0x01
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(2), mpu.cycles)
	assertEqual(t, uint8(0), uint8(mpu.regSet.status & Z))
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0x01), mpu.Read(uint16(0x0010)))
}

func TestBitZpStoresResultOfAndWhenZeroInZPreservesA(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(Z)

	// $0000 BIT $0010
	write(mpu, uint16(0x0000), []uint8{0x24, 0x10})
	mpu.Write(uint16(0x0010), uint8(0x00))
	mpu.regSet.a = 0x01
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(2), mpu.cycles)
	assertEqual(t, uint8(Z), uint8(mpu.regSet.status & Z))
	assertEqual(t, uint8(0x01), mpu.regSet.a)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
}

// BMI

func TestBmiNegativeSetBranchesRelativeForward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(N)

	// $0000 BMI +6
	write(mpu, uint16(0x0000), []uint8{0x30, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.pc)
}

func TestBmiNegativeSetBranchesRelativeBackward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(N)
	mpu.regSet.pc = uint16(0x0050)

	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
	// $0000 BMI -6
	write(mpu, uint16(0x0050), []uint8{0x30, rel})
	mpu.Clock()

	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.pc)
}

func TestBmiNegativeClearDoesNotBranch(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(N)

	// $0000 BMI +6
	write(mpu, uint16(0x0000), []uint8{0x30, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
}

// BNE

func TestBneZeroSetBranchesRelativeForward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= uint8(Z)

	// $0000 BNE +6
	write(mpu, uint16(0x0000), []uint8{0xD0, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.pc)
}

func TestBneZeroSetBranchesRelativeBackward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(Z)
	mpu.regSet.pc = uint16(0x0050)

	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
	// $0000 BNE -6
	write(mpu, uint16(0x0050), []uint8{0xD0, rel})
	mpu.Clock()

	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.pc)
}

func TestBneZeroClearDoesNotBranch(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(Z)

	// $0000 BMI +6
	write(mpu, uint16(0x0000), []uint8{0xD0, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
}

// BPL

func TestBplNegativeSetBranchesRelativeForward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(N)

	// $0000 BPL +6
	write(mpu, uint16(0x0000), []uint8{0x10, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.pc)
}

func TestBplNegativeSetBranchesRelativeBackward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(N)
	mpu.regSet.pc = uint16(0x0050)

	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
	// $0000 BNE -6
	write(mpu, uint16(0x0050), []uint8{0x10, rel})
	mpu.Clock()

	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.pc)
}

func TestBplNegativeClearDoesNotBranch(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(N)

	// $0000 BMI +6
	write(mpu, uint16(0x0000), []uint8{0x10, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
}

// BRK

func TestBrkPushesPcPlus2AndStatusThenSetsPcToIrqVector(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status = uint8(U)
	write(mpu, uint16(0xFFFE), []uint8{0xCD, 0xAB})

	// $C000 BRK
	mpu.Write(uint16(0xC000), uint8(0x00))
	mpu.regSet.pc = 0xC000
	mpu.Clock()

	assertEqual(t, uint16(0xABCD), mpu.regSet.pc)

	assertEqual(t, uint8(0xC0), mpu.Read(uint16(0x01FF)))  // Is it correct?
	// Should be 0x01FD ?
	assertEqual(t, uint8(0x02), mpu.Read(uint16(0x01FE)))
	assertEqual(t, uint8(B | U | I), mpu.Read(uint16(0x01FD))) // strange flags
	// Should be B | U ?
	assertEqual(t, uint8(0xFC), mpu.regSet.stkp)

	assertEqual(t, uint8(U | I), mpu.regSet.status)
	// Should be B | U | I ?
}

// BVC

func TestBvcOverflowClearBranchesRelativeForward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)

	// $0000 BVC +6
	write(mpu, uint16(0x0000), []uint8{0x50, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.pc)
}

func TestBvcOverflowSetBranchesRelativeBackward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)
	mpu.regSet.pc = uint16(0x0050)

	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
	// $0000 BVC -6
	write(mpu, uint16(0x0050), []uint8{0x50, rel})
	mpu.Clock()

	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.pc)
}

func TestBvcOverflowClearDoesNotBranch(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(V)

	// $0000 BVC +6
	write(mpu, uint16(0x0000), []uint8{0x50, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
}


// BVS

func TestBvsOverflowSetBranchesRelativeForward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(V)

	// $0000 BVC +6
	write(mpu, uint16(0x0000), []uint8{0x70, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002) + uint16(0x06), mpu.regSet.pc)
}

func TestBvsOverflowSetBranchesRelativeBackward(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(V)
	mpu.regSet.pc = uint16(0x0050)

	rel := uint8((0x06 ^ 0xFF) + 1) // two's complement of 6
	// $0000 BVS -6
	write(mpu, uint16(0x0050), []uint8{0x70, rel})
	mpu.Clock()

	assertEqual(t, uint16(0x0052) - uint16(0x06), mpu.regSet.pc)
}

func TestBvsOverflowSetDoesNotBranch(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status &= ^uint8(V)

	// $0000 BVS +6
	write(mpu, uint16(0x0000), []uint8{0x70, 0x06})
	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
}

// CLC

func TestClcClearsCarryFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(C)

	// $0000 CLC
	write(mpu, uint16(0x0000), []uint8{0x18})
	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.status & C)
}

// CLD

// NOWAY

// CLI

func TestCliClearsInterruptMaskFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(I)

	// $0000 CLI
	write(mpu, uint16(0x0000), []uint8{0x58})
	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.status & I)
}

// CLV

func TestClvClearsOverflowFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.status |= uint8(V)

	// $0000 CLV
	write(mpu, uint16(0x0000), []uint8{0xB8})
	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.status & V)
}

//    Compare instructions
//
//    See http://6502.org/tutorials/compare_instructions.html
//    and http://www.6502.org/tutorials/compare_beyond.html
//    Cheat sheet:
//
//    - Comparison is actually subtraction "register - memory"
//    - Z contains equality result (1 equal, 0 not equal)
//    - C contains result of unsigned comparison (0 if A<m, 1 if A>=m)
//    - N holds MSB of subtraction result (*NOT* of signed subtraction)
//    - V is not affected by comparison
//    - D has no effect on comparison

// CMP Immediate

func TestCmpImmSetsZeroCarryClearsNegFlagsIfEqual(t *testing.T) {
	// Comparison: A == m
	mpu := CreateOlc6502()
	// $0000 CMP #10, A will be 10

	write(mpu, uint16(0x0000), []uint8{0xC9, 10})
	mpu.regSet.a = 10

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.status & N)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
}

func TestCmpImmClearsZeroCarryTakesNegIfLessUnsigned(t *testing.T) {
	// Comparison: A < m (unsigned)
	mpu := CreateOlc6502()
	// $0000 CMP #10, A will be 1

	write(mpu, uint16(0x0000), []uint8{0xC9, 10})
	mpu.regSet.a = 1

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}

func TestCmpImmClearsZeroSetsCarryTakesNegIfLessSigned(t *testing.T) {
	// Comparison: A < #nn (signed), A negative
	mpu := CreateOlc6502()
	// $0000 CMP #10, A will be -1

	write(mpu, uint16(0x0000), []uint8{0xC9, 1})
	mpu.regSet.a = 0xFF

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(N), mpu.regSet.status & N)  // 0XFF - 0x01 == 0xFE
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
}

func TestCmpImmClearsZeroCarryTakesNegIfLessSignedNeg(t *testing.T) {
	// Comparison: A < m (signed), A and m both negative
	mpu := CreateOlc6502()
	// $0000 CMP #0xFF (-1), A will be -2 (0xFE)

	write(mpu, uint16(0x0000), []uint8{0xC9, 0xFF})
	mpu.regSet.a = 0xFE

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(N), mpu.regSet.status & N)  // 0XFF - 0xFF == 0xFF
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}

func TestCmpImmClearsZeroSetsCarryTakesNegIfMoreUnsigned(t *testing.T) {
	// Comparison: A > m (unsigned)
	mpu := CreateOlc6502()
	// $0000 CMP #1, A will be 10

	write(mpu, uint16(0x0000), []uint8{0xC9, 1})
	mpu.regSet.a = 10

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0), mpu.regSet.status & N)  // 0XFF - 0xFF == 0xFF
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
}

func TestCmpImmClearsZeroCarryTakesNegIfMoreSigned(t *testing.T) {
	// Comparison: A > m (signed), memory negative
	mpu := CreateOlc6502()
	// $0000 CMP #$FF (-1), A will be 2

	write(mpu, uint16(0x0000), []uint8{0xC9, 0xFF})
	mpu.regSet.a = 2

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0), mpu.regSet.status & N)  // 0XFF - 0xFF == 0xFF
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & C)
}

func TestCmpImmClearsZeroCarryTakesNegIfMoreSignedNeg(t *testing.T) {
	// Comparison: A > m (signed), A and memory both negative
	mpu := CreateOlc6502()
	// $0000 CMP #$FF (-2), A will be -1 (0xFF)

	write(mpu, uint16(0x0000), []uint8{0xC9, 0xFE})
	mpu.regSet.a = 0xFF

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0), mpu.regSet.status & N)  // 0XFF - 0xFE == 0x01
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
}

// CPX Immediate

func TestCpxImmSetsZeroCarryClearsNegFlagsIfEqual(t *testing.T) {
	// Comparison: X == m
	mpu := CreateOlc6502()
	// $0000 CPX #$20

	write(mpu, uint16(0x0000), []uint8{0xE0, 0x20})
	mpu.regSet.x = 0x20

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

// CPY Immediate

func TestCpyImmSetsZeroCarryClearsNegFlagsIfEqual(t *testing.T) {
	// Comparison: Y == m
	mpu := CreateOlc6502()
	// $0000 CPY #$30

	write(mpu, uint16(0x0000), []uint8{0xC0, 0x30})
	mpu.regSet.y = 0x30

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(C), mpu.regSet.status & C)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

// Dec Absolute

func TestDecAbsDecrementsMemory(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0xABCD

	write(mpu, uint16(0x0000), []uint8{0xCE, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0x10})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x0F), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestDecAbsBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0xABCD

	write(mpu, uint16(0x0000), []uint8{0xCE, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0x00})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestDecAbsSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0xABCD

	write(mpu, uint16(0x0000), []uint8{0xCE, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0x01})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

// DEC Zero Page

func TestDecZpDecremensMemory(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010

	write(mpu, uint16(0x0000), []uint8{0xC6, 0x10})
	write(mpu, uint16(0x0010), []uint8{0x10})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x0F), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestDecZpBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010

	write(mpu, uint16(0x0000), []uint8{0xC6, 0x10})
	write(mpu, uint16(0x0010), []uint8{0x00})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestDecZpSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010

	write(mpu, uint16(0x0000), []uint8{0xC6, 0x10})
	write(mpu, uint16(0x0010), []uint8{0x01})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

// DEC Absolute, X-Indexed

func TestDecAbsXDecrementsMemory(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0xABCD,X

	write(mpu, uint16(0x0000), []uint8{0xDE, 0xCD, 0xAB})
	mpu.regSet.x = 0x03
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.x), []uint8{0x10})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x0F), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestDecAbsXBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010

	write(mpu, uint16(0x0000), []uint8{0xDE, 0xCD, 0xAB})
	mpu.regSet.x = 0x03
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.x), []uint8{0x00})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestDecAbsXSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010

	write(mpu, uint16(0x0000), []uint8{0xDE, 0xCD, 0xAB})
	mpu.regSet.x = 0x03
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.x), []uint8{0x01})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

// Dec Zero Page, X-Indexed

func TestDecZpXDecrementsMemory(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010,X

	write(mpu, uint16(0x0000), []uint8{0xD6, 0x10})
	mpu.regSet.x = 0x03
	write(mpu, uint16(0x0010) + uint16(mpu.regSet.x), []uint8{0x10})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x0F), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestDecZpXBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010,X

	write(mpu, uint16(0x0000), []uint8{0xD6, 0x10})
	mpu.regSet.x = 0x03
	write(mpu, uint16(0x0010) + uint16(mpu.regSet.x), []uint8{0x00})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

func TestDecZpXSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010,X

	write(mpu, uint16(0x0000), []uint8{0xD6, 0x10})
	mpu.regSet.x = 0x03
	write(mpu, uint16(0x0010) + uint16(mpu.regSet.x), []uint8{0x01})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

// DEX

func TestDexDecrementsX(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010,X
	mpu.regSet.x = 0x10

	write(mpu, uint16(0x0000), []uint8{0xCA})

	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x0F), mpu.regSet.x)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestDexBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.x = 0x00
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0xCA})

	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.x)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestDexSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.x = 0x01
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0xCA})

	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.x)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

// DEY

func TestDeyDecrementsY(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010,Y
	mpu.regSet.y = 0x10

	write(mpu, uint16(0x0000), []uint8{0x88})

	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x0F), mpu.regSet.y)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestDeyBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.y = 0x00
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x88})

	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.y)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestDeySetsZeroFlagWhenDecrementingToZero(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.y = 0x01
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x88})

	mpu.Clock()

	assertEqual(t, uint16(0x0001), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.y)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

// EOR Absolute

func TestEorAbsoluteFlipsBitsOverSettingZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x4D, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestEorAbsoluteFlipsBitsOverSettingNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x4D, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// EOR Zero Page

func TestEorZpFlipsBitsOverSettingZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x45, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestEorZpFlipsBitsOverSettingNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x45, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// EOR Immediate

func TestEorImmediateFlipsBitsOverSettingZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x49, 0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestEorImmediateFlipsBitsOverSettingNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x49, 0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// EOR Absolute, X-Indexed

func TestEorAbsXIndexedFlipsBitsOverSettingZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.x = 0x03
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x5D, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.x), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestEorAbsXIndexedFlipsBitsOverSettingNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.x = 0x03
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x5D, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.x), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// EOR Absolute, Y-Indexed

func TestEorAbsYIndexedFlipsBitsOverSettingZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.y = 0x03
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x59, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.y), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.y)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestEorAbsYIndexedFlipsBitsOverSettingNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.y = 0x03
	// $0000 DEX

	write(mpu, uint16(0x0000), []uint8{0x59, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.y), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.y)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// EOR Indirect, X-Indexed

func TestEorIndIndexedXFlipsBitsOverSettingZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.x = 0x03

	write(mpu, uint16(0x0000), []uint8{0x41, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestEorIndIndexedXFlipsBitsOverSettingNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.x = 0x03

	write(mpu, uint16(0x0000), []uint8{0x41, 0x10})
	write(mpu, uint16(0x0013), []uint8{0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// EOR Indirect, Y-Indexed

func TestEorIndIndexedYFlipsBitsOverSettingZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.y = 0x03

	write(mpu, uint16(0x0000), []uint8{0x51, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.y), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.y)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestEorIndIndexedYFlipsBitsOverSettingNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.y = 0x03

	write(mpu, uint16(0x0000), []uint8{0x51, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xCD, 0xAB})
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.y), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.y)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// EOR Zero Page, X-Indexed

func TestEorZpXFlipsBitsOverSettingZFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0xFF
	mpu.regSet.x = 0x03

	write(mpu, uint16(0x0000), []uint8{0x55, 0x10})
	write(mpu, uint16(0x0010) + uint16(mpu.regSet.x), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
}

func TestEorZpXFlipsBitsOverSettingNFlag(t *testing.T) {
	mpu := CreateOlc6502()
	mpu.regSet.a = 0x00
	mpu.regSet.x = 0x03

	write(mpu, uint16(0x0000), []uint8{0x55, 0x10})
	write(mpu, uint16(0x0010) + uint16(mpu.regSet.x), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0xFF), mpu.regSet.a)
	assertEqual(t, uint8(0xFF), mpu.Read(uint16(0x0010) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(N), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

// INC Absolute

func TestIncAbsIncrementsMemory(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010,X

	write(mpu, uint16(0x0000), []uint8{0xEE, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0x09})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x0A), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestIncAbsBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()
	// $0000 DEC 0x0010,X

	write(mpu, uint16(0x0000), []uint8{0xEE, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestIncAbsSetsZeroFlagWhenIncrementingAbove7F(t *testing.T) {
	mpu := CreateOlc6502()

	write(mpu, uint16(0x0000), []uint8{0xEE, 0xCD, 0xAB})
	write(mpu, uint16(0xABCD), []uint8{0x7F})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD)))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

// INC Zero Page

func TestIncZpIncrementsMemory(t *testing.T) {
	mpu := CreateOlc6502()

	write(mpu, uint16(0x0000), []uint8{0xE6, 0x10})
	write(mpu, uint16(0x0010), []uint8{0x09})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x0A), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestIncZpBelow00RollsOverAndSetsNegativeFlag(t *testing.T) {
	mpu := CreateOlc6502()

	write(mpu, uint16(0x0000), []uint8{0xE6, 0x10})
	write(mpu, uint16(0x0010), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestIncZpSetsZeroFlagWhenDecrementingToZero(t *testing.T) {
	mpu := CreateOlc6502()

	write(mpu, uint16(0x0000), []uint8{0xE6, 0x10})
	write(mpu, uint16(0x0010), []uint8{0x7F})

	mpu.Clock()

	assertEqual(t, uint16(0x0002), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0x0010)))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}

// INC Absolute, X-Indexed

func TestIncAbsXIncrementsMemory(t *testing.T) {
	mpu := CreateOlc6502()

	write(mpu, uint16(0x0000), []uint8{0xFE, 0xCD, 0xAB})
	mpu.regSet.x = 0x03
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.x), []uint8{0x09})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x0A), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(0), mpu.regSet.status & N)
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
}

func TestIncAbsXBelow00RollsOverAndSetsZeroFlag(t *testing.T) {
	mpu := CreateOlc6502()

	write(mpu, uint16(0x0000), []uint8{0xFE, 0xCD, 0xAB})
	mpu.regSet.x = 0x03
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.x), []uint8{0xFF})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x00), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(Z), mpu.regSet.status & Z)
	assertEqual(t, uint8(0), mpu.regSet.status & N)
}

func TestIncAbsXSetsNegativeFlagWhenIncrementsAbove7F(t *testing.T) {
	mpu := CreateOlc6502()

	write(mpu, uint16(0x0000), []uint8{0xFE, 0xCD, 0xAB})
	mpu.regSet.x = 0x03
	write(mpu, uint16(0xABCD) + uint16(mpu.regSet.x), []uint8{0x7F})

	mpu.Clock()

	assertEqual(t, uint16(0x0003), mpu.regSet.pc)
	assertEqual(t, uint8(0x80), mpu.Read(uint16(0xABCD) + uint16(mpu.regSet.x)))
	assertEqual(t, uint8(0), mpu.regSet.status & Z)
	assertEqual(t, uint8(N), mpu.regSet.status & N)
}
