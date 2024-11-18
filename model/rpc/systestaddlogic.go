func (l *SysTestAddLogic) SysTestAdd(in *userclient.SysTestAddReq) (*userclient.CommonResp, error) {

	_, err := l.svcCtx.SysTestModel.Insert(l.ctx,&model.SysTest{
		Id:            IDx.NextId(),  // 雪花ID
		CreateTime:    time.Now(), // 创建时间
        Name:	 in.Name, // 名称
		CreateUserUid:	 in.CreateUserUid, // 创建者userUid
	})
	if err != nil {
		return nil, err
	}

	return &userclient.CommonResp{}, nil
}
