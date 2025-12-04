package requests;

import (
	"fmt"
);

type TransferRequest struct {
	Receiving bool,
	MasterPassword string
}

func (req *TransferRequest) VerifyRequest() error {
	if MasterPassword == ""  { return ErrMasterPasswordRequired }
	return nil
}
