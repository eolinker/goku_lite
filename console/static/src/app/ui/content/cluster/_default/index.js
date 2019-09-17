(function () {
    'use strict';
    angular.module('eolinker')
        .component('clusterDefault', {
            templateUrl: 'app/ui/content/cluster/_default/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$rootScope', '$state'];

    function indexController($scope, GatewayResource, $rootScope, $state) {
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
         * @name 进入集群内页
         */
        privateFun.inToCluster=(inputArg)=>{
            $state.go('home.cluster.node.default',{
                cluster:inputArg.item.name,
                clusterName:inputArg.item.title
            });
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['网关节点']
            });
            vm.component.menuObject = {
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
                        key: '数据库地址',
                        html: `{{item.db.host+':'+item.db.port+'/'+item.db.database}}`
                    }],
                    operate: {
                        funArr: [{
                                key: '查看详情',
                                fun: privateFun.inToCluster
                            }
                        ],
                        power: -1,
                        class:'w_100'
                    }
                },
                baseFun:{
                    click:privateFun.inToCluster
                }
            }
        }
    }
})();