package rp

import (
	"database/sql"
	"testing"
	"log"
	"fmt"
)

func TestCreateFeature(t *testing.T) {
	db, err := sql.Open("pgx", "postgresql://postgres:postgres@localhost:5432/rightprism_db")
	if err != nil {
		log.Fatal(err)
	}

	feature := &Feature{}
	feature.ProductId = 2
	feature.FeatureTitle = "Some test feature 2"

	returnVal := feature.CreateFeature(db)

	if returnVal == "error" {
		t.Error("test failed due to error")
	} else {
		fmt.Println("TestCreateFeature", returnVal)
	}
}