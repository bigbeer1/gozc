func (l *SysAdminUpLogic) SysAdminUp(req *types.SysAdminUpRequest) (*types.Response, error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err := l.svcCtx.AdminRpc.SysAdminUpdate(l.ctx, &SysAdminclient.SysAdminUpdateReq{
	    Id:	 req.Id, // 系统管理员ID
		UpdatedName:	 tokenData.NickName, // 更新人
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
