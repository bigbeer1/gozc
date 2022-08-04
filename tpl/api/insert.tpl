func (l *{{.filename}}AddLogic) {{.filename}}Add(req types.{{.filename}}AddRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.{{.modelname}}Rpc.{{.filename}}Add(l.ctx, &{{.modelname}}.{{.filename}}AddReq{
	    {{.data}}
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}
