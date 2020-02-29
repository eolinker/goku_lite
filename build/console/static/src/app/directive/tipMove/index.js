(function (window, angular) {
    'use strict';
    angular.module('eolinker.directive')
        .directive('moveTipDirective', ['$rootScope', '$compile', '$filter', function ($rootScope, $compile, $filter) {
            return {
                restrict: 'A',
                scope: {},
                link: function ($scope, elem, $attrs, $ctrl) {
                    var uuid = $filter('uuidFilter')();
                    var html = '<div class="move-tip-directive' + uuid + ' move-tip-directive">' +
                        '<div class="message-li">' + ($attrs.tipsObject ? ('{{' + $attrs.tipsObject + '}}') : $attrs.tipsText) + '</div><div class="arrow-li"></div>' +
                        '</div>';
                    angular.element(document.body).append($compile(html)($scope.$parent));
                    var rootElement = angular.element(document.getElementsByClassName($attrs.parentClassName));
                    var tipsDom = null,
                        tipsHeight = 0,
                        tipsStyle = {},
                        ableTarget = null,
                        disableTarget = null,
                        tipsDisabled = null;
                    try {
                        tipsStyle = JSON.parse($attrs.tipsStyle);
                    } catch (e) {

                    }
                    $rootScope.global.$watch.push($scope.$watch($attrs.tipsDisabled, function (currentVal, beforeVal) {
                        tipsDisabled = currentVal;
                    }))

                    tipsDom = angular.element(document.getElementsByClassName('move-tip-directive' + uuid)[0]);
                    rootElement.on('mouseover', function (event) {
                        if (!tipsHeight) {
                            tipsHeight = tipsDom[0].offsetHeight;
                        }
                        if (!$attrs.ableClassName) {
                            ableTarget = event.target;
                        } else {
                            ableTarget = getTarget(event.target, $attrs.ableClassName);
                        }
                        if (!tipsDisabled && event.target.className.indexOf($attrs.disableClassName) == -1 && (!$attrs.ableClassName || $attrs.ableClassName && !!ableTarget)) {
                            var ableTargetRect = ableTarget.getBoundingClientRect();
                            tipsDom.css({
                                left: ableTargetRect.left + (Number(tipsStyle.left) || 0) + 'px',
                                top: ableTargetRect.top + (Number(tipsStyle.top) || 0) - tipsHeight + 5 + 'px',
                                visibility: "visible"
                            });
                        } else {
                            tipsDom.css({
                                visibility: "hidden"
                            });
                        }
                    });
                    tipsDom.on('mouseover', function (event) {
                        tipsDom.css({
                            visibility: "visible"
                        });
                    });
                    tipsDom.on('mouseleave', function (event) {
                        tipsDom.css({
                            visibility: "hidden"
                        });
                    });
                    rootElement.on('mouseleave', function (event) {
                        tipsDom.css({
                            visibility: "hidden"
                        });
                    })
                    $scope.$on('$destroy', function () {
                        tipsDom.remove()
                    });
                }
            };
        }])

    function getTarget(event, ableClassName) {
        if (event.className.indexOf(ableClassName) > -1) return event;
        else if(event.parentElement){
            return getTarget(event.parentElement,ableClassName);
        }
        return false;
    }

})(window, window.angular);