/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MachineTranslateSettingsDto struct {
	Id                  string            `json:"id,omitempty"`
	Uid                 string            `json:"uid,omitempty"`
	BaseName            string            `json:"baseName,omitempty"`
	Name                string            `json:"name,omitempty"`
	Type_               string            `json:"type,omitempty"`
	Default_            bool              `json:"default_,omitempty"`
	IncludeTags         bool              `json:"includeTags,omitempty"`
	MtQualityEstimation bool              `json:"mtQualityEstimation,omitempty"`
	Args                map[string]string `json:"args,omitempty"`
}
