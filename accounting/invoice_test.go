package accounting

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
)

// Xero validates header totals when they are present, so an invoice built
// from line items only must not send SubTotal/TotalTax/Total as "0".
func TestInvoiceMarshalOmitsZeroTotals(t *testing.T) {
	inv := Invoices{Invoices: []Invoice{{
		Type:   "ACCREC",
		Status: "AUTHORISED",
		LineItems: []LineItem{{
			Description: "Sales",
			LineAmount:  decimal.NewFromFloat(100),
			TaxAmount:   decimal.Zero,
		}},
	}}}

	b, err := json.Marshal(inv)
	if err != nil {
		t.Fatal(err)
	}
	body := string(b)

	for _, field := range []string{"CurrencyRate", "SubTotal", "TotalTax", "Total", "TotalDiscount", "AmountDue", "AmountPaid", "AmountCredited"} {
		if strings.Contains(body, `"`+field+`":`) {
			t.Errorf("expected %s to be omitted, got %s", field, body)
		}
	}

	// A zero line TaxAmount still means "no tax" and has to be sent.
	if !strings.Contains(body, `"TaxAmount":"0"`) {
		t.Errorf("expected line TaxAmount to be sent, got %s", body)
	}
}

func TestInvoiceMarshalKeepsNonZeroTotals(t *testing.T) {
	inv := Invoice{Total: decimal.NewFromFloat(110)}

	b, err := json.Marshal(inv)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(b), `"Total":"110"`) {
		t.Errorf("expected Total to be sent, got %s", b)
	}
}
