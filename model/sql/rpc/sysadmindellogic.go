func (l *SysAdminDeleteLogic) SysAdminDelete(in *adminclient.SysAdminDeleteReq) (*adminclient.CommonResp, error) {

	res, err := l.svcCtx.SysAdminModel.FindOne(l.ctx,in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysAdmin没有该ID:%v" ,in.Id)
		}
		return nil, err
	}

    // 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysAdmin该ID已被删除：%v",in.Id)
	}
    

    res.DeletedAt.Time = time.Now()
	res.DeletedAt.Valid = true
	res.DeletedName.String = in.DeletedName
	res.DeletedName.Valid = true

	err = l.svcCtx.SysAdminModel.Update(l.ctx,res)
	if err != nil {
		return nil, err
	}

	return &adminclient.CommonResp{}, nil
}