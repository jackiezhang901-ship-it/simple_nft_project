# Simple NFT Project

## 项目简介

Simple NFT Project是一个基于区块链技术的NFT交易系统，采用订单簿(OrderBook)模式实现NFT的去中心化交易。项目包含智能合约、后端API服务、数据同步服务和基础工具库，构建了一个完整的NFT交易生态系统。

## 项目背景

- **2020年**：DeFi(去中心化金融)元年，各类去中心化金融应用蓬勃发展
- **2021年**：NFT(非同质化代币)元年，NFT市场呈现爆发式增长

随着数字资产的兴起和去中心化市场的需求增长，NFT交易系统不仅是一个基于区块链的应用，更是链上技术与链下服务高度结合的典型范例。本项目探索如何将区块链的去中心化、透明性、不可篡改等特点与传统的链下业务流程进行有机融合，创建一个灵活、可扩展的系统架构。

## 项目意义

1. **技术架构的通用性和可扩展性**：设计了一套可复用的技术架构，不仅服务于NFT交易市场，还能支持未来其他潜在的链上应用，如Bitcoin上的铭文、符文等新兴数字资产。

2. **链上技术原理与链下服务的结合**：通过智能合约实现链上交易逻辑，同时通过后端服务提供高效的数据查询和业务处理能力。

3. **去中心化应用的场景扩展**：探索NFT交易在不同场景下的应用，为去中心化应用的发展提供实践参考。

## 系统架构

项目采用微服务架构，主要包含以下几个核心组件：

### 1. EasySwapContract - 智能合约层

EasySwapContract是基于Solidity开发的智能合约项目，实现了基于订单簿模型的NFT交易系统。

**核心功能：**
- 订单簿交易逻辑
- 资产托管和管理
- 订单验证和执行

**合约组件：**
- `EasySwapOrderBook`：实现完整的订单簿交易逻辑
  - `OrderStorage`：用于存储订单信息的模块
  - `OrderValidator`：用于处理订单逻辑验证的模块
  - `ProtocolManager`：用于管理协议费的模块
- `EasySwapVault`：独立存储订单相关资产的模块

**支持的交易类型：**
1. 限价卖单 (Limit Sell Order)
2. 限价买单 (Limit Buy Order)
3. 市价卖单 (Market Sell Order)
4. 市价买单 (Market Buy Order)
5. 订单编辑/取消 (Edit/Cancel Order)

### 2. NFTAuctionBackend - 后端API服务

NFTAuctionBackend是基于Go语言和Gin框架开发的后端API服务，提供NFT交易相关的数据查询和业务处理功能。

**核心功能：**
- NFT集合(Collection)管理
- NFT物品(Item)查询
- 订单(Order)管理
- 交易活动(Activity)记录
- 用户资产组合(Portfolio)管理
- 排行榜(Ranking)服务

**主要API接口：**
- `/collections/{address}/items`：获取集合中的物品列表
- `/collections/{address}/bids`：获取集合的出价信息
- `/collections/{address}/{token_id}/bids`：获取集合中特定物品的出价信息
- `/collections/{address}/{token_id}`：获取物品详情
- `/collections/{address}/top-trait`：获取物品特性的最高价格信息
- `/collections/{address}/history-sales`：获取历史销售价格信息
- `/collections/{address}/{token_id}/traits`：获取物品特性信息
- `/activities`：查询多链活动记录
- `/portfolio/bid-orders`：批量查询出价信息

### 3. NFTAuctionSync - 数据同步服务

NFTAuctionSync是基于Go语言开发的数据同步服务，负责将区块链上的EasySwap合约事件同步到数据库中。

**核心功能：**
- 订单簿事件同步
- 拍卖事件同步
- NFT信息同步
- 集合过滤和管理

**同步流程：**
1. 监听区块链上的合约事件
2. 解析事件数据
3. 将数据写入数据库
4. 更新缓存(Redis)

### 4. NFTAuctionBase - 基础工具库

NFTAuctionBase是提供基础功能和工具的Go语言库，被其他服务共享使用。

**核心模块：**
- 链客户端(chainclient)：与区块链交互的客户端
- 订单管理(ordermanager)：订单状态管理和查询
- 存储接口(stores)：数据库和缓存访问接口
- HTTP客户端(xhttp)：HTTP请求封装
- 日志记录(logger)：统一的日志记录接口
- 错误处理(errcode)：错误码和错误处理
- 工具包(kit)：常用工具函数

## NFT数据模型

系统使用以下核心实体来描述NFT交易：

1. **Collection**：NFT集合的实体
   - 包含集合的基本信息，如名称、创建者、合约地址等
   - 记录集合的统计信息，如拥有者数量、NFT总量、地板价等

2. **Item**：代表交易系统中代表NFT的实体
   - 记录每个NFT的基本信息，如Token ID、名称、拥有者等
   - 记录NFT的交易信息，如上架价格、成交价格等

3. **Ownership**：代表NFT的所有权
   - 即Item和Wallet的关联关系
   - 记录NFT的流转历史

4. **Order**：代表出售或购买NFT意愿的实体
   - 记录订单的类型、价格、数量等信息
   - 记录订单的状态，如活跃、已完成、已取消等

5. **Activity**：代表NFT状态下发生的事件
   - 包括mint、transfer、list、buy等事件类型
   - 记录事件的详细信息，如交易双方、价格、时间等

## NFT交易模式

系统支持两种主要的NFT交易模式：

1. **订单簿(OrderBook)模式**：
   - Maker和Taker都是用户
   - 价格由订单确定
   - 支持限价单和市价单

2. **做市商(AMM)模式**：
   - Maker是池子，Taker是用户
   - 价格随池子变化
   - 适用于ERC721-AMM等场景

## 技术栈

### 前端
- (待补充)

### 后端
- Go语言
- Gin框架
- GORM
- Redis
- MySQL

### 智能合约
- Solidity
- Hardhat
- OpenZeppelin

### 区块链
- Ethereum
- Optimism
- Sepolia测试网

## 快速开始

### 环境准备

1. 安装Node.js和npm
2. 安装Go语言环境
3. 安装Docker和Docker Compose
4. 准备MySQL和Redis环境

### 智能合约部署

1. 进入EasySwapContract目录
2. 安装依赖：`npm install`
3. 配置环境变量：复制`.env.example`到`.env`并修改配置
4. 编译合约：`npx hardhat compile`
5. 部署合约：`npx hardhat run --network sepolia scripts/deploy.js`

### 后端服务启动

1. 进入NFTAuctionBackend目录
2. 配置数据库和Redis连接信息
3. 配置合约地址和区块链节点
4. 启动服务：`cd src && go run main.go`

### 数据同步服务启动

1. 进入NFTAuctionSync目录
2. 配置数据库和Redis连接信息
3. 配置合约地址和区块链节点
4. 启动服务：`go run main.go daemon`

## 项目结构

```
simple_nft_project/
├── EasySwapContract/      # 智能合约项目
│   ├── contracts/         # Solidity合约代码
│   ├── scripts/           # 部署和交互脚本
│   └── test/              # 合约测试
├── NFTAuctionBackend/     # 后端API服务
│   ├── src/
│   │   ├── api/          # API接口
│   │   ├── service/      # 业务逻辑
│   │   └── main.go       # 程序入口
│   └── config/           # 配置文件
├── NFTAuctionSync/       # 数据同步服务
│   ├── service/          # 同步逻辑
│   ├── model/            # 数据模型
│   └── config/           # 配置文件
└── NFTAuctionBase/       # 基础工具库
    ├── chain/            # 区块链相关
    ├── ordermanager/     # 订单管理
    └── stores/           # 存储接口
```

## 数据库设计

系统使用MySQL存储业务数据，主要包含以下表：

1. `ob_collection_sepolia`：NFT集合信息表
2. `ob_item_sepolia`：NFT物品信息表
3. `ob_order_sepolia`：订单信息表
4. `ob_activity_sepolia`：交易活动记录表

详细的表结构请参考EasySwapContract/README.md中的SQL定义。

## 开发指南

### 智能合约开发

1. 编写Solidity合约代码
2. 编写单元测试
3. 使用Hardhat进行本地测试
4. 部署到测试网验证

### 后端开发

1. 实现业务逻辑
2. 编写API接口
3. 编写单元测试
4. 使用Swagger生成API文档

### 数据同步开发

1. 实现事件监听逻辑
2. 实现数据解析和存储逻辑
3. 编写单元测试
4. 监控同步状态和性能

## 贡献指南

欢迎贡献代码、报告问题或提出建议。请遵循以下步骤：

1. Fork本仓库
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建Pull Request

## 许可证

本项目采用Apache 2.0许可证。详见LICENSE文件。

## 联系方式

- 项目主页：[GitHub](https://github.com/jackiezhang901-ship-it/simple_nft_project)
- 问题反馈：[Issues](https://github.com/jackiezhang901-ship-it/simple_nft_project/issues)
- 邮箱：support@nftauction.com

## 致谢

感谢所有为本项目做出贡献的开发者和社区成员。
