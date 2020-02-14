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

	if err == nil && version != "" {
		err := isInstalledKustomizeSupported(o, version)
		return errors.Wrapf(err, "Problem finding if installed version of kustomize is supported")
	}

	err = o.InstallKustomize()
	if err != nil {
		return errors.Wrap(err, "Failed to install Kustomize")
	}

	log.Logger().Info("Installed Kustomize")
	return nil
}

func isInstalledKustomizeSupported(o *CommonOptions, version string) error {
	// get the stable jx supported version of kustomize to be install
	versionResolver, err := o.GetVersionResolver()
	if err != nil {
		return errors.Wrapf(err, "Unable to get version resolver for jenkins-x-versions")
	}

	stableVersion, err := versionResolver.StableVersion(versionstream.KindPackage, "kustomize")
	if err != nil {
		return errors.Wrapf(err, "Unable to get stable version from the jenkins-x-versions for github.com/%s/%s %v ", "kubernetes-sigs", "kustomize", err)
	}

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

	if currVersion.GTE(lowerLimit) && currVersion.LTE(upperLimit) {
		log.Logger().Info("Kustomize is already installed version")
		return nil
	}

	return errors.Wrapf(err, "Unsupported version of Kustomize installed. Install kustomize version above %s or below %s ", lowerLimit, upperLimit)
}
