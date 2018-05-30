(function() {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 列表组件---依赖侧边（eg：分组）
     * @extend {object} authorityObject 权限类{operate}
     * @extend {object} activeObject 活动/聚焦标志
     * @extend {object} showObject 显示标志
     * @extends {array} list 列表
     * @extend {object} otherObject 不可预期辅助类
     */
    angular.module('goku')
        .component('listRequireCommonComponent', {
            templateUrl: 'app/component/common/list/require/index.html',
            controller: indexController,
            bindings: {
                authorityObject: '<',
                otherObject: '<',
                activeObject: '<',
                showObject: '<',
                list: '<',
                mainObject:'<'
            }
        })

    indexController.$inject = ['$scope', '$compile'];

    function indexController($scope, $compile) {
        var vm = this;
        vm.data = {
            info:{
                btnList:null,
                html:null
            },
            fun: {
                common: null
            }
        }

        var assistantFun={};
        /**
         * @description 统筹绑定调用页面列表功能单击函数
         * @param {extend} obejct 方式值
         * @param {object} arg 共用体变量，后根据传值函数回调方法
         */
        vm.data.fun.common = function(extend, arg) {
            if(!extend)return;
            var template = {
                params: arg
            }
            switch (typeof(extend.params)) {
                case 'string':
                    {
                        return eval('extend.default(' + extend.params + ')');
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
                        return extend.default(template.params);
                        break;
                    }
            }
        }

        assistantFun.initHtml=function(){
            if(!vm.mainObject.tdList)return;
            var template={
                btnString:'',
                html:'<tr class="{{$ctrl.mainObject.baseInfo.class}}" ng-repeat="($outerIndex,item) in $ctrl.list|filter:$ctrl.mainObject.fun.filter" ng-click="$ctrl.mainObject.fun.click({item:item,$index:$index})" ng-class="{\'elem-active\':item.isClick&&$ctrl.activeObject[$ctrl.mainObject.baseInfo.active],\'hover-tr\':$ctrl.activeObject[$ctrl.mainObject.baseInfo.active]}">'
            }
            vm.mainObject.tdList.map(function(val,key){
                switch(val.keyType){
                    case 'customized-html':{
                        template.html+='<td '+(val.showVariable?'ng-show="$ctrl.showObject.'+val.showVariable+'==$ctrl.mainObject.tdList['+key+'].show" ':' ')+(val.authority?'ng-if="$ctrl.authorityObject[$ctrl.mainObject.tdList['+key+'].authority]"':'')+' >'+val.keyHtml+'</td>';
                        break;
                    }
                    case 'btn':{
                        vm.data.info.btnList=val.btnList;
                        val.btnList.map(function(childVal,childKey){
                            template.btnString+='<a class="btn-a" ng-click="$ctrl.data.fun.common($ctrl.data.info.btnList['+childKey+'].fun,{item:item,$index:$outerIndex,$event:$event})" ng-show="$ctrl.data.info.btnList['+childKey+'].show==-1||($ctrl.showObject.'+val.showPoint+'.'+val.showVariable+'==$ctrl.data.info.btnList['+childKey+'].show)"><span class="iconfont icon-'+childVal.icon+'"></span>'+childVal.name+'</a>';
                        })
                        template.html+='<td '+(val.showVariable?'ng-show="$ctrl.showObject.'+val.showVariable+'==$ctrl.mainObject.tdList['+key+'].show" ':' ')+(val.authority?'ng-if="$ctrl.authorityObject[$ctrl.mainObject.tdList['+key+'].authority]"':'')+' >'+template.btnString+'</td>';
                        break;
                    }
                    default:{
                        template.html+='<td '+(val.showVariable?'ng-show="$ctrl.showObject.'+val.showVariable+'==$ctrl.mainObject.tdList['+key+'].show" ':' ')+(val.authority?'ng-if="$ctrl.authorityObject[$ctrl.mainObject.tdList['+key+'].authority]"':'')+' ><span>{{item.'+val.key+'}}</span></td>';
                        break;
                    }
                }
            });
            vm.data.info.html=template.html+'</tr>';
        }
        $scope.$watch('$ctrl.mainObject.tdList',assistantFun.initHtml);
    }
})();
