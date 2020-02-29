(function () {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 分组/表组件
     * @extend {object} authorityObject 权限类{edit}
     * @extend {object} funObject 第一部分功能集类{showVar,btnGroupList{edit:{fun,key,class,showable,icon,tips},sort:{default,cancel,confirm:{fun,key,showable,class,icon,tips}}}}
     * @extend {object} sortObject 排序信息{sortable,groupForm}
     * @extend {object} mainObject 主类{level,extend,query,baseInfo:{name,id,child,fun:{edit,delete},parentFun:{addChild}}}
     */
    angular.module('eolinker')
        .component('groupCommonComponent', {
            templateUrl: 'app/component/common/group/old/index.html',
            controller: indexController,
            bindings: {
                authorityObject: '<',
                funObject: '<',
                sortObject: '<',
                mainObject: '<',
                list: '<'
            }
        })

    indexController.$inject = ['$scope','$state'];

    function indexController($scope,$state) {
        var vm = this;
        vm.data = {
            info: {
                sortForm: {
                    containment: '.group-form-ul',
                    child: {
                        containment: '.child-group-form-ul'
                    },
                    grandson: {
                        containment: '.third-level-group-form-ul'
                    }
                }
            },
            fun: {
                more: null,
                common: null
            }
        }
        var fun={};
        /**
         * 单项列更多操作
         * @param  {object} arg 传参object{item:单列表项,$event:dom文本}
         */
        vm.data.fun.more = function (arg) {
            arg.$event.stopPropagation();
            arg.item.listIsClick = true;
        }

        /**
         * @description 统筹绑定调用页面列表功能单击函数
         * @param {extend} obejct 方式值
         * @param {object} arg 共用体变量，后根据传值函数回调方法
         */
        vm.data.fun.common = function (extend, arg) {
            if (!extend) return;
            var template = {
                params: arg
            }
            switch (typeof (extend.params)) {
                case 'string':
                    {
                        return eval('extend.fun(' + extend.params + ')');
                        break;
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
                        return extend.fun(template.params);
                        break;
                    }
            }

        }
        fun.show=function(){
            if(!vm.mainObject.showRouterList||vm.mainObject.showRouterList.indexOf($state.current.name)>-1){
                if(vm.mainObject.baseFun.init)vm.mainObject.baseFun.init();
                vm.data.info.show=true;
            }else{
                vm.data.info.show=false;
            }
        }
        vm.$onInit=function(){
            fun.show();
            $scope.$on('$stateChangeSuccess',function(){
                fun.show();
            })
        }
    }
})();