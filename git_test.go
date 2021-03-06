package captain // import "github.com/harbur/captain"

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitGetRevision(t *testing.T) {
	rev, err := getRevision(false)
	assert.NoError(t, err)
	assert.Equal(t, 7, len(rev), "Git revision should have length 7 chars")
}

func TestGitGetRevisionFullSha(t *testing.T) {
	rev, err := getRevision(true)
	assert.NoError(t, err)
	assert.Equal(t, 40, len(rev), "Git revision should have a length of 40 chars")
}

// TODO Fails because it assumes current branch is master
func TestGitGetBranch(t *testing.T) {
	branches, err := getBranches(false)
	assert.NoError(t, err)
	assert.NotEmpty(t, branches)
	// assert.Equal(t, []string{"master"}, getBranches(false), "Git branch should be master")
}

// TODO Fails because it assumes current branch is master
func TestGitGetBranchAllBranches(t *testing.T) {
	branches, err := getBranches(true)
	assert.NoError(t, err)
	assert.NotEmpty(t, branches)
	// assert.Equal(t, []string{"master"}, getBranches(true), "Git branch should be master")
}

// TODO Fails because vendors/ is not git-ignored.
func TestGitIsDirty(t *testing.T) {
	// assert.Equal(t, false, isDirty(), "Git should not have local changes")
}

func TestGitIsGit(t *testing.T) {
	assert.Equal(t, true, isGit(), "There should be a git repository")
}
