package maps

import "fmt"

func checkAnk() {
	visitedPages := make(map[string]struct{})

	visitedPages["https://example.com"] = struct{}{}

	visitedPages["https://codesignal.com"] = struct{}{}

	if _, visited := visitedPages["https://codesignal.com"]; visited {
		fmt.Println("The user visited https://codesignal.com before")
	}
}
