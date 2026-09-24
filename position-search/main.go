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

	return &IO{
		s: s,
		w: w,
	}
}

func (o *IO) Close() error {
	return o.w.Flush()
}

func (o *IO) ReadConditions() (int, []int, int) {
	n := 0
	nums := make([]int, n)
	target := 0

	o.s.Scan()
	n, _ = strconv.Atoi(o.s.Text())

	for i := 0; i < n; i++ {
		o.s.Scan()
		nums[i], _ = strconv.Atoi(o.s.Text())
	}

	o.s.Scan()
	target, _ = strconv.Atoi(o.s.Text())

	return n, nums, target
}

func (o *IO) WriteResult(res int) {
	fmt.Fprintf(o.w, "%v", res)
}

func main() {
	o := NewIO()
	defer o.Close()

	n, nums, target := o.ReadConditions()

	res := PositionSearch(n, nums, target)

	o.WriteResult(res)
}

func PositionSearch(n int, nums []int, target int) int {
	return 0
}
