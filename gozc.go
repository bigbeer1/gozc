package main

import (
	"flag"
	"fmt"
	"gozc/gen"
	"gozc/parser"
)

func main() {

	flag.Parse()
	srcPath := "E:\\Gopath\\src\\sql2pb-main\\model\\sql\\sys_admin.sql"
	// 讲sql文件 转换成tables
	tables, err := parser.Parse(srcPath)
	if err != nil {
		fmt.Println(err)
	}

	m := make(map[string]*gen.CodeTuple)

	for _, e := range tables {
		Api, err := gen.GenApiModel(*e, "admin", "api")
		if err != nil {
			fmt.Println(err)
		}
		ApiInsert, err := gen.GenApiModel(*e, "admin", "Insert")
		if err != nil {
			fmt.Println(err)
		}
		ApiDelete, err := gen.GenApiModel(*e, "admin", "delete")
		if err != nil {
			fmt.Println(err)
		}
		ApiUpdate, err := gen.GenApiModel(*e, "admin", "update")
		if err != nil {
			fmt.Println(err)
		}
		ApiFindOne, err := gen.GenApiModel(*e, "admin", "findOne")
		if err != nil {
			fmt.Println(err)
		}
		ApiFindList, err := gen.GenApiModel(*e, "admin", "findList")
		if err != nil {
			fmt.Println(err)
		}

		m[e.Name.Source()] = &gen.CodeTuple{
			Api:         Api,
			ApiInsert:   ApiInsert,
			ApiDelete:   ApiDelete,
			ApiUpdate:   ApiUpdate,
			ApiFindOne:  ApiFindOne,
			ApiFindList: ApiFindList,
		}
	}

	cc := gen.CreateFile(m, srcPath)
	fmt.Println(cc)

}
