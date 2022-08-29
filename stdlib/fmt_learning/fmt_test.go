package fmt_learning

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

var someHere = 2

var _ = struct {
	Bool       bool        `json:"Bool" xml:"Bool" toml:"Bool"`
	Int8       int8        `json:"Int8" xml:"Int8" toml:"Int8"`
	UInt8      uint8       `json:"UInt8" xml:"UInt8" toml:"UInt8"`
	Int16      int16       `json:"Int16" xml:"Int16" toml:"Int16"`
	UInt16     uint16      `json:"UInt16" xml:"UInt16" toml:"UInt16"`
	Int32      int32       `json:"Int32" xml:"Int32" toml:"Int32"`
	UInt32     uint32      `json:"UInt32" xml:"UInt32" toml:"UInt32"`
	Int        int         `json:"Int" xml:"Int" toml:"Int"`
	UInt       uint        `json:"UInt" xml:"UInt" toml:"UInt"`
	Int64      int64       `json:"Int64" xml:"Int64" toml:"Int64"`
	UInt64     uint64      `json:"UInt64" xml:"UInt64" toml:"UInt64"`
	Float32    float32     `json:"Float32" xml:"Float32" toml:"Float32"`
	Float64    float64     `json:"Float64" xml:"Float64" toml:"Float64"`
	Complex64  complex64   `json:"Complex64" xml:"Complex64" toml:"Complex64"`
	Complex128 complex128  `json:"Complex128" xml:"Complex128" toml:"Complex128"`
	String     string      `json:"String" xml:"String" toml:"String"`
	Uintptr    uintptr     `json:"Uinptr" xml:"Uinptr" toml:"Uinptr"`
	Byte       byte        `json:"Byte" xml:"Byte" toml:"Byte"`
	Rune       rune        `json:"Rune" xml:"Rune" toml:"Rune"`
	Any        any         `json:"Any" xml:"Any" toml:"Any"`
	Interface  interface{} `json:"Interface" xml:"Interface" toml:"Interface"`
	Nil        *int        `json:"Nil" xml:"Nil" toml:"Nil"`
}{
	Bool:       true,
	Int8:       1,
	UInt8:      2,
	Int16:      3,
	UInt16:     4,
	Int32:      5,
	UInt32:     6,
	Int:        7,
	UInt:       8,
	Int64:      9,
	UInt64:     10,
	Float32:    11.111,
	Float64:    12.2222222,
	Complex64:  13 + 14i, // We can also create them using the builtin complex64(13, 14)
	Complex128: 15 + 16i,
	String:     "Hello world",
	Uintptr:    uintptr(unsafe.Pointer(&someHere)),
	Byte:       0x43,
	Rune:       0x123456,
	Any:        struct{ a, b string }{a: "Hello", b: "World"},
	Nil:        nil,
}

func TestFmtBinaryRepresentation(t *testing.T) {
	for _, each := range []struct {
		description string
		input       any
		want        string
	}{
		{description: "int8 in binary representation", input: int8(39), want: "00100111"},
		{description: "uint8 in binary representation", input: uint8(39), want: "00100111"},
		{description: "int16 in binary representation", input: int16(27017), want: "0110100110001001"},                                                             // Mongodb Port
		{description: "uint16 in binary representation", input: uint16(27017), want: "0110100110001001"},                                                           // Mongodb Port
		{description: "int32 in binary representation", input: int32(281061476), want: "00010000110000001010100001100100"},                                         // 192.168.0.100
		{description: "uint32 in binary representation", input: uint32(281061476), want: "00010000110000001010100001100100"},                                       // 192.168.0.100
		{description: "int64 in binary representation", input: int64(151930230829876), want: "0000000000000000100010100010111000000011011100000111001100110100"},   // last 6 Bytes from IPv6 => 2001:0db8:0000:0000:0000:8a2e:0370:7334
		{description: "uint64 in binary representation", input: uint64(151930230829876), want: "0000000000000000100010100010111000000011011100000111001100110100"}, // last 6 Bytes from IPv6 => 2001:0db8:0000:0000:0000:8a2e:0370:7334
		{description: "int in binary representation", input: int(151930230829876), want: "0000000000000000100010100010111000000011011100000111001100110100"},       // last 6 Bytes from IPv6 => 2001:0db8:0000:0000:0000:8a2e:0370:7334
		{description: "uint in binary representation", input: uint(151930230829876), want: "0000000000000000100010100010111000000011011100000111001100110100"},     // last 6 Bytes from IPv6 => 2001:0db8:0000:0000:0000:8a2e:0370:7334
	} {
		t.Run(each.description, func(t *testing.T) {
			assert.Equal(t, each.want, fmt.Sprintf("%0"+strconv.Itoa(reflect.TypeOf(each.input).Bits())+"b", each.input))
		})
	}
}
