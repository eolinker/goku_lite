![](https://data.eolinker.com/6IvqUL3cb40efeca8cf4fdc034286bd946b130b45d50bd8.jpg)

[![Gitter](https://badges.gitter.im/goku-api-gateway/community.svg)](https://gitter.im/goku-api-gateway/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/eolinker/goku-api-gateway)](https://goreportcard.com/report/github.com/eolinker/goku-api-gateway) [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3214/badge)](https://bestpractices.coreinfrastructure.org/projects/3214) ![](https://img.shields.io/badge/license-GPL3.0-blue.svg)

Goku API Gateway （中文名：悟空 API 网关）是一个基于 Golang 开发的微服务网关，能够实现高性能 HTTP API 转发、多租户管理、API 访问权限控制等目的，拥有强大的自定义插件系统可以自行扩展，并且提供友好的图形化配置界面，能够快速帮助企业进行 API 服务治理、提高 API 服务的稳定性和安全性。

# 概况

- [为什么要使用Goku](#为什么要使用Goku "为什么要使用Goku")
- [产品特性](#产品特性 "产品特性")
- [为什么要做Goku网关](#为什么要做Goku网关 "为什么要做Goku网关")
- [基准测试](#基准测试 "基准测试")
- [产品截图](#产品截图 "产品截图")
- [安装使用](#安装使用 "安装使用")
- [企业支持](#企业支持 "企业支持")
- [关于我们](#关于我们 "关于我们")
- [授权协议](#授权协议 "授权协议")

# 为什么要使用Goku
Goku API Gateway （悟空 API 网关）是运行在企业系统服务边界上的微服务网关。当您构建网站、App、IOT甚至是开放API交易时，Goku API Gateway 能够帮你将内部系统中重复的组件抽取出来并放置在Goku网关上运行，如进行用户授权、访问控制、防火墙、数据转换等；并且Goku 提供服务编排的功能，让企业可以快速从各类服务上获取需要的数据，对业务实现快速响应。

Goku API Gateway 的社区版本（CE）拥有完善的使用指南和二次开发指南，代码使用纯 Go 语言编写，拥有良好的性能和扩展性，并且内置的插件系统能够让企业针对自身业务进行定制开发。

并且 Goku API Gateway 支持与 EOLINKER 旗下的 API Studio 接口管理平台结合，对 API 进行全面的管理、自动化测试、监控和运维。

总而言之，Goku API Gateway 能让业务开发团队更加专注地实现业务。

[![Stargazers over time](https://starchart.cc/eolinker/goku-api-gateway.svg)](#)

# 产品特性
- **控制台**：通过清晰的UI界面对网关集群进行各项配置。
- **集群管理**：Goku网关节点是无状态的，配置信息自动同步，支持节点水平拓展和多集群部署。
- **热更新**：无需重启服务，即可持续更新配置和插件。
- **服务编排**：一个编排API对应多个backend，backend的入参支持客户端传入，也支持backend间的参数传递；backend的返回数据支持字段的过滤、删除、移动、重命名、拆包和封包；编排API能够设定编排调用失败时的异常返回。
- **数据转换**：支持将返回数据转换成JSON或XML。
- **负载均衡**：支持有权重的round-robin负载平衡。
- **服务发现**：从 Consul、Eureka 等注册中心发现后端服务器。
- **HTTP(S)反向代理**：隐藏真实后端服务，支持 Rest API、Webservice。
- **多租户管理**：根据不同的访问终端或用户来判断。
- **访问策略**：支持不同策略访问不同的API、配置不同的鉴权（匿名、Apikey、Basic）等。
- **灵活的转发规则**：支持模糊匹配请求路径，支持改写转发路径等，可为不同访问策略或集群设置不同的负载。
- **IP黑白名单**。
- **自定义插件**：允许插件挂载在常见阶段，例如before match，access和proxy。
- **CLI**: 使用命令行来启动、关闭和重启Goku。
- **Serverless**: 在转发过程的每一个阶段，都可以添加并调用自定义的插件。
- **请求日志(access log)**：仅记录转发的基本内容，自定义记录字段与排序顺序，定期自动清理日志。
- **运行日志(system log)**：提供控制台和节点的运行日志，默认仅记录ERROR等级的信息，可将等级按实际情况调成INFO、WARN或DEBUG。
- **可扩展**：简单易用的插件机制方便扩展功能。
- **高性能**：性能在众多网关之中表现优异。
- **Open API**：提供 API 对网关进行操作，便于集成。
- **版本控制**：支持操作的发布和多次回滚。
- **监控和指标**：支持Prometheus、Graphite。

#### 迭代计划
- **Open Tracing**：支持Zipkin
- **动态路由**：不同参数值不同转发
- **gRPC 协议转换**：支持协议的转换，客户端可以通过 HTTP/JSON 来访问 gRPC API

# 为什么要做Goku网关
我们 EOLINKER 自2017年成立以来，立志于做全球领先的 API 管理平台，我们先是做了目前国内最大的在线API管理平台（API Studio），然后在18年发布了支持API场景（多个API关联和数据传递）的API监控（API Beacon），今年我们在思考还能为企业客户提供什么更加深度的服务时，认为API网关是一个关键的环节，能够帮助企业综合管理企业内部的微服务API、更方便地对接第三方API以及更好地维护对外的API等。

可以说API网关是我们在深入API管理领域几年之后自然而然要做的事情，而既然要做就努力往大了做，于是我们做了更加大胆的决定：将核心代码全部开源，并且不限制网关的节点，还提供了完整的管理界面，让用户可以部署完成后立即投入使用。

可能有人不理解为什么开源代码是一个大胆的决定，首先我们是一个商业公司而不是公益开源基金会，开源意味着有一大部分收入的流失，其次放眼全球的开源产品几乎都是不盈利的，每年还需要投入大量的研发和维护成本等。

**那我们为什么还要将一个公司的核心产品开源？**

因为一个公司的力量实在有限，如果我们希望把 Goku API Gateway 做到全球一流的水平，将中国的技术产品输出到海外去，开源社区和开发者的力量是必不可少的，因此这产品里面包含着我们的希望和情怀，希望证明在中国，像我们一样专注基础技术领域的企业也能有好的未来。所幸的是我们并不孤独，在我们前面有 Dubbo、TiDB 等优秀的开源项目，相信他们也和我们一样抱有希望在做着类似的事情。

因此我们将 Goku API Gateway 开源，正如它的中文名称 “悟空” 一般，能在开源社区和我们的共同努力下完成72变。

# 基准测试
![](https://data.eolinker.com/p7NFG6lb4c73b26cc880e838fe45aa31bc037b7415e3770.jpg)

[基准测试详情](https://help.eolinker.com/#/tutorial/?groupID=c-362&productID=19#tip7 "Benchmark Detail")

# 产品截图

[查看产品截图](https://github.com/eolinker/goku-api-gateway/blob/master/docs/CONSOLE_PREVIEW_CN.md "查看产品截图")

# 安装使用
* 直接部署：[部署教程](https://help.eolinker.com/#/tutorial/?groupID=c-371&productID=19 "部署教程")
* Docker部署：[控制台Docker](https://hub.docker.com/r/eolinker/goku-api-gateway-ce-console "控制台Docker")、[网关节点Docker](https://hub.docker.com/r/eolinker/goku-api-gateway-ce-node "网关节点Docker")
* [快速入门教程](https://help.eolinker.com/#/tutorial/?groupID=c-307&productID=19 "快速入门教程")
* [源码编译教程](https://help.eolinker.com/#/tutorial/?groupID=c-350&productID=19 "源码编译")

# 企业支持
Goku API Gateway EE（企业版本）拥有更强大的功能、插件库以及专业的技术支持服务，如您需要了解可以通过以下方式联系我们。
- **中国大陆服务支持电话**：400-616-0330 法定工作日（9:30-18:00）
- **申请企业版免费试用及演示**：[预约试用](https://www.eolinker.com/#/survey/applyAmsCloud "预约试用")
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
