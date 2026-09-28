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
	o.s.Scan()
	n, _ := strconv.Atoi(o.s.Text())

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		o.s.Scan()
		nums[i], _ = strconv.Atoi(o.s.Text())
	}

	o.s.Scan()
	target, _ := strconv.Atoi(o.s.Text())

	return n, nums, target
}

func (o *IO) WriteResult(res []int) {
	for _, v := range res {
		fmt.Fprintf(o.w, "%v ", v)
	}
}

func main() {
	o := NewIO()
	defer o.Close()

	n, nums, target := o.ReadConditions()

	res := ExponentialSearch(n, nums, target)

	o.WriteResult(res)
}

func ExponentialSearch(n int, nums []int, target int) []int {
	if n == 0 {
		return []int{-1, -1}
	}

	if nums[0] > target {
		return []int{0, 0}
	}

	var i int

	for i = 1; i < n; i *= 2 {
		if nums[i] > target {
			break
		}
	}

	if i >= n {
		return []int{i / 2, n - 1}
	} else {
		return []int{i / 2, i}
	}
}

type OP struct {
	OK bool
}

func New() {}
