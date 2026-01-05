func (l *{{.filename}}DelLogic) {{.filename}}Del(req *types.{{.filename}}DelRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.{{.modelname}}Rpc.{{.filename}}Delete(l.ctx, &{{.xmodelname}}client.{{.filename}}DeleteReq{
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
