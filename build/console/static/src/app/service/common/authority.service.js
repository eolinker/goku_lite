(function () {
    /**
     * @description 权限设置
     * @author 广州银云信息科技有限公司
     */
    /**
     * @version 1.4 新增权限路由
     */
    'use strict';
    angular.module('eolinker')
        .factory('Authority_CommonService', Authority_CommonService);

    Authority_CommonService.$inject = []

    function Authority_CommonService() {
        var data = {
                permission: {
                    default: null
                }
            },
            fun = {};
        fun.clear = function (mark) {
            data.permission[mark] = null;
        }
        fun.setPermission = function (mark, arg) {
            switch (mark) {
                case 'default':
                    {
                        switch (arg.userType) {
                            case 0:
                                {
                                    data.permission.default = {
                                        "authorityManagement":{
                                            "edit":true,
                                            "look":true
                                        },
                                        "teammateManagement":{
                                            "edit":true,
                                            "editManager":true
                                        },
                                        "apiManagement": {
                                            "edit": true
                                        },
                                        "loadBalance": {
                                            "edit": true
                                        },
                                        "strategyManagement": {
                                            "edit": true
                                        },
                                        "nodeManagement": {
                                            "edit": true
                                        },
                                        "pluginManagement": {
                                            "edit": true
                                        },
                                        "gatewayConfig": {
                                            "edit": true
                                        }
                                    }
                                    break;
                                }
                            case 1:
                                {
                                    data.permission.default = {
                                        "authorityManagement":{
                                            "edit":true,
                                            "look":true
                                        },
                                        "teammateManagement":{
                                            "edit":true,
                                            "editManager":false
                                        },
                                        "apiManagement": {
                                            "edit": true
                                        },
                                        "loadBalance": {
                                            "edit": true
                                        },
                                        "strategyManagement": {
                                            "edit": true
                                        },
                                        "nodeManagement": {
                                            "edit": true
                                        },
                                        "pluginManagement": {
                                            "edit": true
                                        },
                                        "gatewayConfig": {
                                            "edit": true
                                        }
                                    }
                                    break;
                                }
                            case 2:
                                {
                                    data.permission.default = arg.permission;
                                    break;
                                }
                        }
                        break;
                    }
            }
        }
        return {
            permission: data.permission,
            fun: fun
        };
    }
})();