/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SearchTmSegmentDtoV2 struct {
	Id              string                  `json:"id,omitempty"`
	Text            string                  `json:"text,omitempty"`
	Lang            string                  `json:"lang,omitempty"`
	Rtl             bool                    `json:"rtl,omitempty"`
	ModifiedAt      int64                   `json:"modifiedAt,omitempty"`
	CreatedAt       int64                   `json:"createdAt,omitempty"`
	ModifiedBy      *UserReference          `json:"modifiedBy,omitempty"`
	CreatedBy       *UserReference          `json:"createdBy,omitempty"`
	Filename        string                  `json:"filename,omitempty"`
	Project         *SearchTmProjectDtoV2   `json:"project,omitempty"`
	Client          *SearchTmClientDtoV2    `json:"client,omitempty"`
	Domain          *SearchTmDomainDtoV2    `json:"domain,omitempty"`
	SubDomain       *SearchTmSubDomainDtoV2 `json:"subDomain,omitempty"`
	TagMetadata     []TagMetadata           `json:"tagMetadata,omitempty"`
	PreviousSegment string                  `json:"previousSegment,omitempty"`
	NextSegment     string                  `json:"nextSegment,omitempty"`
	Key             string                  `json:"key,omitempty"`
}
