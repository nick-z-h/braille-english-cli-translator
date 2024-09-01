package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var ErrMissingArguments = errors.New("missing required cli arguments")
var ErrArgumentsNotBraille = errors.New("the supplied arguments do not match Braille format")
var ErrArgumentsNotAlphanumeric = errors.New("the supplied arguments are not Alphanumeric")
var ErrArgumentsNotBrailleOrAlphanumeric = errors.New("the supplied arguments are not Alphanumeric or Braille")
var ErrTokenNotBraille = errors.New("invalid braille alphabet token")

type Translator struct {
	text               string
	letterToBrailleMap map[string]string
	numberToBrailleMap map[string]string
}

// processes arguments and joins them together
func handleParams() (string, error) {
	if len(os.Args) < 2 {
		return "", ErrMissingArguments
	}
	args := os.Args[1:]
	return strings.TrimSpace(strings.Join(args, " ")), nil
}

// creates new translator unit
func NewTranslator(text string) *Translator {
	var t Translator
	t.text = text
	t.letterToBrailleMap = map[string]string{
		"CAPITAL_FOLLOWS": ".....O",
		"NUMBER_FOLLOWS":  ".O.OOO",
		"a":               "O.....",
		"b":               "O.O...",
		"c":               "OO....",
		"d":               "OO.O..",
		"e":               "O..O..",
		"f":               "OOO...",
		"g":               "OOOO..",
		"h":               "O.OO..",
		"i":               ".OO...",
		"j":               ".OOO..",
		"k":               "O...O.",
		"l":               "O.O.O.",
		"m":               "OO..O.",
		"n":               "OO.OO.",
		"o":               "O..OO.",
		"p":               "OOO.O.",
		"q":               "OOOOO.",
		"r":               "O.OOO.",
		"s":               ".OO.O.",
		"t":               ".OOOO.",
		"u":               "O...OO",
		"v":               "O.O.OO",
		"w":               ".OOO.O",
		"x":               "OO..OO",
		"y":               "OO.OOO",
		"z":               "O..OOO",
		" ":               "......",
	}
	t.numberToBrailleMap = map[string]string{
		"0": ".OOO..",
		"1": "O.....",
		"2": "O.O...",
		"3": "OO....",
		"4": "OO.O..",
		"5": "O..O..",
		"6": "OOO...",
		"7": "OOOO..",
		"8": "O.OO..",
		"9": ".OO...",
	}
	return &t
}

// checks if the provided arguments are Braille
// not necessarily VALID Braille
func (t *Translator) isBraille() bool {
	for _, rune := range t.text {
		if rune != 'O' && rune != '.' {
			return false
		}
	}
	return true
}

// checks if the provided arguments are Alphanumeric
func (t *Translator) isAlphanumeric() bool {
	for _, rune := range t.text {
		if !unicode.IsLetter(rune) && !unicode.IsNumber(rune) {
			return false
		}
	}
	return true
}

func (t *Translator) toBraille() (string, error) {
	if !t.isBraille() {
		return "", fmt.Errorf("expected Braille input: %s", ErrArgumentsNotBraille)
	}
	// convert to braille
	// token not braille
	return "braille", nil
}

func (t *Translator) toEnglish() (string, error) {
	if !t.isAlphanumeric() {
		return "", fmt.Errorf("expected Alphanumeric input: %s", ErrArgumentsNotAlphanumeric)
	}
	// convert to english
	return "english", nil
}

func (t *Translator) Translate() (string, error) {
	if t.isBraille() {
		return t.toBraille()
	} else if t.isAlphanumeric() {
		return t.toEnglish()
	}
	return "", fmt.Errorf("expected Alphanumeric or Braille input: %s", ErrArgumentsNotBrailleOrAlphanumeric)
}

func main() {
	text, err := handleParams()

	if err != nil {
		log.Fatalf("invalid args: %s\n", err)
	}

	translator := NewTranslator(text)

	convertedText, err := translator.Translate()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(convertedText)
}
