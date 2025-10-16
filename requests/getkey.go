package requests

type GetKey struct {
	PublicKey []byte
}

func (req *GetKey) VerifyRequest() error {
	if len(req.PublicKey) != 2048 {
		return ErrInvalidPublicKey
	}
	return nil
}