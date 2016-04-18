package fixtures

import (
	"github.com/xeipuuv/gojsonreference"
	"os"
	"path/filepath"
)

const FIXTURE_DIR = "fixtures"

var wd, _ = os.Getwd()

var BasicJsonSchemaRef, _ = gojsonreference.NewJsonReference(GetSchemaPath("basic.json"))
var InternalReferenceSchemaRef, _ = gojsonreference.NewJsonReference(GetSchemaPath("internal_reference.json"))
var ExternalReferenceSchemaRef, _ = gojsonreference.NewJsonReference(GetSchemaPath("external_reference.json"))
var NotExistsSchemaRef, _ = gojsonreference.NewJsonReference(GetSchemaPath("not_exists.json"))

func GetSchemaPath(name string) string {
	return filepath.Join(wd, FIXTURE_DIR, "schema", name)
}
