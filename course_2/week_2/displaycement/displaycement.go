package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	msgs := []string{
		"[INFO] Enter the acceleration",
		"[INFO] Enter the initial velocity",
		"[INFO] Enter the initial displacement",
		"[INFO] Enter the time",
	}

	a, v0, s0, t := getAllPars(msgs)
	s := genDisplaceFn(a, v0, s0)
	fmt.Printf("Displaycement = %f\n", s(t))
	fmt.Printf("Displaycement = %f\n", s(t+5))

}

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func getAllPars(msgs []string) (float64, float64, float64, float64) {
	var s []float64
	for _, m := range msgs {
		fmt.Println(m)
		s = append(s, getPar())
	}
	return s[0], s[1], s[2], s[3]
}

func getPar() float64 {
	for {
		inp := readStdIn()
		num, err := strToFloat(inp)
		if err != nil {
			fmt.Println("[ERROR] Type error. Enter float number")
			continue
		}
		return num
	}
}

func readStdIn() string {
	inp, err := bufio.NewReader(os.Stdin).ReadString('\n')
	isErr(err)
	return strings.TrimSuffix(inp, "\n")
}

func strToFloat(inp string) (float64, error) {
	return strconv.ParseFloat(inp, 64)
}

func isErr(e error) {
	if e != nil {
		panic(e)
	}
}
