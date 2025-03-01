package main

import (
	"fmt"
	"log"
	"strconv"
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
	var calcul string = "2-(3+5)*3-2"
	fmt.Println(calcul)

	// Test tree
	var root BinaryOp = *NewBinaryOp("+", &BinaryOp{"5", nil, nil}, &BinaryOp{"3", nil, nil})
	fmt.Println(root.Value, root.Left.Value, root.Right.Value)

	var tokenizedCalcul []string = tokenize(calcul)
	fmt.Println(tokenizedCalcul)

	var parsedCalcul BinaryOp = *parse(tokenizedCalcul)
	printTree(&parsedCalcul, 2)

	var result int = evaluate(&parsedCalcul)
	fmt.Println(result)
}

func parse(tokenizedCalcul []string) *BinaryOp {
	var pos int
	return expr(tokenizedCalcul, &pos)
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
			panic("Mismatched parentheses")
		}
		*pos++
		return node
	} else {
		return NewBinaryOp(token, nil, nil)
	}
}

func evaluate(parsedCalcul *BinaryOp) int {
	if parsedCalcul.Right == nil && parsedCalcul.Left == nil {
		return strToInt(parsedCalcul.Value)
	}
	var left int = evaluate(parsedCalcul.Left)
	var right int = evaluate(parsedCalcul.Right)

	if parsedCalcul.Value == "+" {
		return left + right
	} else if parsedCalcul.Value == "-" {
		return left - right
	} else if parsedCalcul.Value == "*" {
		return left * right
	} else if parsedCalcul.Value == "/" {
		return left / right
	}
	log.Fatal("err in evaluate()")
	return 1
}

///////////
// UTILS //
///////////

func printTree(node *BinaryOp, depth int) {
	if node == nil {
		return
	}
	fmt.Printf("%s%s\n", string(make([]rune, depth*2, depth*2)), node.Value)
	printTree(node.Left, depth+1)
	printTree(node.Right, depth+1)
}

func tokenize(calcul string) []string {
	var tokenizedCalcul []string

	for _, char := range calcul {
		tokenizedCalcul = append(tokenizedCalcul, string(char))
	}

	return tokenizedCalcul
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
