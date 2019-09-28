![Goku API Gateway 悟空网关](https://data.eolinker.com/course/gBTEV2s29e16630bb4dc553bec35ad33914d19aa410a8bf "Goku API Gateway 悟空网关")

[![Gitter](https://badges.gitter.im/goku-api-gateway/community.svg)](https://gitter.im/goku-api-gateway/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/eolinker/goku-api-gateway)](https://goreportcard.com/report/github.com/eolinker/goku-api-gateway) [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3214/badge)](https://bestpractices.coreinfrastructure.org/projects/3214) ![](https://img.shields.io/badge/license-GPL3.0-blue.svg)

Goku API Gateway （中文名：悟空 API 网关）是一个基于 Golang 开发的微服务网关，能够实现高性能 HTTP API 转发、多租户管理、API 访问权限控制等目的，拥有强大的自定义插件系统可以自行扩展，并且提供友好的图形化配置界面，能够快速帮助企业进行 API 服务治理、提高 API 服务的稳定性和安全性。

# 概况

- [为什么要使用Goku](#为什么要使用Goku "为什么要使用Goku")
- [产品特性](#产品特性 "产品特性")
- [产品截图](#产品截图 "产品截图")
- [安装使用](#安装使用 "安装使用")
- [企业支持](#企业支持 "企业支持")
- [关于我们](#关于我们 "关于我们")
- [授权协议](#授权协议 "授权协议")

# 为什么要使用Goku
Goku API Gateway （悟空 API 网关）是运行在企业系统服务边界上的微服务网关。当您构建网站、App、IOT甚至是开放API交易时，Goku API Gateway 能够帮你将内部系统中重复的组件抽取出来并放置在Goku网关上运行，如进行用户授权、访问控制、流量监控、防火墙、静态数据缓存、数据转换等。

Goku API Gateway 的社区版本（CE）拥有完善的使用指南和二次开发指南，代码使用纯 Go 语言编写，拥有良好的性能和扩展性，并且内置的插件系统能够让企业针对自身业务进行定制开发。

并且 Goku API Gateway 支持与 EOLINKER 旗下的 API Studio 接口管理平台结合，对 API 进行全面的管理、自动化测试、监控和运维。

总而言之，Goku API Gateway 能让业务开发团队更加专注地实现业务。

[![Stargazers over time](https://starchart.cc/eolinker/goku-api-gateway.svg)](https://starchart.cc/eolinker/goku-api-gateway)

# 产品特性
- **集群管理**：多个 Goku API Gateway 节点，配置信息自动同步，支持多集群部署。
- **界面管理后台**：通过清晰的UI界面对网关的各项配置进行管理。
- **负载均衡**：对后端服务器进行负载均衡。
- **服务发现**：从 Consul、Eureka 等注册中心发现后端服务器。
- **转发代理**：通过转发请求来隐藏真实后端服务，支持 Rest API、Webservice。
- **多租户管理**：根据不同的访问终端或用户来判断。
- **访问鉴权**：Basic、API Key等。
- **API监控**：请求数据统计。
- **API告警**：支持通过API、邮件方式对异常的服务进行告警。
- **健康检查**：动态发现异常的网关节点以及后端节点，自动切断转发流量并转到其他正常后端服务。
- **异常自动重启**：网关节点异常时会自动尝试重载重启。
- **灵活的转发规则**：支持模糊匹配请求路径，支持改写转发路径等。
- **插件系统**：基于 Go 语言的插件系统，可以快速开发高性能的插件。
- **性能扩展**：网关节点拥有良好的处理性能，支持水平扩展节点数量满足不同的性能需求。
- **日志**：详细的系统日志、请求日志等。
- **Open API**：提供 API 对网关进行操作，便于集成。
- ...

# 产品截图
* 【首页】
首页可以了解网关的基本信息，例如访问策略数、API数等，还可以了解请求和转发的情况，例如成功率等。

![](http://data.eolinker.com/course/p8qL49u6c8adce6b345915b3fd77bf5812a40fe7dd0a8a2)

* 【网关节点】
网关支持集群化管理，进入不同的集群可以管理相应的节点。

![](http://data.eolinker.com/course/wEa9yEI2bf086f3873b55bbdaec32f3b4ce1eb23dfe44ea)

* 【服务注册方式】
您可以通过静态或动态的方式来注册（发现）您的后端服务，创建好服务注册方式后，您可以在某个方式的基础上创建一个或多个负载（Upstream）。

![](http://data.eolinker.com/course/1elb5mF4d3fd6141919001293e0119557b3d5ef0cea0719)

* 【负载配置】
配置API的转发目标服务器（负载后端），创建之后可以设置为 API 的转发地址 / 负载后端（Target / Upstream）。

![](http://data.eolinker.com/course/4tHYXR23abc26b914ca763aac4871ed9d60a3aeb819941f)

* 【接口管理】
支持创建并管理API文档，并且支持导入API文档项目。

![](http://data.eolinker.com/course/WlTJ2kB1cd03ddf839ea1d489890a0bd5b0572efeff6043)

* 【访问策略】
您可以给不同的调用方或应用设置访问策略，不同的访问策略可以设置不同的 API 访问权限、鉴权方式以及插件功能等。

![](http://data.eolinker.com/course/fUrHmVd0d2d88b7f72d985b0e93e434ed528648d2dd34db)

* 【告警设置】
针对异常API可以设置告警提醒，支持邮件和Webhook通知。

![](http://data.eolinker.com/course/9eQ3Lmv64e5cedc1ad4745dfa2895f6657441d874f6c7f4)

* 【扩展插件】
插件系统除了提供官方插件，也可以添加自定义的网关插件。

![](http://data.eolinker.com/course/sQhUflpcebf65dc43cb7e2e838e8d1ecf3e52e9a5a6c566)

* 【日志设置】
提供详细的请求日志和系统运行日志，请求日志可以自定义记录字段；运行日志可以根据情况调整记录等级：ERROR、INFO、DEBUG等。

![](http://data.eolinker.com/course/iyifFJ2809fe63e27df709ddc1a22f94d983c5ecbf8cc29)

# 安装使用
* [部署教程](https://help.eolinker.com/#/tutorial/?groupID=c-351&productID=19 "部署教程")
* [快速入门教程](https://help.eolinker.com/#/tutorial/?groupID=c-307&productID=19 "快速入门教程")
* [源码编译教程](https://help.eolinker.com/#/tutorial/?groupID=c-350&productID=19 "源码编译")

# 企业支持
Goku API Gateway EE（企业版本）拥有更强大的功能、插件库以及专业的技术支持服务，如您需要了解可以通过以下方式联系我们。
- **中国大陆服务支持电话**：400-616-0330 法定工作日（9:30-18:00）
- **申请企业版免费试用及演示**：[预约试用](https://wj.qq.com/s2/2150032/4b5e "预约试用")
- **市场合作邮箱**：market@eolinker.com
- **购买咨询邮箱**：sales@eolinker.com
- **帮助文档**：[help.eolinker.com](help.eolinker.com "help.eolinker.com")
- **QQ群**: 725853895

# 关于我们
EOLINKER 是领先的 API 管理服务供应商，为全球超过3000家企业提供专业的 API 研发管理、API自动化测试、API监控、API网关等服务。是首家为ITSS（中国电子工业标准化技术协会）制定API研发管理行业规范的企业。

官方网站：[https://www.eolinker.com](https://www.eolinker.com "EOLINKER官方网站")
免费下载PC桌面端：[https://www.eolinker.com/pc/](https://www.eolinker.com/pc/ "免费下载PC客户端")

# 授权协议
```
Copyright 2017-2019 Eolinker Inc.

Licensed under the GNU General Public License v3.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.gnu.org/licenses/gpl-3.0.html

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.
```
