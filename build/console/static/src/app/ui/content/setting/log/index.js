(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 网关设置，基础配置相关js
     */
    angular.module('eolinker')
        .component('settingLog', {
            templateUrl: 'app/ui/content/setting/log/index.html',
            controller: indexController
        })

    indexController.$inject = ['$rootScope', 'CODE', 'GatewayResource', '$scope', 'Authority_CommonService'];

    function indexController($rootScope, CODE, GatewayResource, $scope, Authority_CommonService) {
        var vm = this,
            cache = {
                accessLog: null,
                nodeLog: null,
                consoleLog: null
            },
            privateFun = {};
        vm.data = {
            nodeText: 'info',
            consoleText: 'warn'
        }
        vm.ajaxResponse = {
            nodeLog:{},
            consoleLog:{},
            accessLog:{}
        }
        vm.fun = {};
        vm.component = {
            progressBarObj: {
                setting:{
                    value:'name',
                    key:'title'
                }
            },
            listBlockObj: {
                tdList: [{
                    type: 'sort',
                    itemExpression: 'ng-show="$ctrl.authorityObject.edit"'
                }, {
                    type: 'checkbox',
                    modelKey:'select',
                    authority:'edit'
                }, {
                    type: 'text',
                    thKey: '字段名',
                    modelKey: 'name',
                    class:'w_20percent'
                }, {
                    type: 'html',
                    thKey: '描述',
                    html:`<span title="{{item.desc}}">{{item.desc}}</span>`
                }]
            },
            navigationMenuObject: {}
        }

        vm.service = {
            authority: Authority_CommonService
        }
        vm.fun.saveForm = function (inputWhich) {
            let tmpOprWhich=inputWhich.charAt(0).toUpperCase() + inputWhich.slice(1);
            vm.data.submitted = true;
            if ($scope[tmpOprWhich+'Form'].$invalid) return;
            let tmpResponseObj=vm.ajaxResponse[inputWhich+'Log'],tmpAjaxRequest={
                enable:tmpResponseObj.enable,
                dir:tmpResponseObj.dir,
                file:tmpResponseObj.file,
                period:tmpResponseObj.period,
                expire:tmpResponseObj.expire
            }
            switch(inputWhich){
                case 'access':{
                    tmpAjaxRequest.fields=JSON.stringify(tmpResponseObj.fields,(key,val)=>{
                        if(/(desc)|(isHide)/.test(key))return undefined;
                        return val;
                    });
                    break;
                }
                default:{
                    tmpAjaxRequest.level=tmpResponseObj.level;
                    break;
                }
            }
            let tmpPromise=GatewayResource.ConfigLog['Set'+tmpOprWhich](tmpAjaxRequest).$promise;
            tmpPromise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS: {
                        $rootScope.InfoModal('保存成功！', 'success');
                        vm.data[inputWhich+'IsEdit']=false;
                        cache[inputWhich+'Log']=angular.copy(vm.ajaxResponse[inputWhich+'Log']);
                        break;
                    }
                }
            })
            return tmpPromise;
        }
        vm.fun.cancelEdit = function (inputWhich) {
            vm.data.submitted = false;
            vm.data[inputWhich+'IsEdit']=false;
            vm.ajaxResponse[inputWhich+'Log'] = angular.copy(cache[inputWhich+'Log']);
        }
        privateFun.changeMenu = (inputArg) => {
            vm.data.menuType = inputArg.item.active;
            switch (vm.data.menuType) {
                case 0: {
                    if(!cache.accessLog){
                        GatewayResource.ConfigLog.Access({
                            t: new Date().getTime()
                        }).$promise.then((response) => {
                            vm.ajaxResponse.accessLog = response.data;
                            cache.accessLog = angular.copy(vm.ajaxResponse.accessLog);
                        })
                    }
                    break;
                }
                case 1: {
                    if(!cache.consoleLog){
                        GatewayResource.ConfigLog.Console({
                            t: new Date().getTime()
                        }).$promise.then((response) => {
                            vm.ajaxResponse.consoleLog = response.data;
                            cache.consoleLog = angular.copy(vm.ajaxResponse.consoleLog);
                        })
                    }
                    if(!cache.nodeLog){
                        GatewayResource.ConfigLog.Node({
                            t: new Date().getTime()
                        }).$promise.then((response) => {
                            vm.ajaxResponse.nodeLog = response.data;
                            cache.nodeLog = angular.copy(vm.ajaxResponse.nodeLog);
                        })
                    }
                    break;
                }
            }
        }
        vm.fun.wantToEdit=(inputWhich)=>{
            vm.data[inputWhich+'IsEdit']=true;
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['日志设置', '网关设置']
            });
            vm.component.navigationMenuObject = {
                list: [{
                    type: 'navigation',
                    class: 'menu-navigation',
                    activePoint: 'menuType',
                    tabList: [{
                        name: '请求日志',
                        active: 0,
                        fun: {
                            default: privateFun.changeMenu
                        }
                    }, {
                        name: '运行日志',
                        active: 1,
                        fun: {
                            default: privateFun.changeMenu
                        }
                    }]
                }],
                setting: {
                    class: 'common-menu-fixed-seperate common-menu-only-navigation'
                }
            }
            privateFun.changeMenu({
                item:{
                    active:0
                }
            })
        };
    }
})();