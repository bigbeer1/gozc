func (l *{{.filename}}FindOneLogic) {{.filename}}FindOne(in *{{.xmodelname}}client.{{.filename}}FindOneReq) (*{{.xmodelname}}client.{{.filename}}FindOneResp, error) {

	res, err := l.svcCtx.{{.filename}}Model.FindOne(l.ctx,in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("{{.filename}}没有该ID：" + in.Id)
		}
		return nil, err
	}

    {{.deletedAtData}}
    {{.tenant}}

	return &{{.xmodelname}}client.{{.filename}}FindOneResp{
		{{.findoneData}}
	}, nil

}