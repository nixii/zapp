package requests

type GetAllRequest struct {
	MasterPassword string
}

func (req *GetAllRequest) VerifyRequest() error {
	if req.MasterPassword == "" { return ErrMissingMasterPassword }
	return nil
}