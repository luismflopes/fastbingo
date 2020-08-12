package models

// CmsPage entity
type CmsPage struct {
	ID               int64  `json:"id"`
	Title            string `json:"title"`
	SeoName          string `json:"seo_name"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
}

func (me *CmsPage) GetEntityName() string {
	return "cms_pages"
}

func (me *CmsPage) ApplyTranslations(translations map[string]string) {
	me.Title = translations["title"]
	me.SeoName = translations["seo_name"]
	me.ShortDescription = translations["short_description"]
	me.LongDescription = translations["long_description"]
}

// CmsPage entity
type CmsPageUpdate struct {
	Title            *string `json:"title,omitempty"`
	ShortDescription *string `json:"short_description,omitempty"`
	LongDescription  *string `json:"long_description,omitempty"`
}

type CmsPageList struct {
	Total   int64     `json:"total,omitempty"`
	Results []CmsPage `json:"results,omitempty"`
}
