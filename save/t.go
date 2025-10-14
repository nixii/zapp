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
	SiteDoesntExistError error = errors.New("Site doesn't exist")
	UserThisUsernameNotFoundError error = errors.New("The user does not have this account on this site")
)

// Get stuff about the password file
func (self *PasswordFile) GetSite(site string) (map[string]Password, error) {
	siteInfo, ok := (*self)[site]
	if !ok {
		return nil, SiteDoesntExistError
	}

	return siteInfo, nil
}

// Get a specific password
func (self *PasswordFile) GetAccountInfo(site string, username string) (*Password, error) {
	siteinfo, err := self.GetSite(site)
	if err != nil {
		return nil, err
	}

	user, ok := siteinfo[username]
	if !ok {
		return nil, UserThisUsernameNotFoundError
	}

	return &user, nil
}

// Write a password
func (self *PasswordFile) SetPassword(site string, username string, password Password) error {
	siteinfo := (*self)[site]
	if siteinfo == nil {
		siteinfo = make(map[string]Password)
	}

	siteinfo[username] = password
	(*self)[site] = siteinfo

	return nil
}