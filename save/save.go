/*
 * Handle saving and making those files
 */
package save

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

const (
	folderName string = "zapp"
	fileName string = "saves.zapp"
)

var (
	saveDir string
	saveFile string
)

func Init() error {

	// Get the home directory 
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Check the operating systems to load the correct save directory
	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("AppData")
		if appData == "" {
			appData = filepath.Join(homeDir, "AppData", "Roaming")
		}

		saveDir = filepath.Join(appData, folderName)
	case "darwin":
		saveDir = filepath.Join(homeDir, "Library", "Application Support", folderName)
	default:
		saveDir = filepath.Join(homeDir, ".local", "share", folderName)
	}

	// Make the directory needed
	if err := os.MkdirAll(saveDir, 0770); err != nil {
		return err
	}

	// Get the path of the save ifle
	filePath := filepath.Join(saveDir, fileName)

	// Create the basic file
	if _, err := os.Stat(filePath); err != nil {
		err := os.WriteFile(filePath, []byte("{}"), 0770)
		if err != nil {
			return err
		}
	}

	// The file path exists!
	saveFile = filePath

	// No errors :D
	return nil
}

// Read the save file
func ReadSaveFile(mpwd string) (*PasswordFile, error) {
	
	// Read the file's data
	data, err := os.ReadFile(saveFile)
	if err != nil {
		return nil, err
	}

	// Decrypt the file
	// decrypted, err := crypt.Decrypt(data, mpwd)
	// if err != nil {
	// 	return nil, err
	// }
	decrypted := data

	// Special info
	if decrypted == nil {
		decrypted = []byte("{}")
	}

	// Unmarshal the JSON
	var pFile PasswordFile
	err = json.Unmarshal(decrypted, &pFile)
	if err != nil {
		return nil, err
	}

	// Return the data
	return &pFile, nil
}

// Write the save file
func WriteSaveFile(pFile *PasswordFile, mpwd string) error {

	marshaled, err := json.Marshal(pFile)
	if err != nil {
		return err
	}

	// encrypted, err := crypt.Encrypt(marshaled, mpwd)
	// if err != nil {
	// 	return err
	// }
	encrypted := marshaled

	err = os.WriteFile(saveFile, encrypted, 0770)
	if err != nil {
		return err
	}

	return nil
}