(function () {
    /**
     * @name 滚动加载
     * @author 广州银云信息科技有限公司
     */
    'use strict';
    angular.module('eolinker.directive')
        .value('THROTTLE_MILLISECONDS', null)
        .directive('infiniteScroll', ['$rootScope', '$window', '$interval', 'THROTTLE_MILLISECONDS', function ($rootScope, $window, $interval, THROTTLE_MILLISECONDS) {
            return {
                scope: {
                    infiniteScroll: '&',
                    infiniteScrollContainer: '=',
                    infiniteScrollDistance: '=',
                    infiniteScrollDisabled: '=',
                    infiniteScrollCancel: '=',
                    infiniteScrollUseDocumentBottom: '=',
                    infiniteScrollListenForEvent: '@'
                },
                link: function (scope, elem, attrs) {
                    var changeContainer, lastScrollTop, checkInterval, checkWhenEnabled, container, handleInfiniteScrollContainer, handleInfiniteScrollDisabled, handleInfiniteScrollDistance, handleInfiniteScrollUseDocumentBottom, handler, height, immediateCheck, offsetTop, pageYOffset, scrollDistance, scrollEnabled, throttle, unregisterEventListener, useDocumentBottom, windowElement;
                    windowElement = angular.element($window);
                    scrollDistance = null;
                    scrollEnabled = null;
                    checkWhenEnabled = null;
                    container = null;
                    immediateCheck = true;
                    useDocumentBottom = false;
                    unregisterEventListener = null;
                    checkInterval = false;
                    lastScrollTop = 0;
                    height = function (elem) {
                        elem = elem[0] || elem;
                        if (isNaN(elem.offsetHeight)) {
                            return elem.document.documentElement.clientHeight;
                        } else {
                            return elem.offsetHeight;
                        }
                    };
                    offsetTop = function (elem) {
                        if (!elem[0].getBoundingClientRect || elem.css('none')) {
                            return;
                        }
                        return elem[0].getBoundingClientRect().top + pageYOffset(elem);
                    };
                    pageYOffset = function (elem) {
                        elem = elem[0] || elem;
                        if (isNaN(window.pageYOffset)) {
                            return elem.document.documentElement.scrollTop;
                        } else {
                            return elem.ownerDocument.defaultView.pageYOffset;
                        }
                    };
                    handler = function () {
                        var containerBottom, containerTopOffset, elementBottom, remaining, shouldScroll, scrollTop;
                        if (container === windowElement) {
                            containerBottom = height(container) + pageYOffset(container[0].document.documentElement);
                            var elemHeight = height(elem)
                            elementBottom = offsetTop(elem) + elemHeight;
                            // scrollTop = elemHeight - elementBottom;
                        } else {
                            containerBottom = height(container);
                            containerTopOffset = 0;
                            if (offsetTop(container) !== void 0) {
                                containerTopOffset = offsetTop(container);
                            }
                            var elemHeight = height(elem);
                            elementBottom = offsetTop(elem) - containerTopOffset + elemHeight;
                            // scrollTop = elemHeight - elementBottom;
                        }
                        if (useDocumentBottom) {
                            elementBottom = height((elem[0].ownerDocument || elem[0].document).documentElement);
                        }
                        // if(scrollTop>lastScrollTop){
                        //     shouldScroll = remaining <= height(container) * scrollDistance + 1;
                        //     //向下滚动
                        // }else{
                        //     //向上滚动
                        //     shouldScroll = remaining <= height(container) * scrollDistance + 1;
                        // }
                        remaining = elementBottom - containerBottom;
                        shouldScroll = (remaining <= (height(container) * scrollDistance*0.1))&&elementBottom;
                        // lastScrollTop=scrollTop;
                        if (shouldScroll) {
                            checkWhenEnabled = true;
                            if (scrollEnabled) {
                                if (scope.$$phase || $rootScope.$$phase) {
                                    return scope.infiniteScroll();
                                } else {
                                    return scope.$apply(scope.infiniteScroll);
                                }
                            }
                        } else {
                            if (checkInterval) {
                                $interval.cancel(checkInterval);
                            }
                            return checkWhenEnabled = false;
                        }
                    };
                    throttle = function (func, wait) {
                        var later, previous, timeout;
                        timeout = null;
                        previous = 0;
                        later = function () {
                            previous = new Date().getTime();
                            $interval.cancel(timeout);
                            timeout = null;
                            return func.call();
                        };
                        return function () {
                            var now, remaining;
                            now = new Date().getTime();
                            remaining = wait - (now - previous);
                            if (remaining <= 0) {
                                $interval.cancel(timeout);
                                timeout = null;
                                previous = now;
                                return func.call();
                            } else {
                                if (!timeout) {
                                    return timeout = $interval(later, remaining, 1);
                                }
                            }
                        };
                    };
                    if (THROTTLE_MILLISECONDS != null) {
                        handler = throttle(handler, THROTTLE_MILLISECONDS);
                    }
                    scope.$on('$destroy', function () {
                        container.unbind('scroll', handler);
                        if (unregisterEventListener != null) {
                            unregisterEventListener();
                            unregisterEventListener = null;
                        }
                        if (checkInterval) {
                            return $interval.cancel(checkInterval);
                        }
                    });
                    handleInfiniteScrollDistance = function (v) {
                        return scrollDistance = parseFloat(v) || 0;
                    };
                    scope.$watch('infiniteScrollDistance', handleInfiniteScrollDistance);
                    handleInfiniteScrollDistance(scope.infiniteScrollDistance);
                    handleInfiniteScrollDisabled = function (v) {
                        scrollEnabled = !v;
                        if (scrollEnabled && checkWhenEnabled) {
                            checkWhenEnabled = false;
                            return handler();
                        }
                    };
                    scope.$watch('infiniteScrollDisabled', handleInfiniteScrollDisabled);
                    handleInfiniteScrollDisabled(scope.infiniteScrollDisabled);
                    handleInfiniteScrollUseDocumentBottom = function (v) {
                        return useDocumentBottom = v;
                    };
                    scope.$watch('infiniteScrollUseDocumentBottom', handleInfiniteScrollUseDocumentBottom);
                    handleInfiniteScrollUseDocumentBottom(scope.infiniteScrollUseDocumentBottom);
                    changeContainer = function (newContainer) {
                        if (container != null) {
                            container.unbind('scroll', handler);
                        }
                        container = newContainer;
                        if (newContainer != null) {
                            return container.bind('scroll', handler);
                        }
                    };
                    changeContainer(windowElement);
                    if (scope.infiniteScrollListenForEvent) {
                        unregisterEventListener = $rootScope.$on(scope.infiniteScrollListenForEvent, handler);
                    }
                    handleInfiniteScrollContainer = function (newContainer) {
                        if ((newContainer == null) || newContainer.length === 0) {
                            return;
                        }
                        if (newContainer.nodeType && newContainer.nodeType === 1) {
                            newContainer = angular.element(newContainer);
                        } else if (typeof newContainer.append === 'function') {
                            newContainer = angular.element(newContainer[newContainer.length - 1]);
                        } else if (typeof newContainer === 'string') {
                            newContainer = angular.element(document.querySelector(newContainer));
                        }
                        if (newContainer != null) {
                            return changeContainer(newContainer);
                        } else {
                            throw new Error("invalid infinite-scroll-container attribute.");
                        }
                    };
                    scope.$watch('infiniteScrollContainer', handleInfiniteScrollContainer);
                    scope.$watch('infiniteScrollCancel', function (isCancel) {
                        if (isCancel) {
                            container.unbind('scroll', handler);
                            if (unregisterEventListener != null) {
                                unregisterEventListener();
                                unregisterEventListener = null;
                            }
                            if (checkInterval) {
                                return $interval.cancel(checkInterval);
                            }
                        } else {
                            changeContainer(angular.element(elem.parent()));
                        }
                    });
                    handleInfiniteScrollContainer(scope.infiniteScrollContainer || []);
                    if (attrs.infiniteScrollParent != null && attrs.infiniteScrollCancel == null) {
                        changeContainer(angular.element(elem.parent()));
                    }
                    if (attrs.infiniteScrollImmediateCheck != null) {
                        immediateCheck = scope.$eval(attrs.infiniteScrollImmediateCheck);
                    }

                    function getFirstFunction(immediateCheck) {
                        if (immediateCheck) {
                            return $interval((function () {
                                return handler();
                            }), 0)
                        } else {
                            return false;
                        }
                    }
                    return checkInterval = getFirstFunction(immediateCheck);
                }
            };
        }]);
})();