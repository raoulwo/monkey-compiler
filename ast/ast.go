package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (program *Program) String() string {
	var out bytes.Buffer

	for _, statement := range program.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (letStatement *LetStatement) statementNode()       {}
func (letStatement *LetStatement) TokenLiteral() string { return letStatement.Token.Literal }
func (letStatement *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(letStatement.TokenLiteral() + " ")
	out.WriteString(letStatement.Name.String())
	out.WriteString(" = ")

	if letStatement.Value != nil {
		out.WriteString(letStatement.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (identifier *Identifier) expressionNode()      {}
func (identifier *Identifier) TokenLiteral() string { return identifier.Token.Literal }
func (identifier *Identifier) String() string       { return identifier.Value }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (returnStatement *ReturnStatement) statementNode()       {}
func (returnStatement *ReturnStatement) TokenLiteral() string { return returnStatement.Token.Literal }
func (returnStatement *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(returnStatement.TokenLiteral() + " ")

	if returnStatement.ReturnValue != nil {
		out.WriteString(returnStatement.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (expressionStatement *ExpressionStatement) statementNode() {}
func (expressionStatement *ExpressionStatement) TokenLiteral() string {
	return expressionStatement.Token.Literal
}
func (expressionStatement *ExpressionStatement) String() string {
	if expressionStatement.Expression != nil {
		return expressionStatement.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (integerLiteral *IntegerLiteral) expressionNode()      {}
func (integerLiteral *IntegerLiteral) TokenLiteral() string { return integerLiteral.Token.Literal }
func (integerLiteral *IntegerLiteral) String() string       { return integerLiteral.Token.Literal }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (prefixExpression *PrefixExpression) expressionNode() {}
func (prefixExpression *PrefixExpression) TokenLiteral() string {
	return prefixExpression.Token.Literal
}
func (prefixExpression *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(prefixExpression.Operator)
	out.WriteString(prefixExpression.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (infixExpression *InfixExpression) expressionNode()      {}
func (infixExpression *InfixExpression) TokenLiteral() string { return infixExpression.Token.Literal }
func (infixExpression *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(infixExpression.Left.String())
	out.WriteString(" " + infixExpression.Operator + " ")
	out.WriteString(infixExpression.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (boolean *Boolean) expressionNode()      {}
func (boolean *Boolean) TokenLiteral() string { return boolean.Token.Literal }
func (boolean *Boolean) String() string       { return boolean.Token.Literal }

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ifExpression *IfExpression) expressionNode()      {}
func (ifExpression *IfExpression) TokenLiteral() string { return ifExpression.Token.Literal }
func (ifExpression *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ifExpression.Condition.String())
	out.WriteString(" ")
	out.WriteString(ifExpression.Consequence.String())

	if ifExpression.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ifExpression.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (blockStatement *BlockStatement) statementNode()       {}
func (blockStatement *BlockStatement) TokenLiteral() string { return blockStatement.Token.Literal }
func (blockStatement *BlockStatement) String() string {
	var out bytes.Buffer

	for _, statement := range blockStatement.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (functionLiteral *FunctionLiteral) expressionNode()      {}
func (functionLiteral *FunctionLiteral) TokenLiteral() string { return functionLiteral.Token.Literal }
func (functionLiteral *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, parameter := range functionLiteral.Parameters {
		params = append(params, parameter.String())
	}

	out.WriteString(functionLiteral.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(functionLiteral.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (callExpression *CallExpression) expressionNode()      {}
func (callExpression *CallExpression) TokenLiteral() string { return callExpression.Token.Literal }
func (callExpression *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, arg := range callExpression.Arguments {
		args = append(args, arg.String())
	}

	out.WriteString(callExpression.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (stringLiteral *StringLiteral) expressionNode()      {}
func (stringLiteral *StringLiteral) TokenLiteral() string { return stringLiteral.Token.Literal }
func (stringLiteral *StringLiteral) String() string       { return stringLiteral.Token.Literal }

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (arrayLiteral *ArrayLiteral) expressionNode()      {}
func (arrayLiteral *ArrayLiteral) TokenLiteral() string { return arrayLiteral.Token.Literal }
func (arrayLiteral *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, element := range arrayLiteral.Elements {
		elements = append(elements, element.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (indexExpression *IndexExpression) expressionNode()      {}
func (indexExpression *IndexExpression) TokenLiteral() string { return indexExpression.Token.Literal }
func (indexExpression *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(indexExpression.Left.String())
	out.WriteString("[")
	out.WriteString(indexExpression.Index.String())
	out.WriteString("])")

	return out.String()
}

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (hashLiteral *HashLiteral) expressionNode()      {}
func (hashLiteral *HashLiteral) TokenLiteral() string { return hashLiteral.Token.Literal }
func (hashLiteral *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hashLiteral.Pairs {
		pairs = append(pairs, key.String()+": "+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
