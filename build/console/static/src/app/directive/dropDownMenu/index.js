(function () {
    'use strict';
    /**
     * @description 手动控制按钮focus状态，使用与mac os 非chrome浏览器
     * @author 广州银云信息科技有限公司
     */
    angular.module('eolinker.directive')

        .directive('dropDownMenuCommonDirective', ['$rootScope', function ($rootScope) {
            return {
                restrict: 'AE',
                scope: {
                    dirDisable: '<'
                },
                link: function ($scope, elem, attrs, ctrl) {
                    $scope.data={
                        elemArr:elem[0].getElementsByClassName('eo_more_btn')
                    }
                    let privateFun={};
                    privateFun.initWatchDom=()=>{
                        $rootScope.global.$watch.push($scope.$watch('data.elemArr.length', ()=>{
                            if($scope.data.elemArr){
                                let domArr=Array.prototype.slice.call($scope.data.elemArr);
                                domArr.map((val)=>{
                                    let tmpElem=val;
                                    angular.element(tmpElem).bind('click',(event)=>{
                                        tmpElem.focus();
                                    })
                                })
                            }
                            
                        }));
                    }
                    var main=(function() {
                        if(/macintosh|mac os x/i.test(navigator.userAgent)&&!/Chrome/i.test(navigator.userAgent)){
                            privateFun.initWatchDom();
                            $scope.$on('$stateChangeStart',()=>{
                                $scope.data.elemArr=null;
                            })
                            $scope.$on('$stateChangeSuccess',()=>{
                                if(!$scope.data.elemArr){
                                    $scope.data.elemArr=elem[0].getElementsByClassName('eo_more_btn');
                                    privateFun.initWatchDom();
                                }
                            })
                            console.log(navigator.userAgent)
                        }
                        
                        
                    })();
                    
                }
            };
        }]);
})();