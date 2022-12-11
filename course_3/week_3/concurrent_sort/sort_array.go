package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortSlice(s []int, c chan []int) {
	fmt.Println("Slice to be sorted:\t", s)
	if !sort.IntsAreSorted(s) {
		sort.Ints(s)
	} else {
		fmt.Println("Slice is already sorted.")
	}
	c <- s
}

func splitSliceToFour(inS []int) (s1, s2, s3, s4 []int) {
	const PARTS = 4.0

	if r, size := len(inS)%PARTS, len(inS)/PARTS; r == 0 {
		return inS[:size], inS[size : size*2], inS[size*2 : size*3], inS[size*3 : size*4]
	} else {
		var ss [][]int
		j := 0
		for i := 0; i < PARTS-r; i++ {
			ss = append(ss, inS[j:j+size])
			j += size
		}
		for i := 0; i < r; i++ {
			ss = append(ss, inS[j:j+size+1])
			j += size + 1
		}
		return ss[0], ss[1], ss[2], ss[3]
	}
}

// TODO: rewrite to merge sort concept
func mergeRearrangeSlices(ss [][]int, res []int) {
	i := 0
	for _, s := range ss {
		for _, e := range s {
			res[i] = e
			i++
		}
	}
	sort.Ints(res)
}

func readStdIn() string {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("[INFO] Enter the integer array:\n> ")
	inp, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(inp)
}

func fillSlice() []int {
	for {
		var s []int
		isOk := true
		inp := readStdIn()
		for _, e := range strings.Fields(inp) {
			num, err := strconv.Atoi(e)
			if err != nil {
				fmt.Printf("[ERROR] Wrong type. Enter the integer value.\n")
				isOk = false
				break
			}
			s = append(s, num)
		}
		if isOk {
			return s
		}
	}
}

func main() {
	c := make(chan []int, 4)

	rawS := fillSlice()
	sortedS := make([]int, len(rawS))

	s1, s2, s3, s4 := splitSliceToFour(rawS)
	go sortSlice(s1, c)
	go sortSlice(s2, c)
	go sortSlice(s3, c)
	go sortSlice(s4, c)
	s1, s2, s3, s4 = <-c, <-c, <-c, <-c

	mergeRearrangeSlices([][]int{s1, s2, s3, s4}, sortedS)
	fmt.Println("Sorted slice:\t\t", sortedS)
}
