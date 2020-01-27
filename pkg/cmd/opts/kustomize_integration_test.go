// +build integration

package opts_test

import (
	"fmt"
	"github.com/jenkins-x/jx/pkg/cmd/testhelpers"
	"github.com/jenkins-x/jx/pkg/packages"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/log"
)

// TestCommonOptions_EnsureKustomize tests that Kustomize is properly installed
func TestCommonOptions_EnsureKustomizeNoBinaryInstalled(t *testing.T) {
	o := opts.NewCommonOptionsWithFactory(clients.NewFactory())
	o.NoBrew = true


	_, jxHome, err := testhelpers.CreateTestJxHomeDir()
	assert.NoError(t, err)


	errK := o.EnsureKustomize()

	// test that tmpJXHome/bin contains kustomize binary
	assert.FileExists(t, filepath.Join(jxHome, "bin", "kustomize"))

	cmd := exec.Command("kustomize", "version")
	version, err := cmd.Output()
	assert.NoError(t, err)
	assert.Contains(t, "4.5.1", version)

	if err == nil {
		if errK != nil {
			t.Errorf("EnsureKustomize() error = %v", errK)
		} else {
			return
		}
	}

	log.Logger().Info(err.Error())
	if errK == nil {
		return
	} else {
		t.Errorf("EnsureKustomize() error = %v", errK)
	}
}

func TestCommonOptions_EnsureKustomizeBinaryExistsWrongVersion(t *testing.T) {
	o := opts.NewCommonOptionsWithFactory(clients.NewFactory())

	origJXHome, testJXHome, err := testhelpers.CreateTestJxHomeDir()
	defer func() {
		err := os.Setenv("JX_HOME", origJXHome)
		log.Logger().Warnf("Unable to reset JX_HOME for test TestCommonOptions_EnsureKustomizeBinaryExistsWrongVersion: %s", err)

		err = os.RemoveAll(testJXHome)
		log.Logger().Warnf("Unable to remove tmp directory %s: %s", testJXHome, err)
	}()
	assert.NoError(t, err)

	// install a non stable version of kustmize
	version := "4.4.0"
	clientURL := fmt.Sprintf("https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%%2Fv%s/kustomize_v%s_%s_%s.tar.gz", version, version, runtime.GOOS, runtime.GOARCH)
	fullPath := filepath.Join(testJXHome, "bin", "kustomize")
	tmpFile := fullPath + ".tmp"
	err = packages.DownloadFile(clientURL, tmpFile)
	assert.NoError(t, err)
	err = util.RenameFile(tmpFile, fullPath)
	assert.NoError(t, err)
	os.Chmod(fullPath, 0755)
	assert.NoError(t, err)

	errK := o.EnsureKustomize()

	// test that tmpJXHome/bin contains kustomize binary

	if err == nil {
		if errK != nil {
			t.Errorf("EnsureKustomize() error = %v", errK)
		} else {
			return
		}
	}

	if errK == nil {
		return
	} else {
		t.Errorf("EnsureKustomize() error = %v", errK)
	}
}

