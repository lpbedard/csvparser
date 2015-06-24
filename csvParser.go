package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//Parser parse a csv file and returns an array of pointers of the type specified
type Parser interface {
	Parse(outType interface{})
}

//CsvParser parses a csv file and returns an array of pointers the type specified
type CsvParser struct {
	CsvFile      string
	CsvSeparator rune
}

//Parse creates the array of the given type from the csv file
func (parser CsvParser) Parse(f interface{}) ([]interface{}, error) {

	var lines, err = getCsvLines(parser.CsvFile, parser.CsvSeparator)

	var result = make([]interface{}, 0, 0)

	fn := reflect.ValueOf(f)
	outType := fn.Type()

	for l := 0; l < len(lines); l++ {
		//create the new item
		var x = reflect.New(outType).Interface()

		// set all the struct fields
		for i := 0; i < outType.NumField(); i++ {
			var currentField = outType.Field(i)

			var csvTag = currentField.Tag.Get("csv")
			var csvColumn, csvTagErr = strconv.Atoi(csvTag)

			if csvTagErr != nil {
				if csvTag == "" {
					csvColumn = i
				} else {
					return nil, csvTagErr
				}
			}

			if csvColumn < 0 {
				return nil, fmt.Errorf("csv tag in struct field %v is less than zero", currentField.Name)
			}

			if len(lines[l]) < csvColumn {
				return nil, fmt.Errorf("Trying to access csv column %v for field %v, but csv has only %v column", csvColumn, currentField.Name, len(lines[l]))
			}

			var csvElement = lines[l][csvColumn]
			var settableField = reflect.ValueOf(x).Elem().FieldByName(currentField.Name)

			switch currentField.Type.Name() {

			case "bool":
				var parsedBool, err = strconv.ParseBool(csvElement)
				if err != nil {
					return nil, err
				}
				settableField.SetBool(parsedBool)

			case "uint", "uint8", "uint16", "uint32", "uint64":
				var parsedUint, err = strconv.ParseUint(csvElement, 10, 64)
				if err != nil {
					return nil, err
				}
				settableField.SetUint(uint64(parsedUint))

			case "int", "int32", "int64":
				var parsedInt, err = strconv.Atoi(csvElement)
				if err != nil {
					return nil, err
				}
				settableField.SetInt(int64(parsedInt))

			case "float32":
				var parsedFloat, err = strconv.ParseFloat(csvElement, 32)
				if err != nil {
					return nil, err
				}
				settableField.SetFloat(parsedFloat)

			case "float64":
				var parsedFloat, err = strconv.ParseFloat(csvElement, 64)
				if err != nil {
					return nil, err
				}
				settableField.SetFloat(parsedFloat)

			case "string":
				settableField.SetString(csvElement)
			}

		}
		//append the result
		result = append(result, x)
	}
	return result, err
}

//GetCsvLines returns all the lines from a given csvfile using the given separator
func getCsvLines(csvFile string, separator rune) ([][]string, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}

	var csvReader = csv.NewReader(file)
	csvReader.Comma = separator

	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, err
}
