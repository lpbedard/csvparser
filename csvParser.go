package parser

import (
	"encoding/csv"
	"log"
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

			//log.Print(currentField.Type)

			var csvElement = lines[l][csvColumn]
			var settableField = reflect.ValueOf(x).Elem().FieldByName(currentField.Name)

			switch currentField.Type.Name() {

			case "bool":
				var parsedBool, err = strconv.ParseBool(csvElement)
				if err != nil {
					log.Fatalf("Cannot convert %v to bool", csvElement)
				}
				settableField.SetBool(parsedBool)

			case "uint", "uint8", "uint16", "uint32", "uint64":
				var parsedUint, err = strconv.ParseUint(csvElement, 10, 64)
				if err != nil {
					log.Fatalf("Cannot convert %v to uint", csvElement)
				}
				settableField.SetUint(uint64(parsedUint))

			case "int", "int32", "int64":
				var parsedInt, err = strconv.Atoi(csvElement)
				if err != nil {
					log.Fatalf("Cannot convert %v to int", csvElement)
				}
				settableField.SetInt(int64(parsedInt))

			case "float32":
				var parsedFloat, err = strconv.ParseFloat(csvElement, 32)
				if err != nil {
					log.Fatalf("Cannot convert %v to float32", csvElement)
				}
				settableField.SetFloat(parsedFloat)

			case "float64":
				var parsedFloat, err = strconv.ParseFloat(csvElement, 64)
				if err != nil {
					log.Fatalf("Cannot convert %v to float64", csvElement)
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
