// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     PluginGoTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package dataquery

// Defines values for PyroscopeQueryType.
const (
	PyroscopeQueryTypeBoth    PyroscopeQueryType = "both"
	PyroscopeQueryTypeMetrics PyroscopeQueryType = "metrics"
	PyroscopeQueryTypeProfile PyroscopeQueryType = "profile"
)

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
type DataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any `json:"datasource,omitempty"`

	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
}

// GrafanaPyroscopeDataQuery defines model for GrafanaPyroscopeDataQuery.
type GrafanaPyroscopeDataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any `json:"datasource,omitempty"`

	// Allows to group the results.
	GroupBy []string `json:"groupBy,omitempty"`

	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`

	// Specifies the query label selectors.
	LabelSelector *string `json:"labelSelector,omitempty"`

	// Sets the maximum number of time series.
	Limit *int64 `json:"limit,omitempty"`

	// Sets the maximum number of nodes in the flamegraph.
	MaxNodes *int64 `json:"maxNodes,omitempty"`

	// Specifies the type of profile to query.
	ProfileTypeId *string `json:"profileTypeId,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId *string `json:"refId,omitempty"`

	// Specifies the query span selectors.
	SpanSelector []string `json:"spanSelector,omitempty"`
}

// PyroscopeQueryType defines model for PyroscopeQueryType.
type PyroscopeQueryType string
