func (l *{{.filename}}InfoLogic) {{.filename}}Info(req types.{{.filename}}InfoRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	res, err := l.svcCtx.{{.modelname}}Rpc.{{.filename}}FindOne(l.ctx, &{{.modelname}}.{{.filename}}FindOneReq{
		{{.data}}
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	
	var result {{.filename}}FindOneResp
	_ = copier.Copy(&result, res)
	
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}


type {{.filename}}FindOneResp struct {
	{{.modeldata}}
}

