(function () {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 温馨提示组件
     */
    angular.module('eolinker')
        .component('tipCommonComponent', {
            templateUrl: 'app/component/common/tip/index.html',
            controller: indexController,
            bindings: {
                version: '@',
                status: '@',
                user: '@',
                interaction: '<'
            }
        })

    indexController.$inject = ['$scope', '$window', '$state'];

    function indexController($scope, $window, $state) {
        var vm = this;
        vm.data = {
            balance: {
                general: {
                    title: '注意事项',
                    titleStyle:{
                        float:'none'
                    },
                    class:'warning-ul common-ul',
                    content: '<div class="lh_1point75">'+
                             '1. 静态后端支持：IP+端口或域名，如：127.0.0.1:8080 或 www.eolinker.com    权重范围：0-999<br/>'+
                             '2. 静态IP列表，IP与权重之间使用空格分隔，多个IP之间使用英文分号分隔，例如：127.1.1.1:8080 10;<br/>'+
                             '3. 若IP不加权重，则默认该IP权重为1。例如：“ip1:1111 ; ip2:2222 10” ，这里的ip1权重为1，ip2权重为10</div>'
                }
            },
            plug: {
                general: {
                    title: '注意事项',
                    titleStyle:{
                        float:'none'
                    },
                    class:'warning-ul common-ul',
                    content: '<p style="line-height:1em;">1. 如需让自定义的插件生效，必须先重启/重载网关。</p><span>2. 检测插件仅针对自定义插件，用于检测该自定义插件是否可用</span>'
                },
                operate:{
                    title:'温馨提示',
                    class:'warning-ul common-ul',
                    content:'<span>新增自定义插件后，请立即对该插件进行检测，确保插件可用后再重载或重启网关节点，以使插件生效</span>'
                },
                official:{
                    title:'温馨提示',
                    class:'warning-ul common-ul',
                    content:'<span>若要使网关类型插件的最新配置生效，保存后须 重载/重启 网关节点</span>'
                }
            },
            gepditAuth: {
                general: {
                    title: '温馨提示',
                    class:'warning-ul common-ul',
                    content: '尚未对访问策略设置鉴权方式，请在 <a class="eo_link" ui-sref="home.gpedit.inside.plugin.gpedit">策略插件</a> 处添加相应的鉴权插件'
                }
            },
            authorityConfig: {
                general: {
                    title: '权限提示',
                    class:'warning-ul common-ul',
                    content: '<span>系统管理员以及管理员用户拥有最高读写操作权限，因此不会出现在权限管理的列表中，列表仅会列出普通成员</span>'
                }
            }
        }
        vm.service = {
            $window: $window
        }

    }
})();