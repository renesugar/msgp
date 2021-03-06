// Package msgp is the runtime support library for the msgp code generator (http://github.com/dchenk/msgp).
//
// This package defines the utilities used by the msgp code generator for encoding and decoding MessagePack
// from []byte and io.Reader/io.Writer types. Most things here are intended to be used only in programs that
// use the msgp code generator, the point being to avoid runtime reflection.
//
// This package defines four families of functions:
// 	- AppendXxxx() appends an object to a []byte in MessagePack encoding.
// 	- ReadXxxxBytes() reads an object from a []byte and returns the remaining bytes.
// 	- (*Writer).WriteXxxx() writes an object to the buffered *Writer type.
// 	- (*Reader).ReadXxxx() reads an object from a buffered *Reader type.
//
// Types that implement the msgp.Encoder and msgp.Decoder interfaces can be written and read from any io.Writer
// and io.Reader types using
// 		msgp.Encode(io.Writer, msgp.Encoder)
// and
//		msgp.Decode(io.Reader, msgp.Decoder)
//
// There are also methods for converting MessagePack to JSON without an explicit de-serialization step.
//
// For more tips and tricks please visit the wiki at http://github.com/dchenk/msgp
//
package msgp

// Here are all the byte prefixes in the MessagePack standard.
const (
	mfixint   uint8 = 0x00 // 0XXXXXXX
	mnfixint  uint8 = 0xe0 // 111XXXXX
	mfixmap   uint8 = 0x80 // 1000XXXX
	mfixarray uint8 = 0x90 // 1001XXXX
	mfixstr   uint8 = 0xa0 // 101XXXXX
	mnil      uint8 = 0xc0
	mfalse    uint8 = 0xc2
	mtrue     uint8 = 0xc3
	mbin8     uint8 = 0xc4
	mbin16    uint8 = 0xc5
	mbin32    uint8 = 0xc6
	mext8     uint8 = 0xc7
	mext16    uint8 = 0xc8
	mext32    uint8 = 0xc9
	mfloat32  uint8 = 0xca
	mfloat64  uint8 = 0xcb
	muint8    uint8 = 0xcc
	muint16   uint8 = 0xcd
	muint32   uint8 = 0xce
	muint64   uint8 = 0xcf
	mint8     uint8 = 0xd0
	mint16    uint8 = 0xd1
	mint32    uint8 = 0xd2
	mint64    uint8 = 0xd3
	mfixext1  uint8 = 0xd4
	mfixext2  uint8 = 0xd5
	mfixext4  uint8 = 0xd6
	mfixext8  uint8 = 0xd7
	mfixext16 uint8 = 0xd8
	mstr8     uint8 = 0xd9
	mstr16    uint8 = 0xda
	mstr32    uint8 = 0xdb
	marray16  uint8 = 0xdc
	marray32  uint8 = 0xdd
	mmap16    uint8 = 0xde
	mmap32    uint8 = 0xdf
)

const (
	last4  = 0x0f
	first4 = 0xf0
	last5  = 0x1f
	first3 = 0xe0
	last7  = 0x7f
)

func isfixint(b byte) bool {
	return b>>7 == 0
}

func isnfixint(b byte) bool {
	return b&first3 == mnfixint
}

func isfixmap(b byte) bool {
	return b&first4 == mfixmap
}

func isfixarray(b byte) bool {
	return b&first4 == mfixarray
}

func isfixstr(b byte) bool {
	return b&first3 == mfixstr
}

func wfixint(u uint8) byte {
	return u & last7
}

func rfixint(b byte) uint8 {
	return b
}

func wnfixint(i int8) byte {
	return byte(i) | mnfixint
}

func rnfixint(b byte) int8 {
	return int8(b)
}

func rfixmap(b byte) uint8 {
	return b & last4
}

func wfixmap(u uint8) byte {
	return mfixmap | (u & last4)
}

func rfixstr(b byte) uint8 {
	return b & last5
}

func wfixstr(u uint8) byte {
	return (u & last5) | mfixstr
}

func rfixarray(b byte) uint8 {
	return b & last4
}

func wfixarray(u uint8) byte {
	return (u & last4) | mfixarray
}
