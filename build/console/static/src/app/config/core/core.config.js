(function() {
    'use strict';
    //author：广州银云信息科技有限公司
    angular
        .module('eolinker')
        .config(coreConfig);

    coreConfig.$inject = ['$controllerProvider', '$compileProvider', '$filterProvider', '$provide'];

    function coreConfig($controllerProvider, $compileProvider, $filterProvider, $provide) {

        var data = {
            info: {
                core: angular.module('eolinker')
            }
        };

        // registering components 
        data.info.core.controller = $controllerProvider.register; //动态定义controller
        data.info.core.directive = $compileProvider.directive;
        data.info.core.filter = $filterProvider.register;
        data.info.core.factory = $provide.factory;
        data.info.core.service = $provide.service;
        data.info.core.constant = $provide.constant;
        data.info.core.value = $provide.value;

    }

})();
