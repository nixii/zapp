/*
 * Handle a put request
 * and its response
 */
package requests

// Create the type of a request
type PutRequest struct {
	Website string
	Username string
	MasterPassword string
	Password *string
	Email *string
}

// Verify a put reqest has all the information needed
func (req *PutRequest) VerifyRequest() error {
	if req.Website == ""        { return ErrWebsiteRequired }
	if req.Username == ""       { return ErrUsernameRequired }
	if req.MasterPassword == "" {return ErrMissingMasterPassword}
	return nil
}