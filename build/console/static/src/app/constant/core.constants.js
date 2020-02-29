(function() {
    'use strict';
    /*
     * aurhor:广州银云信息科技有限公司
     * 用户权限常量集
     */
    angular
        .module('eolinker.constant')
        .constant('AUTH_EVENTS', {
            //--登录成功--
            LOGIN_SUCCESS: 'auth-login-success',
            //--登录失败--
            LOGIN_FAILED: 'auth-login-failed',
            //--退出成功--
            LOGOUT_SUCCESS: 'auth-logout-success',
            //--认证超时--
            SESSION_TIMEOUT: 'auth-session-timeout',
            //--未认证权限--
            UNAUTHENTICATED: 'auth-not-authenticated',
            //--未登录--
            UNAUTHORIZED: 'auth-not-authorized',
            //--服务器出错--
            SYSTEM_ERROR: 'something-wrong-system'
        })
        .constant('USER_ROLES', {
            USER: 'guest'
        })
        .constant('PATH_INFO', {
            HOSTNAME:'nj.goku.com:1204',
            HOST:window.location.origin+'/',
            INHERIT_HOST:window.location.origin+window.location.pathname,
            MOCK: {
                DEFAULT: 'http://result.goku.com/',
                HIGH: 'http://mock.goku.com/',
            }

        })
})();
