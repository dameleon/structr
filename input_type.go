package main

type InputType string

const (
	INPUT_TYPE_UNKNOWN InputType = "unknown"
	INPUT_TYPE_JSON = "json"
	INPUT_TYPE_API_BLUEPRINT = "api_blueprint"
)

func (it InputType) extNames() []string {
	switch it {
	case INPUT_TYPE_JSON:
		return []string{".json"}
	case INPUT_TYPE_API_BLUEPRINT:
		return []string{".md", ".apib"}
	default:
		return []string{}
	}
}
