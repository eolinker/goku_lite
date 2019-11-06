(function () {
    'use strict';
    angular.module('eolinker')
        .component('listBlockCommonComponent', {
            templateUrl: 'app/component/listBlock/index.html',
            controller: indexController,
            bindings: {
                otherObject: '=',
                authorityObject: '<',
                mainObject: '<',
                list: '=',
                activeObject: '=',
                pageObject: '<'
            }
        })
    indexController.$inject = ['$rootScope', '$scope'];

    function indexController($rootScope, $scope) {
        var vm = this,
            fun = {};
        vm.data = {
            sortForm: {
                parentContainment: 'tbody-div',
                containment: '.tbody-div'
            },
            sortAuthorityVar: '',
            sort: false,
            isEditTable: false,
            html: '',
            partHtml: {},
            movePart: null,
            checkboxTdObject: {
                selectAll: false,
                indexAddress: {},
                query: []
            }
        };
        vm.fun = {};
        var data = {
            radioOriginalIndex: 0,
            movePart: null
        }
        vm.fun.sort = function (arg) {
            var tmpPartModule = data.movePart;
            if (!vm.data.sort) return;
            if (vm.mainObject.setting.hasOwnProperty('unSortIndex') && vm.mainObject.setting.unSortIndex === arg.targetIndex) {
                return;
            }
            switch (arg.where) {
                case 'before':
                case 'in':
                case 'after': {
                    break;
                }
                default: {
                    return;
                }
            }
            arg = arg || {};
            var tmp = {
                list: [],
                oldList: angular.copy(vm.mainObject.setting.isPartModule ? vm.list[tmpPartModule] : vm.list),
                index: arg.originIndex + 1,
                targetIndex: arg.targetIndex
            }
            tmp.list.push(Object.assign({}, arg.from, {
                listDepth: arg.where === 'in' ? (arg.to.listDepth + 1) : arg.to.listDepth,
                isHide: arg.where === 'in' && arg.to.isShrink ? true : false
            }))
            let tmpFunListParse = () => {
                var val = tmp.oldList[tmp.index];
                if (arg.where === 'in') {
                    val.listDepth = arg.to.listDepth + val.listDepth - arg.from.listDepth + 1;
                } else {
                    val.listDepth = val.listDepth - (arg.from.listDepth - arg.to.listDepth);
                }
                tmp.list.push(val);
                tmp.index++;
            }
            if (vm.mainObject.baseFun.sortPartLastIndex) {
                while (tmp.index < arg.groupList.length && (vm.mainObject.baseFun.sortPartLastIndex(arg.originIndex + 1, arg.groupList[tmp.index]) || arg.groupList[tmp.index].listDepth > arg.from.listDepth)) {
                    tmpFunListParse();
                }
            } else {
                while (tmp.index < arg.groupList.length && arg.groupList[tmp.index].listDepth > arg.from.listDepth) {
                    tmpFunListParse();
                }
            }
            if (arg.targetIndex > arg.originIndex && arg.targetIndex < tmp.index) return;
            tmp.oldList.splice(arg.originIndex, tmp.index - arg.originIndex);
            if (arg.targetIndex > arg.originIndex) {
                arg.targetIndex = arg.targetIndex - (tmp.index - arg.originIndex) + 1;
                tmp.targetIndex = arg.targetIndex - 1;
            }
            if (tmp.targetIndex < 0) return;
            var tmpResultList = null;
            switch (arg.where) {
                case 'before': {
                    if (arg.originIndex < arg.targetIndex) {
                        tmpResultList = tmp.oldList.slice(0, arg.targetIndex - 1).concat(tmp.list).concat(tmp.oldList.slice(arg.targetIndex - 1, tmp.oldList.length));
                    } else {
                        tmpResultList = tmp.oldList.slice(0, arg.targetIndex || 0).concat(tmp.list).concat(tmp.oldList.slice(arg.targetIndex || 0, tmp.oldList.length));
                    }
                    break;
                }
                case 'in': {
                    if (arg.to.listDepth >= 4) {
                        return;
                    }
                    if (vm.mainObject.baseFun.sortIn) {
                        vm.mainObject.baseFun.sortIn(tmp.oldList[tmp.targetIndex])
                    }
                    if (arg.originIndex < arg.targetIndex) {
                        tmpResultList = tmp.oldList.slice(0, arg.targetIndex || 1).concat(tmp.list).concat(tmp.oldList.slice(arg.targetIndex || 1, tmp.oldList.length));
                    } else {
                        tmpResultList = tmp.oldList.slice(0, arg.targetIndex + 1).concat(tmp.list).concat(tmp.oldList.slice(arg.targetIndex + 1, tmp.oldList.length));
                    }
                    break;
                }
                case 'after': {
                    tmpResultList = tmp.oldList.slice(0, arg.targetIndex || 1).concat(tmp.list).concat(tmp.oldList.slice(arg.targetIndex || 1, tmp.oldList.length));
                    break;
                }
                default: {
                    return;
                }
            }

            if (vm.mainObject.setting.isPartModule) {
                vm.list[tmpPartModule] = tmpResultList;
            } else {
                vm.list = tmpResultList;
            }
            if (vm.mainObject.baseFun.sort) {
                vm.mainObject.baseFun.sort(tmpResultList, tmpPartModule);
            }
        }
        fun.getTargetEvent = function ($event, inputPointAttr) {
            var itemIndex = $event.getAttribute(inputPointAttr || 'eo-attr-index');
            if (itemIndex) {
                return $event;
            } else {
                return fun.getTargetEvent($event.parentNode, inputPointAttr);
            }
        }
        fun.getTargetIndex = function ($event, inputPointAttr) {
            var itemIndex = $event.getAttribute(inputPointAttr || 'eo-attr-index');
            if (itemIndex) {
                return itemIndex;
            } else {
                return fun.getTargetIndex($event.parentNode, inputPointAttr);
            }
        }
        fun.deleteItem = function (inputIndex) {
            if (vm.data.isDepth) {
                vm.list.splice(inputIndex, (fun.getLastItemIndex(inputIndex, vm.list) - inputIndex) || 1);
            } else {
                vm.list.splice(inputIndex, 1);
            }
        }
        fun.insertItem = function (inputObject) {
            vm.list.splice(inputObject.$index, 0, Object.assign({}, {
                listDepth: inputObject.item.listDepth
            }, vm.mainObject.itemStructure));
        }
        fun.addChildItem = function (inputObject) {
            if (vm.mainObject.baseFun.reduceItemWhenAddChildItem) {
                vm.mainObject.baseFun.reduceItemWhenAddChildItem(inputObject.item)
            }
            vm.list.splice(fun.getLastItemIndex(inputObject.$index, vm.list) || 1, 0, Object.assign({}, {
                listDepth: (inputObject.item.listDepth || 0) + 1,
                isHide: inputObject.item.isShrink ? true : false
            }, vm.mainObject.itemStructure));
        }
        fun.clickCheckbox = function (inputTdObject, inputItemIndex, inputPartIndex) {
            var tmpAuthority = inputTdObject.authority;
            if (tmpAuthority && !vm.authorityObject[tmpAuthority]) return;
            var tmpList = vm.mainObject.setting.isPartModule ? vm.list[inputPartIndex] : vm.list;
            if (inputTdObject.fun) {
                inputTdObject.fun({
                    item: tmpList[inputItemIndex],
                    $index: inputItemIndex
                })
                return;
            }
            if (vm.mainObject.baseFun.checkIsValidItem) {
                if (!vm.mainObject.baseFun.checkIsValidItem({
                        item: tmpList[inputItemIndex],
                        $index: inputItemIndex,
                        type: inputTdObject.type
                    })) return;
            }
            if (inputTdObject.modelKey) {
                let tmpBatchObj = null;
                tmpList[inputItemIndex][inputTdObject.modelKey] = !tmpList[inputItemIndex][inputTdObject.modelKey];
                switch (inputTdObject.type) {
                    case 'checkbox': {
                        tmpBatchObj = vm.data.checkboxTdObject;
                        break;
                    }
                    case 'relationalCheckbox': {
                        tmpBatchObj = inputTdObject;
                        if (inputTdObject.checkIsValidToRelate(tmpList[inputItemIndex])) {
                            fun.clickCheckbox(vm.mainObject.tdList[data.checkboxTdIndex], inputItemIndex, inputPartIndex);
                        }
                        break;
                    }
                }
                if (tmpList[inputItemIndex][inputTdObject.modelKey]) {
                    data.queryLength++;
                    if (data.queryLength === (vm.list || []).length) {
                        tmpBatchObj.selectAll = true;
                    }
                } else {
                    data.queryLength--;
                    tmpBatchObj.selectAll = false;
                }
            } else {
                var tmpItemActiveKeyValue = tmpList[inputItemIndex][inputTdObject.activeKey];
                if (tmpItemActiveKeyValue === null) return;
                if (vm.data.checkboxTdObject.indexAddress[tmpItemActiveKeyValue]) {
                    vm.data.checkboxTdObject.query.splice(vm.data.checkboxTdObject.query.indexOf(tmpItemActiveKeyValue), 1);
                    delete vm.data.checkboxTdObject.indexAddress[tmpItemActiveKeyValue];
                    vm.data.checkboxTdObject.selectAll = false;
                    if (vm.mainObject.baseFun.clickCheckbox) {
                        vm.mainObject.baseFun.clickCheckbox('minus-single');
                    }
                } else {
                    vm.data.checkboxTdObject.indexAddress[tmpItemActiveKeyValue] = inputTdObject.hasOwnProperty('activeValue') ? inputTdObject.activeValue : (parseInt(inputItemIndex) + 1);
                    vm.data.checkboxTdObject.query.push(tmpItemActiveKeyValue);
                    var tmpQuery = [];
                    if (vm.mainObject.setting && vm.mainObject.setting.isScrollLoad) {
                        tmpQuery = vm.otherObject.allQuery.filter(function (val, key) {
                            if (val[inputTdObject.activeKey]) {
                                return true;
                            }
                            return false;
                        });
                    } else if (vm.mainObject.setting.isPartModule) {
                        for (let key in vm.list) {
                            tmpQuery = tmpQuery.concat(vm.list[key]);
                        }
                    } else {
                        tmpQuery = vm.list;
                    }
                    if (tmpQuery.length === (vm.data.checkboxTdObject.query || []).length) {
                        vm.data.checkboxTdObject.selectAll = true;
                    }
                    if (vm.mainObject.baseFun.clickCheckbox) {
                        vm.mainObject.baseFun.clickCheckbox('plus-single');
                    }
                }
                if (vm.mainObject.baseFun.teardownWhenCheckboxIsClick) {
                    vm.mainObject.baseFun.teardownWhenCheckboxIsClick(vm.data.checkboxTdObject, tmpList);
                }
            }
        }
        vm.fun.moreItemClick = function ($event, inputPartIndex) {
            $event.stopPropagation();
            var tmp = {};
            tmp.itemIndex = parseInt(fun.getTargetIndex($event.target));
            tmp.btnObject = vm.mainObject.tdList[fun.getTargetIndex($event.target, 'eo-attr-td-index')].btnList[fun.getTargetIndex($event.target, 'eo-attr-btn-index')];
            tmp.btnObject = tmp.btnObject.funArr[fun.getTargetIndex($event.target, 'eo-attr-btn-fun-index')];
            if (tmp.btnObject.fun) {
                var inputArg = {
                    item: vm.mainObject.setting.isPartModule ? vm.list[inputPartIndex][tmp.itemIndex] : vm.list[tmp.itemIndex],
                    $index: tmp.itemIndex
                };
                switch (typeof tmp.btnObject.param) {
                    case 'string': {
                        eval('tmp.btnObject.fun(' + tmp.btnObject.param + ')');
                        return;
                    }
                    default: {
                        tmp.btnObject.fun(Object.assign(inputArg, tmp.btnObject.param));
                        return;
                    }
                }
            }
        }
        vm.fun.itemClick = function ($event, inputPartIndex) {
            // $event.stopPropagation();
            var tmp = {};
            try {
                tmp.point = $event.target.classList[0];
                if ($event.target.classList.value.indexOf('input-checkbox') > -1) {
                    tmp.point = 'input-checkbox';
                }
            } catch (e) {
                console.log(e)
                tmp.point = 'default';
            }
            if (/container-tbd/.test(tmp.point)) return;
            if (/^(btn-)|(fbtn-)|(cbtn-)/.test(tmp.point)) {
                tmp.itemIndex = parseInt(fun.getTargetIndex($event.target));
                tmp.btnObject = vm.mainObject.tdList[fun.getTargetIndex($event.target, 'eo-attr-td-index')].btnList[fun.getTargetIndex($event.target, 'eo-attr-btn-index')];
                if (!tmp.btnObject.isUnWantToStopPropagation) {
                    $event.stopPropagation();
                }
                if (tmp.point === 'btn-funItem') {
                    tmp.btnObject = tmp.btnObject.funArr[fun.getTargetIndex($event.target, 'eo-attr-btn-fun-index')];
                }
                if (tmp.btnObject.fun) {
                    var inputArg = {
                        item: vm.mainObject.setting.isPartModule ? vm.list[inputPartIndex][tmp.itemIndex] : vm.list[tmp.itemIndex],
                        $index: tmp.itemIndex
                    };
                    if (/^(fbtn-)/.test(tmp.point)) {
                        inputArg.callback = vm.fun.watchFormLastChange;
                    }
                    switch (typeof tmp.btnObject.param) {
                        case 'string': {
                            eval('tmp.btnObject.fun(' + tmp.btnObject.param + ')');
                            return;
                        }
                        default: {
                            tmp.btnObject.fun(Object.assign(inputArg, tmp.btnObject.param));
                            return;
                        }
                    }
                }
                switch (tmp.point) {
                    case 'btn-delete':
                    case 'cbtn-delete': {
                        fun.deleteItem(tmp.itemIndex);
                        break;
                    }
                    case 'btn-addChild': {
                        fun.addChildItem({
                            item: vm.list[tmp.itemIndex],
                            $index: tmp.itemIndex
                        });
                        break;
                    }
                    case 'btn-insert': {
                        fun.insertItem({
                            item: vm.list[tmp.itemIndex],
                            $index: tmp.itemIndex
                        });
                        break;
                    }
                }
            } else {
                $event.stopPropagation();
                if (data.checkboxClickAffectTotalItem && vm.data.checkboxTdObject.isOperating) {
                    tmp.point = 'input-checkbox';
                } else if (data.radioClickAffectTotalItem) {
                    tmp.point = 'input-radio';
                }
                switch (tmp.point) {
                    case 'input-checkbox': {
                        fun.clickCheckbox(vm.mainObject.tdList[data.checkboxTdIndex], fun.getTargetIndex($event.target), inputPartIndex);
                        break;
                    }
                    case 'relational-checkbox': {
                        fun.clickCheckbox(vm.mainObject.tdList[data.relationalCheckboxTdIndex], fun.getTargetIndex($event.target), inputPartIndex);
                        break;
                    }
                    case 'input-radio': {
                        tmp.tdObject = vm.mainObject.tdList[data.radioTdIndex];
                        tmp.itemIndex = fun.getTargetIndex($event.target);
                        if (tmp.tdObject.disabledModelKey && vm.list[tmp.itemIndex][tmp.tdObject.disabledModelKey]) {
                            return;
                        }
                        if ((data.radioOriginalIndex || 0).toString() === tmp.itemIndex && tmp.tdObject.isCanBecancel) {
                            vm.list[tmp.itemIndex][tmp.tdObject.modelKey] = !vm.list[tmp.itemIndex][tmp.tdObject.modelKey];
                            data.radioOriginalIndex = 0;
                        } else {
                            vm.list[data.radioOriginalIndex][tmp.tdObject.modelKey] = false;
                            vm.list[tmp.itemIndex][tmp.tdObject.modelKey] = true;
                            data.radioOriginalIndex = tmp.itemIndex;
                        }

                        break;
                    }
                }
                if (!vm.data.checkboxTdObject.isOperating && vm.mainObject.baseFun.trClick) {
                    tmp.itemIndex = parseInt(fun.getTargetIndex($event.target));
                    vm.mainObject.baseFun.trClick({
                        item: vm.list[tmp.itemIndex],
                        $index: tmp.itemIndex
                    })
                }
            }
        }
        vm.fun.selectAll = function (inputTdIndex) {
            let tmpTdObj = vm.mainObject.tdList[inputTdIndex],
                tmpBatchObj = tmpTdObj.type === "checkbox" ? vm.data.checkboxTdObject : tmpTdObj;
            var tmpAuthority = tmpTdObj.authority;
            if (tmpAuthority && !vm.authorityObject[tmpAuthority]) return;
            var tmp = {
                modelKey: tmpTdObj.modelKey,
                activeKey: tmpTdObj.activeKey
            }
            tmpBatchObj.selectAll = !tmpBatchObj.selectAll;
            switch (tmpTdObj.type) {
                case 'relationalCheckbox': {
                    if (tmpTdObj.checkIsValidToRelateAll(tmpBatchObj.selectAll)) {
                        vm.fun.selectAll(data.checkboxTdIndex);
                    }
                }
            }
            if (tmp.modelKey) {
                if (!tmpBatchObj.selectAll) {
                    if (vm.list.length === 1 && vm.mainObject.setting.isStaticFirstIndex) {
                        tmpBatchObj.selectAll = true;
                        return;
                    }
                }
                for (var key in vm.list) {
                    if (vm.mainObject.setting.isStaticFirstIndex && key === '0') {
                        continue;
                    }
                    vm.list[key][tmp.modelKey] = tmpBatchObj.selectAll;
                }
                data.queryLength = tmpBatchObj.selectAll ? (vm.list || []).length : 0;
                if (vm.mainObject.baseFun.clickCheckbox) {
                    vm.mainObject.baseFun.clickCheckbox(`${tmpBatchObj.selectAll?'plus':'minus'}-all`);
                }
            } else {
                let tmpIndexAddress = vm.data.checkboxTdObject.indexAddress;
                let tmpOldListLength = vm.data.checkboxTdObject.query.length;
                vm.data.checkboxTdObject.indexAddress = {};
                vm.data.checkboxTdObject.query = [];
                if (vm.data.checkboxTdObject.selectAll) {
                    if (vm.mainObject.setting.isPartModule) {
                        for (var moduleKey in vm.list) {
                            var moduleVal = vm.list[moduleKey];
                            for (let key in moduleVal) {
                                vm.data.checkboxTdObject.query.push(moduleVal[key][tmp.activeKey]);
                                vm.data.checkboxTdObject.indexAddress[moduleVal[key][tmp.activeKey]] = tmpTdObj.hasOwnProperty('activeValue') ? tmpTdObj.activeValue : (parseInt(key) + 1);
                            }
                        }
                    } else {
                        if (vm.mainObject.setting.disabledSelectModelKey) {
                            for (let key in vm.list) {
                                if (vm.list[key][vm.mainObject.setting.disabledSelectModelKey] !== vm.mainObject.setting.disabledSelectVal) {
                                    vm.data.checkboxTdObject.query.push(vm.list[key][tmp.activeKey]);
                                    vm.data.checkboxTdObject.indexAddress[vm.list[key][tmp.activeKey]] = tmpTdObj.hasOwnProperty('activeValue') ? tmpTdObj.activeValue : (parseInt(key) + 1);
                                }
                            }
                        } else {
                            let tmpQuery = []
                            if (vm.mainObject.setting && vm.mainObject.setting.isScrollLoad) {
                                tmpQuery = vm.otherObject.allQuery;
                            } else {
                                tmpQuery = vm.list;
                            }
                            if (vm.mainObject.baseFun.checkIsValidItem) {
                                for (let key in tmpQuery) {
                                    if (tmpQuery[key][tmp.activeKey] === null || !vm.mainObject.baseFun.checkIsValidItem({
                                            item: tmpQuery[key],
                                            indexAddress: tmpIndexAddress,
                                            isSelectAll: true
                                        })) continue;
                                    vm.data.checkboxTdObject.query.push(tmpQuery[key][tmp.activeKey]);
                                    vm.data.checkboxTdObject.indexAddress[tmpQuery[key][tmp.activeKey]] = tmpTdObj.hasOwnProperty('activeValue') ? tmpTdObj.activeValue : (parseInt(key) + 1);
                                }
                            } else {
                                for (let key in tmpQuery) {
                                    if (tmpQuery[key][tmp.activeKey] === null) continue;
                                    vm.data.checkboxTdObject.query.push(tmpQuery[key][tmp.activeKey]);
                                    vm.data.checkboxTdObject.indexAddress[tmpQuery[key][tmp.activeKey]] = tmpTdObj.hasOwnProperty('activeValue') ? tmpTdObj.activeValue : (parseInt(key) + 1);
                                }
                            }
                        }
                    }
                    if (vm.mainObject.baseFun.clickCheckbox) {
                        vm.mainObject.baseFun.clickCheckbox('plus-all', {
                            oldLength: tmpOldListLength,
                            currentLenght: vm.data.checkboxTdObject.query.length
                        });
                    }
                } else if (vm.mainObject.baseFun.cancelToSelectAll) {
                    vm.mainObject.baseFun.cancelToSelectAll();
                } else if (vm.mainObject.baseFun.clickCheckbox) {
                    vm.mainObject.baseFun.clickCheckbox('minus-all', {
                        oldLength: tmpOldListLength
                    });
                }
                if (vm.mainObject.baseFun.teardownWhenCheckboxIsClick) {
                    vm.mainObject.baseFun.teardownWhenCheckboxIsClick(vm.data.checkboxTdObject, vm.list);
                }
            }
        }
        fun.getLastItemIndex = function (inputIndex, inputArray) {
            var key = inputIndex + 1;
            while (key < inputArray.length) {
                if ((inputArray[inputIndex].listDepth || 0) >= (inputArray[key].listDepth || 0)) {
                    return key;
                }
                key++;
            }
            return key;
        }
        fun.checkIsLastItem = function (inputIndex, inputArray) {
            var key = inputIndex + 1;
            while (key < inputArray.length) {
                if (inputArray[inputIndex].listDepth === inputArray[key].listDepth) {
                    return false;
                } else if (inputArray[inputIndex].listDepth > inputArray[key].listDepth) {
                    return key;
                }
                key++;
            }
            return key;
        }
        vm.fun.watchFormLastChange = function (inputArg, callback) {
            if (!vm.mainObject.setting.munalAddRow && !inputArg.item.cancelAutomaticAddRow) {
                if (vm.data.isDepth) {
                    if (!(vm.mainObject.setting.munalHideOperateColumn && inputArg.$index === 0)) {
                        var tmpIndex = fun.checkIsLastItem(inputArg.$index, vm.list);
                        if (tmpIndex !== false && (!vm.mainObject.setting.illegalAutomaticAddRowModelKey || (vm.mainObject.setting.illegalAutomaticAddRowModelKey && !inputArg.item.hasOwnProperty(vm.mainObject.setting.illegalAutomaticAddRowModelKey)))) {
                            vm.list.splice(tmpIndex, 0, Object.assign({}, {
                                listDepth: inputArg.item.listDepth
                            }, vm.mainObject.itemStructure));
                        }
                    }
                } else if (inputArg.$index === vm.list.length - 1) {
                    vm.list.splice(inputArg.$index + 1, 0, Object.assign({}, {
                        listDepth: inputArg.item.listDepth
                    }, vm.mainObject.itemStructure));
                }
            }
            if (vm.mainObject.baseFun.watchFormLastChange) {
                vm.mainObject.baseFun.watchFormLastChange(inputArg);
            }
            if (callback) {
                callback(inputArg);
            }
        }
        $scope.importFile = function (inputArg) {
            inputArg.$index = this.$parent.$index;
            vm.mainObject.baseFun.importFile(inputArg);
        }
        vm.fun.shrinkList = function ($event) {
            $event.stopPropagation();
            var tmp = {};
            tmp.targetDom = fun.getTargetEvent($event.target);
            tmp.itemIndex = fun.getTargetIndex($event.target);
            vm.list[tmp.itemIndex].isShrink = !vm.list[tmp.itemIndex].isShrink;
            fun.operateLevel(tmp.targetDom.getAttribute('eo-attr-depth'), tmp.targetDom.nextElementSibling, parseInt(tmp.itemIndex) + 1);
        }
        vm.fun.range = function (inputLength, inputObject) {
            inputLength = inputLength || 1;
            if (!vm.list[inputObject.$index + 1] || ((vm.list[inputObject.$index + 1].listDepth || 0) <= (inputObject.item.listDepth || 0))) inputLength--;
            return new Array(inputLength);
        };
        vm.fun.sortMouseDown = function ($event, inputIndex) {
            if (vm.mainObject.setting.unsortableVar && vm.otherObject && vm.otherObject[vm.mainObject.setting.unsortableVar]) return;
            data.mouseEventElem = angular.element($event.target);
            data.mouseEventElem.bind('mousemove', function () {
                vm.data.movePart = inputIndex;
            })
        }
        vm.fun.mouseUp = function () {
            if (data.mouseEventElem) data.mouseEventElem.unbind('mousemove');
            data.movePart = vm.data.movePart;
            vm.data.movePart = null;
        }
        fun.operateLevel = function (inputDepth, $event, inputIndex) {
            var tmp = {
                    operateName: angular.element($event).hasClass('ng-hide') ? 'removeClass' : 'addClass'
                },
                tmpParentIsShrinkIndex = inputIndex,
                itemIndex = inputIndex;
            while ($event && inputDepth < $event.getAttribute('eo-attr-depth')) {
                switch (tmp.operateName) {
                    case 'addClass': {
                        vm.list[itemIndex].isHide = true;
                        break;
                    }
                    case 'removeClass': {
                        var tmpParentShrinkObject = vm.list[tmpParentIsShrinkIndex];
                        if (vm.list[itemIndex].isShrink && vm.list[itemIndex].listDepth <= tmpParentShrinkObject.listDepth) {
                            vm.list[itemIndex].isHide = false;
                            tmpParentIsShrinkIndex = itemIndex;
                        } else if (vm.list[itemIndex].listDepth <= tmpParentShrinkObject.listDepth) {
                            vm.list[itemIndex].isHide = false;
                            tmpParentIsShrinkIndex = itemIndex;
                        } else if (!tmpParentShrinkObject.isShrink) {
                            vm.list[itemIndex].isHide = false;
                        }
                        break;
                    }
                }
                itemIndex++;
                $event = $event.nextElementSibling;
            }
        }
        fun.parseFloatBtnGroupHtml = function (inputWhich, inputIndex, inputArray) {
            var tmpOutputHtml = '';
            if (inputArray) {
                tmpOutputHtml += '<div class="float-btngroup-tbd float-btngroup-' + inputWhich + '-tbd">';
                for (var btnKey in inputArray) {
                    var btnVal = inputArray[btnKey];
                    tmpOutputHtml += '<button type="button" class="fbtn-' + btnVal.operateName + ' float-btn-lbt ' + (btnVal.class || '') + '" ' + (btnVal.itemExpression || '') + ' eo-attr-btn-index="' + btnKey + '" eo-attr-td-index="' + inputIndex + '">' + (btnVal.key || btnVal.html) + '</button>';
                }
                tmpOutputHtml += '</div>';
            }
            return tmpOutputHtml;
        }
        fun.initItemHtml = function (inputVal, inputKey) {
            var tmpHtml = '',
                tmpThHtml = '';
            switch (inputVal.type) {
                case 'depthText': {
                    vm.data.isDepth = true;
                    tmpThHtml += '<div class="plr5 {{class}}">' + inputVal.thKey + '</div>';
                    tmpHtml += '<div class="td-tbd text-td-tbd plr5 {{class}}">' +
                        `<div class="depth-td-tbd" ng-style="{'padding-left':(15*item.listDepth+($ctrl.data.shrinkBtnLength?25:0))+'px'}" ng-init="item.listDepthArray=$ctrl.fun.range(item.listDepth+1,{item:item,$index:$outerIndex})">` +
                        '<button type="button" class="btn-shrink iconfont" ng-click="$ctrl.fun.shrinkList($event)" ng-class="{\'icon-pinleizengjia\':item.isShrink,\'icon-pinleijianshao\':!item.isShrink}" ng-if="$ctrl.list[$index+1].listDepth>item.listDepth"></button>' +
                        '<span class="divide-td-tbd" ng-class="{\'first-divide-td-tbd\':item.listDepth==$index}" ng-repeat="key in item.listDepthArray track by $index" ng-style="{\'left\':(15*$index+30)+\'px\'}" ng-hide="item.isShrink&&item.listDepth==$index"></span>' +
                        '<span>{{item.' + inputVal.modelKey + '}}</span>' +
                        '</div>' +
                        '</div>';
                    break;
                }
                case 'depthHtml': {
                    vm.data.isDepth = true;
                    tmpThHtml += '<div class="plr5 {{class}}">' + inputVal.thKey + '</div>';
                    tmpHtml += '<div class="td-tbd text-td-tbd plr5 {{class}}">' +
                        `<div class="depth-td-tbd" ng-style="{'padding-left':(15*item.listDepth+($ctrl.data.shrinkBtnLength?25:0))+'px'}" ng-init="item.listDepthArray=$ctrl.fun.range(item.listDepth+1,{item:item,$index:$outerIndex})">` +
                        '<button type="button" class="btn-shrink iconfont" ng-click="$ctrl.fun.shrinkList($event)" ng-class="{\'icon-pinleizengjia\':item.isShrink,\'icon-pinleijianshao\':!item.isShrink}" ng-if="$ctrl.list[$index+1].listDepth>item.listDepth"></button>' +
                        '<span class="divide-td-tbd" ng-class="{\'first-divide-td-tbd\':item.listDepth==$index}" ng-repeat="key in item.listDepthArray track by $index" ng-style="{\'left\':(15*$index+30)+\'px\'}" ng-hide="item.isShrink&&item.listDepth==$index"></span>' +
                        inputVal.html +
                        '</div>' +
                        '</div>';
                    break;
                }
                case 'depthInput': {
                    vm.data.isEditTable = true;
                    vm.data.isDepth = true;
                    tmpThHtml += '<div class="plr5 {{class}}">' + inputVal.thKey + '</div>';
                    tmpHtml += '<div class="td-tbd va-top-td-tbd depth-td-tdb plr5 {{class}}">' +
                        `<div class="depth-td-tbd" ng-style="{'padding-left':(15*item.listDepth+($ctrl.data.shrinkBtnLength?25:0))+'px'}">` +
                        '<button type="button" class="btn-shrink iconfont" ng-click="$ctrl.fun.shrinkList($event)" ng-class="{\'icon-pinleizengjia\':item.isShrink,\'icon-pinleijianshao\':!item.isShrink}" ng-if="$ctrl.list[$index+1].listDepth>item.listDepth"></button>' +
                        '<span class="divide-td-tbd" ng-class="{\'first-divide-td-tbd\':item.listDepth==$index}" ng-repeat="key in $ctrl.fun.range(item.listDepth+1,{item:item,$index:$outerIndex}) track by $index" ng-style="{\'left\':(15*$index+30)+\'px\'}" ng-hide="item.isShrink&&item.listDepth==$index"></span>' +
                        '<input autocomplete="off" ' + (inputVal.itemExpression || '') + ' type="text" class="eo-input" ng-model="item.' + inputVal.modelKey + '" ng-change="$ctrl.fun.watchFormLastChange({item:item,$index:$index})" {{placeholder}}>' +
                        '<p class="eo-error-tips">' + (inputVal.errorTip || ('请填写' + inputVal.thKey)) + '</p>' +
                        '</div>' +
                        fun.parseFloatBtnGroupHtml('input', inputKey, inputVal.btnList) +
                        '</div>';
                    break;
                }
                case 'html': {
                    tmpThHtml += '<div class="plr5 {{class}}" ' + (inputVal.itemExpression || '') + ' >' + inputVal.thKey + '</div>';
                    if ((typeof inputVal.html) === 'string') {
                        tmpHtml += '<div class="td-tbd text-td-tbd plr5 {{class}}" ' + (inputVal.itemExpression || '') + ' >' + inputVal.html + '</div>';
                    } else if (inputVal.html) {
                        tmpHtml = [];
                        for (var key in inputVal.html) {
                            tmpHtml.push('<div class="td-tbd text-td-tbd plr5 {{class}}">' + inputVal.html[key] + '</div>');
                        }
                    }
                    break;
                }
                case 'text': {
                    tmpThHtml += '<div class="plr5 {{class}}" ' + (inputVal.itemExpression || '') + ' >' + (inputVal.thKey || '') + '</div>';
                    if ((typeof inputVal.modelKey) === 'string') {
                        tmpHtml += '<div class="td-tbd text-td-tbd plr5 {{class}}" ' + (inputVal.itemExpression || '') + ' ' + (inputVal.title ? ('title="' + inputVal.title + '"') : '') + '>{{item.' + inputVal.modelKey + '}}</div>';
                    } else if (inputVal.modelKey) {
                        tmpHtml = [];
                        for (let key in inputVal.modelKey) {
                            tmpHtml.push('<div class="td-tbd text-td-tbd plr5 {{class}}" ' + (inputVal.itemExpression || '') + '>{{item.' + inputVal.modelKey[key] + '}}</div>');
                        }
                    }
                    break;
                }
                case 'sort': {
                    vm.data.sort = true;
                    vm.data.sortAuthorityVar = inputVal.authority || '';
                    tmpThHtml += '<div ' + (inputVal.itemExpression || '') + '   class="sort-handle-th" ' + (inputVal.authority ? ('ng-if="$ctrl.authorityObject.' + inputVal.authority + '"') : '') + '>' + (inputVal.thKey || '') + '</div>';
                    tmpHtml += '<div ' + (inputVal.itemExpression || '') + '   class="sort-handle-td" ' + (inputVal.authority ? ('ng-if="$ctrl.authorityObject.' + inputVal.authority + '"') : '') + '><div class="dp_ib" sv-group-handle  ' + (inputVal.itemHandleExpression || '') + '  ><span class="iconfont icon-jiantou_shangxiaqiehuan" sv-handle ' + (inputVal.isWantToPrepareWhenSort ? 'ng-mousedown="$ctrl.fun.sortMouseDown($event,$partIndex)"' : '') + '></span></div></div>';
                    break;
                }
                case 'radio': {
                    data.radioClickAffectTotalItem = inputVal.radioClickAffectTotalItem || false;
                    data.radioOriginalIndex = vm.mainObject.setting.radioOriginalType || 0;
                    data.radioTdIndex = inputKey;
                    tmpThHtml += '<div class="checkbox-th {{class}}">' + inputVal.thKey + '</div>';
                    tmpHtml += '<div class="checkbox-td td-tbd va-top-td-tbd {{class}}"><span class="input-radio eo-checkbox iconfont ' + (inputVal.thKey ? 'inline_cth' : 'block_cth') + '" eo-attr-td-index="' + inputKey + '">{{item.' + inputVal.modelKey + '?"&#xeb14;":""}}</span></div>';
                    break;
                }
                case 'relationalCheckbox': { //关系型checkbox
                    data.relationalCheckboxTdIndex = inputKey;
                    $rootScope.global.$watch.push($scope.$watch('$ctrl.list', fun.watchRelationalCheckboxChange, true));
                    let tmpAuthorityHtml = inputVal.authority ? ('ng-class="{\'disable-checkbox\':!$ctrl.authorityObject.' + inputVal.authority + '}" ') : '';
                    tmpThHtml += '<div class="checkbox-th {{class}}" ' + (inputVal.thItemExpression || '') + '  ><span class="eo-checkbox iconfont ' + (inputVal.thKey ? 'inline_cth' : 'block_cth') + '" ng-click="$ctrl.fun.selectAll(' + inputKey + ')" ' + (tmpAuthorityHtml || '') + '>{{$ctrl.mainObject.tdList[' + inputKey + '].selectAll?"&#xeb14;":"&nbsp;"}}</span>' + (inputVal.thKey ? ('<span class="desc-cth">' + inputVal.thKey + '</span>') : '') + '</div>';
                    tmpHtml += '<div class="checkbox-td td-tbd va-top-td-tbd {{class}}" ' + (inputVal.itemExpression || '') + ' ><span class="' + (inputVal.itemDisabledExpression ? ('{{' + inputVal.itemDisabledExpression + '}}') : 'relational-checkbox') + ' eo-checkbox iconfont ' + (inputVal.thKey ? 'inline_cth' : 'block_cth') + '" eo-attr-td-index="' + inputKey + '" ' + (tmpAuthorityHtml || '') + '>{{item.' + inputVal.modelKey + '?"&#xeb14;":""}}</span></div>';
                    break;
                }
                case 'checkbox': {

                    data.checkboxClickAffectTotalItem = inputVal.checkboxClickAffectTotalItem || false;
                    data.checkboxTdIndex = inputKey;
                    $rootScope.global.$watch.push($scope.$watch('$ctrl.data.checkboxTdObject.isOperating', fun.watchCheckboxChange));
                    if (inputVal.wantToWatchListLength) $rootScope.global.$watch.push($scope.$watch('$ctrl.list.length', fun.watchCheckboxChange, true));
                    if (inputVal.modelKey) $rootScope.global.$watch.push($scope.$watch('$ctrl.list', fun.watchCheckboxChange, true));
                    if (inputVal.isWantedToExposeObject) { //是否希望绑定/暴露内置变量
                        vm.data.checkboxTdObject = vm.activeObject = Object.assign({}, vm.data.checkboxTdObject, vm.activeObject);
                    }
                    vm.data.checkboxTdObject.isOperating = vm.data.checkboxTdObject.hasOwnProperty('isOperating') ? vm.data.checkboxTdObject.isOperating : true;
                    var tmpAuthorityHtml = inputVal.authority ? ('ng-class="{\'disable-checkbox\':!$ctrl.authorityObject.' + inputVal.authority + '}" ') : '';
                    tmpThHtml += '<div class="checkbox-th {{class}}" ' + (inputVal.thItemExpression || '') + '  ><span class="eo-checkbox iconfont ' + (inputVal.thKey ? 'inline_cth' : 'block_cth') + '" ng-click="$ctrl.fun.selectAll(' + inputKey + ')" ' + (tmpAuthorityHtml || '') + '>{{$ctrl.data.checkboxTdObject.selectAll?"&#xeb14;":"&nbsp;"}}</span>' + (inputVal.thKey ? ('<span class="desc-cth">' + inputVal.thKey + '</span>') : '') + '</div>';
                    tmpHtml += '<div class="checkbox-td td-tbd va-top-td-tbd {{class}}" ' + (inputVal.itemExpression || '') + ' ><span class="' + (inputVal.itemDisabledExpression ? ('{{' + inputVal.itemDisabledExpression + '}}') : 'input-checkbox') + ' eo-checkbox iconfont ' + (inputVal.thKey ? 'inline_cth' : 'block_cth') + '" eo-attr-td-index="' + inputKey + '" ' + (tmpAuthorityHtml || '') + '>' + (inputVal.modelKey ? ('{{item.' + inputVal.modelKey + '?"&#xeb14;":""}}') : ('{{$ctrl.data.checkboxTdObject.indexAddress[item.' + inputVal.activeKey + ']?"&#xeb14;":""}}')) + '</span></div>';
                    break;
                }
                case 'cbtn':
                case 'btn': {
                    tmpThHtml += '<div ' + (inputVal.itemExpression || '') + '   class="{{class}} plr5" ' + (inputVal.authority ? ('ng-if="$ctrl.authorityObject.' + inputVal.authority + '"') : '') + '>' + (inputVal.thKey || '操作') + '</div>';
                    tmpHtml += '<div ' + (inputVal.itemExpression || '') + '   class="operate-td-tbd va-top-td-tbd td-tbd {{class}} plr5" ' + (inputVal.authority ? ('ng-if="$ctrl.authorityObject.' + inputVal.authority + '"') : '') + '><div class="f_row" ng-hide="$last&&!$ctrl.mainObject.setting.munalHideOperateColumn&&$ctrl.data.isEditTable">';
                    for (var btnKey in inputVal.btnList) {
                        var btnVal = inputVal.btnList[btnKey];
                        var tmpBtnFunHtml = '';
                        switch (btnVal.type) {
                            case 'more': {
                                for (var btnFunKey in btnVal.funArr) {
                                    var btnFunVal = btnVal.funArr[btnFunKey];
                                    tmpBtnFunHtml += '<p><button type="button" class="btn-funItem" ng-mousedown="$ctrl.fun.moreItemClick($event,$partIndex)" eo-attr-btn-fun-index="' + btnFunKey + '">' + btnFunVal.key + '</button></p>'
                                }
                                tmpHtml += '<div class="more-btn-container" ' + (btnVal.itemExpression || '') + ' ' + (btnVal.authority ? ('ng-if="$ctrl.authorityObject.' + btnVal.authority + '"') : '') + ' eo-attr-btn-index="' + btnKey + '" eo-attr-td-index="' + inputKey + '">' +
                                    '<button type="button" class="more-btn eo-operate-btn ' + (btnVal.class || '') + '" ' + (btnVal.btnItemExpression || '') + '><span>' + (btnVal.key || btnVal.html) + '</span><span class="iconfont icon-xuanzeqizhankai_o"></span></button>' +
                                    '<div class="more-div-btngroup-tbd">' + tmpBtnFunHtml + '</div></div>';
                                break;
                            }
                            case 'html': {
                                tmpHtml += '<div ' + (btnVal.class ? ('class="' + btnVal.class + '"') : '') + ' ' + (btnVal.itemExpression || '') + ' >' + btnVal.html + '</div>';
                                break;
                            }
                            default: {
                                tmpHtml += '<button type="button" ' + (btnVal.authority ? ('ng-if="$ctrl.authorityObject.' + btnVal.authority + '"') : '') + ' class="' + (inputVal.type === 'cbtn' ? 'c' : '') + 'btn-' + btnVal.operateName + ' eo-operate-btn ' + (btnVal.class || '') + '" ' + (btnVal.itemExpression || '') + ' eo-attr-btn-index="' + btnKey + '" eo-attr-td-index="' + inputKey + '">' + (btnVal.key || btnVal.html) + '</button>';
                                break;
                            }
                        }
                    }
                    tmpHtml += '</div></div>';
                    break;
                }
                case 'selectMulti': {
                    vm.data.isEditTable = true;
                    tmpThHtml += '<div class="plr5 {{class}}" ' + (inputVal.thItemExpression || '') + '>' + inputVal.thKey + '</div>';
                    tmpHtml += '<div class="td-tbd plr5 va-top-td-tbd {{class}}" ' + (inputVal.itemExpression || '') + '><select-multi-common-component  model-arr="item.' + inputVal.modelKey + '" main-object="$ctrl.mainObject.tdList[' + inputKey + '].selectMultiObject" list="$ctrl.mainObject.tdList[' + inputKey + '].selectQuery"></select-multi-common-component></div>';
                    break;
                }
                case 'select': {
                    vm.data.isEditTable = true;
                    tmpThHtml += '<div class="plr5 {{class}}" ' + (inputVal.thItemExpression || '') + '>' + inputVal.thKey + '</div>';
                    tmpHtml += '<div class="td-tbd plr5 va-top-td-tbd {{class}}" ' + (inputVal.itemExpression || '') + '><select-default-common-component output="item" ' + (vm.mainObject.setting.readonly ? 'disabled=true' : (inputVal.disabled ? ('disabled="' + inputVal.disabled + '"') : '')) + ' input="{query:' + (inputVal.selectQueryExpression || '$ctrl.mainObject.tdList[' + inputKey + '].selectQuery') + ',key:\'' + (inputVal.key || 'key') + '\', value:\'' + (inputVal.value || 'value') + '\'' + (inputVal.initialData ? (',initialData:' + inputVal.initialData) : '') + '}" model-key="' + inputVal.modelKey + '" input-change-fun="$ctrl.fun.watchFormLastChange({item:item,$index:$index},$ctrl.mainObject.tdList[' + inputKey + '].fun)" required=true></select-default-common-component></div>';
                    break;
                }
                case 'input': {
                    vm.data.isEditTable = true;
                    tmpThHtml += '<div class="plr5 {{class}}" ' + (inputVal.tdItemExpression || '') + ' >' + inputVal.thKey + '</div>';
                    tmpHtml += '<div class="td-tbd input-tbd va-top-td-tbd plr5 {{class}}" ' + (inputVal.tdItemExpression || '') + ' >' +
                        '<input ng-readonly="$ctrl.mainObject.setting.readonly" autocomplete="off" type="text" ' + (inputVal.itemExpression || '') + '  class="eo-input" ng-model="item.' + inputVal.modelKey + '" ng-change="' + (inputVal.changeFun ? ('$ctrl.mainObject.tdList[' + inputKey + '].changeFun({item:item,$index:$index},$ctrl.fun.watchFormLastChange)"') : '$ctrl.fun.watchFormLastChange({item:item,$index:$index})"') + ' {{placeholder}}>' +
                        '<p class="eo-error-tips">' + (inputVal.errorTip || ('请填写' + inputVal.thKey)) + '</p>' +
                        fun.parseFloatBtnGroupHtml('input', inputKey, inputVal.btnList) +
                        '</div>';
                    break;
                }
                case 'autoComplete': {
                    vm.data.isEditTable = true;
                    tmpThHtml += '<div class="plr5 {{class}}">' + inputVal.thKey + '</div>';
                    tmpHtml += '<div class="td-tbd acp-tbd va-top-td-tbd plr5 {{class}}" ' + (inputVal.itemExpression || '') + '>' +
                        '<auto-complete-ams-component model="item" key-name="' + inputVal.modelKey + '" array="' + ((typeof inputVal.selectQuery) === 'string' ? inputVal.selectQuery : ('$ctrl.mainObject.tdList[' + inputKey + '].selectQuery')) + '" input-change-fun="$ctrl.fun.watchFormLastChange({item:item,$index:$index})" {{placeholder}}></auto-complete-ams-component>' +
                        '<p class="eo-error-tips">' + (inputVal.errorTip || ('请填写' + inputVal.thKey)) + '</p>' +
                        fun.parseFloatBtnGroupHtml('acp', inputKey, inputVal.btnList) +
                        '</div>';
                    break;
                }
                case 'autoCompleteAndFile': {
                    vm.data.isEditTable = true;
                    let tmpFileInputHtml = "",
                        tmpFilePlaceholder = inputVal.filePlaceholder || '请选择文件',
                        tmpFileBtnText = inputVal.fileBtnText || '选择文件';
                    if (inputVal.munalDefineFileFun) {
                        tmpFileInputHtml = '<input autocomplete="off" class="eo-input text-input" type="text" ng-model="item.' + inputVal.modelKey + '" disabled="true" placeholder="' + tmpFilePlaceholder + '">' +
                            '<button type="button" class="file-btn-lbt" ng-click="importFile({item:item})">' + tmpFileBtnText + '</button>';
                    } else {
                        tmpFileInputHtml = '<input autocomplete="off" class="eo-input text-input" type="text" ng-model="item.' + inputVal.modelKey + '" disabled="true" placeholder="' + tmpFilePlaceholder + '">' +
                            '<input autocomplete="off" type="file" class="file-input" onchange="angular.element(this).scope().importFile({file:this.files})" multiple="multiple">' +
                            '<button type="button" class="file-btn-lbt">' + tmpFileBtnText + '</button>';
                    }
                    tmpThHtml += '<div class="plr5 {{class}}">' + inputVal.thKey + '</div>';
                    tmpHtml += '<div class="td-tbd plr5 acp-and-file-tbd va-top-td-tbd {{class}}" ' + (inputVal.itemExpression || '') + ' ng-switch="item.' + inputVal.switchVar + '">' +
                        '<div class="file-div" ng-switch-when="' + inputVal.swicthFile + '">' +
                        tmpFileInputHtml + '</div>' +
                        '<div ng-switch-default><auto-complete-ams-component model="item" key-name="' + inputVal.modelKey + '" array="' + ((typeof inputVal.selectQuery) === 'string' ? inputVal.selectQuery : ('$ctrl.mainObject.tdList[' + inputKey + '].selectQuery')) + '" input-change-fun="$ctrl.fun.watchFormLastChange({item:item,$index:$index})"  {{placeholder}}></auto-complete-ams-component>' +
                        fun.parseFloatBtnGroupHtml('acp', inputKey, inputVal.btnList) +
                        '</div></div>';
                    break;
                }
            }
            return {
                thHtml: tmpThHtml,
                tdHtml: tmpHtml
            }
        }
        fun.initPartHtml = function () {
            var tmp = {
                html: new Array(vm.mainObject.setting.partNum),
                thHtml: ''
            }
            for (var partIndex = 0; partIndex < vm.mainObject.setting.partNum; partIndex++) {
                tmp.html[partIndex] = '<div ng-repeat="($outerIndex,item) in $ctrl.list.' + vm.mainObject.setting.partModule[partIndex] + '" ng-hide="item.isHide&&$ctrl.data.isDepth" mouse-up="$ctrl.fun.mouseUp" sv-group-element="$ctrl.data.sortForm" eo-attr-index="{{$index}}" eo-attr-depth="{{item.listDepth}}"  {{trExpression}}><div class="tr-tbd {{trClass}}" {{trDirective}} {{trNgClass}}>';
                try {
                    tmp.html[partIndex] = tmp.html[partIndex].replace('{{trClass}}', vm.mainObject.setting.trClass || '').replace('{{trNgClass}}', vm.mainObject.setting.trNgClass || '').replace('{{trDirective}}', vm.mainObject.setting.trDirective || '').replace('{{trExpression}}', vm.mainObject.setting.trExpression || '');
                } catch (e) {
                    console.error(e)
                }
            }
            for (var key in vm.mainObject.tdList) {
                var val = vm.mainObject.tdList[key];
                var tmpHtmlObject = fun.initItemHtml(val, key);
                tmp.thHtml += tmpHtmlObject.thHtml.replace('{{class}}', val.class || '');
                for (let key in tmp.html) {
                    if (typeof (tmpHtmlObject.tdHtml) === 'string') {
                        tmp.html[key] += tmpHtmlObject.tdHtml.replace('{{class}}', val.class || '');
                    } else {
                        tmp.html[key] += tmpHtmlObject.tdHtml[key].replace('{{class}}', val.class || '');
                    }
                }
            }
            for (let key in tmp.html) {
                vm.data.partHtml[vm.mainObject.setting.partModule[key]] = tmp.html[key] + '</div></div>'
            }
            vm.data.thHtml = tmp.thHtml;
        }
        fun.initHtml = function () {
            var tmp = {
                    html: '',
                    thHtml: ''
                },
                tmpStaticHtml = '<div class="tr-tbd {{trClass}}" {{trDirective}} {{trNgClass}} >';
            try {
                tmpStaticHtml = tmpStaticHtml.replace('{{trClass}}', vm.mainObject.setting.trClass || '').replace('{{trNgClass}}', vm.mainObject.setting.trNgClass || '').replace('{{trDirective}}', vm.mainObject.setting.trDirective || '');
            } catch (REPLACE_ERR) {
                console.error(REPLACE_ERR)
            }
            for (var key in vm.mainObject.tdList) {
                var val = vm.mainObject.tdList[key];
                let tmpHtmlObject = fun.initItemHtml(val, key);
                tmp.thHtml += tmpHtmlObject.thHtml.replace('{{class}}', val.class || '');
                tmpStaticHtml += tmpHtmlObject.tdHtml.replace('{{class}}', val.class || '').replace('{{placeholder}}', val.placeholder ? ('placeholder="' + val.placeholder + '"') : '');
            }
            tmp.html = (vm.mainObject.setting.isForm ? '<ng-form name="ListBlockCommonComponentForm">' : `<div get-Dom-Length-Common-Directive ${vm.mainObject.setting.tbodyExpression||""} bind-Class="btn-shrink" model="$ctrl.data.shrinkBtnLength">`) + `<div class="tr_container_tbd" ng-repeat="($outerIndex,item) in $ctrl.list track by $index" ${vm.data.isDepth?'ng-hide="item.isHide"':''} sv-group-element="$ctrl.data.sortForm" eo-attr-index="{{$index}}" eo-attr-depth="{{item.listDepth}}"  {{trExpression}}>`;
            try {
                tmp.html = tmp.html.replace('{{trExpression}}', vm.mainObject.setting.trExpression || '');
            } catch (REPLACE_ERR) {
                console.error(REPLACE_ERR)
            }
            vm.data.html = tmp.html + (vm.mainObject.extraTrHtml || "") + tmpStaticHtml + '</div></div>' + (vm.mainObject.setting.isForm ? '</ng-form>' : '</div>');
            vm.data.thHtml = tmp.thHtml;
        }
        vm.$onInit = function () {
            if (vm.mainObject) {
                vm.mainObject.setting = vm.mainObject.setting || {};
                vm.mainObject.baseFun = vm.mainObject.baseFun || {};
                if (vm.mainObject.setting.isPartModule) {
                    fun.initPartHtml();
                } else {
                    fun.initHtml();
                }

            }
        }
        fun.watchRelationalCheckboxChange = function () {
            if ((vm.list || []).length <= 0) return;
            var tmpModelItem = vm.mainObject.tdList[data.relationalCheckboxTdIndex],
                tmpModelKey = tmpModelItem.modelKey;
            data.queryLength = 0;
            for (var key in vm.list) {
                if (vm.list[key][tmpModelKey]) {
                    data.queryLength++;
                }
            }
            if ((vm.list || []).length === data.queryLength) {
                tmpModelItem.selectAll = true;
            } else {
                tmpModelItem.selectAll = false;
            }
        }
        fun.watchCheckboxChange = function () {
            if ((vm.list || []).length <= 0 && !vm.mainObject.setting.isPartModule) return;
            if (vm.data.checkboxTdObject.isOperating) {
                var tmpModelKey = vm.mainObject.tdList[data.checkboxTdIndex].modelKey;
                data.queryLength = 0;
                if (tmpModelKey) {
                    for (var key in vm.list) {
                        if (vm.list[key][tmpModelKey]) {
                            data.queryLength++;
                        }
                    }
                    if ((vm.list || []).length === data.queryLength) {
                        vm.data.checkboxTdObject.selectAll = true;
                    } else {
                        vm.data.checkboxTdObject.selectAll = false;
                    }
                } else {
                    if (vm.mainObject.setting.isPartModule) {
                        for (let key in vm.list) {
                            data.queryLength += (vm.list[key] || []).length;
                        }
                    } else {
                        data.queryLength = (vm.list || []).length;
                    }
                    vm.data.checkboxTdObject.query = [];
                    for (let key in vm.data.checkboxTdObject.indexAddress) {
                        vm.data.checkboxTdObject.query.push(vm.mainObject.setting.checkboxKeyIsNum ? parseInt(key) : key);
                    }
                    if ((vm.data.checkboxTdObject.query || []).length >= data.queryLength) {
                        vm.data.checkboxTdObject.selectAll = true;
                    } else {
                        vm.data.checkboxTdObject.selectAll = false;
                    }
                }
            }
        }
    }
})();