// +build integration

package opts_test

import (
<<<<<<< HEAD
	"github.com/blang/semver"
	"github.com/jenkins-x/jx/pkg/cmd/testhelpers"
	"github.com/jenkins-x/jx/pkg/versionstream"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
=======
	"os"
	"path/filepath"
	"testing"

	"github.com/jenkins-x/jx/pkg/cmd/testhelpers"
	"github.com/jenkins-x/jx/pkg/versionstream"
	"github.com/stretchr/testify/assert"
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c

	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/log"
)

<<<<<<< HEAD
// TestCommonOptions_EnsureKustomize tests that Kustomize is properly installed
func TestCommonOptions_EnsureKustomizeNoBinaryInstalled(t *testing.T) {
	o := opts.NewCommonOptionsWithFactory(clients.NewFactory())
	o.NoBrew = true
	version, _ := o.Kustomize().Version()
	if version != "" {
		return
	}

	origJXHome, testJXHome, err := testhelpers.CreateTestJxHomeDir()
	if err != nil {
		t.Errorf("Failed to create a test JX Home directory")
	}
=======
// TestInstallKustomize tests that Kustomize gets properly installed into JX_HOME
func TestInstallKustomize(t *testing.T) {
	origJXHome, testJXHome, err := testhelpers.CreateTestJxHomeDir()
	assert.NoError(t, err, "failed to create a test JX Home directory")
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c

	defer func() {
		err := testhelpers.CleanupTestJxHomeDir(origJXHome, testJXHome)
		if err != nil {
<<<<<<< HEAD
			log.Logger().Warnf("Unable to remove temporary directory %s: %s", testJXHome, err)
		}
	}()

	err = o.EnsureKustomize()
	if err != nil {
		t.Errorf("EnsureKustomize() error for test case TestCommonOptions_EnsureKustomizeNoBinaryInstalled = %v", err)
	}
	// test that tmpJXHome/bin contains kustomize binary
	version, err = o.Kustomize().Version()

	if !assert.FileExists(t, filepath.Join(testJXHome, "bin", "kustomize")) || err != nil {
		t.Errorf("Kustomize was not installed in the temp dir of test TestCommonOptions_EnsureKustomizeNoBinaryInstalled %s: %s", testJXHome, err)
	}

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

func TestCommonOptions_EnsureKustomizeBinaryExistsWrongVersion(t *testing.T) {
	o := opts.NewCommonOptionsWithFactory(clients.NewFactory())
	version, _ := o.Kustomize().Version()
	if version == "" {
		return
	}

	origJXHome, testJXHome, err := testhelpers.CreateTestJxHomeDir()
	if err != nil {
		t.Errorf("Failed to create a test JX Home directory")
	}

	defer func() {
		err := testhelpers.CleanupTestJxHomeDir(origJXHome, testJXHome)
		if err != nil {
			log.Logger().Warnf("Unable to remove temporary directory %s: %s", testJXHome, err)
		}
	}()

	// get the stable jx supported version of kustomize to be install
	versionResolver, err := o.GetVersionResolver()
	if err != nil {
		log.Logger().Warnf("Unable to get version resolver for jenkins-x-versions %s", err)
	}

	stableVersion, err := versionResolver.StableVersion(versionstream.KindPackage, "kustomize")
	if err != nil {
		log.Logger().Warnf("Unable to get stable version from the jenkins-x-versions for github.com/%s/%s %v ", "kubernetes-sigs", "kustomize", err)
	}

	currVersion, err := semver.Make(version[12:17])
	if err != nil {
		log.Logger().Warnf("Unable to get currently installed Kustomize sem-version %s", err)
	}
	lowerLimit, err := semver.Make(stableVersion.Version)
	if err != nil {
		log.Logger().Warnf("Unable to get lowest supported stable Kustomize sem-version %s", err)
	}
	upperLimit, err := semver.Make(stableVersion.UpperLimit)
	if err != nil {
		log.Logger().Warnf("Unable to get highest supported stable Kustomize sem-version %s", err)
	}

	err = o.EnsureKustomize()
	if (currVersion.LT(lowerLimit) || currVersion.GT(upperLimit)) && err == nil {
		t.Errorf("TestCommonOptions_EnsureKustomizeBinaryExistsWrongVersion Failed")
	}
=======
			log.Logger().Warnf("unable to remove temporary directory %s: %s", testJXHome, err)
		}
	}()

	o := opts.NewCommonOptionsWithFactory(clients.NewFactory())
	o.NoBrew = true

	err = o.InstallKustomize()
	assert.NoError(t, err, "error installing kustomize")

	// test that tmpJXHome/bin contains kustomize binary and retrieve its version
	// unset PATH to ensure no unwanted binary will be used
	origPath := os.Getenv("PATH")
	err = os.Setenv("PATH", "")
	assert.NoError(t, err, "unable to temporarily unset PATH: %s", err)

	version, err := o.Kustomize().Version()
	assert.FileExists(t, filepath.Join(testJXHome, "bin", "kustomize"))
	assert.NoError(t, err, "kustomize was not installed in the temp JX_HOME %s", testJXHome)

	err = os.Setenv("PATH", origPath)
	assert.NoError(t, err, "unable to reset PATH: %s", err)

	// get the stable jx supported version of kustomize to be install
	versionResolver, err := o.GetVersionResolver()
	assert.NoError(t, err, "unable to retrieve version resolver")

	stableVersion, err := versionResolver.StableVersion(versionstream.KindPackage, "kustomize")
	assert.NoError(t, err, "unable to retrieve stable version for kustomize from version resolver")
	assert.Contains(t, version, stableVersion.Version)
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c
}
