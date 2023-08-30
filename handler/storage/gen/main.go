package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"reflect"
	"strings"
	"text/template"
	"unicode"

	"github.com/data-preservation-programs/singularity/storagesystem"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const header = `//lint:file-ignore U1000 Ignore all unused code, it's generated
// Code generated. DO NOT EDIT.
package storage

import (
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/lib/encoder"
)
`

const structTemplate = `
type {{.Name}} struct {
	{{- range .Fields }}
    {{.Name}} {{.Type}} {{.Tag}} // {{.Description}}
	{{- end }}
}
`

const createStructTemplate = `
type {{.Name}} struct {
	Name string ` + "`json:\"name\" example:\"my-storage\"`" + ` // Name of the storage, must be unique
	Path string ` + "`json:\"path\"`" + ` // Path of the storage
	Config {{.ConfigType}}
}
`

type CreateStructType struct {
	Name       string
	ConfigType string
}

const s3Template = `
// @Summary Create {{.Name}} storage with {{.Provider}} - {{.ProviderDescription}}
// @Tags Storage
// @Accept json
// @Produce json
// @Param provider path string true "Provider name"
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body {{.StructName}} true "Request body"
// @Router /storage/{{.Type}}/{{.ProviderLower}} [post]
func {{.FuncName}}() {}

`

const otherTemplate = `
// @Summary Create {{.Name}} storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body {{.StructName}} true "Request body"
// @Router /storage/{{.Type}} [post]
func {{.FuncName}}() {}

`

type S3Func struct {
	FuncName            string
	Provider            string
	ProviderLower       string
	ProviderDescription string
	Name                string
	StructName          string
	Type                string
}
type OtherFunc struct {
	FuncName   string
	Name       string
	StructName string
	Type       string
}

type Field struct {
	Name        string
	Type        string
	Tag         string
	Description string
}

type Type struct {
	Name   string
	Fields []Field
}

func main() {
	f := bytes.NewBuffer(nil)

	// Write package lines
	f.WriteString(header)

	structTemplate, err := template.New("struct").Parse(structTemplate)
	if err != nil {
		panic(err)
	}

	createStructTemplate, err := template.New("struct").Parse(createStructTemplate)
	if err != nil {
		panic(err)
	}

	s3Template, err := template.New("s3").Parse(s3Template)
	if err != nil {
		panic(err)
	}

	otherTemplate, err := template.New("other").Parse(otherTemplate)
	if err != nil {
		panic(err)
	}

	for _, backend := range storagesystem.Backends {
		for _, providerOptions := range backend.ProviderOptions {
			var fields []Field
			for _, option := range providerOptions.Options {
				lowerCamel := snakeToLowerCamel(option.Name)
				upperCamel := snakeToUpperCamel(option.Name)
				tag := fmt.Sprintf("`json:\"%s\"", lowerCamel)
				defaultValue := fmt.Sprintf("%v", option.Default)
				if defaultValue != "" {
					tag += fmt.Sprintf(" default:\"%v\"", defaultValue)
				}
				if len(option.Examples) > 0 {
					example := option.Examples[0].Value
					tag += fmt.Sprintf(" example:\"%v\"", example)
				}
				tag += "`"
				if upperCamel == "2fa" {
					upperCamel = "TwoFA"
				}
				tp := "string"
				if option.Default != nil {
					tp = reflect.TypeOf(option.Default).String()
				}
				description := strings.Split(option.Help, "\n")[0]
				fields = append(fields, Field{
					Name:        upperCamel,
					Type:        tp,
					Tag:         tag,
					Description: description,
				})
			}
			typeName := upperFirst(backend.Prefix) + upperFirst(providerOptions.Provider) + "Config"
			t := Type{
				Name:   typeName,
				Fields: fields,
			}
			err = structTemplate.Execute(f, t)
			if err != nil {
				panic(err)
			}

			if len(backend.ProviderOptions) > 1 {
				cobj := CreateStructType{
					Name:       "Create" + upperFirst(backend.Prefix) + upperFirst(providerOptions.Provider) + "StorageRequest",
					ConfigType: typeName,
				}
				err = createStructTemplate.Execute(f, cobj)
				if err != nil {
					panic(err)
				}
				fobj := S3Func{
					FuncName:            "Create" + upperFirst(backend.Prefix) + upperFirst(providerOptions.Provider) + "Storage",
					Name:                upperFirst(backend.Prefix),
					Provider:            providerOptions.Provider,
					ProviderDescription: strings.Split(providerOptions.ProviderDescription, "\n")[0],
					ProviderLower:       strings.ToLower(providerOptions.Provider),
					StructName:          cobj.Name,
					Type:                backend.Prefix,
				}
				err = s3Template.Execute(f, fobj)
				if err != nil {
					panic(err)
				}
			} else {
				cobj := CreateStructType{
					Name:       "Create" + upperFirst(backend.Prefix) + "StorageRequest",
					ConfigType: typeName,
				}
				err = createStructTemplate.Execute(f, cobj)
				if err != nil {
					panic(err)
				}
				oobj := OtherFunc{
					FuncName:   "create" + upperFirst(backend.Prefix) + "Storage",
					Name:       upperFirst(backend.Prefix),
					StructName: cobj.Name,
					Type:       backend.Prefix,
				}
				err = otherTemplate.Execute(f, oobj)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		panic(err)
	}
	// Create generated file in the same directory
	//nolint:gosec
	err = os.WriteFile("handler/storage/types_gen.go", formatted, 0644)
	if err != nil {
		panic(err)
	}
}

func snakeToLowerCamel(s string) string {
	words := strings.Split(s, "_")

	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
			continue
		}

		// Only capitalize the first letter of subsequent words (not the first word)
		if i != 0 {
			words[i] = cases.Title(language.English).String(words[i])
		} else {
			// Ensure the first word starts with a lowercase letter
			r := []rune(words[i])
			r[0] = unicode.ToLower(r[0])
			words[i] = string(r)
		}
	}

	return strings.Join(words, "")
}
func snakeToUpperCamel(s string) string {
	words := strings.Split(s, "_")

	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
			continue
		}
		// Capitalize the first letter of each word
		r := []rune(words[i])
		r[0] = unicode.ToUpper(r[0])
		words[i] = string(r)
	}

	return strings.Join(words, "")
}

func upperFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
