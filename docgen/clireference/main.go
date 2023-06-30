package main

import (
	"bytes"
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
	"os"
	"os/exec"
	"path"
	"strings"
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
	err = os.WriteFile("docs/en/cli-reference/README.md", []byte(sb.String()), 0644)
	if err != nil {
		panic(err)
	}
	for _, command := range app.Commands {
		saveMarkdown(command, path.Join("docs/en/cli-reference"), []string{command.Name})
	}

	currentSummary, err := os.ReadFile("docs/en/SUMMARY.md")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(currentSummary), "\n")
	cliReferenceLineIndex := slices.IndexFunc(lines, func(line string) bool {
		return strings.Contains(line, "CLI Reference")
	})
	webReferenceLineIndex := slices.IndexFunc(lines, func(line string) bool {
		return strings.Contains(line, "Web API Reference")
	})
	if err != nil {
		panic(err)
	}

	lines = append(lines[:cliReferenceLineIndex+1], append([]string{"", summary.String()}, lines[webReferenceLineIndex:]...)...)
	err = os.WriteFile("docs/en/SUMMARY.md", []byte(strings.Join(lines, "\n")), 0644)
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
	err = os.WriteFile(outFile, []byte(sb.String()), 0644)
	if err != nil {
		panic(err)
	}

	var margin string
	for i := 0; i < len(args)-1; i++ {
		margin += "  "
	}

	name := convertHyphenatedString(command.Name)
	if strings.Contains(stdout, "singularity datasource add") && len(command.Subcommands) == 0 {
		name = command.Usage
	}
	if newName, ok := overrides[command.Name]; ok {
		name = newName
	}

	summary.WriteString(fmt.Sprintf("%s* [%s](%s)\n", margin, name, outFile[8:]))
	for _, subcommand := range command.Subcommands {
		saveMarkdown(subcommand, path.Join(outDir, command.Name), append(args, subcommand.Name))
	}
}

func getStdout(args []string) string {
	args = append(args, "--help")
	c := exec.Command("./singularity", args...)
	var stdout bytes.Buffer
	c.Stdout = &stdout

	err := c.Run()
	if err != nil {
		panic(err)
	}

	return stdout.String()
}
