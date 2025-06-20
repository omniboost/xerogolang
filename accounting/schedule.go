package accounting

// Schedule is an element on a Repeating Invoice - do not use it separately
type Schedule struct {

	// Integer used with the unit e.g. 1 (every 1 week), 2 (every 2 months)
	Period float64 `json:"Period,omitempty" xml:"Period,omitempty"`

	// One of the following : WEEKLY or MONTHLY
	Unit string `json:"Unit,omitempty" xml:"Unit,omitempty"`

	// Integer used with due date type e.g 20 (of following month), 31 (of current month)
	DueDate float64 `json:"DueDate,omitempty" xml:"DueDate,omitempty"`

	// Date the first invoice of the current version of the repeating schedule was generated (changes when repeating invoice is edited)
	StartDate string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`

	// The calendar date of the next invoice in the schedule to be generated
	NextScheduledDate string `json:"NextScheduledDate,omitempty" xml:"NextScheduledDate,omitempty"`

	// Invoice end date – only returned if the template has an end date set
	EndDate string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}
