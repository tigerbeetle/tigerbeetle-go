// Code generated by "stringer -type=CreateAccountResult -trimprefix=Account"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AccountLinkedEventFailed-1]
	_ = x[AccountReservedFlag-2]
	_ = x[AccountReservedField-3]
	_ = x[AccountIdMustNotBeZero-4]
	_ = x[AccountIdMustNotBeMaxInt-5]
	_ = x[AccountLedgerMustNotBeZero-6]
	_ = x[AccountCodeMustNotBeZero-7]
	_ = x[AccountMutuallyExclusiveFlags-8]
	_ = x[AccountOverflowsDebits-9]
	_ = x[AccountOverflowsCredits-10]
	_ = x[AccountExceedsCredits-11]
	_ = x[AccountExceedsDebits-12]
	_ = x[AccountExistsWithDifferentFlags-13]
	_ = x[AccountExistsWithDifferentUserData-14]
	_ = x[AccountExistsWithDifferentLedger-15]
	_ = x[AccountExistsWithDifferentCode-16]
	_ = x[AccountExistsWithDifferentDebitsPending-17]
	_ = x[AccountExistsWithDifferentDebitsPosted-18]
	_ = x[AccountExistsWithDifferentCreditsPending-19]
	_ = x[AccountExistsWithDifferentCreditsPosted-20]
	_ = x[AccountExists-21]
}

const _CreateAccountResult_name = "LinkedEventFailedReservedFlagReservedFieldIdMustNotBeZeroIdMustNotBeMaxIntLedgerMustNotBeZeroCodeMustNotBeZeroMutuallyExclusiveFlagsOverflowsDebitsOverflowsCreditsExceedsCreditsExceedsDebitsExistsWithDifferentFlagsExistsWithDifferentUserDataExistsWithDifferentLedgerExistsWithDifferentCodeExistsWithDifferentDebitsPendingExistsWithDifferentDebitsPostedExistsWithDifferentCreditsPendingExistsWithDifferentCreditsPostedExists"

var _CreateAccountResult_index = [...]uint16{0, 17, 29, 42, 57, 74, 93, 110, 132, 147, 163, 177, 190, 214, 241, 266, 289, 321, 352, 385, 417, 423}

func (i CreateAccountResult) String() string {
	i -= 1
	if i >= CreateAccountResult(len(_CreateAccountResult_index)-1) {
		return "CreateAccountResult(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CreateAccountResult_name[_CreateAccountResult_index[i]:_CreateAccountResult_index[i+1]]
}