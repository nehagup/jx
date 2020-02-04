package kustomize

import (
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/util"
	"os"
	"path/filepath"
	"strings"
)

// KustomizeCLI implements common kustomize actions based on kustomize CLI
type KustomizeCLI struct {
	Runner util.Commander
	CWD string
}

// NewKustomizeCLI creates a new KustomizeCLI instance configured to use the provided kustomize CLI in
// the given current working directory
func NewKustomizeCLI(cwd string) *KustomizeCLI {
	runner := &util.Command{
		Name: "kustomize",
	}
	cli := &KustomizeCLI{
		Runner: runner,
		CWD: cwd,
	}

	return cli
}

// Version executes the Kustomize version command and returns its output
func (k *KustomizeCLI) Version(extraArgs ...string) (string, error) {
	args := []string{"version", "--short"}
	args = append(args, extraArgs...)
	return k.runKustomizeWithOutput(args...)
}

func (k *KustomizeCLI) runKustomizeWithOutput(args ...string) (string, error) {
	k.Runner.SetArgs(args)
	return k.Runner.RunWithoutRetry()
}

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
