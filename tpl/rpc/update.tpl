func (l *{{.filename}}UpdateLogic) {{.filename}}Update(in *{{.xmodelname}}client.{{.filename}}UpdateReq) (*{{.xmodelname}}client.CommonResp, error) {

	res, err := l.svcCtx.{{.filename}}Model.FindOne(l.ctx,in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil,fmt.Errorf("{{.filename}}没有该ID: %v" , in.Id)
		}
		return nil, err
	}

    {{.deletedAtData}}
	{{.tenant}}

	{{.updateData}}

	res.ModifiedUserUid.String = in.ModifiedUserUid
	res.ModifiedUserUid.Valid = true
	res.ModifiedTime.Time = time.Now()
	res.ModifiedTime.Valid = true

	err = l.svcCtx.{{.filename}}Model.Update(l.ctx,res)

	if err != nil {
		return nil, err
	}
	return &{{.xmodelname}}client.CommonResp{}, nil

}