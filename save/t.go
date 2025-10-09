/*
 * Types for the saves
 */
package save

// Basic save file type
type PasswordFile map[string]map[string]Password

// The type of a password saved
type Password struct {
	Password string
	Email string
}