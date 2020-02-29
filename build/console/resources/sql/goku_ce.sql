PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for goku_admin
-- ----------------------------
DROP TABLE IF EXISTS "goku_admin";
CREATE TABLE "goku_admin" (
                            "userID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                            "loginCall" text(255) NOT NULL,
                            "loginPassword" text(255) NOT NULL,
                            "userType" integer(4) NOT NULL DEFAULT 0,
                            "groupID" integer(11) NOT NULL DEFAULT 0,
                            "remark" text(255),
                            "permissions" text
);

-- ----------------------------
-- Table structure for goku_balance
-- ----------------------------
DROP TABLE IF EXISTS "goku_balance";
CREATE TABLE "goku_balance" (
                              "balanceID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                              "balanceName" text(255) NOT NULL,
                              "serviceName" text(255) NOT NULL,
                              "balanceConfig" text,
                              "createTime" text,
                              "updateTime" text,
                              "balanceDesc" text(255),
                              "defaultConfig" text NOT NULL,
                              "clusterConfig" text NOT NULL DEFAULT '',
                              "appName" text(255) NOT NULL DEFAULT '',
                              "static" text,
                              "staticCluster" text
);

-- ----------------------------
-- Table structure for goku_cluster
-- ----------------------------
DROP TABLE IF EXISTS "goku_cluster";
CREATE TABLE "goku_cluster" (
                              "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                              "name" text(20) NOT NULL,
                              "title" text(50) NOT NULL,
                              "note" text(255),
                              "db" text,
                              "redis" text
);

-- ----------------------------
-- Table structure for goku_config_log
-- ----------------------------
DROP TABLE IF EXISTS "goku_config_log";
CREATE TABLE "goku_config_log" (
                                 "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                 "name" text(20) NOT NULL,
                                 "enable" integer(11) NOT NULL DEFAULT 1,
                                 "dir" text(255) NOT NULL,
                                 "file" text(255) NOT NULL,
                                 "period" text(10) NOT NULL,
                                 "level" text(10) NOT NULL,
                                 "fields" text NOT NULL,
                                 "expire" integer(11) NOT NULL DEFAULT 3
);

-- ----------------------------
-- Table structure for goku_conn_plugin_api
-- ----------------------------
DROP TABLE IF EXISTS "goku_conn_plugin_api";
CREATE TABLE "goku_conn_plugin_api" (
                                      "connID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                      "apiID" integer(11) NOT NULL,
                                      "pluginName" text(255) NOT NULL,
                                      "pluginConfig" text,
                                      "strategyID" text(255) NOT NULL,
                                      "pluginInfo" text,
                                      "createTime" text,
                                      "updateTime" text,
                                      "pluginStatus" integer(4),
                                      "updateTag" text(32),
                                      "updaterID" integer(11) NOT NULL DEFAULT 0
);

-- ----------------------------
-- Table structure for goku_conn_plugin_strategy
-- ----------------------------
DROP TABLE IF EXISTS "goku_conn_plugin_strategy";
CREATE TABLE "goku_conn_plugin_strategy" (
                                           "connID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                           "strategyID" text(255) NOT NULL,
                                           "pluginName" text(255) NOT NULL,
                                           "pluginConfig" text,
                                           "pluginInfo" text,
                                           "createTime" text,
                                           "updateTime" text,
                                           "pluginStatus" integer(4),
                                           "updateTag" text(32),
                                           "updaterID" integer(11) NOT NULL DEFAULT 0
);

-- ----------------------------
-- Table structure for goku_conn_strategy_api
-- ----------------------------
DROP TABLE IF EXISTS "goku_conn_strategy_api";
CREATE TABLE "goku_conn_strategy_api" (
                                        "connID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                        "strategyID" text(255) NOT NULL,
                                        "apiID" integer(11) NOT NULL,
                                        "apiMonitorStatus" integer(11) NOT NULL DEFAULT 0,
                                        "strategyMonitorStatus" integer(11) NOT NULL DEFAULT 0,
                                        "target" text(255),
                                        "updateTime" text
);

-- ----------------------------
-- Table structure for goku_gateway
-- ----------------------------
DROP TABLE IF EXISTS "goku_gateway";
CREATE TABLE "goku_gateway" (
                              "id" integer(11) NOT NULL,
                              "successCode" text(255) NOT NULL,
                              "nodeUpdatePeriod" integer(11) NOT NULL,
                              "monitorUpdatePeriod" integer(11) NOT NULL,
                              "alertStatus" integer(4) NOT NULL,
                              "alertPeriodType" integer(4) NOT NULL,
                              "alertAddress" text(255),
                              "alertLogPath" text(255),
                              "sender" text(255),
                              "senderPassword" text(255),
                              "smtpAddress" text(255),
                              "smtpPort" integer(11) NOT NULL,
                              "smtpProtocol" integer(4) NOT NULL,
                              "receiverList" text(255),
                              "monitorTimeout" integer(4) NOT NULL,
                              "apiAlertInfo" text,
                              "nodeAlertInfo" text,
                              "redisAlertInfo" text,
                              "versionID" INTEGER NOT NULL DEFAULT 0,
                              PRIMARY KEY ("id")
);

INSERT INTO "goku_gateway" VALUES (1, 200, 1, 30, 0, 0, NULL, NULL, NULL, NULL, NULL, 25, 0, NULL, 0, NULL, NULL, NULL, 0);

-- ----------------------------
-- Table structure for goku_gateway_alert
-- ----------------------------
DROP TABLE IF EXISTS "goku_gateway_alert";
CREATE TABLE "goku_gateway_alert" (
                                    "alertID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                    "requestURL" text(255) NOT NULL,
                                    "targetServer" text(255) NOT NULL,
                                    "alertPeriodType" integer(4) NOT NULL,
                                    "alertCount" integer(11) NOT NULL,
                                    "updateTime" text,
                                    "targetURL" text(255) NOT NULL,
                                    "clusterName" text(255),
                                    "nodeIP" text(255)
);

-- ----------------------------
-- Table structure for goku_gateway_api
-- ----------------------------
DROP TABLE IF EXISTS "goku_gateway_api";
CREATE TABLE "goku_gateway_api" (
                                  "apiID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                  "groupID" integer(11) NOT NULL,
                                  "projectID" integer(11) NOT NULL,
                                  "requestURL" text(255) NOT NULL,
                                  "apiName" text(255) NOT NULL,
                                  "requestMethod" text(255) NOT NULL,
                                  "targetServer" text(255),
                                  "targetURL" text(255),
                                  "targetMethod" text(255),
                                  "isFollow" text(32) NOT NULL,
                                  "stripPrefix" text(32),
                                  "timeout" integer(11),
                                  "retryCount" integer(11),
                                  "createTime" text,
                                  "updateTime" text,
                                  "alertValve" integer(11) NOT NULL DEFAULT 0,
                                  "monitorStatus" integer(11) NOT NULL DEFAULT 1,
                                  "managerID" integer(11) NOT NULL,
                                  "lastUpdateUserID" integer(11) NOT NULL,
                                  "createUserID" integer(11) NOT NULL,
                                  "balanceName" text(255),
                                  "protocol" text(20),
                                  "stripSlash" text(32),
                                  "apiType" integer NOT NULL DEFAULT 0,
                                  "responseDataType" text NOT NULL DEFAULT origin,
                                  "linkApis" TEXT,
                                  "staticResponse" TEXT
);

-- ----------------------------
-- Table structure for goku_gateway_api_group
-- ----------------------------
DROP TABLE IF EXISTS "goku_gateway_api_group";
CREATE TABLE "goku_gateway_api_group" (
                                        "groupID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                        "projectID" integer(11) NOT NULL,
                                        "groupName" text(255) NOT NULL,
                                        "groupPath" text(255),
                                        "groupDepth" text(255),
                                        "parentGroupID" integer(11) NOT NULL DEFAULT 0
);

-- ----------------------------
-- Table structure for goku_gateway_permission_group
-- ----------------------------
DROP TABLE IF EXISTS "goku_gateway_permission_group";
CREATE TABLE "goku_gateway_permission_group" (
                                               "groupID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                               "groupName" text(255) NOT NULL,
                                               "permissions" text
);

-- ----------------------------
-- Table structure for goku_gateway_project
-- ----------------------------
DROP TABLE IF EXISTS "goku_gateway_project";
CREATE TABLE "goku_gateway_project" (
                                      "projectID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                      "projectName" text(255) NOT NULL,
                                      "createTime" text,
                                      "updateTime" text
);

-- ----------------------------
-- Table structure for goku_gateway_strategy
-- ----------------------------
DROP TABLE IF EXISTS "goku_gateway_strategy";
CREATE TABLE "goku_gateway_strategy" (
                                       "strategyID" text(32) NOT NULL,
                                       "strategyName" text(255) NOT NULL,
                                       "updateTime" text,
                                       "createTime" text,
                                       "auth" text(255),
                                       "groupID" integer(11) NOT NULL DEFAULT 0,
                                       "monitorStatus" integer(4) NOT NULL DEFAULT 0,
                                       "enableStatus" integer(11) NOT NULL DEFAULT 0,
                                       "strategyType" integer(11) NOT NULL DEFAULT 0,
                                       PRIMARY KEY ("strategyID")
);

-- ----------------------------
-- Records of "goku_gateway_strategy"
-- ----------------------------
INSERT INTO "goku_gateway_strategy" VALUES ('RGAtKBd', '开放策略', '2019-10-17 00:00:00', '2019-10-17 00:00:00', NULL, 0, 0, 0, 1);

-- ----------------------------
-- Table structure for goku_gateway_strategy_group
-- ----------------------------
DROP TABLE IF EXISTS "goku_gateway_strategy_group";
CREATE TABLE "goku_gateway_strategy_group" (
                                             "groupID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                             "groupName" text(255) NOT NULL,
                                             "groupType" integer(11) NOT NULL DEFAULT 0
);

-- ----------------------------
-- Records of "goku_gateway_strategy_group"
-- ----------------------------
INSERT INTO "goku_gateway_strategy_group" VALUES (1, '开放分组', 1);

-- ----------------------------
-- Table structure for goku_gateway_version_config
-- ----------------------------
DROP TABLE IF EXISTS "goku_gateway_version_config";
CREATE TABLE "goku_gateway_version_config" (
                                             "versionID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                             "name" TEXT NOT NULL,
                                             "version" TEXT,
                                             "remark" TEXT,
                                             "createTime" TEXT,
                                             "updateTime" TEXT,
                                             "publishTime" TEXT,
                                             "config" TEXT,
                                             "balanceConfig" TEXT,
                                             "discoverConfig" TEXT
);

-- ----------------------------
-- Table structure for goku_monitor_cluster
-- ----------------------------
DROP TABLE IF EXISTS "goku_monitor_cluster";
CREATE TABLE "goku_monitor_cluster" (
                                      "recordID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                      "strategyID" text(20) NOT NULL,
                                      "apiID" integer(11) NOT NULL,
                                      "clusterID" integer(11) NOT NULL,
                                      "hour" integer(11) NOT NULL,
                                      "gatewayRequestCount" integer(11) NOT NULL,
                                      "gatewaySuccessCount" integer(11) NOT NULL,
                                      "gatewayStatus2xxCount" integer(11) NOT NULL,
                                      "gatewayStatus4xxCount" integer(11) NOT NULL,
                                      "gatewayStatus5xxCount" integer(11) NOT NULL,
                                      "proxyRequestCount" integer(11) NOT NULL,
                                      "proxySuccessCount" integer(11) NOT NULL,
                                      "proxyStatus2xxCount" integer(11) NOT NULL,
                                      "proxyStatus4xxCount" integer(11) NOT NULL,
                                      "proxyStatus5xxCount" integer(11) NOT NULL,
                                      "proxyTimeoutCount" integer(11) NOT NULL,
                                      "updateTime" text NOT NULL
);

-- ----------------------------
-- Table structure for goku_node_group
-- ----------------------------
DROP TABLE IF EXISTS "goku_node_group";
CREATE TABLE "goku_node_group" (
                                 "groupID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                 "groupName" text(255) NOT NULL,
                                 "groupType" integer(4) NOT NULL,
                                 "clusterID" integer(11) NOT NULL
);

-- ----------------------------
-- Table structure for goku_node_info
-- ----------------------------
DROP TABLE IF EXISTS "goku_node_info";
CREATE TABLE "goku_node_info" (
                                "nodeID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                "nodeIP" text(255) NOT NULL,
                                "updateStatus" integer(4) NOT NULL DEFAULT 0,
                                "createTime" text,
                                "updateTime" text,
                                "groupID" integer(11) NOT NULL DEFAULT 0,
                                "nodeName" text(255) NOT NULL,
                                "nodePort" text(255),
                                "nodeStatus" integer(11) NOT NULL,
                                "version" text(255),
                                "sshPort" text(255) DEFAULT 22,
                                "userName" text(255),
                                "password" text(255),
                                "gatewayPath" text(255),
                                "key" text,
                                "authMethod" integer(4) NOT NULL DEFAULT 0,
                                "clusterID" integer(11) NOT NULL DEFAULT 0
);

-- ----------------------------
-- Table structure for goku_plugin
-- ----------------------------
DROP TABLE IF EXISTS "goku_plugin";
CREATE TABLE "goku_plugin" (
                             "pluginID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                             "pluginName" text(255) NOT NULL,
                             "chineseName" text(255),
                             "pluginStatus" integer(4) NOT NULL,
                             "pluginPriority" integer(4) NOT NULL,
                             "pluginConfig" text,
                             "pluginInfo" text,
                             "isStop" integer(4) NOT NULL,
                             "pluginType" integer(4) NOT NULL,
                             "official" text(255) NOT NULL,
                             "pluginDesc" text(255),
                             "version" text(255) NOT NULL,
                             "isCheck" integer(4) NOT NULL
);


-- ----------------------------
-- Table structure for goku_service_config
-- ----------------------------
DROP TABLE IF EXISTS "goku_service_config";
CREATE TABLE "goku_service_config" (
                                     "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                     "name" text(255) NOT NULL,
                                     "default" integer(4),
                                     "driver" text(20) NOT NULL,
                                     "desc" text NOT NULL,
                                     "config" text NOT NULL,
                                     "clusterConfig" text NOT NULL,
                                     "healthCheck" integer(4) NOT NULL,
                                     "healthCheckPath" text(255) NOT NULL,
                                     "healthCheckPeriod" integer(11) NOT NULL,
                                     "healthCheckCode" text(255) NOT NULL,
                                     "healthCheckTimeOut" integer(11) NOT NULL,
                                     "createTime" text NOT NULL,
                                     "updateTime" text NOT NULL
);

-- ----------------------------
-- Table structure for goku_service_discovery
-- ----------------------------
DROP TABLE IF EXISTS "goku_service_discovery";
CREATE TABLE "goku_service_discovery" (
                                        "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
                                        "name" text(30),
                                        "type" text(20),
                                        "remark" text(500),
                                        "config" text,
                                        "default" text(255),
                                        "createTime" text,
                                        "updateTime" text
);

-- ----------------------------
-- Table structure for goku_table_update_record
-- ----------------------------
DROP TABLE IF EXISTS "goku_table_update_record";
CREATE TABLE "goku_table_update_record" (
                                          "name" text(64) NOT NULL,
                                          "updateTime" text NOT NULL,
                                          "tableID" integer NOT NULL PRIMARY KEY AUTOINCREMENT
);


-- ----------------------------
-- Records of "sqlite_sequence"
-- ----------------------------
INSERT INTO "sqlite_sequence" VALUES ('goku_admin', 1);
INSERT INTO "sqlite_sequence" VALUES ('goku_plugin', 30);
INSERT INTO "sqlite_sequence" VALUES ('goku_balance', 0);
INSERT INTO "sqlite_sequence" VALUES ('goku_config_log', 0);
INSERT INTO "sqlite_sequence" VALUES ('goku_gateway_strategy_group', 1);

-- ----------------------------
-- Auto increment value for goku_admin
-- ----------------------------
UPDATE "sqlite_sequence" SET seq = 1 WHERE name = 'goku_admin';

-- ----------------------------
-- Auto increment value for goku_balance
-- ----------------------------

-- ----------------------------
-- Indexes structure for table goku_balance
-- ----------------------------
CREATE INDEX "balanceName"
  ON "goku_balance" (
                     "balanceName" ASC
    );

-- ----------------------------
-- Indexes structure for table goku_cluster
-- ----------------------------
CREATE INDEX "name"
  ON "goku_cluster" (
                     "name" ASC
    );

-- ----------------------------
-- Auto increment value for goku_config_log
-- ----------------------------

-- ----------------------------
-- Auto increment value for goku_gateway_strategy_group
-- ----------------------------
UPDATE "sqlite_sequence" SET seq = 1 WHERE name = 'goku_gateway_strategy_group';

-- ----------------------------
-- Auto increment value for goku_plugin
-- ----------------------------
UPDATE "sqlite_sequence" SET seq = 30 WHERE name = 'goku_plugin';

PRAGMA foreign_keys = true;
