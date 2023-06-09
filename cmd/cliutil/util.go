package cliutil

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"reflect"
)

var fieldNamesToSkip = []string{
	"CreatedAt", "UpdatedAt",
}

func PrintToConsole(obj interface{}, jsonOutput bool) {
	if jsonOutput {
		objJSON, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			fmt.Println("Error: Unable to marshal object to JSON.")
			return
		}
		fmt.Println(string(objJSON))
		return
	}

	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Slice {
		printTable(obj)
	} else {
		printSingleObject(obj)
	}
}

func isEligibleType(field reflect.StructField) bool {
	return (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) || field.Type.Kind() == reflect.Slice || field.Name == "CreatedAt" || field.Name == "UpdatedAt"
}

func getValue(fieldValue reflect.Value) interface{} {
	var finalValue interface{}
	if fieldValue.Kind() == reflect.Ptr {
		if fieldValue.IsNil() {
			finalValue = ""
		} else {
			finalValue = fieldValue.Elem().Interface()
		}
	} else {
		finalValue = fieldValue.Interface()
	}
	return finalValue
}

func printTable(objects interface{}) {
	value := reflect.ValueOf(objects)
	if value.Len() == 0 {
		return
	}
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	// Get the first object's type and value
	firstObj := reflect.Indirect(value.Index(0))
	objType := firstObj.Type()

	// Prepare headers using the first object
	headers := make([]interface{}, 0, objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if isEligibleType(field) {
			continue
		}
		headers = append(headers, field.Name)
	}

	tbl := table.New(headers...).WithWriter(os.Stdout)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for i := 0; i < value.Len(); i++ {
		objValue := reflect.Indirect(value.Index(i))
		row := make([]interface{}, 0, objType.NumField())
		for j := 0; j < objType.NumField(); j++ {
			field := objType.Field(j)
			fieldValue := objValue.Field(j)
			if isEligibleType(field) {
				continue
			}
			row = append(row, getValue(fieldValue))
		}
		tbl.AddRow(row...)
	}

	tbl.Print()
	fmt.Println()
}
func printSingleObject(obj interface{}) {
	value := reflect.Indirect(reflect.ValueOf(obj))
	objType := value.Type()

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	headers := make([]interface{}, 0, objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if isEligibleType(field) {
			continue
		}
		headers = append(headers, field.Name)
	}

	tbl := table.New(headers...).WithWriter(os.Stdout)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	row := make([]interface{}, 0, objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := value.Field(i)
		if isEligibleType(field) {
			continue
		}
		row = append(row, getValue(fieldValue))
	}
	tbl.AddRow(row...)

	tbl.Print()
	fmt.Println()
}
