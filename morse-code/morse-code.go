package morsecode

import (
	"fmt"
	"strings"
)

var toMorse = map[string]rune{
	".-":   'A',
	"-...": 'B',
	"-.-.": 'C',
	"-..":  'D',
	".":    'E',
	"..-.": 'F',
	"--.":  'G',
	"....": 'H',
	"..":   'I',
	".---": 'J',
	"-.-":  'K',
	".-..": 'L',
	"--":   'M',
	"-.":   'N',
	"---":  'O',
	".--.": 'P',
	"--.-": 'Q',
	".-.":  'R',
	"...":  'S',
	"-":    'T',
	"..-":  'U',
	"...-": 'V',
	".--":  'W',
	"-..-": 'X',
	"-.--": 'Y',
	"--..": 'Z',

	".----": '1',
	"..---": '2',
	"...--": '3',
	"....-": '4',
	".....": '5',
	"-....": '6',
	"--...": '7',
	"---..": '8',
	"----.": '9',
	"-----": '0',

	"..--..": '?',
	"-.-.--": '!',
	".-.-.-": '.',
	"--..--": ',',
	"-.-.-.": ';',
	"---...": ':',
	".-.-.":  '+',
	"-....-": '-',
	"-..-.":  '/',
	"-...-":  '=',
	"/":      ' ',

	//custom types i needed to work with terminal.shop
	".--..--": '[',
	"--..--.": ']',
	"-..--.":  '|',
	"..--..-": '{',
	"--..-..": '}',
	"-.....-": '(',
	".-----.": ')',
	".......": '`',
	"......-": '@',
	".....--": '#',
	"....---": '$',
	"...----": '%',
	"..-----": '^',
	".------": '&',
	"-------": '*',
	"------.": '~',
	"-----..": '"',
	"----...": '\'',
	"---....": '\\',
	"--.....": '|',
	"-......": '_',
	".-.-.-.": '<',
	"-.-.-.-": '>',
}

func reverseMap() map[rune]string {
	fromMorse := make(map[rune]string)
	for morse, char := range toMorse {
		fromMorse[char] = morse
	}
	return fromMorse
}

func ReadMorseIntoString(msg string) (string, error) {
	for _, e := range msg {
		if e != ' ' && e != '.' && e != '-' && e != '/' {
			return "", fmt.Errorf("Morse code contains invalid character: %s", string(e))
		}
	}

	parts := strings.Split(msg, " ")
	translated := ""
	for _, e := range parts {
		if e == "" || e == "\t" || e == " " {
			continue
		}
		char, ok := toMorse[e]
		if !ok {
			return "", fmt.Errorf("Invalid morse code character: %s", e)
		}
		translated += string(char)
	}

	return translated, nil
}

func ReadStringIntoMorse(msg string) (string, error) {
	msg = strings.ToUpper(msg)
	converter := reverseMap()

	if len(msg) < 1 {
		return "", nil
	}

	morseCode := ""
	for _, e := range msg {
		code, ok := converter[e]
		if !ok {
			return "", fmt.Errorf("Character cannot be translated to morse code: %s", string(e))
		}
		morseCode += code + " "
	}

	return morseCode[:len(morseCode)-1], nil
}
