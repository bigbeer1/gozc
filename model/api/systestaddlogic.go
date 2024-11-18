func (l *SysTestAddLogic) SysTestAdd(req *types.SysTestAddRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.UserRpc.SysTestAdd(l.ctx, &userclient.SysTestAddReq{
	    Name:	 req.Name, // 名称
		CreateUserUid:	 tokenData.NickName, // 创建者userUid
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
