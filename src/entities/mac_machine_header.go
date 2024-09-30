package entities

import (
	"fmt"
	"reflect"

	"github.com/lakshyabatman/mach-o-parser/src/parser"
)

type MachHeader struct {
	Magic      uint32
	CpuType    uint32
	CpuSubType uint32
	FileType   uint32
	Ncmds      uint32
	Sizeofcmds uint32
	Flags      Flags
	Reserved   uint32
}

func ParseMachHeader(p *parser.Parser, ch chan any) {
	mch := MachHeader{}

	p.ParseLiteral(ch, 4, reflect.Uint32)
	mch.Magic = uint32((<-ch).(uint32))
	p.ParseLiteral(ch, 4, reflect.Uint32)
	mch.CpuType = uint32((<-ch).(uint32))

	p.ParseLiteral(ch, 4, reflect.Uint32)
	mch.CpuSubType = uint32((<-ch).(uint32))

	p.ParseLiteral(ch, 4, reflect.Uint32)
	mch.FileType = uint32((<-ch).(uint32))

	p.ParseLiteral(ch, 4, reflect.Uint32)
	mch.Ncmds = uint32((<-ch).(uint32))

	p.ParseLiteral(ch, 4, reflect.Uint32)
	mch.Sizeofcmds = uint32((<-ch).(uint32))

	ParseIntoFlag(p, ch)
	if v, ok := (<-ch).(Flags); ok {
		mch.Flags = v
	}

	p.ParseLiteral(ch, 4, reflect.Uint32)
	mch.Reserved = uint32((<-ch).(uint32))
	ch <- mch
}

func (m *MachHeader) Print() {

	fmt.Println("Printing Mach header")
	fmt.Println("---------------------")

	fmt.Printf("	Magic: 0x%x \n", m.Magic)

	fmt.Printf("	CPU type: %d\n", m.CpuType)

	fmt.Printf("	CPU sub type: %d\n", m.CpuSubType)

	fmt.Printf("	File type: %d\n", m.FileType)

	fmt.Printf("	Number of commands: %d\n", m.Ncmds)

	fmt.Printf("	Size of commands: %d\n", m.Sizeofcmds)

	fmt.Println("	Printing flag")
	m.Flags.PrintFlags()

	fmt.Printf("	Reserved: %d\n", m.Reserved)

	fmt.Println("---------------------")
}
