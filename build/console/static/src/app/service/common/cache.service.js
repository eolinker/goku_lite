(function() {
    'use strict';
    /*
     * @author 广州银云信息科技有限公司
     * @description 缓存公用服务js
     */
    angular.module('eolinker')
        .factory('Cache_CommonService', index);

    index.$inject = []

    function index() {
        var data = {
            info: {
                cache: null,
                statusCache:null
            },
            fun: {
                get: null,
                set: null
            }
        }

        data.fun.clear = function(status) {
            if (status) {
                try {
                    data.info.statusCache[status] = null;
                } catch (e) {}
            } else {
                data.info.cache = null;
            }
        }

        /**
         * 获取缓存信息
         * @return {any} 缓存内容
         */
        data.fun.get = function(status) {
            if (status) {
                try {
                    return data.info.statusCache[status];
                } catch (e) {
                    return null;
                }
                return data.info.statusCache;
            }else{
                return data.info.cache;
            }
        }

        /**
         * 设置缓存信息
         * @param {any} input 传参内容
         */
        data.fun.set = function(input, status) {
            if (status) {
                try {
                    data.info.statusCache[status] = input;

                } catch (e) {
                    data.info.statusCache = {};
                    data.info.statusCache[status] = input;
                }
            } else {
                data.info.cache = input;
            }
        }
        return data.fun;
    }
})();
