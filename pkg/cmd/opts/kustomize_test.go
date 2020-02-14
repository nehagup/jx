package opts

import (
	"github.com/jenkins-x/jx/pkg/versionstream"
	"github.com/stretchr/testify/assert"
	"testing"
)

var versionTests = []struct {
	currentVersion string
	stableVersion  *versionstream.StableVersion
	supported      bool
}{
	{"3.5.1", &versionstream.StableVersion{Version: "3.5.4", UpperLimit: "3.6.0"}, false},
	{"3.5.4", &versionstream.StableVersion{Version: "3.5.4", UpperLimit: "3.6.0"}, true},
	{"3.6.0", &versionstream.StableVersion{Version: "3.5.4", UpperLimit: "3.6.0"}, false},
	{"{3.5.4  2020-01-17T14:28:58+00:00  }", &versionstream.StableVersion{Version: "3.5.4", UpperLimit: "3.6.0"}, true},
	{"{kustomize/3.5.4  2020-01-17T14:28:58+00:00  }", &versionstream.StableVersion{Version: "3.5.4", UpperLimit: "3.6.0"}, true},
}

func Test_isInstalledKustomizeVersionSupported(t *testing.T) {
	for _, versionTest := range versionTests {
		t.Run(versionTest.currentVersion, func(t *testing.T) {
			supported, err := isInstalledKustomizeVersionSupported(versionTest.currentVersion, versionTest.stableVersion)
			assert.NoError(t, err)
			if versionTest.supported {
				assert.True(t, supported, "%s should be a supported version", versionTest.currentVersion)
			} else {
				assert.False(t, supported, "%s should not be a supported version", versionTest.currentVersion)
			}
		})
	}
}
