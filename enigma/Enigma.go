package main

import (
	"fmt"
)

// Alphabet constants
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Rotor struct
type Rotor struct {
	wiring      string // Represente la cablage interne dur rotor ordre de connaixion de A a Z
	notch       byte   // indique la postion ou le carctère qu'atteinte, provoque le passage dur rotor suivant
	position    int    // la position actuelle du rotor
	ringSetting int    // le réglage du disque de la bague qui permet d'ajuster le décalge fixe
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
	offset := r.position - r.ringSetting // decalage offset
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
	rotor1 := NewRotor("AQWZSXEDCRFVBGTYHNJUIKOLPM", 'Q')
	rotor2 := NewRotor("MPLOIKJUYHNBGTRFVCDEZSXWQA", 'L')
	rotor3 := NewRotor("WQSXCDFVGBNHJKLMPOUIYTERZA", 'T')

	plugboard := NewPlugboard(map[byte]byte{
		'A': 'K', // ces letres seront trasformer par des lettre correspondant avant d'envoyer dans le rotor
		'B': 'Z',
		'C': 'W',
		'D': 'N',
	})

	// Input message
	message := "SALUT TOUT LE MONDE "
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
