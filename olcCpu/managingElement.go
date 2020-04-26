package olcCpu

import (
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

type managingElement struct {
	mBus *bus

	fetched uint8

	addrAbs uint16
	addrRel uint16

	opcode uint8
	cycles uint8

	lookup []instruction
}

func (manEl *managingElement) fetch() uint8 {
	addrmodeName := manEl.lookup[manEl.opcode].addrmodeName
	if addrmodeName != "imp" {
		log.Debugf(
			"fetch() fetch from addrAbs, addrAbs: %x",
			manEl.addrAbs,
		)
		manEl.fetched = manEl.mBus.Read(manEl.addrAbs, true)
	}
	log.Debugf(
		"fetch() fetched: %x",
		manEl.fetched,
	)
	return manEl.fetched
}

func CreateManagingElement(mBus *bus) *managingElement {
	manEl := &managingElement{}

	manEl.fetched = 0x00

	manEl.addrAbs = 0x0000
	manEl.addrRel = 0x0000

	manEl.opcode = 0x00
	manEl.cycles = 0x00

	manEl.mBus = mBus

	return manEl
}
