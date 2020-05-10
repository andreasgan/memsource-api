/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SearchTmRequestDto struct {
	Segment         string  `json:"segment"`
	WorkflowLevel   int32   `json:"workflowLevel,omitempty"`
	ScoreThreshold  float64 `json:"scoreThreshold,omitempty"`
	PreviousSegment string  `json:"previousSegment,omitempty"`
	NextSegment     string  `json:"nextSegment,omitempty"`
	ContextKey      string  `json:"contextKey,omitempty"`
	// Default: 5
	MaxSegments int32 `json:"maxSegments,omitempty"`
	// Default: 5
	MaxSubSegments int32            `json:"maxSubSegments,omitempty"`
	TagMetadata    []TagMetadataDto `json:"tagMetadata,omitempty"`
	TargetLangs    []string         `json:"targetLangs"`
}
