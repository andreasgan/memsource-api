/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SetProjectTransMemoryDto struct {
	TransMemory *IdReference `json:"transMemory"`
	// Default: false
	ReadMode bool `json:"readMode,omitempty"`
	// Can be set only for Translation Memory with read == true.<br/>         Max 2 write TMs allowed per project.<br/>         Default: false
	WriteMode bool    `json:"writeMode,omitempty"`
	Penalty   float64 `json:"penalty,omitempty"`
	// Can be set only for penalty == 1<br/>Default: false
	ApplyPenaltyTo101Only bool `json:"applyPenaltyTo101Only,omitempty"`
}
