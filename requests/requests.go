/*
 * Basic information that all requests will need
 */
package requests

import "errors"

// Define errors
var (
	WebsiteRequired error = errors.New("This request needs a website")
	UsernameRequired error = errors.New("This request needs a username")
	IncorrectMasterPasswordOrMangledJson error = errors.New("The master password is incorrect, or your JSON data is mangled.")
	MissingMasterPassword error = errors.New("This request needs a master password")
)