package rpc

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"gozc/tools/stringx"
	"strings"
)

func genFindOne(table Table, modelName stringx.String) (string, error) {
	datas := make([]string, 0)
	var tenantCount = 0
	var tenantData = ""
	var deletedCount = 0
	var deletedAtData = ""
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if camel == "DeletedName" {
			continue
		}
		var model string
		switch camel {
		case "TenantId":
			tenantCount++
			continue
		case "DeletedAt":
			deletedCount++
			continue
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("%s:\tres.%s.String, //%s", camel, camel, field.Comment)
			case "sql.NullInt64":
				model = fmt.Sprintf("%s:\tres.%s.Int64, //%s", camel, camel, field.Comment)
			case "sql.NullInt32":
				model = fmt.Sprintf("%s:\tres.%s.Int32, //%s", camel, camel, field.Comment)
			case "sql.NullFloat64":
				model = fmt.Sprintf("%s:\tres.%s.Float64, //%s", camel, camel, field.Comment)
			case "sql.NullFloat32":
				model = fmt.Sprintf("%s:\tres.%s.Float32, //%s", camel, camel, field.Comment)
			case "sql.NullTime":
				model = fmt.Sprintf("%s:\tres.%s.Time.UnixMilli(), //%s", camel, camel, field.Comment)
			case "time.Time":
				model = fmt.Sprintf("%s:\tres.%s.UnixMilli(), //%s", camel, camel, field.Comment)
			default:
				model = fmt.Sprintf("%s:\tres.%s, //%s", camel, camel, field.Comment)
			}
		}

		datas = append(datas, model)
	}
	if tenantCount > 0 {
		tenantData = "if res.TenantId != in.TenantId {\n\t\treturn nil, errors.New(\"不是一个租户非法操作\")\n\t}"
	}

	camel := table.Name.ToCamel()

	if deletedCount > 0 {
		deletedAtData = fmt.Sprintf("// 判断该数据是否被删除\n\tif res.DeletedAt.Valid == true {\n\t\treturn nil, fmt.Errorf(\"%s该ID已被删除：%s\",in.Id)\n\t}", camel, "%v")
	}

	data := strings.Join(datas, "\n\t\t")

	xmodelname := modelName.Lower()
	text, err := pathx.LoadTemplate(category, findOneTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("findOne").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":      camel,
			"xmodelname":    xmodelname,
			"deletedAtData": deletedAtData,
			"tenant":        tenantData,
			"findoneData":   data,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
