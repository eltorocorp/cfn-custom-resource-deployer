package customresources

// CustomResource is a resource to be managed by the deployer.
type CustomResource interface {
	GetResourceName() string
	Create() error
	Update() error
	Delete() error
}
