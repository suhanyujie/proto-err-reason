package render

const ErrCodeTpl = `package {{.PackageName}}

type ErrCode int

const (
{{.ListStr}}
)
`
