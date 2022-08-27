package main

import "fmt"

type Plug1 struct {
	params string
}

func (plug *Plug1) Init(params string) {
	fmt.Printf("Init plugin Plug1 with params %s\n", params)
	plug.params = params
}

func (plug *Plug1) Info() string {
	return fmt.Sprintf("Plugin Plug1 having parameters %s", plug.params)
}

// exported as symbol named "VeradcoPlugin"
var VeradcoPlugin Plug1