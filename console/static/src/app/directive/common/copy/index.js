(function () {
    'use strict';
    /**
     * @description 复制指令
     * @author 广州银云信息科技有限公司
     * @param [string][optional] copyModel 复制绑定模块  
     * @param [string][optional] cacheVariable 复制缓存绑定，用于大数据不方便页面交互行数据
     * @extends $rootScope
     * @extends Cache_CommonService
     */

    angular.module('eolinker')
        .directive('copyCommonDirective', ['$rootScope','Cache_CommonService',function ($rootScope,Cache_CommonService) {
            return {
                restrict: 'A',
                scope:{
                    copyModel:'<'
                },
                link: function ($scope, elem, attrs, ctrl) {
                    var data = {
                        elem:null
                    },fun={};
                    var service={
                        cache:Cache_CommonService
                    }
                    fun.btnFun = function ($event) {
                        $event.stopPropagation();
                        data.elem.value = $scope.copyModel||(attrs.cacheVariable?service.cache.get(attrs.cacheVariable):'');
                        data.elem.select();
                        data.elem.click();
                        try {
                            if (document.execCommand('copy')) {
                                $rootScope.InfoModal("复制成功", 'success');
                            } else {
                                $rootScope.InfoModal("复制失败", 'error');
                            }

                        } catch (err) {
                            $rootScope.InfoModal("复制失败", 'error');
                        }
                    }
                    fun.init = (function () {
                        data.elem=document.getElementById('template_textarea_js')||document.createElement('textarea');
                        data.elem.setAttribute("style", "position:fixed,left:0,top:0,opacity:0;height:0;width:0;");
                        data.elem.setAttribute("id", "template_textarea_js");
                        document.body.appendChild(data.elem);
                        elem.bind(attrs.buttonFunction || 'click', fun.btnFun);
                    })()
                    $scope.$on('$destory',function(){
                        document.body.removeChild(data.elem);
                    })
                }
            };
        }]);
})();