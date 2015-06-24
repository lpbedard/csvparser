# CSV Parser [![Build Status](https://travis-ci.org/empatica/csvparser.svg?branch=master)](https://travis-ci.org/empatica/csvparser) [![codecov.io](http://codecov.io/github/empatica/csvparser/coverage.svg?branch=master)](http://codecov.io/github/empatica/csvparser?branch=master)

Simple library that parse a CSV file and returns an array of pointers to the given type.

### Limitations

- Works only with a struct that contains string, int, uint, bool and float fields.

## Getting started

### Install

    go get -u github.com/empatica/csvparser

### Usage

Define your struct:

    type YourStruct struct{
      Field1 string
      Field2 int
      Field3 bool
      Field4 float64
    }

If you don't add 'csv' tags close to each struct's field, the lib will set the first field using the first column of csv's row, and so on. So the previous struct is the same as:

    type YourStruct struct{
      Field1 string  `csv:"0"`
      Field2 int     `csv:"1"`
      Field3 bool    `csv:"2"`
      Field4 float64 `csv:"3"`
    }

You can always define 'csv' tags (for all or some of the struct's fields) that will tell the lib which column to use:

    type YourStruct struct{
      Field1 string  `csv:"1"`
      Field2 int     `csv:"0"`
      Field3 bool    `csv:"3"`
      Field4 float64 `csv:"2"`
    }

Parse the file:

    var csvParser = csvparser.CsvParser{
        CsvFile:      "path_to_your_file.csv",
        CsvSeparator: ',',
    }

    var parsedItems, err = csvParser.Parse(YourStruct{})

    for i:=0;i<len(parsedItems);i++{
      log.Print(parsedItems[0].(*YourStruct))
    }
