// +build ignore

// This generates countries.go by running "go generate"
package main

import (
	"bufio"
	"os"
	"strings"
	"text/template"
)

type Country struct {
	Code string
	Name string
}

var countries []Country

func main() {
	file, err := os.Open("countries.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")

		if len(parts) < 2 {
			continue
		}

		countries = append(countries, Country{Code: strings.ToLower(parts[0]), Name: strings.ToLower(parts[1])})
	}

	f, err := os.Create("countries.go")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	packageTemplate.Execute(f, struct {
		Countries []Country
	}{
		Countries: countries,
	})
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package photon

var Countries = map[string]string{
{{- range .Countries }}
	{{ printf "%q" .Name }}: {{ printf "%q" .Code }},
{{- end }}
}`))
