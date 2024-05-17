package ast

import (
	"bytes"

	"github.com/tredstart/laml/def/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
    Type() string
}

type Object struct {
	Statements []Statement
}

func (p *Object) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Object) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Object
}

type FieldStatement struct {
	Token token.Token
	Value *Identifier
}

type DefStatement struct {
	Token token.Token
	Value *Identifier
}

func (vs *VarStatement) Type() string {return "Var"}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }
func (vs *VarStatement) String() string {
	var out bytes.Buffer
	out.WriteString(vs.TokenLiteral() + " ")
	out.WriteString(vs.Name.Value + ": ")
	out.WriteString(vs.Value.String() + ";")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (fs *FieldStatement) Type() string {return "Field"}
func (fs *FieldStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *FieldStatement) String() string {
	var out bytes.Buffer
	out.WriteString(fs.TokenLiteral() + " = ")
	out.WriteString(fs.Value.String() + ",\n")
	return out.String()
}

func (ds *DefStatement) Type() string { return "Def"}
func (ds *DefStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DefStatement) String() string {
	var out bytes.Buffer
    out.WriteString(ds.TokenLiteral() + " ")
	out.WriteString(ds.Value.String() + ",")
	return out.String()
}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
