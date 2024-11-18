func (l *SysTestUpLogic) SysTestUp(req *types.SysTestUpRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)
	
	_, err := l.svcCtx.UserRpc.SysTestUpdate(l.ctx, &userclient.SysTestUpdateReq{
	    Id:	 req.Id, // 雪花算法ID
		Name:	 req.Name, // 名称
		ModifiedUserUid:	 tokenData.NickName, // 修改者userUid
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
