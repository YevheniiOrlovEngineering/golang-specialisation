package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MaxStrSize = 20

type Person struct {
	fname string
	lname string
}

func (p *Person) Set(fname string, lname string) {
	p.fname = fname
	p.lname = lname

	if len(fname) > MaxStrSize {
		p.fname = p.fname[:MaxStrSize]
	}

	if len(lname) > MaxStrSize {
		p.lname = p.lname[:MaxStrSize]
	}
}

func (p *Person) Get() (string, string) {
	return p.fname, p.lname
}

func main() {
	var people []Person

	for _, line := range getPeopleList() {
		fullName := strings.Split(line, " ")
		pp := &Person{}
		pp.Set(fullName[0], fullName[1])
		people = append(people, *pp)
	}
	printPeople(people)
}

func getPeopleList() []string {
	fmt.Println("[INFO] Enter the filename to open: ")

	fn, err := bufio.NewReader(os.Stdin).ReadString('\n')
	isErr(err)
	pplRaw, err := os.ReadFile(strings.TrimSuffix(fn, "\n"))
	isErr(err)

	return splitPeople(string(pplRaw))
}

func splitPeople(pplRaw string) []string {
	return strings.Split(
		strings.TrimSuffix(pplRaw, "\n"), "\n")
}

func printPeople(ppl []Person) {
	for _, p := range ppl {
		pp := &p
		name, lastname := pp.Get()
		fmt.Printf("Name: %s, Last Name: %s\n", name, lastname)
	}
}

func isErr(e error) {
	if e != nil {
		panic(e)
	}
}
