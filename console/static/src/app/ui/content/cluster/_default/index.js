(function () {
    'use strict';
    angular.module('eolinker')
        .component('clusterDefault', {
            templateUrl: 'app/ui/content/cluster/_default/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource','CODE', '$rootScope', '$state'];

    function indexController($scope, GatewayResource,CODE, $rootScope, $state) {
        var vm = this;
        vm.ajaxResponse = {
            query: []
        }
        vm.fun = {};
        vm.component = {
            listRequireObject: null
        }
        let privateFun={};
        vm.fun.init = function () {
            $rootScope.global.ajax.Query_Cluster = GatewayResource.Cluster.Query();
            $rootScope.global.ajax.Query_Cluster.$promise.then(function (response) {
                vm.ajaxResponse.query = response.clusters || [];
            })
            return $rootScope.global.ajax.Query_Cluster.$promise;
        }
        /**
         * @desc 进入集群内页
         */
        privateFun.inToCluster=(inputArg)=>{
            $state.go('home.cluster.node.default',{
                cluster:inputArg.item.name,
                clusterName:inputArg.item.title
            });
        }
        /**
         * @desc 编辑集群
         */
        privateFun.editCluster = function (inputArg) {
            let tmpObj={
                edit:{
                    title:"修改集群",
                    opr:"Edit",
                    disabled:true,
                    item:inputArg.item
                },
                add:{
                    title:"新增集群",
                    opr:"Add",
                    item:{},
                    tip:"创建后不可修改"
                }
            };
            inputArg.item=tmpObj[inputArg.status].item;
            let tmpModal = {
                title: tmpObj[inputArg.status].title,
                resource: GatewayResource.Cluster[tmpObj[inputArg.status].opr],
                textArray: [{
                    type: 'input',
                    title: `Primary Key`,
                    value: inputArg.item.name||"",
                    key:"name",
                    disabled:tmpObj[inputArg.status].disabled,
                    placeholder:"具有唯一性,支持英文(不区分大小写)、下划线、数字",
                    pattern:"^[a-zA-Z][a-zA-z0-9_]*$",
                    tip:tmpObj[inputArg.status].tip||"",
                    required:true
                },{
                    type: 'input',
                    title: '集群名称',
                    value: inputArg.item.title||"",
                    key:"title",
                    placeholder:"具有唯一性,支持中文、英文(不区分大小写)、下划线、数字",
                    required:true
                },{
                    type: 'input',
                    title: '备注',
                    key:"note",
                    value: inputArg.item.note||"",
                    placeholder:"集群备注"
                }]
            }
            $rootScope.MixInputModal(tmpModal,(callback)=>{
                if(callback){
                    vm.fun.init();
                    $rootScope.InfoModal(`${tmpModal.title}成功`,'success');
                }
            })
        }
        privateFun.delete = function (inputArg) {
            let tmpAjaxRequest={
                name:inputArg.item.name
            },tmpModal={
                title:"删除集群"
            }
            $rootScope.EnsureModal(tmpModal.title, null, '确认删除？', {btnMessage:'删除'}, function (callback) {
                if (callback) {
                    GatewayResource.Cluster.Delete(tmpAjaxRequest).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    vm.ajaxResponse.query.splice(inputArg.$index, 1);
                                    $rootScope.InfoModal(tmpModal.title + '成功', 'success');
                                    break;
                                }
                        }
                    })
                }
            });
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['网关节点']
            });
            vm.component.menuObject = {
                list: [{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    btnList: [{
                        name: '新建集群',
                        icon: 'jiahao',
                        class: 'eo_theme_btn_success block-btn',
                        fun: {
                            default: privateFun.editCluster,
                            params: {
                                status: 'add'
                            }
                        }
                    }]

                }],
                setting: {
                    class: "common-menu-fixed-seperate common-menu-lg",
                    titleAuthority: 'showTitle',
                    title: '网关节点',
                    secondTitle:"网关支持分集群管理节点"
                }
            }
            vm.component.listDefaultCommonObject = {
                setting:{
                    warning: '尚无任何集群',
                    defaultFoot:true
                },
                item: {
                    default: [{
                        key: '集群名称',
                        html: '{{item.title}}'
                    }, {
                        key: '备注',
                        html: '{{item.note}}'
                    }, {
                        key: 'Primary Key',
                        html: '{{item.name}}'
                    }],
                    operate: {
                        funArr: [{
                                key: '查看详情',
                                fun: privateFun.inToCluster
                            },
                            {
                                key: '修改',
                                fun: privateFun.editCluster,
                                params: {
                                    status: 'edit'
                                }
                            }, {
                                key: `<span>删除</span><tip-directive ng-if="item.nodeCount" input="删除集群前需先移除集群内的节点"></tip-directive>`,
                                fun: privateFun.delete,
                                itemExpression:`ng-disabled="item.nodeCount"`
                            }
                        ],
                        power: -1,
                        class:'w_200'
                    }
                },
                baseFun:{
                    click:privateFun.inToCluster
                }
            }
        }
    }
})();