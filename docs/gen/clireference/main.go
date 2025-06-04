package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"slices"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd"
	"github.com/mattn/go-shellwords"
	"github.com/urfave/cli/v2"
)

var overrides = map[string]string{
	"s3":    "AWS S3 and compliant",
	"gcs":   "Google Cloud Storage",
	"koofr": "Koofr / Digi Storage",
}

var summary strings.Builder

func main() {
	app := cmd.App
	var sb strings.Builder
	sb.WriteString("# CLI Reference\n\n")
	sb.WriteString("{% code fullWidth=\"true\" %}\n")
	sb.WriteString("```\n")
	sb.WriteString(getStdout([]string{}))
	sb.WriteString("```\n")
	sb.WriteString("{% endcode %}\n")
	err := os.MkdirAll("docs/en/cli-reference", 0755)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("docs/en/cli-reference/README.md", []byte(sb.String()), 0644) //nolint:gosec
	if err != nil {
		panic(err)
	}
	summary.WriteString("* [Menu](cli-reference/README.md)\n")
	for _, command := range app.Commands {
		if command.Name == "help" {
			continue
		}
		saveMarkdown(command, path.Join("docs/en/cli-reference"), []string{command.Name})
	}

	currentSummary, err := os.ReadFile("docs/en/SUMMARY.md")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(currentSummary), "\n")
	beginIndex := slices.IndexFunc(lines, func(line string) bool {
		return strings.Contains(line, "<!-- cli begin -->")
	})
	endIndex := slices.IndexFunc(lines, func(line string) bool {
		return strings.Contains(line, "<!-- cli end -->")
	})
	if err != nil {
		panic(err)
	}

	lines = append(lines[:beginIndex+1], append([]string{"", summary.String()}, lines[endIndex:]...)...)
	err = os.WriteFile("docs/en/SUMMARY.md", []byte(strings.Join(lines, "\n")), 0644) //nolint:gosec
	if err != nil {
		panic(err)
	}
}

func convertHyphenatedString(input string) string {
	words := strings.Split(input, "-")

	for i, word := range words {
		// Convert the first character to uppercase and concatenate it with the rest of the word.
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}

	return strings.Join(words, " ")
}

func saveMarkdown(command *cli.Command, outDir string, args []string) {
	var err error
	var outFile string
	var sb strings.Builder

	if len(command.Subcommands) == 0 {
		outFile = path.Join(outDir, command.Name+".md")
	} else {
		outFile = path.Join(outDir, command.Name, "README.md")
		err = os.MkdirAll(path.Join(outDir, command.Name), 0755)
		if err != nil {
			panic(err)
		}
	}

	sb.WriteString(fmt.Sprintf("# %s\n\n", command.Usage))
	sb.WriteString("{% code fullWidth=\"true\" %}\n")
	sb.WriteString("```\n")
	stdout := getStdout(args)
	sb.WriteString(stdout)
	sb.WriteString("```\n")
	sb.WriteString("{% endcode %}\n")
	err = os.WriteFile(outFile, []byte(sb.String()), 0644) //nolint:gosec
	if err != nil {
		panic(err)
	}

	var margin string
	for range len(args) - 1 {
		margin += "  "
	}

	name := convertHyphenatedString(command.Name)
	if strings.Contains(stdout, "singularity datasource add") && len(command.Subcommands) <= 1 {
		name = command.Usage
	}
	if newName, ok := overrides[command.Name]; ok {
		name = newName
	}

	i := strings.Index(outFile, "cli-reference")
	summary.WriteString(fmt.Sprintf("%s* [%s](%s)\n", margin, name, outFile[i:]))
	for _, subcommand := range command.Subcommands {
		if subcommand.Name == "help" {
			continue
		}
		saveMarkdown(subcommand, path.Join(outDir, command.Name), append(args, subcommand.Name))
	}
}

func getStdout(args []string) string {
	args = append([]string{"singularity"}, args...)
	args = append(args, "--help")
	command := strings.Join(args, " ")
	stdout, stderr, err := runArgsInTest(context.TODO(), command)
	if err != nil {
		panic(err)
	}
	if stderr != "" {
		panic(stderr)
	}
	return stdout
}

func runArgsInTest(ctx context.Context, args string) (string, string, error) {
	// Create a clone of the app so that we can run from different tests concurrently
	parser := shellwords.NewParser()
	parser.ParseEnv = true // Enable environment variable parsing
	parsedArgs, err := parser.Parse(args)
	if err != nil {
		return "", "", errors.WithStack(err)
	}

	outWriter := bytes.NewBuffer(nil)
	errWriter := bytes.NewBuffer(nil)

	// Overwrite the stdout and stderr
	cmd.App.Writer = outWriter
	cmd.App.ErrWriter = errWriter

	err = cmd.App.RunContext(ctx, parsedArgs)
	return outWriter.String(), errWriter.String(), err
}
