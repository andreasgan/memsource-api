/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SetProjectTransMemoriesDto struct {
	ReadTransMemories []IdReference `json:"readTransMemories,omitempty"`
	// Write translation memory must be included in the read translation memories, too; max 2 write TMs allowed
	WriteTransMemories []IdReference `json:"writeTransMemories,omitempty"`
	// A list of penalties for each read translation memory
	Penalties []float64 `json:"penalties,omitempty"`
	// Set translation memories only for the specific project target language
	TargetLang string `json:"targetLang,omitempty"`
	// set translation memories only for the specific workflow step
	WorkflowStep *IdReference `json:"workflowStep,omitempty"`
}
