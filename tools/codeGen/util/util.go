package util

import (
	"bytes"
	"io"
	"log"
	big "math/big"
	"unsafe"
)

// TODO: finish
type CID struct {
}

// TODO: finish
// type BigInt struct {
// }

// Check that sizeof(int) == 8 (we only support 64-bit builds for now)
type _UnusedCompileAssert1 = [unsafe.Sizeof(int(0)) - 8]byte
type _UnusedCompileAssert2 = [8 - unsafe.Sizeof(int(0))]byte

type UVarint = uint64
type Varint = int64
type UInt = uint64
type UInts = []uint64 // Subject to change when array support comes.
type Float = float64
type BigInt = big.Int
type Bytes = []byte
type BytesArray = []Bytes // Subject to change when array support comes.
type BytesKey = string    // to use Bytes in map keys.
type BytesAmount = UVarint
type String = string
type T = struct{} // For use in generic definitions.

func Assert(b bool) {
	if !b {
		panic("Assertion failed")
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic("Error check failed")
	}
}

func CompareBytesStrict(x []byte, y []byte) int {
	panic("TODO")
}

func HashBlake2bInternal(x []byte) []byte {
	panic("TODO")
}

type Timestamp UVarint

func CurrentTime() Timestamp {
	panic("TODO")
}

func TextAbbrev(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func IntMin(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func IntMax(x, y int) int {
	if y > x {
		return y
	}
	return x
}

type IntOption struct {
	isSome bool
	value  int
}

func IntOptionSome(x int) IntOption {
	return IntOption{
		isSome: true,
		value:  x,
	}
}

func IntOptionNone() IntOption {
	return IntOption{
		isSome: false,
		value:  -1,
	}
}

func IntOptionMin(a, b IntOption) IntOption {
	if a.IsNone() || b.IsNone() {
		return IntOptionNone()
	} else {
		return IntOptionSome(IntMin(a.Get(), b.Get()))
	}
}

func IntOptionMax(a, b IntOption) IntOption {
	if a.IsNone() || b.IsNone() {
		return IntOptionNone()
	} else {
		return IntOptionSome(IntMax(a.Get(), b.Get()))
	}
}

func IntOptionAdd(a, b IntOption) IntOption {
	if a.IsNone() || b.IsNone() {
		return IntOptionNone()
	} else {
		return IntOptionSome(a.Get() + b.Get())
	}
}

func (x IntOption) IsSome() bool {
	return x.isSome
}

func (x IntOption) IsNone() bool {
	return !x.isSome
}

func (x IntOption) Get() int {
	Assert(x.IsSome())
	return x.value
}

func WriteRepeat(dst io.Writer, x string, n int) {
	for i := 0; i < n; i++ {
		dst.Write([]byte(x))
	}
}

func WriteRepeatString(x string, n int) string {
	buf := bytes.NewBuffer([]byte{})
	WriteRepeat(buf, x, n)
	return buf.String()
}

func SliceContainsString(s []string, x string) bool {
	for _, si := range s {
		if x == si {
			return true
		}
	}
	return false
}

func DerefCheckString(s *string) string {
	Assert(s != nil)
	return *s
}

func RefString(s string) *string {
	return &s
}
