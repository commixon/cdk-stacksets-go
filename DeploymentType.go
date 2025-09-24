package cdkstacksets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/commixon/cdk-stacksets-go/jsii"
)

// Experimental.
type DeploymentType interface {
}

// The jsii proxy struct for DeploymentType
type jsiiProxy_DeploymentType struct {
	_ byte // padding
}

// Experimental.
func NewDeploymentType_Override(d DeploymentType) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-stacksets.DeploymentType",
		nil, // no parameters
		d,
	)
}

// StackSets deployed using the self managed model require you to create the necessary IAM roles in the source and target AWS accounts and to setup the required IAM permissions.
//
// Using this model you can only deploy to AWS accounts that have the necessary IAM roles/permissions
// pre-created.
// Experimental.
func DeploymentType_SelfManaged(options *SelfManagedOptions) DeploymentType {
	_init_.Initialize()

	if err := validateDeploymentType_SelfManagedParameters(options); err != nil {
		panic(err)
	}
	var returns DeploymentType

	_jsii_.StaticInvoke(
		"cdk-stacksets.DeploymentType",
		"selfManaged",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// StackSets deployed using service managed permissions allow you to deploy StackSet instances to accounts within an AWS Organization.
//
// Using this module
// AWS Organizations will handle creating the necessary IAM roles and setting up the
// required permissions.
//
// This model also allows you to enable automated deployments which allows the StackSet
// to be automatically deployed to new accounts that are added to your organization in the future.
//
// This model requires you to be operating in either the AWS Organizations management account
// or the delegated administrator account.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html#stacksets-concepts-stackset-permission-models
//
// Experimental.
func DeploymentType_ServiceManaged(options *ServiceManagedOptions) DeploymentType {
	_init_.Initialize()

	if err := validateDeploymentType_ServiceManagedParameters(options); err != nil {
		panic(err)
	}
	var returns DeploymentType

	_jsii_.StaticInvoke(
		"cdk-stacksets.DeploymentType",
		"serviceManaged",
		[]interface{}{options},
		&returns,
	)

	return returns
}

