package main

import (
	"bytes"
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path"
	"strings"
)

var overrides = map[string]string{
	"s3": "AWS S3 and compliant",
}

func main() {
	app := cmd.App
	for _, command := range app.Commands {
		saveMarkdown(command, path.Join("docs/cli-reference"), []string{command.Name})
	}
	var sb strings.Builder
	sb.WriteString("# CLI Reference\n\n")
	sb.WriteString("```\n")
	sb.WriteString(getStdout([]string{}))
	sb.WriteString("```\n")
	err := os.WriteFile("docs/cli-reference/README.md", []byte(sb.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func saveMarkdown(command *cli.Command, outDir string, args []string) {
	var err error
	var outFile string
	if len(command.Subcommands) == 0 {
		outFile = path.Join(outDir, command.Name+".md")
	} else {
		outFile = path.Join(outDir, command.Name, "README.md")
		err = os.MkdirAll(path.Join(outDir, command.Name), 0755)
		if err != nil {
			panic(err)
		}
		for _, subcommand := range command.Subcommands {
			saveMarkdown(subcommand, path.Join(outDir, command.Name), append(args, subcommand.Name))
		}
	}

	var sb strings.Builder

	name := command.Usage
	if newName, ok := overrides[command.Name]; ok {
		name = newName
	}
	sb.WriteString(fmt.Sprintf("# %s\n\n", name))
	sb.WriteString("```\n")
	sb.WriteString(getStdout(args))
	sb.WriteString("```\n")
	err = os.WriteFile(outFile, []byte(sb.String()), 0644)
	if err != nil {
		panic(err)
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
