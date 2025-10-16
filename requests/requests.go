/*
 * Basic information that all requests will need
 */
package requests

import "errors"

// Define errors
var (
	ErrWebsiteRequired error = errors.New("this request needs a website")
	ErrUsernameRequired error = errors.New("this request needs a username")
	ErrIncorrectMasterPasswordOrMangledJson error = errors.New("the master password is incorrect, or your JSON data is mangled")
	ErrMissingMasterPassword error = errors.New("this request needs a master password")
	ErrInvalidPublicKey error = errors.New("public keys must be 2048 bytes")
)