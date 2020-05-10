/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BackgroundTasksTbDto struct {
	Status            string            `json:"status,omitempty"`
	FinishedDataText  string            `json:"finishedDataText,omitempty"`
	AsyncRequest      *AsyncRequestDto  `json:"asyncRequest,omitempty"`
	LastTaskString    string            `json:"lastTaskString,omitempty"`
	Metadata          *MetadataResponse `json:"metadata,omitempty"`
	LastTaskOk        string            `json:"lastTaskOk,omitempty"`
	LastTaskError     string            `json:"lastTaskError,omitempty"`
	LastTaskErrorHtml string            `json:"lastTaskErrorHtml,omitempty"`
}