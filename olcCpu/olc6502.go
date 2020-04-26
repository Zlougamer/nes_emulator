package olcCpu

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type olc6502 struct {
	mBus *bus
	instrSet *instructionSet
	addrSet *addressingModes
	regSet *registerSet
	manEl *managingElement
}

type Olc6502 interface {
	Read(addr uint16) uint8
	Write(addr uint16, data uint8)

	Clock()
	Reset()
	Irq()
	Nmi()
}

func (cpu *olc6502) fillLookup() {
	filepath := getLookupTableFilepath()
	file, err := os.Open(filepath)
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
		cpu.manEl.lookup = append(cpu.manEl.lookup, instruction{
			name, addrModeName, operate, addrMode, uint8(cycles),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getLookupTableFilepath() string {
	_, filename, _, ok := runtime.Caller(1)
	if ok != true {
		panic("runtime.Caller failed")
	}
	return path.Join(path.Dir(filename), "lookup_table")
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}


func (cpu *olc6502) Read(addr uint16) uint8 {
	return cpu.mBus.Read(addr, false)
}

func (cpu *olc6502) Write(addr uint16, data uint8) {
	cpu.mBus.Write(addr, data)
}

func (cpu *olc6502) Clock() {
	regSet := cpu.regSet
	manEl := cpu.manEl

	if manEl.cycles == 0 {
		manEl.opcode = cpu.Read(cpu.regSet.Pc)
		regSet.Pc++

		manEl.cycles = manEl.lookup[manEl.opcode].cycles

		additionalCycleAddr := manEl.lookup[manEl.opcode].addrmode()
		additionalCycleOp := manEl.lookup[manEl.opcode].operate()

		manEl.cycles += additionalCycleAddr & additionalCycleOp
		log.Debugf(
			"Clock() opcode: %x, pc: %x, cycles: %d, " +
				"cpu.lookup[cpu.opcode]: %v",
			manEl.opcode,
			regSet.Pc,
			manEl.cycles,
			manEl.lookup[manEl.opcode],
		)
	}

	manEl.cycles--
}

func (cpu *olc6502) Reset() {
	regSet := cpu.regSet
	manEl := cpu.manEl

	regSet.A = 0x00
	regSet.X = 0x00
	regSet.Y = 0x00
	regSet.Stkp = 0xFD
	regSet.Status = 0x00 | uint8(U)

	manEl.addrAbs = 0xFFFC
	hi := uint16(cpu.Read(manEl.addrAbs + 0))
	lo := uint16(cpu.Read(manEl.addrAbs + 1))

	cpu.regSet.Pc = (hi << 8) | lo

	manEl.addrAbs = 0x0000
	manEl.addrRel = 0x0000
	manEl.fetched = 0x00

	manEl.cycles = 8
}


func (cpu *olc6502) Irq() {
	regSet := cpu.regSet
	manEl := cpu.manEl

	if regSet.getFlag(I) > 0 {
		cpu.Write(BASE_STKP + uint16(regSet.Stkp), uint8((regSet.Pc >> 8) & 0x00FF))
		regSet.Stkp--
		cpu.Write(BASE_STKP + uint16(regSet.Stkp), uint8(regSet.Pc & 0x00FF))
		regSet.Stkp--

		regSet.setFlag(B, false)
		regSet.setFlag(U, true)
		regSet.setFlag(I, true)
		cpu.Write(BASE_STKP + uint16(regSet.Stkp), regSet.Status)
		regSet.Stkp--

		manEl.addrAbs = 0xFFFE
		hi := uint16(cpu.Read(manEl.addrAbs + 0))
		lo := uint16(cpu.Read(manEl.addrAbs + 1))
		regSet.Pc = (hi << 8) | lo

		manEl.cycles = 7
	}
}


func (cpu *olc6502) Nmi() {
	regSet := cpu.regSet
	manEl := cpu.manEl

	cpu.Write(BASE_STKP + uint16(regSet.Stkp), uint8((regSet.Pc >> 8) & 0x00FF))
	regSet.Stkp--
	cpu.Write(BASE_STKP + uint16(regSet.Stkp), uint8(regSet.Pc & 0x00FF))
	regSet.Stkp--

	regSet.setFlag(B, false)
	regSet.setFlag(U, true)
	regSet.setFlag(I, true)
	cpu.Write(BASE_STKP + uint16(regSet.Stkp), regSet.Status)
	regSet.Stkp--

	manEl.addrAbs = 0xFFFA
	hi := uint16(cpu.Read(manEl.addrAbs + 0))
	lo := uint16(cpu.Read(manEl.addrAbs + 1))
	regSet.Pc = (hi << 8) | lo

	manEl.cycles = 8
}


func CreateOlc6502() *olc6502 {
	cpu := &olc6502{}

	cpu.mBus = CreateBus(cpu)

	cpu.regSet = CreateRegisterSet()
	cpu.manEl = CreateManagingElement(cpu.mBus)

	cpu.addrSet = CreateAdressingModes(cpu.mBus, cpu.regSet, cpu.manEl)
	cpu.instrSet = CreateInstructionSet(cpu.mBus, cpu.regSet, cpu.manEl)

	cpu.fillLookup()

	return cpu
}

func CreateOlc6502ByParams(
	regSet *registerSet,
	manEl *managingElement,
	addrSet *addressingModes,
	instrSet *instructionSet,
) *olc6502 {
	cpu := &olc6502{}

	cpu.mBus = CreateBus(cpu)

	if regSet != nil {
		cpu.regSet = regSet
	} else {
		cpu.regSet = CreateRegisterSet()
	}
	if manEl != nil {
		cpu.manEl = manEl
	} else {
		cpu.manEl = CreateManagingElement(cpu.mBus)
	}

	if addrSet != nil {
		cpu.addrSet = addrSet
	} else {
		cpu.addrSet = CreateAdressingModes(cpu.mBus, cpu.regSet, cpu.manEl)
	}
	if instrSet != nil {
		cpu.instrSet = instrSet
	} else {
		cpu.instrSet = CreateInstructionSet(cpu.mBus, cpu.regSet, cpu.manEl)
	}

	cpu.fillLookup()

	return cpu
}
