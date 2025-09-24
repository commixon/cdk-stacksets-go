package cdkstacksets


// The type of concurrency to use when deploying the StackSet to regions.
// Experimental.
type RegionConcurrencyType string

const (
	// Deploy the StackSet to regions sequentially in the order specified in {@link StackSetProps.operationPreferences.regionOrder }.
	//
	// This is the default behavior.
	// Experimental.
	RegionConcurrencyType_SEQUENTIAL RegionConcurrencyType = "SEQUENTIAL"
	// Deploy the StackSet to all regions in parallel.
	// Experimental.
	RegionConcurrencyType_PARALLEL RegionConcurrencyType = "PARALLEL"
)

