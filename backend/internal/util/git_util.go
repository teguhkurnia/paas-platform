package util

import (
	"os"

	"github.com/go-git/go-git/v6"
	"github.com/sirupsen/logrus"
)

type GitUtil struct {
	Log *logrus.Logger
}

func NewGitUtil(
	log *logrus.Logger,
) *GitUtil {
	return &GitUtil{
		Log: log,
	}
}

func (u *GitUtil) GitClone(repoURL string, destinationPath string) error {
	u.Log.Debugf("Cloning repository from %s to %s", repoURL, destinationPath)
	_, err := git.PlainClone(destinationPath, &git.CloneOptions{
		URL:               repoURL,
		Progress:          os.Stdout,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		u.Log.Errorf("Failed to clone repository: %v", err)
		return err
	}

	return nil
}
