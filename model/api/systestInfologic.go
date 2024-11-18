func (l *SysTestInfoLogic) SysTestInfo(req *types.SysTestInfoRequest) (*types.Response, error) {

	res, err := l.svcCtx.UserRpc.SysTestFindOne(l.ctx, &userclient.SysTestFindOneReq{
		Id:	 req.Id, // 雪花算法ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	
	var result SysTestFindOneResp
	_ = copier.Copy(&result, res)
	
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}


type SysTestFindOneResp struct {
	Id  int64  `json:"id"`  // 雪花算法ID,
	Name  string  `json:"name"`  // 名称,
	CreateUserUid  string  `json:"create_user_uid"`  // 创建者userUid,
	CreateTime  int64  `json:"create_time"`  // 创建时间,
	ModifiedUserUid  string  `json:"modified_user_uid"`  // 修改者userUid,
	ModifiedTime  int64  `json:"modified_time"`  // 修改时间
}

