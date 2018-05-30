(function() {
    /*
     * author：广州银云信息科技有限公司
     *register/login/forget/os.guide/transaction/error
     * 全局navbar指令相关js
     */
    angular.module('goku')
        .component('eoNavbar1', {
            templateUrl: 'app/ui/navbar/nav1/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope'];

    function indexController($scope) {
        var vm = this;
    }

})();
