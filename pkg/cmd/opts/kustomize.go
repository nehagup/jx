package opts

import (
	"regexp"

	"github.com/blang/semver"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/versionstream"
	"github.com/pkg/errors"
)

// EnsureKustomize ensures kustomize is installed
func (o *CommonOptions) EnsureKustomize() error {
	version, err := o.Kustomize().Version()

	// if there is an error we assume that the binary is not installed
	if err != nil {
		return o.InstallKustomize()
	}

	stableVersion, err := o.stableKustomizeVersion()
	if err != nil {
		return errors.Wrap(err, "failed to install Kustomize")
	}

	supported, err := isInstalledKustomizeVersionSupported(version, stableVersion)
	if err != nil {
		return errors.Wrapf(err, "problem finding if installed version of kustomize is supported")
	}

	if !supported {
		err = o.InstallKustomize()
		if err != nil {
			return errors.Wrap(err, "failed to install Kustomize")
		}
		log.Logger().Info("Installed Kustomize")
	}

	log.Logger().Infof("Keeping currently installed Kustomize version: %s", version)
	return nil
}

func (o *CommonOptions) stableKustomizeVersion() (*versionstream.StableVersion, error) {
	versionResolver, err := o.GetVersionResolver()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get version resolver for jenkins-x-versions")
	}

	// get the stable jx supported version of kustomize to be install
	stableVersion, err := versionResolver.StableVersion(versionstream.KindPackage, "kustomize")
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get stable version from the jenkins-x-versions for github.com/%s/%s %v ", "kubernetes-sigs", "kustomize", err)
	}
	return stableVersion, nil
}

func isInstalledKustomizeVersionSupported(version string, stableVersion *versionstream.StableVersion) (bool, error) {
	regex := regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+`)
	currVersion, err := semver.Make(regex.FindString(version))
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

	if currVersion.GTE(lowerLimit) && currVersion.LT(upperLimit) {
		log.Logger().Debugf("Kustomize is already installed version")
		return true, nil
	}

	return false, errors.Wrapf(err, "unsupported version of Kustomize installed. Install kustomize version above %s or below %s ", lowerLimit, upperLimit)
}
