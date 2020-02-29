(function () {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 侧边栏公用服务
     * @required GroupService 
     */
    angular.module('eolinker.service')
        .factory('Group_MultistageService', index)
        .factory('GroupService', GroupFactory);

    GroupFactory.$inject = ['$rootScope']

    function GroupFactory($rootScope) {
        var data = {
                group: null
            },
            fun = {}
        fun.get = function () {
            return data.group;
        }
        fun.set = function (request, boolean) {
            data.group = request;
            if (boolean) {
                $rootScope.$broadcast('$SidebarFinish');
            }
        }
        fun.clear = function () {
            data.group = null;
        }
        return fun;
    }

    index.$inject = ['GroupService']

    function index(GroupService) {
        var data = {
            requestGroupOrder: [],
            time: {},
            service: GroupService,
            generalFun: {},
            fun: {
                clear: null,
                spreed: null,
                operate: null
            },
            sort: {
                operate: null,
                init: null,
            }
        }
        var fun = {};
        var groupInfo = {
            locationArr: [],
            parentGroupPath: {},
            childGroupPath: {
                0: [],
            },
            groupObj: {},
        };
        data.generalFun.resetGroupInfo = function (originData) {
            if (originData) {
                groupInfo = originData;
            } else {
                groupInfo = {
                    locationArr: [],
                    parentGroupPath: {},
                    childGroupPath: {
                        0: [],
                    },
                    groupObj: {},
                };
            }
        }
        data.generalFun.getNextNotChildGroup = function (arg) {
            data.generalFun.resetGroupInfo(arg.groupInfo || groupInfo);
            var currentGroup = groupInfo.groupObj[arg.currentGroupID];
            var index = groupInfo.childGroupPath[currentGroup.parentGroupID].indexOf(currentGroup.groupID);
            if (index != groupInfo.childGroupPath[currentGroup.parentGroupID].length - 1) {
                return groupInfo.childGroupPath[currentGroup.parentGroupID][index + 1]
            } else {
                return fun.getNextNotParentGroup(currentGroup);
            }
        }
        data.generalFun.getGroupLastChildIndex = function (arg) {
            data.generalFun.resetGroupInfo(arg.groupInfo || groupInfo);
            var currentChildGroup = groupInfo.childGroupPath[arg.currentGroupID] || [];
            if (currentChildGroup.length) {
                var nextNotChildGroupID = data.generalFun.getNextNotChildGroup({
                    currentGroupID: currentChildGroup[currentChildGroup.length - 1]
                });
                var nextNoChildGroupIndex = groupInfo.locationArr.indexOf(nextNotChildGroupID);
                return (nextNoChildGroupIndex == -1 ? groupInfo.locationArr.length - 1 : nextNoChildGroupIndex - 1);
            } else {
                return groupInfo.locationArr.indexOf(arg.currentGroupID);
            }
        }
        fun.getNextNotParentGroup = function (currentGroup) {
            if (currentGroup.groupDepth < 2) return;
            var parentGroup = groupInfo.groupObj[currentGroup.parentGroupID];
            var lastGroupArray = groupInfo.childGroupPath[parentGroup.parentGroupID];
            var parentLevelNextGroupID = lastGroupArray[lastGroupArray.indexOf(currentGroup.parentGroupID) + 1];
            if (parentLevelNextGroupID) {
                return parentLevelNextGroupID;
            } else {
                return fun.getNextNotParentGroup(parentGroup);
            }
        }
        fun.insertGroupToParentLast = function (currentGroup) {
            if (currentGroup.parentGroupID) {
                var lastChildIndex = data.generalFun.getGroupLastChildIndex({
                    currentGroupID: currentGroup.parentGroupID
                });
                groupInfo.locationArr.splice(lastChildIndex + 1, 0, currentGroup.groupID);
            } else {
                groupInfo.locationArr.push(currentGroup.groupID);
            }
        }
        fun.orderByGroupOrder = function (arg) {
            if (angular.equals({}, groupOrder)) {
                return;
            }
            var groupInfo = arg.groupInfo;
            var groupOrder = {};
            try {
                groupOrder = JSON.parse(arg.response.groupOrder);
            } catch (e) {

            }
            var obj = [];
            var status = groupInfo.locationArr.length == 0 ? 'reset' : 'child';
            angular.forEach(groupOrder, function (childVal, childKey) {
                obj[childVal] = Number(childKey);
            })
            if (status != 'reset') {
                var index = groupInfo.locationArr.indexOf(arg.response.groupID) + 1;
                groupInfo.childGroupPath[arg.response.groupID] = [];
            } else {
                groupInfo.childGroupPath[0] = [];
                arg.response.groupID = 0;
            }
            angular.forEach(obj, function (childVal, childKey) {
                if (childVal in groupInfo.groupObj) {
                    var childGroup = groupInfo.groupObj[childVal];
                    if (childGroup.parentGroupID == arg.response.groupID) {
                        if (status != 'reset') {
                            groupInfo.parentGroupPath[childVal] = [arg.response.groupID].concat(groupInfo.parentGroupPath[arg.response.groupID] || 0);
                            groupInfo.childGroupPath[arg.response.groupID].push(childVal);
                            groupInfo.locationArr.splice(index, 0, childVal);
                            index++;
                        } else {
                            groupInfo.parentGroupPath[childVal] = [0];
                            groupInfo.childGroupPath[0].push(childVal);
                            groupInfo.locationArr.push(childVal);
                        }
                    }
                }
            })
        }
        data.generalFun.openGroup = function (arg) {
            groupInfo = arg.groupInfo || groupInfo;
            index = groupInfo.locationArr.indexOf(arg.currentGroupID);
            arg.list[index].isOpen = true;
            angular.forEach(groupInfo.childGroupPath[arg.currentGroupID], function (val, key) {
                index = groupInfo.locationArr.indexOf(val);
                if (index != -1) {
                    arg.list[index].hideStatus = false;
                }
            })
        }
        fun.closeGroup = function (arg) {
            angular.forEach(groupInfo.childGroupPath[arg.currentGroupID], function (val, key) {
                index = groupInfo.locationArr.indexOf(val);
                arg.list[index].isOpen = false;
                arg.list[index].hideStatus = true;
                if (groupInfo.childGroupPath[val]) {
                    arg.currentGroupID = val;
                    fun.closeGroup(arg);
                }
            })
        }
        fun.deleteChildGroup = function (currentGroupID) {
            if (groupInfo.childGroupPath[currentGroupID] && groupInfo.childGroupPath[currentGroupID].length) {
                angular.forEach(groupInfo.childGroupPath[currentGroupID], function (val, key) {
                    fun.deleteChildGroup(val);
                })
            }
            groupInfo.groupObj[currentGroupID] = null;
            groupInfo.childGroupPath[currentGroupID] = [];
            groupInfo.parentGroupPath[currentGroupID] = [];
        }
        data.generalFun.sortByLocationArr = function (arg) {
            data.generalFun.resetGroupInfo(arg.groupInfo || groupInfo);
            angular.forEach(groupInfo.locationArr, function (val, key) {
                arg.list[key] = groupInfo.groupObj[val];
            })
            var unnecessaryLength = arg.list.length - groupInfo.locationArr.length;
            if (unnecessaryLength > 0) {
                arg.list.splice(groupInfo.locationArr.length, unnecessaryLength)
            }
        }
        data.fun.deleteGroup = function (arg) {
            data.generalFun.resetGroupInfo(arg.groupInfo);
            var currentIndex = groupInfo.locationArr.indexOf(arg.currentGroup.groupID);
            var nextNoChildGroupID = data.generalFun.getNextNotChildGroup({
                currentGroupID: arg.currentGroup.groupID
            });
            var groupLength = (nextNoChildGroupID ? groupInfo.locationArr.indexOf(nextNoChildGroupID) - currentIndex : groupInfo.locationArr.length - currentIndex);
            groupInfo.locationArr.splice(currentIndex, groupLength);
            groupInfo.childGroupPath[arg.currentGroup.parentGroupID].splice(groupInfo.childGroupPath[arg.currentGroup.parentGroupID].indexOf(arg.currentGroup.groupID), 1);
            groupInfo.parentGroupPath[arg.currentGroup.groupID] = [];
            groupInfo.groupObj[arg.currentGroup.groupID] = null;
            fun.deleteChildGroup(arg.currentGroup.groupID);
            data.generalFun.sortByLocationArr({
                list: arg.list,
            });
        }
        data.generalFun.initGroupStatus = function (arg) {
            if (!arg.list || !arg.list.length) return;
            //分组开关状态、子分组显示状态
            var initGroupDepth = arg.initGroupDepth == 0 ? 0 : 1;
            data.generalFun.resetGroupInfo(arg.groupInfo || groupInfo)
            var currentGroupID = Number(arg.currentGroupID);
            if (arg.status == 'reset') {
                angular.forEach(arg.list, function (val, key) {
                    val.isOpen = false;
                    if (val.groupDepth > initGroupDepth) val.hideStatus = true;
                });
            }
            if (initGroupDepth != 0 && (currentGroupID < 0||currentGroupID==0 || !(currentGroupID in groupInfo.groupObj))) return;
            var currentGroup = groupInfo.groupObj[currentGroupID];
            var index = 0;
            index = groupInfo.locationArr.indexOf(currentGroupID);
            arg.list[index].hideStatus = false;
            arg.list[index].isOpen = true;
            if (currentGroupID != 0) {
                angular.forEach(groupInfo.parentGroupPath[currentGroupID], function (val, key) {
                    if (val || initGroupDepth == 0) {
                        index = groupInfo.locationArr.indexOf(val);
                        arg.list[index].isOpen = true;
                        arg.list[index].hideStatus = false;
                        angular.forEach(groupInfo.childGroupPath[val], function (childVal, childKey) {
                            index = groupInfo.locationArr.indexOf(childVal);
                            arg.list[index].hideStatus = false;
                        })
                    }
                })
            }
            angular.forEach(groupInfo.childGroupPath[currentGroupID], function (val, key) {
                index = groupInfo.locationArr.indexOf(val);
                arg.list[index].hideStatus = false;
            })
        }
        data.fun.generalGroupInfo = function (arg) {
            data.generalFun.resetGroupInfo();
            var groupList = arg.list;
            angular.forEach(groupList, function (val, key) {
                val.groupDepth = val.groupDepth;
                groupInfo.groupObj[val.groupID] = val;
            })
            angular.forEach(groupList, function (val, key) {
                val.groupID = Number(val.groupID);
                val.parentGroupID = Number(val.parentGroupID || 0);
                if (groupInfo.locationArr.indexOf(val.groupID) == -1) {
                    if (val.parentGroupID) {
                        fun.insertGroupToParentLast(val);
                        groupInfo.parentGroupPath[val.groupID] = [val.parentGroupID].concat(groupInfo.parentGroupPath[val.parentGroupID]);
                        groupInfo.childGroupPath[val.parentGroupID] ? groupInfo.childGroupPath[val.parentGroupID].push(val.groupID) : groupInfo.childGroupPath[val.parentGroupID] = [val.groupID];
                    } else {
                        groupInfo.parentGroupPath[val.groupID] = [0];
                        groupInfo.childGroupPath[0].push(val.groupID);
                        groupInfo.locationArr.push(val.groupID);
                    }
                }
            })
            return groupInfo;
        }
        data.sort.init = function (response, currentGroupID, options) {
            var groupList = [];
            options = options || {};
            data.generalFun.resetGroupInfo()
            if (options.responseKey) {
                groupList = response[options.responseKey];
            } else {
                groupList = response.groupList;
            }
           if (!groupList) return {
                groupList: [],
                groupInfo: {},
            };;
            var template = {
                output: []
            }
            template.sortArr = groupList.sort(function (a, b) {
                return a.groupDepth - b.groupDepth;
            })
            angular.forEach(groupList, function (val, key) {
                val.isOpen = false;
                val.groupDepth = val.groupDepth||0;
                groupInfo.groupObj[val.groupID] = val;
            })
            fun.orderByGroupOrder({
                groupInfo: groupInfo,
                response: response
            });
            angular.forEach(template.sortArr, function (val, key) {
                val.groupID = Number(val.groupID);
                val.parentGroupID = Number(val.parentGroupID || 0);
                if (groupInfo.locationArr.indexOf(val.groupID) == -1) {
                    if (val.parentGroupID) {
                        fun.insertGroupToParentLast(val);
                        groupInfo.parentGroupPath[val.groupID] = [val.parentGroupID].concat(groupInfo.parentGroupPath[val.parentGroupID]);
                        groupInfo.childGroupPath[val.parentGroupID] ? groupInfo.childGroupPath[val.parentGroupID].push(val.groupID) : groupInfo.childGroupPath[val.parentGroupID] = [val.groupID];
                    } else {
                        groupInfo.parentGroupPath[val.groupID] = [0];
                        groupInfo.childGroupPath[0].push(val.groupID);
                        groupInfo.locationArr.push(val.groupID);
                    }
                }
                fun.orderByGroupOrder({
                    groupInfo: groupInfo,
                    response: val
                });
            })
            angular.forEach(groupInfo.locationArr, function (val, key) {
                template.output.push(groupInfo.groupObj[val]);
            })
            data.generalFun.initGroupStatus({
                currentGroupID: currentGroupID,
                status: 'reset',
                initGroupDepth: options.initGroupDepth,
                list: template.output,
            });
            return {
                groupList: template.output,
                groupInfo: groupInfo,
            };

        }
        /**
         * 父分组展开功能函数
         * @param {object} arg 参数，eg:{$event:dom,item:单击所处父分组项}
         */
        data.fun.spreed = function (arg) {
            if (arg.$event) {
                arg.$event.stopPropagation();
            }
            data.generalFun.resetGroupInfo(arg.groupInfo);
            var params = {
                currentGroupID: arg.item.groupID,
                list: arg.list,
            }
            if (arg.item.isOpen) {
                var index = groupInfo.locationArr.indexOf(arg.item.groupID);
                arg.list[index].isOpen = false;
                fun.closeGroup(params);
            } else {
                data.generalFun.openGroup(params)
            }
        }
        data.fun.clear = function () {
            data.service.clear();
        };
        data.fun.getGroupPath = function (arg) {
            data.generalFun.resetGroupInfo(arg.groupInfo);
            var template = {
                output: []
            }
            var currentGroupID = arg.currentGroupID;
            var currentGroup = groupInfo.groupObj[currentGroupID];
            if (groupInfo.parentGroupPath[currentGroupID] && groupInfo.parentGroupPath[currentGroupID].length) {
                angular.forEach(groupInfo.parentGroupPath[currentGroupID], function (val, key) {
                    if (val) {
                        var group = groupInfo.groupObj[val];
                        template.output.unshift({
                            groupName: group.groupName,
                            groupID: group.groupID,
                        })
                    }
                })
                template.output.push({
                    groupName: currentGroup.groupName,
                    groupID: currentGroup.groupID,
                });
                return template.output;
            } else {
                return [];
            }
        }
        return data;
    }
})();