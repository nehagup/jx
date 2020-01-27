package opts

import (
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/pkg/errors"
)

// EnsureKustomize ensures kustomize is installed
func (o *CommonOptions) EnsureKustomize() error {
	_, err := o.Kustomize().Version()
	if err == nil {
		log.Logger().Info("Kustomize is already installed")
		return nil
	}

	err = o.InstallKustomize()
	if err != nil {
		return errors.Wrap(err, "Failed to install Kustomize")
	}

	log.Logger().Info("Installed Kustomize")
	return nil
}
