package main

import (
	"fmt"
)

// Alphabet constants
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Rotor struct
type Rotor struct {
	wiring      string
	notch       byte
	position    int
	ringSetting int
}

// Plugboard struct
type Plugboard struct {
	mapping map[byte]byte
}

// Reflector wiring (e.g., B reflector)
const reflector = "GFHDJSKQLMPAZOIERUYTBVNCXW"

// Create a new rotor
func NewRotor(wiring string, notch byte) Rotor {
	return Rotor{
		wiring:      wiring,
		notch:       notch,
		position:    0,
		ringSetting: 0,
	}
}

// Rotate the rotor
func (r Rotor) Rotate() bool {
	r.position = (r.position + 1) % 26
	return r.wiring[r.position] == r.notch
}

// Encode a character through the rotor
func (r *Rotor) Encode(c byte, reverse bool) byte {
	offset := r.position - r.ringSetting
	if offset < 0 {
		offset += 26
	}

	index := (int(c-'A') + offset) % 26
	if reverse {
		index = (26 + index - offset) % 26
	}

	return r.wiring[index]
}

// Plugboard setup
func NewPlugboard(pairs map[byte]byte) Plugboard {
	return Plugboard{mapping: pairs}
}

// Encode a character through the plugboard
func (p *Plugboard) Encode(c byte) byte {
	if mapped, ok := p.mapping[c]; ok {
		return mapped
	}
	return c
}

// Main function
func main() {
	// Example rotors and plugboard
	rotor1 := NewRotor("AQWZSXEDCRFVBGTYHN?JUIKOLPM", 'Q')
	rotor2 := NewRotor("MPLOIK?JUYHNBGTRFVCDEZSXWQA", 'L')
	rotor3 := NewRotor("WQSXCDFVGBNHJ?KLMPOUIYTERZA", 'T')

	plugboard := NewPlugboard(map[byte]byte{
		'A': 'K',
		'B': 'Z',
		'C': 'W',
		'D': 'N',
		'a': 'l',
		'k': 'c',
		'h': 'i',
	})

	// Input message
	message := "HELLO WORD"
	encoded := ""

	for _, char := range message {
		if char < 'A' || char > 'Z' {
			encoded += string(char)
			continue
		}

		// Plugboard pass
		c := plugboard.Encode(byte(char))

		// Rotor forward
		c = rotor1.Encode(c, false)
		c = rotor2.Encode(c, false)
		c = rotor3.Encode(c, false)

		// Reflector
		c = reflector[c-'A']

		// Rotor reverse
		c = rotor3.Encode(c, true)
		c = rotor2.Encode(c, true)
		c = rotor1.Encode(c, true)

		// Plugboard pass
		c = plugboard.Encode(c)

		encoded += string(c)

		// Rotate rotors
		if rotor1.Rotate() {
			if rotor2.Rotate() {
				rotor3.Rotate()
			}
		}
	}

	fmt.Println("Encoded message:", encoded)
}
