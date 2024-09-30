package entities

import (
	"encoding/binary"
	"fmt"

	"github.com/lakshyabatman/mach-o-parser/src/parser"
)

type Flags struct {
	Noundefs                   bool
	IncrLink                   bool
	DyldLink                   bool
	Bindatload                 bool
	Prebound                   bool
	SplitSegs                  bool
	LazyInit                   bool
	TwoLevel                   bool
	ForceFlat                  bool
	Nomultidefs                bool
	Nofixprebindin             bool
	Prebindable                bool
	Allmodsbound               bool
	SubsectionsViaSymbols      bool
	Canonical                  bool
	WeakDefines                bool
	BindsToWeak                bool
	AllowStackExecution        bool
	RootSafe                   bool
	SetuidSafe                 bool
	NoReexportedDylibs         bool
	Pie                        bool
	DeadStrippableDylib        bool
	HasTlvDescriptors          bool
	NoHeapExecution            bool
	AppExtensionSafe           bool
	NlistOutofsyncWithDyldinfo bool
	SimSupport                 bool
	DylibInCache               bool
}

func ParseIntoFlag(p *parser.Parser, data chan any) chan any {
	b := p.Data[0:4]
	p.Data = p.Data[4:]
	flags := Flags{}
	flagInt := binary.LittleEndian.Uint32(b)
	flags.Noundefs = (flagInt&1 == 1)
	flagInt >>= 1
	flags.IncrLink = (flagInt&1 == 1)
	flagInt >>= 1
	flags.DyldLink = (flagInt&1 == 1)
	flagInt >>= 1
	flags.Bindatload = (flagInt&1 == 1)
	flagInt >>= 1
	flags.Prebound = (flagInt&1 == 1)
	flagInt >>= 1
	flags.SplitSegs = (flagInt&1 == 1)
	flagInt >>= 1
	flags.LazyInit = (flagInt&1 == 1)
	flagInt >>= 1
	flags.TwoLevel = (flagInt&1 == 1)
	flagInt >>= 1
	flags.ForceFlat = (flagInt&1 == 1)
	flagInt >>= 1
	flags.Nomultidefs = (flagInt&1 == 1)
	flagInt >>= 1
	flags.Nofixprebindin = (flagInt&1 == 1)
	flagInt >>= 1
	flags.Prebindable = (flagInt&1 == 1)
	flagInt >>= 1
	flags.Allmodsbound = (flagInt&1 == 1)
	flagInt >>= 1
	flags.SubsectionsViaSymbols = (flagInt&1 == 1)
	flagInt >>= 1
	flags.Canonical = (flagInt&1 == 1)
	flagInt >>= 1
	flags.WeakDefines = (flagInt&1 == 1)
	flagInt >>= 1
	flags.BindsToWeak = (flagInt&1 == 1)
	flagInt >>= 1
	flags.AllowStackExecution = (flagInt&1 == 1)
	flagInt >>= 1
	flags.RootSafe = (flagInt&1 == 1)
	flagInt >>= 1
	flags.SetuidSafe = (flagInt&1 == 1)
	flagInt >>= 1
	flags.NoReexportedDylibs = (flagInt&1 == 1)
	flagInt >>= 1
	flags.Pie = (flagInt&1 == 1)
	flagInt >>= 1
	flags.DeadStrippableDylib = (flagInt&1 == 1)
	flagInt >>= 1
	flags.HasTlvDescriptors = (flagInt&1 == 1)
	flagInt >>= 1
	flags.NoHeapExecution = (flagInt&1 == 1)
	flagInt >>= 1
	flags.AppExtensionSafe = (flagInt&1 == 1)
	flagInt >>= 1
	flags.NlistOutofsyncWithDyldinfo = (flagInt&1 == 1)
	flagInt >>= 1
	flags.SimSupport = (flagInt&1 == 1)
	flagInt >>= 1
	flags.DylibInCache = (flagInt&1 == 1)
	flagInt >>= 1
	data <- flags
	return data

}

// Method to print the Flags struct
func (f *Flags) PrintFlags() {
	fmt.Printf("		Noundefs: %v\n", f.Noundefs)
	fmt.Printf("		IncrLink: %v\n", f.IncrLink)
	fmt.Printf("		DyldLink: %v\n", f.DyldLink)
	fmt.Printf("		Bindatload: %v\n", f.Bindatload)
	fmt.Printf("		Prebound: %v\n", f.Prebound)
	fmt.Printf("		SplitSegs: %v\n", f.SplitSegs)
	fmt.Printf("		LazyInit: %v\n", f.LazyInit)
	fmt.Printf("		TwoLevel: %v\n", f.TwoLevel)
	fmt.Printf("		ForceFlat: %v\n", f.ForceFlat)
	fmt.Printf("		Nomultidefs: %v\n", f.Nomultidefs)
	fmt.Printf("		Nofixprebindin: %v\n", f.Nofixprebindin)
	fmt.Printf("		Prebindable: %v\n", f.Prebindable)
	fmt.Printf("		Allmodsbound: %v\n", f.Allmodsbound)
	fmt.Printf("		SubsectionsViaSymbols: %v\n", f.SubsectionsViaSymbols)
	fmt.Printf("		Canonical: %v\n", f.Canonical)
	fmt.Printf("		WeakDefines: %v\n", f.WeakDefines)
	fmt.Printf("		BindsToWeak: %v\n", f.BindsToWeak)
	fmt.Printf("		AllowStackExecution: %v\n", f.AllowStackExecution)
	fmt.Printf("		RootSafe: %v\n", f.RootSafe)
	fmt.Printf("		SetuidSafe: %v\n", f.SetuidSafe)
	fmt.Printf("		NoReexportedDylibs: %v\n", f.NoReexportedDylibs)
	fmt.Printf("		Pie: %v\n", f.Pie)
	fmt.Printf("		DeadStrippableDylib: %v\n", f.DeadStrippableDylib)
	fmt.Printf("		HasTlvDescriptors: %v\n", f.HasTlvDescriptors)
	fmt.Printf("		NoHeapExecution: %v\n", f.NoHeapExecution)
	fmt.Printf("		AppExtensionSafe: %v\n", f.AppExtensionSafe)
	fmt.Printf("		NlistOutofsyncWithDyldinfo: %v\n", f.NlistOutofsyncWithDyldinfo)
	fmt.Printf("		SimSupport: %v\n", f.SimSupport)
	fmt.Printf("		DylibInCache: %v\n", f.DylibInCache)
}
