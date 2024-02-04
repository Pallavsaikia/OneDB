package proto

import (
	"fmt"
	"io"
	"onedb-core/engine/schema"
	"onedb-core/filesys"
	"strconv"
	"strings"
)

func generateFieldProtoContent(fields []schema.Field) string {
	field_lines := []string{}
	for _, field := range fields {
		currentline := ""
		if field.NOT_NULL {
			currentline = " 	" + "required " + field.DATATYPE.String() + " " + field.NAME + " = " + strconv.Itoa(field.COLUMN_INDEX+1) + ";"
		} else {
			currentline = " 	" + field.DATATYPE.String() + " " + field.NAME + " = " + strconv.Itoa(field.COLUMN_INDEX+1) + ";"
		}
		field_lines = append(field_lines, currentline)
	}
	return strings.Join(field_lines, "\n")
}
func generateProtoContent(schema schema.Schema) (string, error) {
	proto_lines := []string{}
	proto_lines = append(proto_lines, "syntax = \"proto3\";")
	proto_lines = append(proto_lines, "")
	proto_lines = append(proto_lines, "message "+schema.SchemaName+"{")
	proto_lines = append(proto_lines, generateFieldProtoContent(schema.Fields))
	proto_lines = append(proto_lines, "}")
	proto_content := strings.Join(proto_lines, "\n")
	return proto_content, nil
}


func GenerateProtoFile(schema schema.Schema, filepath string) error {
	content, _ := generateProtoContent(schema)
	file, err := filesys.CreateFileAndPathIfNotExist(filepath + schema.SchemaName + ".proto")
	if err != nil {
		return fmt.Errorf("error creating proto file:%e", err)
	}
	defer file.Close()
	_, err = io.Copy(file, strings.NewReader(content))
	if err != nil {
		return fmt.Errorf("error writing to proto file:%e", err)
	}
	return nil
}
