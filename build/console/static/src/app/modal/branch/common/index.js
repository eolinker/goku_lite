(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 公用弹窗controller js
     */
    angular.module('eolinker.modal')

        .directive('eoCommonModal', [function () {
            return {
                restrict: 'AE',
                templateUrl: 'app/modal/branch/common/index.html'
            }
        }])

        .controller('Common_LoginModalCtrl', Common_LoginModalCtrl)

        .controller('Common_SingleInputModalCtrl', Common_SingleInputModalCtrl)

        .controller('InfoModalCtrl', InfoModalCtrl)

        .controller('EnsureModalCtrl', EnsureModalCtrl)

        .controller('GroupModalCtrl', GroupModalCtrl)

        .controller('TableModalCtrl', TableModalCtrl)

        .controller('ImportModalCtrl', ImportModalCtrl)

        .controller('SelectVisualGroupModalCtrl', SelectVisualGroupModalCtrl)

        .controller('SingleSelectModalCtrl', SingleSelectModalCtrl)
        .controller('CommonChangePasswordModalCtrl', CommonChangePasswordModalCtrl)
        .controller('ExportModalCtrl', ExportModalCtrl)
        .controller('MixInputModalCtrl', MixInputModalCtrl)

    MixInputModalCtrl.$inject = ['$scope', '$uibModalInstance', '$rootScope', 'CODE', 'uibDateParser', 'input'];

    function MixInputModalCtrl($scope, $uibModalInstance, $rootScope, CODE, uibDateParser, input) {
        $scope.input = input;
        $scope.fun = {};
        $scope.data = {
            submitted: false,
            admin: '',
            datePickerObject: {
                request: {
                    startTime: input.datePickerObject ? (input.datePickerObject.startTime ? new Date(input.datePickerObject.startTime) : '') : '',
                    endTime: input.datePickerObject ? (input.datePickerObject.endTime ? new Date(input.datePickerObject.endTime) : '') : ''
                },
                show: false
            }
        }
        $scope.component = {
            selectPersonCommonComponentObject: {}
        }
        var data = {
                singleTextObject: angular.copy(input.singleTextObject)
            },
            fun = {};
        $scope.fun.showSearchList = function ($event) {
            $event.stopPropagation();
            $scope.data.showSearchList = !$scope.data.showSearchList;
        }
        $scope.fun.searchMemberList = function (value, $event) {
            $event.stopPropagation();
            $scope.data.showSearchList = true;
            $scope.input.singleTextObject.selectOptions = [];
            data.singleTextObject.selectOptions.map(function (val, key) {
                if (fun.filterMemberList(val, value)) {
                    $scope.input.singleTextObject.selectOptions.push(val);
                }
            })
        }
        $scope.fun.datePickerSelect = function (arg) {
            if (!$scope.data.datePickerObject.request.startTime) {
                $rootScope.InfoModal('请选择' + $scope.input.datePickerObject.title + '开始日期', 'error');
            } else if (!$scope.data.datePickerObject.request.endTime) {
                $rootScope.InfoModal('请选择' + $scope.input.datePickerObject.title + '结束日期', 'error');
            } else if (input.datePickerObject.notSameDay && $scope.data.datePickerObject.request.startTime.getTime() == $scope.data.datePickerObject.request.endTime.getTime()) {
                $rootScope.InfoModal('开始和结束日期不能是同一天，请重新选择!', 'error');
                $scope.data.datePickerObject.request.endTime = "";
            } else {
                if (arg) arg.$event.stopPropagation();
                $scope.data.datePickerObject.show = false;
            }
        }
        fun.filterMemberList = function (arg, value) {
            if ((arg.inviteCall || '').indexOf(value) > -1 || (arg.memberNickName || '').indexOf(value) > -1 || (arg.userNickName || '').indexOf(value) > -1) {
                return true;
            } else {
                return false;
            }
        }
        $scope.fun.confirm = function (inputSwitch) {
            var template = {
                request: angular.copy(input.request || {}),
                promise: null
            }
            $scope.data.submitted = true;
            if ($scope.ConfirmForm.$valid) {
                input.textArray.map(function (val, key) {
                    template.request[val.key] = val.value;
                })
                if (input.datePickerObject) {
                    if (!$scope.data.datePickerObject.request.startTime) {
                        $rootScope.InfoModal('请选择' + $scope.input.datePickerObject.title + '开始日期', 'error');
                        return;
                    } else if (!$scope.data.datePickerObject.request.endTime) {
                        $rootScope.InfoModal('请选择' + $scope.input.datePickerObject.title + '结束日期', 'error');
                        return;
                    } else if (input.datePickerObject.notSameDay && $scope.data.datePickerObject.request.startTime.getTime() == $scope.data.datePickerObject.request.endTime.getTime()) {
                        $rootScope.InfoModal('开始和结束日期不能是同一天，请重新选择!', 'error');
                        $scope.data.datePickerObject.request.endTime = "";
                        return;
                    } else {
                        template.request.startTime = uibDateParser.filter($scope.data.datePickerObject.request.startTime, 'yyyy-M!-dd');
                        template.request.endTime = uibDateParser.filter($scope.data.datePickerObject.request.endTime, 'yyyy-M!-dd');
                    }
                }
                if (input.singleTextObject && $scope.component.selectPersonCommonComponentObject.value) {
                    template.request[input.singleTextObject.key] = $scope.component.selectPersonCommonComponentObject.value;
                }
                if (input.ensure) {
                    $rootScope.EnsureModal(input.ensureInfo.title, input.ensureInfo.necessity, input.ensureInfo.info, input.ensureInfo.input || {}, function (callback) {
                        if (callback) {
                            if (!input.resource) {
                                $uibModalInstance.close(Object.assign({}, template.request));
                                return;
                            }
                            template.promise = input.resource(template.request).$promise;
                            template.promise.then(function (response) {
                                switch (response.statusCode) {
                                    case CODE.COMMON.SUCCESS: {
                                        $uibModalInstance.close(Object.assign({}, template.request, response));
                                        break;
                                    }
                                }
                            });
                        }
                    })
                } else {
                    if(input.promiseFun){
                        template.promise=input.promiseFun(Object.assign({}, template.request));
                        if(!template.promise)return;
                    }else  if (input.resource) {
                        let tmpAjaxRequest=template.request;
                        if(inputSwitch==="tmp"){
                            tmpAjaxRequest=Object.assign({},tmpAjaxRequest,input.tmpBtnObj.ajaxRequest||{})
                        }
                        template.promise = input.resource(tmpAjaxRequest).$promise;
                    }else {
                        $uibModalInstance.close(Object.assign({}, template.request));
                        return;
                    }
                    template.promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                $uibModalInstance.close(Object.assign({}, template.request, response));
                                break;
                            }
                        }
                    });
                }

            }
            return template.promise;
        };

        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
    }

    ExportModalCtrl.$inject = ['$scope', '$uibModalInstance', 'CODE', '$rootScope', 'input', '$state'];

    function ExportModalCtrl($scope, $uibModalInstance, CODE, $rootScope, input, $state) {
        $scope.info = {
            projectHashKey: $state.params.projectHashKey
        }
        $scope.data = {
            exportType: 0,
            granularityList: [{
                name: '天',
                active: 0
            }, {
                name: '小时',
                active: 1
            }],
            granularity: 0
        }
        $scope.input = input;
        $scope.fun = {};
        var fun = {};
        $scope.fun.changeMenu = function (mark, arg) {
            switch (mark) {
                case 'granularity': {
                    $scope.data.granularity = arg.item.active;
                    break;
                }
            }
        }
        fun.response = function (arg) {
            switch (arg.response.statusCode) {
                case CODE.COMMON.SUCCESS: {
                    $scope.$broadcast('$DumpDirective_Click_' + arg.switch.toString(), {
                        response: arg.response,
                        fileName: $scope.input.fileName,
                        window: arg.window
                    });
                    $uibModalInstance.close(true);
                    break;
                }
                default: {
                    if (arg.window) arg.window.close();
                    break;
                }
            }
        }
        $scope.fun.dumpDirective = function (arg) {
            var template = {
                promise: null,
                request: input.request,
                window: null
            }
            template.request.granularity = $scope.data.granularity;
            template.request.fileType = 'excel';
            template.promise = input.resource.Download(template.request).$promise;
            template.promise.then(function (response) {
                fun.response({
                    response: response,
                    switch: 'export-by-one-key',
                    window: template.window
                });
            })
            return template.promise;
        }
        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
    }
    CommonChangePasswordModalCtrl.$inject = ['$scope', '$uibModalInstance', 'input'];

    function CommonChangePasswordModalCtrl($scope, $uibModalInstance, input) {
        $scope.data = {
            input: {},
            fun: {
                cancel: null,
                confirm: null
            }
        }
        var data = {
            fun: {
                init: null
            }
        }
        $scope.data.fun.confirm = function () {
            if ($scope.editForm.$invalid || $scope.data.input.inputObject.confirmNewPassword != $scope.data.input.inputObject.key) return;
            $uibModalInstance.close({
                key: $scope.data.input.inputObject.key
            });
        };
        $scope.data.fun.cancel = function () {
            //$uibModalInstance.dismiss(false);
            $uibModalInstance.close(false);
        };
        data.fun.init = (function () {
            angular.copy(input, $scope.data.input);
        })()

    }

    SingleSelectModalCtrl.$inject = ['$scope', '$uibModalInstance', 'input'];

    function SingleSelectModalCtrl($scope, $uibModalInstance, input) {
        $scope.data = {
            title: input.title,
            query: input.query,
            position: input.position
        }
        $scope.output = {
            $index: '0'
        }
        $scope.ok = function () {
            $uibModalInstance.close($scope.output);
        };

        $scope.cancel = function () {
            //$uibModalInstance.dismiss(false);
            $uibModalInstance.close(false);
        };
    }

    SelectVisualGroupModalCtrl.$inject = ['$scope', '$uibModalInstance', '$state', 'GroupService', 'Group_MultistageService', 'input'];

    function SelectVisualGroupModalCtrl($scope, $uibModalInstance, $state, GroupService, Group_MultistageService, input) {
        $scope.title = input.title;
        $scope.secondTitle = input.secondTitle || "分组名称";
        $scope.modalType = input.modalType;
        $scope.data = {
            initialGroupData: [],
            apiGroup: null
        };
        $scope.interaction = {
            request: {}
        };
        $scope.fun = {};
        var fun = {},
            data = {
                isBreak: false,
                groupIDArr: []
            },
            service = {
                groupCommon: Group_MultistageService
            },
            groupInfo = {}

        $scope.fun.initGroup = function () {
            var template = {
                promise: null,
                groupList: [],
            }
            template.promise = input.resource(input.request).$promise;
            template.promise.then(function (response) {
                response.groupList.unshift({
                    groupID: 0,
                    parentGroupID: -1,
                    groupName: "根目录",
                    groupDepth: 0,
                })
                $scope.interaction.request.groupID = 0;
                template.response = service.groupCommon.sort.init(response, 0, {
                    initGroupDepth: 0
                });
                groupInfo = template.response.groupInfo;
                if (input.item.groupID) {
                    switch (typeof (input.item.groupID)) {
                        case 'object': {
                            angular.copy(input.item.groupID).map(function (val, key) {
                                service.groupCommon.fun.deleteGroup({
                                    currentGroup: groupInfo.groupObj[val],
                                    list: template.response.groupList,
                                    groupInfo: groupInfo,
                                });
                            })
                            break;
                        }
                        default: {
                            service.groupCommon.fun.deleteGroup({
                                currentGroup: groupInfo.groupObj[input.item.groupID],
                                list: template.response.groupList,
                                groupInfo: groupInfo,
                            });
                            break;
                        }
                    }
                }
                $scope.list = template.response.groupList;
                $scope.component.groupCommonObject.mainObject.baseInfo.resetFlag = !$scope.component.groupCommonObject.mainObject.baseInfo.resetFlag;
            })
            return template.promise;
        }
        fun.click = function (arg) {
            var template = {
                uri: null
            }
            $scope.interaction.request.groupID = arg.item.groupID;
            $scope.interaction.request.groupName = arg.item.groupName;
            service.groupCommon.generalFun.initGroupStatus({
                currentGroupID: arg.item.groupID,
                initGroupDepth: input.modalType == 'projectGroup' ? 0 : 1,
                groupInfo: groupInfo,
                list: $scope.list,
            });
        }
        fun.init = (function () {
            //返回groupName,列表有显示分组名的组件所用
            input.staticQuery = input.staticQuery || [];
            $scope.list = input.list || [];
            input.current = input.current || {};
            $scope.component = {
                groupCommonObject: {
                    funObject: {
                        unTop: true,
                        baseFun: {
                            click: fun.click,
                        }
                    },
                    mainObject: {
                        baseInfo: {
                            initGroupDepth: input.modalType == 'projectGroup' ? 0 : 1,
                            status: 'cancelRequest',
                            name: 'groupName',
                            id: 'groupID',
                            current: $scope.interaction.request,
                            hasIcon: input.hasIcon,
                        },
                        staticQuery: input.staticQuery || [],
                    }
                }
            }
            if (input.modalType == 'projectGroup') return;
            if (input.staticQuery.length) {
                $scope.interaction.request.groupID = input.current.groupID || input.staticQuery[0].groupID;
            } else {
                $scope.interaction.request.groupID = input.current.groupID > 0 ? input.current.groupID : input.list[0].groupID;
            }
            groupInfo = service.groupCommon.fun.generalGroupInfo({
                list: $scope.list
            });
            if ($scope.interaction.request.groupID > 0) {
                $scope.interaction.request.groupName = groupInfo.groupObj[$scope.interaction.request.groupID].groupName;
            }
        })();
        $scope.fun.confirm = function () {
            $uibModalInstance.close({
                groupID: $scope.interaction.request.groupID,
                groupName: $scope.interaction.request.groupName,
                groupInfo: groupInfo
            });
        };
        $scope.fun.cancel = function () {
            //$uibModalInstance.dismiss(false);
            $uibModalInstance.close(false);
        };
    }

    ImportModalCtrl.$inject = ['$scope', '$uibModalInstance', 'CODE', '$rootScope', 'input'];

    function ImportModalCtrl($scope, $uibModalInstance, CODE, $rootScope, input) {
        $scope.data = {
            title: input.title
        }
        $scope.fun = {}
        $scope.fun.getFile = function (arg) {
            $scope.$broadcast('$Init_LoadingCommonComponent', {
                file: arg.$file[0]
            });
        }
        $scope.fun.import = function (arg) {
            var template = {
                request: new FormData(),
                promise: null
            }
            for (var key in input.request) {
                template.request.append(key, input.request[key]);
            }
            template.request.append('file', arg.file);
            template.promise = input.resource(template.request).$promise;
            template.promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS: {
                        $uibModalInstance.close(true);
                        break;
                    }
                }
            })
            return template.promise;
        }
        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };

    }

    Common_SingleInputModalCtrl.$inject = ['$scope', '$uibModalInstance', 'input'];

    function Common_SingleInputModalCtrl($scope, $uibModalInstance, input) {
        $scope.data = {
            input: angular.copy(input)
        }
        $scope.fun = {}

        $scope.fun.ok = function () {
            if ($scope.editGroupForm.$invalid) return;
            $uibModalInstance.close({
                text: $scope.data.input.text
            });
        };

        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
    }

    Common_LoginModalCtrl.$inject = ['$scope', 'CODE', '$rootScope', 'md5', 'CommonResource', '$uibModalInstance'];

    function Common_LoginModalCtrl($scope, CODE, $rootScope, md5, CommonResource, $uibModalInstance) {
        $scope.data = {
            info: {
                submitted: false
            },
            interaction: {
                request: {
                    loginCall: '',
                    loginPassword: ''
                }
            },
            fun: {
                close: null, //关闭功能函数
                confirm: null, //确认功能函数
            }
        }
        var data = {
            fun: {
                init: null, //初始化功能函数
            }
        }
        $scope.data.fun.close = function () {
            $uibModalInstance.close(false);
        }
        $scope.data.fun.confirm = function () {
            var template = {
                request: {
                    loginCall: $scope.data.interaction.request.loginCall,
                    loginPassword: md5.createHash($scope.data.interaction.request.loginPassword)
                }
            }

            if ($scope.confirmForm.$valid) {
                $scope.data.info.submitted = false;
                $rootScope.global.ajax.Login_Guest = CommonResource.Guest.Login(template.request);
                $rootScope.global.ajax.Login_Guest.$promise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS: {
                            $uibModalInstance.close({
                                loginCall: template.request.loginCall
                            });
                            break;
                        }
                    }
                })
            } else {
                $scope.data.info.submitted = true;
            }
        }
    }

    EnsureModalCtrl.$inject = ['$scope', '$uibModalInstance', 'title', 'necessity', 'info', 'input'];

    function EnsureModalCtrl($scope, $uibModalInstance, title, necessity, info, input) {

        $scope.title = title;
        $scope.necessity = necessity;
        $scope.info = {
            message: info || "此操作无法恢复，确认操作？",
            btnType: input.btnType || 0, //0：warning 1：info,2:success,
            btnMessage: input.btnMessage || "删除",
            btnGroup: input.btnGroup || [],
            hideDefaultBtn:input.hideDefaultBtn,
            timer: {
                limit: input.timeLimit,
                value: input.timeLimit / 1000,
                fun: null,
            },
            btnCancelMessage: input.btnCancelMessage || "取消"
        }
        $scope.fun = {};
        var data = {
            fun: {
                init: null
            }
        }
        $scope.data = {
            input: {}
        }

        /**
         * 初始化
         */
        data.fun.init = (function () {
            angular.copy(input, $scope.data.input);
            if ($scope.info.btnType == 3) {
                $scope.info.timer.fun = setInterval(function () {
                    if ($scope.info.timer == 0) clearInterval($scope.info.timer.fun);
                    $scope.info.timer.value--;
                    $scope.$root && $scope.$root.$$phase || $scope.$apply();
                }, 1000)
            }
        })()
        $scope.fun.btnClick = function (inputItem) {
            inputItem.confirm();
            $uibModalInstance.close(false);
        }
        $scope.ok = function () {
            if ($scope.sureForm.$valid || !$scope.necessity) {
                $uibModalInstance.close(true);
            } else {
                $scope.submited = true;
            }
        };

        $scope.cancel = function () {
            //$uibModalInstance.dismiss(false);
            clearInterval($scope.info.timer.fun);
            $uibModalInstance.close(false);
        };

    }
    InfoModalCtrl.$inject = ['$scope', '$uibModalInstance', '$timeout', 'info', 'type'];

    function InfoModalCtrl($scope, $uibModalInstance, $timeout, info, type) {

        $scope.type = type || 'info';
        $scope.info = info;
        var timer = $timeout(function () {
            $uibModalInstance.close(true);
        }, 1500, true);
        $scope.$on('$destroy', function () {
            if (timer) {
                $timeout.cancel(timer);
            }
        });
    }


    GroupModalCtrl.$inject = ['$scope', '$uibModalInstance', 'input'];

    function GroupModalCtrl($scope, $uibModalInstance, input) {
        $scope.title = input.title;
        $scope.placeholder = input.placeholder || '1~32位字符串';
        $scope.secondTitle = input.secondTitle || '分组';
        $scope.required = input.hasOwnProperty('required') ? input.required : true;
        $scope.info = {
            groupName: '',
            groupID: '',
            $index: '0',
            isAdd: true
        }
        $scope.params = {
            query: [{
                groupName: '--不设置一级分组--',
                groupID: '0'
            }].concat(input.group),
            status: input.status || 'first-level'
        }
        $scope.fun = {};
        $scope.fun.change = function (ISINIT) {
            if (!ISINIT) $scope.info.parentGroupID = '0';
            if ($scope.info.grandParentGroupID == '0') {
                $scope.params.secondLevelQuery = [{
                    groupName: '--不设置二级分组--',
                    groupID: '0'
                }]
                return;
            }
            for (var key in input.group) {
                if (input.group[key].groupID == $scope.info.grandParentGroupID) {
                    $scope.params.secondLevelQuery = [{
                        groupName: '--不设置二级分组--',
                        groupID: '0'
                    }].concat(input.group[key].childGroupList)
                }
            }
        }
        var init = (function () {
            if (input.data) {
                $scope.info = {
                    groupName: input.data.groupName,
                    groupID: input.data.groupID,
                    isAdd: false
                }
            }
            angular.merge($scope.info, input.parentObject);
            switch (input.status) {
                case 'third-level': {
                    $scope.fun.change(true);
                    break;
                }
            }
        })();
        $scope.ok = function () {
            if ($scope.editGroupForm.$valid) {
                $uibModalInstance.close($scope.info);
            } else {
                $scope.submited = true;
            }
        };

        $scope.cancel = function () {
            //$uibModalInstance.dismiss(false);
            $uibModalInstance.close(false);
        };
    }

    TableModalCtrl.$inject = ['$scope', '$uibModalInstance', 'title', 'info', 'databaseHashKey'];

    function TableModalCtrl($scope, $uibModalInstance, title, info, databaseHashKey) {
        $scope.title = title;
        $scope.info = {
            databaseHashKey: databaseHashKey,
            tableID: '',
            tableName: '',
            tableDesc: '',
            isAdd: true
        }

        function init() {
            if (info) {
                $scope.info = {
                    databaseHashKey: databaseHashKey,
                    tableID: info.tableID,
                    tableName: info.tableName,
                    tableDesc: info.tableDesc,
                    isAdd: false
                }
            }
        }
        init();
        $scope.ok = function () {
            if ($scope.editTableForm.$valid) {
                $uibModalInstance.close($scope.info);
            } else {
                $scope.submited = true;
            }
        };

        $scope.cancel = function () {
            //$uibModalInstance.dismiss(false);
            $uibModalInstance.close(false);
        };
    }


})();