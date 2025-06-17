package accounting

import "github.com/shopspring/decimal"

// JournalLine is a line on a Journal
type JournalLine struct {
	//Xero identifier
	JournalLineID string `json:"JournalLineID" xml:"JournalLineID"`

	//Xero identifier for an account
	AccountID string `json:"AccountID" xml:"AccountID"`

	// See Accounts
	AccountCode string `json:"AccountCode" xml:"AccountCode"`

	// See Accounts
	AccountType string `json:"AccountType" xml:"AccountType"`

	// See Accounts
	AccountName string `json:"AccountName" xml:"AccountName"`

	// The description from the source transaction line item. Only returned if populated.
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// Net amount of journal line. This will be a positive value for a debit and negative for a credit
	NetAmount decimal.Decimal `json:"NetAmount" xml:"NetAmount"`

	// 	Gross amount of journal line (NetAmount + TaxAmount).
	GrossAmount decimal.Decimal `json:"GrossAmount" xml:"GrossAmount"`

	// The calculated tax amount based on the TaxType and LineAmount
	TaxAmount decimal.Decimal `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`

	// Used as an override if the default Tax Code for the selected <AccountCode> is not correct – see TaxTypes.
	TaxType string `json:"TaxType,omitempty" xml:"TaxType,omitempty"`

	//see tax TaxTypes
	TaxName string `json:"TaxName,omitempty" xml:"TaxName,omitempty"`

	// Optional Tracking Category – see Tracking. Any JournalLine can have a maximum of 2 <TrackingCategory> elements.
	TrackingCategories []TrackingCategory `json:"TrackingCategories,omitempty" xml:"TrackingCategories>TrackingCategory,omitempty"`
}
