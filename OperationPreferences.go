package cdkstacksets


// CloudFormation operation preferences.
//
// This maps to `aws_cloudformation.CfnStackSet.OperationPreferencesProperty`.
// Experimental.
type OperationPreferences struct {
	// The number of stack instances that can fail before the operation is considered failed.
	// Experimental.
	FailureToleranceCount *float64 `field:"optional" json:"failureToleranceCount" yaml:"failureToleranceCount"`
	// The percentage of stack instances that can fail before the operation is considered failed.
	// Experimental.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// The maximum number of stack instances that can be created or updated concurrently.
	// Experimental.
	MaxConcurrentCount *float64 `field:"optional" json:"maxConcurrentCount" yaml:"maxConcurrentCount"`
	// The maximum percentage of stack instances that can be created or updated concurrently.
	// Experimental.
	MaxConcurrentPercentage *float64 `field:"optional" json:"maxConcurrentPercentage" yaml:"maxConcurrentPercentage"`
	// Whether to deploy multiple regions sequentially or in parallel.
	// Default: RegionConcurrencyType.SEQUENTIAL
	//
	// Experimental.
	RegionConcurrencyType RegionConcurrencyType `field:"optional" json:"regionConcurrencyType" yaml:"regionConcurrencyType"`
	// The order in which to deploy the stack instances to the regions.
	// Experimental.
	RegionOrder *[]*string `field:"optional" json:"regionOrder" yaml:"regionOrder"`
}

