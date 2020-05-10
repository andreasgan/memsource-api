# PreTranslateSettingsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TranslationMemory** | **bool** | Pre-translate from translation memory. Default: false | [optional] [default to null]
**TranslationMemoryThreshold** | **float64** | Pre-translation threshold percent | [optional] [default to null]
**AutoPropagateRepetitions** | **bool** | Propagate repetitions. Default: false | [optional] [default to null]
**MachineTranslation** | **bool** | Pre-translate from machine translation. Default: false | [optional] [default to null]
**NonTranslatables** | **bool** | Pre-translate non-translatables. Default: false | [optional] [default to null]
**RepetitionsAsConfirmed** | **bool** | Set segment status to confirmed for: Repetitions. Default: false | [optional] [default to null]
**Matches100AsTranslated** | **bool** | Set segment status to confirmed for: 100% translation memory matches. Default: false | [optional] [default to null]
**Matches101AsTranslate** | **bool** | Set segment status to confirmed for: 101% translation memory matches. Default: false | [optional] [default to null]
**NonTranslatablesAsTranslated** | **bool** | Set segment status to confirmed for: 100% non-translatable matches. Default: false | [optional] [default to null]
**PreTranslateOnJobCreation** | **bool** | Pre-translate &amp; set job to completed: Pre-translate on job creation. Default: false | [optional] [default to null]
**SetJobStatusCompleted** | **bool** | Pre-translate &amp; set job to completed: Set job to completed once pre-translated. Default: false | [optional] [default to null]
**SetJobStatusCompletedWhenConfirmed** | **bool** | Pre-translate &amp; set job to completed when all segments confirmed: Set job to completed once pre-translated and all segments are confirmed. Default: false | [optional] [default to null]
**SetProjectStatusCompleted** | **bool** | Pre-translate &amp; set job to completed: Set project to completed once all jobs pre-translated.         Default: false | [optional] [default to null]
**LockNonTranslatables** | **bool** | Lock 100% non-translatable matches. Default: false | [optional] [default to null]
**Lock100** | **bool** | Lock 100% translation memory matches. Default: false | [optional] [default to null]
**Lock101** | **bool** | Lock 101% translation memory matches. Default: false | [optional] [default to null]
**NonTranslatablesInEditors** | **bool** | Non translatables enabled in Editors. Default: false | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


