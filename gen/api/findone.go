package api

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"gozc/tools/stringx"
	"strings"
)

func genFindOne(table Table, modelName stringx.String) (string, error) {
	datas := make([]string, 0)
	modeldatas := make([]string, 0)
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		var model string
		switch camel {
		case "TenantId":
			model = fmt.Sprintf("%s:\t tokenData.%s, // %s", camel, camel, field.Comment)
		case "Id":
			model = fmt.Sprintf("%s:\t req.%s, // %s", camel, camel, field.Comment)
		default:
			continue
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
			case "int32":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "int64", field.Name.Source(), field.Comment)
			case "float64":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "float64", field.Name.Source(), field.Comment)
			case "float32":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "float32", field.Name.Source(), field.Comment)
			case "NullFloat32":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "float32", field.Name.Source(), field.Comment)
			case "NullFloat64":
				model = fmt.Sprintf("%s  %s  `json:\"%s\"`  // %s", camel, "float64", field.Name.Source(), field.Comment)
			default:
				continue
			}
		}
		modeldatas = append(modeldatas, model)

	}

	data := strings.Join(datas, "\n\t\t")
	modeldata := strings.Join(modeldatas, ",\n\t")
	camel := table.Name.ToCamel()
	amodelname := modelName.ToCamel()
	xmodelname := modelName.Lower()
	text, err := pathx.LoadTemplate(category, findOneTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("findOne").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":   camel,
			"modelname":  amodelname,
			"xmodelname": xmodelname,
			"data":       data,
			"modeldata":  modeldata,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
