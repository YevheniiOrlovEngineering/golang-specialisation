package main

func main() {
	zoo := make(map[string]Animal)
	var commands []string
	var animals []string

	animalsMap := map[string]Animal{
		"cow":   &Cow{food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  &Bird{food: "worms", locomotion: "fly", noise: "peep"},
		"snake": &Snake{food: "mice", locomotion: "slither", noise: "hsss"}}

	commandsMap := map[string]func(a Animal){
		"eat": Animal.Eat, "move": Animal.Move, "speak": Animal.Speak}

	requests := []string{"newanimal", "query"}

	for c := range commandsMap {
		commands = append(commands, c)
	}
	for a := range animalsMap {
		animals = append(animals, a)
	}

	for {
		inp := readInput()
		ok, r, arg1, arg2 := validateInput(inp, requests, animals, commands, zoo)
		if !ok {
			continue
		}
		switch r {
		case requests[0]:
			addAnimal(arg1, arg2, zoo, animalsMap)
			break
		case requests[1]:
			queryAnimal(zoo[arg1], arg2, commandsMap)
			break
		default:
			break
		}
	}
}
