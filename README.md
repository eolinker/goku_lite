![GOKU API Gateway](https://camo.githubusercontent.com/f859a59b436a665a1551c2909393a91615344836/68747470733a2f2f646174612e656f6c696e6b65722e636f6d2f636f757273652f36486c4658786263323833333934376462666136383430626262613334383731346466626533333031386664366363 "GOKU API Gateway")

[![Gitter](https://badges.gitter.im/goku-api-gateway/community.svg)](https://gitter.im/goku-api-gateway/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/eolinker/goku-api-gateway)](https://goreportcard.com/report/github.com/eolinker/goku-api-gateway) [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3214/badge)](https://bestpractices.coreinfrastructure.org/projects/3214) ![](https://img.shields.io/badge/license-GPL3.0-blue.svg)

Goku API Gateway is a Golang-based microservice gateway that enables high-performance dynamic routing, multi-tenancy management, API access control, etc. It's also suitable for API management under micro-service system. 

Goku provides graphic interface and plug-in system to make configuration easier and expand more convenient.

# Summary / [中文介绍](https://github.com/eolinker/goku-api-gateway/blob/master/README_CN.md "中文介绍")

- [WhyGoku](#WhyGoku "WhyGoku")
- [Features](#Features "Features")
- [Benchmark](#Benchmark "Benchmark")
- [ConsolePreview](#ConsolePreview "ConsolePreview")
- [QuickStart](#QuickStart "QuickStart")
- [EnterpriseSupport](#EnterpriseSupport "EnterpriseSupport")
- [AboutUs](#AboutUs "AboutUs")
- [License](#License "License")

# Why Goku
 
Goku API Gateway is a microservice gateway that runs on the boundaries of enterprise system services. When you build websites, apps, IOT, and even API transactions, Goku API Gateway can help you extract duplicate components from your internal system and place them on the Goku gateway, such as user authorization, access control, traffic monitoring, firewalls, data cache, data conversion and so on.

Goku API Gateway CE provides comprehensive usage guide and customization guide. Goku is written in pure Go language, with good performance and scalability, and the built-in plug-in system enables enterprises to customize development for their own business.

Goku API Gateway also can combine with EOLINK API Studio to enhance API Management,API Monitor and Automated test.

All in all, Goku API Gateway enables enterprise to focus on their business.

[![Stargazers over time](https://starchart.cc/eolinker/goku-api-gateway.svg)](https://starchart.cc/eolinker/goku-api-gateway)

# Product Features
- **Cluster Management**：Mutiple  Goku API Gateway  node，Configuration information is automatically synchronized and can support multi-cluster deployment.
- **UI Management Background**: Manage various configurations of the network through clear UI.
- **Load balancing**: Load balancing for back-end servers.
- **Service Discovery**: Find back-end servers from registries such as Consul and Eureka.
- **Forwarding Agent**: Hide Real Backend Services by Forwarding Requests, Support Rest API, Webservice.
- **Multi-tenant management**: According to different access terminals or users.
- **Access Authentication**: Basic, API Key, etc.
- **API Monitor**：Request data statistics.
- **API Alert**: Support the webhook and email to alert abnormal services.
- **Health check**: Dynamic discovery of exceptional network joints and back-end nodes, automatically cut off forwarding traffic and transfer to other normal back-end services.
- **Exception auto-restart**: When a gateway node is abnormal, it will automatically attempt to restart.
- **Flexible transmit rules**: support fuzzy matching request path, support rewriting transmit path, etc.
- **Plug-in**: Plug-in system based on Go language can rapidly develop high-performance plug-ins.
- **Extension**: Gateway nodes have good processing performance, supporting the number of horizontal extension nodes to meet different performance requirements.
- **Log**: Detailed system log, http log, etc.
- **Open API**：Provide OPEN API for users to operate on the gateway for easy integration.
- ...

# Benchmark
[Benchmark Detail](https://help.eolinker.com/#/tutorial/?groupID=c-362&productID=19#tip7 "Benchmark Detail")

# Console Preview
[See Console Preview](https://github.com/eolinker/goku-api-gateway/blob/master/docs/CONSOLE_PREVIEW.md "See Console Preview")


# Quick Start
* [Deployment Tutorial](https://help.eolinker.com/#/tutorial/?groupID=c-351&productID=19 "Deployment Tutorial")
* [Quick Start](https://help.eolinker.com/#/tutorial/?groupID=c-307&productID=19 "Quick Start Tutorial")
* [Source Code Compilation](https://help.eolinker.com/#/tutorial/?groupID=c-350&productID=19 "Source Code Compilation")

# Enterprise Support
Goku API Gateway EE (Enterprise Version) has more powerful functions, plug-in libraries and professional technical support services. If you want to know more details, you can contact us in the following ways.
- Apply for free trial and demonstration of Enterprise Version：[Appointment trial](https://wj.qq.com/s2/2150032/4b5e "Appointment trial")
- Market Cooperation Mail：market@eolinker.com
- Purchase consultation Mail：sales@eolinker.com
- Help Center：[help.eolinker.com](help.eolinker.com "help.eolinker.com")
- QQ Group: 725853895

# About Us
EOLINK is a leading API management service provider, providing professional API research and development management, API automated test service, API monitor service, API gateway and other services for more than 3000 enterprises worldwide. It is the first enterprise to formulate API R&D management industry norms for ITSS.

Official website :[https://www.eolinker.com](https://www.eolinker.com "EOLINK Official Site")
Free download of PC client :[https://www.eolinker.com/pc](https://www.eolinker.com/pc/ "Free download of PC client")

# License
```
Copyright 2017-2019 Eolink Inc.

Licensed under the GNU General Public License v3.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.gnu.org/licenses/gpl-3.0.html

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.
```
