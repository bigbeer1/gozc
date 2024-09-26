func (l *SysAdminInfoLogic) SysAdminInfo(req *types.SysAdminInfoRequest) (*types.Response, error) {

	res, err := l.svcCtx.AdminRpc.SysAdminFindOne(l.ctx, &adminclient.SysAdminFindOneReq{
		Id:	 req.Id, // 系统管理员ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	
	var result SysAdminFindOneResp
	_ = copier.Copy(&result, res)
	
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}


type SysAdminFindOneResp struct {
	Id  string  `json:"id"`  // 系统管理员ID,
	CreatedAt  int64  `json:"created_at"`  // 创建时间,
	UpdatedAt  int64  `json:"updated_at"`  // 更新时间,
	CreatedName  string  `json:"created_name"`  // 创建人,
	UpdatedName  string  `json:"updated_name"`  // 更新人,
	Name  string  `json:"name"`  // 用户名,
	NickName  string  `json:"nick_name"`  // 姓名,
	Avatar  string  `json:"avatar"`  // 头像,
	Password  string  `json:"password"`  // 密码,
	Email  string  `json:"email"`  // 邮箱,
	Telephone  string  `json:"telephone"`  // 手机号,
	State  int64  `json:"state"`  // 状态
}

