func (l *SysAdminListLogic) SysAdminList(req *types.SysAdminListRequest) (*types.Response, error) {

	all, err := l.svcCtx.AdminRpc.SysAdminList(l.ctx, &adminclient.SysAdminListReq{
		Current:	 req.Current, // 页码
		PageSize:	 req.PageSize, // 页数
		Name:	 req.Name, // 用户名
		NickName:	 req.NickName, // 姓名
		Avatar:	 req.Avatar, // 头像
		Password:	 req.Password, // 密码
		Email:	 req.Email, // 邮箱
		Telephone:	 req.Telephone, // 手机号
		State:	 req.State, // 状态
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	
	var result SysAdminListResp
	_ = copier.Copy(&result, all)
	
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}


type SysAdminListResp struct {
	Total int64                     `json:"total"`
	List  []*SysAdminDataList  `json:"list"`
}

type SysAdminDataList struct {
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

