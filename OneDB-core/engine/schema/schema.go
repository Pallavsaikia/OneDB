package schema

type Schema struct {
	Fields           []Field `json:"fields"`
	SchemaName       string  `json:"schema_name"`
	RelativeLocation string  `json:"relative_location"`
}
