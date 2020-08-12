package storage

import (
	_ "github.com/mattn/go-sqlite3"
)

// type Sqlite3Adaper struct {
// 	Db *sqlx.DB
// }

// // InsertProduct do someting...
// func (me *Sqlite3Adaper) InsertProduct(p *entities.Product) {
// 	_, err := me.Db.NamedExec(`INSERT INTO products (name, seo_name, short_description, long_description)
// 	VALUES (:name, :seo_name, :short_description, :long_description)`, &p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// // InsertTranslations it does...
// func (me *Sqlite3Adaper) InsertTranslations(ti *entities.TranslationItem) {

// 	params := map[string]interface{}{
// 		"entity_type": ti.EntityType,
// 		"entity_id":   ti.EntityID,
// 	}
// 	subQueries := []string{}
// 	for key, value := range ti.Fields {
// 		paramKey := fmt.Sprintf("key_%s", key)
// 		paramValue := fmt.Sprintf("value_%s", key)
// 		subQueries = append(subQueries, fmt.Sprintf("SELECT :%s as field, :%s as translation ", paramKey, paramValue))

// 		params[paramKey] = key
// 		params[paramValue] = value
// 	}

// 	query := fmt.Sprintf(`INSERT INTO translations(locale,entity_type,entity_id,field,translation)
// 		SELECT l.locale, :entity_type, :entity_id, t.field, t.translation FROM locales l, (
// 		%s
// 	) as t
// 	`, strings.Join(subQueries, "UNION "))

// 	_, err := me.Db.NamedExec(query, params)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// func (me *Sqlite3Adaper) UpdateTranslations(ti *entities.TranslationItem) {

// 		params := map[string]interface{}{
// 		"entity_type": ti.EntityType,
// 		"entity_id":   ti.EntityID,
// 	}
// 	subQueries := []string{}
// 	for key, value := range ti.Fields {
// 		paramKey := fmt.Sprintf("key_%s", key)
// 		paramValue := fmt.Sprintf("value_%s", key)
// 		subQueries = append(subQueries, fmt.Sprintf("SELECT :%s as field, :%s as translation ", paramKey, paramValue))

// 		params[paramKey] = key
// 		params[paramValue] = value
// 	}

// 	query := fmt.Sprintf(`UPDATE translations (locale,entity_type,entity_id,field,translation)
// 		SELECT l.locale, :entity_type, :entity_id, t.field, t.translation FROM locales l, (
// 		%s
// 	) as t
// 	`, strings.Join(subQueries, "UNION "))

// 	_, err := me.Db.NamedExec(query, params)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// // func (me *Sqlite3Adaper) FindProductById(id int) *models.Product {
// // 	return me.FindProductsBy(map[string]interface{}{"id": id})[0]
// // }

// func (me *Sqlite3Adaper) FindProductsBy(criteria map[string]interface{}) []*entities.Product {

// 	where := ""
// 	for f := range criteria {
// 		where = where + fmt.Sprintf("%s=:%s", f, f)
// 	}

// 	if where != "" {
// 		where = fmt.Sprintf("WHERE %s", where)
// 	}

// 	r, err := me.Db.NamedQuery(fmt.Sprintf(`SELECT * FROM products %s`, where), criteria)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	items := make([]*entities.Product, 0, 10)
// 	for r.Rows.Next() {
// 		p := entities.Product{}
// 		err = r.StructScan(&p)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		items = append(items, &p)
// 	}

// 	return items
// }

// func (me *Sqlite3Adaper) Migrate() {
// 	_, err := me.Db.Exec("CREATE TABLE IF NOT EXISTS migrations (name VARCHAR NOT NULL);")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//Run new migrations
// 	files, err := ioutil.ReadDir("./migrations")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	migrated := 0
// 	for _, f := range files {
// 		sql, err := ioutil.ReadFile(fmt.Sprintf("./migrations/%s", f.Name()))
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		var found int
// 		row := me.Db.QueryRow("SELECT count(1) as found FROM migrations WHERE name=$1", f.Name())
// 		err = row.Scan(&found)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if found == 0 {
// 			_, err := me.Db.Exec(string(sql))
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			_, err = me.Db.Exec("INSERT INTO migrations(name) VALUES($1)", f.Name())
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			migrated++
// 			log.Println("Migrated: ", f.Name())
// 		}
// 	}

// 	if migrated == 0 {
// 		log.Println("No migrations")
// 	}
// }
