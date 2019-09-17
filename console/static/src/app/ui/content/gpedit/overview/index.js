(function () {
    'use strict';
    angular.module('eolinker')
        .component('gpeditOverview', {
            templateUrl: 'app/ui/content/gpedit/overview/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope','$state','GatewayResource'];

    function indexController($scope,$state,GatewayResource) {
        var vm = this;
        vm.component = {
            menuObject: null
        }
        vm.fun={};
        vm.ajaxResponse={};
        let privateFun={};
        privateFun.ajaxOpenGpedit=()=>{
            GatewayResource.Strategy.Info({
                strategyType:1
            }).$promise.then((response)=>{
                vm.ajaxResponse.openGpeditInfo=response.strategyInfo||{};
            })
        }
        vm.fun.enterGpedit=(inputType)=>{
            switch(inputType){
                case "open":{
                    $state.go("home.gpedit.inside.api.default",{
                        groupType:"open",
                        strategyID:vm.ajaxResponse.openGpeditInfo.strategyID,
                        strategyName:vm.ajaxResponse.openGpeditInfo.strategyName
                    })
                    break;
                }
                case "common":{
                    $state.go("home.gpedit.common.list",{
                        groupType:"common"
                    });
                    break;
                }
            }
        }
        
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['访问策略']
            });
            vm.component.menuObject = {
                setting: {
                    class: "common-menu-fixed-seperate common-menu-lg",
                    titleAuthority: 'showTitle',
                    title: '访问策略',
                    secondTitle:"您可以给不同的调用方或应用设置访问策略，不同的访问策略可以设置不同的 API 访问权限、鉴权方式以及插件功能等"
                }
            }
            privateFun.ajaxOpenGpedit();
        }
    }
})();