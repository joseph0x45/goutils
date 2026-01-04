package goutils

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"golang.org/x/crypto/bcrypt"
)

var appName string = ""

func appNameIsNotEmpty() {
	if appName == "" {
		panic("app name has not been set")
	}
}

func SetAppName(name string) {
	appName = name
}

func EnsureDirExists(path string, perm ...os.FileMode) error {
	perms := os.FileMode(0755)
	if len(perm) != 0 {
		perms = perm[0]
	}
	return os.MkdirAll(path, perms)
}

func AppDataDir(user *user.User) string {
	appNameIsNotEmpty()
	return path.Join(
		user.HomeDir,
		".local/share",
		appName,
	)
}

// Ensures app data dir and config dir exist
// Returns path to the database
func Setup() string {
	appNameIsNotEmpty()
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if err := EnsureDirExists(AppDataDir(user)); err != nil {
		panic(err)
	}
	if err := EnsureDirExists(AppConfigDir(user)); err != nil {
		panic(err)
	}
	return AppDatabasePath(user)

}

func AppConfigDir(user *user.User) string {
	appNameIsNotEmpty()
	return path.Join(
		user.HomeDir,
		".config",
		appName,
	)
}

func AppDatabasePath(user *user.User) string {
	appNameIsNotEmpty()
	return path.Join(
		AppDataDir(user),
		appName+".db",
	)
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("Error while hashing password: %w", err)
	}
	return string(hashed), nil
}

func HashMatchesPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	) == nil
}
