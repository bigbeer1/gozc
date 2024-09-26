func (l *SysAdminFindOneLogic) SysAdminFindOne(in *adminclient.SysAdminFindOneReq) (*adminclient.SysAdminFindOneResp, error) {

	res, err := l.svcCtx.SysAdminModel.FindOne(l.ctx,in.Id)
	if err != nil {
		if  errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysAdmin没有该ID:%v" ,in.Id)
		}
		return nil, err
	}

    // 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysAdmin该ID已被删除：%v",in.Id)
	}
    

	return &adminclient.SysAdminFindOneResp{
		Id:	res.Id, //系统管理员ID
		CreatedAt:	res.CreatedAt.UnixMilli(), //创建时间
		UpdatedAt:	res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName:	res.CreatedName, //创建人
		UpdatedName:	res.UpdatedName.String, //更新人
		Name:	res.Name, //用户名
		NickName:	res.NickName, //姓名
		Avatar:	res.Avatar.String, //头像
		Password:	res.Password, //密码
		Email:	res.Email, //邮箱
		Telephone:	res.Telephone, //手机号
		State:	res.State, //状态
	}, nil

}