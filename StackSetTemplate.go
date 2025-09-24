package cdkstacksets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/commixon/cdk-stacksets-go/jsii"
)

// Represents a StackSet CloudFormation template.
// Experimental.
type StackSetTemplate interface {
	// The S3 URL of the StackSet template.
	// Experimental.
	TemplateUrl() *string
}

// The jsii proxy struct for StackSetTemplate
type jsiiProxy_StackSetTemplate struct {
	_ byte // padding
}

func (j *jsiiProxy_StackSetTemplate) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}


// Experimental.
func NewStackSetTemplate_Override(s StackSetTemplate) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-stacksets.StackSetTemplate",
		nil, // no parameters
		s,
	)
}

// Creates a StackSetTemplate from a StackSetStack.
//
// Returns: StackSetTemplate.
// Experimental.
func StackSetTemplate_FromStackSetStack(stack StackSetStack) StackSetTemplate {
	_init_.Initialize()

	if err := validateStackSetTemplate_FromStackSetStackParameters(stack); err != nil {
		panic(err)
	}
	var returns StackSetTemplate

	_jsii_.StaticInvoke(
		"cdk-stacksets.StackSetTemplate",
		"fromStackSetStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

