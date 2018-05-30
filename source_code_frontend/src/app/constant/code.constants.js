(function() {
    'use strict';
    /*
     * aurhor:广州银云信息科技有限公司
     * 状态码常量集
     */
    angular
        .module('goku.constant')
        .constant('CODE', {
            COMMON: {
                SUCCESS: '000000', //请求成功
                UNLOGIN: '100000', //用户未登录
                UNAUTH: '100002' //用户缺乏相应的操作权限
            },
            USER: {
                ERROR: '120000', //操作失败，处理失败及服务器出错
                ILLIGLE_PASSWORD: '120002', //密码格式非法
                ERROR_PASSWORD: '120003', //用户不存在或用户名密码错误
                ILLIGLE_INFO: '120004', //用户登录信息格式非法，或者不为（手机号、用户名、邮箱）之一
                ERROR_LOGIN: '120005', //用户尚未登录
                EXIST: '120006', //（手机号、用户名、邮箱）已存在
                UNCHANGE: '120007', //新旧密码相同
                ILLIGLE_NICKNAME: '120008', //用户昵称格式非法，长度须小于20位
                ILLIGLE_NAME: '120009', //用户名格式非法
            },
            SMS: {
                ILLEGAL_PHONE: "110001", //接收短信的手机号格式非法
                UNMATCHED_PHONE: "110002", //手机号码与获取验证码的号码不一致    
                SMS_LIMIT: "110003", //一天只能发送10条，运营商限制
                SMS_NOT_ENOUGHT_MONEY: "110004", //账户余额不足
                ILLEGAL_CHECKCODE: "110005" //未获取验证码、验证码非数字组合、验证码不正确或者验证码已过期(300s)
            },
            API_TEST: {
                ERROR: '210000', //操作失败/列表为空等    
                ILLEGAL_URI: '210001', //URI地址格式非法     
                ILLEGAL_REQUEST_TYPE: '210002', //请求参数的类型非法     
                ERROR_ADD_HISTORY: '210003', //添加测试记录失败 
                ILLEGAL_HISTORY_ID: '210004', //测试记录ID格式非法  
            },
            API_GROUP: {
                ERROR: '150000', //操作失败   
                ILLEGAL_NAME: '150001', //分组名称格式非法    
                ILLEGAL_PARENT_ID: '150002', //父分组ID格式非法 
                ILLEGAL_ID: '150003' //分组ID格式非法
            },
            PARTNER: {
                ERROR: '250000', //操作失败/列表为空等    
                ILLEGAL_USERCALL: '250001', //目标邀请人员的信息填写有误，不为（手机号、邮箱、用户名）之一
                EXIST: '250002', //该用户已经被邀请过     
                ILLEGAL_ID: '250003', //协作成员的关联ID（connID）格式非法    
                ILLEGAL_NICKNAME: '250004', //协作成员的昵称格式非法  
                ILLEGAL_TYPE: '250005' //设置的协作成员类型有误    
            },
            MESSAGE: {
                ERROR: '260000', //操作失败/列表为空等      
                ILLEGAL_ID: '260001' //消息ID格式非法    
            },
            ENV: {
                ERROR: '170000', //操作失败/列表为空等    
                ILLEGAL_NAME: '170001', //环境名称格式非法   
                ILLEGAL_ID: '170002', //环境ID格式非法        
                ILLEGAL_URI: '170003', //前置URI地址格式非法      
                ILLEGAL_HEADER_ID: '170004' //header头部ID格式非法   
            },
            PROJECT: {
                ERROR: "140000", //请求失败 
                ILLEGAL_PROJECT_NAME: "140001", //项目名称格式不合法 
                ILLEGAL_PROJECT_TYPE: "140002", //项目类型不合法   
                ILLEGAL_PROJECT_VERSION: "140003", //项目版本不合法    
                ILLEGAL_PROJECT_ID: "140004", //项目ID不合法 
                ILLEGAL_PROJECT_DESCRIPTION: "140005", //项目描述长度不合法  
                ILLEGAL_PROJECT_SHARE_STATUS: "140006", //分享状态格式非法，只能为0/1   
                ILLEGAL_PROJECT_LOCK_STATUS: "140007" //加密选项不合法，只能为0/1  
            },
            PROJECT_API: {
                ERROR: '160000', //操作失败/列表为空等 
                ILLEGAL_ID: '160001', //接口ID格式非法  
                EXIST: '160002', //已存在相同接口    
                ILLEGAL_SEARCH: '160003', //搜索关键字格式非法 
                ILLEGAL_HISTORY_ID: '160004' //接口历史记录ID格式非法
            },
            STATUS_CODE: {
                ERROR: '190000', //操作失败/列表为空等 
                ILLEGAL_NAME: '190001', //状态码格式非法   
                ILLEGAL_DESC: '190002', //状态码描述格式非法      
                ILLEGAL_ID: '190003', //状态码ID格式非法   
                ILLEGAL_SEARCH: '190004' //状态码搜索提示格式非法  
            },
            STATUS_CODE_GROUP: {
                ERROR: '180000', //操作失败/列表为空等 
                ILLEGAL_NAME: '180001', //分组名称格式非法    
                ILLEGAL_ID: '180002', //分组ID格式非法       
                ILLEGAL_PARENT_ID: '180003' //父分组ID格式非法 
            },
            DOC: {
                ERROR: '230000', //操作失败/列表为空等 
                ILLEGAL_GROUP_ID: '230001', //文档分组ID格式非法   
                ILLEGAL_DESC: '230002', //文档描述格式不合法，必须为[0/1]=>[富文本/MARKDOWN]      
                ILLEGAL_ID: '230003', //文档ID格式非法   
                ILLEGAL_SEARCH: '230004', //搜索关键字长度非法，必须为1-255 
                GROUP: {
                    ERROR: '220000', //操作失败/列表为空等 
                    ILLEGAL_NAME: '220001', //分组名称格式非法    
                    ILLEGAL_ID: '220003', //分组ID格式非法       
                    ILLEGAL_PARENT_ID: '220002' //父分组ID格式非法 
                }
            },
            IMPORT_EXPORT: {
                ERROR: '310000', //导入/导出失败
                EMPTY: '310001', //导入的数据为空
                ILLEGAL_VERSION: '310002', //导入postman数据缺少版本号[1/2]或者版本号不正确
                ILLEGAL_IMPORT: '310003' //导入的RAP数据为空或者数据格式有误
            },
            EMPTY: '150008'
        })
        .constant('ERROR_WARNING', {
            COMMON: '请联系您的客服人员或者刷新进行重试'
        })
})();
