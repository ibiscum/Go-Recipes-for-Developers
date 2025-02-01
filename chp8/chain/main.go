package main

import (
	"errors"
	"fmt"
)

type ErrSyntax struct {
	Line int
	Col  int
	Err  error
}

func (err ErrSyntax) Error() string {
	return fmt.Sprintf("syntax error at line %d col %d: %s", err.Line, err.Col, err.Err.Error())
}

func (err ErrSyntax) Is(e error) bool {
	_, ok := e.(ErrSyntax)
	return ok
}

func (err ErrSyntax) As(target any) bool {
	if tgt, ok := target.(*ErrSyntax); ok {
		*tgt = err
		return true
	}
	return false
}

// Simulate an error return, unwrapped
func Parse1() error {
	return ErrSyntax{
		Line: 1,
		Col:  2,
		Err:  errors.New("simulated error"),
	}
}

// Simulate an error return, wrapped once
func Parse2() error {
	return fmt.Errorf("Error wrapped once: %w", Parse1())
}

// Simulate an error return, wrapped twice
func Parse3() error {
	return fmt.Errorf("Error wrapped twice: %w", Parse2())
}

func main() {
	// The base error
	err := Parse1()
	fmt.Printf("Base error is syntax error: %t\n", errors.Is(err, ErrSyntax{}))
	var syntax ErrSyntax
	if errors.As(err, &syntax) {
		fmt.Printf("Syntax error: %+v\n", syntax)
	}
	err = Parse2()
	fmt.Printf("Wrapped once, error type: %T, base error is syntax error: %t\n", err, errors.Is(err, ErrSyntax{}))
	if errors.As(err, &syntax) {
		fmt.Printf("Syntax error: %+v\n", syntax)
	}
	err = Parse3()
	fmt.Printf("Wrapped twice, error type: %T, base error is syntax error: %t\n", err, errors.Is(err, ErrSyntax{}))
	if errors.As(err, &syntax) {
		fmt.Printf("Syntax error: %+v\n", syntax)
	}
}
