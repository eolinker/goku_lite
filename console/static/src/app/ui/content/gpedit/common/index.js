(function () {
    'use strict';
    angular.module('eolinker')
        .component('gpeditCommon', {
            templateUrl: 'app/ui/content/gpedit/common/index.html',
            controller: indexController
        })

    indexController.$inject = ['Cache_CommonService','$state'];

    function indexController(Cache_CommonService,$state) {
        let vm=this,service={
            cache:Cache_CommonService
        };
        vm.component={};
        vm.$onInit=()=>{
            service.cache.clear('gpeditGroup');
            vm.component.menuObject={
                list:[{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    btnList: [{
                        name: '返回策略列表',
                        icon: 'chexiao',
                        fun: {
                            default: ()=>{
                                $state.go("home.gpedit.default");
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