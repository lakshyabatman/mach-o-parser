package entities

import (
	"fmt"
	"reflect"

	"github.com/lakshyabatman/mach-o-parser/src/parser"
)

type SegmentCommand64 struct {
	Segname  string
	Vmaddr   uint64
	VmSize   uint64
	FileOff  uint64
	FileSize uint64
	Maxprot  VmProt
	Initprot VmProt
	Nsects   uint32
	Flags    SegmentCommandFlags
	Sections []SegmentSection
}

type SegmentCommandFlags struct {
	HighVm            bool
	Fvmlib            bool
	Noreloc           bool
	ProtectedVersion1 bool
	ReadOnly          bool
}

type SegmentSection struct {
	SectName  string
	Segname   string
	Addr      uint64
	Size      uint64
	Offset    uint32
	Align     uint32
	Reloff    uint32
	Nreloc    uint32
	Flags     uint32
	Reserved1 uint32
	Reserved2 uint32
	Reserved3 uint32 // Added for 64-bit structure

}

// Method to print SegmentCommandFlags
func (f SegmentCommandFlags) Print() {
	fmt.Printf("		HighVm: %v\n", f.HighVm)
	fmt.Printf("		Fvmlib: %v\n", f.Fvmlib)
	fmt.Printf("		Noreloc: %v\n", f.Noreloc)
	fmt.Printf("		ProtectedVersion1: %v\n", f.ProtectedVersion1)
	fmt.Printf("		ReadOnly: %v\n", f.ReadOnly)
}

// Method to print SegmentSection
func (s SegmentSection) Print() {
	fmt.Printf("		SectName: %s\n", s.SectName)
	fmt.Printf("		Segname: %s\n", s.Segname)
	fmt.Printf("		Addr: %v\n", s.Addr)
	fmt.Printf("		Size: %v\n", s.Size)
	fmt.Printf("		Offset: %v\n", s.Offset)
	fmt.Printf("		Align: %v\n", s.Align)
	fmt.Printf("		Reloff: %v\n", s.Reloff)
	fmt.Printf("		Nreloc: %v\n", s.Nreloc)
	fmt.Printf("		Flags: %v\n", s.Flags)
	fmt.Printf("		Reserved1: %v\n", s.Reserved1)
	fmt.Printf("		Reserved2: %v\n", s.Reserved2)
	fmt.Printf("		Reserved3: %v\n", s.Reserved3)
}

// Method to print SegmentCommand64
func (c SegmentCommand64) Print() {
	fmt.Printf("	Segname: %s\n", c.Segname)
	fmt.Printf("	Vmaddr: %v\n", c.Vmaddr)
	fmt.Printf("	VmSize: %v\n", c.VmSize)
	fmt.Printf("	FileOff: %v\n", c.FileOff)
	fmt.Printf("	FileSize: %v\n", c.FileSize)
	fmt.Printf("	Maxprot: %v\n", c.Maxprot)
	fmt.Printf("	Initprot: %v\n", c.Initprot)
	fmt.Printf("	Nsects: %v\n", c.Nsects)

	// Print flags by calling the method on SegmentCommandFlags
	fmt.Println("	Flags:")
	c.Flags.Print()

	// Print sections by calling the method on each SegmentSection
	fmt.Println("	Sections:")
	for _, section := range c.Sections {
		section.Print()
	}
}

func parseSegmentCommand64(p *parser.Parser, data chan any) {
	segmentCommand64 := SegmentCommand64{}

	p.ParseLiteral(data, 16, reflect.String)
	segmentCommand64.Segname = (<-data).(string)

	p.ParseLiteral(data, 8, reflect.Uint64)
	segmentCommand64.VmSize = (<-data).(uint64)

	p.ParseLiteral(data, 8, reflect.Uint64)
	segmentCommand64.Vmaddr = (<-data).(uint64)

	p.ParseLiteral(data, 8, reflect.Uint64)
	segmentCommand64.FileOff = (<-data).(uint64)

	p.ParseLiteral(data, 8, reflect.Uint64)
	segmentCommand64.FileSize = (<-data).(uint64)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentCommand64.Maxprot = ParseVmProt((<-data).(uint32))

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentCommand64.Initprot = ParseVmProt((<-data).(uint32))

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentCommand64.Nsects = (<-data).(uint32)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentCommand64.Flags = ParseSegmentCommandFlag((<-data).(uint32))

	// fmt.Println(segmentCommand64)
	// skip sections here!

	for i := uint32(0); i < segmentCommand64.Nsects; i++ {
		parseSegmentSection(p, data)
		segmentCommand64.Sections = append(segmentCommand64.Sections, (<-data).(SegmentSection))
	}
	data <- &segmentCommand64

}

type VmProt struct {
	Read    bool
	Write   bool
	Execute bool
}

func ParseVmProt(u uint32) VmProt {
	vmProt := VmProt{
		Read:    ((u & 1) == 1),
		Write:   (((u >> 1) & 1) == 1),
		Execute: ((u >> 2) & 1) == 1,
	}
	return vmProt
}

func ParseSegmentCommandFlag(u uint32) SegmentCommandFlags {
	return SegmentCommandFlags{
		HighVm:            (u & 1) == 1,
		Fvmlib:            (u & 2) == 1,
		Noreloc:           (u & 4) == 1,
		ProtectedVersion1: (u & 8) == 1,
		ReadOnly:          (u & 16) == 1,
	}
}

func parseSegmentSection(p *parser.Parser, data chan any) {
	segmentSection := SegmentSection{}

	p.ParseLiteral(data, 16, reflect.String)
	segmentSection.SectName = (<-data).(string)

	p.ParseLiteral(data, 16, reflect.String)
	segmentSection.Segname = (<-data).(string)

	p.ParseLiteral(data, 8, reflect.Uint64)
	segmentSection.Addr = (<-data).(uint64)

	p.ParseLiteral(data, 8, reflect.Uint64)
	segmentSection.Size = (<-data).(uint64)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentSection.Offset = (<-data).(uint32)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentSection.Align = (<-data).(uint32)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentSection.Reloff = (<-data).(uint32)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentSection.Nreloc = (<-data).(uint32)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentSection.Flags = (<-data).(uint32)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentSection.Reserved1 = (<-data).(uint32)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentSection.Reserved2 = (<-data).(uint32)

	p.ParseLiteral(data, 4, reflect.Uint32)
	segmentSection.Reserved3 = (<-data).(uint32)
	data <- segmentSection

}
