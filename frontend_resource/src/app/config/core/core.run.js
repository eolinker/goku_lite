(function() {
    'use strict';
    //author：广州银云信息科技有限公司
    angular
        .module('goku')
        .run(appRun);

    appRun.$inject = ['$rootScope', '$state', 'CommonResource', 'AUTH_EVENTS', 'USER_ROLES', 'CODE', '$location'];

    function appRun($rootScope, $state, CommonResource, AUTH_EVENTS, USER_ROLES, CODE, $location) {
        var data = {
            info: {
                title: {
                    root: 'GoKu Gateway（Lite）| 国内首个轻量级Go语言开源微服务API网关 '
                },
                _hmt: []
            },
            modal: {
                loginIsExist: null
            },
            ajax: {
                checkLogin: null
            },
            fun: {
                cancelRequest: null
            }
        };
        $rootScope.global = {
            ajax: {}
        }
        CommonResource.Install.CheckConfig().$promise.then(function(response) {
            switch (response.statusCode) {
                case '200000':
                    {
                        $state.go('arrange');
                        break;
                    }
            }
        })
        /**
         * 跳转页面时取消当前页面ajax请求
         */
        data.fun.cancelRequest = function() {
            for (var key in $rootScope.global.ajax) {
                var val = $rootScope.global.ajax[key];
                if (!val) return;
                val.$cancelRequest();
            }
            $rootScope.global.ajax = {};
        }
        $rootScope.$on('$stateChangeStart', function(_default, arg) {
            window.scrollTo(0, 0);
            data.fun.cancelRequest();
        });

        /**
         * @description 监听广播未登录重新登录弹窗
         * @params {any} default 方法默认参数
         * @param {object} arg 传值
         */
        $rootScope.$on('$LoginAgain_Core', function(_default, arg) {
            if (data.modal.loginIsExist) return;
            data.modal.loginIsExist = true;
            $rootScope.Common_LoginModal(function(callback) {
                if (callback) {
                    $rootScope.InfoModal('登录成功！', 'success', function() {
                        if (callback.loginCall == window.localStorage['LOGINCALL']) {
                            if (!/(edit)|(add)/i.test(arg.url)) {
                                window.location.reload();
                            }
                        } else {
                            window.localStorage.removeItem('VERSIONINFO');
                            window.localStorage.setItem('LOGINCALL', callback.loginCall);
                            if (/home\/$/.test(window.location.href)) {
                                window.location.reload();
                            } else {
                                window.location.href = "./#/home/";
                            }
                        }
                    });
                } else {
                    $state.go('index');
                }
                data.modal.loginIsExist = false;
            });
        });
        $rootScope.$on('$TransferStation', function(_default, arg) { //转换交互
            $rootScope.$broadcast(arg.state, arg.data);
        });
        $rootScope.$on('$WindowTitleSet', function(_default, arg) { //设置title
            arg = arg || { list: [] };
            if (arg.list.length > 0) {
                window.document.title = arg.list.join('-') + (arg.list.length >= 1 ? '-' : '') + data.info.title.root;
            } else {
                window.document.title = data.info.title.root;
            }
        });
        $rootScope.$on(AUTH_EVENTS.SYSTEM_ERROR, function(_default) {
            console.log("error");
        })
        $rootScope.$on(AUTH_EVENTS.UNAUTHENTICATED, function(_default) {
            $state.go('login');
        })

        $rootScope.$on(AUTH_EVENTS.UNAUTHORIZED, function(_default) {
            $state.go('login');
        })
    }

})();