## easynode_collect

easynode是对区块链节点进一步抽象和简化，封装了普通区块链节点的技术难点，是面向业务方的节点。使业务方即使不了解区块链概念和技术的情况下，基于现有技术栈和经验，仍然可以设计自己基于区块链的业务。具有一下特征：
- 支持多种公链：目前已支持 ether、tron等公链
- 功能可配制化：通过配置文件可以开启或关闭部分功能
- 动态横向扩展：在不重启系统情况下，可以启动新节点，分担当前系统的压力
- 冗余备份：可以配置多个区块链节点，提供系统稳定性，防止部分节点的异常导致业务系统异常
- 自带负载均衡：在不依赖第三方组件的情况下，通过2层均衡设计，解决了单节点性能瓶颈问题
- 过程可视化：整个系统执行过程都是可视化的，方便监控和维护
- 同步数据方式多样：除了通过配置文件指定自动同步数据外，用户还可以通过HTTP协议自定义需要同步的数据
- 系统健壮性：系统自带异常重试、错误检查等功能，解决因网络异常等客观原因导致数据丢失的问题

easynode_collect是easynode系统的基础和核心服务，是其他服务的运行的必要条件。该服务负责同步主网区块数据的服务，根据用户配置的规则，自动同步主网数据到本地，easynode_collect服务支持横向扩展和负载均衡。


## Prerequisites

- go version: >=1.16
- clickhouse 部署和配置
- kafka 部署和配置
- mysql 部署和配置

## Building the source
(以linux系统为例)
- mkdir easynode & cd easynode
- git clone https://github.com/uduncloud/easynode_collect.git
- cd easynode_collect
- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o easynode_collect app.go
(mac下编译linux程序为例，其他交叉编译的命令请自行搜索)

- ./easynode_collect -config ./config.json

## usages
