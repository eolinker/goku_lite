(function () {
    'use strict';
    /**
     * @description 返回dom元素以及其祖先寻找某个属性
     * @author 广州银云信息科技有限公司
     */
    angular.module('eolinker.filter')
        .filter('findAttr', [function (dom,attr,level) {
            return function (dom,attr,level) {
                level=level||4;
                function getAttr(dom,attr){
                    level--;
                    var value=dom.getAttribute(attr);
                    if(value) return value;
                    if(level){
                        return getAttr(dom.parentNode,attr);
                    }
                }
                return getAttr(dom,attr);
            }
        }])

})();