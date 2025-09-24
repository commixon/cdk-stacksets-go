//go:build no_runtime_type_checking

package cdkstacksets

// Building without runtime type checking enabled, so all the below just return nil

func validateStackSetTarget_FromAccountsParameters(options *AccountsTargetOptions) error {
	return nil
}

func validateStackSetTarget_FromOrganizationalUnitsParameters(options *OrganizationsTargetOptions) error {
	return nil
}

