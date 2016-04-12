package main

import (
	"text/template"
	"strings"
	"regexp"
	"bytes"
)

type TypeData struct {
	Type string
	InnerType string
}

var commonFuncMap = template.FuncMap{
	"toUpperCamelCase": toUpperCamelCase,
	"toLowerCamelCase": toLowerCamelCase,
}

func NewCommonTemplate(name string, templateString string) (*template.Template) {
	return template.Must(template.New(name).Funcs(commonFuncMap).Parse(templateString))
}

func NewContextualTemplate(context Context, name string) (*template.Template) {
	tmpl := template.Must(template.New(name).Funcs(createFuncMap(context)).Parse(context.Config.StructureTemplate))
	return tmpl
}

func createFuncMap(context Context) (template.FuncMap) {
	var translateTypeName func(typ TypeNode) (string)
	translateTypeName = func(typ TypeNode) (string) {
		if tmpl, ok := context.Config.TypeTranslateMap[typ.Name]; ok {
			var inner string
			if typ.InnerType != nil {
				inner = translateTypeName(*typ.InnerType)
			}
			t := NewCommonTemplate(typ.Name, tmpl)
			var res bytes.Buffer
			t.Execute(&res, TypeData{ typ.Name, inner })
			return res.String()
		}
		return typ.Name
	}
	var extractStructures = func(structures []StructureNode) (string) {
		var res bytes.Buffer
		for _, node := range structures {
			tmpl := NewContextualTemplate(context, node.Name)
			tmpl.Execute(&res, node)
		}
		return res.String()
	}

	return template.FuncMap{
		"toUpperCamelCase": toUpperCamelCase,
		"toLowerCamelCase": toLowerCamelCase,
		"translateTypeName": translateTypeName,
		"extractStructures": extractStructures,
	}
}

/// helpers
func toUpperCamelCase(str string) (string) {
	return strings.Replace(strings.Title(replaceSnakeBodyToSpace(str)), " ", "", -1)
}

func toLowerCamelCase(str string) (string) {
	s := toUpperCamelCase(str)
	f := string(s[0])
	return strings.Replace(s, f, strings.ToLower(f), 1)
}

func replaceSnakeBodyToSpace(str string) (string) {
	r := regexp.MustCompile(`[_-]`)
	return r.ReplaceAllString(str, " ")
}
