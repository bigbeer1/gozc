func (l *SysAdminListLogic) SysAdminList(in *adminclient.SysAdminListReq) (*adminclient.SysAdminListResp, error) {

  	whereBuilder := l.svcCtx.SysAdminModel.RowBuilder()

    whereBuilder = whereBuilder.Where("deleted_at is null")
    whereBuilder = whereBuilder.OrderBy("created_at DESC")

    

    // 用户名
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%"+in.Name+"%",
		})
	}
	// 姓名
	if len(in.NickName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"nick_name ": "%"+in.NickName+"%",
		})
	}
	// 头像
	if len(in.Avatar) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"avatar ": "%"+in.Avatar+"%",
		})
	}
	// 密码
	if len(in.Password) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"password ": "%"+in.Password+"%",
		})
	}
	// 邮箱
	if len(in.Email) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"email ": "%"+in.Email+"%",
		})
	}
	// 手机号
	if len(in.Telephone) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"telephone ": "%"+in.Telephone+"%",
		})
	}
	// 状态
	if in.State != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}

	all, err := l.svcCtx.SysAdminModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
    if err != nil {
    	return nil, err
    }

    countBuilder := l.svcCtx.SysAdminModel.CountBuilder("id")

    countBuilder = countBuilder.Where("deleted_at is null")

    

    // 用户名
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%"+in.Name+"%",
		})
	}
	// 姓名
	if len(in.NickName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"nick_name ": "%"+in.NickName+"%",
		})
	}
	// 头像
	if len(in.Avatar) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"avatar ": "%"+in.Avatar+"%",
		})
	}
	// 密码
	if len(in.Password) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"password ": "%"+in.Password+"%",
		})
	}
	// 邮箱
	if len(in.Email) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"email ": "%"+in.Email+"%",
		})
	}
	// 手机号
	if len(in.Telephone) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"telephone ": "%"+in.Telephone+"%",
		})
	}
	// 状态
	if in.State != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}
    count, err := l.svcCtx.SysAdminModel.FindCount(l.ctx, countBuilder)
    if err != nil {
    	return nil, err
    }

    var list []*adminclient.SysAdminListData
    for _, item := range all {
    	list = append(list, &adminclient.SysAdminListData{
    		Id:	item.Id, //系统管理员ID
			CreatedAt:	item.CreatedAt.UnixMilli(), //创建时间
			UpdatedAt:	item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName:	item.CreatedName, //创建人
			UpdatedName:	item.UpdatedName.String, //更新人
			Name:	item.Name, //用户名
			NickName:	item.NickName, //姓名
			Avatar:	item.Avatar.String, //头像
			Password:	item.Password, //密码
			Email:	item.Email, //邮箱
			Telephone:	item.Telephone, //手机号
			State:	item.State, //状态
    	})
    }

    return &adminclient.SysAdminListResp{
    	Total: count,
    	List:  list,
    }, nil
}
