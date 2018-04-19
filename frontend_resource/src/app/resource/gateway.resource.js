(function() {
    'use strict';
    /*
     * author：riverLethe
     * 定义接口网关resource服务定义js
     */
    angular.module('goku.resource')

    .factory('GatewayResource', GatewayResource)

    GatewayResource.$inject = ['$resource', 'serverUrl'];

    function GatewayResource($resource, serverUrl) {
        var data = {
            info: {
                api: [],
                method: 'POST'
            }
        }
        data.info.api['Gateway'] = $resource(serverUrl + 'Web/Gateway/:operate', {

            }, {
                Add: {
                    params: { operate: 'addGateway' },
                    method: data.info.method
                },
                Query: {
                    params: { operate: 'getGatewayList' },
                    method: data.info.method
                },
                Delete: {
                    params: { operate: 'deleteGateway' },
                    method: data.info.method
                },
                Edit: {
                    params: { operate: 'editGateway' },
                    method: data.info.method
                },
                Info: {
                    params: { operate: 'getGateway' },
                    method: data.info.method
                }
            }

        );

        data.info.api['ApiGroup'] = $resource(serverUrl + 'Web/Group/:operate', {

        }, {
            Add: {
                params: { operate: 'addGroup' },
                method: data.info.method
            },
            Query: {
                params: { operate: 'getGroupList' },
                method: data.info.method
            },
            Delete: {
                params: { operate: 'deleteGroup' },
                method: data.info.method
            },
            Edit: {
                params: { operate: 'editGroup' },
                method: data.info.method
            }
        });

        data.info.api['Api'] = $resource(serverUrl + 'Web/Api/:operate', {

        }, {
            Add: {
                params: { operate: 'addApi' },
                method: data.info.method
            },
            All: {
                params: { operate: 'getAllApiList' },
                method: data.info.method
            },
            Query: {
                params: { operate: 'getApiList' },
                method: data.info.method
            },
            Delete: {
                params: { operate: 'deleteApi' },
                method: data.info.method
            },
            Edit: {
                params: { operate: 'editApi' },
                method: data.info.method
            },
            Search: {
                params: { operate: 'searchApi' },
                method: data.info.method
            },
            Detail: {
                params: { operate: 'getApi' },
                method: data.info.method
            }
        });

        data.info.api['Backend'] = $resource(serverUrl + 'Web/Backend/:operate', {

        }, {
            Add: {
                params: { operate: 'addBackend' },
                method: data.info.method
            },
            Query: {
                params: { operate: 'getBackendList' },
                method: data.info.method
            },
            Delete: {
                params: { operate: 'deleteBackend' },
                method: data.info.method
            },
            Edit: {
                params: { operate: 'editBackend' },
                method: data.info.method
            }
        });
        data.info.api['Frequency'] = $resource(serverUrl + 'Web/Frequency/:operate', {

        }, {
            Add: {
                params: { operate: 'addFrequency' },
                method: data.info.method
            },
            Query: {
                params: { operate: 'getFrequencyList' },
                method: data.info.method
            },
            Delete: {
                params: { operate: 'deleteFrequency' },
                method: data.info.method
            },
            Edit: {
                params: { operate: 'editFrequency' },
                method: data.info.method
            }
        });
        data.info.api['Ip'] = $resource(serverUrl + 'Web/IP/:operate', {

        }, {
            Info: {
                params: { operate: 'getIPInfo' },
                method: data.info.method
            },
            Edit: {
                params: { operate: 'editIPList' },
                method: data.info.method
            }
        });
        return data.info.api;
    }
})();
