(function () {
    'use strict';
    /**
     * @name 侧边栏服务
     * @author 广州银云信息科技有限公司
     */
    angular.module('eolinker')
        .factory('Sidebar_CommonService', index);

    index.$inject = []

    function index() {
        var data={}
        try{
            data.isShrink=JSON.parse(window.localStorage.getItem('EO_CONFIG_SIDEBAR_SHRINK_STATUS'))||false;
        }catch(JSON_PARSE_ERR){
            data.isShrink=false;
        }
        return data;
    }
})();