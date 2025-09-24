package cdkstacksets


// Options for StackSets that are managed by AWS Organizations.
// Experimental.
type ServiceManagedOptions struct {
	// Whether or not the StackSet should automatically create/remove the Stack from AWS accounts that are added/removed from an organizational unit.
	//
	// This has no effect if {@link StackSetTarget.fromAccounts} is used
	// Default: true.
	//
	// Experimental.
	AutoDeployEnabled *bool `field:"optional" json:"autoDeployEnabled" yaml:"autoDeployEnabled"`
	// Whether stacks should be removed from AWS accounts that are removed from an organizational unit.
	//
	// By default the stack will be retained (not deleted)
	//
	// This has no effect if {@link StackSetTarget.fromAccounts} is used
	// Default: true.
	//
	// Experimental.
	AutoDeployRetainStacks *bool `field:"optional" json:"autoDeployRetainStacks" yaml:"autoDeployRetainStacks"`
	// Whether or not the account this StackSet is deployed from is the delegated admin account.
	//
	// Set this to `false` if you are using the AWS Organizations management account instead.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html
	//
	// Default: true.
	//
	// Experimental.
	DelegatedAdmin *bool `field:"optional" json:"delegatedAdmin" yaml:"delegatedAdmin"`
}

