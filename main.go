package main

import (
	"fmt"
	"github.com/suhanyujie/proto-err-reason/parser"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "hello",
		Usage: "hello world example",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dir",
				Value: ".",
				Usage: "proto file dir",
			},
			&cli.StringFlag{
				Name:  "file",
				Value: "",
				Usage: "proto file name",
			},
		},
		Action: func(c *cli.Context) error {
			protoFile := c.String("file")
			//content, err := ioutil.ReadFile(protoFile)
			//if err != nil {
			//	log.Fatalf("%v", err)
			//}
			//contentStr := string(content)"
			jsonText, err := parser.ParseToJson(protoFile)
			if err != nil {
				log.Fatal(err)
			}
			enumData := GetEnumErrorCode(jsonText)
			fmt.Printf("hello world content: %s\n", jsonText)
			fmt.Printf("hello world content: %v\n", enumData)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type OneEnumVar struct {
	VarName string
	CodeVal int
	Comment string
}

// todo
func GetEnumErrorCode(protoJson string) []*OneEnumVar {
	list := make([]*OneEnumVar, 0)

	return list
}
