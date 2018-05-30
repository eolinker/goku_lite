(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 环境管理相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.config', {
                    url: '/config',
                    template: '<config></config>'
                });
        }])
        .component('config', {
            templateUrl: 'app/ui/content/home/config/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', '$rootScope', '$state', 'CODE', 'GatewayResource'];

    function indexController($scope, $rootScope, $state, CODE, GatewayResource) {
        var vm = this;
        vm.data = {
            info: {},
            interaction: {
                response: {
                    confInfo: {}
                }
            },
            fun: {
                cancle: null, //取消功能函数
                edit: null //编辑功能函数
            }
        }
        var data={
            template:{
                confInfo:{}
            }
        }
        vm.data.fun.operate=function(status){
            var template={
                promise:null,
                successMessage:'',
                modal:{
                    title:'',
                    secondTitle:''
                }
            }
            switch(status){
                case 'reload':{
                    template.modal.title='重载网关';
                    template.modal.secondTitle='确认重载网关？';
                    
                    break;
                }
                case 'restart':{
                    template.modal.title='重启网关';
                    template.modal.secondTitle='确认重启网关？';
                    
                    break;
                }
                case 'basic':{
                    template.modal.title=(vm.data.interaction.response.confInfo.gatewayServiceStatus?'停止':'开启')+'网关';
                    template.modal.secondTitle='确认'+(vm.data.interaction.response.confInfo.gatewayServiceStatus?'停止':'开启')+'网关？';
                    break;
                }
            }
            $rootScope.EnsureModal(template.modal.title, false, template.modal.secondTitle, { btnType: 2, btnMessage: '确定' }, function(callback) {
                if (callback) {
                    switch(status){
                        case 'reload':{
                            template.successMessage='重载成功';
                            template.promise=GatewayResource.Service.Reload().$promise;
                            break;
                        }
                        case 'restart':{
                            template.successMessage='重启成功';
                            template.promise=GatewayResource.Service.Restart().$promise;
                            break;
                        }
                        case 'basic':{
                            if(vm.data.interaction.response.confInfo.gatewayServiceStatus){
                                template.successMessage='网关已停止运行';
                                template.promise=GatewayResource.Service.Stop().$promise;
                            }else{
                                template.successMessage='开启成功';
                                template.promise=GatewayResource.Service.Start().$promise;
                            }
                            break;
                        }
                    }
                    template.promise.then(function(response){
                        switch(response.statusCode){
                            case '000000':{
                                if(status=='basic'){
                                    vm.data.interaction.response.confInfo.gatewayServiceStatus=!vm.data.interaction.response.confInfo.gatewayServiceStatus;
                                }
                                $rootScope.InfoModal(template.successMessage,'success');
                                break;
                            }
                            default:{
                                $rootScope.InfoModal('设置失败！','error');
                                break;
                            }
                        }
                    })
                }
            });
        }
        vm.data.fun.edit = function () {
            var template = {
                request: {
                    gatewayPort: vm.data.interaction.response.confInfo.gatewayPort
                }
            }
            GatewayResource.Config.Edit(template.request).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            $rootScope.InfoModal('修改成功！', 'success');
                            vm.data.info.status.config.isEdit = false;
                            break;
                        }
                    default:
                        {
                            $rootScope.InfoModal('操作失败！', 'error');
                        }
                }
            })
        }
        vm.data.fun.cancle = function () {
            vm.data.info.status.config.isEdit = false;
            vm.data.interaction.response.confInfo=angular.copy(data.template.confInfo);
        }
        vm.$onInit=function () {
            GatewayResource.Config.Info().$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.data.interaction.response.confInfo = response.confInfo;
                            angular.copy(response.confInfo, data.template.confInfo);
                            break;
                        }
                }
            })
        };
    }
})();