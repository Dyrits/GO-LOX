package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

type Scanner struct {
	line      int
	index     int
	exception bool
	tokens    []string
}

func (scanner *Scanner) PrintTokens() {
	for _, token := range scanner.tokens {
		fmt.Println(token)
	}
	fmt.Println("EOF  null")
}

func (scanner *Scanner) TokenizeSingle(character rune) {
	singles := map[rune]string{
		'(': "LEFT_PAREN",
		')': "RIGHT_PAREN",
		'-': "MINUS",
		'*': "STAR",
		'+': "PLUS",
		',': "COMMA",
		'.': "DOT",
		';': "SEMICOLON",
		'{': "LEFT_BRACE",
		'}': "RIGHT_BRACE",
		'=': "EQUAL",
		'!': "BANG",
		'>': "GREATER",
		'<': "LESS",
		'/': "SLASH",
	}

	if name, exists := singles[character]; exists {
		token := fmt.Sprintf("%s %c null", name, character)
		fmt.Println(token)
		scanner.tokens = append(scanner.tokens, token)
	}
}

func (scanner *Scanner) HandleUnexpected(character rune) {
	errors := map[rune]string{
		'#': "HASH",
		'$': "DOLLAR",
		'%': "PERCENT",
		'@': "AT",
	}

	if _, exists := errors[character]; exists {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("[line %d] Error: Unexpected character: %c", scanner.line, character))
		scanner.exception = true
	}
}

func (scanner *Scanner) Scan(content string) {

	doubles := map[string]string{
		"==": "EQUAL_EQUAL",
		"!=": "BANG_EQUAL",
		">=": "GREATER_EQUAL",
		"<=": "LESS_EQUAL",
	}

	comments := map[string]string{
		"//": "COMMENT",
	}

	whitespaces := map[rune]string{
		' ':  "SPACE",
		'\t': "TAB",
	}

	wrappers := map[rune]string{
		'"': "STRING",
	}

	for index := 0; index < len(content); index++ {
		character := rune(content[index])

		// Check for new lines.
		if character == '\n' {
			scanner.line++
			continue
		}
		// Skip whitespaces.
		if _, exists := whitespaces[character]; exists {
			continue
		}
		// Check for numbers.
		if unicode.IsDigit(character) {
			start := index
			for index < len(content) && unicode.IsDigit(rune(content[index])) {
				index++
			}
			if index < len(content) && content[index] == '.' {
				index++
				for index < len(content) && unicode.IsDigit(rune(content[index])) {
					index++
				}
			}
			number := content[start:index]
			float, _ := strconv.ParseFloat(number, 64)
			if math.Mod(float, 1.0) == 0 {
				number = strconv.FormatFloat(float, 'f', -1, 64) + ".0"
			} else {
				number = strconv.FormatFloat(float, 'f', -1, 64)
			}
			token := fmt.Sprintf("NUMBER %s %s", content[start:index], number)
			fmt.Println(token)
			scanner.tokens = append(scanner.tokens, token)
			index--
			continue
		}
		// Check for identifiers.
		if unicode.IsLetter(character) || character == '_' {
			start := index
			for index < len(content) && (unicode.IsLetter(rune(content[index])) || unicode.IsDigit(rune(content[index])) || rune(content[index]) == '_') {
				index++
			}
			identifier := content[start:index]
			token := fmt.Sprintf("IDENTIFIER %s null", identifier)
			fmt.Println(token)
			scanner.tokens = append(scanner.tokens, token)
			index--
			continue
		}
		if (index + 1) < len(content) {
			characters := content[index : index+2]
			// Check for comments.
			if _, exists := comments[characters]; exists {
				// Skip the rest of the line.
				for index < len(content) && content[index] != '\n' {
					index++
				}
				scanner.line++
				continue
			}
			// Check for double characters.
			if name, exists := doubles[characters]; exists {
				token := fmt.Sprintf("%s %s null", name, characters)
				fmt.Println(token)
				scanner.tokens = append(scanner.tokens, token)
				index++
				continue
			}
		}
		// Check for wrappers.
		if name, exists := wrappers[character]; exists {
			// Get the rest of the string.
			start := index
			end := index + 1
			for end < len(content) && rune(content[end]) != character {
				end++
			}
			if end == len(content) {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("[line %d] Error: Unterminated string.", scanner.line))
				scanner.exception = true
				break
			}
			token := fmt.Sprintf("%s %s %s", name, content[start:end+1], content[start+1:end])
			fmt.Println(token)
			scanner.tokens = append(scanner.tokens, token)
			index = end
		}
		// Check for single characters.
		scanner.TokenizeSingle(character)
		// Check for errors.
		scanner.HandleUnexpected(character)
	}
	fmt.Println("EOF  null")
	if scanner.exception {
		os.Exit(65)
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(content) > 0 {
		scanner := &Scanner{
			line:      1,
			index:     0,
			exception: false,
			tokens:    []string{},
		}
		scanner.Scan(string(content))
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
