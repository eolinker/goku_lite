(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 交互拦截相关服务js
     */
    angular.module('eolinker')
        .factory('AuthInterceptor', AuthInterceptor);

    AuthInterceptor.$inject = ['$rootScope', 'FILTER_WARNING_CODE_ARR_COMMON_CONST', '$q', 'AUTH_EVENTS', 'ERR_CODE_ARR_COMMON_CONST', '$filter', 'CODE']

    function AuthInterceptor($rootScope, FILTER_WARNING_CODE_ARR_COMMON_CONST, $q, AUTH_EVENTS, ERR_CODE_ARR_COMMON_CONST, $filter, CODE) {
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
        data.fun.request = function (config) { //交互请求
            return config;
        };
        data.fun.response = function (response) { //交互响应
            if (response.data) {
                $rootScope.$broadcast({
                    901: AUTH_EVENTS.UNAUTHENTICATED,
                    401: AUTH_EVENTS.UNAUTHORIZED
                } [response.data.code], response);
                try {
                    if (!/^\/nodeHttpServer/.test(response.config.url) && (typeof response.data == 'object')) {
                        response.data = JSON.parse($filter('HtmlFilter')(angular.toJson(response.data)));
                        switch (response.data.statusCode) {
                            case CODE.COMMON.UNLOGIN: {
                                if (/(home)|(transaction)|(guide)/.test(window.location.href)) {
                                    $rootScope.$broadcast('$LoginAgain_Core', response.config);
                                    return $q.reject('Sorry, the current status is not logged in. If you want to continue the operation, please login first!');
                                }
                                break;
                            }
                            case CODE.COMMON.SUCCESS: {
                                break;
                            }
                            default: {
                                if (FILTER_WARNING_CODE_ARR_COMMON_CONST.indexOf(response.data.statusCode) === -1) {
                                    var tmpErrCodeString = ERR_CODE_ARR_COMMON_CONST[response.data.statusCode];
                                    if (tmpErrCodeString) {
                                        $rootScope.InfoModal(tmpErrCodeString, 'error');
                                    } else if (response.data.statusCode) {
                                        $rootScope.InfoModal(response.data.resultDesc || '操作失败，请稍候再试！', 'error');
                                    }
                                    response.data.statusCode = CODE.COMMON.HAD_WARNING;
                                }
                                break;
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
        data.fun.responseError = function (rejection) { //交互响应出错
            $rootScope.$broadcast(AUTH_EVENTS.SYSTEM_ERROR);
            return $q.reject(rejection);
        };
        return data.fun;
    }
})();