syntax = "proto3";

package {{.modelname}}client;

option go_package = "./{{.modelname}}client";

//{{.filename}} start-------------------

// {{.filename}} 添加
message {{.filename}}AddReq{
	{{.Add}}
}

// {{.filename}} 删除
message {{.filename}}DeleteReq{
	{{.Del}}
}

// {{.filename}} 更新
message {{.filename}}UpdateReq{
	{{.Up}}
}

// {{.filename}} 单个查询
message {{.filename}}FindOneReq{
	{{.FindOne}}
}

// {{.filename}} 单个查询返回
message {{.filename}}FindOneResp{
	{{.FindOneResp}}
}


// {{.filename}} 分页查询
message {{.filename}}ListReq{
	{{.List}}
}

// {{.filename}} 分页查询返回
message {{.filename}}ListResp{
	{{.ListResp}}
}

// {{.filename}} 列表信息
message {{.filename}}ListData{
	{{.ListData}}
}

//{{.filename}} end---------------------

service {{.amodelname}} {

  rpc {{.filename}}Add({{.filename}}AddReq) returns(CommonResp);
  rpc {{.filename}}Delete({{.filename}}DeleteReq) returns(CommonResp);
  rpc {{.filename}}Update({{.filename}}UpdateReq) returns(CommonResp);
  rpc {{.filename}}FindOne({{.filename}}FindOneReq) returns({{.filename}}FindOneResp);
  rpc {{.filename}}List({{.filename}}ListReq) returns({{.filename}}ListResp);

}