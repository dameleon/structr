package main

import "encoding/json"

type DrafterParseResult struct {
	AST Blueprint `json:"ast"`
	Error Annotation `json:"error"`
	Warnings []Annotation `json:"warnings"`
}

type Blueprint struct {
	Version string `json:"_version"`
	Metadata []BlueprintMetadata `json:"meta"`
	Name string `json:"name"`
	Description string `json:"description"`
	Element string `json:"element"`
	ResourceGroup []ResourceGroup `json:"resourceGroup"`
	Content []Element `json:"content"`
}

type BlueprintMetadata struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type ResourceGroup struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Element string `json:"element"`
	Name string `json:"name"`
	Description string `json:"description"`
	UriTemplate string `json:"uriTemplate"`
	Model Payload `json:"model"`
	Parameters []Parameter `json:"parameters"`
	Actions []Action `json:"actions"`
	Content []DataStructure `json:"content"`
}

type Payload struct {
	Reference Reference `json:"reference"`
	Name string `json:"name"`
	Description string `json:"description"`
	Headers []Header `json:"headers"`
	Body string `json:"body"`
	Schema string `json:"schema"`
	Content interface{} `json:"content"`
	rawContent []byte
}

func (p *Payload) getRawContent() []byte {
	if p.rawContent == nil {
		p.rawContent, _ = json.Marshal(p.Content)
	}
	return p.rawContent
}

func (p *Payload) ContentAsDataStructure() []DataStructure {
	var res []DataStructure
	json.Unmarshal(p.getRawContent(), &res)
	return res
}

func (p *Payload) ContentAsAsset() []Asset {
	var res []Asset
	json.Unmarshal(p.getRawContent(), &res)
	return res
}

type Parameter struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	Required bool `json:"required"`
	Default string `json:"default"`
	Example string `json:"example"`
	Values []ParameterValue `json:"values"`
}

type ParameterValue struct {
	Value string `json:"value"`
}

type Action struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Method string `json:"method"`
	Parameters []Parameter `json:"parameters"`
	Attributes ActionAttributes `json:"attributes"`
	Relation string `json:"relation"`
	UriTemplate string `json:"uriTemplate"`
	Content []DataStructure `json:"content"`
	Examples []TransactionExample `json:"examples"`
}

type ActionAttributes struct {
	Relation string `json:"relation"`
	UriTemplate string `json:"uriTemplate"`
}

type TransactionExample struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Requests []Payload `json:"requests"`
	Responses []Payload `json:"responses"`
}

type Reference struct {
	Id string `json:"id"`
}

type Header struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type Annotation struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Location []Location `json:"location"`
}

type Location struct {
	Index int `json:"index"`
	Length int `json:"length"`
}

type Asset struct {
	Element string `json:"element"`
	Attributes AssetAttributes `json:"attributes"`
	Content string `json:"content"`
}

type AssetAttributes struct {
	Role string `json:"role"`
}

type DataStructure struct {
	Element string `json:"element"`
	Content []DataStructureElement `json:"content"`
}

type DataStructureElement struct {
	Element string `json:"element"`
	Metadata DataStructureElementMetadata `json:"meta"`
	Content []DataStructureElementType `json:"content"`
}

type DataStructureElementMetadata struct {
	Id string `json:"id"`
	Description string `json:"description"`
}

type DataStructureElementType struct {
	Element string `json:"element"`
	Metadata DataStructureElementTypeMetadata `json:"meta"`
	Attributes DataStructureElementTypeAttributes `json:"attributes"`
	Content DataStructureElementTypeProperty `json:"content"`
}

type DataStructureElementTypeMetadata struct {
	Description string `json:"description"`
}

type DataStructureElementTypeAttributes struct {
	TypeAttribute []string `json:"typeAttribute"`
	Samples []interface{} `json:"samples"`
	Default interface{} `json:"default"`
}

type DataStructureElementTypeProperty struct {
	Key DataStructureElementTypePropertyKey `json:"key"`
	Value DataStructureElementTypePropertyValue `json:"value"`
}

type DataStructureElementTypePropertyKey struct {
	Element string `json:"element"`
	Content string `json:"content"`
}

type DataStructureElementTypePropertyValue struct {
	Element string `json:"element"`
	Content string `json:"content"`
}

type Element struct {
	Element string `json:"element"`
	Attributes ElementAttributes `json:"attributes"`
	Content interface{} `json:"content"`
	rawContent []byte `json:"rawContent"`
}

func (e *Element) getRawContent() []byte {
	if e.rawContent == nil {
		e.rawContent, _ = json.Marshal(e.Content)
	}
	return e.rawContent
}

func (e *Element) ContentAsCopyElement() string {
	var res string
	json.Unmarshal(e.getRawContent(), &res)
	return res
}

func (e *Element) ContentAsCategoryElement() []Element {
	var res []Element
	json.Unmarshal(e.getRawContent(), &res)
	return res
}

func (e *Element) ContentAsDataStructureElement() DataStructure {
	var res DataStructure
	json.Unmarshal(e.getRawContent(), &res)
	return res
}

func (e *Element) ContentAsResourceElement() Resource {
	var res Resource
	json.Unmarshal(e.getRawContent(), &res)
	return res
}

type ElementAttributes struct {
	Name string `json:"name"`
}
