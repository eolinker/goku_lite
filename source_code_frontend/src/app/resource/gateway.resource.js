(function () {
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
        data.info.api['Config'] = $resource(serverUrl + 'Web/Global/:operate', {

        }, {
            Info: {
                params: {
                    operate: 'getConfInfo'
                },
                method: data.info.method
            },
            Edit: {
                params: {
                    operate: 'editConfInfo'
                },
                method: data.info.method
            }
        });
        data.info.api['Service'] = $resource(serverUrl + 'Web/GatewayService/:operate', {

        }, {
            Stop: {
                params: {
                    operate: 'stop'
                },
                method: data.info.method
            },
            Restart: {
                params: {
                    operate: 'restart'
                },
                method: data.info.method
            },
            Reload: {
                params: {
                    operate: 'reload'
                },
                method: data.info.method
            },
            Start: {
                params: {
                    operate: 'start'
                },
                method: data.info.method
            }
        });
        data.info.api['Auth'] = $resource(serverUrl + 'Web/Auth/:operate', {

        }, {
            Info: {
                params: {
                    operate: 'getAuthInfo'
                },
                method: data.info.method
            },
            Edit: {
                params: {
                    operate: 'editAuth'
                },
                method: data.info.method
            }
        });
        data.info.api['Strategy'] = $resource(serverUrl + 'Web/Strategy/:operate', {

        }, {
            Add: {
                params: {
                    operate: 'addStrategy'
                },
                method: data.info.method
            },
            Query: {
                params: {
                    operate: 'getStrategyList'
                },
                method: data.info.method
            },
            SimpleQuery: {
                params: {
                    operate: 'getSimpleStrategyList'
                },
                method: data.info.method
            },
            Delete: {
                params: {
                    operate: 'deleteStrategy'
                },
                method: data.info.method
            },
            Edit: {
                params: {
                    operate: 'editStrategy'
                },
                method: data.info.method
            }
        });
        data.info.api['Gateway'] = $resource(serverUrl + 'Web/Gateway/:operate', {

            }, {
                Add: {
                    params: {
                        operate: 'addGateway'
                    },
                    method: data.info.method
                },
                Query: {
                    params: {
                        operate: 'getGatewayList'
                    },
                    method: data.info.method
                },
                Delete: {
                    params: {
                        operate: 'deleteGateway'
                    },
                    method: data.info.method
                },
                Edit: {
                    params: {
                        operate: 'editGateway'
                    },
                    method: data.info.method
                },
                Info: {
                    params: {
                        operate: 'getGateway'
                    },
                    method: data.info.method
                },
                CheckAlias: {
                    params: {
                        operate: 'checkGatewayAliasIsExist'
                    },
                    method: data.info.method,
                    cancellable: true
                }
            }

        );

        data.info.api['ApiGroup'] = $resource(serverUrl + 'Web/ApiGroup/:operate', {

        }, {
            Add: {
                params: {
                    operate: 'addGroup'
                },
                method: data.info.method
            },
            Query: {
                params: {
                    operate: 'getGroupList'
                },
                method: data.info.method
            },
            Delete: {
                params: {
                    operate: 'deleteGroup'
                },
                method: data.info.method
            },
            Edit: {
                params: {
                    operate: 'editGroup'
                },
                method: data.info.method
            }
        });

        data.info.api['Api'] = $resource(serverUrl + 'Web/Api/:operate', {

        }, {
            Add: {
                params: {
                    operate: 'addApi'
                },
                method: data.info.method
            },
            All: {
                params: {
                    operate: 'getAllApiList'
                },
                method: data.info.method
            },
            Query: {
                params: {
                    operate: 'getApiList'
                },
                method: data.info.method
            },
            Delete: {
                params: {
                    operate: 'deleteApi'
                },
                method: data.info.method
            },
            Edit: {
                params: {
                    operate: 'editApi'
                },
                method: data.info.method
            },
            Search: {
                params: {
                    operate: 'searchApi'
                },
                method: data.info.method
            },
            Detail: {
                params: {
                    operate: 'getApi'
                },
                method: data.info.method
            }
        });

        data.info.api['Backend'] = $resource(serverUrl + 'Web/Backend/:operate', {

        }, {
            Add: {
                params: {
                    operate: 'addBackend'
                },
                method: data.info.method
            },
            Query: {
                params: {
                    operate: 'getBackendList'
                },
                method: data.info.method
            },
            Delete: {
                params: {
                    operate: 'deleteBackend'
                },
                method: data.info.method
            },
            Edit: {
                params: {
                    operate: 'editBackend'
                },
                method: data.info.method
            }
        });
        data.info.api['RateLimit'] = $resource(serverUrl + 'Web/RateLimit/:operate', {

        }, {
            Add: {
                params: {
                    operate: 'addRateLimit'
                },
                method: data.info.method
            },
            Query: {
                params: {
                    operate: 'getRateLimitList'
                },
                method: data.info.method
            },
            Delete: {
                params: {
                    operate: 'deleteRateLimit'
                },
                method: data.info.method
            },
            Edit: {
                params: {
                    operate: 'editRateLimit'
                },
                method: data.info.method
            }
        });
        data.info.api['Ip'] = $resource(serverUrl + 'Web/IP/:operate', {

        }, {
            Info: {
                params: {
                    operate: 'getStrategyIPList'
                },
                method: data.info.method
            },
            Edit: {
                params: {
                    operate: 'editStrategyIPList'
                },
                method: data.info.method
            },
            GlobalInfo: {
                params: {
                    operate: 'getGatewayIPList'
                },
                method: data.info.method
            },
            GlobalEdit: {
                params: {
                    operate: 'editGatewayIPList'
                },
                method: data.info.method
            }
        });
        return data.info.api;
    }
})();