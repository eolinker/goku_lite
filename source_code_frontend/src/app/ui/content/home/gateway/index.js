(function() {
    'use strict';
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway', {
                    url: '/gateway',
                    template: '<eo-navbar3></eo-navbar3>' +
                        '<div class="home-content">' +
                        '    <div class="home-div">' +
                        '        <div ui-view></div>' +
                        '    </div>' +
                        '</div>'
                });
        }])
})();
