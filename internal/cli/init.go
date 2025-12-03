package cli

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates
var TemplateFile embed.FS

type ConfigFile struct {
	Home string
}

func InitConfig(path string) {

	tmpl, err := template.ParseFS(TemplateFile, "templates/config.yaml.tmpl")
	if err != nil {
		log.Println("error to get template file config.yaml")
	}

	var content ConfigFile
	var homedir string
	var configFilePath string

	if path == "" {
		homedir, err = os.UserHomeDir()
		if err != nil {
			log.Println("error to get home directory path")
		}
		configFilePath = filepath.Join(homedir, ".config", "cliconfig.yaml")
	} else {
		homedir = path
		configFilePath = filepath.Join(homedir, "cliconfig.yaml")
	}

	content.Home = homedir

	var tmplFile *os.File
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		tmplFile, err = os.Create(configFilePath)
		if err != nil {
			log.Println("error to create cli config file")
		}

		err = tmpl.Execute(tmplFile, content)
		if err != nil {
			log.Println("error to process the template")
		}
		fmt.Println("Config created successfuly")
	} else {
		fmt.Println("Config file already exists")
	}
}
