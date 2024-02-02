package datatype

type DataType struct {
	Type string
	Size string
}


var Str = DataType{
	Type: "string",
	Size: "65535",
}

var Integer = DataType{
	Type: "integer",
	Size: "64",
}
var Float = DataType{
	Type: "integer",
	Size: "64",
}

var Bool = DataType{
	Type: "boolean",
	Size: "1",
}

var DateTime = DataType{
	Type: "datetime",
	Size: "1",
}
