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

func (o *IO) ReadConditions() []int {
	o.s.Scan()
	n, _ := strconv.Atoi(o.s.Text())

	nums := make([]int, n)

	for i := range n - 1 {
		o.s.Scan()
		nums[i], _ = strconv.Atoi(o.s.Text())
	}

	return nums
}

func (o *IO) WriteResult(res []int) {
	for _, v := range res {
		fmt.Fprintf(o.w, "%v ", v)
	}
}

func main() {
	o := NewIO()

	nums := o.ReadConditions()

	QuickSort(0, len(nums)-1, nums)

	fmt.Println(nums)

	o.WriteResult(nums)
}

func QuickSort(low int, high int, nums []int) {
	if low >= high {
		return
	}

	pivot := nums[high]

	i, j := low, low

	for j < high {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}

		j++
	}

	nums[i], nums[j] = nums[j], nums[i]

	QuickSort(low, i-1, nums)
	QuickSort(i+1, high, nums)
}
