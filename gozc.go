package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/urfave/cli/v2"
	"gozc/gen/api"
	"gozc/gen/rpc"
	"gozc/parser"
	"os"
	"strings"
)

func main() {

	flag.Parse()

	local := []*cli.Command{
		runCmd,
	}
	app := &cli.App{
		Name:                 "gozc",
		Usage:                "go-zero CRUD生成器",
		Version:              "1.0",
		EnableBashCompletion: true,
		Commands:             local,
	}
	app.Setup()

	//os.Args = append(os.Args, "run", "--sql=./sys_auth.sql")

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		return
	}

}

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Print worker info",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "sql",
			Usage: "sql文件路径",
			Value: "",
		},
	},
	Before: func(cctx *cli.Context) error {
		data := os.Getenv("GOZC_PATH")
		if data == "" {
			return errors.New("请先设置环境变量GOZC_PATH路径指向Tpl文件夹")
		}
		fmt.Println("GOZC_PATH路径:" + data)
		return nil
	},
	Action: func(cctx *cli.Context) error {
		var srcPath string
		srcPath = cctx.String("sql")
		if srcPath == "" {
			return errors.New("请传入SQL文件路径地址")
		}
		if strings.Contains(srcPath, ".\\") {
			dir, _ := os.Getwd()
			println(dir)
			srcPath = strings.Replace(srcPath, ".\\", dir+"\\", -1)
		}
		if strings.Contains(srcPath, "./") {
			dir, _ := os.Getwd()
			println(dir)
			srcPath = strings.Replace(srcPath, "./", dir+"\\", -1)
		}
		fmt.Println(srcPath)
		// 讲sql文件 转换成tables
		tables, err := parser.Parse(srcPath)
		if err != nil {
			return err
		}
		err = CreateApiData(tables, srcPath)
		if err != nil {
			return err
		}

		err = CreateRpcData(tables, srcPath)
		if err != nil {
			return err
		}
		return nil
	},
}

func CreateApiData(tables []*parser.Table, srcPath string) error {
	// 生成API
	apiData := make(map[string]*api.CodeTuple)
	for _, e := range tables {
		Api, err := api.GenApiModel(*e, "admin", "api")
		if err != nil {
			return err
		}
		ApiInsert, err := api.GenApiModel(*e, "admin", "Insert")
		if err != nil {
			return err
		}
		ApiDelete, err := api.GenApiModel(*e, "admin", "delete")
		if err != nil {
			return err
		}
		ApiUpdate, err := api.GenApiModel(*e, "admin", "update")
		if err != nil {
			return err
		}
		ApiFindOne, err := api.GenApiModel(*e, "admin", "findOne")
		if err != nil {
			return err
		}
		ApiFindList, err := api.GenApiModel(*e, "admin", "findList")
		if err != nil {
			return err
		}
		apiData[e.Name.Source()] = &api.CodeTuple{
			Api:         Api,
			ApiInsert:   ApiInsert,
			ApiDelete:   ApiDelete,
			ApiUpdate:   ApiUpdate,
			ApiFindOne:  ApiFindOne,
			ApiFindList: ApiFindList,
		}
	}
	err := api.CreateFileApi(apiData, srcPath)
	return err
}

func CreateRpcData(tables []*parser.Table, srcPath string) error {
	// 生成Rpc
	rpcData := make(map[string]*rpc.CodeTuple)
	for _, e := range tables {
		Rpc, err := rpc.GenRpcModel(*e, "admin", "rpc")
		if err != nil {
			return err
		}
		RpcInsert, err := rpc.GenRpcModel(*e, "admin", "Insert")
		if err != nil {
			return err
		}
		RpcDelete, err := rpc.GenRpcModel(*e, "admin", "delete")
		if err != nil {
			return err
		}
		RpcUpdate, err := rpc.GenRpcModel(*e, "admin", "update")
		if err != nil {
			return err
		}
		RpcFindOne, err := rpc.GenRpcModel(*e, "admin", "findOne")
		if err != nil {
			return err
		}
		RpcFindList, err := rpc.GenRpcModel(*e, "admin", "findList")
		if err != nil {
			return err
		}
		rpcData[e.Name.Source()] = &rpc.CodeTuple{
			Rpc:         Rpc,
			RpcInsert:   RpcInsert,
			RpcDelete:   RpcDelete,
			RpcUpdate:   RpcUpdate,
			RpcFindOne:  RpcFindOne,
			RpcFindList: RpcFindList,
		}
	}
	err := rpc.CreateFileRpc(rpcData, srcPath)
	return err
}
