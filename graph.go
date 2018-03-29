package main

import (
	"fmt"

	"github.com/knighthawkbro/graphs/linkedgraph"
)

// Set (Public) - I think there is a keyword or package called set.
type Set interface {
	Add(vert, item string)
	Contains(item string) bool
	Size() int
	String() string

	GetNeighbors(vert string) string
	RemoveVert(vert string) string
	ContainsNeighbor(vert, neigh string) bool
}

func main() {
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as an linked graph...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	graph := linkedgraph.New()
	driver(graph)
}

func driver(cities Set) {
	bostonTowns := []string{
		"Somerville", "Natick", "Concord", "Brookline",
	}
	for _, town := range bostonTowns {
		cities.Add("Boston", town)
	}
	northOfBostonTowns := []string{
		"Beverly", "Peabody",
	}
	for _, town := range northOfBostonTowns {
		cities.Add("Andover", town)
	}
	southOfBostonTowns := []string{
		"Norfolk", "Fall River", "Foxboro",
	}
	for _, town := range southOfBostonTowns {
		cities.Add("Bridgewater", town)
	}

	fmt.Println("Here are the cities in the graph and some towns",
		"in their region")
	fmt.Println(cities)

	fmt.Print("\nIs Boston in the graph?")
	if cities.Contains("Boston") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Print("Is Worcester in the graph?")
	if cities.Contains("Worcester") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Println("\nHere are the cities near Boston:", cities.GetNeighbors("Boston"))
	fmt.Println("Here are the cities near Andover:", cities.GetNeighbors("Andover"))

	fmt.Print("\nIs Natick a neighbor of Boston?")
	if cities.ContainsNeighbor("Boston", "Natick") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Print("Is Foxboro a neighbor of Boston?")
	if cities.ContainsNeighbor("Boston", "Foxboro") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Println("\nRemoving Andover from the graph")
	fmt.Println("Neighbors retrned:", cities.RemoveVert("Andover"))
	fmt.Println(cities)
	fmt.Println("\nRemoving Boston from the graph")
	fmt.Println("Neighbors retrned:", cities.RemoveVert("Boston"))
	fmt.Println(cities)
	fmt.Println("\nRemoving Bridgewater from the graph")
	fmt.Println("Neighbors retrned:", cities.RemoveVert("Bridgewater"))
	fmt.Println(cities)
	fmt.Println("\nTrying to remove Boston again?")
	fmt.Println("Neighbors retrned:", cities.RemoveVert("Boston"))
	fmt.Println(cities)
}
