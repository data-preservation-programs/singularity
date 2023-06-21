package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

type SwaggerSpec struct {
	Paths map[string]Path `json:"paths"`
}

type Path = map[string]Operation

type Operation struct {
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}

func main() {
	content, err := os.ReadFile("./api/docs/swagger.json")
	if err != nil {
		panic(err)
	}

	var spec SwaggerSpec
	err = json.Unmarshal(content, &spec)
	if err != nil {
		panic(err)
	}

	contentMap := map[string]*strings.Builder{}
	var summaries []string

	var pathsSorted []string
	for path := range spec.Paths {
		pathsSorted = append(pathsSorted, path)
	}
	slices.Sort(pathsSorted)
	for _, pathName := range pathsSorted {
		pathObj := spec.Paths[pathName]
		var methodsSorted []string
		for method := range pathObj {
			methodsSorted = append(methodsSorted, method)
		}
		slices.Sort(methodsSorted)
		for _, method := range methodsSorted {
			operation := pathObj[method]
			tag := operation.Tags[0]
			if contentMap[tag] == nil {
				contentMap[tag] = &strings.Builder{}
				contentMap[tag].WriteString("# " + tag + "\n\n")
			}
			contentMap[tag].WriteString(fmt.Sprintf("{%% swagger src=\"https://raw.githubusercontent.com/data-preservation-programs/singularity/main/api/docs/swagger.yaml\" path=\"%s\" method=\"%s\" %%}\n", pathName, method))
			contentMap[tag].WriteString("[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/api/docs/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/api/docs/swagger.yaml)\n")
			contentMap[tag].WriteString("{% endswagger %}\n\n")
		}
	}

	err = os.MkdirAll("./docs/en/web-api-reference", 0755)
	if err != nil {
		panic(err)
	}

	var contentMapSorted []string
	for tag := range contentMap {
		contentMapSorted = append(contentMapSorted, tag)
	}
	slices.Sort(contentMapSorted)
	for _, tag := range contentMapSorted {
		builder := contentMap[tag]
		err := os.WriteFile("./docs/en/web-api-reference/"+convertStringToHyphenated(tag)+".md", []byte(builder.String()), 0644)
		if err != nil {
			panic(err)
		}
		summaries = append(summaries, fmt.Sprintf("* [%s](web-api-reference/%s.md)", tag, convertStringToHyphenated(tag)))
	}

	currentSummary, err := os.ReadFile("docs/en/SUMMARY.md")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(currentSummary), "\n")
	webReferenceLineIndex := slices.IndexFunc(lines, func(line string) bool {
		return strings.Contains(line, "Web API Reference")
	})
	faqIndex := slices.IndexFunc(lines, func(line string) bool {
		return strings.Contains(line, "FAQ")
	})

	slices.Sort(summaries)
	summaries = append(summaries, "* [Specification](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/api/docs/swagger.yaml)", "")
	lines = append(lines[:webReferenceLineIndex+1], append([]string{"", strings.Join(summaries, "\n")}, lines[faqIndex:]...)...)
	err = os.WriteFile("docs/en/SUMMARY.md", []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}

func convertStringToHyphenated(input string) string {
	// Replace spaces with hyphens
	withHyphens := strings.ReplaceAll(input, " ", "-")
	// Convert the string to lowercase
	return strings.ToLower(withHyphens)
}
