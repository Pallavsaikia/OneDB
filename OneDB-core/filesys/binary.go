package filesys

import (
	"bytes"
	"encoding/gob"
	"os"
)

type Serializable interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

func GobEncode(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	return buffer.Bytes(), err
}

func GobDecode(data []byte, target interface{}) error {
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)
	return decoder.Decode(target)
}

func WriteSchemaToFile(data Serializable, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoded, err := data.Encode()
	if err != nil {
		return err
	}

	_, err = file.Write(encoded)
	return err
}


// LoadFromFile loads a Serializable struct from a file.
func ReadSchemaToFile( data Serializable,filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}
	return data.Decode(buffer)
}