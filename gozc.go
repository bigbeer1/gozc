package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/urfave/cli/v2"
	"gozc/gen/api"
	"gozc/gen/http"
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

	os.Args = append(os.Args, "run", "--m=admin", "--sql=E:\\Gopath\\src\\gozc\\model\\sql\\sys_admin.sql")

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
		&cli.StringFlag{
			Name:  "m",
			Usage: "模型名称",
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
		// 拿sql文件地址
		srcPath = cctx.String("sql")
		if srcPath == "" {
			return errors.New("请传入SQL文件路径地址--sql")
		}

		// 拿微服务名称
		modelName := cctx.String("m")
		if srcPath == "" {
			return errors.New("请传入模型名称--m")
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

		// 生成API
		err = CreateApiData(tables, modelName, srcPath)
		if err != nil {
			return err
		}

		// 生成RPC
		err = CreateRpcData(tables, modelName, srcPath)
		if err != nil {
			return err
		}

		// 生成HTTP
		err = CreateHttpData(tables, modelName, srcPath)
		if err != nil {
			return err
		}
		return nil
	},
}

// 去生成API里的代码
func CreateApiData(tables []*parser.Table, modelName, srcPath string) error {
	// 生成API
	apiData := make(map[string]*api.CodeTuple)
	for _, e := range tables {
		Api, err := api.GenApiModel(*e, modelName, "api")
		if err != nil {
			return err
		}
		ApiInsert, err := api.GenApiModel(*e, modelName, "Insert")
		if err != nil {
			return err
		}
		ApiDelete, err := api.GenApiModel(*e, modelName, "delete")
		if err != nil {
			return err
		}
		ApiUpdate, err := api.GenApiModel(*e, modelName, "update")
		if err != nil {
			return err
		}
		ApiFindOne, err := api.GenApiModel(*e, modelName, "findOne")
		if err != nil {
			return err
		}
		ApiFindList, err := api.GenApiModel(*e, modelName, "findList")
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

// 去生成RPC里的代码
func CreateRpcData(tables []*parser.Table, modelName, srcPath string) error {
	// 生成Rpc
	rpcData := make(map[string]*rpc.CodeTuple)
	for _, e := range tables {
		Rpc, err := rpc.GenRpcModel(*e, modelName, "rpc")
		if err != nil {
			return err
		}
		RpcInsert, err := rpc.GenRpcModel(*e, modelName, "Insert")
		if err != nil {
			return err
		}
		RpcDelete, err := rpc.GenRpcModel(*e, modelName, "delete")
		if err != nil {
			return err
		}
		RpcUpdate, err := rpc.GenRpcModel(*e, modelName, "update")
		if err != nil {
			return err
		}
		RpcFindOne, err := rpc.GenRpcModel(*e, modelName, "findOne")
		if err != nil {
			return err
		}
		RpcFindList, err := rpc.GenRpcModel(*e, modelName, "findList")
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

// 去生成SwaggerAPI 规范的东西
func CreateHttpData(tables []*parser.Table, modelName, srcPath string) error {
	// 生成Http
	HttpData := make(map[string]*http.CodeTuple)
	for _, e := range tables {
		Api, err := http.GenHttpModel(*e, modelName, "http-api")
		if err != nil {
			return err
		}
		HttpData[e.Name.Source()] = &http.CodeTuple{
			Api: Api,
		}
	}
	err := http.CreateFileHttp(HttpData, srcPath)
	return err
}
