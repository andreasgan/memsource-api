/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UserCreateDtoV2 struct {
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Timezone  string `json:"timezone"`
	Note      string `json:"note,omitempty"`
	// In previous version as terminologist. Default: false
	MayEditApprovedTerms bool `json:"mayEditApprovedTerms,omitempty"`
	// Default: false
	MayRejectJobs bool `json:"mayRejectJobs,omitempty"`
	// Applies only to Linguist or Guest. Default: true
	EditorMachineTranslateEnabled bool `json:"editorMachineTranslateEnabled,omitempty"`
	// Default: true
	ReceiveNewsletter bool `json:"receiveNewsletter,omitempty"`
	// Default: false
	MayEditTranslationMemory bool          `json:"mayEditTranslationMemory,omitempty"`
	SourceLangs              []string      `json:"sourceLangs,omitempty"`
	TargetLangs              []string      `json:"targetLangs,omitempty"`
	WorkflowSteps            []IdReference `json:"workflowSteps,omitempty"`
	Clients                  []IdReference `json:"clients,omitempty"`
	Domains                  []IdReference `json:"domains,omitempty"`
	SubDomains               []IdReference `json:"subDomains,omitempty"`
	ProjectBusinessUnits     []IdReference `json:"projectBusinessUnits,omitempty"`
}
