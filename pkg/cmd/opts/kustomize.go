package opts

import (
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/pkg/errors"
)

// EnsureKustomize ensures kustomize is installed
func (o *CommonOptions) EnsureKustomize(installationBinDir ...string) error {
	_, err := o.Kustomize().Version()
	if err == nil {
		log.Logger().Info("Kustomize is already installed")
		return nil
	}

	binDir := ""
	if len(installationBinDir) != 0 {
		binDir = installationBinDir[0]
	}

	err = o.InstallKustomize(binDir)
	if err != nil {
		return errors.Wrap(err, "Failed to install Kustomize")
	}

	log.Logger().Info("Installed Kustomize")
	return nil
}
