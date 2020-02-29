(function () {
    /**
     * @description 产品通用型概况页
     * @author 广州银云信息科技有限公司
     */
    'use strict';
    angular.module('eolinker')
        .component('overviewProductComponent', {
            templateUrl: 'app/component/product/overview/index.html',
            controller: indexController,
            bindings:{
                listAuthorityObject:'<',
                authorityObject:'<',
                mainObject:'<',
                otherObject:'<'
            }
        })

    indexController.$inject = [];

    function indexController(){
        var vm = this;
        vm.fun={};
        vm.fun.common = function (extend, arg) {
            let tmpParam={};
            switch (typeof (extend.params)) {
                case 'string':
                    {
                        return eval('extend.fun(' + extend.params + ')');
                    }
                default:
                    {
                        for (var key in extend.params) {
                            if (extend.params[key] == null) {
                                tmpParam[key] = arg[key];
                            } else {
                                tmpParam[key] = extend.params[key];
                            }
                        }
                        return extend.fun(tmpParam);
                    }
            }
        }
    }
})();