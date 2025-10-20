package requests

type PatchRequest struct {
	From struct {
		Username string
		Website string
	}
	To struct {
		Username string
		Email string
		Password string
		Website string
	}
	MasterPassword string
}

func (req *PatchRequest) VerifyRequest() error {
	if req.MasterPassword == "" {
		return ErrMissingMasterPassword
	} else if req.From.Website == "" {
		return ErrWebsiteRequired
	} else if req.From.Username == "" {
		return ErrUsernameRequired
	}
	return nil
}