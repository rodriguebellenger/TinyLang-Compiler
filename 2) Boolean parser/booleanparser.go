package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type BinaryOp struct {
	Value string
	Left  *BinaryOp
	Right *BinaryOp
}

// NewBinaryOp is a constructor function for BinaryOp
func NewBinaryOp(value string, left *BinaryOp, right *BinaryOp) *BinaryOp {
	return &BinaryOp{Value: value, Left: left, Right: right}
}

func main() {
	// Test calcul
	var calcul string = "1<2<3<4"
	var tokenizedCalcul []string = tokenize(calcul)
	fmt.Println(tokenizedCalcul)

	var parsedExpr *BinaryOp = parse(tokenizedCalcul)
	printTree(parsedExpr, 2)
	//var result int = evaluateCalcul(&parsedExpr)
	//fmt.Println(result)
}

func parse(tokenizedExpr []string) *BinaryOp {
	var parsedExpr *BinaryOp
	if tokenizedExpr[0] == "BoolExpr" {
		var pos int
		parsedExpr = parseBoolean(tokenizedExpr[1:], &pos)
	} else {
		var pos int
		parsedExpr = parseCalcul(tokenizedExpr, &pos)
	}
	return parsedExpr
}

func parseBoolean(tokenizedExpr []string, pos *int) *BinaryOp {
	left := parseCalcul(tokenizedExpr, pos)

	for *pos < len(tokenizedExpr) && inList([]string{">=", "<=", "==", "!=", "<", ">"}, tokenizedExpr[*pos]) {
		op := tokenizedExpr[*pos]
		*pos++
		right := parseCalcul(tokenizedExpr, pos)
		left = NewBinaryOp(op, left, right) // Chain comparisons into a single tree
	}

	return left
}

func parseCalcul(tokenizedCalcul []string, pos *int) *BinaryOp {
	return expr(tokenizedCalcul, pos)
}

func expr(tokenizedCalcul []string, pos *int) *BinaryOp {
	var node *BinaryOp = term(tokenizedCalcul, pos)
	for *pos < len(tokenizedCalcul) && inList([]string{"+", "-"}, tokenizedCalcul[*pos]) {
		var op string = tokenizedCalcul[*pos]
		*pos++
		var right *BinaryOp = term(tokenizedCalcul, pos)
		node = NewBinaryOp(op, node, right)
	}
	return node
}

func term(tokenizedCalcul []string, pos *int) *BinaryOp {
	var node *BinaryOp = factor(tokenizedCalcul, pos)
	for *pos < len(tokenizedCalcul) && inList([]string{"*", "/"}, tokenizedCalcul[*pos]) {
		var op string = tokenizedCalcul[*pos]
		*pos++
		var right *BinaryOp = factor(tokenizedCalcul, pos)
		node = NewBinaryOp(op, node, right)
	}
	return node
}

func factor(tokenizedCalcul []string, pos *int) *BinaryOp {
	var token string = tokenizedCalcul[*pos]
	*pos++
	if token == "(" {
		var node *BinaryOp = expr(tokenizedCalcul, pos)
		if *pos >= len(tokenizedCalcul) || tokenizedCalcul[*pos] != ")" {
			log.Fatal("Mismatched parentheses")
		}
		*pos++
		return node
	} else if isInt(token) {
		return NewBinaryOp(token, nil, nil)
	} else {
		log.Fatal("Wrong token found (expected a number) : " + token)
		return NewBinaryOp("nil", nil, nil)
	}
}

///////////
// UTILS //
///////////

func tokenize(calcul string) []string {
	var tokenizedCalcul []string

	for i := 0; i < len(calcul); i++ {
		if isInt(string(calcul[i])) {
			var number string
			for i < len(calcul) && isInt(string(calcul[i])) {
				number += string(calcul[i])
				i++
			}
			i--
			tokenizedCalcul = append(tokenizedCalcul, number)
		} else if i+1 < len(calcul) && inList([]string{">=", "<=", "==", "!="}, calcul[i:i+2]) {
			tokenizedCalcul = append(tokenizedCalcul, calcul[i:i+2])
			if tokenizedCalcul[0] != "BoolExpr" {
				tokenizedCalcul = append([]string{"BoolExpr"}, tokenizedCalcul...)
			}
			i++
		} else {
			if calcul[i] != ' ' {
				tokenizedCalcul = append(tokenizedCalcul, string(calcul[i]))
				if inList([]string{"<", ">"}, string(calcul[i])) && tokenizedCalcul[0] != "BoolExpr" {
					tokenizedCalcul = append([]string{"BoolExpr"}, tokenizedCalcul...)
				}
			}
		}
	}

	return tokenizedCalcul
}

func evaluateCalcul(parsedExpr *BinaryOp) int {
	if parsedExpr.Right == nil && parsedExpr.Left == nil {
		return strToInt(parsedExpr.Value)
	}
	var left int = evaluateCalcul(parsedExpr.Left)
	var right int = evaluateCalcul(parsedExpr.Right)

	if parsedExpr.Value == "+" {
		return left + right
	} else if parsedExpr.Value == "-" {
		return left - right
	} else if parsedExpr.Value == "*" {
		return left * right
	} else if parsedExpr.Value == "/" {
		if right == 0 {
			log.Fatal("Cannot divide by 0")
		}
		return left / right
	}
	log.Fatal("err in evaluate()")
	return 1
}

func printTree(node *BinaryOp, depth int) {
	if node == nil {
		return
	}
	fmt.Printf("%s%s\n", string(make([]rune, depth*2, depth*2)), node.Value)
	printTree(node.Left, depth+1)
	printTree(node.Right, depth+1)
}

func inList(liste []string, item string) bool {
	for _, element := range liste {
		if element == item {
			return true
		}
	}
	return false
}

func strToInt(x string) int {
	num, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Error in strToInt : " + x)
		return -1
	}
	return num
}

func isInt(x string) bool {
	for _, char := range x {
		if !(strings.Contains("0123456789", string(char))) {
			return false
		}
	}
	return true
}
