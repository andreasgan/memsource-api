/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MultilingualXlsSettingsDto struct {
	SourceColumn string `json:"sourceColumn,omitempty"`
	// Format: \"language\":\"column\"; example: {\"en\": \"A\", \"sk\": \"B\"}
	TargetColumns     map[string]string `json:"targetColumns,omitempty"`
	ContextNoteColumn string            `json:"contextNoteColumn,omitempty"`
	ContextKeyColumn  string            `json:"contextKeyColumn,omitempty"`
	TagRegexp         string            `json:"tagRegexp,omitempty"`
	// Default: false
	HtmlSubFilter bool `json:"htmlSubFilter,omitempty"`
	// Default: true
	Segmentation bool   `json:"segmentation,omitempty"`
	ImportRows   string `json:"importRows,omitempty"`
	MaxLenColumn string `json:"maxLenColumn,omitempty"`
}
