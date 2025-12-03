package cli

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates
var TemplateFile embed.FS

type ConfigFile struct {
	Home string
}

func InitConfig(path string) (string, error) {
	tmpl, err := template.ParseFS(TemplateFile, "templates/config.yaml.tmpl")
	if err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}

	var homedir string
	if path == "" {
		homedir, err = os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting home directory: %w", err)
		}
	} else {
		homedir = path
	}

	configFilePath := filepath.Join(homedir, ".config", "cliconfig.yaml")
	if path != "" {
		configFilePath = filepath.Join(homedir, "cliconfig.yaml")
	}

	if _, err := os.Stat(configFilePath); !os.IsNotExist(err) {
		return "Config file already exists", nil
	}

	tmplFile, err := os.Create(configFilePath)
	if err != nil {
		return "", fmt.Errorf("error creating config file: %w", err)
	}

	content := ConfigFile{Home: homedir}
	err = tmpl.Execute(tmplFile, content)
	if err != nil {
		// Attempt to close the file, but we'll return the execution error.
		_ = tmplFile.Close()
		return "", fmt.Errorf("error executing template: %w", err)
	}

	if err := tmplFile.Close(); err != nil {
		return "", fmt.Errorf("error closing config file: %w", err)
	}

	return "Config created successfully", nil
}
