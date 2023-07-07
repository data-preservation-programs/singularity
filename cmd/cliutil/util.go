package cliutil

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"golang.org/x/exp/slices"
)

func PrintAsJSON(obj interface{}) {
	objJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println("Error: Unable to marshal object to JSON.")
		return
	}
	fmt.Println(string(objJSON))
}

func PrintToConsole(obj interface{}, useJSON bool, except []string) {
	if useJSON {
		PrintAsJSON(obj)
		return
	}
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Slice {
		printTable(obj, except)
	} else {
		printSingleObject(obj, except)
	}
}

func isNotEligibleType(field reflect.StructField, except []string) bool {
	return field.Name == "_" || slices.Contains(except, field.Name) ||
		(field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) ||
		(field.Type.Kind() == reflect.Slice && field.Type.Elem().Kind() == reflect.Struct) ||
		(field.Type.Kind() == reflect.Slice && field.Type.Elem().Kind() == reflect.Uint8)
}

func getValue(fieldValue reflect.Value) interface{} {
	var finalValue interface{}
	if fieldValue.Kind() == reflect.Ptr {
		if fieldValue.IsNil() {
			finalValue = ""
		} else {
			finalValue = fieldValue.Elem().Interface()
		}
	} else if timeValue, ok := fieldValue.Interface().(time.Time); ok {
		finalValue = timeValue.UTC().Format("2006-01-02 15:04:05Z")
	} else {
		finalValue = fieldValue.Interface()
	}
	return finalValue
}

func printTable(objects interface{}, except []string) {
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
		if isNotEligibleType(field, except) {
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
			if isNotEligibleType(field, except) {
				continue
			}
			row = append(row, getValue(fieldValue))
		}
		tbl.AddRow(row...)
	}

	tbl.Print()
	fmt.Println()
}
func printSingleObject(obj interface{}, except []string) {
	value := reflect.Indirect(reflect.ValueOf(obj))
	objType := value.Type()

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	headers := make([]interface{}, 0, objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if isNotEligibleType(field, except) {
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
		if isNotEligibleType(field, except) {
			continue
		}
		row = append(row, getValue(fieldValue))
	}
	tbl.AddRow(row...)

	tbl.Print()
	fmt.Println()
}
