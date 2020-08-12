package services

import (
	"fmt"
	"fastbingo/models"
	"strings"

	"github.com/jinzhu/gorm"
)

// TranslationService Decoupled translations handling
type TranslationService struct {
	Storage *gorm.DB
}

func (me *TranslationService) CheckLocale(locale string) (string, error) {

	type Locale struct {
		Locale string
	}

	var l Locale

	me.Storage.Raw(`SELECT locale FROM locales WHERE locale=?
	UNION ALL 
	SELECT locale FROM locales WHERE is_default=1 
	LIMIT 1`, locale).Scan(&l)

	return l.Locale, nil
}

func (me *TranslationService) AddDefaultTranslations(ti *models.TranslationItem) bool {

	params := []interface{}{
		ti.EntityType,
		ti.EntityID,
	}
	subQueries := []string{}
	for key, value := range ti.Fields {
		subQueries = append(subQueries, "SELECT ? as field, ? as translation ")
		params = append(params, key)
		params = append(params, value)
	}

	query := fmt.Sprintf(`INSERT INTO translations(locale,entity_type,entity_id,field,translation)
				SELECT l.locale, ?, ?, t.field, t.translation FROM locales l, (
				%s
			) as t
			`, strings.Join(subQueries, "UNION "))

	err := me.Storage.Exec(query, params...).Error
	if err != nil {
		panic(err)
	}

	return true
}

func (me *TranslationService) UpdateTranslation(ti *models.TranslationItem) bool {

	query := "UPDATE translations SET translation=? WHERE locale=? AND entity_type=? AND entity_id=? AND field=?"

	for key, value := range ti.Fields {
		params := []interface{}{}
		params = append(params, value)
		params = append(params, ti.Locale)
		params = append(params, ti.EntityType)
		params = append(params, ti.EntityID)
		params = append(params, key)

		err := me.Storage.Exec(query, params...).Error
		if err != nil {
			panic(err)
		}
	}

	return true
}

func (me *TranslationService) DeleteTranslations(entityType string, entityID int64) bool {

	query := "DELETE FROM translations WHERE entity_type=? AND entity_id=?"
	err := me.Storage.Exec(query, entityType, entityID).Error
	if err != nil {
		panic(err)
	}

	return true
}

func (me *TranslationService) GetTranslations(locale, entityType string, entityIds []int64) map[int64]map[string]string {

	query := "SELECT entity_id, field, translation FROM translations WHERE locale=? AND entity_type=? AND entity_id IN(?)"
	rows, err := me.Storage.Raw(query, locale, entityType, entityIds).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	t := make(map[int64]map[string]string)
	var id int64
	var field, translation string
	for rows.Next() {
		id = 0
		field = ""
		translation = ""
		rows.Scan(&id, &field, &translation)
		if _, ok := t[id]; !ok {
			t[id] = map[string]string{}
		}

		t[id][field] = translation
	}

	return t
}
