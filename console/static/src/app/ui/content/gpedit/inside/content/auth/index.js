(function () {
    'use strict';
    /*
     * author：riverLethe
     * 鉴权相关js
     */
    angular.module('eolinker')
        .component('gpeditInsideAuth', {
            templateUrl: 'app/ui/content/gpedit/inside/content/auth/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope','Authority_CommonService'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope,Authority_CommonService) {
        var vm = this;
        vm.data = {
            menu: [{
                key: '无认证',
                id: 'none'
            }, {
                key: 'API Key',
                id: 'apiKey'
            }, {
                key: 'Basic认证',
                id: 'basic'
            }],
            deleteClientIDList: []
        }
        vm.ajaxRequest={
            strategyID: $state.params.strategyID
        };
        vm.ajaxResponse={};
        vm.service={
            authority:Authority_CommonService
        }
        vm.fun = {};
        vm.constant = {
            twoLengthArray: [{
                key: 'true',
                value: true
            }, {
                key: 'false',
                value: false
            }],
            algorithmArray: [{
                key: 'HS256',
                value: 'HS256'
            }, {
                key: 'HS384',
                value: 'HS384'
            },{
                key: 'HS512',
                value: 'HS512'
            }, {
                key: 'RS256',
                value: 'RS256'
            },{
                key: 'RS384',
                value: 'RS384'
            }, {
                key: 'RS512',
                value: 'RS512'
            },{
                key: 'ES256',
                value: 'ES256'
            }, {
                key: 'ES384',
                value: 'ES384'
            }, {
                key: 'ES512',
                value: 'ES512'
            }]
        }
        vm.component = {
            menuObject: {}
        }
        var data = {
            templateResponse: null
        }
        var assistantFun = {}
        assistantFun.getStatus = (function () {
            GatewayResource.Auth.Status({
                strategyID: vm.ajaxRequest.strategyID
            }).$promise.then(function (response) {
                vm.ajaxResponse.authAuthority = response;
                if (response.apiKeyStatus) {
                    vm.data.menuType = 'apiKey';
                } else if (response.basicAuthStatus) {
                    vm.data.menuType = 'basicAuth';
                } else if (response.jwtStatus) {
                    vm.data.menuType = 'jwt';
                } else if (response.oAuthStatus) {
                    vm.data.menuType = 'oauth2.0';
                }
            })
        })();
        vm.fun.last = function (status, arg) {
            var template = {
                apiKey: {
                    Apikey: '',
                    hideCredential: false,
                    remark: ''
                },
                basicAuth: {
                    username: '',
                    password: '',
                    hideCredential: false,
                    remark: ''
                },
                jwt: {
                    iss: '',
                    secret: '',
                    rsaPublicKey: '',
                    algorithm: 'HS256',
                    remark: ''
                }
            }
            template.uuidAssistantFun = function () {
                return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
            }
            template.uuid = function () {
                return (template.uuidAssistantFun() + template.uuidAssistantFun() + "-" + template.uuidAssistantFun() + "-" + template.uuidAssistantFun() + "-" + template.uuidAssistantFun() + "-" + template.uuidAssistantFun() + template.uuidAssistantFun() + template.uuidAssistantFun());
            }
            template.oAuth = {
                credentialID: template.uuid(),
                clientID: template.uuid(),
                clientSecret: '',
                redirectUri: '',
                remark: ''
            };
            if (arg.$last) {
                switch (status) {
                    case 'apiKey':
                        {
                            vm.ajaxResponse.authInfo.apiKeyList.push(template.apiKey);
                            break;
                        }
                    case 'basicAuth':
                        {
                            vm.ajaxResponse.authInfo.basicAuthList.push(template.basicAuth);
                            break;
                        }
                    case 'oauth':
                        {
                            vm.ajaxResponse.authInfo.oauth2CredentialList.push(template.oAuth);
                            break;
                        }
                    case 'jwt':
                        {
                            vm.ajaxResponse.authInfo.jwtCredentialList.push(template.jwt);
                            break;
                        }
                    case 'all':
                        {
                            vm.ajaxResponse.authInfo.apiKeyList.push(template.apiKey);
                            vm.ajaxResponse.authInfo.basicAuthList.push(template.basicAuth);
                            vm.ajaxResponse.authInfo.oauth2CredentialList.push(template.oAuth);
                            vm.ajaxResponse.authInfo.jwtCredentialList.push(template.jwt);
                            break;
                        }
                }
            }
        }
        vm.fun.delete = function (status, arg) {
            switch (status) {
                case 'apiKey':
                    {
                        vm.ajaxResponse.authInfo.apiKeyList.splice(arg.$index, 1);
                        break;
                    }
                case 'basicAuth':
                    {
                        vm.ajaxResponse.authInfo.basicAuthList.splice(arg.$index, 1);
                        break;
                    }
                case 'oauth':
                    {
                        if (arg.item.clientID) {
                            vm.data.deleteClientIDList.push(arg.item.clientID);
                        }
                        vm.ajaxResponse.authInfo.oauth2CredentialList.splice(arg.$index, 1);
                        break;
                    }
                case 'jwt':
                    {
                        vm.ajaxResponse.authInfo.jwtCredentialList.splice(arg.$index, 1);
                        break;
                    }
                case 'all':
                    {
                        vm.ajaxResponse.authInfo.apiKeyList.splice(vm.ajaxResponse.authInfo.apiKeyList.length - 1, 1);
                        vm.ajaxResponse.authInfo.basicAuthList.splice(vm.ajaxResponse.authInfo.basicAuthList.length - 1, 1);
                        vm.ajaxResponse.authInfo.oauth2CredentialList.splice(vm.ajaxResponse.authInfo.oauth2CredentialList.length - 1, 1);
                        vm.ajaxResponse.authInfo.jwtCredentialList.splice(vm.ajaxResponse.authInfo.jwtCredentialList.length - 1, 1);
                        break;
                    }
            }
        }
        vm.fun.edit = function () {
            var template = {
                request: {
                    strategyID: $state.params.strategyID,
                    strategyName: vm.ajaxResponse.authInfo.strategyName,
                    deleteClientIDList: vm.data.deleteClientIDList.join(',')
                },
                promise: null
            }
            template.request.jwtCredentialList = JSON.stringify(vm.ajaxResponse.authInfo.jwtCredentialList.slice(0, vm.ajaxResponse.authInfo.jwtCredentialList.length - 1), function (key, val) {
                switch (typeof (val)) {
                    case 'object':
                        {
                            return val;
                            break;
                        }
                    default:
                        {
                            if (/(iss)|(secret)|(rsaPublicKey)|(algorithm)|(remark)/.test(key)) {
                                return val;
                            }
                            break;
                        }
                }
            });
            template.request.apiKeyList = JSON.stringify(vm.ajaxResponse.authInfo.apiKeyList.slice(0, vm.ajaxResponse.authInfo.apiKeyList.length - 1), function (key, val) {
                switch (typeof (val)) {
                    case 'object':
                        {
                            return val;
                            break;
                        }
                    default:
                        {
                            if (/(Apikey)|(hideCredential)|(remark)/.test(key)) {
                                return val;
                            }
                            break;
                        }
                }
            });
            template.request.basicAuthList = JSON.stringify(vm.ajaxResponse.authInfo.basicAuthList.slice(0, vm.ajaxResponse.authInfo.basicAuthList.length - 1), function (key, val) {
                switch (typeof (val)) {
                    case 'object':
                        {
                            return val;
                            break;
                        }
                    default:
                        {
                            if (/(userName)|(hideCredential)|(password)|(remark)/.test(key)) {
                                return val;
                            }
                            break;
                        }
                }
            });
            template.request.oauth2CredentialList = JSON.stringify(vm.ajaxResponse.authInfo.oauth2CredentialList.slice(0, vm.ajaxResponse.authInfo.oauth2CredentialList.length - 1), function (key, val) {
                switch (typeof (val)) {
                    case 'object':
                        {
                            return val;
                            break;
                        }
                    default:
                        {
                            if (/(credentialID)|(clientID)|(clientSecret)|(redirectURI)|(remark)/.test(key)) {
                                return val;
                            }
                            break;
                        }
                }
            });
            template.request.basicUserName = vm.ajaxResponse.authInfo.basicUserName;
            template.promise = GatewayResource.Auth.Edit(template.request).$promise;
            template.promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            $rootScope.InfoModal('保存鉴权信息成功！', 'success');
                            vm.component.menuObject.show.status.writable = false;
                            vm.fun.delete('all');
                            break;
                        }
                }
            })
        }

        assistantFun.init = (function () {
            var template = {
                request: {
                    strategyID: vm.ajaxRequest.strategyID
                }
            }
            GatewayResource.Auth.Info(template.request).$promise.then(function (response) {
                vm.ajaxResponse.authInfo = response;
            })
        })();
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['鉴权方式', '策略']
            });
            vm.component.menuObject={
                show: {
                    status: {
                        writable: false
                    }
                },
                list:[{
                    type: 'btn',
                    class: 'btn-group-li pull-left ml15',
                    showVariable: 'writable',
                    showPoint: 'status',
                    authority:'edit',
                    btnList: [{
                        name: '编辑',
                        class: 'eo_theme_btn_success block-btn',
                        show: false,
                        fun: {
                            default: function () {
                                vm.component.menuObject.show.status.writable = true;
                                data.templateResponse=angular.copy(vm.ajaxResponse.authInfo);
                                vm.fun.last('all', {
                                    $last: true
                                });
                            }
                        }
                    }]
                },{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    showVariable: 'writable',
                    showPoint: 'status',
                    authority:'edit',
                    btnList: [{
                        name: '保存',
                        class: 'eo_theme_btn_success block-btn',
                        show: true,
                        fun: {
                            default: vm.fun.edit
                        }
                    }, {
                        name: '取消',
                        class: 'eo_theme_btn_default',
                        show: true,
                        fun: {
                            default: function () {
                                vm.component.menuObject.show.status.writable = false;
                                vm.fun.delete('all');
                                vm.ajaxResponse.authInfo=angular.copy(data.templateResponse);
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