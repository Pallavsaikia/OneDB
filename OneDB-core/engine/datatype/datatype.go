package datatype

type DataType int

const (
	Integer DataType = iota + 1
	FLoat
	String
	DateTime
)

func (d DataType) ToString() string {
	return [...]string{"Integer", "FLoat", "String", "DateTime"}[d-1]
}

func (d DataType) MaxSize() string {
	switch d {
	case Integer:
		return "64"
	case FLoat:
		return "128"
	case String:
		return "65535"
	case DateTime:
		return "8"
	default:
		return "0"
	}
}
func (d DataType) IsValidDataType() bool {
	return (d >= 1 && d <= 4)
}

func (d DataType) GetEnumIndex() int {
	return int(d)
}
