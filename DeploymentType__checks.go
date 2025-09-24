//go:build !no_runtime_type_checking

package cdkstacksets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDeploymentType_SelfManagedParameters(options *SelfManagedOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDeploymentType_ServiceManagedParameters(options *ServiceManagedOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

