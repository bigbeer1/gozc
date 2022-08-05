func (l *{{.filename}}AddLogic) {{.filename}}Add(in *{{.xmodelname}}client.{{.filename}}AddReq) (*{{.xmodelname}}client.CommonResp, error) {

	_, err := l.svcCtx.{{.filename}}Model.Insert(l.ctx,&model.{{.filename}}{
		Id:           uuid.NewV4().String(),  // ID
		CreatedAt:    time.Now(), // 创建时间
        {{.data}}
	})
	if err != nil {
		return nil, err
	}

	return &{{.xmodelname}}client.CommonResp{}, nil
}
