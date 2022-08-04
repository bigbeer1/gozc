func (l *{{.filename}}UpLogic) {{.filename}}Up(req types.{{.filename}}UpRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.{{.modelname}}Rpc.{{.filename}}Update(l.ctx, &{{.modelname}}.{{.filename}}UpdateReq{
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
