package git

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/viper"
)

var (
	hiRed     = color.New(color.FgHiRed).SprintFunc()
	yellow    = color.New(color.FgYellow).SprintFunc()
	cyan      = color.New(color.FgCyan).SprintFunc()
	green     = color.New(color.FgGreen).SprintFunc()
	red       = color.New(color.FgRed).SprintFunc()
	blue      = color.New(color.FgBlue).SprintFunc()
	boldGreen = color.New(color.FgGreen, color.Bold).SprintFunc()
)

func GetDevStatus(showChange, showBranch, showAllBranches bool) {
	m := viper.GetString("mainPath")
	viperToCheck := viper.Get("ToCheck")
	t, _ := viperToCheck.([]any)

	for _, i := range t {
		toCheck := i.(map[string]any)

		if toCheck["is_repo"] == true {
			path := m + "/" + toCheck["path"].(string)
			getRepoStatus(path, showChange)
			continue
		}

		path := m + "/" + toCheck["path"].(string)
		dir, err := os.ReadDir(path)
		if err != nil {
			log.Printf("error during read Directory at path %v with error: %v", path, err)
			continue
		}

		for _, d := range dir {
			dirName := d.Name()
			fullPath := path + "/" + dirName
			getRepoStatus(fullPath, showChange)
			listLocalBranch(fullPath, showBranch, showAllBranches)
		}
	}
}

func getRepoStatus(repoPath string, verbose bool) {
	repo, err := git.PlainOpen(repoPath)

	fmt.Println(repoPath)
	if err != nil {
		//fmt.Printf("%v is not a git repository \n\n", repoPath)
		return
	}
	w, err := repo.Worktree()
	if err != nil {
		log.Printf("error during open git Worktree %v with error %v \n\n", repoPath, err)
	}

	status, err := w.Status()
	if err != nil {
		log.Printf("error during  git status Worktree %v with error %v \n\n", repoPath, err)
	}

	if status.IsClean() {
		fmt.Println(boldGreen("The repo is clean"))
		fmt.Println()
	} else {
		fmt.Println(hiRed("The repo is not clean"))
		fmt.Println()
	}
	if verbose {
		for k, v := range status {
			switch {
			case v.Worktree == git.StatusCode('?'):
				r := string(v.Worktree) + " " + k
				fmt.Println(yellow(r))
			case v.Worktree == git.StatusCode('M'):
				r := string(v.Worktree) + " " + k
				fmt.Println(cyan(r))
			case v.Worktree == git.StatusCode('A'):
				r := string(v.Worktree) + " " + k
				fmt.Println(green(r))
			case v.Worktree == git.StatusCode('D'):
				r := string(v.Worktree) + " " + k
				fmt.Println(red(r))
			case v.Worktree == git.StatusCode('R'):
				r := string(v.Worktree) + " " + k
				fmt.Println(yellow(r))
			case v.Worktree == git.StatusCode('C'):
				r := string(v.Worktree) + " " + k
				fmt.Println(yellow(r))
			case v.Worktree == git.StatusCode('U'):
				r := string(v.Worktree) + " " + k
				fmt.Println(hiRed(r))
			case v.Worktree == git.StatusCode(' '):
				r := string(v.Worktree) + " " + k
				fmt.Println(blue(r))
			}
		}
		fmt.Println()
	}
}

func listLocalBranch(repoPath string, showBranch, showAllBranches bool) {

	repo, err := git.PlainOpen(repoPath)

	if err != nil {
		log.Printf("%v is not a git repository \n\n", repoPath)
		return
	}
	head, err := repo.Head()
	if err != nil {
		log.Println("error to get HEAD")
	}
	if showBranch {
		fmt.Println(green(fmt.Sprintf("Attach on branch: %v", head.Name().String())))
	}

	refs, err := repo.References()
	if err != nil {
		log.Println("error to get ref")
	}

	var remote []string
	var local []string

	err = refs.ForEach(func(r *plumbing.Reference) error {
		if r.Name().IsRemote() {
			b := strings.Split(string(r.Name().Short()), "/")
			if b[1] != "HEAD" {
				remote = append(remote, b[1])
			}
		} else if r.Name().IsBranch() {
			local = append(local, string(r.Name().Short()))
		}
		return nil
	})
	if err != nil {
		log.Println("error during refs.ForEach:", err)
	}

	if showAllBranches {
		fmt.Printf("Remote: %v\n", remote)
		fmt.Printf("Local: %v\n", local)
	}
	fmt.Println()
}
