package cdkstacksets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// StackSet stack props.
// Experimental.
type StackSetStackProps struct {
	// The common prefix for the asset bucket names used by this StackSetStack.
	//
	// Required if `assetBuckets` is provided.
	// Default: - No Buckets provided and Assets will not be supported.
	//
	// Experimental.
	AssetBucketPrefix *string `field:"optional" json:"assetBucketPrefix" yaml:"assetBucketPrefix"`
	// An array of Buckets can be passed to store assets, enabling StackSetStack Asset support.
	//
	// One Bucket is required per target region. The name must be `${assetBucketPrefix}-<region>`, where
	// `<region>` is the region targeted by the StackSet.
	// Default: - No Buckets provided and Assets will not be supported.
	//
	// Experimental.
	AssetBuckets *[]awss3.IBucket `field:"optional" json:"assetBuckets" yaml:"assetBuckets"`
}

