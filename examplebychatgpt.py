import re

class Node:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right

    def __repr__(self):
        return self.to_parentheses()
    
    def to_parentheses(self):
        if self.left is None and self.right is None:
            return str(self.value)
        left_str = self.left.to_parentheses() if self.left else ''
        right_str = self.right.to_parentheses() if self.right else ''
        return f'({left_str} {self.value} {right_str})'

def tokenize(expression):
    tokens = re.findall(r'\d+\.\d+|\d+|[+\-*/()]', expression)
    return tokens

class Parser:
    def __init__(self, tokens):
        self.tokens = tokens
        self.pos = 0

    def parse(self):
        return self.expr()
    
    def expr(self):
        node = self.term()
        while self.pos < len(self.tokens) and self.tokens[self.pos] in ('+', '-'):
            op = self.tokens[self.pos]
            self.pos += 1
            node = Node(op, node, self.term())
        return node
    
    def term(self):
        node = self.factor()
        while self.pos < len(self.tokens) and self.tokens[self.pos] in ('*', '/'):
            op = self.tokens[self.pos]
            self.pos += 1
            node = Node(op, node, self.factor())
        return node
    
    def factor(self):
        token = self.tokens[self.pos]
        self.pos += 1
        if token == '(':
            node = self.expr()
            self.pos += 1  # Skip closing parenthesis ')'
            return node
        else:
            return Node(float(token) if '.' in token else int(token))

def parse_expression(expression):
    tokens = tokenize(expression.replace(' ', ''))
    parser = Parser(tokens)
    return parser.parse()

# Example usage:
expr = "3 + 5 * ( 2 - 8 ) / 2"
print(parse_expression(expr))
