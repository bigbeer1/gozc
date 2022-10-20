package api

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"gozc/tools/stringx"
	"strings"
)

func genApi(table Table, pkgName stringx.String) (string, error) {

	camel := table.Name.ToCamel()
	xcamel := table.Name.ToCamelWithStartLower()

	var modelname = pkgName.Lower()
	var amodelname = pkgName.ToCamel()

	add := GetApiData(table, insertTemplateFile)
	del := GetApiData(table, deleteTemplateFile)
	up := GetApiData(table, updateTemplateFile)
	list := GetApiData(table, findListTemplateFile)
	info := GetApiData(table, findOneTemplateFile)

	text, err := pathx.LoadTemplate(category, apiTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("api").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":   camel,
			"xfilename":  xcamel,
			"modelname":  modelname,
			"Amodelname": amodelname,
			"Add":        add,
			"Del":        del,
			"Up":         up,
			"List":       list,
			"Info":       info,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func GetApiData(table Table, dataType string) string {
	modeldatas := make([]string, 0)
	var initmodel string
	var reqType = "json"
	var reqTypeData string
	var reqTypeDataInt string
	if dataType == findListTemplateFile {
		//添加分页
		initmodel = fmt.Sprintf("%s  %s  `form:\"%s\"`  // %s", "Current", "int64", "current,default=1,optional", "页码")
		modeldatas = append(modeldatas, initmodel)
		initmodel = fmt.Sprintf("%s  %s  `form:\"%s\"`  // %s", "PageSize", "int64", "page_size,default=10,optional", "页数")
		modeldatas = append(modeldatas, initmodel)
		reqType = "form"
	}
	if dataType == deleteTemplateFile {
		reqType = "path"
	}
	if dataType == findOneTemplateFile {
		reqType = "form"
	}

	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		switch dataType {
		case findListTemplateFile:
			if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" || camel == "TenantId" || camel == "Sort" {
				continue
			}
			reqTypeData = field.Name.Source() + ",optional"
			reqTypeDataInt = field.Name.Source() + ",default=99,optional"
		case insertTemplateFile:
			if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" || camel == "TenantId" {
				continue
			}
			reqTypeData = field.Name.Source() + ",optional"
			reqTypeDataInt = field.Name.Source() + ",optional"
		case deleteTemplateFile:
			if camel != "Id" {
				continue
			}
			reqTypeData = field.Name.Source()
			reqTypeDataInt = field.Name.Source()

		case updateTemplateFile:
			if camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" || camel == "TenantId" {
				continue
			}
			if camel != "Id" {
				reqTypeData = field.Name.Source() + ",optional"
				reqTypeDataInt = field.Name.Source() + ",optional"
			} else {
				reqTypeData = field.Name.Source()
				reqTypeDataInt = field.Name.Source()
			}
		case findOneTemplateFile:
			if camel != "Id" {
				continue
			}
			reqTypeData = field.Name.Source()
			reqTypeDataInt = field.Name.Source()
		}
		var model string
		switch camel {
		case "TenantId":
			continue
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "string", reqType, reqTypeData, field.Comment)
			case "string":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "string", reqType, reqTypeData, field.Comment)
			case "sql.NullTime":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "int64", reqType, reqTypeData, field.Comment)
			case "time.Time":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "int64", reqType, reqTypeData, field.Comment)
			case "sql.NullInt64":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "int64", reqType, reqTypeDataInt, field.Comment)
			case "sql.NullInt32":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "int64", reqType, reqTypeDataInt, field.Comment)
			case "int64":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "int64", reqType, reqTypeDataInt, field.Comment)
			case "int32":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "int64", reqType, reqTypeDataInt, field.Comment)
			case "float64":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "float64", reqType, reqTypeData, field.Comment)
			case "float32":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "float32", reqType, reqTypeData, field.Comment)
			case "NullFloat32":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "float32", reqType, reqTypeData, field.Comment)
			case "NullFloat64":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "float64", reqType, reqTypeData, field.Comment)
			default:
				continue
			}
		}
		modeldatas = append(modeldatas, model)

	}
	modeldata := strings.Join(modeldatas, "\n\t\t")
	return modeldata

}
