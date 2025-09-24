// cdk-stacksets
package cdkstacksets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-stacksets.AccountsTargetOptions",
		reflect.TypeOf((*AccountsTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-stacksets.Capability",
		reflect.TypeOf((*Capability)(nil)).Elem(),
		map[string]interface{}{
			"NAMED_IAM": Capability_NAMED_IAM,
			"IAM": Capability_IAM,
			"AUTO_EXPAND": Capability_AUTO_EXPAND,
		},
	)
	_jsii_.RegisterClass(
		"cdk-stacksets.DeploymentType",
		reflect.TypeOf((*DeploymentType)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DeploymentType{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-stacksets.IStackSet",
		reflect.TypeOf((*IStackSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IStackSet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-stacksets.OperationPreferences",
		reflect.TypeOf((*OperationPreferences)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-stacksets.OrganizationsTargetOptions",
		reflect.TypeOf((*OrganizationsTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-stacksets.RegionConcurrencyType",
		reflect.TypeOf((*RegionConcurrencyType)(nil)).Elem(),
		map[string]interface{}{
			"SEQUENTIAL": RegionConcurrencyType_SEQUENTIAL,
			"PARALLEL": RegionConcurrencyType_PARALLEL,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-stacksets.SelfManagedOptions",
		reflect.TypeOf((*SelfManagedOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-stacksets.ServiceManagedOptions",
		reflect.TypeOf((*ServiceManagedOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-stacksets.StackSet",
		reflect.TypeOf((*StackSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTarget", GoMethod: "AddTarget"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StackSet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStackSet)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-stacksets.StackSetProps",
		reflect.TypeOf((*StackSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-stacksets.StackSetStack",
		reflect.TypeOf((*StackSetStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addTransform", GoMethod: "AddTransform"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "bundlingRequired", GoGetter: "BundlingRequired"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "exportStringListValue", GoMethod: "ExportStringListValue"},
			_jsii_.MemberMethod{JsiiMethod: "exportValue", GoMethod: "ExportValue"},
			_jsii_.MemberMethod{JsiiMethod: "formatArn", GoMethod: "FormatArn"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "nested", GoGetter: "Nested"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackParent", GoGetter: "NestedStackParent"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackResource", GoGetter: "NestedStackResource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "regionalFact", GoMethod: "RegionalFact"},
			_jsii_.MemberMethod{JsiiMethod: "renameLogicalId", GoMethod: "RenameLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContextKey", GoMethod: "ReportMissingContextKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "splitArn", GoMethod: "SplitArn"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateFile", GoGetter: "TemplateFile"},
			_jsii_.MemberProperty{JsiiProperty: "templateOptions", GoGetter: "TemplateOptions"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYamlString", GoMethod: "ToYamlString"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_StackSetStack{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkStack)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-stacksets.StackSetStackProps",
		reflect.TypeOf((*StackSetStackProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-stacksets.StackSetStackSynthesizer",
		reflect.TypeOf((*StackSetStackSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBootstrapVersionRule", GoMethod: "AddBootstrapVersionRule"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberProperty{JsiiProperty: "assetBucketPrefix", GoGetter: "AssetBucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "assetBuckets", GoGetter: "AssetBuckets"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapQualifier", GoGetter: "BootstrapQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "boundStack", GoGetter: "BoundStack"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromDockerImageAsset", GoMethod: "CloudFormationLocationFromDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromFileAsset", GoMethod: "CloudFormationLocationFromFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "emitArtifact", GoMethod: "EmitArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "emitStackArtifact", GoMethod: "EmitStackArtifact"},
			_jsii_.MemberProperty{JsiiProperty: "lookupRole", GoGetter: "LookupRole"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeStackTemplate", GoMethod: "SynthesizeStackTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeTemplate", GoMethod: "SynthesizeTemplate"},
		},
		func() interface{} {
			j := jsiiProxy_StackSetStackSynthesizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkStackSynthesizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-stacksets.StackSetTarget",
		reflect.TypeOf((*StackSetTarget)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StackSetTarget{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-stacksets.StackSetTemplate",
		reflect.TypeOf((*StackSetTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "templateUrl", GoGetter: "TemplateUrl"},
		},
		func() interface{} {
			return &jsiiProxy_StackSetTemplate{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-stacksets.TargetOptions",
		reflect.TypeOf((*TargetOptions)(nil)).Elem(),
	)
}
