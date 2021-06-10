package render

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/suhanyujie/proto-err-reason/model/vo"
)

func Render(list []*vo.OneEnumVar, outputPath, targetFilePkgName string) error {
	tmpl, err := template.New("error_code.tpl").Parse(ErrCodeTpl)
	if err != nil {
		return errors.Wrap(err, "new tpl error")
	}

	_, err = os.Stat(outputPath)
	if err != nil {
		if os.IsNotExist(err) {
			// 创建文件
			log.Printf("输出文件不存在，创建新文件")
			fh, err := os.Create(outputPath)
			if err != nil {
				return errors.Wrap(err, "create file error.")
			}
			fh.Close()
		} else {
			return errors.Wrap(err, "file stat error.")
		}
	}
	fh, err := os.OpenFile(outputPath, os.O_RDWR, 0777)
	if err != nil {
		return errors.Wrap(err, "open file error.")
	}
	defer fh.Close()
	valueObj := vo.OutputVar{
		ListStr:     getOutputListStr(list),
		PackageName: targetFilePkgName,
	}
	err = tmpl.Execute(fh, valueObj)
	if err != nil {
		return errors.Wrap(err, "exec tpl error")
	}
	return nil
}

func getOutputListStr(list []*vo.OneEnumVar) string {
	textList := make([]string, len(list))
	for _, item := range list {
		textList = append(textList, fmt.Sprintf("\t%s ErrCode = %d // %s\n", item.VarName, item.CodeVal, item.Comment))
	}
	return strings.Join(textList, "")
}
