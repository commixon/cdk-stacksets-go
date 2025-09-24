//go:build no_runtime_type_checking

package cdkstacksets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StackSetStack) validateAddDependencyParameters(target awscdk.Stack) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateAddTransformParameters(transform *string) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateAllocateLogicalIdParameters(cfnElement awscdk.CfnElement) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateExportStringListValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateExportValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateFormatArnParameters(components *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateGetLogicalIdParameters(element awscdk.CfnElement) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateRegionalFactParameters(factName *string) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateResolveParameters(obj interface{}) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateSplitArnParameters(arn *string, arnFormat awscdk.ArnFormat) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateToJsonStringParameters(obj interface{}) error {
	return nil
}

func (s *jsiiProxy_StackSetStack) validateToYamlStringParameters(obj interface{}) error {
	return nil
}

func validateStackSetStack_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStackSetStack_IsStackParameters(x interface{}) error {
	return nil
}

func validateStackSetStack_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_StackSetStack) validateSetTerminationProtectionParameters(val *bool) error {
	return nil
}

func validateNewStackSetStackParameters(scope constructs.Construct, id *string, props *StackSetStackProps) error {
	return nil
}

