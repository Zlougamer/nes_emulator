package genTestOlcCpu

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "path"
    "path/filepath"
    "runtime"
    "strings"
    "testing"
    "text/template"
)

type MemoryEntry struct {
    Addr string `json:"Addr"`
    Data string `json:"Data"`
}

type TestCase struct{
    TestName string `json:"TestName"`
    InitAccum string `json:"InitAccum"`
    InitStkp string `json:"InitStkp"`
    InitPc string `json:"InitPc"`
    InitStatus string `json:"InitStatus"`
    InitX string `json:"InitX"`
    InitY string `json:"InitY"`
    MemoryWrite []MemoryEntry `json:"MemoryWrite"`
    ExpPc string `json:"ExpPc"`
    ExpAccum string `json:"ExpAccum"`
    ExpX string `json:"ExpX"`
    ExpY string `json:"ExpY"`
    ExpStkp string `json:"ExpStkp"`
    ExpC string `json:"ExpC"`
    ExpN string `json:"ExpN"`
    ExpZ string `json:"ExpZ"`
    ExpV string `json:"ExpV"`
    ExpI string `json:"ExpI"`
    ExpB string `json:"ExpB"`
    ExpU string `json:"ExpU"`
    ExpMemory []MemoryEntry `json:"ExpMemory"`
}

type TestSuite struct {
    TestCaseArr []TestCase `json:"TestCaseArr"`
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
    mpu := olcCpu.CreateOlc6502ByParams(regSet)

    regSet.A = {{.InitAccum}}
    {{if .InitStkp}}regSet.Stkp = {{.InitStkp}}
    {{end}}
    {{if .InitPc}}regSet.Pc = {{.InitPc}}
    {{end}}
    regSet.Status = {{.InitStatus}}
    regSet.X = {{.InitX}}
    regSet.Y = {{.InitY}}

    {{range .MemoryWrite}}write(mpu, uint16({{.Addr}}), []uint8{{"{"}}{{.Data}}{{"}"}})
    {{end}}
    mpu.Clock()

	assertEqual(t, uint16({{.ExpPc}}), regSet.Pc)
	assertEqual(t, uint8({{.ExpAccum}}), regSet.A)
    {{if .ExpX}}assertEqual(t, uint8({{.ExpX}}), regSet.X)
    {{end}}
    {{if .ExpY}}assertEqual(t, uint8({{.ExpY}}), regSet.Y)
    {{end}}
    {{if .ExpStkp}}assertEqual(t, uint8({{.ExpStkp}}), regSet.Stkp)
    {{end}}
	assertEqual(t, {{.ExpC}}, regSet.Status & olcCpu.C != 0)
	assertEqual(t, {{.ExpN}}, regSet.Status & olcCpu.N != 0)
	assertEqual(t, {{.ExpZ}}, regSet.Status & olcCpu.Z != 0)
	assertEqual(t, {{.ExpV}}, regSet.Status & olcCpu.V != 0)
    {{if .ExpI}}assertEqual(t, {{.ExpI}}, regSet.Status & olcCpu.I != 0)
    {{end}}
    {{if .ExpB}}assertEqual(t, {{.ExpB}}, regSet.Status & olcCpu.B != 0)
    {{end}}
    {{if .ExpU}}assertEqual(t, {{.ExpU}}, regSet.Status & olcCpu.U != 0)
    {{end}}
    {{range .ExpMemory}}assertEqual(t, uint8({{.Data}}), mpu.Read(uint16({{.Addr}})))
    {{end}}
}

{{end}}

`

func TestGenerateAdcTests(t *testing.T) {
    staticInfoPath := getStaticInfoTestsOlcFilepath()
    files, err := filepath.Glob(path.Join(staticInfoPath, "*"))
    check(err)

    for _, f := range files {
        GenerateTestFromJson(f)
    }
}

func getStaticInfoTestsOlcFilepath() string {
    rootPath := getRootFilepath()
    testsPath := path.Join(rootPath, "staticInfo", "tests", "testOlcCpu")
    return testsPath
}


func getRootFilepath() string {
    _, filename, _, ok := runtime.Caller(1)
    if ok != true {
        panic("runtime.Caller failed")
    }
    rootPath := path.Join(path.Dir(filename), "..", "..", "..")
    return rootPath
}


func GenerateTestFromJson(fIn string) {
    jsonFile, err := os.Open(fIn)
    check(err)
    defer closeFile(jsonFile)

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var data TestSuite

    err = json.Unmarshal(byteValue, &data)
    check(err)

    testsPath := getTestsOlcFilepath()
    basename := filepath.Base(fIn)
    name := strings.TrimSuffix(basename, filepath.Ext(basename))
    outputFilename := "generated" + strings.Title(name) + "_test.go"

    adcTestsPath := path.Join(testsPath, outputFilename)

    fOut, err := os.OpenFile(adcTestsPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
    check(err)
    defer closeFile(fOut)

    templ := template.Must(template.New("oneTestTemplate").Parse(oneTestTemplate))
    if err := templ.Execute(fOut, data); err != nil {
        panic(err)
    }
}


func closeFile(f *os.File) {
    fmt.Println("closing file")
    err := f.Close()
    check(err)
}


func check(e error) {
    if e != nil {
        fmt.Printf("error: %v\n", e)
        panic(e)
    }
}


func getTestsOlcFilepath() string {
    rootPath := getRootFilepath()
    testsPath := path.Join(rootPath, "tests", "testOlcCpu")
    return testsPath
}

