package kustomize

// Kustomizer defines common kustomize actions used within Jenkins X
type Kustomizer interface {
	Version(extraArgs ...string) (string, error)
	HasKustomize() bool
	FindKustomize() (resource []string)
	SetCWD(dir string)
	ApplyKustomize() error
}
