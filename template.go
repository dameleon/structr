package main

import (
	"text/template"
	"strings"
	"regexp"
	"bytes"
	"io"
)

var commonFuncMap = template.FuncMap{
	"toUpperCamelCase": toUpperCamelCase,
	"toLowerCamelCase": toLowerCamelCase,
}

func NewTemplate(name string) (*template.Template) {
	return template.New(name).Funcs(commonFuncMap)
}

type StructGenerator interface {
	Generate(wr io.Writer, node StructureNode) (error)
}

func NewStructGenerator(templateString string, typeTranslateMap map[string]string) (StructGenerator, error) {
	g := &structGenerator{ typeTranslateMap, nil }
	var err error
	g.template, err = NewTemplate("StructTemplate").Funcs(template.FuncMap{
		"translateTypeName": g.translateTypeName,
		"extractStructures": g.extractStructures,
	}).Parse(templateString)
	if err != nil {
		return nil, err
	}
	return g, nil
}

type structGenerator struct {
	typeTranslateMap map[string]string
	template *template.Template
}

func (g *structGenerator) Generate(wr io.Writer, node StructureNode) (error) {
	return g.template.Execute(wr, node)
}

func (g *structGenerator) extractStructures(nodes []StructureNode) (string) {
	res := []string{}
	for _, node := range nodes {
		var buf bytes.Buffer
		err := g.Generate(&buf, node)
		if err != nil {
			panic(err)
		}
		res = append(res, buf.String())
	}
	return strings.Join(res, "")
}

func (g *structGenerator) translateTypeName(typ TypeNode) (string) {
	if tmpl, ok := g.typeTranslateMap[typ.Name]; ok {
		var data = struct {
			Type string
			InnerType string
		}{ typ.Name, "" }
		if typ.InnerType != nil {
			data.InnerType = g.translateTypeName(*typ.InnerType)
		}
		t := template.Must(NewTemplate(typ.Name).Parse(tmpl))
		var res bytes.Buffer
		err := t.Execute(&res, data)
		if err != nil {
			panic(err)
		}
		return res.String()
	}
	return typ.Name
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
