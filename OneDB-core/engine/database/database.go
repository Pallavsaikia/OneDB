package database

import (
	"fmt"
	"onedb-core/libs"
)

type Database struct {
	DATABASE_NAME           string `json:"dbname"`
	DATABASE_FOLDER_ADDRESS string `json:"dbfolder_address"`
}


func (db Database) Validate() error {
	if db.DATABASE_NAME == "" {
		return fmt.Errorf("error:DATABASE_NAME cannot be empty")
	}
	if libs.ContainsSpace(db.DATABASE_NAME) {
		return fmt.Errorf("error:DATABASE_NAME cannot have spaces:'%s'", db.DATABASE_NAME)
	}
	if libs.ContainsNumber(db.DATABASE_NAME) {
		return fmt.Errorf("error:DATABASE_NAME cannot have number:'%s'", db.DATABASE_NAME)
	}
	if libs.ContainsSpecialCharacters(db.DATABASE_NAME) {
		return fmt.Errorf("error:DATABASE_NAME cannot contain special characters other than a hyphen'_':'%s'", db.DATABASE_NAME)
	}

	return nil
}
func (db *Database) Initialize() error {
	err := db.Validate()
	if err != nil {
		return err
	}
	db.DATABASE_FOLDER_ADDRESS = "database/" + db.DATABASE_NAME
	return nil
}
