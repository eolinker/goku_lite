(function () {
    'use strict';
    /*
     * author：riverLethe
     * 鉴权相关js
     */
    angular.module('goku')
        .component('gatewayGpeditAuth', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/authority/gpedit/auth/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope) {
        var vm = this;
        vm.data = {
            info: {
                menu: [{
                    key: '无认证',
                    id: 'none'
                }, {
                    key: 'API Key',
                    id: 'apiKey'
                }, {
                    key: 'Basic认证',
                    id: 'basic'
                }]
            },
            interaction: {
                request: {
                    gatewayAlias: $state.params.gatewayAlias,
                    strategyID: $state.params.strategyID
                },
                response: {

                }
            }
        }
        vm.fun = {};
        vm.component = {
            menuObject: {
                show: {
                    status: {
                        isEdit: false
                    }
                },
                list: null
            }
        }
        var data = {
            templateResponse: null
        }
        var assistantFun = {}
        vm.fun.cancle = function () {
            vm.component.menuObject.show.status.isEdit = false;
            vm.data.interaction.response.authInfo = angular.copy(data.templateResponse);
        }
        vm.fun.refreshApiKey = function () {
            vm.data.interaction.response.authInfo.apiKey = CryptoJS.MD5((new Date()).getTime().toString()).toString();
        }
        vm.fun.edit = function () {
            if ($scope.ConfirmForm.$invalid) {
                $scope.ConfirmForm.$submitted = true;
                return;
            }
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    strategyID: $state.params.strategyID,
                    auth: vm.data.interaction.response.authInfo.auth
                },
                promise: null
            }
            switch (template.request.auth) {
                case 'basic':
                    {
                        template.request.basicUserPassword = vm.data.interaction.response.authInfo.basicUserPassword;
                        template.request.basicUserName = vm.data.interaction.response.authInfo.basicUserName;
                        break;
                    }
                case 'apiKey':
                    {
                        template.request.apiKey = vm.data.interaction.response.authInfo.apiKey;
                        break;
                    }
            }
            template.promise = GatewayResource.Auth.Edit(template.request).$promise;
            template.promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                    case '190000':
                        {
                            $rootScope.InfoModal('保存鉴权信息成功！', 'success');
                            vm.component.menuObject.show.status.isEdit = false;
                            break;
                        }
                    default:
                        {
                            $rootScope.InfoModal('保存鉴权信息操作失败！', 'error');
                        }
                }
            })
            return template.promise;
        }

        assistantFun.init = (function () {
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    strategyID: vm.data.interaction.request.strategyID
                }
            }
            GatewayResource.Auth.Info(template.request).$promise.then(function (response) {
                vm.data.interaction.response.authInfo = response.authInfo || {
                    auth: 'none',
                    basicUserName: '',
                    basicUserPassword: ''
                };
                vm.data.interaction.response.authInfo.apiKey = vm.data.interaction.response.authInfo.apiKey || CryptoJS.MD5((new Date()).getTime().toString()).toString();
                data.templateResponse = angular.copy(vm.data.interaction.response.authInfo);
            })
        })();
        vm.$onInit = function () {
            vm.component.menuObject.list = [{
                type: 'btn',
                class: 'btn-group-li pull-left',
                showVariable: 'isEdit',
                showPoint: 'status',
                btnList: [{
                    name: '策略组列表',
                    icon: 'xiangzuo',
                    show: -1,
                    fun: {
                        default: function () {
                            $state.go('home.gateway.inside.gpedit.default', {
                                strategyID: null
                            });
                        }
                    }
                }, {
                    name: '编辑',
                    class: 'eo-button-success',
                    show: false,
                    fun: {
                        default: function () {
                            vm.component.menuObject.show.status.isEdit = true;
                        }
                    }
                }, {
                    name: '保存',
                    class: 'eo-button-success',
                    show: true,
                    fun: {
                        default: vm.fun.edit
                    }
                }, {
                    name: '取消',
                    class: 'eo-button-default',
                    show: true,
                    fun: {
                        default: vm.fun.cancle
                    }
                }]
            }]
        }
    }
})();