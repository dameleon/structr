package main

import (
	"os"
	"os/exec"
	"strings"
)

type ParseType int;

const (
	PARSE_TYPE_JSON_SCHEMA ParseType = iota
	PARSE_TYPE_API_BLUEPRINT
)

func (pt ParseType) String() string {
	switch pt {
	case PARSE_TYPE_JSON_SCHEMA:
		return "parse_type_json_schema"
	case PARSE_TYPE_API_BLUEPRINT:
		return "parse_type_api_blueprint"
	default:
		return "unknown"
	}
}

var DRAFTER_BIN_ENV_VAR = "DRAFTER_BIN"

type Environment struct {
	JsonSchema bool
	ApiBlueprint bool
	drafterBinPath string
}

func NewEnvironment() Environment {
	// NOTE: JsonSchema is always true
	env := Environment{true, false, ""}
	if binPath := os.Getenv(DRAFTER_BIN_ENV_VAR); binPath != "" {
		info, err := os.Stat(binPath)
		if err == nil && (info.Mode() & 0111) != 0 {
			env.ApiBlueprint = true
			env.drafterBinPath = binPath
		}
	} else {
		// TODO: windows
		out, err := exec.Command("which", "drafter").Output()
		if err == nil && out != nil {
			env.ApiBlueprint = true
			env.drafterBinPath = strings.TrimSpace(string(out))
		}
	}
	return env
}
