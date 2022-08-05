func (l *{{.filename}}ListLogic) {{.filename}}List(req *types.{{.filename}}ListRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	all, err := l.svcCtx.{{.modelname}}Rpc.{{.filename}}List(l.ctx, &{{.xmodelname}}client.{{.filename}}ListReq{
		{{.data}}
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	
	var result {{.filename}}ListResp
	_ = copier.Copy(&result, all)
	
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}


type {{.filename}}ListResp struct {
	Total int64                     `json:"total"`
	List  []*{{.filename}}DataList  `json:"list"`
}

type {{.filename}}DataList struct {
	{{.modeldata}}
}

