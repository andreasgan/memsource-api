/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type JobPartExtendedDto struct {
	Uid                   string                        `json:"uid,omitempty"`
	Status                string                        `json:"status,omitempty"`
	Providers             []ProviderReference           `json:"providers,omitempty"`
	SourceLang            string                        `json:"sourceLang,omitempty"`
	TargetLang            string                        `json:"targetLang,omitempty"`
	WorkflowLevel         int32                         `json:"workflowLevel,omitempty"`
	WorkflowStep          *ProjectWorkflowStepReference `json:"workflowStep,omitempty"`
	Filename              string                        `json:"filename,omitempty"`
	DateDue               time.Time                     `json:"dateDue,omitempty"`
	WordsCount            int32                         `json:"wordsCount,omitempty"`
	BeginIndex            int32                         `json:"beginIndex,omitempty"`
	EndIndex              int32                         `json:"endIndex,omitempty"`
	IsParentJobSplit      bool                          `json:"isParentJobSplit,omitempty"`
	DateCreated           time.Time                     `json:"dateCreated,omitempty"`
	Project               *ProjectReference             `json:"project,omitempty"`
	LastWorkflowLevel     int32                         `json:"lastWorkflowLevel,omitempty"`
	WorkUnit              *IdReference                  `json:"workUnit,omitempty"`
	Imported              bool                          `json:"imported,omitempty"`
	Continuous            bool                          `json:"continuous,omitempty"`
	ContinuousJobInfo     *ContinuousJobInfoDto         `json:"continuousJobInfo,omitempty"`
	OriginalFileDirectory string                        `json:"originalFileDirectory,omitempty"`
}
