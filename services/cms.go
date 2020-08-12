package services

import (
	"encoding/json"
	"fastbingo/models"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

// type ProductServiceContract interface {
// 	ListProduct() ([]*models.CmsPage, error)
// 	DetailProduct(slug string) (*models.CmsPage, error)
// 	CreateProduct(req *models.CmsPageCreateRequest)
// }

type CmsService struct {
	Storage            *gorm.DB
	TranslationService *TranslationService
}

// // ListProduct fetch a list of products translated with the requested locale.
// // Details filter here....
// func (me *CmsService) ListProduct(f *models.Filters, locale string) *models.CmsPageList {

// 	results := models.CmsPageList{}

// 	err := me.Storage.Model(&models.CmsPage{}).Where(f.FilterFields, f.FilterValues...).Count(&results.Total).Error
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = me.Storage.Where(f.FilterFields, f.FilterValues...).Limit(f.Limit).Offset(f.Offset).Find(&results.Results).Error
// 	if err != nil {
// 		panic(err)
// 	}

// 	ids := []int64{}
// 	for _, p := range results.Results {
// 		ids = append(ids, p.ID)
// 	}

// 	if len(ids) > 0 {
// 		t := me.TranslationService.GetTranslations(locale, "products", ids)
// 		for _, p := range results.Results {
// 			if _, ok := t[p.ID]; ok {
// 				p.ApplyTranslations(t[p.ID])
// 			}
// 		}
// 	}

// 	return &results
// }

// DetailPage fetch a product by its seo name translated with the requested locale
func (me *CmsService) DetailPage(slug, locale string) *models.CmsPage {

	p := models.CmsPage{}

	query := `select p.*
	from translations t
	inner join cms_pages p ON p.id=t.entity_id
	where t.field='seo_name' AND t.translation=?`

	err := me.Storage.Raw(query, slug).Scan(&p).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			panic(models.AppError{Code: models.ErrorNotFound, Message: "Cms Page not found."})
		}

		panic(err.Error())
	}

	t := me.TranslationService.GetTranslations(locale, p.GetEntityName(), []int64{p.ID})
	if _, ok := t[p.ID]; ok {
		p.ApplyTranslations(t[p.ID])
	}

	return &p
}

// CreatePage creates a cms page and initializes all translations with default values
func (me *CmsService) CreatePage(p *models.CmsPage) *models.CmsPage {

	p.SeoName = slug.Make(p.Title)

	err := me.Storage.Create(&p).Error
	if err != nil {
		panic(err)
	}

	ti := models.TranslationItem{
		EntityType: p.GetEntityName(),
		EntityID:   p.ID,
		Fields: map[string]string{
			"title":             p.Title,
			"seo_name":          p.SeoName,
			"short_description": p.ShortDescription,
			"long_description":  p.LongDescription,
		},
	}

	me.TranslationService.AddDefaultTranslations(&ti)

	return p
}

// UpdatePage Updates a cms page on the requested language
func (me *CmsService) UpdatePage(id int64, pu *models.CmsPageUpdate, locale string) *models.CmsPage {

	p := models.CmsPage{}
	err := me.Storage.First(&p, id).Error
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(pu)
	json.Unmarshal(b, &p)

	p.SeoName = slug.Make(p.Title)

	err = me.Storage.Save(&p).Error
	if err != nil {
		panic(err)
	}

	ti := models.TranslationItem{
		Locale:     locale,
		EntityType: p.GetEntityName(),
		EntityID:   p.ID,
		Fields: map[string]string{
			"title":             p.Title,
			"seo_name":          p.SeoName,
			"short_description": p.ShortDescription,
			"long_description":  p.LongDescription,
		},
	}

	me.TranslationService.UpdateTranslation(&ti)

	return &p
}

// DeleteProduct removes permanently a cms page and its translations
func (me *CmsService) DeletePage(id int64) *models.CmsPage {

	p := models.CmsPage{ID: id}
	db := me.Storage.Delete(&p)
	if db.Error != nil {
		panic(db.Error)
	}

	if db.RowsAffected == 0 {
		return nil
	}

	me.TranslationService.DeleteTranslations(p.GetEntityName(), id)

	return &p
}
