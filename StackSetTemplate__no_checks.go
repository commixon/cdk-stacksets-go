//go:build no_runtime_type_checking

package cdkstacksets

// Building without runtime type checking enabled, so all the below just return nil

func validateStackSetTemplate_FromStackSetStackParameters(stack StackSetStack) error {
	return nil
}

