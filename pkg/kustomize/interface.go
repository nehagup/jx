package kustomize

// Kustomizer defines common kustomize actions used within Jenkins X
type Kustomizer interface {
	Version(extraArgs ...string) (string, error)
<<<<<<< HEAD
	HasKustomize() bool
	FindKustomize() (resource []string)
	SetCWD(dir string)
	ApplyKustomize() error
=======
>>>>>>> ce287f89a20c832068d763f94f9f8c94c1b6696c
}
