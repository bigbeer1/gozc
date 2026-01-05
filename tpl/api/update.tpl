func (l *{{.filename}}UpLogic) {{.filename}}Up(req *types.{{.filename}}UpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.{{.modelname}}Rpc.{{.filename}}Update(l.ctx, &{{.xmodelname}}client.{{.filename}}UpdateReq{
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
