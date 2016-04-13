# structr

[![wercker status](https://app.wercker.com/status/b61410dc565d9d7c6348d926068b5382/s "wercker status")](https://app.wercker.com/project/bykey/b61410dc565d9d7c6348d926068b5382)

## Installation

```shell

$ go get github.com/dameleon/structr

```

## Usage

### 1. Write your configuration file in yaml

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
  struct {{.Name}} {
  {{range .Properties}}
      var {{.Name}}: {{.Type|translateTypeName}}?
  {{end}}

  {{.Children|extractStructures}}
  }

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

## LICENSE

The MIT License (MIT)
