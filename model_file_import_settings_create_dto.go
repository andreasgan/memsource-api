/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type FileImportSettingsCreateDto struct {
	InputCharset  string `json:"inputCharset,omitempty"`
	OutputCharset string `json:"outputCharset,omitempty"`
	ZipCharset    string `json:"zipCharset,omitempty"`
	// default: auto-detect
	FileFormat string `json:"fileFormat,omitempty"`
	// Default: false
	TargetLength bool `json:"targetLength,omitempty"`
	// default: 1000
	TargetLengthMax int32 `json:"targetLengthMax,omitempty"`
	// Default: false
	TargetLengthPercent bool `json:"targetLengthPercent,omitempty"`
	// default: 130
	TargetLengthPercentValue float64                     `json:"targetLengthPercentValue,omitempty"`
	SegmentationRuleId       int64                       `json:"segmentationRuleId,omitempty"`
	TargetSegmentationRuleId int64                       `json:"targetSegmentationRuleId,omitempty"`
	Android                  *AndroidSettingsDto         `json:"android,omitempty"`
	Csv                      *CsvSettingsDto             `json:"csv,omitempty"`
	Dita                     *DitaSettingsDto            `json:"dita,omitempty"`
	DocBook                  *DocBookSettingsDto         `json:"docBook,omitempty"`
	Doc                      *DocSettingsDto             `json:"doc,omitempty"`
	Html                     *HtmlSettingsDto            `json:"html,omitempty"`
	Idml                     *IdmlSettingsDto            `json:"idml,omitempty"`
	Json                     *JsonSettingsDto            `json:"json,omitempty"`
	Mac                      *MacSettingsDto             `json:"mac,omitempty"`
	Md                       *MdSettingsDto              `json:"md,omitempty"`
	Mif                      *MifSettingsDto             `json:"mif,omitempty"`
	MultilingualXls          *MultilingualXlsSettingsDto `json:"multilingualXls,omitempty"`
	MultilingualCsv          *MultilingualCsvSettingsDto `json:"multilingualCsv,omitempty"`
	MultilingualXml          *MultilingualXmlSettingsDto `json:"multilingualXml,omitempty"`
	Pdf                      *PdfSettingsDto             `json:"pdf,omitempty"`
	Php                      *PhpSettingsDto             `json:"php,omitempty"`
	Po                       *PoSettingsDto              `json:"po,omitempty"`
	Ppt                      *PptSettingsDto             `json:"ppt,omitempty"`
	Properties               *PropertiesSettingsDto      `json:"properties,omitempty"`
	Psd                      *PsdSettingsDto             `json:"psd,omitempty"`
	QuarkTag                 *QuarkTagSettingsDto        `json:"quarkTag,omitempty"`
	Resx                     *ResxSettingsDto            `json:"resx,omitempty"`
	SdlXlf                   *SdlXlfSettingsDto          `json:"sdlXlf,omitempty"`
	TmMatch                  *TmMatchSettingsDto         `json:"tmMatch,omitempty"`
	Ttx                      *TtxSettingsDto             `json:"ttx,omitempty"`
	Txt                      *TxtSettingsDto             `json:"txt,omitempty"`
	Xlf2                     *Xlf2SettingsDto            `json:"xlf2,omitempty"`
	Xlf                      *XlfSettingsDto             `json:"xlf,omitempty"`
	Xls                      *XlsSettingsDto             `json:"xls,omitempty"`
	Xml                      *XmlSettingsDto             `json:"xml,omitempty"`
	Yaml                     *YamlSettingsDto            `json:"yaml,omitempty"`
}
