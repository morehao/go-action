# 服务器指纹采集服务

一个基于 Go + Gin 框架的服务器指纹采集服务，支持物理服务器、Docker 容器和 Kubernetes 环境部署。


## 核心功能
- 全面的指纹信息采集：

系统信息：主机名、OS、架构、内核版本
硬件信息：CPU、内存、磁盘、网卡 MAC
唯一标识：机器 ID、主机 ID
环境信息：自动检测物理服务器/Docker/K8s


- 智能环境检测：

自动识别运行环境
容器 ID 提取
K8s Pod/Node 信息获取


- RESTful API：

/health - 健康检查
/fingerprint - 获取完整指纹信息

## 服务部署
### 本地运行
``` bash
make deps && make run
```
### Docker 运行
``` bash
make deploy-docker
```
### Kubernetes 运行
``` bash
make deploy-k8s
```