(function() {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 侧边栏公用服务
     * @required GroupService 
     */
    angular.module('goku.service')
        .factory('Group_GatewayCommonService', index);

    index.$inject = ['GroupService', 'CODE', '$rootScope','$filter']

    function index(GroupService, CODE, $rootScope,$filter) {
        var data = {
            service: GroupService,
            fun: {
                clear: null,
                spreed: null,
                operate: null
            },
            sort: {
                operate: null,
                init: null,
            }
        }


        /**
         * 分组操作
         * @param {string} status 操作状态
         * @param {object} arg 传参
         * @param {object} options 选项{callback:回调函数(选填),resource:请求资源,originGroupQuery:原始分组队列,status:状态（child），默认父分组}
         */
        data.fun.operate = function(status, arg, options) {
            var template = {
                modal: {},
                $index: null
            }
            switch (status) {
                case 'edit':
                    {
                        template.modal = {
                            title: (options.status.indexOf('edit') > -1 ? '修改' : '新增') + (options.status.indexOf('child') > -1 ? '子分组' : '分组'),
                            secondTitle: '分组名称',
                            group: options.status.indexOf('parent-edit') > -1 ? null : options.originGroupQuery
                        }
                        $rootScope.GroupModal(template.modal.title, arg.item, template.modal.secondTitle, template.modal.group, function(callback) {
                            if (callback) {
                                angular.merge(callback, callback, options.baseRequest);
                                switch (options.status) {
                                    case 'parent-edit':
                                        {
                                            break;
                                        }
                                    default:
                                        {
                                            template.$index = parseInt(callback.$index) - 1;
                                            if (template.$index > -1) {
                                                callback.parentGroupID = options.originGroupQuery[template.$index].groupID;
                                            }
                                            break;
                                        }
                                }
                                if (options.status.indexOf('edit') > -1) {
                                    $filter('CurrentTime_CommonFilter')('object',callback,['$index','isAdd']);
                                    options.resource.Edit(callback).$promise.then(function(response) {
                                        switch (response.statusCode) {
                                            case CODE.COMMON.SUCCESS:
                                                {
                                                    $rootScope.InfoModal(template.modal.title + '成功', 'success');
                                                    options.callback();
                                                    break;
                                                }
                                        }
                                    });
                                } else {
                                    $filter('CurrentTime_CommonFilter')('object',callback,['$index','groupID','isAdd']);
                                    options.resource.Add(callback).$promise.then(function(response) {
                                        switch (response.statusCode) {
                                            case CODE.COMMON.SUCCESS:
                                                {
                                                    $rootScope.InfoModal(template.modal.title + '成功', 'success');
                                                    options.callback();
                                                    break;
                                                }
                                        }
                                    });
                                }
                            }
                        });
                        break;
                    }
                
            }
        }

        /**
         * 清空分组服务
         */
        data.fun.clear = function() {
            data.service.clear();
        };

        return data;
    }
})();
