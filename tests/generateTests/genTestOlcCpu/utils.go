package genTestOlcCpu

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

type MemoryEntry struct {
	Addr string `json:"Addr"`
	Data string `json:"Data"`
}

type TestCase struct{
	TestName string `json:"TestName"`
	InitAccum string `json:"InitAccum"`
	InitStatus string `json:"InitStatus"`
	InitX string `json:"InitX"`
	InitY string `json:"InitY"`
	MemoryWrite []MemoryEntry `json:"MemoryWrite"`
	ExpPc string `json:"ExpPc"`
	ExpAccum string `json:"ExpAccum"`
	ExpC string `json:"ExpC"`
	ExpN string `json:"ExpN"`
	ExpZ string `json:"ExpZ"`
	ExpV string `json:"ExpV"`
	ExpMemory []MemoryEntry `json:"ExpMemory"`
}

type TestSuite struct {
	TestCaseArr []TestCase `json:"TestCaseArr"`
}

func check(e error) {
	if e != nil {
		fmt.Printf("error: %v\n", e)
		panic(e)
	}
}

func getRootFilepath() string {
	_, filename, _, ok := runtime.Caller(1)
	if ok != true {
		panic("runtime.Caller failed")
	}
	rootPath := path.Join(path.Dir(filename), "..", "..", "..")
	return rootPath
}

func getTestsOlcFilepath() string {
	rootPath := getRootFilepath()
	testsPath := path.Join(rootPath, "tests", "testOlcCpu")
	return testsPath
}

func getStaticInfoTestsOlcFilepath() string {
	rootPath := getRootFilepath()
	testsPath := path.Join(rootPath, "staticInfo", "tests", "testOlcCpu")
	return testsPath
}


func closeFile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()
	check(err)
}

var oneTestTemplate = `// This file is generated! Do not edit it!!!!
package testOlcCpu

import (
	"testing"
	"github.com/Zlougamer/nes_emulator/olcCpu"
)


{{range .TestCaseArr}}
func Test{{.TestName}}(t *testing.T) {
    regSet := olcCpu.CreateRegisterSet()
    mpu := olcCpu.CreateOlc6502ByParams(regSet, nil)

    regSet.A = {{.InitAccum}}
    regSet.Status = {{.InitStatus}}
    regSet.X = {{.InitX}}
    regSet.Y = {{.InitY}}

    {{range .MemoryWrite}}write(mpu, uint16({{.Addr}}), []uint8{{"{"}}{{.Data}}{{"}"}})
    {{end}}
    mpu.Clock()

	assertEqual(t, uint16({{.ExpPc}}), regSet.Pc)
	assertEqual(t, uint8({{.ExpAccum}}), regSet.A)
	assertEqual(t, {{.ExpC}}, regSet.Status & olcCpu.C != 0)
	assertEqual(t, {{.ExpN}}, regSet.Status & olcCpu.N != 0)
	assertEqual(t, {{.ExpZ}}, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, {{.ExpV}}, regSet.Status & olcCpu.V != 0)
    {{range .ExpMemory}}
	assertEqual(t, uint8({{.Data}}), mpu.Read(uint16({{.Addr}})))
    {{end}}
}

{{end}}

`
