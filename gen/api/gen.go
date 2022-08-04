package api

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
	findListTemplateFile = "find-list.tpl"
)

func GenApiModel(in parser.Table, pkgName string, apiType string) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}
	var table Table
	table.Table = in

	var modelName stringx.String
	modelName = stringx.From(pkgName)

	switch apiType {
	case "api":
		res, err := genApi(table, modelName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "Insert":
		res, err := genInsert(table, modelName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "delete":
		res, err := genDelete(table, modelName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "update":
		res, err := genUpdate(table, modelName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "findOne":
		res, err := genFindOne(table, modelName)
		if err != nil {
			return "", err
		}
		return res, nil
	case "findList":
		res, err := genFindList(table, modelName)
		if err != nil {
			return "", err
		}
		return res, nil
	}
	return "", nil
}
