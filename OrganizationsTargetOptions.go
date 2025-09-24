package cdkstacksets


// Options for deploying a StackSet to a set of Organizational Units (OUs).
// Experimental.
type OrganizationsTargetOptions struct {
	// A list of regions the Stack should be deployed to.
	//
	// If {@link StackSetProps.operationPreferences.regionOrder } is specified
	// then the StackSet will be deployed sequentially otherwise it will be
	// deployed to all regions in parallel.
	// Experimental.
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// Parameter overrides that should be applied to only this target.
	// Default: - use parameter overrides specified in {@link StackSetProps.parameterOverrides }
	//
	// Experimental.
	ParameterOverrides *map[string]*string `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// A list of organizational unit ids to deploy to.
	//
	// The StackSet will
	// deploy the provided Stack template to all accounts in the OU.
	// This can be further filtered by specifying either `additionalAccounts`
	// or `excludeAccounts`.
	//
	// If the `deploymentType` is specified with `autoDeployEnabled` then
	// the StackSet will automatically deploy the Stack to new accounts as they
	// are added to the specified `organizationalUnits`.
	// Experimental.
	OrganizationalUnits *[]*string `field:"required" json:"organizationalUnits" yaml:"organizationalUnits"`
	// A list of additional AWS accounts to deploy the StackSet to.
	//
	// This can be
	// used to deploy the StackSet to additional AWS accounts that exist in a
	// different OU than what has been provided in `organizationalUnits`.
	// Default: - Stacks will only be deployed to accounts that exist in the
	// specified organizationalUnits.
	//
	// Experimental.
	AdditionalAccounts *[]*string `field:"optional" json:"additionalAccounts" yaml:"additionalAccounts"`
	// A list of AWS accounts to exclude from deploying the StackSet to.
	//
	// This can
	// be useful if there are accounts that exist in an OU that is provided in
	// `organizationalUnits`, but you do not want the StackSet to be deployed.
	// Default: - Stacks will be deployed to all accounts that exist in the OUs
	// specified in the organizationUnits property.
	//
	// Experimental.
	ExcludeAccounts *[]*string `field:"optional" json:"excludeAccounts" yaml:"excludeAccounts"`
}

