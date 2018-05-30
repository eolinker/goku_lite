(function () {
    'use strict';
    /*
     * author：riverLethe
     * 流量控制相关js
     */
    angular.module('goku')
        .component('gatewayGpeditRate', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/authority/gpedit/rate/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope) {
        var vm = this;
        vm.data = {
            interaction: {
                request: {
                    strategyID: $state.params.strategyID,
                    gatewayAlias: $state.params.gatewayAlias
                },
                response: {
                    query: []
                }
            },
            fun: {
                delete: null, //删除请求频率限制功能函数
                edit: null, //编辑请求频率限制功能函数
                init: null //初始化功能函数
            }
        }
        vm.component = {
            listRequireObject: {
                mainObject: {
                    tdList: [],
                    fun: {}
                }
            },
            menuObject: {
                list: null
            }
        }
        vm.data.fun.init = function (arg) {
            var template = {
                promise: null,
                request: {
                    strategyID: vm.data.interaction.request.strategyID,
                    gatewayAlias: vm.data.interaction.request.gatewayAlias
                }
            }
            template.promise = GatewayResource.RateLimit.Query(template.request).$promise;
            template.promise.then(function (response) {
                vm.data.interaction.response.query = response.limitList||[];
            })
            return template.promise;
        }
        vm.data.fun.edit = function (status,arg) {
            arg = arg || {};
            var template = {
                modal: {
                    config:{
                        title: status=='edit' ? '修改流控' : '创建流控',
                    },
                    request:arg.item||{
                        allow:true,
                        period:'sec',
                        startTime:0,
                        endTime:24
                    }
                },
                request: {
                    strategyID: vm.data.interaction.request.strategyID,
                    gatewayAlias: vm.data.interaction.request.gatewayAlias
                },
                promise:null
            }
            template.modal.request.period=template.modal.request.period||0;
            $rootScope.GatewayRateLimitModal(template.modal, function (callback) {
                if (callback) {
                    angular.merge(template.request,callback);
                    switch(status){
                        case 'edit':{
                            template.request.limitID=arg.item.limitID;
                            template.promise=GatewayResource.RateLimit.Edit(template.request).$promise;
                            break;
                        }
                        case 'add':{
                            template.promise=GatewayResource.RateLimit.Add(template.request).$promise;
                            break;
                        }
                    }
                    template.promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                            case '210000':
                                {
                                    $rootScope.InfoModal(template.modal.config.title+'成功！', 'success');
                                    vm.data.fun.init();
                                    break;
                                }
                            default:{
                                $rootScope.InfoModal('操作失败！', 'error');
                                break;
                            }
                        }
                    })
                }
            });
        }
        vm.data.fun.delete = function (arg) {
            $rootScope.EnsureModal('删除流控', false, '确认删除', {}, function (data) {
                if (data) {
                    GatewayResource.RateLimit.Delete({
                            strategyID: vm.data.interaction.request.strategyID,
                            gatewayAlias: vm.data.interaction.request.gatewayAlias,
                            limitID: arg.item.limitID
                        }).$promise
                        .then(function (response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS:
                                    {
                                        vm.data.interaction.response.query.splice(arg.$index, 1);
                                        $rootScope.InfoModal('流控删除成功', 'success');
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
            vm.component.listRequireObject.mainObject.tdList = [{
                name: '访问类型',
                keyType: 'customized-html',
                keyHtml: '{{item.allow?\'允许\':\'禁止\'}}访问'
            }, {
                name: '单位时间',
                keyType: 'customized-html',
                keyHtml: '{{item.period==\'sec\'?\'1秒\':item.period==\'min\'?\'1分钟\':item.period==\'hour\'?\'1小时\':\'1天\'}}'
            }, {
                name: '请求频率限制',
                key: 'limit'
            }, {
                name: '时间段',
                keyType: 'customized-html',
                keyHtml: '{{item.startTime+\'~\'+item.endTime}}'
            }, {
                name: '优先级',
                key: 'priority'
            }, {
                name: '操作',
                keyType: 'btn',
                btnList: [{
                    name: '修改',
                    icon: 'bianji',
                    fun: {
                        default: vm.data.fun.edit,
                        params:'"edit",arg'
                    }
                }, {
                    name: '删除',
                    icon: 'shanchu',
                    fun: {
                        default: vm.data.fun.delete,
                        params: {
                            switch: 0
                        }
                    }
                }]
            }];
            vm.component.menuObject.list = [{
                type: 'btn',
                class: 'btn-group-li pull-left',
                btnList: [{
                    name: '策略组列表',
                    icon: 'xiangzuo',
                    fun: {
                        default: function(){
                            $state.go('home.gateway.inside.gpedit.default',{strategyID:null});
                        }
                    }
                },{
                    name: '创建流控',
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