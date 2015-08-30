package main

import "io"

//Translator is an interface for
//producing output for specific uses
//rather than just printing it to the
//screen as commands do. Essentially
//allows for overriding output of a
//command.
type Translator interface {
	Initialize(w io.Writer, s Storage)
	Execute(w io.Writer, s Storage, list string, key string, value string)
	Finalize(w io.Writer, s Storage)
}
