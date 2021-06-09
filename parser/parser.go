package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	protoparser "github.com/yoheimuta/go-protoparser/v4"
	"os"
	"path/filepath"
)

/// 解析 proto3 的 proto 文件
func ParseToJson(filename string) (string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		errText := fmt.Sprintf("failed to open %s, err %v\n", filename, err)
		fmt.Fprintf(os.Stderr, errText)
		return "", errors.New(errText)
	}
	defer reader.Close()
	isDebug := true
	permissive := true
	unordered := false
	got, err := protoparser.Parse(
		reader,
		protoparser.WithDebug(isDebug),
		protoparser.WithPermissive(permissive),
		protoparser.WithFilename(filepath.Base(filename)),
	)
	if err != nil {
		errText := fmt.Sprintf("failed to parse, err %v\n", err)
		return "", errors.New(errText)
	}

	var v interface{}
	v = got
	if unordered {
		v, err = protoparser.UnorderedInterpret(got)
		if err != nil {
			errText := fmt.Sprintf("failed to interpret, err %v\n", err)
			return "", errors.New(errText)
		}
	}

	gotJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		errText := fmt.Sprintf("failed to marshal, err %v\n", err)
		return "", errors.New(errText)
	}
	jsonContent := string(gotJSON)

	return jsonContent, nil
}
