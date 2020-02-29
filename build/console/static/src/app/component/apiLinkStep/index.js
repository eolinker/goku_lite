(function () {
    'use strict';

    /**
     * loading加载组件
     * @param {function} fun 绑定方法
     * @param {object} interaction 交互参数
     */
    angular.module('eolinker')
        .component('apiLinkStepComponent', {
            templateUrl: 'app/component/apiLinkStep/index.html',
            controller: indexController,
            bindings: {
                mainObject:"<",
                list:"=",
                otherObject:"<"
            }
        })

    indexController.$inject = ['$rootScope'];

    function indexController($rootScope) {
        let vm=this;
        vm.component={
            ipListObj:{
                tdList:[{
                    type:"input",
                    modelKey:"ip"
                },{
                    type:"btn",
                    class:"w_50",
                    btnList:[{
                        key:"删除",
                        operateName:"delete"
                    }]
                }]
            },
            deleteActionObj:{
                tdList:[{
                    type:"input",
                    modelKey:"origin"
                },{
                    type:"btn",
                    class:"w_50",
                    btnList:[{
                        key:"删除",
                        operateName:"delete"
                    }]
                }]
            },
            renameActionObj:{
                tdList:[{
                    type:"input",
                    modelKey:"origin",
                    changeFun:(inputArg,inputCallback)=>{
                        let tmpArr=inputArg.item.origin.split('.');
                        inputArg.item.prefixStr=tmpArr.slice(0,tmpArr.length-1).join('.');
                        if(inputArg.item.prefixStr)inputArg.item.prefixStr+=".";
                        inputCallback(inputArg);
                    },
                    itemExpression:`ng-class="{'eo-input-error':item.target&&!item.origin}"`
                },{
                    type:"html",
                    html:`<span class="iconfont icon-jiantou_xiangyou_o"></span>`,
                    class:"w_50 tac"
                },{
                    type:"html",
                    html:`<div class="f_row_ac"><span class="mr5 c999">{{item.prefixStr}}</span><input class="eo-input" ng-model="item.target" type="text"></div>`
                },{
                    type:"btn",
                    class:"w_50",
                    btnList:[{
                        key:"删除",
                        operateName:"delete"
                    }]
                }]
            },
            moveActionObj:{
                tdList:[{
                    type:"input",
                    modelKey:"origin",
                    itemExpression:`ng-class="{'eo-input-error':item.target&&!item.origin}"`
                },{
                    type:"html",
                    html:`<span class="iconfont icon-jiantou_xiangyou_o"></span>`,
                    class:"w_50 tac"
                },{
                    type:"input",
                    modelKey:"target"
                },{
                    type:"btn",
                    class:"w_50",
                    btnList:[{
                        key:"删除",
                        operateName:"delete"
                    }]
                }]
            }
        }
        let CONST={
            ITEM:{
                decode:"json",
                encode:"origin",
                method:"post",
                proto:"http",
                timeout:"2000",
                retry:"0",
                blackList:[{
                    ip:""
                }],
                whiteList:[{
                    ip:""
                }],
                move:[{
                    origin:"",
                    target:""
                }],
                delete:[{
                    origin:""
                }],
                rename:[{
                    origin:"",
                    target:""
                }]
            }
        },privateFun={};
        vm.fun={};
        privateFun.moveStep=(inputOriginIndex,inputNewIndex)=>{
            let tmpNewItem=vm.list[inputNewIndex],tmpOriginItem=vm.list[inputOriginIndex];
            vm.list.splice(inputOriginIndex,1,tmpNewItem);
            vm.list.splice(inputNewIndex,1,tmpOriginItem);
        }
        vm.fun.oprItem=(inputOpr,inputIndex)=>{
            switch(inputOpr){
                case 'moveUp':{
                    privateFun.moveStep(inputIndex,inputIndex-1);
                    break;
                }
                case 'moveDown':{
                    privateFun.moveStep(inputIndex,inputIndex+1);
                    break;
                }
                case 'delete':{
                    $rootScope.EnsureModal('删除Step', null, '确认删除？', {}, function (callback) {
                        if (callback) {
                            vm.list.splice(inputIndex,1);
                        }
                    });
                    break;
                }
            }
        }
        vm.fun.addStep=()=>{
            vm.list.push(angular.copy(CONST.ITEM))
        }
    }
})();