func (l *{{.filename}}AddLogic) {{.filename}}Add(req *types.{{.filename}}AddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.{{.modelname}}Rpc.{{.filename}}Add(l.ctx, &{{.xmodelname}}client.{{.filename}}AddReq{
	    {{.data}}
	})

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}
