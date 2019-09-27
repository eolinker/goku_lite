![GOKU API Gateway](https://camo.githubusercontent.com/f859a59b436a665a1551c2909393a91615344836/68747470733a2f2f646174612e656f6c696e6b65722e636f6d2f636f757273652f36486c4658786263323833333934376462666136383430626262613334383731346466626533333031386664366363 "GOKU API Gateway")

[![Gitter](https://badges.gitter.im/goku-api-gateway/community.svg)](https://gitter.im/goku-api-gateway/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/eolinker/goku-api-gateway)](https://goreportcard.com/report/github.com/eolinker/goku-api-gateway) [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3214/badge)](https://bestpractices.coreinfrastructure.org/projects/3214) ![](https://img.shields.io/badge/license-GPL3.0-blue.svg)

Goku API Gateway is a Golang-based microservice gateway that enables high-performance dynamic routing, multi-tenancy management, API access control, etc. It's also suitable for API management under micro-service system. 

Goku provides graphic interface and plug-in system to make configuration easier and expand more convenient.

# Summary / [中文介绍](https://github.com/eolinker/goku-api-gateway/blob/master/README_CN.md "中文介绍")

- [WhyGoku](#WhyGoku "WhyGoku")
- [Features](#Features "Features")
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

# Console Preview
* 【Home Page】
Home page can help users understand the basic information of gateway, such as access policy, API number, etc. It can also help understand the situation of request and forwarding, such as success rate.

![](http://data.eolinker.com/course/nCN4Qifbe6f7d197c26dadae4248664ce30693061049f0f)

* 【Gateway Note】
Gateway supports cluster management, and access to different clusters can manage the corresponding nodes.

![](http://data.eolinker.com/course/gJdazCFd5207d6b3b2c8d63cf613e8684a5ce1f3da506fc)

* 【Service Registration Method】
You can register (discover) your back-end services in a static or dynamic way. After creating a service registration method, you can create one or more loads on the basis of one way or another.（Upstream）。

![](http://data.eolinker.com/course/Ny7TmGRaf427ef3b63bae01d7856884247d7a11df865803)

* 【Load Configuration】
Configure the API's transmit target server (load back-end), which can be set to the API's transmit address/load back-end after creation（Target / Upstream）。

![](https://camo.githubusercontent.com/2dd6e6c88049dd7182cd2dc81b745a8d96856423/687474703a2f2f646174612e656f6c696e6b65722e636f6d2f636f757273652f4655414b45413764656137373335653961313534356636373764333430313064653838623535633637636336356463)

* 【API Mangement】
Support the creation and management of API documents, and support the import of API document projects.

![](http://data.eolinker.com/course/7nb8KKEafffa070b5e510b67b1eeb1027c16654bc72f464)

* 【Strategy】
You can set access strategies for different callers or applications. Different access strategies can set different API access rights, authentication and plug-in functions.

![](http://data.eolinker.com/course/e122iUe133714876f2cce05e591dda7adb9e5501ebf7b27)

* 【Alarm Settings】
Alerts can be set for exception APIs to support email and Webhook notifications.

![](http://data.eolinker.com/course/cW6ILWw7c2eae26101ea8d1cc74661e020c98c403d35605)

* 【Extension Plug-in】
In addition to providing official plug-ins, plug-in systems can also add custom gateway plug-ins.

![](https://camo.githubusercontent.com/1496eb6ea6a594c843d50ec243d6f850a8d41b04/687474703a2f2f646174612e656f6c696e6b65722e636f6d2f636f757273652f527a475058343332303265613261656635386334336664323435663166663065636131323265383830613330366231)

* 【Log Setting】
Detailed request logs and system running logs are provided. Request logs can customize recording fields. Running logs can adjust recording levels according to circumstances: ERROR, INFO, DEBUG, etc.

![](http://data.eolinker.com/course/EHNCLtd8f8bee31f86968ee4dfcd8eeff946fe199195dfc)

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
