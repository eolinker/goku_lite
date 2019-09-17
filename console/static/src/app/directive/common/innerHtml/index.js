(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 注入HTML指令js
     * @require {string} html 所需注入的代码
     * @param {string} status 注入代码类型，默认绑定angular元素，若关闭状态为unbind-angular
     */
    angular.module('eolinker.directive')

    .directive('innerHtmlCommonDirective', ['$compile','Cache_CommonService','$rootScope', function($compile,Cache_CommonService,$rootScope) {
        return {
            restrict: 'AE',
            scope:{
                html:'<',
                innerHtmlCommonDirective:'@'
            },
            link: function($scope, elem, attrs, ctrl) {
              
                $rootScope.global.$watch.push($scope.$watch(attrs.html?'html':'innerHtmlCommonDirective',function(){
                    var template={
                        html:attrs.html?$scope.html:$scope.innerHtmlCommonDirective,
                        elemFunName:'append'
                    }
                    if(!template.html)return;
                    if(attrs.remove)elem.empty();
                    if(attrs.position=='front') {
                        template.elemFunName='prepend';
                    }
                    switch(attrs.status){
                        case 'unbind-angular':{
                            elem[template.elemFunName](template.html);
                            break;
                        }
                        case 'cache':{
                            elem[template.elemFunName](Cache_CommonService.get(attrs.cacheVariable));
                            break;
                        }
                        default:{
                            try{
                                elem[template.elemFunName]($compile(template.html)($scope.$parent));
                            }catch(e){
                                elem[template.elemFunName]($compile('<span>'+template.html+'</span>')($scope.$parent));
                            }
                            break;
                        }
                    }
                }))
            }
        };
    }]);
})();