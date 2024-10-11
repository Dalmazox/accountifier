package repositories

import (
	"fmt"
	"log"
	"os"
)

func LoadQuery(queryName string) string {
	content, err := os.ReadFile(fmt.Sprintf("./internal/repositories/sql/%s.sql", queryName))

	if err != nil {
		log.Fatalf("sql file %s not found", queryName)
	}

	return string(content)
}
