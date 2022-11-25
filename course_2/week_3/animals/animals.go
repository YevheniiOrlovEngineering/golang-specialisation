package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) eat() {
	fmt.Printf("Food: %s\n", a.food)
}

func (a Animal) move() {
	fmt.Printf("Locomotion: %s\n", a.locomotion)
}

func (a Animal) speak() {
	fmt.Printf("Noise: %s!\n", a.noise)
}

func main() {
	var animals []string
	var commands []string

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	asMap := map[string]map[string]func(){
		"cow":   {"eat": cow.eat, "move": cow.move, "speak": cow.speak},
		"bird":  {"eat": bird.eat, "move": bird.move, "speak": bird.speak},
		"snake": {"eat": snake.eat, "move": snake.move, "speak": snake.speak},
	}

	for k := range asMap {
		animals = append(animals, k)
	}

	for c := range asMap[animals[0]] {
		commands = append(commands, c)
	}

	for {
		inp := readInput()
		ok, animal, command := validateInput(inp, animals, commands)
		if !ok {
			continue
		}
		asMap[animal][command]()
	}
}

func readInput() string {
	fmt.Print("> ")
	inp, err := bufio.NewReader(os.Stdin).ReadString('\n')
	isErr(err)
	return strings.TrimSpace(inp)
}

func validateInput(s string, animals []string, commands []string) (bool, string, string) {
	m, err := regexp.MatchString("^[a-z]+ [a-z]+$", s)
	isErr(err)
	if !m {
		fmt.Println("[ERROR] Wrong format ('animal command'). Enter 2 strings, space separated in a single line.")
		return false, "", ""
	}
	ws := strings.Split(s, " ")
	a, c := ws[0], ws[1]

	if !contains(a, animals) {
		fmt.Printf("[ERROR] Bad input. Animal should be one from the list: (%s).\n",
			strings.Join(animals, " "))
		return false, "", ""
	}
	if !contains(c, commands) {
		fmt.Printf("[ERROR] Bad input. Command should be one from the list: (%s).\n",
			strings.Join(commands, " "))
		return false, "", ""
	}
	return true, a, c
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
