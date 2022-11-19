package main

import (
	"fmt"
)

func truncNum() (int, error) {
	var inNum float64

	fmt.Println("[INFO] Enter float number: ")
	if _, err := fmt.Scan(&inNum); err != nil {
		return -1, err
	}
	outNum := int(inNum)
	fmt.Println("[INFO] Truncated number:", outNum)
	return outNum, nil
}

func main() {
	if _, err := truncNum(); err != nil {
		fmt.Println(err)
	}
}
