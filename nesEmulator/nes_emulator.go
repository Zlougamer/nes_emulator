package main

import (
    "flag"
    "os"
    "text/template"
)

type data struct {
    Type string
    Name string
}

func main() {
    var d data
    flag.StringVar(&d.Type, "type", "", "The subtype used for the queue being generated")
    flag.StringVar(&d.Name, "name", "", "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
    flag.Parse()

    t := template.Must(template.New("queue").Parse(queueTemplate))
    if err := t.Execute(os.Stdout, d); err != nil {
        panic(err)
    }
}

var queueTemplate = `
package queue

import (
  "container/list"
)

func New{{.Name}}() *{{.Name}} {
  return &{{.Name}}{list.New()}
}

type {{.Name}} struct {
  list *list.List
}

func (q *{{.Name}}) Len() int {
  return q.list.Len()
}

func (q *{{.Name}}) Enqueue(i {{.Type}}) {
  q.list.PushBack(i)
}

func (q *{{.Name}}) Dequeue() {{.Type}} {
  if q.list.Len() == 0 {
    panic(ErrEmptyQueue)
  }
  raw := q.list.Remove(q.list.Front())
  if typed, ok := raw.({{.Type}}); ok {
    return typed
  }
  panic(ErrInvalidType)
}
`


//package main
//
//import "fmt"
//
//func main() {
//    // TODO:
//    // 1. Read about logging
//    // 2. Write tests for different cases
//    // 3. Refactor instructions and adressing modes
//    // 4. Support decompiler and graphical render as in video
//
//    //b := cpu.CreateBus()
//    //olc := cpu.CreateOlc6502()
//    //b.Read(0, false)
//    //b.Write(0, 0)
//    fmt.Println("Hello, nes emulator!")
//    // Read about logging in go
//    //cpu.Read(1)
//    //cpu.Write(1, 1)
//}
