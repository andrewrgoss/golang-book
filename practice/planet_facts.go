/*
 *	This program returns facts about the eight planets in our solar system. 
 *
 *	Author: Andrew Goss, andrewrgoss.com
 *
 * 	Initial: December 10, 2015
 */

package main

import (
	"fmt"
	"log"
	"io"
)

func main() {
	
	// Set up logging...

	var logFile io.Writer
	log.New(logFile, "Log: ", 4)

	// Credits

	fmt.Println("\n\n***************** Planet Facts *****************")
	fmt.Println("\n\nAuthor: Andrew Goss\nWebsite: andrewrgoss.com\n\n")
	
	p := new(Planet)
	p.Name = "Venus"
	p.OrderFromSun = 2
	p.SpewFacts()
	
	e := new(Earth)
	e.Name = "Earth"
	e.OrderFromSun = 3
	e.SpewFacts()
}

type Planet struct {
	Name string
	OrderFromSun uint
	Composition string
	Rings bool
}

type Earth struct {
	Planet // Earth "is-a"" planet
	Moons uint
}

func (p *Planet) rings()bool {
	if p.Name == "Saturn" {
		return true
	}
	return false
}

func (p *Planet) SpewFacts() {
	log.Print(p.Name, " is planet #",p.OrderFromSun, " from the sun.")
	log.Print(p.Name, " has", )
}