package rpc

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
		Rpc         string
		RpcInsert   string
		RpcDelete   string
		RpcUpdate   string
		RpcFindOne  string
		RpcFindList string
	}
)

const (
	category                 = "rpc"
	rpcTemplateFile          = "rpc.tpl"
	insertTemplateFile       = "insert.tpl"
	deleteTemplateFile       = "delete.tpl"
	updateTemplateFile       = "update.tpl"
	findOneTemplateFile      = "find-one.tpl"
	findOneRespTemplateFile  = "find-one-resp"
	findListTemplateFile     = "find-list.tpl"
	findListRespTemplateFile = "find-list-resp"
	findDataTemplateFile     = "find-list-data"
)

func GenRpcModel(in parser.Table, pkgName string, apiType string) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}
	var table Table
	table.Table = in
	var modelName stringx.String
	modelName = stringx.From(pkgName)

	switch apiType {
	case "rpc":
		res, err := genRpc(table, modelName)
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
