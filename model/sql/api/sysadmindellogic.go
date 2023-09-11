func (l *SysAdminDelLogic) SysAdminDel(req *types.SysAdminDelRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.UserRpc.SysAdminDelete(l.ctx, &userclient.SysAdminDeleteReq{
	    Id:	 req.Id, // 系统管理员ID
		DeletedName:	 tokenData.NickName, // 删除人
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
