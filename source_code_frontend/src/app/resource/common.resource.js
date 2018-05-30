(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 定义公用接口resource服务定义js
     */
    angular.module('goku.resource')

    .factory('CommonResource', CommonResource)

    CommonResource.$inject = ['$resource', 'serverUrl'];

    function CommonResource($resource, serverUrl) {
        var data = {
            info: {
                api: [],
                method: 'POST'
            }
        }
        data.info.api['Install'] = $resource(serverUrl + 'Web/Install/:operate', {

            }, {
                CheckConfig: {
                    params: { operate: 'checkIsInstall' },
                    method: data.info.method,
                    cancellable: true
                },
                CheckDatabase: {
                    params: { operate: 'checkDBConnect' },
                    method: data.info.method,
                    cancellable: true
                },
                CheckRedis: {
                    params: { operate: 'checkRedisConnect' },
                    method: data.info.method,
                    cancellable: true
                },
                Post: {
                    params: { operate: 'install' },
                    method: data.info.method,
                    cancellable: true
                },
                PostConfig: {
                    params: { operate: 'installConfigure' },
                    method: data.info.method,
                    cancellable: true
                }
                
            }

        );
        data.info.api['Guest'] = $resource(serverUrl + 'Web/Guest/:operate', {

        }, {
            Login: {
                params: { operate: 'login' },
                method: data.info.method,
                cancellable: true
            }
        });

        data.info.api['User'] = $resource(serverUrl + 'Web/User/:operate', {

        }, {
            Check: {
                params: { operate: 'checkLogin' },
                method: data.info.method,
                cancellable: true
            },
            LoginOut: {
                params: { operate: 'logout' },
                method: data.info.method,
                cancellable: true
            },
            Password: {
                params: { operate: 'editPassword' },
                method: data.info.method,
                cancellable: true
            }
        });


        return data.info.api;
    }
})();
