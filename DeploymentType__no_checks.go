//go:build no_runtime_type_checking

package cdkstacksets

// Building without runtime type checking enabled, so all the below just return nil

func validateDeploymentType_SelfManagedParameters(options *SelfManagedOptions) error {
	return nil
}

func validateDeploymentType_ServiceManagedParameters(options *ServiceManagedOptions) error {
	return nil
}

