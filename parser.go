package main

import (
	"log/slog"
)

func parseStringExpression(tokens []Token, i *int) StringExpression {
	if tokens[*i].kind != tokenKindString {
		slog.Error("Did not find string token in string expression")
	}

	result := StringExpression{value: tokens[*i].value}
	*i++
	return result
}

func parseExpression(tokens []Token, i *int) Expression {
	if tokens[*i].kind == tokenKindString {
		return parseStringExpression(tokens, i)
	}

	slog.Error("Failed to parse expression at token: ", i)
	return nil
}

func parseAssignment(tokens []Token, i *int) AssignmentStatement {
	result := AssignmentStatement{}

	if tokens[*i].kind != tokenKindVariable {
		slog.Error("Did not find variable name in assignment")
	}
	result.assignee = tokens[*i].value
	*i++

	if tokens[*i].kind != tokenKindEqualSign {
		slog.Error("Did not find equal sign in assignment")
	}
	*i++

	// parse the expression that will be the new value of the assignee
	result.expression = parseExpression(tokens, i)

	return result
}

func parseStatement(tokens []Token, i *int) Statement {
	if tokens[*i+1].kind == tokenKindEqualSign {
		return parseAssignment(tokens, i)
	}

	slog.Error("Failed to parse statement at token: ", i)
	*i++
	return nil
}

func parse(source string) Program {
	tokens := lex(source)
	ast := Program{}
	i := 0

	for i < len(tokens)-1 {
		statement := parseStatement(tokens, &i)
		if statement != nil {
			ast.content = append(ast.content, statement)
		}
	}

	return ast
}
