package main

// interfaces
type Statement interface {
	statement()
}

type Expression interface {
	expression()
}

// the Program itself
type Program struct {
	content []Statement
}

// STATEMENTS
type AssignmentStatement struct {
	assignee   string
	expression Expression
}

func (statement AssignmentStatement) statement() {}

// EXPRESSIONS
type StringExpression struct {
	value string
}

func (expression StringExpression) expression() {}

type IdentifierExpression struct {
	value string
}

func (expression IdentifierExpression) expression() {}
