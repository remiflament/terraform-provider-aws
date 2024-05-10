// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waf

// Exports for use in tests only.
var (
	ResourceByteMatchSet  = resourceByteMatchSet
	ResourceGeoMatchSet   = resourceGeoMatchSet
	ResourceIPSet         = resourceIPSet
	ResourceRateBasedRule = resourceRateBasedRule
	ResourceRegexMatchSet = resourceRegexMatchSet
	ResourceWebACL        = resourceWebACL

	DiffIPSetDescriptors  = diffIPSetDescriptors
	FindByteMatchSetByID  = findByteMatchSetByID
	FindGeoMatchSetByID   = findGeoMatchSetByID
	FindIPSetByID         = findIPSetByID
	FindRateBasedRuleByID = findRateBasedRuleByID
	FindRegexMatchSetByID = findRegexMatchSetByID
	FindWebACLByID        = findWebACLByID
	FlattenFieldToMatch   = flattenFieldToMatch
)
