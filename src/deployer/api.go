package deployer

import (
	"github.com/eltorocorp/cfn-custom-resource-deployer/src/customresources"
	"github.com/eltorocorp/cfn-response/cfnhelper"
)

// API handles registration of custom resources that are available to the deployer, and exposes methods for deploying
// the resources.
type API struct {
	registeredResources []customresources.CustomResource
}

// New registers a set of custom resources with the API, and returns a reference to a new deployer API.
// If more than one resource has the same name, an error is returned.
// If any resource has an invalid name-format (Custom::[name]), an error is returned.
func New(registeredResources []customresources.CustomResource) (*API, error) {
	if err := validateResourceNameFormats(registeredResources); err != nil {
		return nil, err
	}
	if err := checkForDuplicateNames(registeredResources); err != nil {
		return nil, err
	}
	return &API{
		registeredResources: registeredResources,
	}, nil
}

// DeployCustomResource finds the registered resource specified in the request, and attempts to deploy that resource.
//
// If the deployer is not able to find the specified resource in the slice of registered resources, it will return an error.
//
// If the deployment action is successful, the method reports success back to CloudFormation.
// If the deployment action fails, the method will attempt to report failure back to CloudFormation
// and will also return an error to the caller.
func (a *API) DeployCustomResource(request *cfnhelper.Request) error {

}