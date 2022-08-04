package rpc

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"gozc/tools/stringx"
	"strings"
)

func genInsert(table Table, modelName stringx.String) (string, error) {
	datas := make([]string, 0)
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" || camel == "UpdatedName" || camel == "DeletedName" {
			continue
		}
		var model string
		switch camel {
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("%s:\tsql.NullString{String: in.%s, Valid: in.%s != \"\"}, // %s", camel, camel, camel, field.Comment)
			case "sql.NullInt64":
				model = fmt.Sprintf("%s:\tsql.NullInt64{Int64: in.%s, Valid: in.%s != \"\"}, // %s", camel, camel, camel, field.Comment)
			case "sql.NullInt32":
				model = fmt.Sprintf("%s:\tsql.NullInt32{Int32: in.%s, Valid: in.%s != \"\"}, // %s", camel, camel, camel, field.Comment)
			case "sql.NullFloat64":
				model = fmt.Sprintf("%s:\tsql.NullFloat64{Float64: in.%s, Valid: in.%s != \"\"}, // %s", camel, camel, camel, field.Comment)
			case "sql.NullFloat32":
				model = fmt.Sprintf("%s:\tsql.NullFloat32{Float32: in.%s, Valid: in.%s != \"\"}, // %s", camel, camel, camel, field.Comment)
			case "sql.NullTime":
				model = fmt.Sprintf("%s:\tsql.NullTime{Time: in.%s, Valid: in.%s != \"\"}, // %s", camel, camel, camel, field.Comment)
			default:
				model = fmt.Sprintf("%s:\t in.%s, // %s", camel, camel, field.Comment)
			}
		}
		datas = append(datas, model)
	}
	data := strings.Join(datas, "\n\t\t")
	camel := table.Name.ToCamel()
	xmodelname := modelName.Lower()

	text, err := pathx.LoadTemplate(category, insertTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("insert").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":   camel,
			"xmodelname": xmodelname,
			"data":       data,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
