(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 定义数据字典resource服务定义js
     */
    angular.module('eolinker.resource')

        .factory('GatewayResource', GatewayResource)

    GatewayResource.$inject = ['$resource', 'serverUrl'];

    function GatewayResource($resource, serverUrl) {
        var data = {
            info: {
                api: [],
                method: 'POST'
            }
        };
        data.info.api['ConfigLog'] = $resource(serverUrl+'config/log/:operate', {

        }, {
            Console: {
                params: {
                    operate: 'console'
                },
                method: 'GET',
                cancellable: true
            },
            SetConsole: {
                params: {
                    operate: 'console'
                },
                method: 'PUT',
                cancellable: true
            },
            Node: {
                params: {
                    operate: 'node'
                },
                method: 'GET',
                cancellable: true
            },
            SetNode: {
                params: {
                    operate: 'node'
                },
                method: 'PUT',
                cancellable: true
            },
            Access: {
                params: {
                    operate: 'access'
                },
                method: 'GET',
                cancellable: true
            },
            SetAccess: {
                params: {
                    operate: 'access'
                },
                method: 'PUT',
                cancellable: true
            }
        });
        data.info.api['ServiceDiscovery'] = $resource(serverUrl+'balance/service/:operate', {

        }, {
            DriverQuery: {
                params: {
                    operate: 'drivers'
                },
                method: 'GET',
                cancellable: true
            },Query: {
                params: {
                    operate: 'list'
                },
                method: 'GET',
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'info'
                },
                method: 'GET',
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'save'
                },
                method: data.info.method,
                cancellable: true
            },
            SimpleQuery: {
                params: {
                    operate: 'simple'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['Cluster'] = $resource(serverUrl +'cluster/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'list'
                },
                method: 'GET',
                cancellable: true
            },
            SimpleQuery: {
                params: {
                    operate: 'simpleList'
                },
                method: 'GET',
                cancellable: true
            }
        });
        data.info.api['AlertMessage'] = $resource(serverUrl + 'alert/msg/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true,
            },
            Empty: {
                params: {
                    operate: 'clear'
                },
                method: data.info.method,
                cancellable: true,
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.info.method,
                cancellable: true,
            }
        });
        data.info.api['Monitor'] = $resource(serverUrl + 'monitor/gateway/:operate', {

        }, {
            Info: {
                params: {
                    operate: 'getSummaryInfo'
                },
                method: data.info.method,
                cancellable: true,
            },
            Refresh: {
                params: {
                    operate: 'refreshInfo'
                },
                method: data.info.method,
                cancellable: true,
            },
            Download: {
                params: {
                    operate: 'download'
                },
                method: data.info.method,
                cancellable: true,
            }
        });
        data.info.api['Config'] = $resource(serverUrl + 'gateway/config/:mark/:operate', {

        }, {
            BaseInfo: {
                params: {
                    mark: 'base',
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true,
            },
            AlertInfo: {
                params: {
                    mark: 'alert',
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true,
            },
            BaseEdit: {
                params: {
                    mark: 'base',
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true,
            },
            AlertEdit: {
                params: {
                    mark: 'alert',
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true,
            }
        });
        data.info.api['ImportAms'] = $resource(serverUrl + 'import/ams/:operate', {

        }, {
            Project: {
                params: {
                    operate: 'project'
                },
                method: data.info.method,
                cancellable: true,
                transformRequest: angular.identity,
                headers: {
                    "Content-Type": undefined
                }
            },
            Api: {
                params: {
                    operate: 'api'
                },
                method: data.info.method,
                cancellable: true,
                transformRequest: angular.identity,
                headers: {
                    "Content-Type": undefined
                }
            },
            Group: {
                params: {
                    operate: 'group'
                },
                method: data.info.method,
                cancellable: true,
                transformRequest: angular.identity,
                headers: {
                    "Content-Type": undefined
                }
            }
        });
        data.info.api['Project'] = $resource(serverUrl + 'project/:operate/:target', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: "GET",
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            QueryAndGroup: {
                params: {
                    operate: 'strategy',
                    target:"getList"
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['Api'] = $resource(serverUrl + 'apis/:mark/:operate', {

        }, {
            IDQuery: {
                params: {
                    mark:"id",
                    operate: 'getList'
                },
                method: "GET",
                cancellable: true
            },
            Query: {
                params: {
                    operate: 'getList'
                },
                method: "GET",
                cancellable: true
            },
            Copy: {
                params: {
                    operate: 'copy'
                },
                method: data.info.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.info.method,
                cancellable: true
            },
            ChangeGroup: {
                params: {
                    operate: 'batchEditGroup'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['ApiGroup'] = $resource(serverUrl + 'apis/group/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['Auth'] = $resource(serverUrl + 'auth/:operate', {

        }, {
            Edit: {
                params: {
                    operate: 'editInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            Status: {
                params: {
                    operate: 'getStatus'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['Node'] = $resource(serverUrl + 'node/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            ChangeGroup: {
                params: {
                    operate: 'batchEditGroup'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            Reload: {
                params: {
                    operate: 'batchReload'
                },
                method: data.info.method,
                cancellable: true
            },
            Restart: {
                params: {
                    operate: 'batchRestart'
                },
                method: data.info.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'batchStart'
                },
                method: data.info.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['NodeGroup'] = $resource(serverUrl + 'node/group/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['Strategy'] = $resource(serverUrl + 'strategy/:mark/:operate', {

        }, {
            IDQuery: {
                params: {
                    mark:"id",
                    operate: 'getList'
                },
                method: "GET",
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: "GET",
                cancellable: true
            },
            Query: {
                params: {
                    operate: 'getList'
                },
                method: "GET",
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Copy: {
                params: {
                    operate: 'copy'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.info.method,
                cancellable: true
            },
            ChangeGroup: {
                params: {
                    operate: 'batchEditGroup'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'batchStart'
                },
                method: data.info.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['StrategyGroup'] = $resource(serverUrl + 'strategy/group/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['StrategyApi'] = $resource(serverUrl + 'strategy/api/:mark/:operate', {

        }, {
            IDQuery: {
                params: {
                    mark:"id",
                    operate: 'getList'
                },
                method: "GET",
                cancellable: true
            },
            UnassignIDQuery: {
                params: {
                    mark:"id",
                    operate: 'getNotInList'
                },
                method: "GET",
                cancellable: true
            },
            All: {
                params: {
                    operate: 'getNotInList'
                },
                method: data.info.method,
                cancellable: true
            },
            Query: {
                params: {
                    operate: 'getList'
                },
                method: "GET",
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            ChangeTarget: {
                params: {
                    operate: 'batchEditTarget'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['Balance'] = $resource(serverUrl + 'balance/:operate', {

        }, {
            CheckIsExistInCluster: {
                params: {
                    operate: 'exits'
                },
                method: data.info.method,
                cancellable: true
            },
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            SimpleQuery:{
                params: {
                    operate: 'simple'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['Plugin'] = $resource(serverUrl + 'plugin/:mark/:operate', {

        }, {
            TagQuery: {
                params: {
                    operate: 'getListByType'
                },
                method: data.info.method,
                cancellable: true
            },
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'start'
                },
                method: data.info.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'stop'
                },
                method: data.info.method,
                cancellable: true
            },
            BatchStop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.info.method,
                cancellable: true
            },
            BatchStart: {
                params: {
                    operate: 'batchStart'
                },
                method: data.info.method,
                cancellable: true
            },
            Check: {
                params: {
                    operate: 'check',
                    mark:'availiable'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['PluginStrategy'] = $resource(serverUrl + 'plugin/strategy/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'addPluginToStrategy'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'batchStart'
                },
                method: data.info.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['GpeditPluginApi'] = $resource(serverUrl + 'strategy/api/plugin/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        data.info.api['PluginApi'] = $resource(serverUrl + 'plugin/api/:mark/:operate', {

        }, {
            QueryByStrategy: {
                params: {
                    mark:'notAssign',
                    operate: 'getList'
                },
                method: data.info.method,
                cancellable: true
            },
            Query: {
                params: {
                    operate: 'getListByStrategy'
                },
                method: data.info.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'addPluginToApi'
                },
                method: data.info.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.info.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.info.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.info.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'batchStart'
                },
                method: data.info.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.info.method,
                cancellable: true
            }
        });
        return data.info.api;
    }
})();