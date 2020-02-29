(() => {
    /**
     * @name 网关设置路由配置
     * @author 广州银云信息科技有限公司
     */
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.setting', {
                    url: '/setting',
                    template: '<div ui-view></div>'
                })
                .state('home.setting.log', {
                    url: '/log',
                    template: '<setting-log></setting-log>'
                });
        }])
})()