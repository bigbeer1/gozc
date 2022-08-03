package gen

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"strings"
)

func genUpdate(table Table, pkgName string) (string, error) {
	datas := make([]string, 0)
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if camel == "CreatedAt" || camel == "UpdatedAt" || camel == "DeletedAt" || camel == "CreatedName" || camel == "DeletedName" {
			continue
		}
		var model string
		switch camel {
		case "UpdatedName":
			model = fmt.Sprintf("%s:\t tokenData.%s, // %s", camel, camel, field.Comment)
		case "TenantId":
			model = fmt.Sprintf("%s:\t tokenData.%s, // %s", camel, camel, field.Comment)
		default:
			model = fmt.Sprintf("%s:\t req.%s, // %s", camel, camel, field.Comment)
		}
		datas = append(datas, model)
	}
	data := strings.Join(datas, "\n\t\t")
	camel := table.Name.ToCamel()
	text, err := pathx.LoadTemplate(category, updateTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("insert").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":  camel,
			"modelname": pkgName,
			"data":      data,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
