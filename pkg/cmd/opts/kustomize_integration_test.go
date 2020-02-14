// +build integration

package opts_test

import (
	"path/filepath"
	"testing"

	"github.com/jenkins-x/jx/pkg/cmd/testhelpers"
	"github.com/jenkins-x/jx/pkg/versionstream"
	"github.com/stretchr/testify/assert"

	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/log"
)

// TestCommonOptions_EnsureKustomize tests that Kustomize is properly installed
func TestCommonOptions_EnsureKustomizeNoBinaryInstalled(t *testing.T) {
	o := opts.NewCommonOptionsWithFactory(clients.NewFactory())
	o.NoBrew = true
	version, _ := o.Kustomize().Version()
	if version != "" {
		return
	}

	origJXHome, testJXHome, err := testhelpers.CreateTestJxHomeDir()
	assert.NoError(t, err, "Failed to create a test JX Home directory")

	defer func() {
		err := testhelpers.CleanupTestJxHomeDir(origJXHome, testJXHome)
		if err != nil {
			log.Logger().Warnf("Unable to remove temporary directory %s: %s", testJXHome, err)
		}
	}()

	err = o.EnsureKustomize()
	assert.NoError(t, err, "EnsureKustomize() error for test case TestCommonOptions_EnsureKustomizeNoBinaryInstalled")

	// test that tmpJXHome/bin contains kustomize binary
	version, err = o.Kustomize().Version()

	assert.FileExists(t, filepath.Join(testJXHome, "bin", "kustomize"))
	assert.NoError(t, err, "Kustomize was not installed in the temp dir of test TestCommonOptions_EnsureKustomizeNoBinaryInstalled %s: %s", testJXHome)

	// get the stable jx supported version of kustomize to be install
	versionResolver, err := o.GetVersionResolver()
	if err != nil {
		log.Logger().Warnf("Unable to get version resolver for jenkins-x-versions %s", err)
	}

	stableVersion, err := versionResolver.StableVersion(versionstream.KindPackage, "kustomize")
	if err != nil {
		log.Logger().Warnf("Unable to get stable version from the jenkins-x-versions for github.com/%s/%s %v ", "kubernetes-sigs", "kustomize", err)
	}

	assert.Contains(t, version, stableVersion.Version)
}
