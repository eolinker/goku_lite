(function () {
    'use strict';
    /**
     * @name 导航栏服务
     * @author 广州银云信息科技有限公司
     */
    angular.module('eolinker')
        .factory('NavbarService', NavbarService);

    NavbarService.$inject = []

    function NavbarService() {
        var data = {
                navigation: {
                    query: [],
                    current: ''
                }
            };
        return {
            info: data
        };
    }
})();