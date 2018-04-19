(function() {
    'use strict';

    angular
        .module('goku')
        .config(lazyloadConfig);

    lazyloadConfig.$inject = ['$ocLazyLoadProvider', 'APP_REQUIRES'];

    function lazyloadConfig($ocLazyLoadProvider, APP_REQUIRES) {
        var template = {
            modules: APP_REQUIRES.MODULES
        }
        // Lazy Load modules configuration
        $ocLazyLoadProvider.config({
            debug: false,
            events: true,
            modules: template.modules
        });

    }
})();