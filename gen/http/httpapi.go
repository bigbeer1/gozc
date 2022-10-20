package http

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"gozc/tools/stringx"
	"strings"
)

func genHttpApi(table Table, pkgName stringx.String) (string, error) {

	xcamel := table.Name.ToCamelWithStartLower()
	var modelname = pkgName.Lower()
	var idType, idCommand string
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if camel != "Id" {
			continue
		}
		idType = field.DataType
		idCommand = field.Comment
	}
	text, err := pathx.LoadTemplate(category, apiTemplateFile, "")
	if err != nil {
		return "", err
	}

	add := GetHttpData(table, "add")
	update := GetHttpData(table, "update")

	query := GetHttpQueryData(table)
	output, err := util.With("http-api").
		Parse(text).
		Execute(map[string]interface{}{
			"modelname": modelname,
			"xfilename": xcamel,
			"idCommand": idCommand,
			"idType":    idType,
			"Add":       add,
			"Update":    update,
			"Query":     query,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func GetHttpData(table Table, dataType string) string {
	modeldatas := make([]string, 0)

	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		xcamel := util.SafeString(field.Name.Lower())
		switch dataType {
		case "add":
			if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" || camel == "TenantId" {
				continue
			}
		case "Id":
			if camel != "Id" {
				continue
			}
			model := fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "string", field.Comment)
			modeldatas = append(modeldatas, model)
			modeldata := strings.Join(modeldatas, ",\n\t\t\t\t")
			return modeldata
		case "update":
			if camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
				camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" || camel == "TenantId" {
				continue
			}
		default:
			continue
		}
		var model string
		switch camel {
		case "TenantId":
			continue
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "string", field.Comment)
			case "string":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "string", field.Comment)
			case "sql.NullTime":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "integer", field.Comment)
			case "time.Time":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "integer", field.Comment)
			case "sql.NullInt64":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "integer", field.Comment)
			case "sql.NullInt32":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "integer", field.Comment)
			case "int64":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "integer", field.Comment)
			case "int32":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "integer", field.Comment)
			case "float64":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "number", field.Comment)
			case "float32":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "number", field.Comment)
			case "NullFloat32":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "number", field.Comment)
			case "NullFloat64":
				model = fmt.Sprintf("  \"%s\": {\n\t\t\t\t\t\"type\": \"%s\",\n\t\t\t\t\t\"description\": \"%s\"\n\t\t\t\t  }", xcamel, "number", field.Comment)
			default:
				continue
			}
		}
		modeldatas = append(modeldatas, model)

	}
	modeldata := strings.Join(modeldatas, ",\n\t\t\t\t")
	return modeldata

}

func GetHttpQueryData(table Table) string {
	modeldatas := make([]string, 0)
	var initmodel string
	// 添加分页
	initmodel = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
		" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", "current", "页码", "integer")
	modeldatas = append(modeldatas, initmodel)
	initmodel = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
		" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", "page_size", "页数", "integer")
	modeldatas = append(modeldatas, initmodel)

	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		xcamel := util.SafeString(field.Name.Lower())
		if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
			camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" || camel == "TenantId" || camel == "Sort" {
			continue
		}

		var model string
		switch camel {
		case "TenantId":
			continue
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "string")
			case "string":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "string")
			case "sql.NullTime":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "integer")
			case "time.Time":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "integer")
			case "sql.NullInt64":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "integer")
			case "sql.NullInt32":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "integer")
			case "int64":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "integer")
			case "int32":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "integer")
			case "float64":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "number")
			case "float32":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "number")
			case "NullFloat32":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "number")
			case "NullFloat64":
				model = fmt.Sprintf("{\n\t\t\t\"name\": \"%s\",\n\t\t\t\"in\": \"query\",\n\t\t\t\"description\":"+
					" \"%s\",\n\t\t\t\"required\": false,\n\t\t\t\"schema\": {\n\t\t\t  \"type\": \"%s\"\n\t\t\t}\n\t\t  }", xcamel, field.Comment, "number")
			default:
				continue
			}
		}
		modeldatas = append(modeldatas, model)

	}
	modeldata := strings.Join(modeldatas, ",\n\t\t  ")
	return modeldata

}
