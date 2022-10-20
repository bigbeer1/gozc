package rpc

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gozc/tools/pathx"
	"gozc/tools/stringx"
	"strings"
)

func genFindList(table Table, modelName stringx.String) (string, error) {

	var tenantCount = 0
	var tenantData = ""
	var tenantDataCount = ""

	var deletedCount = 0
	var deletedData = ""
	var deletedDataCount = ""

	whereBuilderData, tenantCount, deletedCount := getListData(table, "whereBuilder")
	countBuilderData, tenantCount, deletedCount := getListData(table, "countBuilder")

	findListData := getFindListData(table)

	if tenantCount > 0 {
		tenantData = "whereBuilder = whereBuilder.Where(squirrel.Eq{\n  \t\t\"tenant_id\":     in.TenantId,\n  \t})"
		tenantDataCount = "countBuilder = countBuilder.Where(squirrel.Eq{\n    \t\"tenant_id\": in.TenantId,\n  \t})"
	}

	if deletedCount > 0 {
		deletedData = "whereBuilder = whereBuilder.Where(\"deleted_at is null\")"
		deletedDataCount = "countBuilder = countBuilder.Where(\"deleted_at is null\")"
	}
	createdData := "whereBuilder = whereBuilder.OrderBy(\"created_at DESC\")"

	camel := table.Name.ToCamel()
	xmodelname := modelName.Lower()
	text, err := pathx.LoadTemplate(category, findListTemplateFile, "")
	if err != nil {
		return "", err
	}
	output, err := util.With("findOne").
		Parse(text).
		Execute(map[string]interface{}{
			"filename":       camel,
			"xmodelname":     xmodelname,
			"listDeletedAt":  deletedData,
			"listCreatedAt":  createdData,
			"listTenant":     tenantData,
			"listData":       whereBuilderData,
			"countDeletedAt": deletedDataCount,
			"countTenant":    tenantDataCount,
			"countData":      countBuilderData,
			"findlistData":   findListData,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func getListData(table Table, builderName string) (data string, tenantCount, deletedCount int) {
	datas := make([]string, 0)
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		xcamel := util.SafeString(field.Name.Lower())
		if camel == "Id" || camel == "CreatedAt" || camel == "UpdatedAt" || camel == "CreatedName" || camel == "UpdatedName" || camel == "DeletedName" {
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
		case "Sort":
			continue
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("// %s\n\tif len(in.%v) > 0 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "sql.NullInt64":
				model = fmt.Sprintf("// %s\n\tif in.%v != 99 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "sql.NullInt32":
				model = fmt.Sprintf("// %s\n\tif in.%v != 99 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "sql.NullFloat64":
				model = fmt.Sprintf("// %s\n\tif in.%v != 99.0 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "sql.NullFloat32":
				model = fmt.Sprintf("// %s\n\tif in.%v != 99.0 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "string":
				model = fmt.Sprintf("// %s\n\tif len(in.%v) > 0 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "int64":
				model = fmt.Sprintf("// %s\n\tif in.%v != 99 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "int32":
				model = fmt.Sprintf("// %s\n\tif in.%v != 99 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "float32":
				model = fmt.Sprintf("// %s\n\tif in.%v != 99.0 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "float64":
				model = fmt.Sprintf("// %s\n\tif in.%v != 99.0 {\n\t\t%s = %s.Where(squirrel.Eq{\n\t\t\t\"%s \": in.%s,\n\t\t})\n\t}", field.Comment, camel, builderName, builderName, xcamel, camel)
			case "sql.NullTime":
				continue
			case "time.Time":
				continue
			}
		}
		datas = append(datas, model)
	}
	data = strings.Join(datas, "\n\t")
	return data, tenantCount, deletedCount
}

func getFindListData(table Table) string {
	datas := make([]string, 0)
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if camel == "DeletedName" {
			continue
		}
		var model string
		switch camel {
		case "TenantId":
			continue
		case "DeletedAt":
			continue
		default:
			switch field.DataType {
			case "sql.NullString":
				model = fmt.Sprintf("%s:\titem.%s.String, //%s", camel, camel, field.Comment)
			case "sql.NullInt64":
				model = fmt.Sprintf("%s:\titem.%s.Int64, //%s", camel, camel, field.Comment)
			case "sql.NullInt32":
				model = fmt.Sprintf("%s:\titem.%s.Int32, //%s", camel, camel, field.Comment)
			case "sql.NullFloat64":
				model = fmt.Sprintf("%s:\titem.%s.Float64, //%s", camel, camel, field.Comment)
			case "sql.NullFloat32":
				model = fmt.Sprintf("%s:\titem.%s.Float32, //%s", camel, camel, field.Comment)
			case "sql.NullTime":
				model = fmt.Sprintf("%s:\titem.%s.Time.UnixMilli(), //%s", camel, camel, field.Comment)
			case "time.Time":
				model = fmt.Sprintf("%s:\titem.%s.UnixMilli(), //%s", camel, camel, field.Comment)
			default:
				model = fmt.Sprintf("%s:\titem.%s, //%s", camel, camel, field.Comment)
			}
		}

		datas = append(datas, model)
	}

	data := strings.Join(datas, "\n\t\t\t")

	return data

}
