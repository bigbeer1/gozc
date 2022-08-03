package gen

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"strings"
)

func genApi(table Table, pkgName string) (string, error) {

	camel := table.Name.ToCamel()

	add := GetData(table, insertTemplateFile)
	del := GetData(table, deleteTemplateFile)
	up := GetData(table, updateTemplateFile)
	list := GetData(table, findlistTemplateFile)
	info := GetData(table, findOneTemplateFile)

	text, err := pathx.LoadTemplate(category, apiTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("insert").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":  camel,
			"modelname": pkgName,
			"Add":       add,
			"Del":       del,
			"Up":        up,
			"List":      list,
			"Info":      info,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func GetData(table Table, dataType string) string {
	modeldatas := make([]string, 0)
	var initmodel string
	var reqType = "json"
	var reqTypeData string
	if dataType == findlistTemplateFile {
		//添加分页
		initmodel = fmt.Sprintf("%s  %s  `form:\"%s\"`  // %s", "Current", "int64", "current,default=1,optional", "页码")
		modeldatas = append(modeldatas, initmodel)
		initmodel = fmt.Sprintf("%s  %s  `form:\"%s\"`  // %s", "PageSize", "int64", "page_size,default=10,optional", "页数")
		modeldatas = append(modeldatas, initmodel)
		reqType = "form"
	}

	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		switch dataType {
		case findlistTemplateFile:
			if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" || camel == "TenantId" {
				continue
			}
			reqTypeData = field.Name.Source() + ",optional"
		case insertTemplateFile:
			if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" || camel == "TenantId" {
				continue
			}
			reqTypeData = field.Name.Source() + ",optional"
		case deleteTemplateFile:
			if camel != "Id" {
				continue
			}
			reqTypeData = field.Name.Source()
		case updateTemplateFile:
			if camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" || camel == "TenantId" {
				continue
			}
			reqTypeData = field.Name.Source() + ",optional"
		case findOneTemplateFile:
			if camel != "Id" {
				continue
			}
			reqTypeData = field.Name.Source()
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
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "int64", reqType, reqTypeData, field.Comment)
			case "sql.NullInt32":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "int64", reqType, reqTypeData, field.Comment)
			case "int64":
				model = fmt.Sprintf("%s  %s  `%s:\"%s\"`  // %s", camel, "int64", reqType, reqTypeData, field.Comment)
			default:
				continue
			}
		}
		modeldatas = append(modeldatas, model)

	}
	modeldata := strings.Join(modeldatas, "\n\t\t")
	return modeldata

}
