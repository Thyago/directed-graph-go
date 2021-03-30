package main

import "github.com/thyago/directed-graph-go/controllers"

func main() {
	// Create server
	s := controllers.NewServer()
	s.Run()
}
