package main

import (
	// o que muda nessa aula é esse pacote
	// é extremamente recomendável usar esse html template pois ele protege vc de ataques
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	templates := []string{
		"header.html",
		"content.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{
		"ToUpper": ToUpper,
	})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}

}
