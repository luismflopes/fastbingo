package models

// Product entity
type Product struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	SeoName          string `json:"seo_name"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
}

func (me *Product) GetEntityName() string {
	return "products"
}

func (me *Product) ApplyTranslations(translations map[string]string) {
	me.Name = translations["name"]
	me.SeoName = translations["seo_name"]
	me.ShortDescription = translations["short_description"]
	me.LongDescription = translations["long_description"]
}

// Product entity
type ProductUpdate struct {
	Name             *string `json:"name,omitempty"`
	ShortDescription *string `json:"short_description,omitempty"`
	LongDescription  *string `json:"long_description,omitempty"`
}

type ProductList struct {
	Total   int64     `json:"total,omitempty"`
	Results []Product `json:"results,omitempty"`
}
