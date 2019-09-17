(function () {
    'use strict';
    /**
     * 获取某一节点的数量
     * @param {string} bindClass 绑定监听类
     * @param {number} model 绑定视图数据
     */
    angular.module('eolinker.directive')
        .directive('dragChangeSpacingCommonDirective', ['$rootScope', function ($rootScope) {
            return {
                restrict: 'A',
                scope: {
                    mainObject: '<'
                },
                link: function ($scope, elem, attrs, ngModel) {
                    let elemAffect = document.getElementsByClassName(attrs.affectClass), containerElemAffect = null;
                    let privateFun = {}, elemHead = document.getElementsByTagName('head'), elemWidth = '0px';
                    /**
                     * @desc 更改对象高度
                     */
                    privateFun.setHeight = (domEvent) => {
                        if (domEvent.clientY <= $scope.mainObject.setting.clientY) return false;
                        let tmpHeight = (document.body.clientHeight - domEvent.clientY);
                        if (tmpHeight >= $scope.mainObject.setting.minHeight) {
                            elemAffect[0].style.height = tmpHeight + 'px';
                            $scope.$root && $scope.$root.$$phase || $scope.$apply();
                        }
                        return false
                    }
                    /**
                     * @desc 更改对象宽度
                     */
                    privateFun.setWidth = (domEvent) => {
                        let tmpMoveX = domEvent.movementX;
                        let tmpWidth = elemAffect[0].clientWidth + tmpMoveX - 15;
                        if (tmpWidth <= $scope.mainObject.setting.minWidth) return;
                        // if(containerElemAffect){
                        //     containerElemAffect[0].style.width=containerElemAffect[0].scrollWidth+tmpMoveX  + 'px';
                        // }
                        elemWidth = tmpWidth + 'px';
                        for (let key = 0; key < $scope.mainObject.setting.affectCount; key++) {
                            elemAffect[key].style.width = elemWidth;
                        }
                        $scope.$root && $scope.$root.$$phase || $scope.$apply();
                    }
                    elem.bind('mousedown', (inputEvent) => {
                        inputEvent.stopPropagation();
                        elem.top = elem.offsetTop;
                        if ($scope.mainObject && $scope.mainObject.baseFun && $scope.mainObject.baseFun.mouseDown) {
                            $scope.mainObject.baseFun.mouseDown(elemAffect);
                        }
                        switch ($scope.mainObject.setting.object) {
                            case 'height': {
                                document.onmousemove = privateFun.setHeight;
                                break;
                            }
                            case 'width': {
                                angular.element(elemHead).append('<style id="eo_tmp_drag" type="text/css">*{cursor:col-resize!important;user-select: none;}</style>')
                                document.onmousemove = privateFun.setWidth;
                                break;
                            }
                        }

                        document.onmouseup = function () {
                            document.onmousemove = null;
                            document.onmouseup = null;
                            let tmpHeadStyleElem = document.getElementById('eo_tmp_drag');
                            angular.element(tmpHeadStyleElem).remove();
                            elem.releaseCapture && elem.releaseCapture();
                            if ($scope.mainObject && $scope.mainObject.baseFun && $scope.mainObject.baseFun.mouseup) {
                                $scope.mainObject.baseFun.mouseup($scope.mainObject.mark, elemWidth);
                            }
                        };
                        elem.setCapture && elem.setCapture();
                        return false
                    })

                    function main() {
                        if (attrs.containerAffectClass) {
                            containerElemAffect = document.getElementsByClassName(attrs.containerAffectClass);
                        }
                    }
                    main();
                }
            };
        }]);

})();