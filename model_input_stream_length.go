/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type InputStreamLength struct {
	Stream            *InputStream `json:"stream,omitempty"`
	Length            int64        `json:"length,omitempty"`
	Name              string       `json:"name,omitempty"`
	CharacterEncoding string       `json:"characterEncoding,omitempty"`
	Extension         string       `json:"extension,omitempty"`
}
