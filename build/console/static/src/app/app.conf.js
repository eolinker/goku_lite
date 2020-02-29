(function () { 
 return angular.module("eolinker")
.constant("serverUrl", "../")
.constant("nodeServerUrl", "../nodeHttpServer/")
.constant("isDebug", false)
.constant("assetUrl", "")
.constant("PRODUCT_TYPE", "online")
.constant("COOKIE_CONFIG", {"path":"/","domain":".eolinker.com"})
.constant("WEBSOCKET_PORT", 1204);

})();
