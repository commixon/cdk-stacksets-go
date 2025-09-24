package cdkstacksets


// StackSets that contains certain functionality require an explicit acknowledgement that the stack contains these capabilities.
// Experimental.
type Capability string

const (
	// Required if the stack contains IAM resources with custom names.
	// Experimental.
	Capability_NAMED_IAM Capability = "NAMED_IAM"
	// Required if the stack contains IAM resources.
	//
	// If the IAM resources
	// also have custom names then specify {@link Capability.NAMED_IAM} instead.
	// Experimental.
	Capability_IAM Capability = "IAM"
	// Required if the stack contains macros.
	//
	// Not supported if deploying
	// a service managed stackset.
	// Experimental.
	Capability_AUTO_EXPAND Capability = "AUTO_EXPAND"
)

