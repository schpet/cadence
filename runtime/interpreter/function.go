package interpreter

import (
	"github.com/dapperlabs/flow-go/language/runtime/ast"
	"github.com/dapperlabs/flow-go/language/runtime/errors"
	"github.com/dapperlabs/flow-go/language/runtime/sema"
	"github.com/raviqqe/hamt"
	// revive:disable:dot-imports
	. "github.com/dapperlabs/flow-go/language/runtime/trampoline"
	// revive:enable
)

// Invocation

type Invocation struct {
	Arguments     []Value
	ArgumentTypes []sema.Type
	Location      LocationPosition
	Interpreter   *Interpreter
}

// FunctionValue

type FunctionValue interface {
	Value
	isFunctionValue()
	invoke(Invocation) Trampoline
}

// InterpretedFunctionValue

type InterpretedFunctionValue struct {
	Interpreter      *Interpreter
	ParameterList    *ast.ParameterList
	Type             *sema.FunctionType
	Activation       hamt.Map
	BeforeStatements []ast.Statement
	PreConditions    ast.Conditions
	Statements       []ast.Statement
	PostConditions   ast.Conditions
}

func (InterpretedFunctionValue) isValue() {}

func (f InterpretedFunctionValue) Copy() Value {
	return f
}

func (InterpretedFunctionValue) GetOwner() string {
	// value is never owned
	return ""
}

func (InterpretedFunctionValue) SetOwner(owner string) {
	// NO-OP: value cannot be owned
}

func (InterpretedFunctionValue) isFunctionValue() {}

func (f InterpretedFunctionValue) invoke(invocation Invocation) Trampoline {
	return f.Interpreter.invokeInterpretedFunction(f, invocation.Arguments)
}

// HostFunctionValue

type HostFunction func(invocation Invocation) Trampoline

type HostFunctionValue struct {
	Function HostFunction
	Members  map[string]Value
}

func NewHostFunctionValue(
	function HostFunction,
) HostFunctionValue {
	return HostFunctionValue{
		Function: function,
	}
}

func (HostFunctionValue) isValue() {}

func (f HostFunctionValue) Copy() Value {
	return f
}

func (HostFunctionValue) GetOwner() string {
	// value is never owned
	return ""
}

func (HostFunctionValue) SetOwner(owner string) {
	// NO-OP: value cannot be owned
}

func (HostFunctionValue) isFunctionValue() {}

func (f HostFunctionValue) invoke(invocation Invocation) Trampoline {
	return f.Function(invocation)
}

func (f HostFunctionValue) GetMember(interpreter *Interpreter, _ LocationRange, name string) Value {
	return f.Members[name]
}

func (f HostFunctionValue) SetMember(_ *Interpreter, _ LocationRange, _ string, _ Value) {
	panic(errors.NewUnreachableError())
}
