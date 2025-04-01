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
		if e != ' ' && e != '.' && e != '-' {
			return "", fmt.Errorf("Morse code contains invalid character: %q, can only contain '.','-', or ' '", e)
		}
	}

	translated := ""
	command := strings.Split(msg, " ")
	for _, e := range command {
		char, ok := toMorse[e]
		if !ok {
			return "", fmt.Errorf("Invalid morse code character: %s", e)
		}
		translated += string(char)
	}

	return translated, nil
}

func ReadStringIntoMorse(msg string) (string, error) {
	return "", nil
}
