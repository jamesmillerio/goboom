package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestAllCommand(t *testing.T) {

	expectedList := "tests"
	expectedEntryKey := "test1:"
	expectedEntryValue := "Hello world 1\n"
	buffer := new(bytes.Buffer)
	storage := NewCustomListStorage("./.boom")
	all := NewAllCommand()
	all.Execute(buffer, storage)
	actual := buffer.String()

	//Make sure we have a definition and description.
	testDescriptionAndDefinitions(all, t)

	if !strings.Contains(actual, expectedList) {
		t.Errorf("The all command did not return an expected list. Expected: '%v' Actual: '%v'\n", expectedList, actual)
	}

	if !strings.Contains(actual, expectedEntryKey) {
		t.Errorf("The all command did not return an expected entry key. Expected: '%v' Actual: '%v'\n", expectedEntryKey, actual)
	}

	if !strings.Contains(actual, expectedEntryValue) {
		t.Errorf("The all command did not return an expected entry value. Expected: '%v' Actual: '%v'\n", expectedEntryValue, actual)
	}
}

func TestAllListCommand(t *testing.T) {

	list := "tests"
	expectedEntryKey := "test1:"
	expectedEntryValue := "Hello world 1\n"
	buffer := new(bytes.Buffer)
	storage := NewCustomListStorage("./.boom")
	all := NewAllListCommand(list)
	all.Execute(buffer, storage)
	actual := buffer.String()

	//Make sure we have a definition and description.
	testDescriptionAndDefinitions(all, t)

	if !strings.Contains(actual, expectedEntryKey) {
		t.Errorf("The all command did not return an expected entry key. Expected: '%v' Actual: '%v'\n", expectedEntryKey, actual)
	}

	if !strings.Contains(actual, expectedEntryValue) {
		t.Errorf("The all command did not return an expected entry value. Expected: '%v' Actual: '%v'\n", expectedEntryValue, actual)
	}
}

func TestEchoCommand(t *testing.T) {

	expected := "Hello world 1\n"
	buffer := new(bytes.Buffer)
	storage := NewCustomListStorage("./.boom")
	echo := NewEchoEntryCommand("test1")
	echo.Execute(buffer, storage)
	actual := buffer.String()

	//Make sure we have a definition and description.
	testDescriptionAndDefinitions(echo, t)

	if actual != expected {
		t.Errorf("The echo command did not return the correct value. Expected: '%v' Actual: '%v'\n", expected, actual)
	}
}

func TestEchoSpecificEntryCommand(t *testing.T) {

	expected := "Hello world 1\n"
	buffer := new(bytes.Buffer)
	storage := NewCustomListStorage("./.boom")
	echo := NewEchoSpecificEntryCommand("tests", "test1")
	echo.Execute(buffer, storage)
	actual := buffer.String()

	//Make sure we have a definition and description.
	testDescriptionAndDefinitions(echo, t)

	if actual != expected {
		t.Errorf("The echo command did not return the correct value. Expected: '%v' Actual: '%v'\n", expected, actual)
	}
}

func TestEchoCommandNotFound(t *testing.T) {

	list := "test413"
	entry := "test9123"
	expected := fmt.Sprintf("%v not found in %v.\n", entry, list)
	buffer := new(bytes.Buffer)
	storage := NewCustomListStorage("./.boom")
	echo := NewEchoSpecificEntryCommand(list, entry)
	echo.Execute(buffer, storage)
	actual := buffer.String()

	//Make sure we have a definition and description.
	testDescriptionAndDefinitions(echo, t)

	if actual != expected {
		t.Errorf("The echo command did not return the correct value. Expected: '%v' Actual: '%v'\n", expected, actual)
	}
}

/*Common tests and helpers. */
func testDescriptionAndDefinitions(command Command, t *testing.T) {

	if command == nil {
		t.Error("Could not check description or definition. The command was nil.")
	}

	if command.GetDefinition() == "" {
		t.Error("Command has no definition.")
	}

	if command.GetDescription() == "" {
		t.Error("Command has no description.")
	}

}
