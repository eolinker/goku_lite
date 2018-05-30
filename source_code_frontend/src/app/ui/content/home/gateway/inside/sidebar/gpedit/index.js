(function() {
    /*
     * author：riverLethe
     * 网关内页页内包模块sidebar相关js
     */
    angular.module('goku')
        .component('gatewayGpeditSidebar', {
            template: '<group-common-component main-object="$ctrl.component.groupCommonObject.mainObject" fun-object="$ctrl.component.groupCommonObject.funObject"></group-common-component>',
            controller: indexController,
            bindings:{
                top:'<'
            }
        })

    indexController.$inject = ['$scope', '$state'];

    function indexController($scope, $state) {
        var vm = this;
        vm.component={
            groupCommonObject: null
        }
        vm.$onInit = function() {
            vm.component.groupCommonObject = {
                funObject:{
                    unTop:vm.top
                },
                mainObject: {
                    baseInfo: {
                        title:'策略组详情',
                        name: 'name',
                        id:'mark',
                        interaction:{
                            mark:$state.current.mark
                        }
                    },
                    staticQuery: [{
                        name:'鉴权方式',
                        sref:'home.gateway.inside.gpedit.auth',
                        mark:'auth'
                    },{
                        name:'IP黑白名单',
                        sref:'home.gateway.inside.gpedit.ip',
                        mark:'ip'
                    },{
                        name:'流量控制',
                        sref:'home.gateway.inside.gpedit.rate',
                        mark:'rate'
                    }],
                    baseFun: {
                        parentClick: function(arg){
                            $state.go(arg.item.sref);
                        }
                    }
                }
            }
        }
    }

})();
