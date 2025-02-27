package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

////////////
// TOKENS //
////////////

type Token struct {
	value     string
	tokenType string
	line      int
}

const (
	integer         string = "Integer"
	identifier      string = "Identifier"
	reservedKeyword string = "Keyword"
	operator        string = "Operator"
	comparison      string = "Comparison"
)

var keywords []string = []string{"if", "while", "print", "func", "end"}
var comparisons []string = []string{"==", ">=", "<=", "!="}

func main() {
	args := os.Args[1:] // Skip the program name
	content, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal("\rCouldn't read file")
	}

	var startTime time.Time = time.Now()
	var tokenizedProgram []Token = tokenize(string(content))
	var elapsed time.Duration = time.Since(startTime)
	fmt.Println(tokenizedProgram)
	fmt.Printf("Temps : %s\n", elapsed)
}

///////////////
// TOKENIZER //
///////////////

func tokenize(program string) []Token {
	var tokenizedProgram []Token
	var line int = 1
	i := 0
	for i < len(program) {
		ch := program[i]

		if unicode.IsSpace(rune(ch)) {
			if ch == '\n' {
				line += 1
			}
			i += 1
			continue
		}
		if len(program)-i > 1 {
			if inList(comparisons, program[i:i+2]) {
				tokenizedProgram = append(tokenizedProgram, Token{value: program[i : i+2], tokenType: comparison, line: line})
				i += 2
				continue
			}
		}
		switch ch {
		case '+', '-', '=':
			tokenizedProgram = append(tokenizedProgram, Token{value: string(program[i]), tokenType: operator, line: line})
		case '>', '<':
			tokenizedProgram = append(tokenizedProgram, Token{value: string(program[i]), tokenType: comparison, line: line})
		default:
			var word string
			for i < len(program) && (unicode.IsLetter(rune(program[i])) || unicode.IsDigit(rune(program[i]))) {
				word += string(program[i])
				i++
			}
			if isInt(word) {
				tokenizedProgram = append(tokenizedProgram, Token{value: word, tokenType: integer, line: line})
			} else if inList(keywords, word) {
				tokenizedProgram = append(tokenizedProgram, Token{value: word, tokenType: reservedKeyword, line: line})
			} else if isValidIdentifier(word) {
				tokenizedProgram = append(tokenizedProgram, Token{value: word, tokenType: identifier, line: line})
			} else {
				var err string = "Unrecognized word \"" + word + "\" at line " + intToStr(line) + program[i-2:i+3]
				fmt.Println(tokenizedProgram)
				log.Fatal(err)
			}
			continue
		}
		i++
	}
	return tokenizedProgram
}

///////////
// UTILS //
///////////

func inList(liste []string, item string) bool {
	for _, element := range liste {
		if element == item {
			return true
		}
	}
	return false
}

func isValidIdentifier(x string) bool {
	if len(x) == 0 {
		return false
	}
	for _, c := range x {
		if !strings.Contains("azertyuiopqsdfghjklmwxcvbnAZERTYUIOPQSDFGHJKLMWXCVBN", string(c)) {
			return false
		}
	}
	return true
}

func isInt(x string) bool {
	if len(x) == 0 {
		return false
	}
	for _, c := range x {
		if !strings.Contains("0123456789", string(c)) {
			return false
		}
	}
	return true
}

func intToStr(x int) string {
	num := strconv.Itoa(x)
	return num
}

func strToInt(x string) int {
	num, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Error in strToInt : " + x)
		return -1
	}
	return num
}
