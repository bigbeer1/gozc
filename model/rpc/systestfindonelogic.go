func (l *SysTestFindOneLogic) SysTestFindOne(in *userclient.SysTestFindOneReq) (*userclient.SysTestFindOneResp, error) {

	res, err := l.svcCtx.SysTestModel.FindOne(l.ctx,in.Id)
	if err != nil {
		if  errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysTest没有该ID:%v" ,in.Id)
		}
		return nil, err
	}

    // 判断该数据是否被删除
	if res.Deleted == 1 {
		return nil, fmt.Errorf("SysTest该ID已被删除：%v",in.Id)
	}
    

	return &userclient.SysTestFindOneResp{
		Id:	res.Id, //雪花算法ID
		Name:	res.Name, //名称
		CreateUserUid:	res.CreateUserUid, //创建者userUid
		CreateTime:	res.CreateTime.UnixMilli(), //创建时间
		ModifiedUserUid:	res.ModifiedUserUid.String, //修改者userUid
		ModifiedTime:	res.ModifiedTime.Time.UnixMilli(), //修改时间
	}, nil

}