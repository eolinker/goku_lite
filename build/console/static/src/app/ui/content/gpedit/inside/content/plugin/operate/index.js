(function () {
    'use strict';
    /**
     * 测试导入增强插件
     * @param [object] input 输入信息
     * @param [object] output flag：标识是否已经更改,visible：是否可视
     */
    angular.module('eolinker')
        .component('gpeditInsidePluginOperate', {
            templateUrl: 'app/ui/content/gpedit/inside/content/plugin/operate/index.html',
            controller: indexController,
            bindings:{
                oprTarget:'@',
                accept:"<",
                trigger:'=',
                status:"@",
                pluginName:'@',
                chineseName:'@'
            }
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope) {
        var vm = this;
        vm.data = {};
        vm.ajaxRequest={
            strategyID: $state.params.strategyID
        }
        vm.ajaxResponse={
            bindInfo: {
                apiID: $state.params.apiID
            }
        }
        vm.fun={};
        vm.component = {
            menuObject: {
                list: []
            }
        };
        let privateFun = {},
        data={},
        _Resource;
        privateFun.getPluginList = function () {
            if (vm.data.status == 'edit') return;
            switch(vm.oprTarget){
                case 'api':{
                    $rootScope.global.ajax.TagQueryPlugin = GatewayResource.PluginApi.QueryByStrategy({
                        strategyID: vm.ajaxRequest.strategyID
                    });
                    break;
                }
                case 'gpedit':{
                    $rootScope.global.ajax.TagQueryPlugin = GatewayResource.Plugin.TagQuery({
                        pluginType: 1
                    });
                    break;
                }
            }
            
            $rootScope.global.ajax.TagQueryPlugin.$promise.then(function (response) {
                vm.ajaxResponse.pluginList = response.pluginList || [{pluginName:'尚未开启任何'+(vm.oprTarget=='api'?'API':'策略')+'插件',key:-1}];
                if(vm.ajaxResponse.pluginList.length==0){
                    vm.ajaxResponse.pluginList=[{pluginName:'尚未开启任何'+(vm.oprTarget=='api'?'API':'策略')+'插件',key:-1}];
                }
                vm.ajaxResponse.bindInfo.pluginName = vm.ajaxResponse.pluginList[0].pluginName;
                for (var key in response.pluginList) {
                    switch (response.pluginList[key].pluginName) {
                        case 'goku-response_headers':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "matchStatusCode": "200", \n' +
                                '    "responseHeaders": {\n' +
                                '       "Gateway":"GoKu"\n' +
                                '   }\n' +
                                '}';
                                break;
                            }
                        case 'goku-cors':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "allowOrigin": "", \n' +
                                '    "allowMethods": "GET,POST", \n' +
                                '    "allowHeaders": "*", \n' +
                                '    "allowCredentials": "true", \n' +
                                '    "exposeHeaders": "" \n' +
                                '}';
                                break;
                            }
                        case 'goku-replay_attack_defender':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "timestampTTL": 600, \n' +
                                '    "replayAttackToken": ""\n' +
                                '}';
                                break;
                            }
                        case 'goku-proxy_caching':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "responseCodes": "",\n' +
                                '    "requestMethods": "", \n' +
                                '    "contentTypes": "", \n' +
                                '    "cacheTTL": 300 \n' +
                                '}';
                                break;
                            }
                        case 'goku-jwt_auth':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "signatureIsBase64": false, \n' +
                                '    "claimsToVerify": ["exp","nbf"], \n' +
                                '    "runOnPreflight": true,\n' +
                                '    "jwtCredentials": [{\n' +
                                '       "iss": "",\n' +
                                '       "secret": "",\n' +
                                '       "rsaPublicKey": "",\n' +
                                '       "algorithm": "HS256",\n' +
                                '       "remark": ""\n' +
                                '    }],\n' +
                                '    "hideCredentials": false\n' +
                                '}';
                                break;
                            }
                        case 'goku-circuit_breaker':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "failurePercent": 0.5,\n' +
                                '    "monitorPeriod": 20, \n' +
                                '    "minimumRequests": 20, \n' +
                                '    "breakPeriod": 20, \n' +
                                '    "successCounts": 10, \n' +
                                '    "matchStatusCodes": "",\n' +
                                '    "statusCode": 200,\n' +
                                '    "headers": {},\n' +
                                '    "body": ""\n' +
                                '}';
                                break;
                            }
                        case 'goku-extra_params':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "params": [{\n' +
                                '       "paramName": "", \n' +
                                '       "paramPosition": "",\n' +
                                '       "paramValue": "", \n' +
                                '       "paramConflictSolution": "convert"\n'+
                                '    }]\n' +
                                '}';
                                break;
                            }
                        case 'goku-oauth2_auth':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "scopes": [],\n' +
                                '    "mandatoryScope": false, \n' +
                                '    "tokenExpiration": 7200,\n' +
                                '    "enableAuthorizationCode": true,\n' +
                                '    "enableImplicitGrant": false,\n' +
                                '    "enableClientCredentials": false,\n' +
                                '    "hideCredentials": false,\n' +
                                '    "acceptHttpIfAlreadyTerminated": true,\n' +
                                '    "refreshTokenTTL": 1209600,\n' +
                                '    "oauth2CredentialList": []\n' +
                                '}';
                                break;
                            }
                        case 'goku-apikey_auth':
                            {
                                response.pluginList[key].pluginConfig = '[{\n' +
                                '    "Apikey": "",\n' +
                                '    "hideCredential": false\n' +
                                '}]';
                                break;
                            }
                        case 'goku-basic_auth':
                            {
                                response.pluginList[key].pluginConfig = '[{\n' +
                                '    "userName": "",\n' +
                                '    "password": "", \n' +
                                '    "hideCredential": false\n' +
                                '}]';
                                break;
                            }
                        case 'goku-ip_restriction':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "ipListType":"black",\n' +
                                '    "ipWhiteList":["127.0.0.1"],\n' +
                                '    "ipBlackList":["127.0.0.1"]\n' +
                                '}';
                                break;
                            }
                        case 'goku-rate_limiting':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "second": 10, \n' +
                                '    "minute": 50, \n' +
                                '    "hour": 100, \n' +
                                '    "day": 1000, \n' +
                                '    "hideClientHeader": false \n' +
                                '}';
                                break;
                            }
                        case 'goku-request_size_limiting':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "allowedPayLoadSize": 655\n' +
                                '}';
                                break;
                            }
                        case 'goku-service_downgrade':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "matchStatusCodes": "",\n' +
                                '    "statusCode": 200,\n' +
                                '    "headers": {}, \n' +
                                '    "body": ""\n' +
                                '}';
                                break;
                            }
                        case 'goku-default_response':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "statusCode": 200, \n' +
                                '    "headers": {}, \n' +
                                '    "body": "" \n' +
                                '}';
                                break;
                            }
                        case 'goku-params_check':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "params": [{\n' +
                                '       "paramName":"",\n' +
                                '       "paramPosition":"",\n' +
                                '       "regular":""\n' +
                                '    }]\n' +
                                '}';
                                break;
                            }
                        case 'goku-params_transformer':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "params": [{\n' +
                                '       "paramName": "", \n' +
                                '       "paramPosition": "",\n' +
                                '       "proxyParamName": "", \n' +
                                '       "proxyParamPosition": "",\n' +
                                '       "required": true\n' +
                                '    }],\n' +
                                '    "removeAfterTransformed": true\n' +
                                '}';
                                break;
                            }
                        case 'goku-http_log':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "logName":"http",\n' +
                                '    "fileDir":"",\n' +
                                '    "recordPeriod":"hour"\n' +
                                '}';
                                break;
                            }
                        case 'goku-data_format_transformer':
                            {
                                response.pluginList[key].pluginConfig = '{\n' +
                                '    "enableXMLToJSON":true,\n' +
                                '    "enableJSONToXML":true,\n' +
                                '    "continueIfTransformFailed":true,\n' +
                                '    "XMLRootTag":"XML"\n' +
                                '}';
                                break;
                            }

                    }
                }
            })
        };
        privateFun.getBindInfo = function () {
            var template = {
                request: {
                    pluginName: vm.ajaxResponse.bindInfo.pluginName,
                    strategyID: vm.ajaxRequest.strategyID
                },
                pluginConfig: ''
            }
            switch (vm.oprTarget) {
                case 'api':
                    {
                        template.request.apiID = vm.ajaxResponse.bindInfo.apiID;
                        break;
                    }
            }
            _Resource.Info(template.request).$promise.then(function (response) {
                if (vm.data.status == 'edit' && vm.oprTarget == 'api') {
                    vm.ajaxResponse.apiList = [{
                        apiID: vm.ajaxResponse.bindInfo.apiID,
                        apiName: response.apiName,
                        requestURL: response.requestURL
                    }];
                }
                if (response.strategyPluginConfig || response.apiPluginConfig) {
                    try {
                        template.pluginConfig = (response.strategyPluginConfig || response.apiPluginConfig);
                        template.pluginConfig = JSON.stringify(JSON.parse(template.pluginConfig.replace(/(\\)/g, 'author:eoLinkerFormatPosition')), null, 4).replace(/author:eoLinkerFormatPosition/g, '\\');
                    } catch (e) {
                        console.log(e);
                    }
                    $scope.$broadcast('$Maunal_AceEditorAms', template.pluginConfig);
                } else {
                    for (var key in vm.ajaxResponse.pluginList) {
                        if (vm.ajaxResponse.pluginList[key].pluginName == vm.ajaxResponse.bindInfo.pluginName) {
                            $scope.$broadcast('$Maunal_AceEditorAms', vm.ajaxResponse.pluginList[key].pluginConfig);
                            break;
                        }
                    }
                }
            });
        }
        vm.fun.reset = function () {
            privateFun.getBindInfo();
        }
        vm.fun.init = function () {
            if ($rootScope.global.ajax.TagQueryPlugin) {
                $rootScope.global.ajax.TagQueryPlugin.$promise.finally(function(){
                    privateFun.getBindInfo();
                });
            } else {
                privateFun.getBindInfo();
            }
        };
        privateFun.back = function () {
            var template = {
                uri: 'home.gpedit.inside.plugin.' + vm.oprTarget
            }
            $state.go(template.uri);
        }
        privateFun.confirm = function () {
            var template = {
                output: {
                    pluginName: vm.ajaxResponse.bindInfo.pluginName,
                    pluginConfig: vm.ajaxResponse.bindInfo.pluginConfig,
                    strategyID: vm.ajaxRequest.strategyID
                }
            }
            try {
                vm.ajaxResponse.bindInfo.pluginConfig = JSON.stringify(JSON.parse(vm.ajaxResponse.bindInfo.pluginConfig.replace(/(\\)/g, 'author:eoLinkerFormatPosition')), null, 4).replace(/author:eoLinkerFormatPosition/g, '\\');
                $scope.$broadcast('$Maunal_AceEditorAms', vm.ajaxResponse.bindInfo.pluginConfig);
            } catch (e) {
                console.log(e)
            }
            switch (vm.oprTarget) {
                case 'api':
                    {
                        template.output.apiID = vm.ajaxResponse.bindInfo.apiID;
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
                promise: null
            }
            if ($scope.ConfirmForm.$valid && vm.ajaxResponse.bindInfo.pluginConfig) {
                template.request=privateFun.confirm();
                template.promise = privateFun.edit({
                    request: template.request
                });
            } else {
                vm.data.submited = true;
                $rootScope.InfoModal('插件编辑失败，请检查信息是否填写配置信息！', 'error');
            }
            return template.promise;
        }
        privateFun.edit = function (arg) {
            if (vm.data.status == 'edit') {
                data.EditResource = _Resource.Edit(arg.request);
                data.EditResource.$promise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                if(vm.oprTarget==="gpedit")privateFun.back();
                                $rootScope.InfoModal('修改成功', 'success');
                                vm.trigger="saveSuccess";
                                break;
                            }
                    }
                })
            } else {
                data.EditResource = _Resource.Add(arg.request);
                data.EditResource.$promise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                if(vm.oprTarget==="gpedit")privateFun.back();
                                $rootScope.InfoModal('添加插件成功', 'success');
                                vm.trigger="saveSuccess";
                                break;
                            }
                    }
                })
            }
            return data.EditResource.$promise;
        }
        vm.$onInit = function () {
            vm.data.status=vm.status||$state.params.status;
            vm.ajaxResponse.bindInfo.pluginName=vm.pluginName||$state.params.pluginName || '';
            vm.ajaxResponse.pluginList=[{
                pluginName: vm.ajaxResponse.bindInfo.pluginName,
                chineseName:vm.chineseName||$state.params.chineseName || ''
            }]
            vm.component.menuObject={
                list:[{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    btnList: [{
                        name: '返回列表',
                        icon: 'chexiao',
                        fun: {
                            default: privateFun.back
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
            if(vm.oprTarget==="api"){
                $rootScope.global.$watch.push($scope.$watch('$ctrl.accept', ()=>{
                    if(!vm.accept)return;
                    switch(vm.accept){
                        case "save":{
                            if(data.EditResource){
                                data.EditResource.$cancelRequest();
                            }
                            vm.fun.requestProcessing();
                            break;
                        }
                    }
                }, true));
                _Resource=GatewayResource.PluginApi;
            }else{
                _Resource=GatewayResource.PluginStrategy;
            }
            privateFun.getPluginList();
        }
        vm.$onDestroy=()=>{
            if($rootScope.global.ajax.TagQueryPlugin){
                $rootScope.global.ajax.TagQueryPlugin.$cancelRequest();
                delete $rootScope.global.ajax.TagQueryPlugin;
            }
        }
    }
})();