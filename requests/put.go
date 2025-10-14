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
func (self *PutRequest) VerifyRequest() error {
	if self.Website == ""        { return WebsiteRequired }
	if self.Username == ""       { return UsernameRequired }
	if self.MasterPassword == "" {return MissingMasterPassword}
	return nil
}