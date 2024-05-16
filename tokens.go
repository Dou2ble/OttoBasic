package main

type TokenKind uint8

const (
	tokenKindString TokenKind = iota
	tokenKindHash
	tokenKindEqualSign
	tokenKindReturn
	tokenKindVariable
	tokenKindSection
)

type Token struct {
	kind  TokenKind
	value string
}

func newToken(kind TokenKind, value string) Token {
	return Token{kind: kind, value: value}
}
