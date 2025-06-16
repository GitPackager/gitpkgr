package git

import (
	"fmt"
	"strings"

	"github.com/GitPackager/gitpkgr/internals/utils"
)

const (
	DefaultCloneLocation = "/tmp/gitpkgr"
	DefaultBranch        = "main"
)

func CloneRepo(repoURL string, branchName string) error {
}
