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

type AsyncRequestV2Dto struct {
	Id            string              `json:"id,omitempty"`
	CreatedBy     *UserReference      `json:"createdBy,omitempty"`
	DateCreated   time.Time           `json:"dateCreated,omitempty"`
	Action        string              `json:"action,omitempty"`
	AsyncResponse *AsyncResponseV2Dto `json:"asyncResponse,omitempty"`
	Parent        *AsyncRequestV2Dto  `json:"parent,omitempty"`
	Project       *ProjectReference   `json:"project,omitempty"`
}
