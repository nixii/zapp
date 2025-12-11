package requests

type TransferRequest struct {
	Receiving bool
	MasterPassword string
}

func (req *TransferRequest) VerifyRequest() error {
	if req.MasterPassword == ""  { return ErrMissingMasterPassword }
	return nil
}
