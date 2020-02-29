(function() {
    'use strict';
    /*
     * aurhor:广州银云信息科技有限公司
     * 状态码常量集
     */
    angular
        .module('eolinker.constant')
        .constant('RESPONSE_TEXT', {
            FAILURE: "请稍候再试或提交工单反馈"
        })
        .constant('ERROR_WARNING', {
            COMMON: '请稍候再试'
        })
        .constant('FILTER_WARNING_CODE_ARR_COMMON_CONST',[
            "210000",
            "230011",
            "230000"
        ])
        .constant('ERR_CODE_ARR_COMMON_CONST',{
            "230005":"当前已存在相同地址的节点",
            "260002":"当前已存在相同的负载名称",
            "510009":"添加失败，人数已满！",
            "280013":"操作失败，存在运行中的节点",
            "340000":"当前无可下载的报表",
            "120000":"旧密码错误",
            "210001":"插件优先级非法",
            "210002":"插件名称非法",
            "210003":"插件优先级已经存在",
            "210004":"插件名称已经存在",
            "210005":"插件名称不存在",
            "210009":"插件类型非法",
            "190005":"当前请求方式下，该URL已存在！"
        })
        .constant('CODE', {
            COMMON: {
                HAD_WARNING:'xxxxxx',
                SUCCESS: '000000', //请求成功
                UNLOGIN: '100001', //用户未登录
                SERVER_ERROR: '100000', //请求失败，服务器错误，稍后再试
                UNAUTH: '100002' //用户缺乏相应的操作权限
            }
        })
})();
