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

func (scanner *Scanner) BreakLine(content string) {
	if scanner.index < len(content) {
		character := rune(content[scanner.index])

		if character == '\n' {
			scanner.line++
			scanner.index++
		}
	}
}

func (scanner *Scanner) Skip(content string) {
	if scanner.index < len(content) {
		character := rune(content[scanner.index])
		whitespaces := map[rune]string{
			' ':  "SPACE",
			'\t': "TAB",
		}

		if _, exists := whitespaces[character]; exists {
			scanner.index++
		}
	}
}

func (scanner *Scanner) Comment(content string) {
	if scanner.index < len(content) {
		comments := map[string]string{
			"//": "COMMENT",
		}

		if (scanner.index + 1) < len(content) {
			characters := content[scanner.index : scanner.index+2]

			if _, exists := comments[characters]; exists {

				for scanner.index < len(content) && content[scanner.index] != '\n' {
					scanner.index++
				}

				scanner.line++
				scanner.index++
			}
		}
	}
}

func (scanner *Scanner) TokenizeNumber(content string) {
	if scanner.index < len(content) {
		character := rune(content[scanner.index])

		if unicode.IsDigit(character) {
			start := scanner.index

			for scanner.index < len(content) && unicode.IsDigit(rune(content[scanner.index])) {
				scanner.index++
			}

			if scanner.index < len(content) && content[scanner.index] == '.' {
				scanner.index++
				for scanner.index < len(content) && unicode.IsDigit(rune(content[scanner.index])) {
					scanner.index++
				}
			}

			number := content[start:scanner.index]
			float, _ := strconv.ParseFloat(number, 64)

			if math.Mod(float, 1.0) == 0 {
				number = strconv.FormatFloat(float, 'f', -1, 64) + ".0"
			} else {
				number = strconv.FormatFloat(float, 'f', -1, 64)
			}

			token := fmt.Sprintf("NUMBER %s %s", content[start:scanner.index], number)
			fmt.Println(token)
			scanner.tokens = append(scanner.tokens, token)
		}
	}

}

func (scanner *Scanner) TokenizeIdentifierOrKeyword(content string) {
	keywords := map[string]string{
		"and":    "AND",
		"class":  "CLASS",
		"else":   "ELSE",
		"false":  "FALSE",
		"for":    "FOR",
		"fun":    "FUN",
		"if":     "IF",
		"nil":    "NIL",
		"or":     "OR",
		"print":  "PRINT",
		"return": "RETURN",
		"super":  "SUPER",
		"this":   "THIS",
		"true":   "TRUE",
		"var":    "VAR",
		"while":  "WHILE",
	}

	if scanner.index < len(content) {
		character := rune(content[scanner.index])

		if unicode.IsLetter(character) || character == '_' {
			start := scanner.index

			for scanner.index < len(content) && (unicode.IsLetter(rune(content[scanner.index])) || unicode.IsDigit(rune(content[scanner.index])) || rune(content[scanner.index]) == '_') {
				scanner.index++
			}

			word := content[start:scanner.index]

			if name, exists := keywords[word]; exists {
				token := fmt.Sprintf("%s %s null", name, word)
				fmt.Println(token)
				scanner.tokens = append(scanner.tokens, token)
			} else {
				// If not a keyword, treat it as an identifier
				token := fmt.Sprintf("IDENTIFIER %s null", word)
				fmt.Println(token)
				scanner.tokens = append(scanner.tokens, token)
			}
		}
	}

}

func (scanner *Scanner) TokenizeDouble(content string) {
	if scanner.index < len(content) {
		doubles := map[string]string{
			"==": "EQUAL_EQUAL",
			"!=": "BANG_EQUAL",
			">=": "GREATER_EQUAL",
			"<=": "LESS_EQUAL",
		}

		if (scanner.index + 1) < len(content) {
			characters := content[scanner.index : scanner.index+2]

			if name, exists := doubles[characters]; exists {
				token := fmt.Sprintf("%s %s null", name, characters)
				fmt.Println(token)
				scanner.tokens = append(scanner.tokens, token)
				scanner.index += 2
				// Recursive call to handle multiple double characters following each other.
				scanner.TokenizeDouble(content)
			}
		}
	}
}

func (scanner *Scanner) TokenizeString(content string) {
	if scanner.index < len(content) {
		character := rune(content[scanner.index])
		wrappers := map[rune]string{
			'"': "STRING",
		}

		if name, exists := wrappers[character]; exists {
			start := scanner.index
			end := scanner.index + 1

			for end < len(content) && rune(content[end]) != character {
				end++
			}

			if end == len(content) {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("[line %d] Error: Unterminated string.", scanner.line))
				scanner.exception = true
			} else {
				token := fmt.Sprintf("%s %s %s", name, content[start:end+1], content[start+1:end])
				fmt.Println(token)
				scanner.tokens = append(scanner.tokens, token)
			}

			scanner.index = end + 1
		}
	}
}

func (scanner *Scanner) TokenizeSingle(content string) {
	if scanner.index < len(content) {
		character := rune(content[scanner.index])
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
			scanner.index++
		}
	}
}

func (scanner *Scanner) HandleUnexpected(content string) {
	if scanner.index < len(content) {
		character := rune(content[scanner.index])
		errors := map[rune]string{
			'#': "HASH",
			'$': "DOLLAR",
			'%': "PERCENT",
			'@': "AT",
		}

		if _, exists := errors[character]; exists {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("[line %d] Error: Unexpected character: %c", scanner.line, character))
			scanner.exception = true
			scanner.index++
		}
	}
}

func (scanner *Scanner) Scan(content string) {

	for scanner.index < len(content) {
		scanner.BreakLine(content)
		scanner.Skip(content)
		scanner.Comment(content)
		scanner.TokenizeNumber(content)
		scanner.TokenizeIdentifierOrKeyword(content)
		scanner.TokenizeDouble(content)
		scanner.TokenizeString(content)
		scanner.TokenizeSingle(content)
		scanner.HandleUnexpected(content)
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
