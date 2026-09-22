package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, nums := readInput()

	res := MinDifference(n, nums)

	writeOutput(res)
}

func readInput() (int, []int) {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	nums := make([]int, n)

	for i := 0; i <= n-1; i++ {
		if s.Err() != nil {
			s.Scan()
			nums[i], _ = strconv.Atoi(s.Text())
		}
	}

	return n, nums
}

func writeOutput(res []int) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, v := range res {
		fmt.Fprintf(w, "%v ", v)
	}
}

func MinDifference(n int, nums []int) []int {
	res := make([]int, 2)
	res[0], res[1] = nums[0], nums[1]

	min := abs(nums[0], nums[1])

	for i, j := 0, 1; j < n; j++ {
		eq := abs(nums[i], nums[j])

		if eq < min {
			res[0], res[1] = nums[i], nums[j]
			break
		}

		i++
	}

	return res
}

func abs(i, j int) int {
	if i >= j {
		return i - j
	} else {
		return j - i
	}
}
