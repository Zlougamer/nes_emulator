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
	//assertEqual(t, uint8(0), mpu.regSet.status & N)
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
