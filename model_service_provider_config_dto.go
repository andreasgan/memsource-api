/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ServiceProviderConfigDto struct {
	AuthenticationSchemes []AuthSchema `json:"authenticationSchemes,omitempty"`
	Schemas               []string     `json:"schemas,omitempty"`
	Patch                 *Supported   `json:"patch,omitempty"`
	Bulk                  *Supported   `json:"bulk,omitempty"`
	Filter                *Supported   `json:"filter,omitempty"`
	ChangePassword        *Supported   `json:"changePassword,omitempty"`
	Sort                  *Supported   `json:"sort,omitempty"`
	Etag                  *Supported   `json:"etag,omitempty"`
	XmlDataFormat         *Supported   `json:"xmlDataFormat,omitempty"`
}
