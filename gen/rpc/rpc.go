package rpc

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"gozc/tools/stringx"
	"strings"
)

func genRpc(table Table, pkgName stringx.String) (string, error) {

	camel := table.Name.ToCamel()

	var modelname = pkgName.Lower()
	var amodelname = pkgName.ToCamel()

	add := GetRpcData(table, insertTemplateFile)
	del := GetRpcData(table, deleteTemplateFile)
	up := GetRpcData(table, updateTemplateFile)
	findOne := GetRpcData(table, findOneTemplateFile)
	findOneResp := GetRpcData(table, findOneRespTemplateFile)
	list := GetRpcData(table, findListTemplateFile)
	listResp := GetRpcData(table, findListRespTemplateFile)
	listData := GetRpcData(table, findDataTemplateFile)
	text, err := pathx.LoadTemplate(category, rpcTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("rpc").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":    camel,
			"modelname":   modelname,
			"amodelname":  amodelname,
			"Add":         add,
			"Del":         del,
			"Up":          up,
			"FindOne":     findOne,
			"FindOneResp": findOneResp,
			"List":        list,
			"ListResp":    listResp,
			"ListData":    listData,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func GetRpcData(table Table, dataType string) string {
	modeldatas := make([]string, 0)
	var count = 0
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		lowerCamel := util.SafeString(field.Name.Lower())
		switch dataType {
		case insertTemplateFile:
			if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "UpdatedName" || camel == "DeletedName" {
				continue
			}
		case deleteTemplateFile:
			if camel != "Id" && camel != "DeletedName" && camel != "TenantId" {
				continue
			}
		case updateTemplateFile:
			if camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "DeletedName" {
				continue
			}
		case findOneTemplateFile:
			if camel != "Id" && camel != "TenantId" {
				continue
			}
		case findOneRespTemplateFile:
			if camel == "DeletedAt" ||
				camel == "DeletedName" || camel == "TenantId" {
				continue
			}
		case findListTemplateFile:
			if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" {
				continue
			}
		case findDataTemplateFile:
			if camel == "DeletedAt" ||
				camel == "DeletedName" || camel == "TenantId" {
				continue
			}
		case findListRespTemplateFile:
			var model string
			model = fmt.Sprintf("%s  %s = %v;  // %s", "int64", "total", 1, "总数")
			modeldatas = append(modeldatas, model)
			model = fmt.Sprintf("%s  %s = %v;  // %s", "repeated", table.Name.ToCamel()+"ListData list", 2, "内容")
			modeldatas = append(modeldatas, model)
			modeldata := strings.Join(modeldatas, "\n\t")
			return modeldata
		default:
			return ""
		}

		count++
		var model string
		switch camel {
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("%s  %s = %v;  // %s", "string", lowerCamel, count, field.Comment)
			case "string":
				model = fmt.Sprintf("%s  %s = %v;  // %s", "string", lowerCamel, count, field.Comment)
			case "sql.NullTime":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "int64", lowerCamel, count, field.Comment)
			case "time.Time":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "int64", lowerCamel, count, field.Comment)
			case "sql.NullInt64":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "int64", lowerCamel, count, field.Comment)
			case "sql.NullInt32":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "int64", lowerCamel, count, field.Comment)
			case "int64":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "int64", lowerCamel, count, field.Comment)
			case "int32":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "int64", lowerCamel, count, field.Comment)
			case "float64":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "double", lowerCamel, count, field.Comment)
			case "float32":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "float", lowerCamel, count, field.Comment)
			case "sql.NullFloat64":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "double", lowerCamel, count, field.Comment)
			case "sql.NullFloat32":
				model = fmt.Sprintf("%s   %s = %v;  // %s", "float", lowerCamel, count, field.Comment)
			default:
				continue
			}
		}
		modeldatas = append(modeldatas, model)

	}
	modeldata := strings.Join(modeldatas, "\n\t")
	return modeldata

}
