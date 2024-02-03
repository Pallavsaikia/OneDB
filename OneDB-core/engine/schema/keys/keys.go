package keys

import (
	"fmt"
)

type PrimaryKeyGenerateType int

const (
	TimeStamp PrimaryKeyGenerateType = iota + 1
	AutoIncreament
)

type PRIMARY_KEY struct {
	KeyType PrimaryKeyGenerateType `json:"primary_key_gen_type"`
}

func (p PRIMARY_KEY) Validate() error {
	if p.KeyType != TimeStamp && p.KeyType != AutoIncreament {
		return fmt.Errorf("error:primary key type is not valid")
	}
	return nil
}

type FOREIGN_KEY struct {
	RelatedSchemaName string `json:"relation_schema_name"`
}
type UNIQUE_KEY struct {
	Unique bool `json:"unique"`
}
