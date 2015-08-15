package main

import "io"

//Command defines the common functions
//a Command struct should have to facilitate
//store/retrieving/acting on key/values.
type Command interface {
	GetDefinition() string
	GetDescription() string
	Execute(w io.Writer, s Storage)
}
