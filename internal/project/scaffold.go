package project

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

type ProjectData struct {
	ProjectName string
	Description string
}

type fileTemplate struct {
	templateName string
	targetFile   string
}

var filesToGenerate = []fileTemplate{
	{"readme.md.tmpl", "README.md"},
	{"gemini.md.tmpl", "GEMINI.md"},
	{"gitignore.tmpl", ".gitignore"},
	{"pre-commit-config.yaml.tmpl", ".pre-commit-config.yaml"},
	{"markdownlint.yaml.tmpl", ".markdownlint.yaml"},
	{"yamllint.yaml.tmpl", ".yamllint.yaml"},
}

func ScaffoldProject(path string, data ProjectData, force bool) error {
	// Ensure directory exists
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("could not create directory: %w", err)
	}

	// Git Init
	// Check if repository already exists to avoid errors or re-initialization logic if not desired,
	// though PlainInit handles idempotency by returning ErrRepositoryAlreadyExists.
	_, err := git.PlainInit(path, false)
	if err == git.ErrRepositoryAlreadyExists {
		// It's okay if it already exists
	} else if err != nil {
		return fmt.Errorf("could not initialize git repository: %w", err)
	} else {
		fmt.Println("Initialized git repository")
	}

	// Generate files
	for _, ft := range filesToGenerate {
		targetPath := filepath.Join(path, ft.targetFile)

		// Check if exists
		if _, err := os.Stat(targetPath); err == nil && !force {
			fmt.Printf("File %s already exists, skipping...\n", ft.targetFile)
			continue
		}

		if err := generateFile(ft, targetPath, data); err != nil {
			return fmt.Errorf("error generating %s: %w", ft.targetFile, err)
		}
		fmt.Printf("Generated %s\n", ft.targetFile)
	}

	// Initial Commit
	err = initialCommit(path)
	if err != nil {
		return fmt.Errorf("could not create initial commit: %w", err)
	}

	return nil
}

func initialCommit(path string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	// Add all files
	_, err = w.Add(".")
	if err != nil {
		return err
	}

	// Try to load global git config
	var name, email string
	cfg, err := config.LoadConfig(config.GlobalScope)
	if err == nil {
		name = cfg.User.Name
		email = cfg.User.Email
	}

	// Fallbacks
	if name == "" {
		name = "CLI User"
	}
	if email == "" {
		email = "cli@generated.local"
	}

	// Commit
	_, err = w.Commit("feat: init project", &git.CommitOptions{
		Author: &object.Signature{
			Name:  name,
			Email: email,
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Created initial commit: feat: init project (Author: %s <%s>)\n", name, email)
	return nil
}

func generateFile(ft fileTemplate, targetPath string, data ProjectData) error {
	tmpl, err := template.ParseFS(templatesFS, "templates/"+ft.templateName)
	if err != nil {
		return err
	}

	f, err := os.Create(targetPath)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		_ = f.Close()
		return err
	}

	return f.Close()
}
