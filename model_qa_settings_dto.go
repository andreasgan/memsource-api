/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type QaSettingsDto struct {
	EmptyTranslation               bool                    `json:"emptyTranslation,omitempty"`
	InconsistentTranslation        bool                    `json:"inconsistentTranslation,omitempty"`
	JoinTags                       bool                    `json:"joinTags,omitempty"`
	MissingNumbers                 bool                    `json:"missingNumbers,omitempty"`
	SegmentNotConfirmed            bool                    `json:"segmentNotConfirmed,omitempty"`
	Terminology                    bool                    `json:"terminology,omitempty"`
	MultipleSpaces                 bool                    `json:"multipleSpaces,omitempty"`
	TrailingSpace                  bool                    `json:"trailingSpace,omitempty"`
	TrailingPunctuation            bool                    `json:"trailingPunctuation,omitempty"`
	TargetLength                   *TargetLengthDto        `json:"targetLength,omitempty"`
	Formatting                     bool                    `json:"formatting,omitempty"`
	UnresolvedComment              bool                    `json:"unresolvedComment,omitempty"`
	EmptyPairTags                  bool                    `json:"emptyPairTags,omitempty"`
	StrictJobStatus                bool                    `json:"strictJobStatus,omitempty"`
	ForbiddenStrings               *ForbiddenStringsDto    `json:"forbiddenStrings,omitempty"`
	ExcludeLockedSegments          bool                    `json:"excludeLockedSegments,omitempty"`
	IgnoreNotApprovedTerms         bool                    `json:"ignoreNotApprovedTerms,omitempty"`
	SpellCheck                     bool                    `json:"spellCheck,omitempty"`
	RepeatedWords                  bool                    `json:"repeatedWords,omitempty"`
	InconsistentTagContent         bool                    `json:"inconsistentTagContent,omitempty"`
	EmptyTagContent                bool                    `json:"emptyTagContent,omitempty"`
	XliffTags                      bool                    `json:"xliffTags,omitempty"`
	NestedTags                     bool                    `json:"nestedTags,omitempty"`
	ForbiddenTerms                 bool                    `json:"forbiddenTerms,omitempty"`
	TargetLengthPercent            *TargetLengthPercentDto `json:"targetLengthPercent,omitempty"`
	TargetLengthPerSegment         bool                    `json:"targetLengthPerSegment,omitempty"`
	NewerAtPrecedingWorkflowStep   bool                    `json:"newerAtPrecedingWorkflowStep,omitempty"`
	LeadingAndTrailingSpaces       bool                    `json:"leadingAndTrailingSpaces,omitempty"`
	IgnoreInAllWorkflowSteps       bool                    `json:"ignoreInAllWorkflowSteps,omitempty"`
	UnmodifiedFuzzyTranslation     bool                    `json:"unmodifiedFuzzyTranslation,omitempty"`
	UnmodifiedFuzzyTranslationTM   bool                    `json:"unmodifiedFuzzyTranslationTM,omitempty"`
	UnmodifiedFuzzyTranslationMTNT bool                    `json:"unmodifiedFuzzyTranslationMTNT,omitempty"`
	Regexp                         *RegexpCheckDto         `json:"regexp,omitempty"`
	ExtraNumbers                   bool                    `json:"extraNumbers,omitempty"`
	TargetSourceIdentical          bool                    `json:"targetSourceIdentical,omitempty"`
	Moravia                        *MoraviaQaDto           `json:"moravia,omitempty"`
}
