(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 上传文件指令js
     */
    angular.module('eolinker.directive')

    .directive('uploadFileDirective', [function() {
        return {
            restrict: 'AE',
            template:'<input  autocomplete="off" name="file" id="{{inputId}}" class="hidden" type="file" onChange="angular.element(this).scope().uploadFileDirective({arg:{$files:this.files}})" file-reset-directive button-function="change" accept="{{fileType}} "/>',
            scope: {
                fileType:'@',
                inputId:'@',
                uploadFileDirective: '&' //绑定设置回调函数
            }
        };
    }]);
})();