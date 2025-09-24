//go:build no_runtime_type_checking

package cdkstacksets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StackSet) validateAddTargetParameters(target StackSetTarget) error {
	return nil
}

func (s *jsiiProxy_StackSet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_StackSet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_StackSet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateStackSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStackSet_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateStackSet_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewStackSetParameters(scope constructs.Construct, id *string, props *StackSetProps) error {
	return nil
}

