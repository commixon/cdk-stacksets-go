//go:build !no_runtime_type_checking

package cdkstacksets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

func (s *jsiiProxy_StackSetStackSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	if requiredVersion == nil {
		return fmt.Errorf("parameter requiredVersion is required, but nil was provided")
	}

	if bootstrapStackVersionSsmParameter == nil {
		return fmt.Errorf("parameter bootstrapStackVersionSsmParameter is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateAddDockerImageAssetParameters(_asset *awscdk.DockerImageAssetSource) error {
	if _asset == nil {
		return fmt.Errorf("parameter _asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_asset, func() string { return "parameter _asset" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateAddFileAssetParameters(asset *awscdk.FileAssetSource) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(asset, func() string { return "parameter asset" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateBindParameters(stack awscdk.Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(dest, func() string { return "parameter dest" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	if location == nil {
		return fmt.Errorf("parameter location is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(location, func() string { return "parameter location" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateEmitArtifactParameters(session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateEmitStackArtifactParameters(stack awscdk.Stack, session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateSynthesizeStackTemplateParameters(stack awscdk.Stack, session awscdk.ISynthesisSession) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StackSetStackSynthesizer) validateSynthesizeTemplateParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

