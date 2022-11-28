package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func validateInput(s string, requests []string, animals []string, commands []string, zoo map[string]Animal) (
	bool, string, string, string,
) {
	m, err := regexp.MatchString("^[a-z]+ [a-zA-Z]+ [a-z]+$", s)
	isErr(err)
	if !m {
		fmt.Println("[ERROR] Wrong format ('command arg1 arg2'). " +
			"Enter 3 strings, space separated in a single line.")
		return false, "", "", ""
	}
	ws := strings.Split(s, " ")
	r, arg1, arg2 := ws[0], ws[1], ws[2]

	if !contains(r, requests) {
		fmt.Printf("[ERROR] Bad input. Request should be one from the list: (%s).\n",
			strings.Join(requests, " "))
		return false, "", "", ""
	}

	switch r {
	case requests[0]:
		name, animal := arg1, arg2
		if zoo[name] != nil {
			fmt.Printf("[ERROR] Bad input. An animal with name '%s' already exists.\n", name)
			return false, "", "", ""
		}
		if !contains(animal, animals) {
			fmt.Printf("[ERROR] Bad input. Animal should be one from the list: (%s).\n",
				strings.Join(animals, " "))
			return false, "", "", ""
		}
		return true, r, name, animal
	case requests[1]:
		name, command := arg1, arg2
		if zoo[name] == nil {
			fmt.Printf("[ERROR] Bad input. An animal with name '%s' doesn't exist.\n", name)
			return false, "", "", ""
		}
		if !contains(command, commands) {
			fmt.Printf("[ERROR] Bad input. Command should be one from the list: (%s).\n",
				strings.Join(commands, " "))
			return false, "", "", ""
		}
		return true, r, name, command
	default:
		break
	}
	return false, "", "", ""
}

func readInput() string {
	fmt.Print("> ")
	inp, err := bufio.NewReader(os.Stdin).ReadString('\n')
	isErr(err)
	return strings.TrimSpace(inp)
}

func contains(s string, ss []string) bool {
	tmp := make([]string, len(ss))
	copy(tmp, ss)
	sort.Strings(tmp)
	i := sort.SearchStrings(tmp, s)
	if i == len(tmp) || tmp[i] != s {
		return false
	}
	return true
}

func isErr(e error) {
	if e != nil {
		panic(e)
	}
}

func addAnimal(n string, a string, zoo map[string]Animal, aList map[string]Animal) {
	zoo[n] = aList[a]
	fmt.Println("Created it!")
}

func queryAnimal(a Animal, c string, commands map[string]func(a Animal)) {
	commands[c](a)
}
