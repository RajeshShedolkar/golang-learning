package main

//import "fmt"


func ValidateParentheses(s string) bool {
	stack := make([]rune, 0)

	match := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		// opening brackets → push
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
			continue
		}

		// closing brackets → validate
		if ch == ')' || ch == '}' || ch == ']' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top != match[ch] {
				return false
			}
		}
	}

	// stack must be empty
	return len(stack) == 0
}


func ValidateParentheses2(s string) bool {
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
			if top > -1 && !isParenthesisMatchHashStyle(string(stack[top]), sCurrChar) {
				return false
			}
			top -= 1
		}
	}
	if top < -1{
		return false
	}
	return true
}

func isParenthesisMatchHashStyle(p1 string, p2 string) bool {
	hash := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	r1 := rune(p1[0])
	r2 := rune(p2[0])
	if v, ok := hash[r1]; ok {
		if v == r2 {
			return true
		}
	}
	return false
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
