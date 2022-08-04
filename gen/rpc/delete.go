package rpc

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"gozc/tools/stringx"
)

func genDelete(table Table, modelName stringx.String) (string, error) {
	var deletedCount = 0
	var deletedData = ""
	var deletedAtData = ""
	var tenantCount = 0
	var tenantData = ""
	var delType = "delete(id)"
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		switch camel {
		case "DeletedAt":
			deletedCount++
		case "TenantId":
			tenantCount++
		}
	}
	if tenantCount > 0 {
		tenantData = "if res.TenantId != in.TenantId {\n\t\treturn nil, errors.New(\"不是一个租户非法操作\")\n\t}"
	}

	camel := table.Name.ToCamel()

	if deletedCount > 0 {
		deletedData = "res.DeletedAt.Time = time.Now()\n\tres.DeletedAt.Valid = true\n\tres.DeletedName.String = in.DeletedName\n\tres.DeletedName.Valid = true"
		delType = "update(res)"
		deletedAtData = fmt.Sprintf("// 判断该数据是否被删除\n\tif res.DeletedAt.Valid == true {\n\t\treturn nil, errors.New(\"%s该ID已被删除：\" + in.Id)\n\t}", camel)
	}

	xmodelname := modelName.Lower()
	text, err := pathx.LoadTemplate(category, deleteTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("delete").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":      camel,
			"xmodelname":    xmodelname,
			"deletedAtData": deletedAtData,
			"tenant":        tenantData,
			"del":           deletedData,
			"delType":       delType,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
