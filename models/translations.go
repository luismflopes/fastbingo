package models

// TranslationItem provides
type TranslationItem struct {
	Locale     string            `db:"locale"`
	EntityType string            `db:"entity_type"`
	EntityID   int64             `db:"entity_id"`
	Fields     map[string]string `db:"field"`
}
