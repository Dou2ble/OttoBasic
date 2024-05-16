package main

import "unicode"

func lex(source string) []Token {
	result := []Token{}

	for i := 0; i < len(source); i++ {
		if source[i] == '#' {
			result = append(result, newToken(tokenKindHash, "#"))
		} else if source[i] == '=' {
			result = append(result, newToken(tokenKindEqualSign, "="))
		} else if source[i] == '<' {
			result = append(result, newToken(tokenKindReturn, "<"))
		} else if source[i] == '"' {
			// strings
			value := ""
			i++
			for ; source[i] != '"'; i++ {
				value = value + string(source[i])
			}
			result = append(result, newToken(tokenKindString, value))
		} else if !unicode.IsSpace(rune(source[i])) {
			// sections and variables
			value := ""
			for ; !unicode.IsSpace(rune(source[i])); i++ {
				value = value + string(source[i])
			}

			var kind TokenKind
			if unicode.IsUpper(rune(value[0])) {
				kind = tokenKindSection
			} else {
				kind = tokenKindVariable
			}

			result = append(result, newToken(kind, value))
		}
	}

	return result
}
