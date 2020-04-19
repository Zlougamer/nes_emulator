package olcCpu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	log "github.com/sirupsen/logrus"
)

const BASE_STKP = 0x0100

type instruction struct {
	name string
	addrmodeName string
	operate  func() uint8
	addrmode func() uint8
	cycles   uint8
}

type olc6502 struct {
	mBus *bus
	instrSet *instructionSet
	addrSet *addressingModes
	regSet *registerSet

	fetched uint8

	addrAbs uint16
	addrRel uint16

	opcode uint8
	cycles uint8

	lookup []instruction
}

type Olc6502 interface {
	connectBus(b *bus)
	connectInstructionSet(i *instructionSet)
	connectAdressingModes(a *addressingModes)
	connectRegisterSet(r *registerSet)

	Read(addr uint16) uint8
	Write(addr uint16, data uint8)

	Clock()
	Reset()
	Irq()
	Nmi()

	fetch() uint8
}

func (cpu *olc6502) connectBus(b *bus) {
	cpu.mBus = b
}

func (cpu *olc6502) connectInstructionSet(i *instructionSet) {
	cpu.instrSet = i
}

func (cpu *olc6502) connectAdressingModes(a *addressingModes) {
	cpu.addrSet = a
}

func (cpu *olc6502) connectRegisterSet(r *registerSet) {
	cpu.regSet = r
}

func (cpu *olc6502) fillLookup() {
	file, err := os.Open("./lookup_table")
	if err != nil {
        log.Fatal(err)
    }
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitedLine := strings.Split(scanner.Text(), " ")
		name := splitedLine[0]
		addrModeName := splitedLine[1]
		stringCycles := splitedLine[2]
		cycles, err := strconv.Atoi(stringCycles)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		operate := cpu.instrSet.instrByName[name]
		addrMode := cpu.addrSet.modeByName[addrModeName]
		cpu.lookup = append(cpu.lookup, instruction{
			name, addrModeName, operate, addrMode, uint8(cycles),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func (cpu *olc6502) fetch() uint8 {
	addrmodeName := cpu.lookup[cpu.opcode].addrmodeName
	if addrmodeName != "IMP" {
		log.Debugf(
			"fetch() fetch from addrAbs, addrAbs: %x",
			cpu.addrAbs,
		)
		cpu.fetched = cpu.Read(cpu.addrAbs)
	}
	log.Debugf(
		"fetch() fetched: %x",
		cpu.fetched,
	)
	return cpu.fetched
}

func (cpu *olc6502) Read(addr uint16) uint8 {
	return cpu.mBus.Read(addr, false)
}

func (cpu *olc6502) Write(addr uint16, data uint8) {
	cpu.mBus.Write(addr, data)
}

func (cpu *olc6502) Clock() {
	if cpu.cycles == 0 {
		cpu.opcode = cpu.Read(cpu.regSet.pc)
		cpu.regSet.pc++

		cpu.cycles = cpu.lookup[cpu.opcode].cycles

		additionalCycleAddr := cpu.lookup[cpu.opcode].addrmode()
		additionalCycleOp := cpu.lookup[cpu.opcode].operate()

		cpu.cycles += additionalCycleAddr & additionalCycleOp
		log.Debugf(
			"Clock() opcode: %x, pc: %x, cycles: %d, " +
				"cpu.lookup[cpu.opcode]: %v",
			cpu.opcode,
			cpu.regSet.pc,
			cpu.cycles,
			cpu.lookup[cpu.opcode],
		)
	}

	cpu.cycles--
}

func (cpu *olc6502) Reset() {
	cpu.regSet.a = 0x00
	cpu.regSet.x = 0x00
	cpu.regSet.y = 0x00
	cpu.regSet.stkp = 0xFD
	cpu.regSet.status = 0x00 | uint8(U)

	cpu.addrAbs = 0xFFFC
	hi := uint16(cpu.Read(cpu.addrAbs + 0))
	lo := uint16(cpu.Read(cpu.addrAbs + 1))

	cpu.regSet.pc = (hi << 8) | lo

	cpu.addrAbs = 0x0000
	cpu.addrRel = 0x0000
	cpu.fetched = 0x00

	cpu.cycles = 8
}


func (cpu *olc6502) Irq() {
	regSet := cpu.regSet
	if regSet.getFlag(I) > 0 {
		cpu.Write(BASE_STKP + uint16(regSet.stkp), uint8((regSet.pc >> 8) & 0x00FF))
		regSet.stkp--
		cpu.Write(BASE_STKP + uint16(regSet.stkp), uint8(regSet.pc & 0x00FF))
		regSet.stkp--

		regSet.setFlag(B, false)
		regSet.setFlag(U, true)
		regSet.setFlag(I, true)
		cpu.Write(BASE_STKP + uint16(regSet.stkp), regSet.status)
		regSet.stkp--

		cpu.addrAbs = 0xFFFE
		hi := uint16(cpu.Read(cpu.addrAbs + 0))
		lo := uint16(cpu.Read(cpu.addrAbs + 1))
		regSet.pc = (hi << 8) | lo

		cpu.cycles = 7
	}
}


func (cpu *olc6502) Nmi() {
	regSet := cpu.regSet

	cpu.Write(BASE_STKP + uint16(regSet.stkp), uint8((regSet.pc >> 8) & 0x00FF))
	regSet.stkp--
	cpu.Write(BASE_STKP + uint16(regSet.stkp), uint8(regSet.pc & 0x00FF))
	regSet.stkp--

	regSet.setFlag(B, false)
	regSet.setFlag(U, true)
	regSet.setFlag(I, true)
	cpu.Write(BASE_STKP + uint16(regSet.stkp), regSet.status)
	regSet.stkp--

	cpu.addrAbs = 0xFFFA
	hi := uint16(cpu.Read(cpu.addrAbs + 0))
	lo := uint16(cpu.Read(cpu.addrAbs + 1))
	regSet.pc = (hi << 8) | lo

	cpu.cycles = 8
}


func CreateOlc6502() *olc6502 {
	cpu := &olc6502{}

	cpu.fetched = 0x00

	cpu.addrAbs = 0x0000
	cpu.addrRel = 0x0000

	cpu.opcode = 0x00
	cpu.cycles = 0x00

	cpu.mBus = CreateBus(cpu)
	cpu.instrSet = CreateInstructionSet(cpu)
	cpu.addrSet = CreateAdressingModes(cpu)
	cpu.regSet = CreateRegisterSet(cpu)

	cpu.fillLookup()

	return cpu
}
