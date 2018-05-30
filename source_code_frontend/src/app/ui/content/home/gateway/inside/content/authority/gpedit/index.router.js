(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块backend相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.gpedit', {
                    url: '/gpedit?strategyID',
                    template: '<div class="common-table-div" ui-view></div>'
                })
                .state('home.gateway.inside.gpedit.default', {
                    url: '/',
                    template: '<gateway-gpedit-default></gateway-gpedit-default>'
                })
                .state('home.gateway.inside.gpedit.ip', {
                    url: '/ip',
                    template: '<gateway-gpedit-ip></gateway-gpedit-ip>',
                    mark:'ip'
                })
                .state('home.gateway.inside.gpedit.rate', {
                    url: '/rate',
                    template: '<gateway-gpedit-rate></gateway-gpedit-rate>',
                    mark:'rate'
                })
                .state('home.gateway.inside.gpedit.auth', {
                    url: '/auth',
                    template: '<gateway-gpedit-auth></gateway-gpedit-auth>',
                    mark:'auth',
                    resolve: helper.resolveFor('CLIPBOARD')
                });
        }])

})();
