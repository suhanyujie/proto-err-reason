package main

import (
	"encoding/json"
	"fmt"
	"github.com/suhanyujie/proto-err-reason/model/vo"
	"github.com/suhanyujie/proto-err-reason/parser"
	"github.com/suhanyujie/proto-err-reason/render"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	app := &cli.App{
		Name:  "hello",
		Usage: "hello world example",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Value: "",
				Usage: "proto file name",
			},
			&cli.StringFlag{
				Name:  "o",
				Value: "example/output_error_code.go",
				Usage: "output file name",
			},
			&cli.StringFlag{
				Name:  "pkg",
				Value: "example",
				Usage: "target file pkg name",
			},
		},
		Action: func(c *cli.Context) error {
			protoFile := c.String("file")
			outputFile := c.String("o")
			targetFilePkgName := c.String("pkg")
			jsonText, err := parser.ParseToJson(protoFile)
			if err != nil {
				log.Fatal(err)
			}
			enumItemList := GetEnumErrorCode(jsonText)
			if err := render.Render(enumItemList, outputFile, targetFilePkgName); err != nil {
				fmt.Printf("%+v", err)
				return nil
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


// GetEnumErrorCode 从 proto json 中获取数据
func GetEnumErrorCode(protoJson string) []*vo.OneEnumVar {
	list := make([]*vo.OneEnumVar, 0)
	protoObj := &vo.Proto{}
	err := json.Unmarshal([]byte(protoJson), protoObj)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	for _, bodyItem := range protoObj.ProtoBody {
		if bodyItem.EnumName != "" {
			for _, enumItem := range bodyItem.EnumBody {
				codeVal, _ := strconv.ParseInt(enumItem.Number, 10, 64)
				list = append(list, &vo.OneEnumVar{
					VarName: enumItem.Ident,
					CodeVal: int(codeVal),
					Comment: TrimEnumItemComment(enumItem.InlineComment.Raw),
				})
			}
		}
	}
	return list
}

// TrimEnumItemComment 去掉注释文本中的双斜线和空格
func TrimEnumItemComment(text string) string {
	text = strings.Replace(text, "//", "", 1)
	text = strings.TrimLeft(text, " ")
	return text
}
