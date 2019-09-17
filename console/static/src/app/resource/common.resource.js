(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 定义公用接口resource服务定义js
     */
    angular.module('eolinker.resource')

        .factory('CommonResource', CommonResource)

    CommonResource.$inject = ['$resource', 'serverUrl'];

    function CommonResource($resource, serverUrl) {
        var data = {
            info: {
                api: [],
                method: 'POST'
            }
        }
        data.info.api['Guest'] = $resource(serverUrl + 'guest/:operate', {

        }, {
            Check: {
                params: {
                    operate: 'checkLogin'
                },
                method: data.info.method,
                cancellable: true
            },
            Login: {
                params: {
                    operate: 'login'
                },
                method: data.info.method,
                cancellable: true
            }
        });


        data.info.api['User'] = $resource(serverUrl + 'user/:mark/:operate', {

        }, {
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            LoginOut: {
                params: {
                    operate: 'logout'
                },
                method: data.info.method,
                cancellable: true
            },
            ChangePassword: {
                params: {
                    mark: 'password',
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            }
        });

        return data.info.api;
    }
})();