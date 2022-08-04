package rpc

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"gozc/tools/stringx"
	"strings"
)

func genUpdate(table Table, modelName stringx.String) (string, error) {
	datas := make([]string, 0)
	var tenantCount = 0
	var tenantData = ""
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" ||
			camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" {
			continue
		}
		var model string
		switch camel {
		case "TenantId":
			tenantCount++
			continue
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("// %s \n\tif len(in.%v) != 0 {\n\t\tres.%s.String = in.%s\n\t\tres.%s.Valid = true\n\t}", field.Comment, camel, camel, camel, camel)
			case "sql.NullInt64":
				model = fmt.Sprintf("// %s \n\tif in.%v != 0 {\n\t\tres.%s.Int64 = in.%s\n\t\tres.%s.Valid = true\n\t}", field.Comment, camel, camel, camel, camel)
			case "sql.NullInt32":
				model = fmt.Sprintf("// %s \n\tif in.%v != 0 {\n\t\tres.%s.Int32 = in.%s\n\t\tres.%s.Valid = true\n\t}", field.Comment, camel, camel, camel, camel)
			case "sql.NullFloat64":
				model = fmt.Sprintf("// %s \n\tif in.%v != 0.0 {\n\t\tres.%s.Float64 = in.%s\n\t\tres.%s.Valid = true\n\t}", field.Comment, camel, camel, camel, camel)
			case "sql.NullFloat32":
				model = fmt.Sprintf("// %s \n\tif in.%v != 0.0 {\n\t\tres.%s.Float32 = in.%s\n\t\tres.%s.Valid = true\n\t}", field.Comment, camel, camel, camel, camel)
			case "string":
				model = fmt.Sprintf("// %s\n\tif len(in.%v) > 0 {\n\t\tres.%s = in.%s\n\t}", field.Comment, camel, camel, camel)
			case "int64":
				model = fmt.Sprintf("// %s\n\tif in.%v != 0 {\n\t\tres.%s = in.%s\n\t}", field.Comment, camel, camel, camel)
			case "int32":
				model = fmt.Sprintf("// %s\n\tif in.%v != 0 {\n\t\tres.%s = in.%s\n\t}", field.Comment, camel, camel, camel)
			case "float32":
				model = fmt.Sprintf("// %s\n\tif in.%v != 0.0 {\n\t\tres.%s = in.%s\n\t}", field.Comment, camel, camel, camel)
			case "float64":
				model = fmt.Sprintf("// %s\n\tif in.%v != 0.0 {\n\t\tres.%s = in.%s\n\t}", field.Comment, camel, camel, camel)
			case "sql.NullTime":
				continue
			case "time.Time":
				continue
			default:
				continue
			}
		}
		datas = append(datas, model)
	}
	if tenantCount > 0 {
		tenantData = "if res.TenantId != in.TenantId {\n\t\treturn nil, errors.New(\"不是一个租户非法操作\")\n\t}"
	}
	data := strings.Join(datas, "\n\t")
	camel := table.Name.ToCamel()
	xmodelname := modelName.Lower()
	text, err := pathx.LoadTemplate(category, updateTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("update").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":   camel,
			"xmodelname": xmodelname,
			"tenant":     tenantData,
			"updateData": data,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
