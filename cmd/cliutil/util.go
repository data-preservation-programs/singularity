package cliutil

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/fatih/color"
	"github.com/logrusorgru/aurora/v4"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
)

var ErrIncorrectNArgs = errors.New("incorrect number of arguments")

func Success(msg string) string {
	return aurora.BrightGreen(msg).String()
}

func Failure(msg string) string {
	return aurora.BrightRed(msg).String()
}

func Warning(msg string) string {
	return aurora.BrightYellow(msg).String()
}

func CheckNArgs(c *cli.Context) error {
	required := strings.Count(c.Command.ArgsUsage, "<")
	optional := strings.Count(c.Command.ArgsUsage, "[")
	if c.Args().Len() < required || c.Args().Len() > required+optional {
		cli.ShowSubcommandHelp(c)
		return ErrIncorrectNArgs
	}
	return nil
}

var ReallyDotItFlag = &cli.BoolFlag{
	Name:  "really-do-it",
	Usage: "Really do it",
}

func HandleReallyDoIt(context *cli.Context) error {
	if !context.Bool("really-do-it") {
		return cli.Exit("You must pass --really-do-it to do this.", 1)
	}
	return nil
}

func PrintAsJSON(obj any) {
	objJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println("Error: Unable to marshal object to JSON.")
		return
	}
	fmt.Println(string(objJSON))
}

func PrintToConsole(c *cli.Context, obj any) {
	useJSON := c.Bool("json")
	verbose := c.Bool("verbose")
	if useJSON {
		PrintAsJSON(obj)
		return
	}

	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Slice {
		printTable(obj, verbose)
	} else {
		printSingleObject(obj, verbose)
	}
}

func isNotEligibleType(field reflect.StructField, verbose bool) bool {
	if verbose {
		return field.Tag.Get("cli") == "verbose" || field.Tag.Get("cli") == "normal"
	}

	return field.Tag.Get("cli") == "normal"
}

func getValue(fieldValue reflect.Value) any {
	var finalValue any
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

func printTable(objects any, verbose bool) {
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
	headers := make([]any, 0, objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if isNotEligibleType(field, verbose) {
			continue
		}
		headers = append(headers, field.Name)
	}

	tbl := table.New(headers...).WithWriter(os.Stdout)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for i := 0; i < value.Len(); i++ {
		objValue := reflect.Indirect(value.Index(i))
		row := make([]any, 0, objType.NumField())
		for j := 0; j < objType.NumField(); j++ {
			field := objType.Field(j)
			fieldValue := objValue.Field(j)
			if isNotEligibleType(field, verbose) {
				continue
			}
			row = append(row, getValue(fieldValue))
		}
		tbl.AddRow(row...)
	}

	tbl.Print()
	fmt.Println()
}
func printSingleObject(obj any, verbose bool) {
	value := reflect.Indirect(reflect.ValueOf(obj))
	objType := value.Type()

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	headers := make([]any, 0, objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if isNotEligibleType(field, verbose) {
			continue
		}
		headers = append(headers, field.Name)
	}

	tbl := table.New(headers...).WithWriter(os.Stdout)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	row := make([]any, 0, objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := value.Field(i)
		if isNotEligibleType(field, verbose) {
			continue
		}
		row = append(row, getValue(fieldValue))
	}
	tbl.AddRow(row...)

	tbl.Print()
	fmt.Println()
}
