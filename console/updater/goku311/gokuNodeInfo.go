package goku311

import (
	SQL "database/sql"
)

const gokuNodeInfoSQL = `DROP TABLE  IF EXISTS goku_node_info_new;
CREATE TABLE "goku_node_info_new" (
  "nodeID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "createTime" text,
  "updateTime" text,
  "groupID" integer(11) NOT NULL DEFAULT 0,
  "nodeName" text(255) NOT NULL,
  "nodeStatus" integer(11) NOT NULL,
  "version" text(255),
  "sshAddress" text(255) DEFAULT 22,
  "sshUserName" text(255),
  "sshPassword" text(255),
  "gatewayPath" text(255),
  "sshKey" text,
  "authMethod" integer(4) NOT NULL DEFAULT 0,
  "clusterID" integer(11) NOT NULL DEFAULT 0,
  "listenAddress" text(22) NOT NULL DEFAULT '',
  "adminAddress" text(22) NOT NULL DEFAULT '',
  "nodeKey" TEXT(32) NOT NULL DEFAULT ''
);

CREATE UNIQUE INDEX "nodeKey_new" 
ON "goku_node_info_new" (
  "nodeKey" ASC
);`

func createGokuNodeInfo(db *SQL.DB) error {
	_, err := db.Exec(gokuNodeInfoSQL)
	if err != nil {
		return err
	}

	sql := "INSERT INTO goku_node_info_new (`nodeID`,`createTime`,`updateTime`,`groupID`,`nodeName`,`nodeStatus`,`version`,`sshAddress`,`sshUserName`,`sshPassword`,`gatewayPath`,`sshKey`,`authMethod`,`clusterID`,`listenAddress`,`adminAddress`,`nodeKey`) SELECT `nodeID`,`createTime`,`updateTime`,`groupID`,`nodeName`,`nodeStatus`,`version`,`sshPort`,`userName`,`password`,`gatewayPath`,`key`,`authMethod`,`clusterID`,`nodeIP` || ':' || `nodePort`,`nodeIP` || ':' || `nodePort`,`nodeID` || `nodeIP` || ':' || `nodePort` FROM goku_node_info;"
	_, err = db.Exec(sql)
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS goku_node_info")
	if err != nil {
		return err
	}
	_, err = db.Exec("ALTER TABLE goku_node_info_new RENAME TO goku_node_info")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE  IF EXISTS goku_node_info_new")
	if err != nil {
		return err
	}

	return nil
}
