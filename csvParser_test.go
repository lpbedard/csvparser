package parser

import (
	"os"
	"testing"
)

type ExampleContact1 struct {
	FirstName   string `csv:"0"`
	LastName    string `csv:"1"`
	Title       string `csv:"2"`
	Email       string `csv:"3"`
	Birthdate   string `csv:"4"`
	Description string `csv:"5"`
}

type ExampleContact2 struct {
	FirstName   string
	LastName    string `csv:"1"`
	Title       string `csv:"2"`
	Email       string
	Birthdate   string `csv:"4"`
	Description string `csv:"5"`
}

type ExampleContact3 struct {
	FirstName   string
	LastName    string
	Title       string
	Email       string
	Birthdate   string
	Description string
}

var contacts1 []interface{}
var contacts2 []interface{}
var contacts3 []interface{}

var parseErr1 error
var parseErr2 error
var parseErr3 error
var csvParser CsvParser

func TestMain(m *testing.M) {
	//setup

	csvParser = CsvParser{
		CsvFile:      "example.csv",
		CsvSeparator: ',',
	}

	contacts1, parseErr1 = csvParser.Parse(ExampleContact1{})
	contacts2, parseErr2 = csvParser.Parse(ExampleContact2{})
	contacts3, parseErr3 = csvParser.Parse(ExampleContact3{})

	//run all the tests
	os.Exit(m.Run())
}

func TestParsingHasNoError(t *testing.T) {
	if parseErr1 != nil {
		t.Errorf("Parsing1 returns an error: %v", parseErr1)
	}
	if parseErr2 != nil {
		t.Errorf("Parsing2 returns an error: %v", parseErr2)
	}
	if parseErr3 != nil {
		t.Errorf("Parsing3 returns an error: %v", parseErr3)
	}
}

func TestAllContactsHaveBeenParsed(t *testing.T) {
	if len(contacts1) != 2 {
		t.Errorf("Parsed contacts length is wrong: Actual %v, Expected %v", len(contacts1), 2)
	}
	if len(contacts2) != 2 {
		t.Errorf("Parsed contacts length is wrong: Actual %v, Expected %v", len(contacts2), 2)
	}
	if len(contacts3) != 2 {
		t.Errorf("Parsed contacts length is wrong: Actual %v, Expected %v", len(contacts3), 2)
	}
}

func TestContactHaveBeenParsed1(t *testing.T) {

	var c = contacts1[0].(*ExampleContact1)

	if c.FirstName != "Tom" {
		t.Errorf("Parsed contact firstname is wrong: Actual %v, Expected %v", c.FirstName, "Tom")
	}
	if c.LastName != "Jones" {
		t.Errorf("Parsed contact lastname is wrong: Actual %v, Expected %v", c.LastName, "Jones")
	}
	if c.Title != "Senior Director" {
		t.Errorf("Parsed contact title is wrong: Actual %v, Expected %v", c.Title, "Senior Director")
	}
	if c.Email != "buyer@mymail.com" {
		t.Errorf("Parsed contact email is wrong: Actual %v, Expected %v", c.Email, "buyer@mymail.com")
	}
	if c.Birthdate != "1999-06-07" {
		t.Errorf("Parsed contact birthdate is wrong: Actual %v, Expected %v", c.Birthdate, "1999-06-07")
	}
	if c.Description != "Self-described as \"the top\" branding guru on the West Coast" {
		t.Errorf("Parsed contact description is wrong: Actual %v, Expected %v", c.Description, "Self-described as \"the top\" branding guru on the West Coast")
	}
}

func TestContactHaveBeenParsed2(t *testing.T) {

	var c = contacts2[0].(*ExampleContact2)

	if c.FirstName != "Tom" {
		t.Errorf("Parsed contact firstname is wrong: Actual %v, Expected %v", c.FirstName, "Tom")
	}
	if c.LastName != "Jones" {
		t.Errorf("Parsed contact lastname is wrong: Actual %v, Expected %v", c.LastName, "Jones")
	}
	if c.Title != "Senior Director" {
		t.Errorf("Parsed contact title is wrong: Actual %v, Expected %v", c.Title, "Senior Director")
	}
	if c.Email != "buyer@mymail.com" {
		t.Errorf("Parsed contact email is wrong: Actual %v, Expected %v", c.Email, "buyer@mymail.com")
	}
	if c.Birthdate != "1999-06-07" {
		t.Errorf("Parsed contact birthdate is wrong: Actual %v, Expected %v", c.Birthdate, "1999-06-07")
	}
	if c.Description != "Self-described as \"the top\" branding guru on the West Coast" {
		t.Errorf("Parsed contact description is wrong: Actual %v, Expected %v", c.Description, "Self-described as \"the top\" branding guru on the West Coast")
	}
}

func TestContactHaveBeenParsed3(t *testing.T) {

	var c = contacts3[0].(*ExampleContact3)

	if c.FirstName != "Tom" {
		t.Errorf("Parsed contact firstname is wrong: Actual %v, Expected %v", c.FirstName, "Tom")
	}
	if c.LastName != "Jones" {
		t.Errorf("Parsed contact lastname is wrong: Actual %v, Expected %v", c.LastName, "Jones")
	}
	if c.Title != "Senior Director" {
		t.Errorf("Parsed contact title is wrong: Actual %v, Expected %v", c.Title, "Senior Director")
	}
	if c.Email != "buyer@mymail.com" {
		t.Errorf("Parsed contact email is wrong: Actual %v, Expected %v", c.Email, "buyer@mymail.com")
	}
	if c.Birthdate != "1999-06-07" {
		t.Errorf("Parsed contact birthdate is wrong: Actual %v, Expected %v", c.Birthdate, "1999-06-07")
	}
	if c.Description != "Self-described as \"the top\" branding guru on the West Coast" {
		t.Errorf("Parsed contact description is wrong: Actual %v, Expected %v", c.Description, "Self-described as \"the top\" branding guru on the West Coast")
	}
}

func TestParsingANotExistingCsvFile(t *testing.T) {
	var csvNotExistingParser = CsvParser{
		CsvFile:      "example_not_existing.csv",
		CsvSeparator: ',',
	}

	_, err := csvNotExistingParser.Parse(ExampleContact1{})

	if err == nil {
		t.Error("Parsing a not existing csv file should return an error")
	}
}

func TestParsingAnInvalidCsvFile(t *testing.T) {
	var csvNotExistingParser = CsvParser{
		CsvFile:      "example_invalid.csv",
		CsvSeparator: ',',
	}

	_, err := csvNotExistingParser.Parse(ExampleContact1{})

	if err == nil {
		t.Error("TestParsingAnInvalidCsvFile should return an error")
	}
}
