(function () {
    'use strict';


    angular.module('goku')
        .component('gpeditGatewayComponent', {
            templateUrl: 'app/component/gateway/gpedit/index.html',
            controller: indexController,
            bindings: {}
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', '$filter', 'Communicate_CommonService'];

    function indexController($scope, GatewayResource, $state, $filter, Communicate_CommonService) {
        var vm = this;
        vm.data = {
            info: {
                gatewayUrl: '',
                model: null,
                itemStatus: 'hidden'
            },
            interaction: {
                response: {
                    query: null
                }
            },
            fun: {
                click: null
            }
        }
        var data = {
            storage: {},
            interaction: {
                request: {
                    gatewayAlias: $state.params.gatewayAlias
                }
            },
            fun: {
                init: null
            }
        }
        var service = {
            communicate: Communicate_CommonService
        }
        vm.gpeditQueryInit = function (attr) {
            return attr;
        };
        data.fun.init = (function () {
            service.communicate.fun.clear('GPEDIT_ID');
            service.communicate.fun.clear('GATEWAY_URL_SHOULD_BE_CONTACT');
            data.storage = JSON.parse(window.localStorage['GPEDIT_COMPONENT_TABLE'] || '{}');
            GatewayResource.Strategy.SimpleQuery({
                gatewayAlias: data.interaction.request.gatewayAlias
            }).$promise.then(function (response) {
                vm.data.interaction.response.query = response.strategyList || [];
                var template = {
                    cache: data.storage[data.interaction.request.gatewayAlias]
                }
                if (template.cache) {
                    for (var key = 0; key < vm.data.interaction.response.query.length; key++) {
                        var val = vm.data.interaction.response.query[key];
                        if (val.strategyID == template.cache.strategyID) {
                            vm.data.info.model = val;
                            service.communicate.fun.set('GPEDIT_ID', '/' + val.strategyID);
                            break;
                        }
                    }
                    vm.data.info.shouldContactGatewayUrl = template.cache.GATEWAY_URL_SHOULD_BE_CONTACT;
                    service.communicate.fun.set('GATEWAY_URL_SHOULD_BE_CONTACT', template.cache.GATEWAY_URL_SHOULD_BE_CONTACT);
                }
                vm.data.info.model = vm.data.info.model || {
                    strategyName: '请选择策略组'
                }
            })
        })();
        vm.data.fun.click = function (status, gpeditItem) { //下拉按钮单击功能函数
            switch (status) {
                case 'contact':
                    {
                        vm.data.info.shouldContactGatewayUrl = !vm.data.info.shouldContactGatewayUrl;
                        service.communicate.fun.set('GATEWAY_URL_SHOULD_BE_CONTACT', vm.data.info.shouldContactGatewayUrl);
                        break;
                    }
                case 'select':
                    {
                        gpeditItem = gpeditItem || {
                            strategyID: null,
                            strategyName: '请选择策略组'
                        }
                        gpeditItem.changed = true;
                        vm.data.info.itemStatus = 'hidden';
                        vm.data.info.model = gpeditItem;
                        service.communicate.fun.set('GPEDIT_ID', gpeditItem.strategyID ? ('/' + gpeditItem.strategyID) : '');
                        break;
                    }
            }
            data.storage[data.interaction.request.gatewayAlias] = {
                strategyID: vm.data.info.model.strategyID,
                GATEWAY_URL_SHOULD_BE_CONTACT: vm.data.info.shouldContactGatewayUrl
            };
            window.localStorage.setItem('GPEDIT_COMPONENT_TABLE', JSON.stringify(data.storage));
        }
    }
})();