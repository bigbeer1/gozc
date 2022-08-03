package gen

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"strings"
)

func genFindList(table Table, pkgName string) (string, error) {
	datas := make([]string, 0)
	modeldatas := make([]string, 0)
	var initmodel string

	//添加分页
	initmodel = fmt.Sprintf("%s:\t req.%s, // %s", "Current", "Current", "页码")
	datas = append(datas, initmodel)
	initmodel = fmt.Sprintf("%s:\t req.%s, // %s", "PageSize", "PageSize", "页数")
	datas = append(datas, initmodel)

	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" || camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" {
			continue
		}
		var model string
		switch camel {
		case "TenantId":
			model = fmt.Sprintf("%s:\t tokenData.%s, // %s", camel, camel, field.Comment)
		default:
			model = fmt.Sprintf("%s:\t req.%s, // %s", camel, camel, field.Comment)
		}
		datas = append(datas, model)
	}

	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if camel == "DeletedAt" || camel == "DeletedName" {
			continue
		}
		var model string
		switch camel {
		case "TenantId":
			continue
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "string", field.Name.Source(), field.Comment)
			case "string":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "string", field.Name.Source(), field.Comment)
			case "sql.NullTime":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "int64", field.Name.Source(), field.Comment)
			case "time.Time":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "int64", field.Name.Source(), field.Comment)
			case "sql.NullInt64":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "int64", field.Name.Source(), field.Comment)
			case "sql.NullInt32":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "int64", field.Name.Source(), field.Comment)
			case "int64":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "int64", field.Name.Source(), field.Comment)
			default:
				continue
			}
		}
		modeldatas = append(modeldatas, model)

	}

	data := strings.Join(datas, "\n\t\t")
	modeldata := strings.Join(modeldatas, ",\n\t")
	camel := table.Name.ToCamel()
	text, err := pathx.LoadTemplate(category, findlistTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("insert").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":  camel,
			"modelname": pkgName,
			"data":      data,
			"modeldata": modeldata,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
