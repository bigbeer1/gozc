func (l *SysTestListLogic) SysTestList(req *types.SysTestListRequest) (*types.Response, error) {

	all, err := l.svcCtx.UserRpc.SysTestList(l.ctx, &userclient.SysTestListReq{
		Current:	 req.Current, // 页码
		PageSize:	 req.PageSize, // 页数
		Name:	 req.Name, // 名称
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	
	var result SysTestListResp
	_ = copier.Copy(&result, all)
	
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}


type SysTestListResp struct {
	Total int64                     `json:"total"`
	List  []*SysTestDataList  `json:"list"`
}

type SysTestDataList struct {
	Id  int64  `json:"id"`  // 雪花算法ID,
	Name  string  `json:"name"`  // 名称,
	CreateUserUid  string  `json:"create_user_uid"`  // 创建者userUid,
	CreateTime  int64  `json:"create_time"`  // 创建时间,
	ModifiedUserUid  string  `json:"modified_user_uid"`  // 修改者userUid,
	ModifiedTime  int64  `json:"modified_time"`  // 修改时间
}

