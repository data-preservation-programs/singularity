package cliutil

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/table"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var ErrIncorrectNArgs = errors.New("incorrect number of arguments")

func Failure(msg string) string {
	return color.New(color.FgRed).Sprint(msg)
}

func CheckNArgs(c *cli.Context) error {
	required := strings.Count(c.Command.ArgsUsage, "<")
	optional := strings.Count(c.Command.ArgsUsage, "[")
	if c.Args().Len() < required || c.Args().Len() > required+optional {
		err := cli.ShowSubcommandHelp(c)
		if err != nil {
			return errors.WithStack(err)
		}
		return ErrIncorrectNArgs
	}
	return nil
}

var ReallyDotItFlag = &cli.BoolFlag{
	Name:  "really-do-it",
	Usage: "Really do it",
}

var ErrReallyDoIt = errors.New("you must pass --really-do-it to do this")

func HandleReallyDoIt(context *cli.Context) error {
	if !context.Bool("really-do-it") {
		return ErrReallyDoIt
	}
	return nil
}

func PrintAsJSON(c *cli.Context, obj any) {
	objJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println("Error: Unable to marshal object to JSON.")
		return
	}
	_, _ = c.App.Writer.Write(objJSON)
}

func Print(c *cli.Context, obj any) {
	if c.Bool("json") {
		PrintAsJSON(c, obj)
		return
	}

	if c.Bool("verbose") {
		_, _ = c.App.Writer.Write([]byte(table.New(table.WithVerbose()).Render(obj)))
	} else {
		_, _ = c.App.Writer.Write([]byte(table.New().Render(obj)))
	}
}
