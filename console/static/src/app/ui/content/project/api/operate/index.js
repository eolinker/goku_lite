(function () {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块api编辑相关js
     */
    angular.module('eolinker')
        .component('apiOperate', {
            templateUrl: 'app/ui/content/project/api/operate/index.html',
            controller: indexController
        })
        .filter('Filter_SlashSymbol', () => {
            return (inputStr) => {
                return (inputStr || '').replace(/\/{2,}/g, '/');
            }
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', 'GroupService'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, GroupService) {
        var vm = this;
        vm.data = {
            status:$state.params.status,
            isSpreedStaticResponse:true,
            requestMethod: false,
            apiGroup: null,
            requestMethodList: [{
                key: 'POST',
                value: 'post'
            }, {
                key: 'GET',
                value: 'get'
            }, {
                key: 'PUT',
                value: 'put'
            }, {
                key: 'DELETE',
                value: 'delete'
            }, {
                key: 'HEAD',
                value: 'head'
            }, {
                key: 'OPTIONS',
                value: 'options'
            }, {
                key: 'PATCH',
                value: 'patch'
            }]
        };
        vm.fun = {};
        vm.ajaxRequest = {
            projectID: $state.params.projectID,
            groupID: $state.params.groupID || -1,
            apiID: $state.params.apiID
        }
        vm.CONST = {
            PROTOCOL_ARR: [{
                key: 'HTTP',
                value: 'http'
            }, {
                key: 'HTTPS',
                value: 'https'
            }]
        }
        vm.ajaxResponse = {
            apiInfo: {
                apiName: '',
                requestMethod: '',
                requestURL: '/',
                balanceName: '',
                targetURL: '/',
                targetMethod: '-1',
                isFollow: true,
                timeout: '2000',
                retryCount: 0,
                protocol: 'http',
                responseDataType:$state.params.status==="add-link"?"json":"origin",
                linkApis:[],
                apiType:$state.params.status==="add-link"?1:0
            },
            balanceList:[]
        }
        vm.component = {
            selectMultistageCommonComponentObject: {
                new: {},
                original: {}
            },
            menuObject: {
                list: []
            },
            balanceAutoCompleteObj:{required:true,pattern:'[\\w\\._\\/\\-\\:]+'},
            apiLinkStepObj:{
                CONST:{
                    PROTOCOL_ARR:vm.CONST.PROTOCOL_ARR
                }
            }
        };
        
        var privateFun = {};
        vm.fun.back = function () {
            $state.go('home.project.api.default', {
                'groupID': $state.params.groupID
            });
        }
        privateFun.parseLinkApis=()=>{
            if(vm.ajaxResponse.apiInfo.apiType===0)return true;
            let tmpLinkApis=angular.copy(vm.ajaxResponse.apiInfo.linkApis);
            for(let key in tmpLinkApis){
                let val=tmpLinkApis[key];
                if(!/^[1-9]\d*$/.test(val.timeout)){
                    return false;
                }
                val.timeout=parseInt(val.timeout);
                if(!/(^[1-9]\d*$)|(^0$)/.test(val.retry)){
                    return false;
                }
                val.retry=parseInt(val.retry);
                let tmpList=[];
                val.blackList.map((childItem)=>{
                    if(childItem.ip){
                        tmpList.push(childItem.ip);
                    }
                })
                val.blackList=tmpList;
                tmpList=[];
                val.whiteList.map((childItem)=>{
                    if(childItem.ip){
                        tmpList.push(childItem.ip);
                    }
                })
                val.whiteList=tmpList;
                val.delete=val.delete.filter((childItem)=>{
                    if(childItem.origin)return childItem;
                });
                tmpList=[];
                for(let childKey in val.move){
                    let childItem =val.move[childKey];
                    if(childItem.target){
                        if(childItem.origin){
                            tmpList.push(childItem)
                        }else return;
                    }
                }
                val.move=tmpList;
                tmpList=[];
                for(let childKey in val.rename){
                    let childItem =val.rename[childKey];
                    if(childItem.target){
                        if(childItem.origin){
                            childItem.target=(childItem.prefixStr||"")+childItem.target;
                            tmpList.push(childItem)
                        }else return;
                    }
                }
                val.rename=tmpList;
            }
            return JSON.stringify(tmpLinkApis,(tmpInputKey,tmpInputItem)=>{
                if (/(prefixStr)|(prefixStr)|($$hashKey)/.test(tmpInputKey)) {
                    return undefined;
                }
                return tmpInputItem;
            })
        }
        privateFun.confirm = function () {
            vm.data.requestMethod = false;
            var tmpOutput = {
                apiType:vm.ajaxResponse.apiInfo.apiType,
                projectID: vm.ajaxRequest.projectID,
                groupID: vm.component.selectMultistageCommonComponentObject.new.value==-1?0:vm.component.selectMultistageCommonComponentObject.new.value,
                apiName: vm.ajaxResponse.apiInfo.apiName,
                requestMethod: [],
                requestURL: vm.ajaxResponse.apiInfo.requestURL,
                balanceName: vm.ajaxResponse.apiInfo.balanceName,
                targetURL: vm.ajaxResponse.apiInfo.targetURL,
                targetMethod: vm.ajaxResponse.apiInfo.targetMethod,
                responseDataType: vm.ajaxResponse.apiInfo.responseDataType,
                timeout: vm.ajaxResponse.apiInfo.timeout,
                retryCount: vm.ajaxResponse.apiInfo.retryCount,
                protocol:vm.ajaxResponse.apiInfo.protocol,
                staticResponse:vm.ajaxResponse.apiInfo.staticResponse
            }
            for (var key in vm.data.requestMethodList) {
                if (vm.data.requestMethodList[key].checkbox) {
                    tmpOutput.requestMethod.push(vm.data.requestMethodList[key].value);
                }
            }
            if (tmpOutput.requestMethod.length > 0) {
                tmpOutput.requestMethod = tmpOutput.requestMethod.join(',');
            } else {
                vm.data.requestMethod = true;
            }
            switch (vm.data.status) {
                case 'edit': {
                    tmpOutput.apiID = vm.ajaxRequest.apiID
                    break;
                }
            }
            if (vm.ajaxResponse.apiInfo.targetMethod == '-1') {
                tmpOutput.isFollow = true;
                delete tmpOutput.targetMethod;
            }
            return tmpOutput;
        }
        vm.fun.load = function (arg) {
            $scope.$emit('$TransferStation', {
                state: '$Init_LoadingCommonComponent',
                data: arg
            });
        }
        vm.fun.requestProcessing = function (arg) {
            var tmpPromise = null,tmpLinkApis=privateFun.parseLinkApis(),tmpAjaxRequest=privateFun.confirm();
            vm.data.submitted = true;
            if ($scope.ConfirmForm.$valid && !vm.data.requestMethod&&tmpLinkApis) {
                tmpPromise = privateFun.edit({
                    request: tmpLinkApis!==true?Object.assign({},tmpAjaxRequest,{
                        linkApis:tmpLinkApis
                    }):tmpAjaxRequest
                });
            } else {
                $rootScope.InfoModal('API编辑失败，请检查信息是否填写完整！', 'error');

            }
            return tmpPromise;
        }
        privateFun.edit = function (arg) {
            var tmpPromise = null;
            if (vm.data.status == 'edit') {
                tmpPromise = GatewayResource.Api.Edit(arg.request).$promise;
                tmpPromise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS: {
                            vm.fun.back();
                            $rootScope.InfoModal('API修改成功', 'success');
                            break;
                        }
                        case '190005': {
                            vm.data.requestMethod = true;
                            $scope.ConfirmForm.requestURL.$invalid = true;
                            break;
                        }
                    }
                })
            } else {
                tmpPromise = GatewayResource.Api.Add(arg.request).$promise;
                tmpPromise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS: {
                            vm.fun.back();
                            $rootScope.InfoModal('API添加成功', 'success');
                            break;
                        }
                        case '190005': {
                            vm.data.requestMethod = true;
                            $scope.ConfirmForm.requestURL.$invalid = true;
                            break;
                        }
                    }
                })
            }
            return tmpPromise;
        }
        privateFun.initApi = function () {
            var tmpAjaxRequest = {
                projectID: vm.ajaxRequest.projectID,
                apiID: vm.ajaxRequest.apiID
            }
            switch (vm.data.status) {
                case 'edit': {
                    GatewayResource.Api.Info(tmpAjaxRequest).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                vm.ajaxResponse.apiInfo = response.apiInfo;
                                let tmpRequestMethod = response.apiInfo.requestMethod.split(',');
                                vm.ajaxResponse.apiInfo.targetMethod = response.apiInfo.isFollow ? '-1' : vm.ajaxResponse.apiInfo.targetMethod;
                                for (var key in tmpRequestMethod) {
                                    var val = tmpRequestMethod[key];
                                    switch (val.toLowerCase()) {
                                        case 'post': {
                                            vm.data.requestMethodList[0].checkbox = true;
                                            break;
                                        }
                                        case 'get': {
                                            vm.data.requestMethodList[1].checkbox = true;
                                            break;
                                        }
                                        case 'put': {
                                            vm.data.requestMethodList[2].checkbox = true;
                                            break;
                                        }
                                        case 'delete': {
                                            vm.data.requestMethodList[3].checkbox = true;
                                            break;
                                        }
                                        case 'head': {
                                            vm.data.requestMethodList[4].checkbox = true;
                                            break;
                                        }
                                        case 'options': {
                                            vm.data.requestMethodList[5].checkbox = true;
                                            break;
                                        }
                                        case 'batch': {
                                            vm.data.requestMethodList[6].checkbox = true;
                                            break;
                                        }
                                    }
                                }
                                if(response.apiInfo.apiType===1){
                                    vm.ajaxResponse.apiInfo.linkApis.map((val)=>{
                                        let tmpList=[];
                                        val.timeout=(val.timeout||2000).toString();
                                        val.blackList.map((childItem)=>{
                                            tmpList.push({
                                                ip:childItem
                                            })
                                        })
                                        val.blackList=tmpList.concat([{ip:""}]);
                                        tmpList=[];
                                        val.whiteList.map((childItem)=>{
                                            tmpList.push({
                                                ip:childItem
                                            })
                                        })
                                        val.whiteList=tmpList.concat([{ip:""}]);
                                        tmpList=[];
                                        val.delete.map((childItem)=>{
                                            tmpList.push(childItem)
                                        })
                                        val.delete=tmpList.concat([{origin:""}]);

                                        tmpList=[];
                                        val.move.map((childItem)=>{
                                            tmpList.push(childItem)
                                        })
                                        val.move=tmpList.concat([{origin:"",target:""}]);

                                        tmpList=[];
                                        val.rename.map((childItem)=>{
                                            let tmpArr=childItem.target.split('.');
                                            childItem.prefixStr=tmpArr.slice(0,tmpArr.length-1).join('.');
                                            if(childItem.prefixStr)childItem.prefixStr+=".";
                                            childItem.target=tmpArr[tmpArr.length-1];
                                            tmpList.push(childItem)
                                        })
                                        val.rename=tmpList.concat([{origin:"",target:""}]);
                                    })
                                }
                                break;
                            }
                        }
                    });
                    break;
                }
            }
        }
        privateFun.getDefaultBalance=()=>{
            GatewayResource.Balance.SimpleQuery().$promise.then((response)=>{
                vm.ajaxResponse.balanceList=response.balanceNames;
            })
        }
        vm.fun.init = (function () {
            let tmpStaticGroupArr=[{
                groupName:"未分组",
                groupID:-1
            }]
            var tmpCache = GroupService.get();
            if (tmpCache) {
                vm.data.apiGroup = tmpStaticGroupArr.concat(tmpCache);
                privateFun.initApi();
            } else {
                $rootScope.global.ajax.Query_Group.$promise.finally(function () {
                    vm.data.apiGroup = tmpStaticGroupArr.concat(GroupService.get());
                    privateFun.initApi();
                })
            }

        })()
        vm.$onInit = function () {
            privateFun.getDefaultBalance();
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