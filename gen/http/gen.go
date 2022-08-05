package http

import (
	"fmt"
	"gozc/parser"
	"gozc/tools/stringx"
)

type Table struct {
	parser.Table
}

type (
	CodeTuple struct {
		Api string
	}
)

const (
	category        = "Http"
	apiTemplateFile = "http-api.tpl"
)

func GenHttpModel(in parser.Table, pkgName string, apiType string) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}
	var table Table
	table.Table = in
	var modelName stringx.String
	modelName = stringx.From(pkgName)

	switch apiType {
	case "http-api":
		res, err := genHttpApi(table, modelName)
		if err != nil {
			return "", err
		}
		return res, nil

	}
	return "", nil

}
