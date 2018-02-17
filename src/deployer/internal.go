package deployer

import (
	"fmt"
	"regexp"

	"github.com/deckarep/golang-set"
	"github.com/eltorocorp/cfn-custom-resource-deployer/src/customresources"
)

func validateResourceNameFormats(resources []customresources.CustomResource) error {
	// must be Custom:: followed by word characters
	mustMatchThis := regexp.MustCompile(`^(Custom::)\w+`)
	// must not be Custom:: followed by anything containing non-word characters
	mustNotMatchThis := regexp.MustCompile(`^(Custom::).*\W`)
	for _, resource := range resources {
		resourceName := resource.GetResourceName()
		if mustMatchThis.FindStringIndex(resourceName) == nil {
			return fmt.Errorf("resource %v has an invalid name", resourceName)
		}
		if mustNotMatchThis.FindStringIndex(resourceName) != nil {
			return fmt.Errorf("resource name %v contains illegal characters", resourceName)
		}
	}
}

func checkForDuplicateNames(resources []customresources.CustomResource) error {
	resourceNames := mapset.NewSet()
	for _, resource := range resources {
		if resourceNames.Add(resource.GetResourceName()) == false {
			return fmt.Errorf("resource %v was registered more than once", resource.GetResourceName())
		}
	}
	return nil
}
