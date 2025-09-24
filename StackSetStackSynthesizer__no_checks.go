//go:build no_runtime_type_checking

package cdkstacksets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StackSetStackSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateAddDockerImageAssetParameters(_asset *awscdk.DockerImageAssetSource) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateAddFileAssetParameters(asset *awscdk.FileAssetSource) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateBindParameters(stack awscdk.Stack) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateEmitArtifactParameters(session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateEmitStackArtifactParameters(stack awscdk.Stack, session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateSynthesizeStackTemplateParameters(stack awscdk.Stack, session awscdk.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateSynthesizeTemplateParameters(session awscdk.ISynthesisSession) error {
	return nil
}

