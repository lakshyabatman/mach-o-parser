package entities

import (
	"fmt"
	"reflect"

	"github.com/lakshyabatman/mach-o-parser/parser"
)

type DyldChainedFixupsCommand struct {
	DataOff  uint32
	Datasize uint32
}

func (d *DyldChainedFixupsCommand) Print() {
	fmt.Printf("	DataOff: %d\n", d.DataOff)
	fmt.Printf("	Datasize: %d\n", d.Datasize)
}

func ParseDyldChainedFixups(p *parser.Parser, data chan any) {
	dyldChainedFixups := DyldChainedFixupsCommand{}

	p.ParseLiteral(data, 4, reflect.Uint32)
	dyldChainedFixups.DataOff = (<-data).(uint32)

	p.ParseLiteral(data, 4, reflect.Uint32)
	dyldChainedFixups.Datasize = (<-data).(uint32)

	data <- &dyldChainedFixups
}
