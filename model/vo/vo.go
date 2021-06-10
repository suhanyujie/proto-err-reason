package vo

type OneEnumVar struct {
	VarName string
	CodeVal int
	Comment string
}

type Proto struct {
	Syntax    ProtoSyntax           `json:"syntax"`
	ProtoBody []*ProtoProtoBodyItem `json:"protoBody"`
}

type ProtoProtoBodyItem struct {
	Name        string `json:"name"`
	Modifier    int    `json:"modifier"`
	OptionName  string `json:"optionName"`
	MessageName string `json:"messageName"`

	EnumName string      `json:"enumName"` // 只需要确定 EnumName 和 EnumBody 字段
	EnumBody []*EnumItem `json:"enumBody"`
}

type ProtoSyntax struct {
	ProtobufVersion string  `json:"protobufVersion"`
	Comments        *string `json:"comments"`
	InlineComment   *string `json:"inlineComment"`
	Meta            Meta    `json:"meta"`
}

type Meta struct {
	Pos     MetaPos     `json:"pos"`
	LastPos MetaLastPos `json:"lastPos"`
}

type MetaLastPos MetaPos

type MetaPos struct {
	Filename string `json:"filename"`
	Offset   int    `json:"offset"`
	Line     int    `json:"line"`
	Column   int    `json:"column"`
}

type EnumObj struct {
	EnumName string      `json:"enumName"`
	EnumBody []*EnumItem `json:"enumBody"`
}

type EnumItem struct {
	Ident            string                `json:"ident"`
	Number           string                `json:"number"`
	EnumValueOptions interface{}           `json:"enumValueOptions"`
	Comments         interface{}           `json:"comments"`
	InlineComment    EnumItemInlineComment `json:"inlineComment"`
	Meta             Meta                  `json:"meta"`
}

type EnumItemInlineComment struct {
	Raw  string `json:"raw"`
	Meta Meta   `json:"meta"`
}

type OutputVar struct {
	ListStr     string `json:"listStr"`
	PackageName string `json:"packageName"`
}
