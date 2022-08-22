package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func findIAN() (string, error) {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("[INFO] Enter the string: ")
	line, err := in.ReadString('\n')

	if err != nil {
		return "", errors.New("[ERROR] 'Scan' executed with error")
	}

	outString := strings.ToLower(strings.TrimSuffix(line, "\n"))

	if outString[0] == 'i' && outString[len(outString)-1] == 'n' && strings.ContainsRune(outString, 'a') {
		fmt.Println("[INFO] Found!")
	} else {
		fmt.Println("[INFO] Not Found!")
	}
	return outString, nil
}

func main() {
	if _, err := findIAN(); err != nil {
		fmt.Println(err)
	}
}
