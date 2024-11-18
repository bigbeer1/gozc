func (l *SysTestUpdateLogic) SysTestUpdate(in *userclient.SysTestUpdateReq) (*userclient.CommonResp, error) {

	res, err := l.svcCtx.SysTestModel.FindOne(l.ctx,in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil,fmt.Errorf("SysTest没有该ID: %v" , in.Id)
		}
		return nil, err
	}

    // 判断该数据是否被删除
	if res.Deleted == 1 {
		return nil, fmt.Errorf("SysTest该ID已被删除：%v",in.Id)
	}
	

	// 名称
	if len(in.Name) > 0 {
		res.Name = in.Name
	}

	res.ModifiedUserUid.String = in.ModifiedUserUid
	res.ModifiedUserUid.Valid = true
	res.ModifiedTime.Time = time.Now()
	res.ModifiedTime.Valid = true

	err = l.svcCtx.SysTestModel.Update(l.ctx,res)

	if err != nil {
		return nil, err
	}
	return &userclient.CommonResp{}, nil

}