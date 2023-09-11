func (l *SysAdminAddLogic) SysAdminAdd(in *userclient.SysAdminAddReq) (*userclient.CommonResp, error) {

	_, err := l.svcCtx.SysAdminModel.Insert(l.ctx,&model.SysAdmin{
		Id:           uuid.NewV4().String(),  // ID
		CreatedAt:    time.Now(), // 创建时间
        CreatedName:	 in.CreatedName, // 创建人
		Name:	 in.Name, // 用户名
		NickName:	 in.NickName, // 姓名
		Avatar:	sql.NullString{String: in.Avatar, Valid: in.Avatar != ""}, // 头像
		Password:	 in.Password, // 密码
		Email:	 in.Email, // 邮箱
		Telephone:	 in.Telephone, // 手机号
		State:	 in.State, // 状态
	})
	if err != nil {
		return nil, err
	}

	return &userclient.CommonResp{}, nil
}
