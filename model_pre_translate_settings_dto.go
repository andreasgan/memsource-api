/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PreTranslateSettingsDto struct {
	// Pre-translate from translation memory. Default: false
	TranslationMemory bool `json:"translationMemory,omitempty"`
	// Pre-translation threshold percent
	TranslationMemoryThreshold float64 `json:"translationMemoryThreshold,omitempty"`
	// Propagate repetitions. Default: false
	AutoPropagateRepetitions bool `json:"autoPropagateRepetitions,omitempty"`
	// Pre-translate from machine translation. Default: false
	MachineTranslation bool `json:"machineTranslation,omitempty"`
	// Pre-translate non-translatables. Default: false
	NonTranslatables bool `json:"nonTranslatables,omitempty"`
	// Set segment status to confirmed for: Repetitions. Default: false
	RepetitionsAsConfirmed bool `json:"repetitionsAsConfirmed,omitempty"`
	// Set segment status to confirmed for: 100% translation memory matches. Default: false
	Matches100AsTranslated bool `json:"matches100AsTranslated,omitempty"`
	// Set segment status to confirmed for: 101% translation memory matches. Default: false
	Matches101AsTranslate bool `json:"matches101AsTranslate,omitempty"`
	// Set segment status to confirmed for: 100% non-translatable matches. Default: false
	NonTranslatablesAsTranslated bool `json:"nonTranslatablesAsTranslated,omitempty"`
	// Pre-translate & set job to completed: Pre-translate on job creation. Default: false
	PreTranslateOnJobCreation bool `json:"preTranslateOnJobCreation,omitempty"`
	// Pre-translate & set job to completed: Set job to completed once pre-translated. Default: false
	SetJobStatusCompleted bool `json:"setJobStatusCompleted,omitempty"`
	// Pre-translate & set job to completed when all segments confirmed: Set job to completed once pre-translated and all segments are confirmed. Default: false
	SetJobStatusCompletedWhenConfirmed bool `json:"setJobStatusCompletedWhenConfirmed,omitempty"`
	// Pre-translate & set job to completed: Set project to completed once all jobs pre-translated.         Default: false
	SetProjectStatusCompleted bool `json:"setProjectStatusCompleted,omitempty"`
	// Lock 100% non-translatable matches. Default: false
	LockNonTranslatables bool `json:"lockNonTranslatables,omitempty"`
	// Lock 100% translation memory matches. Default: false
	Lock100 bool `json:"lock100,omitempty"`
	// Lock 101% translation memory matches. Default: false
	Lock101 bool `json:"lock101,omitempty"`
	// Non translatables enabled in Editors. Default: false
	NonTranslatablesInEditors bool `json:"nonTranslatablesInEditors,omitempty"`
}
