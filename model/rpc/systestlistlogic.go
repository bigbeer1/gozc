func (l *SysTestListLogic) SysTestList(in *userclient.SysTestListReq) (*userclient.SysTestListResp, error) {

  	whereBuilder := l.svcCtx.SysTestModel.RowBuilder()

    whereBuilder = whereBuilder.Where(squirrel.Eq{"deleted": 0,})
	
    whereBuilder = whereBuilder.OrderBy("create_time DESC, id DESC")

    

    // 名称
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%"+in.Name+"%",
		})
	}

	all, err := l.svcCtx.SysTestModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
    if err != nil {
    	return nil, err
    }

    countBuilder := l.svcCtx.SysTestModel.CountBuilder("id")

    countBuilder = countBuilder.Where(squirrel.Eq{"deleted": 0,})

    

    // 名称
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%"+in.Name+"%",
		})
	}
    count, err := l.svcCtx.SysTestModel.FindCount(l.ctx, countBuilder)
    if err != nil {
    	return nil, err
    }

    var list []*userclient.SysTestListData
    for _, item := range all {
    	list = append(list, &userclient.SysTestListData{
    		Id:	item.Id, //雪花算法ID
			Name:	item.Name, //名称
			CreateUserUid:	item.CreateUserUid, //创建者userUid
			CreateTime:	item.CreateTime.UnixMilli(), //创建时间
			ModifiedUserUid:	item.ModifiedUserUid.String, //修改者userUid
			ModifiedTime:	item.ModifiedTime.Time.UnixMilli(), //修改时间
    	})
    }

    return &userclient.SysTestListResp{
    	Total: count,
    	List:  list,
    }, nil
}
