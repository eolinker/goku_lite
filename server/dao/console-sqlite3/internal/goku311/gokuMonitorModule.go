package goku311

import SQL "database/sql"

const gokuMonitorModuleSQL = `DROP TABLE  IF EXISTS goku_monitor_module;
CREATE TABLE "goku_monitor_module" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" TEXT NOT NULL,
  "config" TEXT NOT NULL,
  "moduleStatus" integer NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX "moduleName"
ON "goku_monitor_module" (
  "name" ASC
);`

func createGokuMonitorModule(db *SQL.DB) error {
	_, err := db.Exec(gokuMonitorModuleSQL)
	if err != nil {
		return err
	}
	return nil
}
