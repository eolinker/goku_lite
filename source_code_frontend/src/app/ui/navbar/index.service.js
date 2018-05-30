(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     *顶部栏（navbar）相关服务js
     */
    angular.module('goku')
        .factory('NavbarService', NavbarService);

    NavbarService.$inject = ['$state', 'CommonResource', 'CODE']

    function NavbarService($state, CommonResource, CODE) {
        var data = {
            info: {
                status: 0, //登录状态 0：没登录，1：已登录
                navigation: {
                    query: [],
                    current: ''
                }
            },
            company: {
                query: null
            },
            fun: {
                logout: null, //退出登录
                init: null, //初始化功能函数
                $router: null //路由更换功能函数
            }
        }
        data.fun.logout = function() {
            var template = {
                promise: null
            }
            template.promise = CommonResource.User.LoginOut().$promise;
            template.promise.then(function(response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            if (/(home)/.test($state.current.name)) { //window.location.href
                                $state.go('index');
                            } else {
                                $state.reload();
                            }
                            data.info.status = 0;
                            break;
                        }
                }
            })
            return template.promise;
        }
        data.fun.$router = function() {
            var template = {
                promise: null
            }
            if (data.info.status == 0) {
                template.promise = CommonResource.User.Check().$promise;
                template.promise.then(function(response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                data.info.status = 1;
                                $state.go('home.gateway.default');
                                break;
                            }
                        default:
                            {
                                data.info.status = 0;
                            }
                    }
                });
            } else {
                $state.go('home.gateway.default');
            }

            return template.promise;
        }
        return data;
    }
})();