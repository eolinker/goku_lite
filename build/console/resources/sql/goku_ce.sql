USE goku_ce;
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for goku_admin
-- ----------------------------
DROP TABLE IF EXISTS `goku_admin`;
CREATE TABLE `goku_admin` (
  `userID` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `loginCall` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `loginPassword` varchar(255) NOT NULL,
  `userType` tinyint(4) NOT NULL DEFAULT '0',
  `groupID` int(11) NOT NULL DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `permissions` text,
  PRIMARY KEY (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_admin
-- ----------------------------

-- ----------------------------
-- Table structure for goku_balance
-- ----------------------------
DROP TABLE IF EXISTS `goku_balance`;
CREATE TABLE `goku_balance` (
  `balanceID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `balanceName` varchar(255) NOT NULL DEFAULT '',
  `serviceName` varchar(255) NOT NULL DEFAULT '',
  `balanceConfig` text,
  `createTime` timestamp NULL DEFAULT NULL,
  `updateTime` timestamp NULL DEFAULT NULL,
  `balanceDesc` varchar(255) DEFAULT NULL,
  `defaultConfig` text NOT NULL,
  `clusterConfig` text NOT NULL,
  `appName` varchar(255) NOT NULL DEFAULT '',
  `static` text,
  `staticCluster` text,
  PRIMARY KEY (`balanceID`),
  UNIQUE KEY `balanceName` (`balanceName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_balance
-- ----------------------------

-- ----------------------------
-- Table structure for goku_cluster
-- ----------------------------
DROP TABLE IF EXISTS `goku_cluster`;
CREATE TABLE `goku_cluster` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  `title` varchar(50) NOT NULL DEFAULT '',
  `note` varchar(255) DEFAULT NULL,
  `db` text,
  `redis` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `titel` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goku_cluster
-- ----------------------------

-- ----------------------------
-- Table structure for goku_config_log
-- ----------------------------
DROP TABLE IF EXISTS `goku_config_log`;
CREATE TABLE `goku_config_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  `enable` int(11) NOT NULL DEFAULT '0',
  `dir` varchar(255) NOT NULL DEFAULT 'logs/',
  `file` varchar(255) NOT NULL DEFAULT '',
  `period` varchar(10) NOT NULL DEFAULT '',
  `level` varchar(10) NOT NULL DEFAULT '',
  `fields` text NOT NULL,
  `expire` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goku_config_log
-- ----------------------------

-- ----------------------------
-- Table structure for goku_conn_plugin_api
-- ----------------------------
DROP TABLE IF EXISTS `goku_conn_plugin_api`;
CREATE TABLE `goku_conn_plugin_api` (
  `connID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `apiID` int(11) NOT NULL,
  `pluginName` varchar(255) NOT NULL,
  `pluginConfig` text,
  `strategyID` varchar(255) NOT NULL,
  `pluginInfo` text,
  `createTime` timestamp NULL DEFAULT NULL,
  `updateTime` timestamp NULL DEFAULT NULL,
  `pluginStatus` tinyint(4) DEFAULT NULL,
  `updateTag` varchar(32) DEFAULT NULL COMMENT '更新标识位',
  `updaterID` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`connID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_conn_plugin_api
-- ----------------------------

-- ----------------------------
-- Table structure for goku_conn_plugin_strategy
-- ----------------------------
DROP TABLE IF EXISTS `goku_conn_plugin_strategy`;
CREATE TABLE `goku_conn_plugin_strategy` (
  `connID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `strategyID` varchar(255) NOT NULL,
  `pluginName` varchar(255) NOT NULL,
  `pluginConfig` text,
  `pluginInfo` text,
  `createTime` timestamp NULL DEFAULT NULL,
  `updateTime` timestamp NULL DEFAULT NULL,
  `pluginStatus` tinyint(4) DEFAULT NULL,
  `updateTag` varchar(32) DEFAULT NULL COMMENT '更新标识位',
  `updaterID` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`connID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_conn_plugin_strategy
-- ----------------------------

-- ----------------------------
-- Table structure for goku_conn_strategy_api
-- ----------------------------
DROP TABLE IF EXISTS `goku_conn_strategy_api`;
CREATE TABLE `goku_conn_strategy_api` (
  `connID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `strategyID` varchar(255) NOT NULL,
  `apiID` int(11) NOT NULL,
  `apiMonitorStatus` int(11) NOT NULL DEFAULT '1',
  `strategyMonitorStatus` int(11) NOT NULL DEFAULT '1',
  `target` varchar(255) DEFAULT NULL,
  `updateTime` datetime DEFAULT NULL,
  PRIMARY KEY (`connID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_conn_strategy_api
-- ----------------------------

-- ----------------------------
-- Table structure for goku_gateway
-- ----------------------------
DROP TABLE IF EXISTS `goku_gateway`;
CREATE TABLE `goku_gateway` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `successCode` varchar(255) NOT NULL,
  `nodeUpdatePeriod` int(11) NOT NULL,
  `monitorUpdatePeriod` int(11) NOT NULL,
  `alertStatus` tinyint(4) NOT NULL DEFAULT '0',
  `sender` varchar(255) DEFAULT NULL,
  `senderPassword` varchar(255) DEFAULT NULL,
  `smtpAddress` varchar(255) DEFAULT NULL,
  `smtpPort` int(11) NOT NULL DEFAULT '25',
  `smtpProtocol` tinyint(4) NOT NULL DEFAULT '0',
  `monitorTimeout` tinyint(4) NOT NULL DEFAULT '5',
  `apiAlertInfo` text,
  `nodeAlertInfo` text,
  `redisAlertInfo` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_gateway
-- ----------------------------
INSERT INTO `goku_gateway` VALUES ('1', '200', '5', '30', '0', null, null, null, '25', '0', '5', '{\"alertAddr\":\"\",\"alertPeriodType\":0,\"logPath\":\"./log/apiAlert\",\"receiverList\":\"\"}', '{\"alertAddr\":\"\",\"logPath\":\"./log/nodeAlert\",\"receiverList\":\"\"}', '{\"alertAddr\":\"\",\"logPath\":\"./log/redisAlert\",\"receiverList\":\"\"}');

-- ----------------------------
-- Table structure for goku_gateway_alert
-- ----------------------------
DROP TABLE IF EXISTS `goku_gateway_alert`;
CREATE TABLE `goku_gateway_alert` (
  `alertID` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `requestURL` varchar(255) NOT NULL,
  `targetServer` varchar(255) NOT NULL,
  `alertPeriodType` tinyint(4) NOT NULL,
  `alertCount` int(11) NOT NULL,
  `updateTime` timestamp NULL DEFAULT NULL,
  `targetURL` varchar(255) NOT NULL,
  `clusterName` varchar(255) DEFAULT NULL,
  `nodeIP` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`alertID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_gateway_alert
-- ----------------------------

-- ----------------------------
-- Table structure for goku_gateway_api
-- ----------------------------
DROP TABLE IF EXISTS `goku_gateway_api`;
CREATE TABLE `goku_gateway_api` (
  `apiID` int(11) NOT NULL AUTO_INCREMENT,
  `groupID` int(11) NOT NULL,
  `projectID` int(11) NOT NULL,
  `requestURL` varchar(255) NOT NULL,
  `apiName` varchar(255) NOT NULL,
  `requestMethod` varchar(255) NOT NULL,
  `targetServer` varchar(255) DEFAULT NULL,
  `targetURL` varchar(255) DEFAULT NULL,
  `targetMethod` varchar(255) DEFAULT NULL,
  `isFollow` varchar(32) NOT NULL,
  `stripPrefix` varchar(32) DEFAULT NULL,
  `timeout` int(11) DEFAULT NULL,
  `retryCount` int(11) DEFAULT NULL,
  `createTime` timestamp NULL DEFAULT NULL,
  `updateTime` timestamp NULL DEFAULT NULL,
  `alertValve` int(11) NOT NULL DEFAULT '0',
  `monitorStatus` int(11) NOT NULL DEFAULT '0',
  `managerID` int(11) NOT NULL DEFAULT '0',
  `lastUpdateUserID` int(11) NOT NULL DEFAULT '0',
  `createUserID` int(11) NOT NULL,
  `balanceName` varchar(255) DEFAULT NULL,
  `protocol` varchar(20) DEFAULT NULL,
  `stripSlash` varchar(32) NOT NULL DEFAULT '1',
  PRIMARY KEY (`apiID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_gateway_api
-- ----------------------------

-- ----------------------------
-- Table structure for goku_gateway_api_group
-- ----------------------------
DROP TABLE IF EXISTS `goku_gateway_api_group`;
CREATE TABLE `goku_gateway_api_group` (
  `groupID` int(11) NOT NULL AUTO_INCREMENT,
  `projectID` int(11) NOT NULL,
  `groupName` varchar(255) NOT NULL,
  `groupPath` varchar(255) DEFAULT NULL,
  `groupDepth` tinyint(4) NOT NULL DEFAULT '1',
  `parentGroupID` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`groupID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_gateway_api_group
-- ----------------------------

-- ----------------------------
-- Table structure for goku_gateway_permission_group
-- ----------------------------
DROP TABLE IF EXISTS `goku_gateway_permission_group`;
CREATE TABLE `goku_gateway_permission_group` (
  `groupID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `groupName` varchar(255) NOT NULL,
  `permissions` text,
  PRIMARY KEY (`groupID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_gateway_permission_group
-- ----------------------------

-- ----------------------------
-- Table structure for goku_gateway_project
-- ----------------------------
DROP TABLE IF EXISTS `goku_gateway_project`;
CREATE TABLE `goku_gateway_project` (
  `projectID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `projectName` varchar(255) NOT NULL,
  `createTime` timestamp NULL DEFAULT NULL,
  `updateTime` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`projectID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_gateway_project
-- ----------------------------

-- ----------------------------
-- Table structure for goku_gateway_strategy
-- ----------------------------
DROP TABLE IF EXISTS `goku_gateway_strategy`;
CREATE TABLE `goku_gateway_strategy` (
  `strategyID` varchar(32) NOT NULL,
  `strategyName` varchar(255) NOT NULL,
  `updateTime` timestamp NULL DEFAULT NULL,
  `createTime` timestamp NULL DEFAULT NULL,
  `auth` varchar(255) NOT NULL DEFAULT '0',
  `groupID` int(11) NOT NULL DEFAULT '1',
  `monitorStatus` int(4) NOT NULL DEFAULT '0',
  `enableStatus` int(11) NOT NULL DEFAULT '1',
  `strategyType` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`strategyID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_gateway_strategy
-- ----------------------------
INSERT INTO `goku_gateway_strategy` VALUES ('tqvka3', '开放策略', '2019-02-20 09:59:18', '2019-02-20 09:59:21', '0', '1', '0', '1', '1');

-- ----------------------------
-- Table structure for goku_gateway_strategy_group
-- ----------------------------
DROP TABLE IF EXISTS `goku_gateway_strategy_group`;
CREATE TABLE `goku_gateway_strategy_group` (
  `groupID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `groupName` varchar(255) NOT NULL,
  `groupType` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`groupID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_gateway_strategy_group
-- ----------------------------
INSERT INTO `goku_gateway_strategy_group` VALUES ('1', '开放策略', '1');

-- ----------------------------
-- Table structure for goku_message
-- ----------------------------
DROP TABLE IF EXISTS `goku_message`;
CREATE TABLE `goku_message` (
  `msgID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `updateTime` timestamp NULL DEFAULT NULL,
  `msg` varchar(255) DEFAULT NULL,
  `msgType` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`msgID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_message
-- ----------------------------

-- ----------------------------
-- Table structure for goku_monitor_cluster
-- ----------------------------
DROP TABLE IF EXISTS `goku_monitor_cluster`;
CREATE TABLE `goku_monitor_cluster` (
  `recordID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `strategyID` varchar(20) NOT NULL,
  `apiID` int(11) NOT NULL,
  `clusterID` int(11) NOT NULL,
  `hour` int(11) NOT NULL,
  `gatewayRequestCount` int(11) NOT NULL DEFAULT '0',
  `gatewaySuccessCount` int(11) NOT NULL DEFAULT '0',
  `gatewayStatus2xxCount` int(11) NOT NULL DEFAULT '0',
  `gatewayStatus4xxCount` int(11) NOT NULL DEFAULT '0',
  `gatewayStatus5xxCount` int(11) NOT NULL DEFAULT '0',
  `proxyRequestCount` int(11) NOT NULL DEFAULT '0',
  `proxySuccessCount` int(11) NOT NULL DEFAULT '0',
  `proxyStatus2xxCount` int(11) NOT NULL DEFAULT '0',
  `proxyStatus4xxCount` int(11) NOT NULL DEFAULT '0',
  `proxyStatus5xxCount` int(11) NOT NULL DEFAULT '0',
  `proxyTimeoutCount` int(11) NOT NULL DEFAULT '0',
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`recordID`),
  UNIQUE KEY `key` (`strategyID`,`apiID`,`clusterID`,`hour`),
  KEY `strategyID_2` (`strategyID`),
  KEY `apiID` (`apiID`),
  KEY `clusterID` (`clusterID`),
  KEY `hour` (`hour`),
  KEY `strategy_api` (`strategyID`,`apiID`,`clusterID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_monitor_cluster
-- ----------------------------

-- ----------------------------
-- Table structure for goku_node_group
-- ----------------------------
DROP TABLE IF EXISTS `goku_node_group`;
CREATE TABLE `goku_node_group` (
  `groupID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `groupName` varchar(255) NOT NULL DEFAULT '',
  `groupType` tinyint(4) NOT NULL DEFAULT '0',
  `clusterID` int(11) NOT NULL,
  PRIMARY KEY (`groupID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_node_group
-- ----------------------------

-- ----------------------------
-- Table structure for goku_node_info
-- ----------------------------
DROP TABLE IF EXISTS `goku_node_info`;
CREATE TABLE `goku_node_info` (
  `nodeID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `nodeIP` varchar(255) NOT NULL,
  `updateStatus` tinyint(4) NOT NULL DEFAULT '0',
  `createTime` timestamp NULL DEFAULT NULL,
  `updateTime` timestamp NULL DEFAULT NULL,
  `groupID` int(11) NOT NULL,
  `nodeName` varchar(255) NOT NULL,
  `nodePort` varchar(255) DEFAULT NULL,
  `nodeStatus` varchar(255) DEFAULT NULL,
  `version` varchar(255) DEFAULT NULL,
  `sshPort` varchar(255) NOT NULL DEFAULT '22',
  `userName` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `gatewayPath` varchar(255) DEFAULT NULL,
  `key` text,
  `authMethod` tinyint(4) NOT NULL DEFAULT '0',
  `clusterID` int(11) NOT NULL,
  PRIMARY KEY (`nodeID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_node_info
-- ----------------------------

-- ----------------------------
-- Table structure for goku_plugin
-- ----------------------------
DROP TABLE IF EXISTS `goku_plugin`;
CREATE TABLE `goku_plugin` (
  `pluginID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pluginName` varchar(255) NOT NULL,
  `chineseName` varchar(255) DEFAULT NULL,
  `pluginStatus` tinyint(4) NOT NULL DEFAULT '0',
  `pluginPriority` int(4) NOT NULL DEFAULT '0',
  `pluginConfig` text,
  `pluginInfo` text,
  `isStop` tinyint(4) NOT NULL DEFAULT '0',
  `pluginType` tinyint(4) NOT NULL DEFAULT '0',
  `official` varchar(255) NOT NULL,
  `pluginDesc` varchar(255) DEFAULT NULL,
  `version` varchar(255) NOT NULL,
  `isCheck` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`pluginID`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for goku_redis_config
-- ----------------------------
DROP TABLE IF EXISTS `goku_redis_config`;
CREATE TABLE `goku_redis_config` (
  `id` int(11) unsigned NOT NULL,
  `cluster_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(20) NOT NULL DEFAULT '',
  `mod` varchar(10) NOT NULL DEFAULT '',
  `addrs` text NOT NULL,
  `password` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `cluster_id` (`cluster_id`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goku_redis_config
-- ----------------------------

-- ----------------------------
-- Table structure for goku_redis_info
-- ----------------------------
DROP TABLE IF EXISTS `goku_redis_info`;
CREATE TABLE `goku_redis_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `server` varchar(20) NOT NULL COMMENT 'ip:port',
  `info` text NOT NULL COMMENT 'info,json',
  `datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `server` (`server`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goku_redis_info
-- ----------------------------

-- ----------------------------
-- Table structure for goku_redis_memory
-- ----------------------------
DROP TABLE IF EXISTS `goku_redis_memory`;
CREATE TABLE `goku_redis_memory` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `server` varchar(20) DEFAULT NULL COMMENT 'ip:port',
  `used` int(11) DEFAULT NULL,
  `peak` int(11) DEFAULT NULL,
  `datetime` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `server` (`server`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goku_redis_memory
-- ----------------------------

-- ----------------------------
-- Table structure for goku_redis_server
-- ----------------------------
DROP TABLE IF EXISTS `goku_redis_server`;
CREATE TABLE `goku_redis_server` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `server` varchar(20) NOT NULL DEFAULT '' COMMENT 'ip:port',
  `password` varchar(20) DEFAULT NULL,
  `clusterID` int(11) NOT NULL DEFAULT '1',
  `status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `server` (`server`,`clusterID`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goku_redis_server
-- ----------------------------

-- ----------------------------
-- Table structure for goku_service_config
-- ----------------------------
DROP TABLE IF EXISTS `goku_service_config`;
CREATE TABLE `goku_service_config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `default` tinyint(4) DEFAULT NULL,
  `driver` varchar(20) NOT NULL DEFAULT '' COMMENT 'driver.name',
  `desc` text NOT NULL,
  `config` text NOT NULL,
  `clusterConfig` text NOT NULL,
  `healthCheck` tinyint(4) NOT NULL DEFAULT '0',
  `healthCheckPath` varchar(255) NOT NULL DEFAULT '/',
  `healthCheckPeriod` int(11) NOT NULL DEFAULT '5',
  `healthCheckCode` varchar(255) NOT NULL DEFAULT '200',
  `healthCheckTimeOut` int(11) NOT NULL DEFAULT '200',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goku_service_config
-- ----------------------------

-- ----------------------------
-- Table structure for goku_service_discovery
-- ----------------------------
DROP TABLE IF EXISTS `goku_service_discovery`;
CREATE TABLE `goku_service_discovery` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL,
  `type` varchar(20) DEFAULT NULL,
  `remark` varchar(500) DEFAULT NULL,
  `config` text,
  `default` varchar(255) DEFAULT NULL,
  `createTime` timestamp NULL DEFAULT NULL,
  `updateTime` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goku_service_discovery
-- ----------------------------

-- ----------------------------
-- Table structure for goku_table_update_record
-- ----------------------------
DROP TABLE IF EXISTS `goku_table_update_record`;
CREATE TABLE `goku_table_update_record` (
  `name` varchar(64) NOT NULL,
  `updateTime` datetime NOT NULL,
  `tableID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`tableID`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goku_table_update_record
-- ----------------------------

-- ----------------------------
-- Table structure for goku_version
-- ----------------------------
DROP TABLE IF EXISTS `goku_version`;
CREATE TABLE `goku_version` (
  `version` varchar(20) NOT NULL DEFAULT '',
  `sol` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goku_version
-- ----------------------------
INSERT INTO `goku_version` VALUES ('3.0.0', '0');
