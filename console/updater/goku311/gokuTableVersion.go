package goku311

import SQL "database/sql"

const gokuTableVersionSQL = `CREATE TABLE "goku_table_version" (
  "tableID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "tableName" TEXT NOT NULL,
  "version" TEXT NOT NULL
);

CREATE UNIQUE INDEX "tableName"
ON "goku_table_version" (
  "tableName"
);`

func createGokuTableVersion(db *SQL.DB) error {
	_, err := db.Exec(gokuTableVersionSQL)
	if err != nil {
		return err
	}

	return nil
}
