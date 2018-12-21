package main

import "fmt"

type Tesla struct {
	ID    string `json:"ID"`
	Make  string `json:"Make"`
	Model string `json:"Model"`
	Cost  int64  `json:"Cost"`
}

func main() {
	fmt.Println("Initializing project")
}
