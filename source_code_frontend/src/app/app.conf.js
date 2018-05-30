(function () { 
 return angular.module("goku")
.constant("serverUrl", "../")
.constant("isDebug", false)
.constant("assetUrl", "")
.constant("COOKIE_CONFIG", {"path":"/","domain":".goku.com"})
.constant("WEBSOCKET_PORT", 1204);

})();
