(function () {
    'use strict';
    angular.module('eolinker')
        .component('alertList', {
            templateUrl: 'app/ui/content/alert/list/index.html',
            controller: indexController
        })

    indexController.$inject = [ '$scope', 'GatewayResource', '$rootScope', 'CODE','Authority_CommonService'];

    function indexController($scope, GatewayResource, $rootScope, CODE,Authority_CommonService) {
        var vm = this;
        vm.data = {
            pagination: {
                maxSize: 10,
                pageSize: 50,
                page: 0,
                msgCount: 0
            }
        }
        vm.ajaxResponse={
            query: []
        };

        vm.component = {
            menuObject: {},
            listRequireObject: null
        }
        vm.fun={};
        var privateFun = {},data={};
        vm.service={
            authority:Authority_CommonService
        }
        vm.fun.init = function () {
            return privateFun.initLoadAjax();
        }
        privateFun.initLoadAjax=()=>{
            let tmpAjaxRequest={
                pageSize: vm.data.pagination.pageSize,
                page:vm.data.pagination.page+1
            },tmpResponseArr=vm.ajaxResponse.query||[];
            data.isQuerying = true;
            $rootScope.global.ajax.Query_AlertMessage = GatewayResource.AlertMessage.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_AlertMessage.$promise.then(function (response) {
                vm.ajaxResponse.query =tmpResponseArr.concat(response.alertMessageList || []);
                vm.data.pagination.msgCount = response.page.totalNum||0;
                vm.data.pagination.page++;
                data.isQuerying = false;
            })
            return $rootScope.global.ajax.Query_AlertMessage.$promise;
        }
        privateFun.scrollLoading = function () {
            let flag = {
                hasItem: vm.ajaxResponse.query && vm.ajaxResponse.query.length !== 0,
                hasNextPage: vm.data.pagination.page < (vm.data.pagination.msgCount / vm.data.pagination.pageSize),
                isQuerying: data.isQuerying
            }
            if (flag.hasItem && flag.hasNextPage && !flag.isQuerying) {
                privateFun.initLoadAjax('preload');
            }
        }
        privateFun.delete = function (status, arg) {
            var template = {
                request: {
                    alertID: status == 'empty' ? null : arg.item.alertID
                },
                modal: {
                    title: status == 'empty' ? '清空告警信息' : '删除告警信息'
                },
                resource:status=='empty'?GatewayResource.AlertMessage.Empty:GatewayResource.AlertMessage.Delete
            }
            $rootScope.EnsureModal(template.modal.title, null, '确认'+(status == 'empty'?'清空':'删除')+'？', {btnMessage:status == 'empty'?'清空':'删除'}, function (callback) {
                if (callback) {
                    template.resource(template.request).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    switch (status) {
                                        case 'empty':
                                            {
                                                vm.ajaxResponse.query=[];
                                                break;
                                            }
                                        case 'single':
                                            {
                                                vm.ajaxResponse.query.splice(arg.$index, 1);
                                                break;
                                            }
                                    }
                                    $rootScope.InfoModal(template.modal.title + '成功', 'success');
                                    break;
                                }
                        }
                    })
                }
            });
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['告警列表']
            });
            vm.component.listDefaultCommonObject = {
                baseFun:{
                    scrollLoading: privateFun.scrollLoading
                },
                item: {
                    default: [{
                        key: '时间',
                        html: '{{item.updateTime}}',
                        keyStyle: {
                            'width': '200px'
                        }
                    },{
                        key: '集群名称',
                        html: '{{item.title}}',
                        keyStyle: {
                            'width': '200px'
                        }
                    },{
                        key:'节点',
                        html:'{{item.addr}}',
                        keyStyle: {
                            'width': '200px'
                        }
                    }, {
                        key: '请求URI',
                        html: '{{item.requestURL}}'
                    }, {
                        key: '转发地址',
                        html: '{{item.targetServer}}'
                    }, {
                        key: '转发URI',
                        html: '{{item.targetURL}}'
                    }, {
                        key: '描述',
                        html: '{{item.msg}}'
                    }],
                    operate: {
                        funArr: [
                            {
                                key: '删除',
                                fun: privateFun.delete,
                                params: '"single",arg'
                            }
                        ],
                        class:'w_150'
                    }
                },
                setting: {
                    page:true,
                    isFixedHeight: true,
                    unhover: true,
                    fixFoot:true,
                    scroll: true,
                    scrollRemainRatio: 7
                }
            }
            vm.component.menuObject = {
                list: [{
                        type: 'btn',
                        authority: 'edit',
                        class: 'pull-left',
                        btnList: [{
                            name: '清空告警信息',
                            icon: 'shanchu',
                            class: 'eo_theme_btn_danger block-btn',
                            fun: {
                                default: privateFun.delete,
                                params:'"empty"'
                            }
                        }]
                    }],
                setting: {
                    titleAuthority: 'showTitle',
                    title: '告警列表'
                }
            }
        }
    }
})();