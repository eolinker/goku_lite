(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 交互拦截相关服务js
     */
    angular.module('goku')
        .factory('AuthInterceptor', AuthInterceptor);

    AuthInterceptor.$inject = ['$rootScope', '$q', 'AUTH_EVENTS', '$cookies', '$filter', 'CODE']

    function AuthInterceptor($rootScope, $q, AUTH_EVENTS, $cookies, $filter, CODE) {
        var Auth;
        var data = {
            info: {
                auth: null
            },
            fun: {
                request: null, //交互请求功能函数
                response: null, //交互响应功能函数
                responseError: null //交互响应出错功能函数
            }
        }
        data.fun.request = function(config) { //交互请求
            config.headers = config.headers || {};
            return config;
        };
        data.fun.response = function(response) { //交互响应
            if (response.data) {
                $rootScope.$broadcast({
                    901: AUTH_EVENTS.UNAUTHENTICATED,
                    401: AUTH_EVENTS.UNAUTHORIZED
                }[response.data.code], response);
                try {
                    if (typeof response.data == 'object') {
                        response.data = JSON.parse($filter('HtmlFilter')(angular.toJson(response.data)));
                        if (response.data.statusCode == CODE.COMMON.UNLOGIN) {
                            if (/(home)/.test(window.location.href)) {
                                $rootScope.$broadcast('$LoginAgain_Core', response.config);
                                return $q.reject('Sorry, the current status is not logged in. If you want to continue the operation, please login first!');
                            }
                        }
                    }
                } catch (e) {
                    response.data = response.data;
                    $rootScope.$broadcast(AUTH_EVENTS.SYSTEM_ERROR);
                }
            }
            return $q.resolve(response);
        };
        data.fun.responseError = function(rejection) { //交互响应出错
            $rootScope.$broadcast(AUTH_EVENTS.SYSTEM_ERROR);
            return $q.reject(rejection);
        };
        return data.fun;
    }
})();