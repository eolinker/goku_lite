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
            api:[],
            method:"POST"
        };
        data.api['Version'] = $resource(serverUrl+'version/config/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: 'GET',
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.method,
                cancellable: true
            },
            Publish: {
                params: {
                    operate: 'publish'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['ConfigLog'] = $resource(serverUrl+'config/log/:operate', {

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
        data.api['ServiceDiscovery'] = $resource(serverUrl+'balance/service/:operate', {

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
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'save'
                },
                method: data.method,
                cancellable: true
            },
            SimpleQuery: {
                params: {
                    operate: 'simple'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['Cluster'] = $resource(serverUrl +'cluster/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'list'
                },
                method: 'GET',
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            SimpleQuery: {
                params: {
                    operate: 'simpleList'
                },
                method: 'GET',
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['MonitorModuleConf'] = $resource(serverUrl + 'monitor/module/config/:operate', {

        }, {
            Get: {
                params: {
                    operate: 'get'
                },
                method: "GET",
                cancellable: true,
            },
            Set: {
                params: {
                    operate: 'set'
                },
                method: data.method,
                cancellable: true,
            }
        });
        data.api['Monitor'] = $resource(serverUrl + 'monitor/gateway/:operate', {

        }, {
            Info: {
                params: {
                    operate: 'getSummaryInfo'
                },
                method: data.method,
                cancellable: true,
            },
            Refresh: {
                params: {
                    operate: 'refreshInfo'
                },
                method: data.method,
                cancellable: true,
            },
            Download: {
                params: {
                    operate: 'download'
                },
                method: data.method,
                cancellable: true,
            }
        });
        data.api['Config'] = $resource(serverUrl + 'gateway/config/:mark/:operate', {

        }, {
            BaseInfo: {
                params: {
                    mark: 'base',
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true,
            },
            BaseEdit: {
                params: {
                    mark: 'base',
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true,
            }
        });
        data.api['ImportAms'] = $resource(serverUrl + 'import/ams/:operate', {

        }, {
            Project: {
                params: {
                    operate: 'project'
                },
                method: data.method,
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
                method: data.method,
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
                method: data.method,
                cancellable: true,
                transformRequest: angular.identity,
                headers: {
                    "Content-Type": undefined
                }
            }
        });
        data.api['Project'] = $resource(serverUrl + 'project/:operate/:target', {

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
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            },
            QueryAndGroup: {
                params: {
                    operate: 'strategy',
                    target:"getList"
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['Api'] = $resource(serverUrl + 'apis/:mark/:operate', {

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
                method: data.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.method,
                cancellable: true
            },
            ChangeGroup: {
                params: {
                    operate: 'batchEditGroup'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['ApiGroup'] = $resource(serverUrl + 'apis/group/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['Auth'] = $resource(serverUrl + 'auth/:operate', {

        }, {
            Edit: {
                params: {
                    operate: 'editInfo'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            },
            Status: {
                params: {
                    operate: 'getStatus'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['Node'] = $resource(serverUrl + 'node/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            ChangeGroup: {
                params: {
                    operate: 'batchEditGroup'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            },
            Reload: {
                params: {
                    operate: 'batchReload'
                },
                method: data.method,
                cancellable: true
            },
            Restart: {
                params: {
                    operate: 'batchRestart'
                },
                method: data.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'batchStart'
                },
                method: data.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['NodeGroup'] = $resource(serverUrl + 'node/group/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['Strategy'] = $resource(serverUrl + 'strategy/:mark/:operate', {

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
                method: data.method,
                cancellable: true
            },
            Copy: {
                params: {
                    operate: 'copy'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.method,
                cancellable: true
            },
            ChangeGroup: {
                params: {
                    operate: 'batchEditGroup'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'batchStart'
                },
                method: data.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['StrategyGroup'] = $resource(serverUrl + 'strategy/group/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['StrategyApi'] = $resource(serverUrl + 'strategy/api/:mark/:operate', {

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
                method: data.method,
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
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            },
            ChangeTarget: {
                params: {
                    operate: 'batchEditTarget'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['Balance'] = $resource(serverUrl + 'balance/:operate', {

        }, {
            CheckIsExistInCluster: {
                params: {
                    operate: 'exits'
                },
                method: data.method,
                cancellable: true
            },
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            },
            SimpleQuery:{
                params: {
                    operate: 'simple'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['Plugin'] = $resource(serverUrl + 'plugin/:mark/:operate', {

        }, {
            TagQuery: {
                params: {
                    operate: 'getListByType'
                },
                method: data.method,
                cancellable: true
            },
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'add'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'delete'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'start'
                },
                method: data.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'stop'
                },
                method: data.method,
                cancellable: true
            },
            BatchStop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.method,
                cancellable: true
            },
            BatchStart: {
                params: {
                    operate: 'batchStart'
                },
                method: data.method,
                cancellable: true
            },
            Check: {
                params: {
                    operate: 'check',
                    mark:'availiable'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['PluginStrategy'] = $resource(serverUrl + 'plugin/strategy/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'addPluginToStrategy'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'batchStart'
                },
                method: data.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['GpeditPluginApi'] = $resource(serverUrl + 'strategy/api/plugin/:operate', {

        }, {
            Query: {
                params: {
                    operate: 'getList'
                },
                method: data.method,
                cancellable: true
            }
        });
        data.api['PluginApi'] = $resource(serverUrl + 'plugin/api/:mark/:operate', {

        }, {
            QueryByStrategy: {
                params: {
                    mark:'notAssign',
                    operate: 'getList'
                },
                method: data.method,
                cancellable: true
            },
            Query: {
                params: {
                    operate: 'getListByStrategy'
                },
                method: data.method,
                cancellable: true
            },
            Add: {
                params: {
                    operate: 'addPluginToApi'
                },
                method: data.method,
                cancellable: true
            },
            Edit: {
                params: {
                    operate: 'edit'
                },
                method: data.method,
                cancellable: true
            },
            Delete: {
                params: {
                    operate: 'batchDelete'
                },
                method: data.method,
                cancellable: true
            },
            Info: {
                params: {
                    operate: 'getInfo'
                },
                method: data.method,
                cancellable: true
            },
            Start: {
                params: {
                    operate: 'batchStart'
                },
                method: data.method,
                cancellable: true
            },
            Stop: {
                params: {
                    operate: 'batchStop'
                },
                method: data.method,
                cancellable: true
            }
        });
        return data.api;
    }
})();