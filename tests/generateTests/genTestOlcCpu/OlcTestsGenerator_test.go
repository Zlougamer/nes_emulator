package genTestOlcCpu

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "path"
    "path/filepath"
    "strings"
    "testing"
    "text/template"
)

func TestGenerateAdcTests(t *testing.T) {
    staticInfoPath := getStaticInfoTestsOlcFilepath()
    files, err := filepath.Glob(path.Join(staticInfoPath, "*"))
    check(err)

    for _, f := range files {
        GenerateTestFromJson(f)
    }
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

    fOut, err := os.OpenFile(adcTestsPath, os.O_RDWR|os.O_CREATE, 0755)
    check(err)
    defer closeFile(fOut)

    templ := template.Must(template.New("oneTestTemplate").Parse(oneTestTemplate))
    if err := templ.Execute(fOut, data); err != nil {
        panic(err)
    }

}
