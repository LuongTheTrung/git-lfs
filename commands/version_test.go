package commands

import (
	"fmt"
	"testing"
)

func TestVersionOnEmptyRepository(t *testing.T) {
	repo := NewRepository(t, "empty")
	repo.AddPath(repo.Path, ".git")
	repo.AddPath(repo.Path, "subdir")

	cmd := repo.Command("version")
	cmd.Output = fmt.Sprintf("git-media v%s", Version)

	cmd = repo.Command("version", "-comics")
	cmd.Output = fmt.Sprintf("git-media v%s\nNothing may see Gah Lak Tus and survive.", Version)

	repo.Test()
}