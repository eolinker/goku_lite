(function () {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 分组/表组件
     * @param {object} authorityObject 权限类{edit}
     * @param {object} funObject 第一部分功能集类{showVar,btnGroupList{edit:{fun,key,class,showable,icon,tips},sort:{default,cancel,confirm:{fun,key,showable,class,icon,tips}}}}
     * @param {object} requestObject 排序信息{sortable,groupForm}
     * @param {object} mainObject 主类{level,extend,query,baseInfo:{name,id,child,fun:{edit,delete},parentFun:{addChild}}}
     */
    angular.module('eolinker')
        .component('groupDefaultCommonComponent', {
            templateUrl: 'app/component/common/group/default/index.html',
            controller: indexController,
            bindings: {
                list: '<',
                authorityObject: '<',
                funObject: '<',
                requestObject: '<',
                mainObject: '<',
            }
        })

    indexController.$inject = ['$scope', '$rootScope', '$state', '$filter', 'Group_MultistageService', 'GroupService', 'CODE', 'RESPONSE_TEXT'];

    function indexController($scope, $rootScope, $state, $filter, Group_MultistageService, GroupService, CODE, RESPONSE_TEXT) {
        var vm = this;
        vm.data = {
            listIsClick: false,
            sortForm: {
                parentContainment: 'sort-container-gdcc',
                containment: '.actural-group-box',
            }
        };
        vm.fun = {
            more: null,
            common: null
        };
        vm.ajaxResponse = {
            query: []
        };
        vm.service = {
            cache: GroupService,
        }
        var privateFun = {},
            data = {
                maxLevel: 5
            },
            service = {
                groupCommon: Group_MultistageService,
            },
            groupInfo = {
                locationArr: [],
                parentGroupPath: {},
                childGroupPath: {
                    0: [],
                },
                groupObj: {},
            };

        /**
         * @description 分组操作后刷新页面函数
         * @param {Object} inputArg 参数，eg:{parentGroupID:父分组ID,opGroup:目标分组}
         */
        privateFun.reloadItem = function (inputArg) {
            //1.同级分组排序
            if (inputArg.parentGroupID.before == inputArg.parentGroupID.after) return;
            let tmpCurrentGroupID = Number($state.params.groupID);
            //2.移动分组为当前分组
            if (inputArg.opGroup.groupID == tmpCurrentGroupID) return;
            //排序前是否属于当前列表所属分组
            let tmpOriginalFlag = inputArg.parentGroupID.before == tmpCurrentGroupID ? true : privateFun.isFatherGroupID({
                child: inputArg.parentGroupID.before,
                parent: tmpCurrentGroupID
            });
            //排序后是否属于当前列表所属分组
            let tmpNewFlag = inputArg.parentGroupID.after == tmpCurrentGroupID ? true : privateFun.isFatherGroupID({
                child: inputArg.parentGroupID.after,
                parent: tmpCurrentGroupID
            });
            //3.移动前后分组都属于当前列表所属分组
            if (tmpOriginalFlag && tmpNewFlag) return;
            //4.移动前后分组都不属于当前列表所属分组
            if (!tmpOriginalFlag && !tmpNewFlag) return;
            $state.reload($state.current.name);
        }
        privateFun.judgeItemRelation = (groupID1, groupID2) => {
            if(groupID2<0) return 'father';
            let isFather = privateFun.isFatherGroupID({
                child: groupID1,
                parent: groupID2
            });
            let isChild = privateFun.isFatherGroupID({
                child: groupID2,
                parent: groupID1
            });
            if (isFather) {
                return 'father';
            } else if (isChild) {
                return 'child';
            } else{
                return 'noRelationShip'
            }
        }
        /**
         * @description 修改父分组后，重置childGroupPath信息
         * @param {Number} inputGroupID 被重置分组ID
         */
        privateFun.resetChildGroup = function (inputGroupID) {
            angular.forEach(groupInfo.childGroupPath[inputGroupID], (val) => {
                groupInfo.groupObj[val].groupDepth = groupInfo.groupObj[inputGroupID].groupDepth + 1;
                groupInfo.parentGroupPath[val] = [inputGroupID].concat(groupInfo.parentGroupPath[inputGroupID]);
                privateFun.resetChildGroup(val);
            })
        }

        /**
         * 获取分组子层级数
         * @param {number} inputGroupID 当前分组ID
         */
        privateFun.getDeepestGroupDepth = function (inputGroupID) {
            var tmpMaxLen = -1;
            angular.forEach(groupInfo.parentGroupPath, (val) => {
                var tmpLen = val.indexOf(inputGroupID);
                if (tmpLen > -1 && tmpLen > tmpMaxLen) {
                    tmpMaxLen = tmpLen;
                }
            })
            return tmpMaxLen + 2;
        }

        /**
         * 移动到其他分组后，重置分组信息
         * @param {object} inputMoveItem 参数，eg:{inputMoveItem:被移动分组}
         * @param {number} inputTargetGroupID 当前分组ID
         */
        privateFun.resetAfterChangeParentGroupID = function (inputMoveItem, inputTargetGroupID) {
            if (inputMoveItem.parentGroupID == inputTargetGroupID) return;
            groupInfo.groupObj[inputMoveItem.groupID].parentGroupID = inputTargetGroupID;
            if (inputTargetGroupID == 0) {
                groupInfo.parentGroupPath[inputMoveItem.groupID] = [0];
                groupInfo.groupObj[inputMoveItem.groupID].groupDepth = 1;
            } else {
                groupInfo.parentGroupPath[inputMoveItem.groupID] = [inputTargetGroupID].concat(groupInfo.parentGroupPath[inputTargetGroupID]);
                groupInfo.groupObj[inputMoveItem.groupID].groupDepth = groupInfo.groupObj[inputTargetGroupID].groupDepth + 1;
            }
            privateFun.resetChildGroup(inputMoveItem.groupID);
        }

        /**
         * @description 判断是否是父节点
         * @param {object} inputArg  参数对象
         */
        privateFun.isFatherGroupID = function (inputArg) {
            if (inputArg.child == 0) return false;
            if (groupInfo.parentGroupPath[inputArg.child].indexOf(inputArg.parent) > -1) return true;
            return false;
        }
        /**
         * 父分组展开功能函数
         * @param {object} inputArg 参数，eg:{groupInfo:分组信息,list:分组列表}
         */
        privateFun.spreed = function (inputArg) {
            if (vm.funObject.baseFun && vm.funObject.baseFun.spreed) {
                vm.funObject.baseFun.spreed(inputArg);
            } else {
                inputArg.list = vm.ajaxResponse.query;
                inputArg.groupInfo = groupInfo;
                service.groupCommon.fun.spreed(inputArg)
            }
        }

        vm.fun.click = function (inputArg) {
            let tmpPoint = "";
            inputArg = inputArg || {};
            try {
                tmpPoint = inputArg.$event.target.classList[0];
            } catch (e) {
                tmpPoint = 'group';
            }
            switch (tmpPoint) {
                case 'group-icon': {
                    if (!inputArg.item) {
                        inputArg.item = vm.ajaxResponse.query[$filter('findAttr')(inputArg.$event.target, 'data-index')];
                    }
                    privateFun.spreed(inputArg);
                    break;
                }
                case 'actural-group-box': {
                    break;
                }
                default: {
                    if (!inputArg.item) {
                        var tmpIndex = $filter('findAttr')(inputArg.$event.target, 'data-index');
                        inputArg.item = vm.ajaxResponse.query[tmpIndex];
                        inputArg.$index = tmpIndex;
                    }
                    let initGroupStatusFun = () => {
                        service.groupCommon.generalFun.initGroupStatus({
                            currentGroupID: inputArg.item.groupID,
                            groupInfo: groupInfo,
                            list: vm.ajaxResponse.query
                        });

                    }
                    if (vm.funObject.baseFun && vm.funObject.baseFun.click) {
                        vm.funObject.baseFun.click(inputArg, initGroupStatusFun);
                    } else {
                        if (vm.mainObject.baseInfo.current.groupID != inputArg.item.groupID) {
                            $state.go($state.current.name, angular.merge({
                                groupID: inputArg.item.groupID
                            }, vm.requestObject ? vm.requestObject.baseRequest : {}));
                            vm.mainObject.baseInfo.current.groupID = inputArg.item.groupID;
                        }
                        initGroupStatusFun();
                        if (vm.funObject && vm.funObject.clickCallback) {
                            vm.funObject.clickCallback(inputArg);
                        }
                    }
                }
            }
        }
        /**
         * @description 排序分组
         * @param  {object} inputArg {from:排序前groupObj，to:排序后分组groupObj，where:排序位置}
         */
        vm.fun.sort = function (inputArg) {
            if (!vm.mainObject.baseInfo.sort) return;
            inputArg = inputArg || {};
            let tmpParentGroupIDObj = {
                    before: inputArg.from.parentGroupID,
                    after: null
                },
                tmpOriginData = angular.copy(groupInfo),
                tmpCurrentGroupID = inputArg.from.groupID,
                tmpTargetGroupID = inputArg.to.groupID;
            if (privateFun.isFatherGroupID({
                    child: tmpTargetGroupID,
                    parent: tmpCurrentGroupID
                })) return;
            let tmpFromIndex = groupInfo.locationArr.indexOf(tmpCurrentGroupID),
                tmpTargetIndex = groupInfo.locationArr.indexOf(tmpTargetGroupID);
            switch (inputArg.where) {
                case 'before': {
                    //同级分组
                    let tmpNextNoChildGroupID = service.groupCommon.generalFun.getNextNotChildGroup({
                            currentGroupID: tmpCurrentGroupID,
                            groupInfo: groupInfo
                        }),
                        tmpGroupLength = (tmpNextNoChildGroupID ? groupInfo.locationArr.indexOf(tmpNextNoChildGroupID) - tmpFromIndex : groupInfo.locationArr.length - tmpFromIndex);
                    if (tmpFromIndex + tmpGroupLength - 1 == tmpTargetIndex - 1 && inputArg.from.groupDepth == inputArg.to.groupDepth) return;
                    let tmpGroupDepth = privateFun.getDeepestGroupDepth(inputArg.from.groupID);
                    if (tmpGroupDepth + inputArg.to.groupDepth - 1 > data.maxLevel) {
                        $rootScope.InfoModal('分组最多支持' + data.maxLevel + '级', 'error');
                        return;
                    }
                    tmpParentGroupIDObj.after = inputArg.to.parentGroupID;
                    let tmpMovingGroup = groupInfo.locationArr.splice(tmpFromIndex, tmpGroupLength);
                    tmpTargetIndex = groupInfo.locationArr.indexOf(tmpTargetGroupID);
                    tmpMovingGroup.forEach(function (val, key) {
                        groupInfo.locationArr.splice(tmpTargetIndex + key, 0, val);
                    })
                    groupInfo.childGroupPath[tmpParentGroupIDObj.before].splice(groupInfo.childGroupPath[tmpParentGroupIDObj.before].indexOf(tmpCurrentGroupID), 1);
                    groupInfo.childGroupPath[tmpParentGroupIDObj.after].splice(groupInfo.childGroupPath[tmpParentGroupIDObj.after].indexOf(tmpTargetGroupID), 0, tmpCurrentGroupID);
                    privateFun.resetAfterChangeParentGroupID(inputArg.from, tmpParentGroupIDObj.after);
                    break;
                }
                case 'in': {
                    let tmpIsLastChild = groupInfo.childGroupPath[tmpParentGroupIDObj.before].indexOf(tmpCurrentGroupID) == groupInfo.childGroupPath[tmpParentGroupIDObj.before].length - 1;
                    if (tmpParentGroupIDObj.before == tmpParentGroupIDObj.after && tmpIsLastChild) return;
                    let tmpGroupDepth = privateFun.getDeepestGroupDepth(inputArg.from.groupID);
                    if (tmpGroupDepth + inputArg.to.groupDepth > data.maxLevel) {
                        $rootScope.InfoModal('分组最多支持' + data.maxLevel + '级', 'error');
                        return;
                    }
                    tmpParentGroupIDObj.after = inputArg.to.groupID;
                    let tmpNextNoChildGroupID = service.groupCommon.generalFun.getNextNotChildGroup({
                            currentGroupID: tmpCurrentGroupID,
                            groupInfo: groupInfo
                        }),
                        tmpGroupLength = (tmpNextNoChildGroupID ? groupInfo.locationArr.indexOf(tmpNextNoChildGroupID) - tmpFromIndex : groupInfo.locationArr.length - tmpFromIndex);
                    let tmpMovingGroup = groupInfo.locationArr.splice(tmpFromIndex, tmpGroupLength);
                    groupInfo.childGroupPath[tmpParentGroupIDObj.before].splice(groupInfo.childGroupPath[tmpParentGroupIDObj.before].indexOf(tmpCurrentGroupID), 1);
                    tmpTargetIndex = service.groupCommon.generalFun.getGroupLastChildIndex({
                        currentGroupID: tmpParentGroupIDObj.after,
                        groupInfo: groupInfo
                    }) + 1;
                    tmpMovingGroup.forEach(function (val, key) {
                        groupInfo.locationArr.splice(tmpTargetIndex + key, 0, val);
                    })
                    groupInfo.childGroupPath[tmpParentGroupIDObj.after] = groupInfo.childGroupPath[tmpParentGroupIDObj.after] || [];
                    groupInfo.childGroupPath[tmpParentGroupIDObj.after].push(tmpCurrentGroupID);
                    privateFun.resetAfterChangeParentGroupID(inputArg.from, tmpParentGroupIDObj.after);
                    //变成子分组
                    break;
                }
                case 'after': {
                    if (tmpFromIndex == tmpTargetIndex + 1 && inputArg.from.groupDepth == inputArg.to.groupDepth) break;
                    let tmpGroupDepth = privateFun.getDeepestGroupDepth(inputArg.from.groupID);
                    if (tmpGroupDepth + inputArg.to.groupDepth - 1 > data.maxLevel) {
                        $rootScope.InfoModal('分组最多支持' + data.maxLevel + '级', 'error');
                        return;
                    }
                    tmpParentGroupIDObj.after = inputArg.to.parentGroupID;
                    let tmpNextNoChildGroupID = service.groupCommon.generalFun.getNextNotChildGroup({
                        currentGroupID: tmpCurrentGroupID,
                        groupInfo: groupInfo
                    });
                    groupInfo.childGroupPath[tmpParentGroupIDObj.before].splice(groupInfo.childGroupPath[tmpParentGroupIDObj.before].indexOf(tmpCurrentGroupID), 1);
                    let tmpGroupLength = (tmpNextNoChildGroupID ? groupInfo.locationArr.indexOf(tmpNextNoChildGroupID) - tmpFromIndex : groupInfo.locationArr.length - tmpFromIndex);
                    let tmpMovingGroup = groupInfo.locationArr.splice(tmpFromIndex, tmpGroupLength);
                    tmpTargetIndex = groupInfo.locationArr.length;
                    tmpMovingGroup.forEach(function (val, key) {
                        groupInfo.locationArr.splice(tmpTargetIndex + key, 0, val);
                    })
                    groupInfo.childGroupPath[tmpParentGroupIDObj.after].push(tmpCurrentGroupID);
                    privateFun.resetAfterChangeParentGroupID(inputArg.from, tmpParentGroupIDObj.after);
                    break;
                }
                default: {
                    return;
                }
            }
            service.groupCommon.generalFun.sortByLocationArr({
                list: vm.ajaxResponse.query,
                groupInfo: groupInfo,
            });
            let tmpAjaxRequest = {
                groupOrder: null
            }
            tmpAjaxRequest.groupOrder = [{
                parentGroupID: tmpParentGroupIDObj.after,
                groupID: [tmpCurrentGroupID]
            }]
            if (tmpParentGroupIDObj.after != tmpParentGroupIDObj.before) {
                tmpAjaxRequest.groupOrder.push({
                    parentGroupID: tmpParentGroupIDObj.before
                })
            }
            angular.forEach(tmpAjaxRequest.groupOrder, (val) => {
                let groupOrder = {};
                angular.forEach(groupInfo.childGroupPath[val.parentGroupID], (childVal, childKey) => {
                    groupOrder[childVal] = childKey;
                })
                val.groupOrder = groupOrder;
            })
            tmpAjaxRequest.groupOrder = JSON.stringify(tmpAjaxRequest.groupOrder);
            angular.merge(tmpAjaxRequest, vm.requestObject.baseRequest);
            vm.requestObject.resource.Sort(tmpAjaxRequest).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS: {
                        privateFun.reloadItem({
                            parentGroupID: tmpParentGroupIDObj,
                            opGroup: inputArg.from
                        });
                        if (inputArg.where == 'in') {
                            service.groupCommon.generalFun.openGroup({
                                currentGroupID: tmpParentGroupIDObj.after,
                                list: vm.ajaxResponse.query,
                                groupInfo: groupInfo,
                            });
                        }
                        break;
                    }
                    default: {
                        $rootScope.InfoModal('排序失败，' + RESPONSE_TEXT.FAILURE, 'error');
                        groupInfo = tmpOriginData;
                        service.groupCommon.generalFun.resetGroupInfo(groupInfo);
                        service.groupCommon.generalFun.sortByLocationArr({
                            list: vm.ajaxResponse.query,
                            groupInfo: groupInfo,
                        });
                    }
                }
                vm.service.cache.set(vm.ajaxResponse.query);
            })
        }
        privateFun.init = function (status) {
            let tmpAjaxRequest = angular.copy(vm.requestObject.baseRequest),
                tmpSortObj = {
                    response: null
                }
            vm.service.cache.clear()
            if (!data.requesting) {
                data.requesting = true;
                $rootScope.global.ajax.Query_Group=vm.requestObject.resource.Query(tmpAjaxRequest);
                $rootScope.global.ajax.Query_Group.$promise.then(function (response) {
                    data.requesting = false;
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS: {
                            if(vm.mainObject.baseInfo.queryName){
                                response.groupList=response[vm.mainObject.baseInfo.queryName]
                            }
                            tmpSortObj.response = service.groupCommon.sort.init(response, vm.mainObject.baseInfo.current.groupID);
                            vm.ajaxResponse.query = tmpSortObj.response.groupList;
                            groupInfo = tmpSortObj.response.groupInfo;
                            if (vm.funObject.callback) vm.funObject.callback.querySuccess(vm.ajaxResponse.query,);
                            break;
                        }
                    }
                })
            }
        }
        /**
         * @description 编辑分组
         * @param  {string} status 状态 eg:edit/add
         * @param  {object} inputArg {item:单列表项}
         */
        privateFun.edit = function (status, inputArg) {
            if (status == 'add-child' && inputArg.item.groupDepth >= data.maxLevel) {
                $rootScope.InfoModal('分组最多支持' + data.maxLevel + '级', 'error');
                return;
            }
            inputArg = inputArg || {};
            let tmpModal = {
                title: (status == 'edit' ? "编辑" : "新建") + (inputArg.item ? "子分组" : "分组"),
                data: status == 'edit' ? inputArg.item : null,
                secondTitle: "分组名称",
            };
            $rootScope.GroupModal(tmpModal, function (callback) {
                if (callback) {
                    angular.merge(callback, vm.requestObject.baseRequest);
                    if (status == 'edit') {
                        callback.groupID = inputArg.item.groupID;
                        $filter('Field_CommonFilter')('object', callback, ['$index']);
                        vm.requestObject.resource.Edit(callback).$promise.then(function (response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS: {
                                    $rootScope.InfoModal(tmpModal.title + '成功！', 'success');
                                    privateFun.init();
                                    break;
                                }
                            }
                        });
                    } else {
                        if (status == 'add-child') {
                            callback.parentGroupID = inputArg.item.groupID;
                        }
                        $filter('Field_CommonFilter')('object', callback, ['$index', 'groupID']);
                        vm.requestObject.resource.Add(callback).$promise.then(function (response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS: {
                                    $rootScope.InfoModal(tmpModal.title + '成功！', 'success');
                                    privateFun.init();
                                    break;
                                }
                            }
                        });
                    }
                }
            });
        }
        /**
         * @description 删除分组
         * @param  {object} inputArg 传参object{item:单列表项,modal:弹框信息,callback:回调函数}
         */
        privateFun.delete = function (inputArg) {
            inputArg = inputArg || {};
            let tmpAjaxRequest = Object.assign({}, {
                    groupID: inputArg.item.groupID
                }, vm.requestObject.baseRequest),
                tmpModal = inputArg.modal
            $rootScope.EnsureModal(tmpModal.title, false, tmpModal.message, {}, function (callback) {
                if (callback) {
                    vm.requestObject.resource.Delete(tmpAjaxRequest).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                let nowGroupWithDelteItemRelation=privateFun.judgeItemRelation(inputArg.item.groupID,vm.mainObject.baseInfo.current.groupID);
                                service.groupCommon.fun.deleteGroup({
                                    currentGroup: inputArg.item,
                                    list: vm.ajaxResponse.query,
                                    groupInfo: groupInfo
                                });
                                vm.service.cache.set(vm.ajaxResponse.query);
                                $rootScope.InfoModal('分组删除成功！', 'success');
                                if (inputArg.callback) {
                                    inputArg.callback(inputArg, angular.copy(response),{
                                        nowGroupWithDelteItemRelation:nowGroupWithDelteItemRelation
                                    });
                                } else if ($state.params.groupID == inputArg.item.groupID) {
                                    vm.mainObject.baseInfo.current[vm.mainObject.baseInfo.id]=inputArg.item.parentGroupID || -1;
                                    $state.go($state.current.name, angular.merge({
                                        groupID: inputArg.item.parentGroupID || -1
                                    }, vm.requestObject ? vm.requestObject.baseRequest : {}));
                                } else {
                                    privateFun.reloadItem({
                                        parentGroupID: {
                                            before: inputArg.item.parentGroupID,
                                            after: 0
                                        },
                                        opGroup: inputArg.item
                                    });
                                }
                                break;
                            }
                            default: {
                                $rootScope.InfoModal('分组删除失败，' + RESPONSE_TEXT.FAILURE, 'error');
                            }
                        }
                    })
                }
            });
        }

        /**
         * 单项列更多操作
         * @param  {object} inputArg 传参object{item:单列表项,$event:点击事件对象}
         */
        vm.fun.more = function (inputArg) {
            let tmpPoint = "";
            inputArg.$event.stopPropagation();
            inputArg.item.listIsClick = true;
            vm.data.listIsClick = true;
            try {
                tmpPoint = inputArg.$event.target.classList[0];
            } catch (e) {
                tmpPoint = 'group';
            }
            switch (tmpPoint) {
                case 'more-icon': {
                    break;
                }
                case 'active': {
                    break;
                }
                default: {
                    let tmpIndex = $filter('findAttr')(inputArg.$event.target, 'data-child-index');
                    vm.fun.common(vm.mainObject.itemFun[tmpIndex], inputArg);
                }
            }

        }

        /**
         * @description 统筹绑定调用页面列表功能单击函数
         * @param {obejct} extend 方式值
         * @param {object} arg 共用体变量，后根据传值函数回调方法
         */
        vm.fun.common = function (extend, arg) {
            if (!extend) return;
            let tmpParam = arg || {};
            if (!extend.fun) {
                extend.fun = privateFun[extend.funName];
            }
            switch (typeof (extend.params)) {
                case 'string': {
                    return eval('extend.fun(' + extend.params + ')');
                }
                default: {
                    for (var key in extend.params) {
                        if (extend.params[key] == null) {
                            tmpParam[key] = arg[key];
                        } else {
                            tmpParam[key] = extend.params[key];
                        }
                    }
                    return extend.fun(tmpParam,privateFun.init);
                }
            }

        }
        vm.$onInit = function () {
            if (vm.mainObject.baseInfo.type != 'modal') {
                vm.service.cache.isShrink = false;
            }
            vm.mainObject.baseInfo.initGroupDepth = vm.mainObject.baseInfo.initGroupDepth == 0 ? 0 : 1;
            if (vm.mainObject.baseInfo.status != 'cancelRequest') {
                privateFun.init();
            }
        }
        privateFun.watchResetFlag = $scope.$watch('$ctrl.mainObject.baseInfo.resetFlag', function () {
            switch (vm.mainObject.baseInfo.status) {
                case 'cancelRequest': {
                    vm.ajaxResponse.query = vm.list;
                    groupInfo = service.groupCommon.fun.generalGroupInfo({
                        list: vm.ajaxResponse.query
                    })
                    service.groupCommon.generalFun.initGroupStatus({
                        currentGroupID: vm.mainObject.baseInfo.current.groupID,
                        groupInfo: groupInfo,
                        list: vm.ajaxResponse.query,
                    });
                    break;
                }
                default: {
                    if (vm.requestObject) {
                        privateFun.init()
                    }
                }
            }
        });
        vm.$onDestroy = function () {
            privateFun.watchResetFlag();
        }

    }
})();