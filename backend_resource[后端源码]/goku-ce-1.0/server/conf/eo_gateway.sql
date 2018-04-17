
USE $mysql_dbname
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for eo_admin
-- ----------------------------
DROP TABLE IF EXISTS `eo_admin`;
CREATE TABLE `eo_admin` (
  `userID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `loginCall` varchar(255) NOT NULL,
  `loginPassword` varchar(255) NOT NULL,
  `type` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`userID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for eo_conn_api_group_strategy
-- ----------------------------
DROP TABLE IF EXISTS `eo_conn_api_group_strategy`;
CREATE TABLE `eo_conn_api_group_strategy` (
  `connID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `groupID` int(11) NOT NULL,
  `strategyID` int(11) NOT NULL,
  PRIMARY KEY (`connID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for eo_conn_gateway
-- ----------------------------
DROP TABLE IF EXISTS `eo_conn_gateway`;
CREATE TABLE `eo_conn_gateway` (
  `connID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `userID` int(10) unsigned NOT NULL,
  `gatewayID` int(10) unsigned NOT NULL,
  PRIMARY KEY (`connID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for eo_gateway
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway`;
CREATE TABLE `eo_gateway` (
  `gatewayID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gatewayName` varchar(255) NOT NULL,
  `gatewayDesc` varchar(200) DEFAULT NULL,
  `gatewayStatus` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `gatewayAlias` varchar(255) NOT NULL,
  `productType` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `endDate` datetime DEFAULT NULL,
  `hashKey` varchar(255) NOT NULL,
  `updateTime` datetime NOT NULL,
  `createTime` datetime NOT NULL,
  PRIMARY KEY (`gatewayID`,`hashKey`,`gatewayAlias`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for eo_gateway_api
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_api`;
CREATE TABLE `eo_gateway_api` (
  `apiID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `apiName` varchar(255) NOT NULL,
  `groupID` int(10) unsigned NOT NULL,
  `gatewayProtocol` tinyint(3) unsigned NOT NULL,
  `gatewayRequestType` tinyint(3) unsigned NOT NULL,
  `gatewayRequestURI` varchar(255) NOT NULL,
  `backendProtocol` tinyint(255) NOT NULL,
  `backendRequestType` tinyint(255) NOT NULL,
  `backendID` int(11) unsigned NOT NULL,
  `backendRequestURI` varchar(255) NOT NULL,
  `isRequestBody` tinyint(3) unsigned NOT NULL,
  `gatewayRequestBodyNote` varchar(255) DEFAULT NULL,
  `gatewayID` int(10) unsigned NOT NULL,
  PRIMARY KEY (`apiID`,`gatewayRequestURI`,`apiName`,`gatewayID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for eo_gateway_api_cache
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_api_cache`;
CREATE TABLE `eo_gateway_api_cache` (
  `cacheID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gatewayID` int(11) unsigned NOT NULL,
  `groupID` int(11) NOT NULL,
  `apiID` int(11) NOT NULL,
  `path` varchar(255) NOT NULL,
  `apiJson` longtext NOT NULL,
  `gatewayHashKey` varchar(255) NOT NULL,
  `redisJson` longtext NOT NULL,
  `backendID` int(11) NOT NULL,
  PRIMARY KEY (`cacheID`,`path`,`gatewayID`,`gatewayHashKey`,`apiID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for eo_gateway_api_constant
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_api_constant`;
CREATE TABLE `eo_gateway_api_constant` (
  `paramID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `paramKey` varchar(255) NOT NULL,
  `paramValue` varchar(255) NOT NULL,
  `paramName` varchar(255) DEFAULT NULL,
  `backendParamPosition` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `apiID` int(10) unsigned NOT NULL,
  PRIMARY KEY (`paramID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for eo_gateway_api_group
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_api_group`;
CREATE TABLE `eo_gateway_api_group` (
  `groupID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `groupName` varchar(255) NOT NULL,
  `gatewayID` int(10) unsigned NOT NULL,
  `parentGroupID` int(10) unsigned NOT NULL DEFAULT '0',
  `isChild` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`groupID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for eo_gateway_api_request_param
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_api_request_param`;
CREATE TABLE `eo_gateway_api_request_param` (
  `paramID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gatewayParamPostion` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `isNotNull` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `paramType` tinyint(3) unsigned NOT NULL,
  `gatewayParamKey` varchar(255) NOT NULL,
  `backendParamPosition` tinyint(3) unsigned NOT NULL,
  `backendParamKey` varchar(255) NOT NULL,
  `apiID` int(10) unsigned NOT NULL,
  PRIMARY KEY (`paramID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for eo_gateway_area
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_area`;
CREATE TABLE `eo_gateway_area` (
  `areaID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `areaName` varchar(255) DEFAULT NULL,
  `areaStatus` tinyint(3) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`areaID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for eo_gateway_auth
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_auth`;
CREATE TABLE `eo_gateway_auth` (
  `authID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `authType` tinyint(4) NOT NULL,
  `apiKey` varchar(255) DEFAULT NULL,
  `userName` varchar(255) DEFAULT NULL,
  `userPassword` varchar(255) DEFAULT NULL,
  `strategyID` int(11) NOT NULL,
  PRIMARY KEY (`authID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for eo_gateway_backend
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_backend`;
CREATE TABLE `eo_gateway_backend` (
  `backendID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `backendName` varchar(255) NOT NULL,
  `backendURI` varchar(255) NOT NULL,
  `gatewayID` int(10) unsigned NOT NULL,
  PRIMARY KEY (`backendID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for eo_gateway_rate_limit
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_rate_limit`;
CREATE TABLE `eo_gateway_rate_limit` (
  `limitID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `intervalType` tinyint(4) DEFAULT NULL,
  `viewType` tinyint(4) NOT NULL,
  `limitCount` int(11) DEFAULT NULL,
  `priorityLevel` int(4) unsigned NOT NULL DEFAULT '1',
  `strategyID` int(11) NOT NULL,
  `createTime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updateTime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `startTime` varchar(255) NOT NULL,
  `endTime` varchar(255) NOT NULL,
  PRIMARY KEY (`limitID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for eo_gateway_strategy_group
-- ----------------------------
DROP TABLE IF EXISTS `eo_gateway_strategy_group`;
CREATE TABLE `eo_gateway_strategy_group` (
  `strategyID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `strategyName` varchar(255) NOT NULL,
  `strategyKey` varchar(255) NOT NULL,
  `gatewayID` int(11) NOT NULL,
  `strategyDesc` varchar(255) DEFAULT NULL,
  `updateTime` timestamp NULL DEFAULT NULL,
  `createTime` timestamp NULL DEFAULT NULL,
  `blackList` text,
  `whiteList` text,
  `chooseType` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`strategyID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

