(function() {
    'use strict';
    /*
     * @author 广州银云信息科技有限公司
     * @description 单向流公用服务js
     */
    angular.module('goku')
        .factory('Communicate_CommonService', index);

    index.$inject = []

    function index() {
        var data={}
        var fun={};
        fun.clear = function(status) {
            data[status]=null;
        }

        /**
         * 设置缓存信息
         * @param {any} input 传参内容
         */
        fun.set = function(status,input) {
            data[status] = angular.copy(input);
        }
        return {
            data:data,
            fun:fun
        };
    }
})();
