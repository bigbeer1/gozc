func (l *SysAdminAddLogic) SysAdminAdd(req *types.SysAdminAddRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.AdminRpc.SysAdminAdd(l.ctx, &adminclient.SysAdminAddReq{
	    CreatedName:	 tokenData.NickName, // 创建人
		Name:	 req.Name, // 用户名
		NickName:	 req.NickName, // 姓名
		Avatar:	 req.Avatar, // 头像
		Password:	 req.Password, // 密码
		Email:	 req.Email, // 邮箱
		Telephone:	 req.Telephone, // 手机号
		State:	 req.State, // 状态
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
