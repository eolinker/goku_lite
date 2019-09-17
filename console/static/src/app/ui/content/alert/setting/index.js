(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 环境管理相关js
     */
    angular.module('eolinker')
        .component('alertSetting', {
            templateUrl: 'app/ui/content/alert/setting/index.html',
            controller: indexController
        })

    indexController.$inject = ['$rootScope', 'CODE', 'GatewayResource', '$scope', 'Authority_CommonService'];

    function indexController($rootScope, CODE, GatewayResource, $scope, Authority_CommonService) {
        var vm = this;
        vm.data = {
            isEdit: false,
            menu:'api'
        }
        vm.ajaxResponse = {
            alertInfo: {}
        }
        vm.fun = {};
        vm.CONST = {
            ALERT_METHOD_ARR:[{
                key:'API告警',
                value:'api',
                tip:"请求成功状态码在 [网关设置>基本设置] 页面设置，返回非成功状态码则视为请求失败；API的告警阀值在API编辑页面设置"
            }],
            ALERT_PROTOCOL_ARR: [{
                key: '不设置任何协议',
                value: 0
            }, {
                key: 'SSL协议',
                value: 1
            }, {
                key: 'TLS协议',
                value: 2
            }],
            alertPeriodTypeQuery: [{
                key: '1分钟',
                value: 0
            }, {
                key: '5分钟',
                value: 1
            }, {
                key: '15分钟',
                value: 2
            }, {
                key: '30分钟',
                value: 3
            }, {
                key: '60分钟',
                value: 4
            }]
        }
        vm.service = {
            authority: Authority_CommonService
        }
        var cache = {
            alertInfo: {}
        },privateFun={};
        vm.fun.startAlert = function () {
            if (vm.data.isEdit) {
                vm.ajaxResponse.alertInfo.alertStatus = vm.ajaxResponse.alertInfo.alertStatus ? 0 : 1;
            }
        }
        vm.fun.changeNotice = function (arg) {
            if (arg.$last) {
                vm.ajaxResponse.alertInfo[vm.data.menu+'AlertInfo'].userEmail.push({
                    value: ''
                });
            }
        }
        vm.fun.authorityToEdit = function () {
            vm.data.isEdit = true;
            if(vm.ajaxResponse.alertInfo.apiAlertInfo.userEmail[vm.ajaxResponse.alertInfo.apiAlertInfo.userEmail.length-1].value){
                vm.ajaxResponse.alertInfo.apiAlertInfo.userEmail.push({
                    value: ''
                });
            }
        }
        vm.fun.changeAlertMenu=(inputMenuType)=>{
            vm.data.menu=inputMenuType;
        }
        vm.fun.deleteNotice = function (arg) {
            vm.ajaxResponse.alertInfo[vm.data.menu+'AlertInfo'].userEmail.splice(arg.$index, 1);
        }
        vm.fun.checkIsValidEmail=(inputEmail)=>{
            return !/^[0-9A-Za-z-_.]+@[0-9a-z-]+\.[a-z]{2,20}(\.[a-z]{2,20}){0,1}$/.test(inputEmail)&&inputEmail;
        }
        privateFun.spliceUserEmailArr=(inputMarkArr)=>{
            for(let val of inputMarkArr){
                if(vm.ajaxResponse.alertInfo[val+'AlertInfo'].userEmail.length>1)vm.ajaxResponse.alertInfo[val+'AlertInfo'].userEmail.splice(vm.ajaxResponse.alertInfo[val+'AlertInfo'].userEmail.length - 1, 1);
            }
        }
        privateFun.setReceiverList=(inputArr)=>{
            let tmpOutput=[];
            for(let val of inputArr){
                if(vm.fun.checkIsValidEmail(val.value))return false;
                if (val.value) tmpOutput.push(val.value);
            }
            return tmpOutput.join(',');
        }
        vm.fun.editAlert = function () {
            vm.data.submitted=true;
            if ($scope.ConfirmForm.$invalid) {
                $rootScope.InfoModal('编辑失败，请检查信息是否填写完整！', 'error');
                return;
            }
            let tmpAjaxRequest = {
                alertStatus: vm.ajaxResponse.alertInfo.alertStatus,
                sender: vm.ajaxResponse.alertInfo.sender,
                senderPassword: vm.ajaxResponse.alertInfo.senderPassword,
                smtpAddress: vm.ajaxResponse.alertInfo.smtpAddress,
                smtpPort: vm.ajaxResponse.alertInfo.smtpPort,
                smtpProtocol: vm.ajaxResponse.alertInfo.smtpProtocol,
                apiAlertInfo: angular.copy(vm.ajaxResponse.alertInfo.apiAlertInfo)
            }
            tmpAjaxRequest.apiAlertInfo.receiverList=privateFun.setReceiverList(tmpAjaxRequest.apiAlertInfo.userEmail);
            if(tmpAjaxRequest.apiAlertInfo.receiverList===false){
                $rootScope.InfoModal('编辑失败，请检查信息是否填写完整！', 'error');
                return;
            }
            delete tmpAjaxRequest.apiAlertInfo.userEmail;
            tmpAjaxRequest.apiAlertInfo=JSON.stringify(tmpAjaxRequest.apiAlertInfo);
            GatewayResource.Config.AlertEdit(tmpAjaxRequest).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS: {
                        $rootScope.InfoModal('修改成功！', 'success');
                        vm.data.isEdit = false;
                        privateFun.spliceUserEmailArr(['api','node','redis']);
                        angular.copy(vm.ajaxResponse.alertInfo, cache.alertInfo);
                        break;
                    }
                }
            })
        }
        vm.fun.cancleAlert = function () {
            vm.data.isEdit = false;
            vm.data.userEmail = angular.copy(cache.userEmail);
            vm.ajaxResponse.alertInfo = angular.copy(cache.alertInfo);
        }
        privateFun.splitMailStr=(inputMailStr)=>{
            let tmpOutput=[];
            (inputMailStr.split(',')).map(function (val, key) {
                tmpOutput.push({
                    value: val
                })
            });
            return tmpOutput;
        }
        vm.$onInit = function () {
            GatewayResource.Config.AlertInfo().$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS: {
                        vm.ajaxResponse.alertInfo = response.gatewayConfig || {
                            "alertStatus": 0,
                            "sender": "",
                            "senderPassword": "",
                            "smtpAddress": "",
                            "smtpPort": '',
                            "smtpProtocol": 0,
                            "apiAlertInfo": {
                                alertPeriodType:0,
                                receiverList:'',
                                alertAddr:''
                            }
                        };
                        vm.ajaxResponse.alertInfo.apiAlertInfo.userEmail=privateFun.splitMailStr(vm.ajaxResponse.alertInfo.apiAlertInfo.receiverList);
                        angular.copy(vm.ajaxResponse.alertInfo, cache.alertInfo);
                        break;
                    }
                }
            })
        };
    }
})();