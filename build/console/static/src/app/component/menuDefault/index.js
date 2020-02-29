(function () {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 工具栏组件
     * @extend {object} authorityObject 权限类{operate}
     * @extend {object} activeObject 活动/聚焦标志
     * @extend {object} showObject 显示标志
     * @extend {array} mainObject item,
     * setting{object}
     * 高度：menuSize：sm |md | lg 默认（sm）
     * 标题:title,titleAuthority
     * 背景：白/灰  background   根据type: seperate| listTop
     * 批量操作: batch,batchAuthority
     * @extend {object} otherObject 不可预期辅助类
     * @description scar:新增tip-li的authority
     */
    angular.module('eolinker')
        .component('menuDefaultCommonComponent', {
            templateUrl: 'app/component/menuDefault/index.html',
            controller: indexController,
            bindings: {
                authorityObject: '<',
                otherObject: '=',
                activeObject: '<',
                showObject: '<',
                mainObject: '<',
                blockListObject:'='
            }
        })

    indexController.$inject = [];

    function indexController() {
        var vm = this;
        vm.fun = {
            common: null
        };
        vm.component={
            blockListObject:{
                listItem:{
                    baseFun:{
                        teardownWhenCheckboxIsClick:(inputCheckboxObject,inputList)=>{
                            let tmpIndexAddress={};
                            inputList.map((val)=>{
                                if(inputCheckboxObject.indexAddress.hasOwnProperty(val.value)){
                                    tmpIndexAddress[val.value]=inputCheckboxObject.indexAddress[val.value];
                                }else{
                                    tmpIndexAddress[val.value]=0;
                                }
                            })
                            window.localStorage.setItem(vm.mainObject.setting.listItemStorageName, JSON.stringify(tmpIndexAddress));
                        }
                    },
                    setting:{
                        trClass: 'hover-tr-lbcc'
                    },
                    tdList: [{
                        type: 'checkbox',
                        isWantedToExposeObject: true,
                        checkboxClickAffectTotalItem: true,
                        activeKey: 'value',
                        activeValue: 1
                    },{
                        thKey: '列表项',
                        type: 'text',
                        modelKey: 'key'
                    }]
                }
            }
        }
        /**
         * @description 统筹绑定调用页面列表功能单击函数
         * @param {extend} obejct 方式值
         * @param {object} arg 共用体变量，后根据传值函数回调方法
         */
        vm.fun.common = function (extend, arg) {
            if (!extend) return;
            var template = {
                params: arg
            }
            switch (typeof (extend.params)) {
                case 'string':
                    {
                        return eval('extend.default(' + extend.params + ')');
                    }
                default:
                    {
                        for (var key in extend.params) {
                            if (extend.params[key] == null) {
                                template.params[key] = arg[key];
                            } else {
                                template.params[key] = extend.params[key];
                            }
                        }
                        return extend.default(template.params);
                    }
            }

        }
        vm.fun.search = function (arg) {
            arg.item.keyword = arg.item.keyword || "";
            if (arg.$event) {
                arg.$event.stopPropagation();
            }
            arg.item.tapShow = true;
            arg.item.fun(arg);
        }
        vm.fun.batchCancel = function () {
            vm.mainObject.setting.batchInitFun();
        }
        vm.fun.batchDefault = function () {
            vm.otherObject.batch.isOperating = true;
            if(vm.mainObject.baseFun&&vm.mainObject.baseFun.batchDefault){
                vm.mainObject.baseFun.batchDefault();
            }
        }
        vm.$onInit = function () {
            if (vm.mainObject.setting.batch) {
                vm.mainObject.setting.batchInitFun();
            }
        }
    }
})();