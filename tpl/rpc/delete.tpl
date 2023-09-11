func (l *{{.filename}}DeleteLogic) {{.filename}}Delete(in *{{.xmodelname}}client.{{.filename}}DeleteReq) (*{{.xmodelname}}client.CommonResp, error) {

	res, err := l.svcCtx.{{.filename}}Model.FindOne(l.ctx,in.Id)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("{{.filename}}没有该ID：" + in.Id)
		}
		return nil, err
	}

    {{.deletedAtData}}
    {{.tenant}}

    {{.del}}

	err = l.svcCtx.{{.filename}}Model.{{.delType}}
	if err != nil {
		return nil, err
	}

	return &{{.xmodelname}}client.CommonResp{}, nil
}