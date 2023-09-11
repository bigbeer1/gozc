func (l *SysAdminDeleteLogic) SysAdminDelete(in *userclient.SysAdminDeleteReq) (*userclient.CommonResp, error) {

	res, err := l.svcCtx.SysAdminModel.FindOne(l.ctx,in.Id)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("SysAdmin没有该ID：" + in.Id)
		}
		return nil, err
	}

    // 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, errors.New("SysAdmin该ID已被删除：" + in.Id)
	}
    

    res.DeletedAt.Time = time.Now()
	res.DeletedAt.Valid = true
	res.DeletedName.String = in.DeletedName
	res.DeletedName.Valid = true

	err = l.svcCtx.SysAdminModel.Update(l.ctx,res)
	if err != nil {
		return nil, err
	}

	return &userclient.CommonResp{}, nil
}