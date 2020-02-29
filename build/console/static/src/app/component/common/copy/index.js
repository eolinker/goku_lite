(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 复制input值到剪贴板js（欠缺及需优化地：样式需要自己个项目调控，无法做到自定义样式模板）
     */

    angular.module('eolinker')
        .component('copyCommonComponent', {
            template: '<inner-html-common-directive html="$ctrl.data.info.html"></inner-html-common-directive>',
            bindings: {
                copyModel: '<',
                buttonHtml: '@',
                isPopup: '@',
                switchTemplet: '@'
            },
            controller: indexController
        });

    indexController.$inject = ['$rootScope', '$scope'];

    function indexController($rootScope, $scope) {
        var vm = this;
        var data = {
            info: {
                templet: {
                    button: '<button class="eo_theme_btn_default copy-common-component-' + $scope.$id + '" data-clipboard-text="{{$ctrl.copyModel}}"><span class=\"iconfont icon-copy\" ><\/span>{{$ctrl.data.info.clipboard.text}}</button>',
                    input: '<input  autocomplete="off" type="text" id="copy-common-component-' + $scope.$id + '" name="link" value="{{$ctrl.copyModel}}" class="eo-input copy-common-component-' + $scope.$id + '" data-clipboard-action="copy" data-clipboard-target="#copy-common-component-' + $scope.$id + '" ng-class="{\'eo-copy\':($ctrl.data.info.clipboard.success)&&($ctrl.data.info.clipboard.isClick)}" data-ng-click="$ctrl.data.fun.click()" readonly>' + '<label for="copy-common-component-' + $scope.$id + '" class="pull-right copy-tips " ng-class="{\'copy-success\':($ctrl.data.info.clipboard.success)&&($ctrl.data.info.clipboard.isClick),\'copy-error\':(!$ctrl.data.info.clipboard.success)&&($ctrl.data.info.clipboard.isClick)}">' + '{{$ctrl.data.info.clipboard.text}}' + '</label>',
                    textarea: '<textarea id="copy-common-component-' + $scope.$id + '" readonly>{{$ctrl.copyModel}}</textarea><button data-clipboard-action="copy" data-clipboard-target="#copy-common-component-' + $scope.$id + '">{{$ctrl.data.info.clipboard.text}}</button>'
                },
                clipboard: null
            },
            fun: {
                init: null, //初始化功能函数
                reset: null, //重置功能函数
                $destory: null, //页面销毁功能函数
            }
        }
        vm.data = {
            info: {
                html: null,
                clipboard: {
                    isClick: false,
                    success: false,
                    text: vm.buttonHtml || '点击复制' //显示button文本（默认文本'点击复制'）
                }
            },
            fun: {
                click: null, //单击功能函数
            }
        }
        data.fun.reset = function(arg) {
            data.info.clipboard = new Clipboard(arg.class);
            data.info.clipboard.on('success', function(_default) {
                vm.data.info.clipboard.success = true;
                vm.data.info.clipboard.isClick = true;
                console.info('Text:', _default.text);
                if (vm.isPopup) { //成功或者失败是否以弹窗形式提醒
                    $rootScope.InfoModal("已复制到剪贴板", 'success');
                } else {
                    vm.data.info.clipboard.text = '复制成功';
                }
                $scope.$root && $scope.$root.$$phase || $scope.$apply();
                _default.clearSelection();
            });

            data.info.clipboard.on('error', function(_default) {
                vm.data.info.clipboard.success = false;
                vm.data.info.clipboard.isClick = true;
                console.info('Text:', _default.text);
                if (vm.isPopup) {
                    $rootScope.InfoModal("复制到剪贴板失败", 'error');
                } else {
                    vm.data.info.clipboard.text = '复制失败';
                }
                $scope.$root && $scope.$root.$$phase || $scope.$apply();
            });
        }
        vm.data.fun.click = function() {
            vm.data.info.clipboard.isClick = false;
        }
        data.fun.$destroy = function() {
            data.info.clipboard.destroy();
        }
        vm.$onInit = function() {
            switch (vm.switchTemplet) { //选择模板（0：button模板，1：input模板，2：textarea模板，默认input模板）
                case '0':
                    {
                        vm.data.info.html = data.info.templet.button;
                        break;
                    }
                case '1':
                    {
                        vm.data.info.html = data.info.templet.input;
                        break;
                    }
                case '2':
                    {
                        vm.data.info.html = data.info.templet.textarea;
                        break;
                    }
                default:
                    {
                        vm.data.info.html = data.info.templet.input;
                        break;
                    }
            }
            $scope.$on('$destroy', data.fun.$destroy);
            data.fun.reset({
                class: ('.copy-common-component-' + $scope.$id)
            });
        }
    }
})();
