(function () {
    'use strict';
    /*
     * author：riverLethe
     * 策略组列表相关js
     */
    angular.module('goku')
        .component('gatewayGpeditDefault', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/authority/gpedit/default/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', '$rootScope', '$state', 'CODE', 'GatewayResource', '$filter'];

    function indexController($scope, $rootScope, $state, CODE, GatewayResource, $filter) {
        var vm = this;
        vm.data = {
            info: {
                copyText: ''
            },
            interaction: {
                request: {
                    gatewayAlias: $state.params.gatewayAlias
                },
                response: {
                    query: []
                }
            },
            fun: {
                delete: null, //删除策略组功能函数
                edit: null, //编辑策略组功能函数
                init: null //初始化功能函数
            }
        }
        vm.component = {
            listDefaultCommonObject: null,
            menuObject: {
                list: null
            }
        }
        vm.data.fun.init = function (arg) {
            var template = {
                promise: null,
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias
                }
            }
            template.promise = GatewayResource.Strategy.Query(template.request).$promise;
            template.promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.data.interaction.response.query = response.strategyList;
                            break;
                        }
                }
            })
            return template.promise;
        }
        vm.data.fun.edit = function (status, arg) {
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    strategyName: ''
                }
            }
            $rootScope.GroupModal(status == 'edit' ? '修改策略组' : '新增策略组', {
                groupName:status=='edit'?arg.item.strategyName:''
            }, '策略名称', null, function (callback) {
                if (callback) {
                    template.request.strategyName = callback.groupName;
                    switch (status) {
                        case 'edit':
                            {
                                template.request.strategyID = arg.item.strategyID;
                                GatewayResource.Strategy.Edit(template.request)
                                .$promise.then(function (response) {
                                    switch (response.statusCode) {
                                        case CODE.COMMON.SUCCESS:
                                            {
                                                $rootScope.InfoModal('修改策略组成功！', 'success');
                                                arg.item.strategyName = callback.groupName;
                                                arg.item.updateTime = $filter('currentTimeFilter')();
                                                break;
                                            }
                                    }
                                })
                                break;
                            }
                        default:
                            {
                                GatewayResource.Strategy.Add(template.request)
                                .$promise.then(function (response) {
                                    switch (response.statusCode) {
                                        case CODE.COMMON.SUCCESS:
                                            {
                                                $rootScope.InfoModal('新增策略组成功！', 'success');
                                                vm.data.interaction.response.query.splice(0, 0, {
                                                    strategyName: callback.groupName,
                                                    strategyID: response.strategyID,
                                                    updateTime: $filter('currentTimeFilter')()
                                                })
                                                break;
                                            }
                                    }
                                })
                                break;
                            }
                    }
                }
            });
        }
        vm.data.fun.delete = function (arg) {
            $rootScope.EnsureModal('删除策略组', false, '确认删除', {}, function (data) {
                if (data) {
                    GatewayResource.Strategy.Delete({
                            gatewayAlias: vm.data.interaction.request.gatewayAlias,
                            strategyID: arg.item.strategyID
                        }).$promise
                        .then(function (response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS:
                                    {
                                        vm.data.interaction.response.query.splice(arg.$index, 1);
                                        $rootScope.InfoModal('策略组删除成功', 'success');
                                        break;
                                    }
                                default:
                                    {
                                        $rootScope.InfoModal('删除失败，请稍候再试或到论坛提交bug', 'error');
                                        break;
                                    }
                            }
                        })
                }
            });
        }
        vm.$onInit = function () {
            vm.component.listDefaultCommonObject = {
                mainObject: {
                    item: {
                        default: [{
                                key: '策略组名称',
                                html: '{{item.strategyName}}'
                            },
                            {
                                key: 'ID',
                                html: '{{item.strategyID}}'
                            },
                            {
                                key: '最后修改时间',
                                html: '{{item.updateTime}}'
                            },
                        ],
                        fun: {
                            array: [{
                                    key: '复制',
                                    fun: function (arg) {
                                        var template = {
                                            elem: document.createElement("textarea")
                                        }
                                        template.elem.style.height = '0px';
                                        template.elem.style.width = '0px';
                                        template.elem.style.opacity = '0';
                                        template.elem.value = arg.item.strategyID;
                                        document.body.appendChild(template.elem);
                                        template.elem.select();
                                        template.elem.click();
                                        try {
                                            if (document.execCommand('copy')) {
                                                $rootScope.InfoModal('复制成功', 'success');
                                            } else {
                                                $rootScope.InfoModal('复制失败', 'error');
                                            }

                                        } catch (err) {
                                            $rootScope.InfoModal('复制失败', 'error');
                                        }
                                    },
                                    show: -1
                                },
                                {
                                    icon: 'bianji',
                                    key: '修改',
                                    fun: vm.data.fun.edit,
                                    show: -1,
                                    params: '"edit",arg'
                                },
                                {
                                    icon: 'shanchu',
                                    key: '删除',
                                    fun: vm.data.fun.delete,
                                    show: -1
                                }
                            ],
                            power: 1
                        }
                    },
                    baseInfo: {
                        colspan: 4,
                        warning: '尚未新建任何策略组'
                    },
                    baseFun: {
                        click: function (arg) {
                            $state.go('home.gateway.inside.gpedit.auth', {
                                strategyID: arg.item.strategyID
                            });
                        }
                    }
                }
            }
            vm.component.menuObject.list = [{
                type: 'btn',
                class: 'btn-group-li pull-left',
                btnList: [{
                    name: '新建策略组',
                    icon: 'tianjia',
                    class: 'eo-button-success',
                    fun: {
                        default: vm.data.fun.edit,
                        params: '"add"'
                    }
                }]
            }]
        }
    }
})();