package main
import "fmt"

func ValidateParenthesis(s string) bool {
	// {(abc+xyz*[1+2])}()
	str := []rune(s)
	n := len(str)
	stack := make([]rune, n)
	top := -1
	for _, currChar := range str {
		sCurrChar := string(currChar)
		if sCurrChar == "(" || sCurrChar == "{" || sCurrChar == "[" {
			top += 1
			stack[top] = currChar
			
		}
		if sCurrChar == ")" || sCurrChar == "}" || sCurrChar == "]" {
			if len(stack) > 0 && !isParenthesisMatchCaseStyle(string(stack[top]), sCurrChar) {
				return false
			}
			top -= 1
		}
	}
	return true
}

func isParenthesisMatchCaseStyle(p1 string, p2 string) bool {
	switch {
	case p1 == "{" && p2 == "}":
		return true
	case p1 == "(" && p2 == ")":
		return true
	case p1 == "[" && p2 == "]":
		return true
	default:
		return false
	}
}