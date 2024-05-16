package main

import (
	"fmt"
	"io/ioutil"
	"log/slog"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("source.ottobasic")
	if err != nil {
		slog.Error("Error reading file:", err)
		os.Exit(1)
		return
	}

	// Convert the byte slice to a string
	source := string(data)

	tokens := lex(source)
	for _, token := range tokens {
		slog.Info(fmt.Sprint(token))
	}

	ast := parse(source)
	slog.Info(fmt.Sprint(ast))
}
