(function () {
    'use strict';
    /**
     * 测试导入增强插件
     * @param [object] input 输入信息
     * @param [object] output flag：标识是否已经更改,visible：是否可视
     */
    angular.module('eolinker')
        .component('pluginOperate', {
            templateUrl: 'app/ui/content/plugin/operate/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope) {
        var vm = this;
        vm.data = {
            from: $state.params.status,
            status: $state.params.status,
            info: {},
            interaction: {
                response: {
                    pluginInfo: {
                        pluginName: $state.params.pluginName || '',
                        pluginPriority: '',
                        pluginConfig: '',
                        pluginDesc: '',
                        isStop: '1',
                        pluginType: '',
                        // version: ''
                    }
                }
            }
        }
        vm.component = {
            menuObject: {
                list: []
            }
        };
        vm.fun={};
        let privateFun={};
        vm.fun.init = function () {
            if ($state.params.status == 'add') {
                vm.data.interaction.response.pluginInfo.pluginType = '0';
                return;
            }
            var template = {
                request: {
                    pluginName: vm.data.interaction.response.pluginInfo.pluginName
                }
            }
            GatewayResource.Plugin.Info(template.request).$promise.then(function (response) {
                response.pluginInfo.isStop = response.pluginInfo.isStop.toString();
                response.pluginInfo.pluginType = response.pluginInfo.pluginType.toString();
                vm.data.interaction.response.pluginInfo = response.pluginInfo;
                $scope.$broadcast('$Maunal_AceEditorAms', response.pluginInfo.pluginConfig || '');
            });


        };
        vm.fun.back = function () {
            $state.go('home.plugin.default');
        }
        privateFun.confirm = function () {
            var template = {
                output: {
                    pluginName: vm.data.interaction.response.pluginInfo.pluginName,
                    pluginPriority: vm.data.interaction.response.pluginInfo.pluginPriority,
                    pluginDesc: vm.data.interaction.response.pluginInfo.pluginDesc,
                    isStop: vm.data.interaction.response.pluginInfo.isStop,
                    pluginType: vm.data.interaction.response.pluginInfo.pluginType,
                    // version: vm.data.interaction.response.pluginInfo.version
                }
            }
            switch (template.output.pluginType) {
                case '0':
                    {
                        template.output.pluginConfig = vm.data.interaction.response.pluginInfo.pluginConfig;
                        break;
                    }
            }
            return template.output;
        }
        vm.fun.load = function (arg) {
            $scope.$emit('$TransferStation', {
                state: '$Init_LoadingCommonComponent',
                data: arg
            });
        }
        vm.fun.requestProcessing = function (arg) { //arg status:（0：继续添加 1：快速保存，2：编辑（修改/新增））
            var template = {
                request: privateFun.confirm(),
                promise: null
            }
            vm.data.info.submited = false;
            if ($scope.editForm.$valid && ((template.request.pluginConfig && template.request.pluginType == '0') || (template.request.pluginType != '0'))) {
                template.promise = privateFun.edit({
                    request: template.request
                });
            } else if (template.request.pluginConfig) {
                vm.data.info.submited = true;
                $rootScope.InfoModal('插件编辑失败，请检查信息是否填写完整！', 'error');
            } else {
                vm.data.info.submited = true;
                $rootScope.InfoModal('插件编辑失败，未填写插件配置文件信息！', 'error');
            }
            return template.promise;
        }
        privateFun.edit = function (arg) {
            var template = {
                promise: null
            }
            if ($state.params.status == 'edit') {
                template.promise = GatewayResource.Plugin.Edit(arg.request).$promise;
            } else {
                template.promise = GatewayResource.Plugin.Add(arg.request).$promise;
            }
            template.promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.fun.back();
                            $rootScope.InfoModal(($state.params.status == 'edit' ? '修改' : '添加') + '插件成功', 'success');
                            break;
                        }
                }
            })
            return template.promise;
        }
        vm.$onInit = function () {
            var template = {
                array: []
            }
            vm.component.menuObject={
                list:[{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    btnList: [{
                        name: '返回列表',
                        icon: 'chexiao',
                        fun: {
                            default: vm.fun.back
                        }
                    }]
                }, {
                    type: 'btn',
                    class: 'btn-group-li',
                    btnList: [{
                        name: '保存',
                        class: 'eo_theme_btn_success block-btn',
                        fun: {
                            disabled: 1,
                            default: vm.fun.requestProcessing,
                            params: {
                                status: 1
                            }
                        }
                    }]
                }],
                setting:{
                    class:'common-menu-fixed-seperate'
                }
            };
        }
    }
})();