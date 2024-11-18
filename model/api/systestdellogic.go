func (l *SysTestDelLogic) SysTestDel(req *types.SysTestDelRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.UserRpc.SysTestDelete(l.ctx, &userclient.SysTestDeleteReq{
	    Id:	 req.Id, // 雪花算法ID
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
