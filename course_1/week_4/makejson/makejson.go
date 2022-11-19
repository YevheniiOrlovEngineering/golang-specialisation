package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	name, address := readNameAddrStdIn()
	db := map[string]string{"name": name, "address": address}

	dbJson, err := json.Marshal(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dbJson))
}

func readNameAddrStdIn() (string, string) {
	var buf = bufio.NewReader(os.Stdin)

	fmt.Println("[INFO] Enter the name: ")
	name, err := buf.ReadString('\n')
	isErr(err)

	fmt.Println("[INFO] Enter the address: ")
	address, err := buf.ReadString('\n')
	isErr(err)

	name = strings.TrimSuffix(name, "\n")
	address = strings.TrimSuffix(address, "\n")

	return name, address
}

func isErr(e error) {
	if e != nil {
		panic(e)
	}
}
