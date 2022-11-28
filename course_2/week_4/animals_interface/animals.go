package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

// ---

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c *Cow) Eat() {
	fmt.Printf("Food: %s\n", c.food)
}

func (c *Cow) Move() {
	fmt.Printf("Locomotion: %s\n", c.locomotion)
}

func (c *Cow) Speak() {
	fmt.Printf("Noise: %s!\n", c.noise)
}

// ---

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b *Bird) Eat() {
	fmt.Printf("Food: %s\n", b.food)
}

func (b *Bird) Move() {
	fmt.Printf("Locomotion: %s\n", b.locomotion)
}

func (b *Bird) Speak() {
	fmt.Printf("Noise: %s!\n", b.noise)
}

// ---

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s *Snake) Eat() {
	fmt.Printf("Food: %s\n", s.food)
}

func (s *Snake) Move() {
	fmt.Printf("Locomotion: %s\n", s.locomotion)
}

func (s *Snake) Speak() {
	fmt.Printf("Noise: %s!\n", s.noise)
}
