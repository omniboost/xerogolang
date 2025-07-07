package accounting

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/markbates/goth"
	"github.com/omniboost/xerogolang"
	"github.com/omniboost/xerogolang/helpers"
)

// Report is an organised set of financial information
type Report struct {
	//The ID of the report
	ReportID string `json:"ReportID,omitempty" xml:"ReportID,omitempty"`
	//The Name of the report
	ReportName string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	//The type of report
	ReportType string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	//A collection of titles for the report
	ReportTitles *[]string `json:"ReportTitles,omitempty" xml:"ReportTitles>ReportTitle,omitempty"`
	//The date of the report
	ReportDate string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// Last modified date UTC format
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`
	//Attributes of the report
	Attributes *[]ReportAttribute `json:"Attributes,omitempty" xml:"Attributes>Attribute,omitempty"`
	//Rows on the report that may contain cells, Attributes, or other rows
	Rows *[]Row `json:"Rows,omitempty" xml:"Rows>Row,omitempty"`
}

// Reports is a collection of reports
type Reports struct {
	Reports []Report `json:"Reports" xml:"Report"`
}

// The Xero API returns Dates based on the .Net JSON date format available at the time of development
// We need to convert these to a more usable format - RFC3339 for consistency with what the API expects to recieve
func (r *Reports) convertDates() error {
	var err error
	for n := len(r.Reports) - 1; n >= 0; n-- {
		if r.Reports[n].UpdatedDateUTC != "" {
			r.Reports[n].UpdatedDateUTC, err = helpers.DotNetJSONTimeToRFC3339(r.Reports[n].UpdatedDateUTC, true)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalReport(reportResponseBytes []byte) (*Reports, error) {
	var reportResponse *Reports
	err := json.Unmarshal(reportResponseBytes, &reportResponse)
	if err != nil {
		return nil, err
	}

	err = reportResponse.convertDates()
	if err != nil {
		return nil, err
	}

	return reportResponse, err
}

// Run1099 will run the 1099 Report and marshal the results to a Report Struct
// This Report will only work for US based Organisations
func Run1099(ctx context.Context, provider xerogolang.IProvider, session goth.Session, reportYear int) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	querystringParameters := map[string]string{
		"reportYear": strconv.Itoa(reportYear),
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/TenNinetyNine", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunAgedPayablesByContact will run the Aged Payables By Contact Report and marshal the results to a Report Struct
// Date, FromDate and ToDate can be added as optional paramters as a map
func RunAgedPayablesByContact(ctx context.Context, provider xerogolang.IProvider, session goth.Session, contactID string, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if querystringParameters != nil {
		querystringParameters["ContactID"] = contactID
	} else {
		querystringParameters = map[string]string{
			"ContactID": contactID,
		}
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/AgedPayablesByContact", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunAgedReceivablesByContact will run the Aged Receivables By Contact Report and marshal the results to a Report Struct
// Date, FromDate and ToDate can be added as optional paramters as a map
func RunAgedReceivablesByContact(ctx context.Context, provider xerogolang.IProvider, session goth.Session, contactID string, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if querystringParameters != nil {
		querystringParameters["ContactID"] = contactID
	} else {
		querystringParameters = map[string]string{
			"ContactID": contactID,
		}
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/AgedReceivablesByContact", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunBalanceSheet will run the Balance Sheet Report and marshal the results to a Report Struct
// date, trackingOptionID1, trackingOptionID2, standardLayout, and paymentsOnly can be added as optional paramters as a map
func RunBalanceSheet(ctx context.Context, provider xerogolang.IProvider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/BalanceSheet", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunBankStatement will run the Bank Statement Report and marshal the results to a Report Struct
// FromDate and ToDate can be added as optional paramters as a map
func RunBankStatement(ctx context.Context, provider xerogolang.IProvider, session goth.Session, bankAccountID string, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	if querystringParameters != nil {
		querystringParameters["bankAccountID"] = bankAccountID
	} else {
		querystringParameters = map[string]string{
			"bankAccountID": bankAccountID,
		}
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/BankStatement", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunBankSummary will run the Bank Summary Report and marshal the results to a Report Struct
// FromDate and ToDate can be added as optional paramters as a map
func RunBankSummary(ctx context.Context, provider xerogolang.IProvider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/BankSummary", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunBASReport will retrieve an individual BAS Report given a reportID and marshal the results to a Report Struct
// Will only work for AU based Organisations
func RunBASReport(ctx context.Context, provider xerogolang.IProvider, session goth.Session, reportID string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/"+reportID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunBASReports will retrieve all BAS Reports and marshal the results to a Report Struct
// Will only work for AU based Organisations
func RunBASReports(ctx context.Context, provider xerogolang.IProvider, session goth.Session) (*Reports, error) {
	return RunBASReport(ctx, provider, session, "")
}

// RunBudgetSummary will run the Budget Summary Report and marshal the results to a Report Struct
func RunBudgetSummary(ctx context.Context, provider xerogolang.IProvider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/BudgetSummary", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunExecutiveSummary will run the Executive Summary Report and marshal the results to a Report Struct
// date can be added as an optional paramter as a map
func RunExecutiveSummary(ctx context.Context, provider xerogolang.IProvider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/ExecutiveSummary", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunGSTReport will retrieve an individual GST Report given a reportID and marshal the results to a Report Struct
// Will only work for NZ based Organisations
func RunGSTReport(ctx context.Context, provider xerogolang.IProvider, session goth.Session, reportID string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/"+reportID, additionalHeaders, nil)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunGSTReports will retrieve all GST Reports and marshal the results to a Report Struct
// Will only work for NZ based Organisations
func RunGSTReports(ctx context.Context, provider xerogolang.IProvider, session goth.Session) (*Reports, error) {
	return RunGSTReport(ctx, provider, session, "")
}

// RunProfitAndLoss will run the Profit And Loss Report and marshal the results to a Report Struct
// date, trackingCategoryID, trackingOptionID, trackingCategoryID2, trackingOptionID2,
// standardLayout, and paymentsOnly can be added as optional paramters as a map
func RunProfitAndLoss(ctx context.Context, provider xerogolang.IProvider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/ProfitAndLoss", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}

// RunTrialBalance will run the TrialBalance Report and marshal the results to a Report Struct
// date and paymentsOnly can be added as optional paramters as a map
func RunTrialBalance(ctx context.Context, provider xerogolang.IProvider, session goth.Session, querystringParameters map[string]string) (*Reports, error) {
	additionalHeaders := map[string]string{
		"Accept": "application/json",
	}

	reportResponseBytes, err := provider.Find(ctx, session, "Reports/TrialBalance", additionalHeaders, querystringParameters)
	if err != nil {
		return nil, err
	}

	return unmarshalReport(reportResponseBytes)
}
