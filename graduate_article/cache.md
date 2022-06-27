分布式缓存
memcached: 多线程吞吐高，但原生不支持集群，且数据类型有限；

redis: 集群扩容，数据类型丰富，但单线程能力较 memcached 弱；

方案：memcached+redis

缓存优化：
一致性 hash+虚拟节点 < 有界负载一致性 hash < redis-cluster

解决数据一致性：
数据更新用消息队列到达某个消费者，再同步更新数据库与缓存，保证一致；

保证 Cache Aside 时的读写一致性：
3(写操作的 write)是 set cache，但 4(读操作的 set cache)为 settex(更新缓存时间)

多级缓存：
先清理下游再上游；

下游 TTL 大于上游；

热点缓存
local cache；

单飞；

小技巧
空缓存保护；

pipeline 减少网络往返；

分布式事务
幂等

两阶段提交