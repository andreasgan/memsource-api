/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HtmlSettingsDto struct {
	// Default: true
	BreakTagCreatesSegment bool `json:"breakTagCreatesSegment,omitempty"`
	// Default: true
	UnknownTagCreatesTag bool `json:"unknownTagCreatesTag,omitempty"`
	// Default: false
	PreserveWhitespace bool `json:"preserveWhitespace,omitempty"`
	// Default: true
	ImportComments bool `json:"importComments,omitempty"`
	// Example: \"script,blockquote\"
	ExcludeElements        string `json:"excludeElements,omitempty"`
	TagRegexp              string `json:"tagRegexp,omitempty"`
	CharEntitiesToTags     string `json:"charEntitiesToTags,omitempty"`
	TranslateMetaTagRegexp string `json:"translateMetaTagRegexp,omitempty"`
}
