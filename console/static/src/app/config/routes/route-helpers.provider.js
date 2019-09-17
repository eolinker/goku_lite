/**=========================================================
 * Module: helpers.js
 * Provides helper functions for routes definition
 =========================================================*/

(function() {
    'use strict';

    angular
        .module('eolinker')
        .provider('RouteHelpers', RouteHelpersProvider)

    RouteHelpersProvider.$inject = ['APP_REQUIRES'];


    function RouteHelpersProvider(APP_REQUIRES) {
        /* jshint validthis:true */
        var data = {
            info:{
                query:{
                    MODULES:APP_REQUIRES.MODULES,
                    SCRIPTS:{}
                }
            },
            fun: {
                // provider access level
                basepath: null, // Set here the base of the relative path;for all app views
                resolveFor: null, // Generates a resolve object by passing script names;previously configured in constant.APP_REQUIRES
                $get: null // controller access level；resolveFor
            }
        }
        data.fun.$get = function() {
            return {
                basepath: data.fun.basepath,
                resolveFor: data.fun.resolveFor
            };
        }

        data.fun.basepath = function(uri) {
            return 'app/' + uri;
        }

        data.fun.resolveFor = function() {
            var _args = arguments;
            return {
                deps: ['$ocLazyLoad', '$q', function($ocLL, $q) {
                    // Creates a promise chain for each argument
                    var promise = $q.when(1); // empty promise
                    for (var i = 0, len = _args.length; i < len; i++) {
                        promise = andThen(_args[i]);
                    }
                    return promise;

                    // creates promise to chain dynamically
                    function andThen(_arg) {
                        // also support a function that returns a promise
                        if (typeof _arg === 'function')
                            return promise.then(_arg);
                        else
                            return promise.then(function() {
                                // if is a module, pass the name. If not, pass the array
                                var whatToLoad = getRequired(_arg);
                                // simple error check
                                if (!whatToLoad) return $.error('Route resolve: Bad resource name [' + _arg + ']');
                                // finally, return a promise
                                return $ocLL.load(whatToLoad);
                            });
                    }
                    // check and returns required data
                    // analyze module items with the form [name: '', files: []]
                    // and also simple array of script files (for not angular js)
                    function getRequired(name) {
                        if (data.info.query.MODULES)
                            for (var m in data.info.query.MODULES)
                                if (data.info.query.MODULES[m].name && data.info.query.MODULES[m].name === name)
                                    return data.info.query.MODULES[m];
                        return data.info.query.SCRIPTS && data.info.query.SCRIPTS[name];
                    }

                }]
            };
        }
        return data.fun;
    }


})();
