/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SearchTmResponseDtoV3 struct {
	SegmentId    string                    `json:"segmentId,omitempty"`
	Source       *SearchTmSegmentDtoV3     `json:"source,omitempty"`
	Translations []SearchTmSegmentDtoV3    `json:"translations,omitempty"`
	TransMemory  *SearchTmTransMemoryDtoV3 `json:"transMemory,omitempty"`
	GrossScore   float64                   `json:"grossScore,omitempty"`
	Score        float64                   `json:"score,omitempty"`
	SubSegment   bool                      `json:"subSegment,omitempty"`
}
