package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"person_name"`
	Title   string `json:"person_title"`
	Address `json:"person_address_obj"`
}

type Address struct {
	Street string `json:"person_addr_street"`
	City   string `json:"person_city"`
	State  string `json:"person_state"`
	Postal string `json:"person_postal_code"`
}

func updateName(p *Person, name string) {
	p.Name = name
}

func main() {
	p := Person{}
	p.Name = "uknown"
	p.Title = "author"
	p.Street = "12345 Main street"
	p.City, p.State = "Goville", "Go"
	p.Postal = "12345"
	fmt.Println(p)
	updateName(&p, "Vladimir Vivien")
	fmt.Println(p)
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}
