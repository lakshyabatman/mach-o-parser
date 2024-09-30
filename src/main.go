package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lakshyabatman/mach-o-parser/src/entities"
	"github.com/lakshyabatman/mach-o-parser/src/parser"
)

func main() {
	if len(os.Args) == 1 {
		panic(errors.New("file argument not provided"))
	}
	filePath := os.Args[1]

	data, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}
	p := parser.Parser{Data: data}
	ch := make(chan interface{}, 1)
	defer close(ch)
	entities.ParseMachHeader(&p, ch)
	mch := (<-ch).(entities.MachHeader)
	mch.Print()
	fmt.Println("Printing sections")
	fmt.Println("---------------------")

	for i := 0; i < int(mch.Ncmds); i++ {
		entities.ParseLoadCommand(&p, ch)
	}
}
