package sample

import (
	"github.com/iesreza/foundation/system"
	"html/template"
)

type component struct {
	Path      string
	Assets    string
	Views     string
	Templates *template.Template
}

func (component component) GetTemplates() *template.Template {
	return component.Templates
}

func (component component) ViewPath() string {
	return component.Views
}

func (component component) AssetsPath() string {
	return component.Assets
}

var Component = component{
	Path:   "components/sample/",
	Assets: "components/sample/assets/",
	Views:  "components/sample/views/",
}

func Register() {
	Component.Register()
}

func (component *component) Register() {
	system.Components["sample"] = component

}

func (component *component) Menu() {

}

func (component *component) Install() {
	panic("implement me")
}

func (component *component) Uninstall() {
	panic("implement me")
}

func (component *component) Update() {
	panic("implement me")
}

func (component *component) ComputeHash() {
	panic("implement me")
}
