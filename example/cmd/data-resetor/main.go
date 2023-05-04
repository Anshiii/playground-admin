package main

import (
	"github.com/anshiii/playground-admin/example/admin"
)

func main() {
	db := admin.ConnectDB()
	tbs := admin.GetNonIgnoredTableNames()
	admin.EmptyDB(db, tbs)
	admin.InitDB(db, tbs)
	return
}
