package commConst

type DataSource string

const (
	ZenData DataSource = "zendata"
	CSV     DataSource = "csv"
	Excel   DataSource = "excel"
)

type ExtractorType string

const (
	Value       ExtractorType = "value"
	XPath       ExtractorType = "xpath"
	JSONPath    ExtractorType = "json_path"
	CssSelector ExtractorType = "sss_selector"
)

type ExtractorSource string

const (
	Body   ExtractorSource = "body"
	Header ExtractorSource = "header"
)
