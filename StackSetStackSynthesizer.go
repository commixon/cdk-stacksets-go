package cdkstacksets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/commixon/cdk-stacksets-go/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/commixon/cdk-stacksets-go/internal"
)

// Deployment environment for an AWS StackSet stack.
//
// Interoperates with the StackSynthesizer of the parent stack.
// Experimental.
type StackSetStackSynthesizer interface {
	awscdk.StackSynthesizer
	// The common prefix for the asset bucket names used by this StackSetStack.
	//
	// Required if `assetBuckets` is provided.
	// Default: - No Buckets provided and Assets will not be supported.
	//
	// Experimental.
	AssetBucketPrefix() *string
	// An array of Buckets can be passed to store assets, enabling StackSetStack Asset support.
	//
	// One Bucket is required per target region. The name must be `${assetBucketPrefix}-<region>`, where
	// `<region>` is the region targeted by the StackSet.
	// Default: - No Buckets provided and Assets will not be supported.
	//
	// Experimental.
	AssetBuckets() *[]awss3.IBucket
	// The qualifier used to bootstrap this stack.
	// Experimental.
	BootstrapQualifier() *string
	// Retrieve the bound stack.
	//
	// Fails if the stack hasn't been bound yet.
	// Experimental.
	BoundStack() awscdk.Stack
	// The role used to lookup for this stack.
	// Experimental.
	LookupRole() *string
	// Add a CfnRule to the bound stack that checks whether an SSM parameter exceeds a given version.
	//
	// This will modify the template, so must be called before the stack is synthesized.
	// Experimental.
	AddBootstrapVersionRule(requiredVersion *float64, bootstrapStackVersionSsmParameter *string)
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	//
	// The synthesizer must rely on some out-of-band mechanism to make sure the given files
	// are actually placed in the returned location before the deployment happens. This can
	// be by writing the instructions to the asset manifest (for use by the `cdk-assets` tool),
	// by relying on the CLI to upload files (legacy behavior), or some other operator controlled
	// mechanism.
	// Experimental.
	AddDockerImageAsset(_asset *awscdk.DockerImageAssetSource) *awscdk.DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	//
	// The synthesizer must rely on some out-of-band mechanism to make sure the given files
	// are actually placed in the returned location before the deployment happens. This can
	// be by writing the instructions to the asset manifest (for use by the `cdk-assets` tool),
	// by relying on the CLI to upload files (legacy behavior), or some other operator controlled
	// mechanism.
	// Experimental.
	AddFileAsset(asset *awscdk.FileAssetSource) *awscdk.FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	// Experimental.
	Bind(stack awscdk.Stack)
	// Turn a docker asset location into a CloudFormation representation of that location.
	//
	// If any of the fields contain placeholders, the result will be wrapped in a `Fn.sub`.
	// Experimental.
	CloudFormationLocationFromDockerImageAsset(dest *cloudassemblyschema.DockerImageDestination) *awscdk.DockerImageAssetLocation
	// Turn a file asset location into a CloudFormation representation of that location.
	//
	// If any of the fields contain placeholders, the result will be wrapped in a `Fn.sub`.
	// Experimental.
	CloudFormationLocationFromFileAsset(location *cloudassemblyschema.FileDestination) *awscdk.FileAssetLocation
	// Write the CloudFormation stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session. The Stack artifact will control the settings for the
	// CloudFormation deployment.
	// Experimental.
	EmitArtifact(session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	// Deprecated: Use `emitArtifact` instead.
	EmitStackArtifact(stack awscdk.Stack, session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Have the stack write out its template.
	// Deprecated: Use `synthesizeTemplate` instead.
	SynthesizeStackTemplate(stack awscdk.Stack, session awscdk.ISynthesisSession)
	// Write the stack template to the given session.
	//
	// Return a descriptor that represents the stack template as a file asset
	// source, for adding to an asset manifest (if desired). This can be used to
	// have the asset manifest system (`cdk-assets`) upload the template to S3
	// using the appropriate role, so that afterwards only a CloudFormation
	// deployment is necessary.
	//
	// If the template is uploaded as an asset, the `stackTemplateAssetObjectUrl`
	// property should be set when calling `emitArtifact.`
	//
	// If the template is *NOT* uploaded as an asset first and the template turns
	// out to be >50KB, it will need to be uploaded to S3 anyway. At that point
	// the credentials will be the same identity that is doing the `UpdateStack`
	// call, which may not have the right permissions to write to S3.
	// Experimental.
	// SynthesizeTemplate(session awscdk.ISynthesisSession, lookupRoleArn *string) *awscdk.FileAssetSource
}

// The jsii proxy struct for StackSetStackSynthesizer
type jsiiProxy_StackSetStackSynthesizer struct {
	internal.Type__awscdkStackSynthesizer
}

func (j *jsiiProxy_StackSetStackSynthesizer) AssetBucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetBucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackSetStackSynthesizer) AssetBuckets() *[]awss3.IBucket {
	var returns *[]awss3.IBucket
	_jsii_.Get(
		j,
		"assetBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackSetStackSynthesizer) BootstrapQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackSetStackSynthesizer) BoundStack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"boundStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackSetStackSynthesizer) LookupRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookupRole",
		&returns,
	)
	return returns
}

// Creates a new StackSetStackSynthesizer.
// Experimental.
func NewStackSetStackSynthesizer(assetBuckets *[]awss3.IBucket, assetBucketPrefix *string) StackSetStackSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_StackSetStackSynthesizer{}

	_jsii_.Create(
		"cdk-stacksets.StackSetStackSynthesizer",
		[]interface{}{assetBuckets, assetBucketPrefix},
		&j,
	)

	return &j
}

// Creates a new StackSetStackSynthesizer.
// Experimental.
func NewStackSetStackSynthesizer_Override(s StackSetStackSynthesizer, assetBuckets *[]awss3.IBucket, assetBucketPrefix *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-stacksets.StackSetStackSynthesizer",
		[]interface{}{assetBuckets, assetBucketPrefix},
		s,
	)
}

func (s *jsiiProxy_StackSetStackSynthesizer) AddBootstrapVersionRule(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) {
	if err := s.validateAddBootstrapVersionRuleParameters(requiredVersion, bootstrapStackVersionSsmParameter); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addBootstrapVersionRule",
		[]interface{}{requiredVersion, bootstrapStackVersionSsmParameter},
	)
}

func (s *jsiiProxy_StackSetStackSynthesizer) AddDockerImageAsset(_asset *awscdk.DockerImageAssetSource) *awscdk.DockerImageAssetLocation {
	if err := s.validateAddDockerImageAssetParameters(_asset); err != nil {
		panic(err)
	}
	var returns *awscdk.DockerImageAssetLocation

	_jsii_.Invoke(
		s,
		"addDockerImageAsset",
		[]interface{}{_asset},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSetStackSynthesizer) AddFileAsset(asset *awscdk.FileAssetSource) *awscdk.FileAssetLocation {
	if err := s.validateAddFileAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns *awscdk.FileAssetLocation

	_jsii_.Invoke(
		s,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSetStackSynthesizer) Bind(stack awscdk.Stack) {
	if err := s.validateBindParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{stack},
	)
}

func (s *jsiiProxy_StackSetStackSynthesizer) CloudFormationLocationFromDockerImageAsset(dest *cloudassemblyschema.DockerImageDestination) *awscdk.DockerImageAssetLocation {
	if err := s.validateCloudFormationLocationFromDockerImageAssetParameters(dest); err != nil {
		panic(err)
	}
	var returns *awscdk.DockerImageAssetLocation

	_jsii_.Invoke(
		s,
		"cloudFormationLocationFromDockerImageAsset",
		[]interface{}{dest},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSetStackSynthesizer) CloudFormationLocationFromFileAsset(location *cloudassemblyschema.FileDestination) *awscdk.FileAssetLocation {
	if err := s.validateCloudFormationLocationFromFileAssetParameters(location); err != nil {
		panic(err)
	}
	var returns *awscdk.FileAssetLocation

	_jsii_.Invoke(
		s,
		"cloudFormationLocationFromFileAsset",
		[]interface{}{location},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackSetStackSynthesizer) EmitArtifact(session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) {
	if err := s.validateEmitArtifactParameters(session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"emitArtifact",
		[]interface{}{session, options},
	)
}

func (s *jsiiProxy_StackSetStackSynthesizer) EmitStackArtifact(stack awscdk.Stack, session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) {
	if err := s.validateEmitStackArtifactParameters(stack, session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (s *jsiiProxy_StackSetStackSynthesizer) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StackSetStackSynthesizer) SynthesizeStackTemplate(stack awscdk.Stack, session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeStackTemplateParameters(stack, session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

// func (s *jsiiProxy_StackSetStackSynthesizer) SynthesizeTemplate(session awscdk.ISynthesisSession, lookupRoleArn *string) *awscdk.FileAssetSource {
// 	if err := s.validateSynthesizeTemplateParameters(session); err != nil {
// 		panic(err)
// 	}
// 	var returns *awscdk.FileAssetSource
//
// 	_jsii_.Invoke(
// 		s,
// 		"synthesizeTemplate",
// 		[]interface{}{session, lookupRoleArn},
// 		&returns,
// 	)
//
// 	return returns
// }
