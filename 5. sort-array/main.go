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

func (o *IO) ReadConditions() (int, []int) {
	o.s.Scan()
	n, _ := strconv.Atoi(o.s.Text())

	nums := make([]int, n)

	for i := range n {
		o.s.Scan()
		nums[i], _ = strconv.Atoi(o.s.Text())
	}

	return n, nums
}

func (o *IO) WriteResult(res []int) {
	for _, v := range res {
		fmt.Fprintf(o.w, "%v ", v)
	}
}

func main() {
	o := NewIO()

	n, nums := o.ReadConditions()

	res := QuickSort(n, nums)

	o.WriteResult(res)
}

func QuickSort(n int, nums []int) []int {
	return nil
}
