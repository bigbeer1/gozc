syntax = "v1"


type (
    {{.filename}}AddRequest {
		{{.Add}}
    }

    {{.filename}}DelRequest {
		{{.Del}}
    }

    {{.filename}}UpRequest {
		{{.Up}}      
    }


    {{.filename}}ListRequest {
		{{.List}}       
    }

    {{.filename}}InfoRequest {
		{{.Info}}          
    }

)

@server(
    //声明当前service下所有路由需要jwt鉴权，且会自动生成包含jwt逻辑的代码
    jwt: Auth
    group: {{.xfilename}}
)

service {{.Amodelname}} {

    // 添加
    @handler {{.filename}}Add
    post /{{.modelname}}/{{.xfilename}} ({{.filename}}AddRequest) returns (Response)

    // 删除
    @handler {{.filename}}Del
    delete /{{.modelname}}/{{.xfilename}}/:id ({{.filename}}DelRequest) returns (Response)

    // 更新
    @handler {{.filename}}Up
    put /{{.modelname}}/{{.xfilename}} ({{.filename}}UpRequest) returns (Response)

    // 分页查询
    @handler {{.filename}}List
    get /{{.modelname}}/{{.xfilename}} ({{.filename}}ListRequest) returns (Response)

    // 查询详细信息
    @handler {{.filename}}Info
    get /{{.modelname}}/{{.xfilename}}Info ({{.filename}}InfoRequest) returns (Response)
}