package main

import (
	"fmt"
	"strings"
)

// A pointer holds he memeory address of a value.
func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // point to print i through the pointer
	fmt.Println(p)  // print the memory location
	*p = 21         // set i throught the pointer
	fmt.Println(i)  // set the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide through the pointer
	fmt.Println(j) // see the new value of j

	var jon *string
	jonStr := "jon"
	jon = &jonStr
	fmt.Println(*jon)    // "jon"
	fmt.Println(jon)     // 0xc42000e2b0
	fmt.Println(jonStr)  // "jon"
	fmt.Println(&jonStr) // 0xc42000e2b0
}

func structs() {
	// A struct is a collection of fields.
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	p := &v // accessing structs through a pointer
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(p.X)    // no need for explicit (*p)
	fmt.Println((*p).X) // need the parens.

	var (
		v1      = Vertex{1, 2}
		v2      = Vertex{X: 1}  // Y:0 is implicit
		v3      = Vertex{}      // X:0 and Y:0
		pointer = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, pointer, v2, v3)
}

// The type [n]T is an array of n values of type T.
func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Unlike arrays which have fixed sizes. A slice is dynamically-sized.
func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int
	s = primes[1:4]
	fmt.Println(s)
}

func slicesAreRefs() {
	names := [4]string{
		"Eugene",
		"Tae",
		"P",
		"I",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	c := names[0:4]
	fmt.Println(a, b)
	fmt.Println(c)

	b[0] = "XXX" // This changes names[1]
	fmt.Println(a, b)
	fmt.Println(names)
}

// Array literal without the length.
// Creates an equivalent array and a slice that refers to it.
func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Length is number of elements the slice contains.
// Capacity is number of elements in underlying array.
func sliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	// You can exten a slice's length by re-slicing it, provided it has sufficient capacity.
	s = s[:4] // len 4. cap 6
	printSlice(s)

	s = s[2:] // len 2. cap 4
	printSlice(s)

	// Adjusting end value decreases len and not cap.
	// Adjusting starting value can drop both since it drops the values.
}

func nilSlice() {
	// Zero value of a slice is nil, with 0 len and cap. No underlying array.
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeSlice() {
	a := make([]int, 5) // [0, 0, 0, 0, 0]
	fmt.Println("a")
	printSlice(a)

	b := make([]int, 0, 5) // [] with cap 5
	fmt.Println("b")
	printSlice(b)

	c := b[:2] // [0, 0] with cap 5
	fmt.Println("c")
	printSlice(c)

	d := c[2:5] // [0, 0, 0] with cap 3 since we dropped from the front.
	fmt.Println("d")
	printSlice(d)
}

func sliceOfSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Take turns playing.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// More on slices: https://blog.golang.org/go-slices-usage-and-internals
func appendToSlice() {

	var s []int
	printSlice(s)

	// Append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// Slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add multiple elements at the same time.
	s = append(s, 2, 3, 4) // Why does cap grow more than len?
	// https://stackoverflow.com/questions/38543825/appending-one-element-to-nil-slice-increases-capacity-by-two
	printSlice(s)
}

func loopOverRange() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	// first is index, second is copy of the element at that index.
	for i, v := range pow {
		fmt.Printf("%d, %d\n", i, v)
		fmt.Printf("2**%d = %d\n", i, v)
	}

}

func moreRanges() {
	new := make([]int, 10)
	for i := range new {
		fmt.Printf("before: %d is i. %d is new[i], and %v is new\n", i, new[i], new)
		// Binary shift. https://stackoverflow.com/questions/5801008/go-and-operators
		new[i] = 1 << uint(i) // == 2**i
		fmt.Printf("after : %d is i. %d is new[i], and %v is new\n", i, new[i], new)
	}
	for _, value := range new {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	// pointers()
	// structs()
	// arrays()
	// slices()
	// slicesAreRefs()
	// sliceLiterals()
	// sliceDefaults()
	// sliceLengthAndCapacity()
	// nilSlice()
	// makeSlice()
	// sliceOfSlice()
	// appendToSlice()
	// loopOverRange()
	moreRanges()
}
