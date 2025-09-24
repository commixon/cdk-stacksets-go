package cdkstacksets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/commixon/cdk-stacksets-go/internal"
)

// Represents a CloudFormation StackSet.
// Experimental.
type IStackSet interface {
	awscdk.IResource
	// Only available on self managed stacksets.
	//
	// The admin role that CloudFormation will use to perform stackset operations.
	// This role should only have permissions to be assumed by the CloudFormation service
	// and to assume the execution role in each individual account.
	//
	// When you create the execution role it must have an assume role policy statement which
	// allows `sts:AssumeRole` from this admin role.
	//
	// To grant specific users/groups access to use this role to deploy stacksets they must have
	// a policy that allows `iam:GetRole` & `iam:PassRole` on this role resource.
	// Experimental.
	Role() awsiam.IRole
}

// The jsii proxy for IStackSet
type jsiiProxy_IStackSet struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IStackSet) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

