(function () {
    /**
     * @name 日期弹窗
     * @author 广州银云信息科技有限公司
     */

    'use strict';
    angular.module('eolinker.directive')

        .directive('datepickerTimesDirective', ['$document', '$rootScope', function ($document, $rootScope) {
            return {
                restrict: 'AE',
                transclude: true,
                templateUrl: 'app/directive/common/datepicker/index.html',
                scope: {
                    interaction: '=',
                    datepickerTimesDirective: '&'
                },
                link: function ($scope, elem, attrs, ctrl) {
                    $scope.data = {
                        datePicker: {
                            duration: {
                                option: {
                                    dateDisabled: null,
                                    formatYear: 'yy',
                                    startingDay: 1
                                }
                            }
                        }
                    }
                    $scope.request = {
                        startTime: "",
                        endTime: "",
                    }
                    $scope.fun = {};
                    var fun = {};
                    $scope.fun.confirm = function (arg) {
                        if (arg) arg.$event.stopPropagation();
                        $scope.interaction.request.startTime = $scope.request.startTime;
                        $scope.interaction.request.endTime = $scope.request.endTime;
                        $scope.datepickerTimesDirective();
                    }
                    $scope.fun.close = function (arg) {
                        if (arg) arg.$event.stopPropagation();
                        $scope.interaction.show = false;
                    }
                    $document.on("click", function (_default) {
                        $scope.fun.close();
                        $scope.$root && $scope.$root.$$phase || $scope.$apply();
                    });
                    $scope.data.datePicker.duration.option.customClass = function (_default) {
                        if (_default.mode === 'day' && !(_default.date.getMonth() === _default.currentMonth)) {
                            return 'uib-day-disabled';
                        }
                    }
                    fun.init = function () {
                        $scope.request.startTime = $scope.interaction.request.startTime;
                        $scope.request.endTime = $scope.interaction.request.endTime;
                    }
                    fun.init();
                    $rootScope.global.$watch.push($scope.$watch('interaction.show', function () {
                        if ($scope.interaction.show) {
                            fun.init();
                            $scope.$broadcast('$RESET_DATEPICKER')
                        }
                    }));
                }
            };
        }]);
})();