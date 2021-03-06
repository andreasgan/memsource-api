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

type JobPartReference struct {
	Uid                           string                 `json:"uid,omitempty"`
	Status                        string                 `json:"status,omitempty"`
	Providers                     []ProviderReference    `json:"providers,omitempty"`
	TargetLang                    string                 `json:"targetLang,omitempty"`
	WorkflowLevel                 int32                  `json:"workflowLevel,omitempty"`
	WorkflowStep                  *WorkflowStepReference `json:"workflowStep,omitempty"`
	Filename                      string                 `json:"filename,omitempty"`
	DateDue                       time.Time              `json:"dateDue,omitempty"`
	DateCreated                   time.Time              `json:"dateCreated,omitempty"`
	Imported                      bool                   `json:"imported,omitempty"`
	JobAssignedEmailTemplate      *ObjectReference       `json:"jobAssignedEmailTemplate,omitempty"`
	NotificationIntervalInMinutes int32                  `json:"notificationIntervalInMinutes,omitempty"`
	Continuous                    bool                   `json:"continuous,omitempty"`
}
