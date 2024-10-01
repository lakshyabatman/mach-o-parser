package entities

import (
	"fmt"
	"reflect"

	"github.com/lakshyabatman/mach-o-parser/parser"
)

type Command uint32

const LC_REQ_DYLD = 0x80000000

const (
	Segment                Command = 0x1
	Symtab                 Command = 0x2
	Symseg                 Command = 0x3
	Thread                 Command = 0x4
	Unixthread             Command = 0x5
	Loadfvmlib             Command = 0x6
	Idfvmlib               Command = 0x7
	Ident                  Command = 0x8
	Fvmfile                Command = 0x9
	Prepage                Command = 0xa
	Dysymtab               Command = 0xb
	LoadDylib              Command = 0xc
	IdDylib                Command = 0xd
	LoadDylinker           Command = 0xe
	IdDylinker             Command = 0xf
	PreboundDylib          Command = 0x10
	Routines               Command = 0x11
	SubFramework           Command = 0x12
	SubUmbrella            Command = 0x13
	SubClient              Command = 0x14
	SubLibrary             Command = 0x15
	TwolevelHints          Command = 0x16
	PrebindCksum           Command = 0x17
	LoadWeakDylib          Command = (0x18 | LC_REQ_DYLD)
	Segment64              Command = 0x19
	Routines64             Command = 0x1a
	UUID                   Command = 0x1b
	Rpath                  Command = (0x1c | LC_REQ_DYLD)
	CodeSignature          Command = 0x1d
	SegmentSplitInfo       Command = 0x1e
	ReexportDylib          Command = (0x1f | LC_REQ_DYLD)
	LazyLoadDylib          Command = 0x20
	EncryptionInfo         Command = 0x21
	DyldInfo               Command = 0x22
	DyldInfoOnly           Command = (0x22 | LC_REQ_DYLD)
	LoadUpwardDylib        Command = (0x23 | LC_REQ_DYLD)
	VersionMinMacosx       Command = 0x24
	VersionMinIphoneos     Command = 0x25
	FunctionStarts         Command = 0x26
	DyldEnvironment        Command = 0x27
	Main                   Command = (0x28 | LC_REQ_DYLD)
	DataInCode             Command = 0x29
	SourceVersion          Command = 0x2A
	DylibCodeSignDrs       Command = 0x2B
	EncryptionInfo64       Command = 0x2C
	LinkerOption           Command = 0x2D
	LinkerOptimizationHint Command = 0x2E
	VersionMinTvos         Command = 0x2F
	VersionMinWatchos      Command = 0x30
	Note                   Command = 0x31
	BuildVersion           Command = 0x32
	DyldExportsTrie        Command = (0x33 | LC_REQ_DYLD)
	DyldChainedFixups      Command = (0x34 | LC_REQ_DYLD)
	FilesetEntry           Command = (0x35 | LC_REQ_DYLD)
)

func CommandToString(cmd Command) string {
	switch cmd {
	case Segment:
		return "Segment"
	case Symtab:
		return "Symtab"
	case Symseg:
		return "Symseg"
	case Thread:
		return "Thread"
	case Unixthread:
		return "Unixthread"
	case Loadfvmlib:
		return "Loadfvmlib"
	case Idfvmlib:
		return "Idfvmlib"
	case Ident:
		return "Ident"
	case Fvmfile:
		return "Fvmfile"
	case Prepage:
		return "Prepage"
	case Dysymtab:
		return "Dysymtab"
	case LoadDylib:
		return "LoadDylib"
	case IdDylib:
		return "IdDylib"
	case LoadDylinker:
		return "LoadDylinker"
	case IdDylinker:
		return "IdDylinker"
	case PreboundDylib:
		return "PreboundDylib"
	case Routines:
		return "Routines"
	case SubFramework:
		return "SubFramework"
	case SubUmbrella:
		return "SubUmbrella"
	case SubClient:
		return "SubClient"
	case SubLibrary:
		return "SubLibrary"
	case TwolevelHints:
		return "TwolevelHints"
	case PrebindCksum:
		return "PrebindCksum"
	case LoadWeakDylib:
		return "LoadWeakDylib"
	case Segment64:
		return "Segment64"
	case Routines64:
		return "Routines64"
	case UUID:
		return "UUID"
	case Rpath:
		return "Rpath"
	case CodeSignature:
		return "CodeSignature"
	case SegmentSplitInfo:
		return "SegmentSplitInfo"
	case ReexportDylib:
		return "ReexportDylib"
	case LazyLoadDylib:
		return "LazyLoadDylib"
	case EncryptionInfo:
		return "EncryptionInfo"
	case DyldInfo:
		return "DyldInfo"
	case DyldInfoOnly:
		return "DyldInfoOnly"
	case LoadUpwardDylib:
		return "LoadUpwardDylib"
	case VersionMinMacosx:
		return "VersionMinMacosx"
	case VersionMinIphoneos:
		return "VersionMinIphoneos"
	case FunctionStarts:
		return "FunctionStarts"
	case DyldEnvironment:
		return "DyldEnvironment"
	case Main:
		return "Main"
	case DataInCode:
		return "DataInCode"
	case SourceVersion:
		return "SourceVersion"
	case DylibCodeSignDrs:
		return "DylibCodeSignDrs"
	case EncryptionInfo64:
		return "EncryptionInfo64"
	case LinkerOption:
		return "LinkerOption"
	case LinkerOptimizationHint:
		return "LinkerOptimizationHint"
	case VersionMinTvos:
		return "VersionMinTvos"
	case VersionMinWatchos:
		return "VersionMinWatchos"
	case Note:
		return "Note"
	case BuildVersion:
		return "BuildVersion"
	case DyldExportsTrie:
		return "DyldExportsTrie"
	case DyldChainedFixups:
		return "DyldChainedFixups"
	case FilesetEntry:
		return "FilesetEntry"
	default:
		return "Unknown"
	}
}

type LoadCommand struct {
	Cmd         string
	CmdSize     uint32
	CommandBody interface{}
}

func (command *LoadCommand) Print() {
	fmt.Printf("Command name: %s\n", command.Cmd)
	fmt.Printf("Command size: %d\n", command.CmdSize)

}

func ParseLoadCommand(p *parser.Parser, data chan any) {
	p.ParseLiteral(data, 4, reflect.Uint32)
	var c Command
	loadCommand := LoadCommand{}
	v, ok := (<-data).(uint32)
	if ok {
		c = Command(v)
		loadCommand.Cmd = CommandToString(Command(v))
	}
	p.ParseLiteral(data, 4, reflect.Uint32)
	v, ok = (<-data).(uint32)
	if ok {
		loadCommand.CmdSize = v
	}
	loadCommand.Print()
	parseCommand(c, loadCommand.CmdSize, p, data)
	// select {
	// // case v, _ := <-data:
	// // 	loadCommand.CommandBody = v
	// // 	data <- loadCommand

	// // }
	// if len(data) != 0 {
	// 	loadCommand.CommandBody = <-data

	// }
}

func parseCommand(c Command, cmdSize uint32, p *parser.Parser, data chan any) {
	switch c {
	case Segment64:
		parseSegmentCommand64(p, data)
	case DyldChainedFixups:
		ParseDyldChainedFixups(p, data)
	default:
		p.Skip((cmdSize - 8))

	}

	if len(data) != 0 {
		entity, ok := (<-data).(Entity)
		if ok {
			entity.Print()
		}
	}

}
