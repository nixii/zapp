/*
 * The JSON get request for a get request
 * as well as the return type
 */
package requests

// Create the type of a get request
type DeleteRequest struct {
	Website string
	Username string
	MasterPassword string
}

// Verify a get request
func (req *DeleteRequest) VerifyRequest() error {
	if req.Website == ""        { return ErrWebsiteRequired }
	if req.Username == ""       { return ErrUsernameRequired } 
	if req.MasterPassword == "" { return ErrMissingMasterPassword }
	return nil
}