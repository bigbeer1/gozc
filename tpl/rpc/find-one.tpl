func (l *{{.filename}}FindOneLogic) {{.filename}}FindOne(in *{{.xmodelname}}client.{{.filename}}FindOneReq) (*{{.xmodelname}}client.{{.filename}}FindOneResp, error) {

	res, err := l.svcCtx.{{.filename}}Model.FindOne(l.ctx,in.Id)
	if err != nil {
		if errors.Is(err,sqlc.ErrNotFound) {
			return nil, fmt.Errorf("{{.filename}}没有该ID: %s", in.Id)
		}
		return nil, err
	}

    {{.deletedAtData}}
    {{.tenant}}

	return &{{.xmodelname}}client.{{.filename}}FindOneResp{
		{{.findoneData}}
	}, nil

}