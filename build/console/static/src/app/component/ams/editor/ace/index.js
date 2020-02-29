(function () {
    'use strict';

    /**
     * author：广州银云信息科技有限公司
     * ace 编辑器指令js
     * @param {string} setVariable 设置绑定setModel的对象键，(optional)，需搭配setModel使用
     * @param {object} setModel 存储值位置，双绑
     * @param {string} type 语言类型{json,javascript}，默认json
     * @param {any} watchModel 监听控件变化切换值[optional]，用于多ace编辑器
     * @param {any} readOnly 监听读写权限[optional]
     */

    angular.module('eolinker')
        .component('aceEditorAmsComponent', {
            controller: indexController,
            template: '<div id="ace-editor-ams-component-js"></div>',
            bindings: {
                setVariable: '<',
                setModel: '=',
                type: '@',
                watchModel: '<',
                readOnly: '<',
                id: '@',
                minLine:'@'
            }
        })

    indexController.$inject = ['$scope'];

    function indexController($scope) {
        var vm = this;
        var data = {
            editor: null,
            fun: {
                init: null
            }
        }

        /**
         * 监听readOnly读写权限操作函数
         */
        data.fun.readOnly = function () {
            data.editor.setReadOnly(vm.readOnly);
        };

        /**
         * 监听watchModel操作函数
         */
        data.fun.render = function () {
            if (vm.setVariable) {
                data.editor.session.setValue(vm.setModel[vm.setVariable] || '');
            } else {
                data.editor.session.setValue(vm.setModel || '');
            }
        };

        /**
         * 公用javascript初始化设置
         */
        data.fun.initJavascriptConfig = function () {
            data.editor.session.setMode("ace/mode/javascript");
            
            data.editor.setAutoScrollEditorIntoView(true);
        }

        /**
         * 自动补全自定义数据
         */
        data.fun.autoCompleteCustom = function (editor, session, pos, prefix, callback) {
            var template = {
                define: [{
                        meta: "custom",
                        caption: "eo.img",
                        value: "eo.img",
                        score: 7
                    }, {
                        meta: "custom",
                        caption: "eo.file",
                        value: "eo.file",
                        score: 7
                    }, {
                        meta: "custom",
                        caption: "eo.execute",
                        value: "eo.execute",
                        score: 6
                    },
                    {
                        meta: "custom",
                        caption: "eo.stop",
                        value: "eo.stop",
                        score: 5
                    },
                    {
                        meta: "custom",
                        caption: "eo.info",
                        value: "eo.info",
                        score: 4
                    },
                    {
                        meta: "custom",
                        caption: "eo.md5",
                        value: "eo.md5",
                        score: 3
                    },
                    {
                        meta: "custom",
                        caption: "eo.sha1",
                        value: "eo.sha1",
                        score: 2
                    },
                    {
                        meta: "custom",
                        caption: "eo.sha256",
                        value: "eo.sha256",
                        score: 1
                    },
                    {
                        meta: "env",
                        caption: "env.baseUrl",
                        value: "env.baseUrl",
                        score: 1
                    },
                    {
                        meta: "env",
                        caption: "env.headers",
                        value: "env.headers",
                        score: 1
                    },
                    {
                        meta: "env",
                        caption: "env.extraParams",
                        value: "env.extraParams",
                        score: 1
                    },
                    {
                        meta: "env",
                        caption: "env.globalParams",
                        value: "env.globalParams",
                        score: 1
                    }
                ]
            }
            if (prefix.length === 0) {
                return callback(null, []);
            } else {
                return callback(null, template.define);
            }
        }

        /**
         * 初始化编辑器
         */
        vm.$onInit = function () {
            data.editor = ace.edit(vm.id || 'ace-editor-ams-component-js');
            data.editor.setOptions({
                minLines: vm.minLine||5,
                maxLines: 30,
                enableLiveAutocompletion: true, //只能补全
            });
            data.editor.setShowPrintMargin(false);
            data.editor.setTheme("ace/theme/monokai");
            data.editor.getSession().on('change', function (e) {
                if (vm.setVariable) {
                    vm.setModel[vm.setVariable] = data.editor.getValue();
                } else {
                    vm.setModel = data.editor.getValue();
                }
                $scope.$root && $scope.$root.$$phase || $scope.$apply();
            });
            switch (vm.type) {
                case 'javascript':
                    {

                        data.fun.initJavascriptConfig();
                        try {
                            ace.require("ace/ext/language_tools")
                        } catch (e) {}
                        break;
                    }
                case 'automated-javascript':
                    {
                        data.fun.initJavascriptConfig();
                        try {
                            ace.require("ace/ext/language_tools")
                                .addCompleter({
                                    getCompletions: data.fun.autoCompleteCustom
                                });
                        } catch (e) {}
                        break;
                    }
                default:
                    {
                        data.editor.session.setMode("ace/mode/rust");
                        break;
                    }
            }
            $scope.$on('$InsertText_AceEditorAms' + (vm.id || ''), function (_default, input) {
                data.editor.insert(input);
            })
            $scope.$on('$Maunal_AceEditorAms' + (vm.id || ''), function (_default, input) {
                switch (typeof (input||'')) {
                    case 'object':
                        {
                            data.editor.session.setValue(input.data || '');
                            break;
                        }
                    default:
                        {
                            data.editor.session.setValue(input || '');
                            break;
                        }
                }
            })
            $scope.$on('$ResetAceEditor_AmsEditor' + (vm.id || ''), function () {
                data.editor.session.setValue('');
            })
        }

        $scope.$watch('$ctrl.watchModel', data.fun.render);

        $scope.$watch('$ctrl.readOnly', data.fun.readOnly);


        $scope.$on('$stateChangeStart', function () {
            if (data.editor) data.editor.destroy();
        })
        vm.$onDestroy = function () {
            if (data.editor) {
                data.editor.destroy();
            }
            
        }
    }
})();