package main

import (
	"fmt"
	"io"
)

//NewAlfredScriptFilterTranslator creates
//an instance of an Alfred script filtertransltor.
func NewAlfredScriptFilterTranslator() *AlfredScriptFilterTranslator {
	return new(AlfredScriptFilterTranslator)
}

//AlfredScriptFilterTranslator is a translator that
//allows for the output of lists/keys/value in the
//expected xml format for an Alfred Script Filter.
type AlfredScriptFilterTranslator struct {
}

//Initialize allows for a chance to do any
//initial actions to the io stream.
func (t *AlfredScriptFilterTranslator) Initialize(w io.Writer, s Storage) {
	fmt.Fprint(w, "<?xml version=\"1.0\"?>\n")
	fmt.Fprint(w, "<items>\n")
}

//Execute performs the requested translation.
func (t *AlfredScriptFilterTranslator) Execute(w io.Writer, s Storage, list string, key string, value string) {
	fmt.Fprintf(w, "<item uid=\"%s:%s:%s\" arg=\"%s %s\">\n", list, key, value, list, key)
	fmt.Fprintf(w, "<title>%s</title>\n", key)
	fmt.Fprintf(w, "<subtitle>%s</subtitle>\n", value)
	fmt.Fprintf(w, "<icon type=\"fileicon\">alfredhat.png</icon>\n")
	fmt.Fprint(w, "</item>\n")
}

//Finalize performs an final actions to the io stream.
func (t *AlfredScriptFilterTranslator) Finalize(w io.Writer, s Storage) {
	fmt.Fprint(w, "</items>\n")
}
