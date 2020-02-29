(function () {
    /**
     * @name 项目设置
     * @author 广州银云信息科技有限公司
     */
    'use strict';
    angular.module('eolinker')
        .component('gpeditSetting', {
            templateUrl: 'app/ui/content/gpedit/inside/content/setting/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', 'Authority_CommonService'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, Authority_CommonService) {
        var vm = this;
        vm.data={};
        vm.ajaxResponse={};
        vm.ajaxRequest={
            strategyID:$state.params.strategyID
        };
        vm.service = {
            authority: Authority_CommonService
        };
        vm.fun = {};
        vm.fun.oprGpeditStatus = function (inputEnableStatus) {
            let tmpPromise,tmpAjaxRequest={
                strategyIDList:vm.ajaxRequest.strategyID
            },tmpTarget=inputEnableStatus?"Stop":"Start";
            tmpPromise=GatewayResource.Strategy[tmpTarget](tmpAjaxRequest).$promise;
            tmpPromise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.ajaxResponse.strategyInfo.enableStatus=inputEnableStatus?0:1;
                            $rootScope.InfoModal(`${inputEnableStatus?"停用":"启用"}成功`, 'success');
                            break;
                        }
                }
            });
            return tmpPromise;
        }
        vm.fun.confirm = function () {
            let tmpAjaxRequest={
                strategyID:vm.ajaxRequest.strategyID,
                strategyName:vm.ajaxResponse.strategyInfo.strategyName,
                groupID:vm.ajaxResponse.strategyInfo.groupID
            },tmpPromise;
            if ($scope.BasicForm.$valid) {
                tmpPromise = GatewayResource.Strategy.Edit(tmpAjaxRequest).$promise;
                tmpPromise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                $rootScope.InfoModal('编辑成功', 'success');
                                break;
                            }
                    }
                })
            }
            return tmpPromise;
        }
        vm.fun.init = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['策略管理']
            });
            $rootScope.global.ajax.Info_Strategy = GatewayResource.Strategy.Info({
                strategyType:$state.params.groupType==="open"?1:0,
                strategyID:vm.ajaxRequest.strategyID
            });
            $rootScope.global.ajax.Info_Strategy.$promise.then(function (response) {
                vm.ajaxResponse.strategyInfo = response.strategyInfo || {};
            })
            return $rootScope.global.ajax.Info_Strategy.$promise;
        }
    }
})();