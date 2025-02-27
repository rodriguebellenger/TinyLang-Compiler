class ASTNode:
    pass

class NumberNode(ASTNode):
    def __init__(self, value):
        self.value = value  # e.g., 5

class VariableNode(ASTNode):
    def __init__(self, name):
        self.name = name  # e.g., "x"

class BinaryOpNode(ASTNode):
    def __init__(self, left, operator, right):
        self.left = left   # left operand (VariableNode or NumberNode)
        self.operator = operator  # e.g., "+"
        self.right = right # right operand

class AssignmentNode(ASTNode):
    def __init__(self, name, value):
        self.name = name  # e.g., "x"
        self.value = value  # right-hand side expression

class PrintNode(ASTNode):
    def __init__(self, expr):
        self.expr = expr  # expression to print

class WhileNode(ASTNode):
    def __init__(self, condition, body):
        self.condition = condition  # comparison expression
        self.body = body  # list of statements

class Parser:
    def __init__(self, tokens):
        self.tokens = tokens
        self.pos = 0  # Track current position in tokens

    def peek(self):
        """Look at the current token without consuming it."""
        if self.pos < len(self.tokens):
            return self.tokens[self.pos]
        return None  # End of input

    def consume(self):
        """Consume and return the current token."""
        token = self.peek()
        self.pos += 1
        return token

    def parse_expression(self):
        """Parses expressions like 'x + 5' or '42'."""
        token = self.consume()
        
        if token[0] == "NUMBER":
            return NumberNode(int(token[1]))
        elif token[0] == "IDENT":
            return VariableNode(token[1])
        else:
            raise SyntaxError("Expected a number or variable")

    def parse_assignment(self):
        """Parses assignments like 'x = 5'."""
        var_name = self.consume()[1]  # Get variable name
        self.consume()  # Consume '='
        value = self.parse_expression()
        return AssignmentNode(var_name, value)

    def parse_print(self):
        """Parses 'print x'."""
        self.consume()  # Consume 'print'
        expr = self.parse_expression()
        return PrintNode(expr)

    def parse_while(self):
        """Parses 'while x > 0 ... end'."""
        self.consume()  # Consume 'while'
        left = self.parse_expression()
        op = self.consume()[1]  # Operator (e.g., '>')
        right = self.parse_expression()
        condition = BinaryOpNode(left, op, right)

        body = []
        while self.peek()[0] != "END":
            body.append(self.parse_statement())

        self.consume()  # Consume 'end'
        return WhileNode(condition, body)

    def parse_statement(self):
        """Parses a single statement like assignment, print, or while."""
        token = self.peek()

        if token[0] == "IDENT":
            return self.parse_assignment()
        elif token[0] == "PRINT":
            return self.parse_print()
        elif token[0] == "WHILE":
            return self.parse_while()
        else:
            raise SyntaxError(f"Unexpected token {token}")

    def parse(self):
        """Parses the full program into an AST."""
        statements = []
        while self.pos < len(self.tokens):
            statements.append(self.parse_statement())
        return statements  # List of AST nodes

tokens = [
    ("IDENT", "x"), ("ASSIGN", "="), ("NUMBER", "5"), ("+", "+"), ("NUMBER", "5"), ("+", "+"), ("NUMBER", "5"),
    ("WHILE", "while"), ("IDENT", "x"), (">", ">"), ("NUMBER", "0"),
    ("NEWLINE", "\n"),
    ("INDENT", "    "), ("PRINT", "print"), ("IDENT", "x"), ("NEWLINE", "\n"),
    ("INDENT", "    "), ("IDENT", "x"), ("ASSIGN", "="), ("IDENT", "x"), ("-", "-"), ("NUMBER", "1"),
    ("NEWLINE", "\n"),
    ("END", "end")
]
parser = Parser(tokens)
ast = parser.parse()