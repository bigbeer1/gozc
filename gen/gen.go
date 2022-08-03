package gen

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"gozc/parser"
	"gozc/tools/pathx"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Table struct {
	parser.Table
	ContainsUniqueCacheKey bool
}

type (
	defaultGenerator struct {
		console.Console
		// source string
		dir          string
		pkg          string
		isPostgreSql bool
	}

	// Option defines a function with argument defaultGenerator
	Option func(generator *defaultGenerator)

	CodeTuple struct {
		Api         string
		ApiInsert   string
		ApiDelete   string
		ApiUpdate   string
		ApiFindOne  string
		ApiFindList string
	}
)

const (
	category             = "api"
	apiTemplateFile      = "api.tpl"
	insertTemplateFile   = "insert.tpl"
	deleteTemplateFile   = "delete.tpl"
	updateTemplateFile   = "update.tpl"
	findOneTemplateFile  = "find-one.tpl"
	findlistTemplateFile = "find-list.tpl"
)

func GenApiModel(in parser.Table, pkgName string, apiType string) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}
	var table Table
	table.Table = in
	switch apiType {
	case "api":
		res, err := genApi(table, pkgName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "Insert":
		res, err := genInsert(table, pkgName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "delete":
		res, err := genDelete(table, pkgName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "update":
		res, err := genUpdate(table, pkgName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "findOne":
		res, err := genFindOne(table, pkgName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "findList":
		res, err := genFindList(table, pkgName)
		if err != nil {
			return "", err
		}
		return res, nil
	}

	return "", nil

}

//func ExecuteModel(pkg string, table Table, code *code) (*bytes.Buffer, error) {
//	text, err := pathx.LoadTemplate(category, modelGenTemplateFile, template.ModelGen)
//	if err != nil {
//		return nil, err
//	}
//	t := util.With("model").
//		Parse(text).
//		GoFmt(true)
//	output, err := t.Execute(map[string]interface{}{
//		"pkg":         pkg,
//		"imports":     code.importsCode,
//		"vars":        code.varsCode,
//		"types":       code.typesCode,
//		"new":         code.newCode,
//		"insert":      code.insertCode,
//		"find":        strings.Join(code.findCode, "\n"),
//		"update":      code.updateCode,
//		"delete":      code.deleteCode,
//		"extraMethod": code.cacheExtra,
//		"tableName":   code.tableName,
//		"data":        table,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return output, nil
////}
//
//func GenModelCustom(in parser.Table, withCache bool) (string, error) {
//	text, err := pathx.LoadTemplate(category, modelCustomTemplateFile, template.ModelCustom)
//	if err != nil {
//		return "", err
//	}
//
//	t := util.With("model-custom").
//		Parse(text).
//		GoFmt(true)
//	output, err := t.Execute(map[string]interface{}{
//		"pkg":                   in.Name.Source(),
//		"withCache":             withCache,
//		"upperStartCamelObject": in.Name.ToCamel(),
//		"lowerStartCamelObject": stringx.From(in.Name.ToCamel()).Untitle(),
//	})
//	if err != nil {
//		return "", err
//	}
//
//	return output.String(), nil
//}

func CreateFile(modelList map[string]*CodeTuple, srcPath string) error {

	dirAbs := filepath.Dir(srcPath)

	err := pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	for tableName, codes := range modelList {
		name := fmt.Sprintf("%v.api", SafeString(tableName))
		filename := filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.Api), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vaddlogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiInsert), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vdellogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiDelete), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vuplogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiUpdate), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vInfologic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiFindOne), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vlistlogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiFindList), os.ModePerm)
		if err != nil {
			return err
		}
	}

	fmt.Println("Done.")
	return nil
}

// SafeString converts the input string into a safe naming style in golang
func SafeString(in string) string {
	if len(in) == 0 {
		return in
	}
	if strings.Contains(in, "_") {
		in = strings.Replace(in, "_", "", -1)
	}
	return in
}
