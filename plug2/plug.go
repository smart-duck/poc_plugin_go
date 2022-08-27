package main

import "fmt"

type Plug struct {
	params string
}

func (plug *Plug) Init(params string) {
	fmt.Printf("Init plugin Plug2 with params %s\n", params)
	plug.params = params
}

func (plug *Plug) Info() string {
	return fmt.Sprintf("Plugin Plug2 having parameters %s", plug.params)
}

// exported as symbol named "VeradcoPlugin"
var VeradcoPlugin Plug