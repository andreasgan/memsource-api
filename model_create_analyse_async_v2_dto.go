/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CreateAnalyseAsyncV2Dto struct {
	Jobs []UidReference `json:"jobs"`
	// default: PreAnalyse
	Type_ string `json:"type,omitempty"`
	// Default: true
	IncludeFuzzyRepetitions bool `json:"includeFuzzyRepetitions,omitempty"`
	// Default: true
	IncludeConfirmedSegments bool `json:"includeConfirmedSegments,omitempty"`
	// Default: true
	IncludeNumbers bool `json:"includeNumbers,omitempty"`
	// Default: true
	IncludeLockedSegments bool `json:"includeLockedSegments,omitempty"`
	// Default: true
	CountSourceUnits bool `json:"countSourceUnits,omitempty"`
	// Default: true
	IncludeTransMemory bool `json:"includeTransMemory,omitempty"`
	// Default: false. Works only for type=PreAnalyse.
	IncludeNonTranslatables bool `json:"includeNonTranslatables,omitempty"`
	// Default: false. Works only for type=PreAnalyse.
	IncludeMachineTranslationMatches bool `json:"includeMachineTranslationMatches,omitempty"`
	// Default: false. Works only for type=PostAnalyse.
	TransMemoryPostEditing bool `json:"transMemoryPostEditing,omitempty"`
	// Default: false. Works only for type=PostAnalyse.
	NonTranslatablePostEditing bool `json:"nonTranslatablePostEditing,omitempty"`
	// Default: false. Works only for type=PostAnalyse.
	MachineTranslatePostEditing bool         `json:"machineTranslatePostEditing,omitempty"`
	Name                        string       `json:"name,omitempty"`
	NetRateScheme               *IdReference `json:"netRateScheme,omitempty"`
	// Required for type=Compare
	CompareWorkflowLevel int32 `json:"compareWorkflowLevel,omitempty"`
	// Default: false. Use default project settings. Will be overwritten with setting sent         in the API call.
	UseProjectAnalysisSettings bool               `json:"useProjectAnalysisSettings,omitempty"`
	CallbackUrl                string             `json:"callbackUrl,omitempty"`
	Provider                   *ProviderReference `json:"provider,omitempty"`
}
