package parser

import (
	"os"
	"testing"
)

var contacts1 []interface{}
var contacts2 []interface{}
var contacts3 []interface{}

var parseErr1 error
var parseErr2 error
var parseErr3 error
var csvParser CsvParser

func TestMain(m *testing.M) {
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
	testSingleContact(t, contacts1[0].(*ExampleContact1))
}

func TestContactHaveBeenParsed2(t *testing.T) {
	testSingleContact(t, contacts2[0].(*ExampleContact2))
}

func TestContactHaveBeenParsed3(t *testing.T) {
	testSingleContact(t, contacts3[0].(*ExampleContact3))
}

func TestContactWithInvalidBoolField(t *testing.T) {
	_, err := csvParser.Parse(ExampleContactInvalidBoolean{})

	if err == nil {
		t.Error("TestContactWithInvalidBoolField should return an error")
	}
}

func TestContactWithInvalidUintField(t *testing.T) {
	_, err := csvParser.Parse(ExampleContactInvalidUint{})

	if err == nil {
		t.Error("TestContactWithInvalidUintField should return an error")
	}
}

func TestContactWithInvalidIntField(t *testing.T) {
	_, err := csvParser.Parse(ExampleContactInvalidInt{})

	if err == nil {
		t.Error("TestContactWithInvalidIntField should return an error")
	}
}

func TestContactWithInvalidFloat32Field(t *testing.T) {
	_, err := csvParser.Parse(ExampleContactInvalidFloat32{})

	if err == nil {
		t.Error("TestContactWithInvalidFloat32Field should return an error")
	}
}

func TestContactWithInvalidFloat64Field(t *testing.T) {
	_, err := csvParser.Parse(ExampleContactInvalidFloat64{})

	if err == nil {
		t.Error("TestContactWithInvalidFloat64Field should return an error")
	}
}

func TestContactWithCsvColumnTooHigh(t *testing.T) {
	_, err := csvParser.Parse(ExampleContactWithCsvColumnFieldTooHigh{})

	if err == nil {
		t.Error("TestContactWithCsvColumnTooHigh should return an error")
	}
}

func TestContactWithCsvTagLessThanZero(t *testing.T) {
	_, err := csvParser.Parse(ExampleContactWithCsvTagLessThanZero{})

	if err == nil {
		t.Error("ExampleContactWithCsvTagLessThanZero should return an error")
	}
}

func TestContactWithCsvTagNotAnInteger(t *testing.T) {
	_, err := csvParser.Parse(ExampleContactWithCsvTagNotAnInteger{})

	if err == nil {
		t.Error("ExampleContactWithCsvTagNotAnInteger should return an error")
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

func testSingleContact(t *testing.T, c ContactGetter) {
	if c.GetFirstName() != "Tom" {
		t.Errorf("Parsed contact firstname is wrong: Actual %v, Expected %v", c.GetFirstName(), "Tom")
	}
	if c.GetLastName() != "Jones" {
		t.Errorf("Parsed contact lastname is wrong: Actual %v, Expected %v", c.GetLastName(), "Jones")
	}
	if c.GetWorking() != true {
		t.Errorf("Parsed contact working is wrong: Actual %v, Expected %v", c.GetWorking(), true)
	}
	if c.GetAge() != 56 {
		t.Errorf("Parsed contact working is wrong: Actual %v, Expected %v", c.GetAge(), 56)
	}
	if c.GetSalary32() != 42000.32 {
		t.Errorf("Parsed contact salary is wrong: Actual %v, Expected %v", c.GetSalary32(), 42000.32)
	}
	if c.GetSalary64() != 42000.64 {
		t.Errorf("Parsed contact salary is wrong: Actual %v, Expected %v", c.GetSalary64(), 42000.64)
	}
	if c.GetVacationDays() != 10 {
		t.Errorf("Parsed contact vacation days is wrong: Actual %v, Expected %v", c.GetVacationDays(), 10)
	}
	if c.GetTitle() != "Senior Director" {
		t.Errorf("Parsed contact title is wrong: Actual %v, Expected %v", c.GetTitle(), "Senior Director")
	}
	if c.GetEmail() != "buyer@mymail.com" {
		t.Errorf("Parsed contact email is wrong: Actual %v, Expected %v", c.GetEmail(), "buyer@mymail.com")
	}
	if c.GetBirthdate() != "1999-06-07" {
		t.Errorf("Parsed contact birthdate is wrong: Actual %v, Expected %v", c.GetBirthdate(), "1999-06-07")
	}
	if c.GetDescription() != "Self-described as \"the top\" branding guru on the West Coast" {
		t.Errorf("Parsed contact description is wrong: Actual %v, Expected %v", c.GetDescription(), "Self-described as \"the top\" branding guru on the West Coast")
	}
}
