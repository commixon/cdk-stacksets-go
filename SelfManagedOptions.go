package cdkstacksets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for StackSets that are not managed by AWS Organizations.
// Experimental.
type SelfManagedOptions struct {
	// The admin role that CloudFormation will use to perform stackset operations.
	//
	// This role should only have permissions to be assumed by the CloudFormation service
	// and to assume the execution role in each individual account.
	//
	// When you create the execution role it must have an assume role policy statement which
	// allows `sts:AssumeRole` from this admin role.
	//
	// To grant specific users/groups access to use this role to deploy stacksets they must have
	// a policy that allows `iam:GetRole` & `iam:PassRole` on this role resource.
	// Default: - a default role will be created.
	//
	// Experimental.
	AdminRole awsiam.IRole `field:"optional" json:"adminRole" yaml:"adminRole"`
	// The name of the stackset execution role that already exists in each target AWS account.
	//
	// This role must be configured with a trust policy that allows `sts:AssumeRole` from the `adminRole`.
	//
	// In addition this role must have the necessary permissions to manage the resources created by the stackset.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html#stacksets-prereqs-accountsetup
	//
	// Default: - AWSCloudFormationStackSetExecutionRole.
	//
	// Experimental.
	ExecutionRoleName *string `field:"optional" json:"executionRoleName" yaml:"executionRoleName"`
}

