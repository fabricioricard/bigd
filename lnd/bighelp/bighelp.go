package pkthelp

type Field struct {
	Name        string
	Description []string
	Repeated    bool
	Type        Type
}

type Varient struct {
	Name        string
	Description []string
}

type Type struct {
	Name        string
	Description []string
	Fields      []Field
}

type Method struct {
	Name             string
	Service          string
	Category         string
	ShortDescription string
	Description      []string
	Req              Type
	Res              Type
}

var EnumVarientType Type = Type{
	Name: "ENUM_VARIENT",
}
