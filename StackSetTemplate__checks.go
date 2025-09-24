//go:build !no_runtime_type_checking

package cdkstacksets

import (
	"fmt"
)

func validateStackSetTemplate_FromStackSetStackParameters(stack StackSetStack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

