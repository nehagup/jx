package kustomize

import (
	"github.com/jenkins-x/jx/pkg/util"
)

// KustomizeCLI implements common kustomize actions based on kustomize CLI
type KustomizeCLI struct {
	Runner util.Commander
}

// NewKustomizeCLI creates a new KustomizeCLI instance configured to use the provided kustomize CLI in
// the given current working directory
func NewKustomizeCLI() *KustomizeCLI {
	runner := &util.Command{
		Name: "kustomize",
	}
	cli := &KustomizeCLI{
		Runner: runner,
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
