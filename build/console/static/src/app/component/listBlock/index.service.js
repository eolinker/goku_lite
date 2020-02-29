(function () {
    'use strict';
    /*
     * @author 广州银云信息科技有限公司
     * @description 缓存公用服务js
     */
    angular.module('eolinker')
        .factory('ListBlock_CommonComponentService', index);

    index.$inject = []

    function index() {
        var publicFun = {}, privateFun = {};
        privateFun.deepCopy = function (inputObject) {
            return JSON.parse(JSON.stringify(inputObject || []));
        }
        privateFun.loopGenerateList = function (inputArray, options) {
            var tmp = {
                output: []
            }
            try {
                for (var key in inputArray) {
                    var val = inputArray[key];
                    if (options.fun) {
                        var tmpVal = options.fun(val, options.listDepth);
                        if (tmpVal) {
                            switch (Object.prototype.toString.call(tmpVal)) {
                                case '[object Array]': {
                                    tmp.output = tmp.output.concat(tmpVal);
                                    continue;
                                }
                                default: {
                                    val = tmpVal;
                                    break;
                                }
                            }
                        }
                    }
                    if (!options.munalConstructListDepth) val.listDepth = options.listDepth;
                    tmp.output.push(val);
                    if (val.childList && val.childList.length > 0) {
                        tmp.output = tmp.output.concat(privateFun.loopGenerateList(privateFun.deepCopy(val.childList), Object.assign({}, options, {
                            listDepth: options.listDepth + 1
                        })));
                        delete val.childList;
                    }
                }
            } catch (UNDEFINED_ERROR) {
                console.error(UNDEFINED_ERROR)
            }

            return tmp.output;
        }
        publicFun.initReadonlyTableList = function (inputArray, options) {
            options = options || {};
            var tmpOutput = privateFun.deepCopy(inputArray);
            options.listDepth = 0;
            tmpOutput = privateFun.loopGenerateList(tmpOutput, options);
            return tmpOutput;
        }
        privateFun.organizeLevelAsChildList = function (inputArray, options, inputParent) {
            options = options || {};
            let tmp = {
                outList: [],
                length: 0
            }, tmpIsNeedToBreakLoop = false;
            for (; options.index < inputArray.length;) {
                var val = inputArray[options.index];
                if (options.fun) {
                    tmp.optionFunStatus = options.fun(val, options.index, inputParent);
                    switch (tmp.optionFunStatus) {
                        case false: {
                            return;
                        }
                        case null: {
                            options.index++;
                        }
                        default: {
                            if (typeof (tmp.optionFunStatus) === 'object' && tmp.optionFunStatus) {
                                options.index = tmp.optionFunStatus.currentIndex;
                                tmp.outList.push(tmp.optionFunStatus.newItem);
                                if ((options.index == inputArray.length) || (val.listDepth > inputArray[options.index].listDepth)) {
                                    tmpIsNeedToBreakLoop = true;
                                    break;
                                } else {
                                    continue;
                                }
                            }
                        }
                    }
                }
                if (tmpIsNeedToBreakLoop) break;
                if (tmp.optionFunStatus !== null) {
                    options.index++;
                    tmp.outList.push(val);
                }
                if ((options.index === inputArray.length) || (val.listDepth > inputArray[options.index].listDepth)) {
                    break;
                }
                if (val.listDepth < inputArray[options.index].listDepth) {
                    val.childList = privateFun.organizeLevelAsChildList(angular.copy(inputArray), options, val);
                    if (!val.childList) {
                        return;
                    } else if (inputArray[options.index] && val.listDepth > inputArray[options.index].listDepth) {
                        break;
                    }
                }

            }
            return tmp.outList;
        }
        publicFun.formatNestList = function (inputArray, options) {
            options = options || {};
            options.index = 0;
            var tmpOutput = privateFun.deepCopy(inputArray);
            if (tmpOutput) {
                return privateFun.organizeLevelAsChildList(tmpOutput, options);
            } else {
                return tmpOutput;
            }
        }
        publicFun.initEditTableList = function (inputArray, options) {
            options = options || {};
            var tmpOutput = privateFun.deepCopy(inputArray);
            if (options.fun || options.isLoop) {
                options.listDepth = options.listDepth || 0;
                tmpOutput = privateFun.loopGenerateList(tmpOutput, options);
            }
            if (!options.munalHideOperateColumn) {
                if (options.lastFilterKey && tmpOutput.length > 0 && !tmpOutput[tmpOutput.length - 1][options.lastFilterKey]) return tmpOutput;
                tmpOutput.push(Object.assign({}, {
                    listDepth: 0
                }, options.itemStructure));
            }
            return tmpOutput;
        }
        return publicFun;
    }
})();
