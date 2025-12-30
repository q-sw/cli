/*
Copyright © 2025 qsw
*/
package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/q-sw/cli/internal/project"
	"github.com/spf13/cobra"
)

var (
	projectName  string
	projectForce bool
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  `Scaffold and manage projects structure and configuration.`,
}

var projectInitCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize a new project structure",
	Long: `Scaffold a new project with standard files:
- README.md
- GEMINI.md
- .gitignore
- Pre-commit configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := "."
		if len(args) > 0 {
			targetPath = args[0]
		}

		absPath, err := filepath.Abs(targetPath)
		if err != nil {
			log.Fatalf("could not resolve path: %v", err)
		}

		if projectName == "" {
			projectName = filepath.Base(absPath)
		}

		data := project.ProjectData{
			ProjectName: projectName,
			Description: "Project created with cli",
		}

		fmt.Printf("Initializing project '%s' in %s\n", data.ProjectName, absPath)

		err = project.ScaffoldProject(absPath, data, projectForce)
		if err != nil {
			log.Fatalf("could not scaffold project: %v", err)
		}

		fmt.Println("Project initialized successfully!")
	},
}

func init() {
	cliCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectInitCmd)

	projectInitCmd.Flags().StringVarP(&projectName, "name", "n", "", "Project name (defaults to directory name)")
	projectInitCmd.Flags().BoolVarP(&projectForce, "force", "f", false, "Overwrite existing files")
}
