(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 分组相关服务js
     */
    angular.module('goku')
        .factory('GroupService', GroupFactory);

    GroupFactory.$inject = ['$rootScope']

    function GroupFactory($rootScope) {
        var data = {
            info: {
                group: null //分组列表
            },
            fun: {
                get: null, //获取分组
                set: null, //设置分组
                clear:null,//清空分组服务信息
            }
        }
        data.fun.get = function() {
            return data.info.group;
        }
        data.fun.set = function(request, boolean) {
            data.info.group = request;
            if (boolean) {
                $rootScope.$broadcast('$SidebarFinish');
            }
        }
        data.fun.clear = function() {
            data.info.group=null;
        }
        return data.fun;
    }
})();
