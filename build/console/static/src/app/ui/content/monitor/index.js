(function () {
    'use strict';
    /**
     * @name API监控设置
     * @author 广州银云信息科技有限公司
     */
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.monitor', {
                    url: '/monitor',
                    template: '<monitor></monitor>'
                })
        }])
        .component('monitor', {
            templateUrl: 'app/ui/content/monitor/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$rootScope', 'CODE'];

    function indexController($scope, GatewayResource, $rootScope, CODE) {
        var vm = this;
        vm.ajaxResponse={}
        vm.fun={};
        let privateFun={};

        /**
         * @desc 取消/关闭当前监控模块
         */
        privateFun.cancelModule=(inputModuleItem)=>{
            GatewayResource.MonitorModuleConf.Set({
                moduleName:inputModuleItem.moduleName,
                moduleStatus:0
            }).$promise.then((response)=>{
                switch(response.statusCode){
                    case CODE.COMMON.SUCCESS:{
                        inputModuleItem.moduleStatus=0;
                        break;
                    }
                }
            })
        }
        /**
         * @desc 保存当前监控模块
         */
        privateFun.saveModule=(inputModuleItem,inputOptions={})=>{
            GatewayResource.MonitorModuleConf.Set({
                moduleName:inputModuleItem.moduleName,
                moduleStatus:1,
                config:JSON.stringify(inputModuleItem.config)
            }).$promise.then((response)=>{
                switch(response.statusCode){
                    case CODE.COMMON.SUCCESS:{
                        if(inputOptions.isStatic){
                            inputModuleItem.moduleStatus=1;
                        }else{
                            $rootScope.InfoModal("保存成功","success");
                        }
                        break;
                    }
                }
            })
        }
        /**
         * @desc 临时开启当前监控模块
         */
        privateFun.tmpOpenModule=(inputModuleItem)=>{
            inputModuleItem.moduleStatus=1;
        }
        vm.fun.oprModule=(inputOpr,inputItem,inputForm)=>{
            switch(inputOpr){
                case "save":{
                    if(inputForm.$invalid)return;
                    privateFun.saveModule(inputItem);
                    break;
                }
                case "cancel":{
                    privateFun.cancelModule(inputItem);
                    break;
                }
                case "on-off":{
                    if(inputItem.moduleStatus===1){
                        privateFun.cancelModule(inputItem);
                    }else if(inputItem.config){
                        privateFun.tmpOpenModule(inputItem);
                    }else{
                        privateFun.saveModule(inputItem,{
                            isStatic:true
                        });
                    }
                    break;
                }
            }
        }
        /**
         * @desc ajax后台，获取配置
         */
        privateFun.ajaxBackEnd=()=>{
            let tmpPromise=GatewayResource.MonitorModuleConf.Get().$promise;
            tmpPromise.then((response)=>{
                vm.ajaxResponse.list=response.moduleList||[];
            })
            return tmpPromise;
        }
        vm.fun.init=privateFun.ajaxBackEnd;
    }
})();