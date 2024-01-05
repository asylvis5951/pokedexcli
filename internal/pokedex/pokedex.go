package pokedex

import (
	"fmt"
)

type Pdex struct {
	entries	map[string]pdexEntry
}

type pdexEntry struct {
	value []byte
}

func (p *Pdex) Add(name string, val []byte) error {
	if _, ok := p.Get(name); ok {
		return fmt.Errorf("%v already exists in Pokedex", name)
	}
	p.entries[name] = pdexEntry{
		value: val,
	}
	// fmt.Println(p.entries[name])
	return nil
}

func (p *Pdex) Get(name string) ([]byte, bool) {
	if _, ok := p.entries[name]; !ok {
		return []byte{}, false
	}
	return p.entries[name].value, true
}

func NewPokedex() *Pdex {
	e := map[string]pdexEntry{}
	p := Pdex{
		entries: e,
	}
	return &p
}

func (p *Pdex) GetAll() {
	for k,_ := range p.entries {
		fmt.Printf("- %v\n", k)
	}
}
