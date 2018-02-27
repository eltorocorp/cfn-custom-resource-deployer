//go:generate mockgen -destination=generated/mocks/customresource/mock_customresource.go github.com/eltorocorp/cfn-custom-resource-deployer/src/customresources CustomResource

package deployer_test

import (
	"strings"
	"testing"

	"github.com/eltorocorp/cfn-custom-resource-deployer/src/customresources"
	"github.com/eltorocorp/cfn-custom-resource-deployer/src/deployer"
	"github.com/eltorocorp/cfn-custom-resource-deployer/src/deployer/generated/mocks/customresource"
	"github.com/golang/mock/gomock"
)

func Test_New_DuplicatesFound_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	duplicatedName := "Custom::DuplicatedName"
	resourceOne := mock_customresources.NewMockCustomResource(controller)
	resourceOne.EXPECT().ResourceName().Return(&duplicatedName).AnyTimes()
	resourceTwo := mock_customresources.NewMockCustomResource(controller)
	resourceTwo.EXPECT().ResourceName().Return(&duplicatedName).AnyTimes()

	customresources := []customresources.CustomResource{
		resourceOne,
		resourceTwo,
	}
	_, err := deployer.New(customresources)

	if err == nil {
		t.Error("Expected an error, but none was received.")
		t.FailNow()
	}

	if !strings.Contains(err.Error(), "registered more than once") {
		t.Error("Expected error to contain text 'registered more than once', but received '", err.Error(), "'")
		t.FailNow()
	}
}

func Test_New_InvalidResourceNameFound_ReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	invalidName := "Custom::Invalid::Name"
	resource := mock_customresources.NewMockCustomResource(controller)
	resource.EXPECT().ResourceName().Return(&invalidName).AnyTimes()

	customresources := []customresources.CustomResource{resource}
	_, err := deployer.New(customresources)

	if err == nil {
		t.Error("Expected an error, but none was received.")
		t.FailNow()
	}

	expectedError := "resource name Custom::Invalid::Name contains illegal characters"
	if err.Error() != expectedError {
		t.Error("Expected:", expectedError, "Received:", err.Error())
		t.FailNow()
	}
}
