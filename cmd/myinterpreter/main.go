package main

import (
	"fmt"
	"os"
)

type Scanner struct{}

func (scanner *Scanner) Scan(content string) {
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

	errors := map[rune]string{
		'#': "HASH",
		'$': "DOLLAR",
		'%': "PERCENT",
		'@': "AT",
	}

	invalid := false

	for index := 0; index < len(content); index++ {
		// Skip whitespaces.
		if _, exists := whitespaces[rune(content[index])]; exists {
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
				continue
			}
			// Check for double characters.
			if name, exists := doubles[characters]; exists {
				fmt.Println(fmt.Sprintf("%s %s null", name, characters))
				index++
				continue
			}
		}
		character := rune(content[index])
		// Check for single characters.
		if name, exists := singles[character]; exists {
			fmt.Println(fmt.Sprintf("%s %c null", name, character))
		}
		// Check for errors.
		if _, exists := errors[character]; exists {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("[line 1] Error: Unexpected character: %c", character))
			invalid = true
		}
	}
	fmt.Println("EOF  null")
	if invalid {
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
		scanner := &Scanner{}
		scanner.Scan(string(content))
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
