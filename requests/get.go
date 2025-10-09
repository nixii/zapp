/*
 * The JSON get request for a get request
 * as well as the return type
 */
package requests

// Create the type of a get request
type GetRequest struct {
	Website string
	Username string
	MasterPassword string
}

// Create the type of a get response
type GetResponse struct {
	Password string
	Email string
}

// Verify a get request
func Get_VerifyRequest(req GetRequest) error {
	if req.Website == ""        { return WebsiteRequired }
	if req.Username == ""       { return UsernameRequired } 
	if req.MasterPassword == "" { return MissingMasterPassword }
	return nil
}