package cdkstacksets


// Common options for deploying a StackSet to a target.
// Experimental.
type TargetOptions struct {
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
}

