package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type IO struct {
	s *bufio.Scanner
	w *bufio.Writer
}

func NewIO() *IO {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	s.Split(bufio.ScanWords)

	return &IO{s: s, w: w}
}

func (io *IO) Flush() {
	io.w.Flush()
}

func (io *IO) ReadInput() (int, []int) {
	n := 0
	nums := make([]int, n)

	io.s.Scan()
	n, _ = strconv.Atoi(io.s.Text())

	for i := 0; i < n; i++ {
		io.s.Scan()
		nums[i], _ = strconv.Atoi(io.s.Text())
	}

	return n, nums
}

func (io *IO) WriteOutput(res int) {
	fmt.Fprintf(io.w, "%v", res)
}

func main() {
	io := NewIO()
	defer io.Flush()

	n, nums := io.ReadInput()

	res := QualityControl(n, nums)

	io.WriteOutput(res)
}

func QualityControl(n int, nums []int) int {
	prev := -1

	for _, v := range nums {
		if v%2 == 0 {
			prev = v
		}
	}

	return prev
}
