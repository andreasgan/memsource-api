/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PageDtoBusinessUnitDto struct {
	TotalElements    int32             `json:"totalElements,omitempty"`
	TotalPages       int32             `json:"totalPages,omitempty"`
	PageSize         int32             `json:"pageSize,omitempty"`
	PageNumber       int32             `json:"pageNumber,omitempty"`
	NumberOfElements int32             `json:"numberOfElements,omitempty"`
	Content          []BusinessUnitDto `json:"content,omitempty"`
}
