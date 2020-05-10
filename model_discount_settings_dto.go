/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DiscountSettingsDto struct {
	Repetition float64 `json:"repetition,omitempty"`
	Tm101      float64 `json:"tm101,omitempty"`
	Tm100      float64 `json:"tm100,omitempty"`
	Tm95       float64 `json:"tm95,omitempty"`
	Tm85       float64 `json:"tm85,omitempty"`
	Tm75       float64 `json:"tm75,omitempty"`
	Tm50       float64 `json:"tm50,omitempty"`
	Tm0        float64 `json:"tm0,omitempty"`
	Mt100      float64 `json:"mt100,omitempty"`
	Mt95       float64 `json:"mt95,omitempty"`
	Mt85       float64 `json:"mt85,omitempty"`
	Mt75       float64 `json:"mt75,omitempty"`
	Mt50       float64 `json:"mt50,omitempty"`
	Mt0        float64 `json:"mt0,omitempty"`
	Nt100      float64 `json:"nt100,omitempty"`
	Nt99       float64 `json:"nt99,omitempty"`
	Nt85       float64 `json:"nt85,omitempty"`
	Nt75       float64 `json:"nt75,omitempty"`
	Nt50       float64 `json:"nt50,omitempty"`
	Nt0        float64 `json:"nt0,omitempty"`
}