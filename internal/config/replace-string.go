package config

import (
	"fmt"
	"html/template"
	"os"
)

type Color struct {
	Color string
	Name  string
}
type FieldInput struct {
	TypeInput, Label string
	Colors           []Color
}
type Input struct {
	InputField []FieldInput
}

func Replace() {
	mock := Input{
		InputField: []FieldInput{
			{
				TypeInput: "text",
				Label:     "Nick name",
			},
			{
				TypeInput: "text",
				Label:     "Kid name",
			},
			{
				TypeInput: "select",
				Label:     "Select",
				Colors: []Color{
					{Color: "Chelsea"},
					{Color: "Arsenal"},
					{Color: "Manchester United"},
					{Color: "Liverpool"},
				},
			},
		}}
	tmpl := template.Must(template.ParseFiles("layout.html"))
	outFile, err := os.Create("vue-test/src/App.vue")
	defer outFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(outFile, mock)

}
