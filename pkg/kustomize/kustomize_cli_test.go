package kustomize_test

import (
	"github.com/jenkins-x/jx/pkg/cmd/testhelpers"
	"github.com/jenkins-x/jx/pkg/kustomize"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/util"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestKustomizeCLI_FindKustomize_HasKustomize(t *testing.T) {
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
	err = os.MkdirAll(filepath.Join(testJXHome, "bin/base/charts"), os.ModePerm)
	err = os.MkdirAll(filepath.Join(testJXHome, "bin/base/templates"), os.ModePerm)
	err = os.MkdirAll(filepath.Join(testJXHome, "bin/staging"), os.ModePerm)
	_, err = os.Create(filepath.Join(testJXHome, "bin/base/Charts.yaml"))
	_, err = os.Create(filepath.Join(testJXHome, "bin/base/values.yaml"))
	_, err = os.Create(filepath.Join(testJXHome, "bin/base/kustomization.yaml"))
	_, err = os.Create(filepath.Join(testJXHome, "bin/staging/values.yaml"))
	_, err = os.Create(filepath.Join(testJXHome, "bin/staging/kustomization.yaml"))
	_, err = os.Create(filepath.Join(testJXHome, "bin/base/charts/kustomization.yaml"))
	if err != nil {
		t.Errorf("Error creating a test kustomize directory in TestKustomizeCLI_FindKustomize %s", err)
	}

	jxBin, err := util.JXBinLocation()
	if err != nil {
		t.Errorf("Error finding JXBin in TestKustomizeCLI_FindKustomize %s", err)
	}

	k := kustomize.NewKustomizeCLI(jxBin)
	wantedOutput := []string{filepath.Join(jxBin, "base/charts/kustomization.yaml"),
		filepath.Join(jxBin,"base/kustomization.yaml"), filepath.Join(jxBin,"staging/kustomization.yaml")}
	output := k.FindKustomize()

	if !reflect.DeepEqual(wantedOutput, output){
		t.Errorf("Not able to find all of the kustomize resource %s", err)
	}

	if !k.HasKustomize() {
		t.Errorf("Failed to find presence of kustomize")
	}

	k.SetCWD(filepath.Join(jxBin, "base/templates"))
	if k.HasKustomize() {
		t.Errorf("Failed to find presence of kustomize")
	}
}
