package cdkstacksets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/commixon/cdk-stacksets-go/jsii"
)

// Which organizational units and/or accounts the stack set should be deployed to.
//
// `fromAccounts` can be used to deploy the stack set to specific AWS accounts
//
// `fromOrganizationalUnits` can be used to deploy the stack set to specific organizational units
// and optionally include additional accounts from other OUs, or exclude accounts from the specified
// OUs.
//
// Example:
//   // deploy to specific accounts
//   cdkstacksets.StackSetTarget_FromAccounts(&AccountsTargetOptions{
//   	Accounts: []*string{
//   		jsii.String("11111111111"),
//   		jsii.String("22222222222"),
//   	},
//   	Regions: []*string{
//   		jsii.String("us-east-1"),
//   		jsii.String("us-east-2"),
//   	},
//   })
//
//   // deploy to OUs and 1 additional account
//   cdkstacksets.StackSetTarget_FromOrganizationalUnits(&OrganizationsTargetOptions{
//   	Regions: []*string{
//   		jsii.String("us-east-1"),
//   		jsii.String("us-east-2"),
//   	},
//   	OrganizationalUnits: []*string{
//   		jsii.String("ou-1111111"),
//   		jsii.String("ou-2222222"),
//   	},
//   	AdditionalAccounts: []*string{
//   		jsii.String("33333333333"),
//   	},
//   })
//
//   // deploy to OUs but exclude 1 account
//   cdkstacksets.StackSetTarget_FromOrganizationalUnits(&OrganizationsTargetOptions{
//   	Regions: []*string{
//   		jsii.String("us-east-1"),
//   		jsii.String("us-east-2"),
//   	},
//   	OrganizationalUnits: []*string{
//   		jsii.String("ou-1111111"),
//   		jsii.String("ou-2222222"),
//   	},
//   	ExcludeAccounts: []*string{
//   		jsii.String("11111111111"),
//   	},
//   })
//
// Experimental.
type StackSetTarget interface {
}

// The jsii proxy struct for StackSetTarget
type jsiiProxy_StackSetTarget struct {
	_ byte // padding
}

// Experimental.
func NewStackSetTarget_Override(s StackSetTarget) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-stacksets.StackSetTarget",
		nil, // no parameters
		s,
	)
}

// Deploy the StackSet to a list of accounts.
//
// Example:
//   cdkstacksets.StackSetTarget_FromAccounts(&AccountsTargetOptions{
//   	Accounts: []*string{
//   		jsii.String("11111111111"),
//   		jsii.String("22222222222"),
//   	},
//   	Regions: []*string{
//   		jsii.String("us-east-1"),
//   		jsii.String("us-east-2"),
//   	},
//   })
//
// Experimental.
func StackSetTarget_FromAccounts(options *AccountsTargetOptions) StackSetTarget {
	_init_.Initialize()

	if err := validateStackSetTarget_FromAccountsParameters(options); err != nil {
		panic(err)
	}
	var returns StackSetTarget

	_jsii_.StaticInvoke(
		"cdk-stacksets.StackSetTarget",
		"fromAccounts",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Deploy the StackSet to a list of AWS Organizations organizational units.
//
// You can optionally include/exclude individual AWS accounts.
//
// Example:
//   cdkstacksets.StackSetTarget_FromOrganizationalUnits(&OrganizationsTargetOptions{
//   	Regions: []*string{
//   		jsii.String("us-east-1"),
//   		jsii.String("us-east-2"),
//   	},
//   	OrganizationalUnits: []*string{
//   		jsii.String("ou-1111111"),
//   		jsii.String("ou-2222222"),
//   	},
//   })
//
// Experimental.
func StackSetTarget_FromOrganizationalUnits(options *OrganizationsTargetOptions) StackSetTarget {
	_init_.Initialize()

	if err := validateStackSetTarget_FromOrganizationalUnitsParameters(options); err != nil {
		panic(err)
	}
	var returns StackSetTarget

	_jsii_.StaticInvoke(
		"cdk-stacksets.StackSetTarget",
		"fromOrganizationalUnits",
		[]interface{}{options},
		&returns,
	)

	return returns
}

