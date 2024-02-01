package schema

type Schema struct {
	Fields           []Field     `json:"fields"`
	SchemaName       string      `json:"schema_name"`
	RelativeLocation string      `json:"relative_location"`
	PrimaryKey       PRIMARY_KEY `json:"primary_key"`
}


