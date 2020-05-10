/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SearchRequestDto struct {
	Query           string           `json:"query"`
	SourceLang      string           `json:"sourceLang"`
	TargetLangs     []string         `json:"targetLangs,omitempty"`
	PreviousSegment string           `json:"previousSegment,omitempty"`
	NextSegment     string           `json:"nextSegment,omitempty"`
	TagMetadata     []TagMetadataDto `json:"tagMetadata,omitempty"`
	// Remove leading and trailing whitespace from query. Default: true
	TrimQuery bool `json:"trimQuery,omitempty"`
	// Return both wildcard and exact search results. Default: true
	PhraseQuery bool `json:"phraseQuery,omitempty"`
}
