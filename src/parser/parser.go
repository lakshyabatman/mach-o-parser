package parser

import (
	"encoding/binary"
	"errors"
	"reflect"
)

type Parser struct {
	Data   []uint8
	Offset int
}

func (p *Parser) ParseLiteral(data chan any, size int, kind reflect.Kind) chan any {
	if len(p.Data) < size {
		panic(errors.New("not enough bytes"))
	}
	bytes := p.Data[0:size]
	p.Data = p.Data[size:]
	p.Offset += size
	switch {
	case kind == reflect.String:
		data <- string(bytes[:])
	case kind == reflect.Uint16:
		data <- binary.LittleEndian.Uint16(bytes)
	case kind == reflect.Uint32:
		data <- binary.LittleEndian.Uint32(bytes)
	case kind == reflect.Uint64:
		data <- binary.LittleEndian.Uint64(bytes)
	default:
		data <- bytes
	}
	return data
}

func (p *Parser) Skip(n uint32) {
	if len(p.Data) < int(n) {
		panic(errors.New("out of index while skipping"))
	}
	p.Data = p.Data[n:]
}
