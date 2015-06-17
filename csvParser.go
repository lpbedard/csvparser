package parser

import (
	"encoding/csv"
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
			var csvColumn, csvTagErr = strconv.Atoi(currentField.Tag.Get("csv"))

			if csvTagErr != nil {
				csvColumn = i
			}

			reflect.ValueOf(x).Elem().FieldByName(currentField.Name).SetString(lines[l][csvColumn])
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
