# structr

[![wercker status](https://app.wercker.com/status/b61410dc565d9d7c6348d926068b5382/s "wercker status")](https://app.wercker.com/project/bykey/b61410dc565d9d7c6348d926068b5382)

## Installation

```shell
$ go get github.com/dameleon/structr
```

## Synopsis

```
structr generate -c ${YOUR_CONFIGURATION_YAML_FILE} ${INPUT_FILE_PATH}...
structr template
structr help
```

## Usage

### 1. Write your configuration in yaml file

For example, translating JSONSchema type to swift type.

```yaml
## structr configulation file

# Definitions for translating each type of JSONSchema.
type_translate_map:
  string: "String"
  integer: "Int"
  number: "Float"
  boolean: "Bool"
  null: "nil"
  array: "[{{.InnerType}}]"
  object: "{{.InnerType}}"

# If execute structr with "outDir" option, structr outputs file(s) with the definition of "output_filename" template.
output_filename: "{{.Name}}.swift"

# To output the structure of JSONSchema in dependency(a.k.a specified of "$ref" key in JSONSchema).
output_dependencies: true

# Templates for output the structure
structure_template: |
  struct {{.Name|toUpperCamelCase}} {
  {{range .Properties}}
      var {{.Name}}: {{.Type|translateTypeName|toUpperCamelCase}}?{{end}}

  {{.Children|extractStructures}}
  }

# If the structure has children, to specify strings to nest structure of children
child_structures_nesting: "    "
```

Also, you can get configulation file template following command

```shell
$ structr template > ${YOUR_CONFIGURATION_YAML_FILE}
```

### 2. Generate structure definition(s)

```shell
# output with stdout
$ structr generate -c ${YOUR_CONFIGURATION_YAML_FILE} ${INPUT_FILE_PATH}

struct YourStruct {

    var hoo: String?   
    var bar: String?  
    var child: ChildStruct?
    var dependency: YourDependencyStruct?

    struct ChildStruct {
    
        var baz: String?

    }

}

struct YourDependencyStruct {

    var hoge: String?
    var fuga: String?
    var piyo: String?

}


# output with file
$ structr generate -c ${YOUR_CONFIGURATION_YAML_FILE} --ourDir ${OUTPUT_DIR_PATH} ${INPUT_FILE_PATH}
$ tree ${OUTPUT_DIR_PATH}
OUTPUT_DIR_PATH
├── YourStruct.swift
└── YourDependencyStruct.swift
```

## NOTE

### configure template

#### passed datas

##### In "type_translate_map" each values

```golang
struct {
    .Type string
    .InnerType string
}
```


##### In "output_filename" and "structure_template"

`StructureNode`

for more details, see [nodes.go](./blob/master/nodes.go)


#### helpers

##### for "type_translate_map", "output_filename", "structure_template"

- `toUpperCamelCase` : returns upper camelized string
    - `{"foo-bar-baz" | toUpperCamelCase} -> "FooBarBaz"`
    - `{"foo bar baz" | toUpperCamelCase} -> "FooBarBaz"`
    - `{"fooBarBaz" | toUpperCamelCase} -> "FooBarBaz"`
- `toLowerCamelCase` : returns lower camelized string
    - `{"foo-bar-baz" | toLowerCamelCase} -> "fooBarBaz"`
    - `{"foo bar baz" | toLowerCamelCase} -> "fooBarBaz"`
    - `{"fooBarBaz" | toLowerCamelCase} -> "fooBarBaz"`

##### for "structure_template"

- `translateTypeName` : returns translated string of `TypeNode`
    - ex: `{"type_translate_map": { "string": "TypedString" }}`
        - `{TypeNode{ Name: "string" } | translateTypeName} -> "TypedString"`
    - ex: `{"type_translate_map": { "string": "TypedString", "array": "[]{{.InnerType}}" }}`
        - `{TypeNode{ Name: "array", InnerType: "string" } | translateTypeName} -> "[]TypedString"`
- `extractStructures` : returns extracted StructureNodes(for StructureNode.Children)


## Author

[dameleon](https://twitter.com/damele0n)<dameleon[at]gmail.com>

## LICENSE

The MIT License (MIT)


## Acknowledgement

- [github.com/codegangsta/cli](https://github.com/codegangsta/cli)
- [github.com/jteeuwen/go-bindata](https://github.com/jteeuwen/go-bindata)
- [github.com/xeipuuv/gojsonpointer](https://github.com/xeipuuv/gojsonpointer)
- [github.com/xeipuuv/gojsonreference](https://github.com/xeipuuv/gojsonreference)
- [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2)

### for debug

- [github.com/k0kubun/pp](https://github.com/k0kubun/pp)
