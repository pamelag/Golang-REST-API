package rp

import(
	"database/sql"
)

type Feature struct {
	FeatureId int "json:featureId"
	RpToken string "json:rpToken"
	ProductId int "json:productId"
	FeatureTitle string "json:featureTitle"
	FeatureDescription string "json:featureDescription"
}

func (f *Feature) CreateFeature(database *sql.DB) string {

	uuid, err := GenerateId()

	if err != nil {
		return "uuid error"
	} 

	f.RpToken = uuid;

	tx, errtx := database.Begin()

	defer tx.Rollback()

	sql := "insert into feature (rp_token, product_id, feature_title, feature_description) VALUES ($1, $2, $3, $4)"
	stmt, _ := tx.Prepare(sql)

	defer stmt.Close()

	_, inserterr := stmt.Exec(f.RpToken, f.ProductId, f.FeatureTitle, f.FeatureDescription)

	if inserterr != nil {
	
		return "error"
	
	} else {

		errtx = tx.Commit()

		if errtx != nil {
			return "error"
		} else {
			
			err2 := database.QueryRow("select feature_id, rp_token, product_id, feature_title, feature_description from feature where rp_token=$1", f.RpToken).Scan(&f.FeatureId, &f.RpToken, &f.ProductId, &f.FeatureTitle, &f.FeatureDescription)
			
			if err2 != nil {
				return "error"
			} else {
				return "success"
			}
		}

		
	}
	
}