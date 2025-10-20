/*
 * Types for the saves
 */
package save

import (
	"errors"
)

// Basic save file type
type PasswordFile map[string]map[string]Password

// The type of a password saved
type Password struct {
	Password string
	Email string
}

// Errors
var (
	ErrSiteDoesntExistError error = errors.New("site doesn't exist")
	ErrUserThisUsernameNotFoundError error = errors.New("the user does not have this account on this site")
	ErrPasswordNotFound error = errors.New("password not found in the file")
)

// Get stuff about the password file
func (pwf *PasswordFile) GetSite(site string) (map[string]Password, error) {
	siteInfo, ok := (*pwf)[site]
	if !ok {
		return nil, ErrSiteDoesntExistError
	}

	return siteInfo, nil
}

// Get a specific password
func (pwf *PasswordFile) GetAccountInfo(site string, username string) (*Password, error) {
	siteinfo, err := pwf.GetSite(site)
	if err != nil {
		return nil, err
	}

	user, ok := siteinfo[username]
	if !ok {
		return nil, ErrUserThisUsernameNotFoundError
	}

	return &user, nil
}

// Write a password
func (pwf *PasswordFile) SetPassword(site string, username string, password Password) error {
	siteinfo := (*pwf)[site]
	if siteinfo == nil {
		siteinfo = make(map[string]Password)
	}

	siteinfo[username] = password
	(*pwf)[site] = siteinfo

	return nil
}

// Remove a password
func (pwf *PasswordFile) RemovePassword(site string, username string) error {
	siteinfo := (*pwf)[site]
	if siteinfo == nil {
		return ErrPasswordNotFound
	}

	delete(siteinfo, username)

	if len(siteinfo) != 0 {
		(*pwf)[site] = siteinfo
	} else {
		(*pwf)[site] = nil
	}

	return nil
}