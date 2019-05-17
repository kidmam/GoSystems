package main

//https://medium.com/yakka/better-go-error-handling-with-xerrors-1987650e0c79

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

var ReallyBadError = errors.New("this is a really bad error")

func someErrorHappens() error {
	return xerrors.Errorf("uh oh! something terrible happened: %w", ReallyBadError)
}

type ComplexError struct {
	Message string
	Code    int
	frame   xerrors.Frame
}

func (ce ComplexError) FormatError(p xerrors.Printer) error {
	p.Printf("%d %s", ce.Code, ce.Message)
	ce.frame.Format(p)
	return nil
}

func (ce ComplexError) Format(f fmt.State, c rune) {
	xerrors.FormatError(ce, f, c)
}

func (ce ComplexError) Error() string {
	return fmt.Sprint(ce)
}

func someComplexErrorHappens() error {
	complexErr := ComplexError{
		Code:    1234,
		Message: "there was way too much tuna",
		frame:   xerrors.Caller(1), // skip the first frame
	}
	return xerrors.Errorf("uh oh! something terribly complex happened: %w", complexErr)
}

func main() {
	err := someErrorHappens()
	if xerrors.Is(err, ReallyBadError) {
		// deal with the really bad error
	}

	cerr := someComplexErrorHappens()
	var originalErr ComplexError
	if xerrors.As(cerr, &originalErr) {
		// deal with the complex error
		// we can now directly interrogate originalErr.Code
		// and originalErr.Message!
	}
}
