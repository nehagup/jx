package kustomize

import (
<<<<<<< HEAD
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/util"
	"os"
	"path/filepath"
	"strings"
=======
	"regexp"

	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c
)

// KustomizeCLI implements common kustomize actions based on kustomize CLI
type KustomizeCLI struct {
	Runner util.Commander
<<<<<<< HEAD
	CWD string
=======
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c
}

// NewKustomizeCLI creates a new KustomizeCLI instance configured to use the provided kustomize CLI in
// the given current working directory
<<<<<<< HEAD
func NewKustomizeCLI(cwd string) *KustomizeCLI {
=======
func NewKustomizeCLI() *KustomizeCLI {
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c
	runner := &util.Command{
		Name: "kustomize",
	}
	cli := &KustomizeCLI{
		Runner: runner,
<<<<<<< HEAD
		CWD: cwd,
	}

=======
	}
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c
	return cli
}

// Version executes the Kustomize version command and returns its output
func (k *KustomizeCLI) Version(extraArgs ...string) (string, error) {
	args := []string{"version", "--short"}
	args = append(args, extraArgs...)
<<<<<<< HEAD
	return k.runKustomizeWithOutput(args...)
=======
	version, err := k.runKustomizeWithOutput(args...)
	if err != nil {
		return "", err
	}
	return extractSemanticVersion(version)
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c
}

func (k *KustomizeCLI) runKustomizeWithOutput(args ...string) (string, error) {
	k.Runner.SetArgs(args)
	return k.Runner.RunWithoutRetry()
}

<<<<<<< HEAD
// HasKustomize finds out if there is any kustomize resource in the cwd or subdirectories
func (k *KustomizeCLI) HasKustomize() bool {
	if len(k.FindKustomize()) != 0 {
		return true
	}

	return false
}

// FindKustomize looks for the kustomization.yaml i.e. kustomize resources in present and sub-directories
func (k *KustomizeCLI) FindKustomize() (resource []string) {
	fp, err := filepath.Abs(k.CWD)
	var resources []string
	err = filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "kustomization.yaml"){
			resources = append(resources, path)
		}
		return nil
	})
	if err != nil {
		log.Logger().Errorf("Problem finding kustomize resources %s ", err)
	}
	return resources
}

// SetCWD configures the common working directory of kustomize CLI
func (k *KustomizeCLI) SetCWD(dir string) {
	k.CWD = dir
}

//
func (k *KustomizeCLI) ApplyKustomize() (err error) {

	return
}
=======
// extractSemanticVersion return the semantic version string out of given version cli output.
// currently tested on {Version:3.5.4 GitCommit ....} and {Version:kustomize/v3.5.4 GitCommit: ...}
func extractSemanticVersion(version string) (string, error) {
	regex, err := regexp.Compile(`[0-9]+\.[0-9]+\.[0-9]+`)
	if err != nil {
		return "", errors.Wrapf(err, "not able to extract a semantic version of kustomize version output")
	}
	return regex.FindString(version), nil
}
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c
