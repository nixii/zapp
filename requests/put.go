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
func Put_VerifyRequest(req PutRequest) error {
	if req.Website == ""        { return WebsiteRequired }
	if req.Username == ""       { return UsernameRequired }
	if req.MasterPassword == "" {return MissingMasterPassword}
	return nil
}