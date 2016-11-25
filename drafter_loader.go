package main

type DrafterLoader interface {
	Load(filepath string) (DrafterParseResult, error)
}
