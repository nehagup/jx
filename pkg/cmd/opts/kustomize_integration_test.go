// +build integration

package opts_test

import (
	"os/exec"
	"runtime"
	"testing"

	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/log"
)

// TestCommonOptions_EnsureKustomize tests that Kustomize is properly installed
func TestCommonOptions_EnsureKustomize(t *testing.T) {
	o := opts.NewCommonOptionsWithFactory(clients.NewFactory())

	cmd := exec.Command("kustomize", "version")
	_, err := cmd.Output()
	errK := o.EnsureKustomize()

	if err == nil {
		if errK != nil {
			t.Errorf("EnsureKustomize() error = %v", errK)
		} else {
			return
		}
	}
	defer func() {
		if runtime.GOOS == "darwin" && !o.NoBrew {
			err = o.RunCommand("brew", "uninstall", "kustomize")
			log.Logger().Error("Error uninstalling kustomize")
		}
	}()

	log.Logger().Info(err.Error())
	if errK == nil {
		return
	} else {
		t.Errorf("EnsureKustomize() error = %v", errK)
	}
}
