func (l *SysAdminUpdateLogic) SysAdminUpdate(in *adminclient.SysAdminUpdateReq) (*adminclient.CommonResp, error) {

	res, err := l.svcCtx.SysAdminModel.FindOne(l.ctx,in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("SysAdmin没有该ID：" + in.Id)
		}
		return nil, err
	}

	

	// 用户名
	if len(in.Name) > 0 {
		res.Name = in.Name
	}
	// 姓名
	if len(in.NickName) > 0 {
		res.NickName = in.NickName
	}
	// 头像 
	if len(in.Avatar) != 0 {
		res.Avatar.String = in.Avatar
		res.Avatar.Valid = true
	}
	// 密码
	if len(in.Password) > 0 {
		res.Password = in.Password
	}
	// 邮箱
	if len(in.Email) > 0 {
		res.Email = in.Email
	}
	// 手机号
	if len(in.Telephone) > 0 {
		res.Telephone = in.Telephone
	}
	// 状态
	if in.State != 0 {
		res.State = in.State
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysAdminModel.Update(l.ctx,res)

	if err != nil {
		return nil, err
	}
	return &adminclient.CommonResp{}, nil

}