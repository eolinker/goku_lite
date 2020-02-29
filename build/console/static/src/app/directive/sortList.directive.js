(function (window, angular) {
    'use strict';
    /*参考文档：https://github.com/kamilkp/angular-sortable-view
     *改进者：广州银云信息科技有限公司
     */
    angular.module('eolinker.directive')
        /*
         *这是所有的逻辑发生的地方。 
         *如果多个列表应该彼此连接，以便元素可以在它们之间移动，并且它们具有共同的祖先，则将此属性放在该元素上。 
         *如果没有，并且您仍然需要可多排序的行为，必须提供该属性的值。 
         *该值将用作将这些根连接在一起的标识符。
         */
        .directive('svRoot', [function () {
            function shouldBeAfter(elem, pointer, isGrid) { //转换节点时最低位置限制
                return isGrid ? elem.x - pointer.x < 0 : elem.y - pointer.y < 0;
            }

            function getSortableElements(key) { //获取排序节点
                return ROOTS_MAP[key];
            }

            function removeSortableElements(key) { //移除排序节点
                delete ROOTS_MAP[key];
            }

            var sortingInProgress;
            var ROOTS_MAP = Object.create(null); //外容器所包含的排序节点集
            // window.ROOTS_MAP = ROOTS_MAP; // for debug purposes

            return {
                restrict: 'A',
                controller: ['$scope', '$attrs', '$interpolate', '$parse', function ($scope, $attrs, $interpolate, $parse) {
                    var mapKey = $interpolate($attrs.svRoot)($scope) || $scope.$id;
                    if (!ROOTS_MAP[mapKey]) ROOTS_MAP[mapKey] = [];

                    var that = this;
                    var candidates; // 设置可能目的地址集
                    var $placeholder; // 占位符节点
                    var options; // 排序选项
                    var $helper; // 协助节点 - 用鼠标指针拖动的节点

                    var $original; // 原始节点
                    var $target; // 最后完美目的地址
                    var isGrid = false; //是否为网格结构
                    var onSort = $parse($attrs.svOnSort); //作为该属性的值传递的表达式将在元素排序后顺序更改时进行计算。

                    // ----- 参考 https://github.com/angular/angular.js/issues/8044
                    $attrs.svOnStart = $attrs.$$element[0].attributes['sv-on-start']; //该表达式作为一个当用户开始移动元素时被计算的参数值传递。
                    $attrs.svOnStart = $attrs.svOnStart && $attrs.svOnStart.value;

                    $attrs.svOnStop = $attrs.$$element[0].attributes['sv-on-stop']; //该表达式作为一个当用户结束移动元素时被计算的参数值传递。
                    $attrs.svOnStop = $attrs.svOnStop && $attrs.svOnStop.value;
                    // -------------------------------------------------------------------

                    var onStart = $parse($attrs.svOnStart);
                    var onStop = $parse($attrs.svOnStop);
                    var inputFun = $parse($attrs.fun)($scope);
                    this.sortingInProgress = function () {
                        return sortingInProgress;
                    };

                    if ($attrs.svGrid) { // sv-grid 是否存在
                        isGrid = $attrs.svGrid === "true" ? true : $attrs.svGrid === "false" ? false : null;
                        if (isGrid === null)
                            throw 'Invalid value of sv-grid attribute';
                    } else {
                        // 检查是否至少一个列表具有网格状布局
                        $scope.$watchCollection(function () { //浅层监视(只监视对象中的第一层元素/属性，如果这些元素/属性还有嵌套属性就不在考虑范围之内)
                            return getSortableElements(mapKey);
                        }, function (collection) {
                            isGrid = false;
                            var array = collection.filter(function (item) {
                                return !item.container;
                            }).map(function (item) {
                                return {
                                    part: item.getPart().id,
                                    y: item.element[0].getBoundingClientRect().top
                                };
                            });
                            var dict = Object.create(null);
                            array.forEach(function (item) {
                                if (dict[item.part])
                                    dict[item.part].push(item.y);
                                else
                                    dict[item.part] = [item.y];
                            });
                            Object.keys(dict).forEach(function (key) {
                                dict[key].sort();
                                dict[key].forEach(function (item, index) {
                                    if (index < dict[key].length - 1) {
                                        if (item > 0 && item === dict[key][index + 1]) {
                                            isGrid = true;
                                        }
                                    }
                                });
                            });
                        });
                    }
                    //移动更新
                    this.$moveUpdate = function (opts, mouse, svElement, svOriginal, svPlaceholder, originatingPart, originatingIndex) {
                        var svRect = svElement[0].getBoundingClientRect();
                        if (opts.tolerance === 'element')
                            mouse = {
                                x: ~~(svRect.left + svRect.width / 2),
                                y: ~~(svRect.top + svRect.height / 2)
                            };

                        sortingInProgress = true;
                        candidates = []; //候选集
                        if (!$placeholder) {
                            if (svPlaceholder) { // 自定义占位符
                                $placeholder = svPlaceholder.clone();
                                $placeholder.removeClass('ng-hide');
                            } else { // 默认占位符
                                $placeholder = svOriginal.clone();
                                $placeholder.addClass('sv-visibility-hidden');
                                $placeholder.addClass('sv-placeholder');
                                // $placeholder.css({//2016-12-8-21:57
                                //     'height': svRect.height + 'px',
                                //     'width': svRect.width + 'px'
                                // });
                                $placeholder.css({
                                    'height': svElement[0].height + 'px',
                                    'width': svElement[0].width + 'px'
                                });
                            }

                            svOriginal.after($placeholder);
                            svOriginal.addClass('ng-hide');

                            // 缓存选项，帮助器和原始元素引用
                            $original = svOriginal;
                            options = opts;
                            $helper = svElement;

                            onStart($scope, {
                                $helper: {
                                    element: $helper
                                },
                                $part: originatingPart.model(originatingPart.scope),
                                $index: originatingIndex,
                                $item: originatingPart.model(originatingPart.scope)[originatingIndex]
                            });
                            $scope.$root && $scope.$root.$$phase || $scope.$apply();
                        }

                        // ----- 移动节点
                        $helper[0].reposition({
                            x: mouse.x + document.body.scrollLeft - mouse.offset.x * svRect.width,
                            y: mouse.y + document.body.scrollTop - mouse.offset.y * svRect.height
                        });

                        // ----- 管理候选集
                        getSortableElements(mapKey).forEach(function (se, index) {
                            if (opts.containment != null) {
                                //优化，移动开始时计算
                                if (!elementMatchesSelector(se.element, opts.containment) &&
                                    !elementMatchesSelector(se.element, opts.containment + ' *')
                                ) return; // 元素不在允许的包含内
                            }
                            var rect = se.element[0].getBoundingClientRect();
                            var center = {
                                x: ~~(rect.left + rect.width / 2),
                                y: ~~(rect.top + rect.height / 2)
                            };
                            if (!se.container && // 不是容器元素
                                (se.element[0].scrollHeight || se.element[0].scrollWidth)) { // 节点可见
                                candidates.push({
                                    element: se.element,
                                    q: (center.x - mouse.x) * (center.x - mouse.x) + (center.y - mouse.y) * (center.y - mouse.y),
                                    view: se.getPart(),
                                    targetIndex: se.getIndex(),
                                    after: shouldBeAfter(center, mouse, isGrid)
                                });
                            }
                            if (se.container && !se.element[0].querySelector('[sv-element]:not(.sv-placeholder):not(.sv-source)')) { // 空容器
                                candidates.push({
                                    element: se.element,
                                    q: (center.x - mouse.x) * (center.x - mouse.x) + (center.y - mouse.y) * (center.y - mouse.y),
                                    view: se.getPart(),
                                    targetIndex: 0,
                                    container: true
                                });
                            }
                        });
                        var pRect = $placeholder[0].getBoundingClientRect();
                        var pCenter = {
                            x: ~~(pRect.left + pRect.width / 2),
                            y: ~~(pRect.top + pRect.height / 2)
                        };
                        candidates.push({
                            q: (pCenter.x - mouse.x) * (pCenter.x - mouse.x) + (pCenter.y - mouse.y) * (pCenter.y - mouse.y),
                            element: $placeholder,
                            placeholder: true
                        });
                        candidates.sort(function (a, b) {
                            return a.q - b.q;
                        });

                        candidates.forEach(function (cand, index) {
                            if (index === 0 && !cand.placeholder && !cand.container) {
                                $target = cand;
                                cand.element.addClass('sv-candidate');
                                if (cand.after)
                                    cand.element.after($placeholder);
                                else
                                    insertElementBefore(cand.element, $placeholder);
                            } else if (index === 0 && cand.container) {
                                $target = cand;
                                cand.element.append($placeholder);
                            } else
                                cand.element.removeClass('sv-candidate');
                        });
                    };

                    this.$drop = function (originatingPart, index, options) { //调整顺序 
                        if (!$placeholder) return;
                        if (options.revert) {
                            var placeholderRect = $placeholder[0].getBoundingClientRect();
                            var helperRect = $helper[0].getBoundingClientRect();
                            var distance = Math.sqrt(
                                Math.pow(helperRect.top - placeholderRect.top, 2) +
                                Math.pow(helperRect.left - placeholderRect.left, 2)
                            );

                            var duration = +options.revert * distance / 200; // 恒定速度：持续时间取决于距离
                            duration = Math.min(duration, +options.revert); // 但是它不再是options.revert

                            ['-webkit-', '-moz-', '-ms-', '-o-', ''].forEach(function (prefix) {
                                if (typeof $helper[0].style[prefix + 'transition'] !== "undefined")
                                    $helper[0].style[prefix + 'transition'] = 'all ' + duration + 'ms ease';
                            });
                            setTimeout(afterRevert, duration);
                            // $helper.css({//2016-12-8-21:57
                            //     'top': placeholderRect.top + document.body.scrollTop + 'px',
                            //     'left': placeholderRect.left + document.body.scrollLeft + 'px'
                            // });
                        } else
                            afterRevert();

                        function afterRevert() { //布局恢复函数
                            sortingInProgress = false;
                            $placeholder.remove();
                            $helper.remove();
                            $original.removeClass('ng-hide');

                            candidates = void 0;
                            $placeholder = void 0;
                            options = void 0;
                            $helper = void 0;
                            $original = void 0;

                            // sv-on-stop 反馈函数
                            onStop($scope, {
                                $part: originatingPart.model(originatingPart.scope),
                                $index: index,
                                $item: originatingPart.model(originatingPart.scope)[index]
                            });

                            if ($target) {
                                $target.element.removeClass('sv-candidate');
                                var spliced = originatingPart.model(originatingPart.scope).splice(index, 1);
                                var targetIndex = $target.targetIndex;
                                if ($target.view === originatingPart && $target.targetIndex > index)
                                    targetIndex--;
                                if ($target.after)
                                    targetIndex++;
                                $target.view.model($target.view.scope).splice(targetIndex, 0, spliced[0]);

                                // sv-on-sort 反馈函数
                                if ($target.view !== originatingPart || index !== targetIndex)
                                    onSort($scope, {
                                        $partTo: $target.view.model($target.view.scope),
                                        $partFrom: originatingPart.model(originatingPart.scope),
                                        $item: spliced[0],
                                        $indexTo: targetIndex,
                                        $indexFrom: index
                                    });
                                if ($attrs.fun) {
                                    inputFun({
                                        $index: index,
                                        $targetIndex: targetIndex
                                    });
                                }
                            }
                            $target = void 0;

                            $scope.$root && $scope.$root.$$phase || $scope.$apply();
                        }
                    };

                    this.addToSortableElements = function (se) { //添加到排序节点集
                        getSortableElements(mapKey).push(se);
                    };
                    this.removeFromSortableElements = function (se) { //从原本排序节点集移除
                        var elems = getSortableElements(mapKey);
                        var index = elems.indexOf(se);
                        if (index > -1) {
                            elems.splice(index, 1);
                            if (elems.length === 0)
                                removeSortableElements(mapKey);
                        }
                    };
                }]
            };
        }])
        /*
         *此属性应放在作为ngRepeat的元素的容器的元素上。 其值应与ng-repeat属性中的右侧表达式相同。
         */
        .directive('svPart', ['$parse', function ($parse) {
            return {
                restrict: 'A',
                require: '^svRoot', //依赖svRoot指令
                controller: ['$scope', function ($scope) {
                    $scope.$svCtrl = this;
                    this.getPart = function () { //获取sv-root $scope.part
                        return $scope.part;
                    };
                    this.$drop = function (index, options) {
                        $scope.$sortableRoot.$drop($scope.part, index, options);
                    };
                }],
                scope: true,
                link: function ($scope, $element, $attrs, $sortable) {
                    if (!$attrs.svPart) throw new Error('no model provided');
                    var model = $parse($attrs.svPart);
                    if (!model.assign) throw new Error('model not assignable');

                    $scope.part = {
                        id: $scope.$id,
                        element: $element,
                        model: model,
                        scope: $scope
                    };
                    $scope.$sortableRoot = $sortable;

                    var sortablePart = {
                        element: $element,
                        getPart: $scope.$svCtrl.getPart,
                        container: true
                    };
                    $sortable.addToSortableElements(sortablePart);
                    $scope.$on('$destroy', function () {
                        $sortable.removeFromSortableElements(sortablePart);
                    });
                }
            };
        }])
        /*
         *此属性应放置在与ng-repeat属性相同的元素上。 
         *它的（可选）值应该是一个计算为options对象的表达式。
         *含：mousedown touchstart mousemove touchmove mouseup touchend touchcancel操作
         */
        .directive('svElement', ['$parse', function ($parse) {
            return {
                restrict: 'A',
                require: ['^svPart', '^svRoot'], //依赖svPart以及svRoot指令
                controller: ['$scope', function ($scope) {
                    $scope.$svCtrl = this;
                }],
                link: function ($scope, $element, $attrs, $controllers) {
                    var sortableElement = {
                        element: $element,
                        getPart: $controllers[0].getPart,
                        getIndex: function () {
                            return $scope.$index;
                        }
                    };
                    $controllers[1].addToSortableElements(sortableElement);
                    $scope.$on('$destroy', function () {
                        $controllers[1].removeFromSortableElements(sortableElement);
                    });

                    var handle = $element;
                    handle.on('mousedown touchstart', onMousedown);
                    $scope.$watch('$svCtrl.handle', function (customHandle) {
                        if (customHandle) {
                            handle.off('mousedown touchstart', onMousedown);
                            handle = customHandle;
                            handle.on('mousedown touchstart', onMousedown);
                        }
                    });

                    var helper;
                    $scope.$watch('$svCtrl.helper', function (customHelper) {
                        if (customHelper) {
                            helper = customHelper;
                        }
                    });

                    var placeholder;
                    $scope.$watch('$svCtrl.placeholder', function (customPlaceholder) {
                        if (customPlaceholder) {
                            placeholder = customPlaceholder;
                        }
                    });

                    var body = angular.element(document.body);
                    var html = angular.element(document.documentElement);

                    var moveExecuted;

                    function onMousedown(e) { //mouseDown函数
                        touchFix(e);

                        if ($controllers[1].sortingInProgress()) return;
                        if (e.button != 0 && e.type === 'mousedown') return;

                        moveExecuted = false;
                        var opts = $parse($attrs.svElement)($scope);
                        opts = angular.extend({}, {
                            tolerance: 'pointer',
                            revert: 200,
                            containment: 'html'
                        }, opts);
                        if (opts.containment) {
                            var containmentRect = closestElement.call($element, opts.containment)[0].getBoundingClientRect();
                        }

                        var target = $element;
                        var clientRect = $element[0].getBoundingClientRect();
                        var clone;

                        if (!helper) helper = $controllers[0].helper;
                        if (!placeholder) placeholder = $controllers[0].placeholder;
                        if (helper) {
                            clone = helper.clone();
                            clone.removeClass('ng-hide');
                            clone.css({
                                'left': clientRect.left + document.body.scrollLeft + 'px',
                                'top': clientRect.top + document.body.scrollTop + 'px'
                            });
                            target.addClass('sv-visibility-hidden');
                        } else {
                            clone = target.clone();
                            clone.addClass('sv-helper').css({
                                'left': clientRect.left + document.body.scrollLeft + 'px',
                                'top': clientRect.top + document.body.scrollTop + 'px',
                                'width': clientRect.width + 'px'
                            });
                        }

                        clone[0].reposition = function (coords) { //克隆元素重定位
                            var targetLeft = coords.x;
                            var targetTop = coords.y;
                            var helperRect = clone[0].getBoundingClientRect();

                            var body = document.body;

                            if (containmentRect) {
                                if (targetTop < containmentRect.top + body.scrollTop) // 上边界
                                    targetTop = containmentRect.top + body.scrollTop;
                                if (targetTop + helperRect.height > containmentRect.top + body.scrollTop + containmentRect.height) // bottom boundary
                                    targetTop = containmentRect.top + body.scrollTop + containmentRect.height - helperRect.height;
                                if (targetLeft < containmentRect.left + body.scrollLeft) // 左边界
                                    targetLeft = containmentRect.left + body.scrollLeft;
                                if (targetLeft + helperRect.width > containmentRect.left + body.scrollLeft + containmentRect.width) // right boundary
                                    targetLeft = containmentRect.left + body.scrollLeft + containmentRect.width - helperRect.width;
                            }
                            this.style.left = targetLeft - body.scrollLeft + 'px';
                            this.style.top = targetTop - body.scrollTop + 'px';
                        };

                        var pointerOffset = {
                            x: (e.clientX - clientRect.left) / clientRect.width,
                            y: (e.clientY - clientRect.top) / clientRect.height
                        };
                        html.addClass('sv-sorting-in-progress');
                        //mouseUp相应函数
                        html.on('mousemove touchmove', onMousemove).on('mouseup touchend touchcancel', function mouseup(e) {
                            html.off('mousemove touchmove', onMousemove);
                            html.off('mouseup touchend touchcancel', mouseup);
                            html.removeClass('sv-sorting-in-progress');
                            if (moveExecuted) {
                                $controllers[0].$drop($scope.$index, opts);
                            }
                            $element.removeClass('sv-visibility-hidden');
                        });

                        // onMousemove(e);
                        function onMousemove(e) {
                            touchFix(e);
                            if (!moveExecuted) {
                                $element.parent().prepend(clone);
                                moveExecuted = true;
                            }
                            $controllers[1].$moveUpdate(opts, {
                                x: e.clientX,
                                y: e.clientY,
                                offset: pointerOffset
                            }, clone, $element, placeholder, $controllers[0].getPart(), $scope.$index);
                        }
                    }
                }
            };
        }])
        /*
         *此属性是可选的。 如果需要，它可以放置在可排序元素内的元素上。 这个元素将是排序操作的句柄。
         */
        .directive('svHandle', function () {
            return {
                require: '?^svElement', //依赖svElement指令
                link: function ($scope, $element, $attrs, $svCtrl) {
                    if ($svCtrl)
                        $svCtrl.handle = $element.add($svCtrl.handle); // 支持添加多级把手
                }
            };
        });

    angular.element(document.head).append([ //头部添加style
        '<style>' +
        '.sv-helper{' +
        'position: fixed !important;' +
        'z-index: 99999;' +
        'margin: 0 !important;' +
        '}' +
        '.sv-candidate{' +
        '}' +
        '.sv-placeholder{' +
        // 'opacity: 0;' +
        '}' +
        '.sv-sorting-in-progress{' +
        '-webkit-user-select: none;' +
        '-moz-user-select: none;' +
        '-ms-user-select: none;' +
        'user-select: none;' +
        '}' +
        '.sv-visibility-hidden{' +
        'visibility: hidden !important;' +
        'opacity: 0 !important;' +
        '}' +
        '</style>'
    ].join(''));

    function touchFix(e) { //拖动位置匹配
        if (!('clientX' in e) && !('clientY' in e)) {
            var touches = e.touches || e.originalEvent.touches;
            if (touches && touches.length) {
                e.clientX = touches[0].clientX;
                e.clientY = touches[0].clientY;
            }
            e.preventDefault();
        }
    }

    function getPreviousSibling(element) { //获取上一级元素
        element = element[0];
        if (element.previousElementSibling)
            return angular.element(element.previousElementSibling);
        else {
            var sib = element.previousSibling;
            while (sib != null && sib.nodeType != 1)
                sib = sib.previousSibling;
            return angular.element(sib);
        }
    }

    function insertElementBefore(element, newElement) { //在被选元素内部的开头插入新节点
        var prevSibl = getPreviousSibling(element);
        if (prevSibl.length > 0) {
            prevSibl.after(newElement);
        } else {
            element.parent().prepend(newElement);
        }
    }

    var dde = document.documentElement,
        matchingFunction = dde.matches ? 'matches' :
        dde.matchesSelector ? 'matchesSelector' :
        dde.webkitMatches ? 'webkitMatches' :
        dde.webkitMatchesSelector ? 'webkitMatchesSelector' :
        dde.msMatches ? 'msMatches' :
        dde.msMatchesSelector ? 'msMatchesSelector' :
        dde.mozMatches ? 'mozMatches' :
        dde.mozMatchesSelector ? 'mozMatchesSelector' : null;
    if (matchingFunction == null)
        throw 'This browser doesn\'t support the HTMLElement.matches method';

    function elementMatchesSelector(element, selector) { //设置节点匹配选择器
        if (element instanceof angular.element) element = element[0];
        if (matchingFunction !== null)
            return element[matchingFunction](selector);
    }

    var closestElement = angular.element.prototype.closest || function (selector) {
        var el = this[0].parentNode;
        while (el !== document.documentElement && !el[matchingFunction](selector))
            el = el.parentNode;

        if (el[matchingFunction](selector))
            return angular.element(el);
        else
            return angular.element();
    };

    /*
        简单实现jQuery .add方法
     */
    if (typeof angular.element.prototype.add !== 'function') {
        angular.element.prototype.add = function (elem) {
            var i, res = angular.element();
            elem = angular.element(elem);
            for (i = 0; i < this.length; i++) {
                res.push(this[i]);
            }
            for (i = 0; i < elem.length; i++) {
                res.push(elem[i]);
            }
            return res;
        };
    }

})(window, window.angular);