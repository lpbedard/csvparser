# CSV Parser

Simple library that parse a CSV file and returns an array of the given type

### Limitations

- It works only with a struct that contains string fields

## Getting started

### Install

    go get -u github.com/empatica/csvparser


### Parse CSV file

    var csvParser = csvparser.CsvParser{
        CsvFile:      "path_to_your_file.csv",
        CsvSeparator: ',',
    }

    var parsedItems, err = csvParser.Parse(YourStruct{})

    //print parsedItems

    for i:=0;i<len(parsedItems);i++{
      log.Print(parsedItems[0].(*YourStruct))
    }
