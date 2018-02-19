package customresources

// CustomResource is a resource to be managed by the deployer.
type CustomResource interface {
	// ActionWasSuccessful returns nil if the Create, Update, or Delete action hasn't been called.
	// Once Create Update or Delete has been called, ActionWasSuccessful should return true
	// if the actiom completed successfully, and false if an error was encountered.
	ActionWasSuccessful() *bool

	//Reason returns nil if an action has not been called or if the action was successful.
	//If an action was called, and not successful, reason will contain a message describing what happened.
	Reason() *string

	// ResourceName returns the name of the resource. ie Custom::SomeResource.
	ResourceName() *string

	// GetNoEcho returns whether or not this resource wants to allow CloudFormation to display its data.
	NoEcho() *bool

	// GetData retuns any data that the resource wishes to emit back to CloudFormation.
	Data() *interface{}

	// Create is called when CloudFormation wishes to create a new instance of a custom resource.
	Create()

	// Update is called when CloudFormation wishes to update an existing instance of a custom resource.
	Update()

	// Delete is called when CloudFormation wishes to delete an existing instance of a custom resrouce.
	Delete()
}
