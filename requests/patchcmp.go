package requests

type PatchCmpRequest struct {
	MasterPassword string
	NewMasterPassword string
}

func (req *PatchCmpRequest) VerifyRequest() error {
	if req.MasterPassword == "" || req.NewMasterPassword == "" { return ErrMissingMasterPassword }
	return nil
}