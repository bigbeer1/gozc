func (l *{{.filename}}UpdateLogic) {{.filename}}Update(in *{{.xmodelname}}client.{{.filename}}UpdateReq) (*{{.xmodelname}}client.CommonResp, error) {

	res, err := l.svcCtx.{{.filename}}Model.FindOne(l.ctx,in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("{{.filename}}没有该ID：" + in.Id)
		}
		return nil, err
	}

	{{.tenant}}

	{{.updateData}}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.{{.filename}}Model.Update(l.ctx,res)

	if err != nil {
		return nil, err
	}
	return &{{.xmodelname}}client.CommonResp{}, nil

}