package config

import (
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

const configDirPath = "/config/"

// LoadEnv class is responsible for loading order of .env files. Uses godotenv.Load to inject env variables from
// file into environment;
func LoadEnv(parentFolderName string) {
	profile := os.Getenv("PROFILE")

	if "" != profile {
		load(".env."+profile, parentFolderName)
		load(".env", parentFolderName) // The Original .env
	}
}

func load(file string, parentFolderName string) {
	err := godotenv.Load("." + configDirPath + file)
	if err != nil {
		rootPath := getRootPath(parentFolderName)
		err := godotenv.Load(rootPath + configDirPath + file)
		if err != nil {
			panic("Error loading file: " + file)
		}
	}
}

func getRootPath(dirName string) string {
	projectName := regexp.MustCompile(`^(.*` + dirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	return string(projectName.Find([]byte(currentWorkDirectory)))
}
