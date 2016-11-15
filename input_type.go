package main

type InputType int

const (
	INPUT_TYPE_UNKNOWN InputType = iota
	INPUT_TYPE_JSON
)

func StringToInputMode(str string) InputType {
	switch str {
	case INPUT_TYPE_JSON.String():
		return INPUT_TYPE_JSON
	default:
		return INPUT_TYPE_UNKNOWN
	}
}

func (it InputType) String() string {
	switch it {
	case INPUT_TYPE_JSON:
		return "json"
	default:
		return "unknown"
	}
}

func (it InputType) extNames() []string {
	switch it {
	case INPUT_TYPE_JSON:
		return []string{".json"}
	default:
		return []string{}
	}
}