package parser

import (
	"fmt"
	"github.com/zeromicro/ddl-parser/console"
	"github.com/zeromicro/ddl-parser/parser"
	"gozc/converter"

	"github.com/zeromicro/go-zero/core/collection"
	"gozc/tools/stringx"
	"path/filepath"
	"strings"
)

// Parse parses ddl into golang structure
func Parse(filename string) ([]*Table, error) {
	p := parser.NewParser()
	tables, err := p.From(filename)
	if err != nil {
		return nil, err
	}

	nameOriginals := parseNameOriginal(tables)
	indexNameGen := func(column ...string) string {
		return strings.Join(column, "_")
	}

	prefix := filepath.Base(filename)
	var list []*Table
	for indexTable, e := range tables {
		var (
			primaryColumn    string
			primaryColumnSet = collection.NewSet()
			uniqueKeyMap     = make(map[string][]string)
			normalKeyMap     = make(map[string][]string)
			columns          = e.Columns
		)

		for _, column := range columns {
			if column.Constraint != nil {
				if column.Constraint.Primary {
					primaryColumnSet.AddStr(column.Name)
				}

				if column.Constraint.Unique {
					indexName := indexNameGen(column.Name, "unique")
					uniqueKeyMap[indexName] = []string{column.Name}
				}

				if column.Constraint.Key {
					indexName := indexNameGen(column.Name, "idx")
					uniqueKeyMap[indexName] = []string{column.Name}
				}
			}
		}

		for _, e := range e.Constraints {
			if len(e.ColumnPrimaryKey) > 1 {
				return nil, fmt.Errorf("%s: unexpected join primary key", prefix)
			}

			if len(e.ColumnPrimaryKey) == 1 {
				primaryColumn = e.ColumnPrimaryKey[0]
				primaryColumnSet.AddStr(e.ColumnPrimaryKey[0])
			}

			if len(e.ColumnUniqueKey) > 0 {
				list := append([]string(nil), e.ColumnUniqueKey...)
				list = append(list, "unique")
				indexName := indexNameGen(list...)
				uniqueKeyMap[indexName] = e.ColumnUniqueKey
			}
		}

		if primaryColumnSet.Count() > 1 {
			return nil, fmt.Errorf("%s: unexpected join primary key", prefix)
		}

		primaryKey, fieldM, err := convertColumns(columns, primaryColumn)
		if err != nil {
			return nil, err
		}

		var fields []*Field
		// sort
		for indexColumn, c := range columns {
			field, ok := fieldM[c.Name]
			if ok {
				field.NameOriginal = nameOriginals[indexTable][indexColumn]
				fields = append(fields, field)
			}
		}

		var (
			uniqueIndex = make(map[string][]*Field)
			normalIndex = make(map[string][]*Field)
		)

		for indexName, each := range uniqueKeyMap {
			for _, columnName := range each {
				uniqueIndex[indexName] = append(uniqueIndex[indexName], fieldM[columnName])
			}
		}

		for indexName, each := range normalKeyMap {
			for _, columnName := range each {
				normalIndex[indexName] = append(normalIndex[indexName], fieldM[columnName])
			}
		}

		checkDuplicateUniqueIndex(uniqueIndex, e.Name)

		list = append(list, &Table{
			Name:        stringx.From(e.Name),
			PrimaryKey:  primaryKey,
			UniqueIndex: uniqueIndex,
			Fields:      fields,
		})
	}

	return list, nil
}

func parseNameOriginal(ts []*parser.Table) (nameOriginals [][]string) {
	var columns []string

	for _, t := range ts {
		columns = []string{}
		for _, c := range t.Columns {
			columns = append(columns, c.Name)
		}
		nameOriginals = append(nameOriginals, columns)
	}
	return
}

type (
	Table struct {
		Name        stringx.String
		Db          stringx.String
		PrimaryKey  Primary
		UniqueIndex map[string][]*Field
		Fields      []*Field
	}

	Primary struct {
		Field
		AutoIncrement bool
	}

	Field struct {
		NameOriginal    string
		Name            stringx.String
		DataType        string
		Comment         string
		SeqInIndex      int
		OrdinalPosition int
	}
)

func convertColumns(columns []*parser.Column, primaryColumn string) (Primary, map[string]*Field, error) {
	var (
		primaryKey Primary
		fieldM     = make(map[string]*Field)
		log        = console.NewColorConsole()
	)

	for _, column := range columns {
		if column == nil {
			continue
		}

		var (
			comment       string
			isDefaultNull bool
		)

		if column.Constraint != nil {
			comment = column.Constraint.Comment
			isDefaultNull = !column.Constraint.NotNull
			if !column.Constraint.NotNull && column.Constraint.HasDefaultValue {
				isDefaultNull = false
			}

			if column.Name == primaryColumn {
				isDefaultNull = false
			}
		}

		dataType, err := converter.ConvertDataType(column.DataType.Type(), isDefaultNull, column.DataType.Unsigned())
		if err != nil {
			return Primary{}, nil, err
		}

		if column.Constraint != nil {
			if column.Name == primaryColumn {
				if !column.Constraint.AutoIncrement && dataType == "int64" {
					log.Warning("[convertColumns]: The primary key %q is recommended to add constraint `AUTO_INCREMENT`", column.Name)
				}
			} else if column.Constraint.NotNull && !column.Constraint.HasDefaultValue {
				log.Warning("[convertColumns]: The column %q is recommended to add constraint `DEFAULT`", column.Name)
			}
		}

		var field Field
		field.Name = stringx.From(column.Name)
		field.DataType = dataType
		field.Comment = stringx.TrimNewLine(comment)

		if field.Name.Source() == primaryColumn {
			primaryKey = Primary{
				Field: field,
			}
			if column.Constraint != nil {
				primaryKey.AutoIncrement = column.Constraint.AutoIncrement
			}
		}

		fieldM[field.Name.Source()] = &field
	}
	return primaryKey, fieldM, nil
}

func checkDuplicateUniqueIndex(uniqueIndex map[string][]*Field, tableName string) {
	log := console.NewColorConsole()
	uniqueSet := collection.NewSet()
	for k, i := range uniqueIndex {
		var list []string
		for _, e := range i {
			list = append(list, e.Name.Source())
		}

		joinRet := strings.Join(list, ",")
		if uniqueSet.Contains(joinRet) {
			log.Warning("[checkDuplicateUniqueIndex]: table %s: duplicate unique index %s", tableName, joinRet)
			delete(uniqueIndex, k)
			continue
		}

		uniqueSet.AddStr(joinRet)
	}
}

const timeImport = "time.Time"

func (t *Table) ContainsTime() bool {
	for _, item := range t.Fields {
		if item.DataType == timeImport {
			return true
		}
	}
	return false
}
