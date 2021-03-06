/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IdmlSettingsDto struct {
	// Default: false
	ExtractNotes bool `json:"extractNotes,omitempty"`
	// Default: true
	SimplifyCodes bool `json:"simplifyCodes,omitempty"`
	// Default: true
	ExtractMasterSpreads bool `json:"extractMasterSpreads,omitempty"`
	// Default: true
	ExtractLockedLayers bool `json:"extractLockedLayers,omitempty"`
	// Default: false
	ExtractInvisibleLayers bool `json:"extractInvisibleLayers,omitempty"`
	// Default: false
	ExtractHiddenConditionalText bool `json:"extractHiddenConditionalText,omitempty"`
	// Default: false
	ExtractHyperlinks bool `json:"extractHyperlinks,omitempty"`
	// Default: false
	KeepKerning bool `json:"keepKerning,omitempty"`
	// Default: false
	KeepTracking bool   `json:"keepTracking,omitempty"`
	TargetFont   string `json:"targetFont,omitempty"`
	// Default: true
	ReplaceFont bool `json:"replaceFont,omitempty"`
	// Default: false
	RemoveXmlElements bool   `json:"removeXmlElements,omitempty"`
	TagRegexp         string `json:"tagRegexp,omitempty"`
	// Default: true
	ExtractVariables bool `json:"extractVariables,omitempty"`
}
