package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//Parser parse a csv file and returns an array of pointers of the type specified
type Parser interface {
	Parse(resultType interface{})
}

//CsvParser parses a csv file and returns an array of pointers the type specified
type CsvParser struct {
	CsvFile         string
	CsvSeparator    rune
	SkipFirstLine   bool
	SkipEmptyValues bool
}

//Parse creates the array of the given type from the csv file
func (parser CsvParser) Parse(f interface{}) ([]interface{}, error) {

	csvFile, err := os.Open(parser.CsvFile)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	var csvReader = csv.NewReader(csvFile)
	csvReader.Comma = parser.CsvSeparator

	var results = make([]interface{}, 0, 0)

	resultType := reflect.ValueOf(f).Type()

	if parser.SkipFirstLine {
		csvReader.Read()
	}

	for {

		rawCSVLine, err := csvReader.Read()
		if err != nil {
			if fmt.Sprint(err) == "EOF" {
				break
			} else {
				return nil, err
			}
		}

		var newResult = reflect.New(resultType).Interface()

		// set all the struct fields
		for fieldIndex := 0; fieldIndex < resultType.NumField(); fieldIndex++ {
			var currentField = resultType.Field(fieldIndex)

			var csvTag = currentField.Tag.Get("csv")

			if csvTag == "-" {
				continue
			}

			var csvColumnIndex, csvTagErr = strconv.Atoi(csvTag)

			if csvTagErr != nil {
				if csvTag == "" {
					csvColumnIndex = fieldIndex
				} else {
					return nil, csvTagErr
				}
			}

			if csvColumnIndex < 0 {
				return nil, fmt.Errorf("csv tag in struct field %v is less than zero", currentField.Name)
			}

			if csvColumnIndex >= len(rawCSVLine) {
				return nil, fmt.Errorf("Trying to access csv column %v for field %v, but csv has only %v column(s)", csvColumnIndex, currentField.Name, len(rawCSVLine))
			}

			var csvElement = rawCSVLine[csvColumnIndex]
			var settableField = reflect.ValueOf(newResult).Elem().FieldByName(currentField.Name)

			if csvElement == "" && parser.SkipEmptyValues {
				continue
			}

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

			case "Time":
				var date time.Time
				var err error
				csvDateFormats := strings.Split(currentField.Tag.Get("csvDate"), ",")

				// loop on formats
				for _, csvDate := range csvDateFormats {
					date, err = time.Parse(csvDate, csvElement)
					if err == nil {
						settableField.Set(reflect.ValueOf(date))
						break
					}
				}
				if err != nil {
					return nil, err
				}
			}
		}

		results = append(results, newResult)
	}
	return results, err
}
