func (l *{{.filename}}ListLogic) {{.filename}}List(in *{{.xmodelname}}client.{{.filename}}ListReq) (*{{.xmodelname}}client.{{.filename}}ListResp, error) {

  	whereBuilder := l.svcCtx.{{.filename}}Model.RowBuilder()

    {{.listDeletedAt}}
    {{.listCreatedAt}}

    {{.listTenant}}

    {{.listData}}

	all, err := l.svcCtx.{{.filename}}Model.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
    if err != nil {
    	return nil, err
    }

    countBuilder := l.svcCtx.{{.filename}}Model.CountBuilder("id")

    {{.countDeletedAt}}

    {{.countTenant}}

    {{.countData}}
    count, err := l.svcCtx.{{.filename}}Model.FindCount(l.ctx, countBuilder)
    if err != nil {
    	return nil, err
    }

    var list []*{{.xmodelname}}client.{{.filename}}ListData
    for _, item := range all {
    	list = append(list, &{{.xmodelname}}client.{{.filename}}ListData{
    		{{.findlistData}}
    	})
    }

    return &{{.xmodelname}}client.{{.filename}}ListResp{
    	Total: count,
    	List:  list,
    }, nil
}
