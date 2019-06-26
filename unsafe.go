package main

// https://itnext.io/manipulating-private-fields-in-go-4da4ca525717
import (
	"fmt"
	"strings"
	"unsafe"
)

const ptrSize = unsafe.Sizeof(new(int))

func main() {
	//fmt.Printf("%i\n", ptrSize)
	bs := make([]byte, 0, 7)
	bs = append(bs, 'u', 'n', 's', 'a', 'f', 'e')

	var sb strings.Builder
	p := unsafe.Pointer(&sb)
	*(*[]byte)(unsafe.Pointer(uintptr(p) + uintptr(ptrSize))) = bs
	fmt.Println(sb.String())

	sb.WriteByte('!')
	fmt.Println(sb.String())

	fmt.Printf("%s\n", bs)
	fmt.Printf("%s\n", bs[:7])
	// Output:
	// unsafe
	// unsafe!
	// unsafe
	// unsafe!
}
