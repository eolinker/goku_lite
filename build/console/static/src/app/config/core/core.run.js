(function () {
    'use strict';
    //author：广州银云信息科技有限公司
    angular
        .module('eolinker')
        .run(appRun);

    appRun.$inject = ['$rootScope', '$state', 'AUTH_EVENTS'];

    function appRun($rootScope, $state, AUTH_EVENTS) {
        var data = {
            info: {
                title: {
                    root: 'GoKu Gateway | 企业微服务架构的首选解决方案，加速企业数字化转型'
                }
            },
            modal: {
                loginIsExist: null
            },
            fun: {
                cancelRequest: null
            }
        };
        $rootScope.global = {
            ajax: {},
            $watch: []
        }
        var fun = {};
        /**
         * @description 销毁$watch队列
         */
        fun.destoryWatch = function () {
            $rootScope.global.$watch.map(function (val, key) {
                val();
            })
            $rootScope.global.$watch = [];
        }
        /**
         * 跳转页面时取消当前页面ajax请求
         */
        data.fun.cancelRequest = function () {
            for (var key in $rootScope.global.ajax) {
                var val = $rootScope.global.ajax[key];
                if (!val) return;
                val.$cancelRequest();
            }
            $rootScope.global.ajax = {};
        }
        /**
         * @description 监听广播未登录重新登录弹窗
         * @params {any} default 方法默认参数
         * @param {object} arg 传值
         */
        $rootScope.$on('$LoginAgain_Core', function (_default, arg) {
            if (data.modal.loginIsExist) return;
            data.modal.loginIsExist = true;
            $rootScope.Common_LoginModal(function (callback) {
                if (callback) {
                    $rootScope.InfoModal('登录成功！', 'success', function () {
                        if (callback.loginCall == window.localStorage['LOGINCALL']) {
                            if (!/(edit)|(add)/i.test(arg.url)) {
                                window.location.reload();
                            }
                        } else {
                            window.localStorage.setItem('LOGINCALL', callback.loginCall);
                            if (/home\/panel$/.test(window.location.href)) {
                                window.location.reload();
                            } else {
                                window.location.href = "./#/home/panel";
                            }
                        }
                    });
                } else {
                    $state.go('index');
                }
                data.modal.loginIsExist = false;
            });
        });
        $rootScope.$on('$stateChangeStart', function (_default, _target,_params,_origin) {
            if (_target.containerRouter) {
                location.href = "./#/";
            }
            if(_target.name!==_origin.name){
                window.sessionStorage.removeItem('COMMON_SEARCH_TIP');
            }
            data.fun.cancelRequest();
        });
        $rootScope.$on('$stateChangeSuccess', function (_default, arg) {
            window.scrollTo(0, 0);
            fun.destoryWatch();
        });

        $rootScope.$on('$TransferStation', function (_default, arg) { //转换交互
            $rootScope.$broadcast(arg.state, arg.data);
        });
        $rootScope.$on('$WindowTitleSet', function (_default, arg) { //设置title
            arg = arg || {
                list: []
            };
            if (arg.list.length > 0) {
                window.document.title = arg.list.join('-') + (arg.list.length >= 1 ? '-' : '') + data.info.title.root;
            } else {
                window.document.title = data.info.title.root;
            }
        });
        $rootScope.$on(AUTH_EVENTS.SYSTEM_ERROR, function (_default) {
            console.log("error");
        })
    }

})();