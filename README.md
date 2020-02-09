![](https://data.eolinker.com/WJ2lfq217421a961efc420d88a7cb6f59586824a8ea2f84.jpg)

[![Gitter](https://badges.gitter.im/goku-api-gateway/community.svg)](https://gitter.im/goku-api-gateway/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/eolinker/goku-api-gateway)](https://goreportcard.com/report/github.com/eolinker/goku-api-gateway) [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3214/badge)](https://bestpractices.coreinfrastructure.org/projects/3214) ![](https://img.shields.io/badge/license-GPL3.0-blue.svg)

Goku API Gateway is a Golang-based microservice gateway that enables high-performance dynamic routing,service orchestration, multi-tenancy management, API access control, etc. It's also suitable for API management under micro-service system. 

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

[![Stargazers over time](https://starchart.cc/eolinker/goku-api-gateway.svg)](#)

# Product Features
- **Dashboard**: Built-in dashboard to configure Goku.
- **Cluster Management**：Goku nodes are stateless and can be expanded horizontally. Also the configuration can be synchronized automatically.
- **Hot Updates**: Continuously updates configurations without restart nodes.
- **Orchestration**：Orchestration can correspond to multiple backends. The backend input parameter supports the client incoming, and also supports the parameter transfer between backend. The return data of backend supports filter, delete, move, rename, target and group. API can set the exception return when the orchestration call fails.
- **Data transform ** :Support for converting returned data to JSON or XML.
- **Load balancing**: Round-robin load balancing with weight.
- **Service Discovery**: Service discorvery from Consul or Eureka.
- **HTTP(S) Forward Proxy**: Hide real backend services, support Rest API, Webservice.
- **Multi-tenant management**: According to different strategies to regnorize different users.
- **Strategies**: Support different strategies to access different APIs, configure different authentication (anonymous, Apikey, Basic) and so on.
- **API Alert**: Support the webhook and email to alert abnormal services.
- **Flexible transmit rules**: support fuzzy matching request path, support rewriting transmit path, etc.
- **IP Whitelist/Blacklist**
- **Custom plugins**: Allow plugins to be mounted in common phases, such as before match, access, and proxy.
- **CLI**: Start\stop\reload Goku through the command line.
- **Serverless**: Invoke functions in each phase in Goku.
- **Access Log**:Only record the basic content in proxy, customize the record fields and sort order, and automatically clean up the logs periodically.
- **System Log**:Provide running logs of consoles and nodes,only record the error information, adjust the level to INFO, WARN or DEBUG according to the actual situation.
- **Scalability**: plug-in mechanism is easy to extend.
- **High performance**: Performance excels among many gateways.
- **Open API**：Provide OPEN API for users to operate on the gateway for easy integration.
- **Configured version management** : Support for the release of operations and multiple rollbacks.
- **Monitoring and indicators**: Support for Prometheus, Graphite.

# Benchmark
![](https://data.eolinker.com/p7NFG6lb4c73b26cc880e838fe45aa31bc037b7415e3770.jpg)
[Benchmark Detail](https://help.eolinker.com/#/tutorial/?groupID=c-362&productID=19#tip7 "Benchmark Detail")

# Console Preview
[Console Preview Detail](https://github.com/eolinker/goku-api-gateway/blob/master/docs/CONSOLE_PREVIEW.md "See Console Preview")

# Quick Start
* [Deployment Tutorial](https://help.eolinker.com/#/tutorial/?groupID=c-371&productID=19 "Deployment Tutorial")
* [Docker for Console](https://hub.docker.com/r/eolinker/goku-api-gateway-ce-console "Docker for Console")、[Docker for Gateway Node](https://hub.docker.com/r/eolinker/goku-api-gateway-ce-node "Docker for Gateway Node")
* [Quick Start](https://help.eolinker.com/#/tutorial/?groupID=c-307&productID=19 "Quick Start Tutorial")
* [Source Code Compilation](https://help.eolinker.com/#/tutorial/?groupID=c-350&productID=19 "Source Code Compilation")

# Enterprise Support
Goku API Gateway EE (Enterprise Version) has more powerful functions, plug-in libraries and professional technical support services. If you want to know more details, you can contact us in the following ways.
- Apply for free trial and demonstration of Enterprise Version：[Appointment trial](https://www.eolinker.com/#/survey/applyAmsCloud "Appointment trial")
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
