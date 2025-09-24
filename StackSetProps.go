package cdkstacksets


// Experimental.
type StackSetProps struct {
	// Which accounts/OUs and regions to deploy the StackSet to.
	// Experimental.
	Target StackSetTarget `field:"required" json:"target" yaml:"target"`
	// The Stack that will be deployed to the target.
	// Experimental.
	Template StackSetTemplate `field:"required" json:"template" yaml:"template"`
	// Specify a list of capabilities required by your stackset.
	//
	// StackSets that contains certain functionality require an explicit acknowledgement
	// that the stack contains these capabilities.
	//
	// If you deploy a stack that requires certain capabilities and they are
	// not specified, the deployment will fail with a `InsufficientCapabilities` error.
	// Default: - no specific capabilities.
	//
	// Experimental.
	Capabilities *[]Capability `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The type of deployment for this StackSet.
	//
	// The deployment can either be managed by
	// AWS Organizations (i.e. DeploymentType.serviceManaged()) or by the AWS account that
	// the StackSet is deployed from.
	//
	// In order to use DeploymentType.serviceManaged() the account needs to either be the
	// organizations's management account or a delegated administrator account.
	// Default: DeploymentType.selfManaged()
	//
	// Experimental.
	DeploymentType DeploymentType `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// An optional description to add to the StackSet.
	// Default: - none.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If this is `true` then StackSets will perform non-conflicting operations concurrently and queue any conflicting operations.
	//
	// This means that you can submit more than one operation per StackSet and they will be
	// executed concurrently. For example you can submit a single request that updates existing
	// stack instances *and* creates new stack instances. Any conflicting operations will be queued
	// for immediate processing once the conflict is resolved.
	// Default: true.
	//
	// Experimental.
	ManagedExecution *bool `field:"optional" json:"managedExecution" yaml:"managedExecution"`
	// The operation preferences for the StackSet.
	//
	// This allows you to control how the StackSet is deployed
	// across the target accounts and regions.
	// Experimental.
	OperationPreferences *OperationPreferences `field:"optional" json:"operationPreferences" yaml:"operationPreferences"`
	// The input parameters for the stack set template.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The name of the stack set.
	// Default: - CloudFormation generated name.
	//
	// Experimental.
	StackSetName *string `field:"optional" json:"stackSetName" yaml:"stackSetName"`
}

