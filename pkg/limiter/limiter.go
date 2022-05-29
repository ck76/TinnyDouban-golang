package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

//在上述代码中，我们声明了 LimiterIface 接口，用于定义当前限流器所必须要的方法。
//
//为什么要这么做呢，实际上需要知道一点，限流器是存在多种实现的，
//可能某一类接口需要限流器 A，另外一类接口需要限流器 B，
//所采用的策略不是完全一致的，因此我们需要声明 LimiterIfac 这类通用接口，
//保证其接口的设计，我们初步的在 Iface 接口中，一共声明了三个方法，如下：
//
//Key：获取对应的限流器的键值对名称。
//GetBucket：获取令牌桶。
//AddBuckets：新增多个令牌桶。
//同时我们定义 Limiter 结构体用于存储令牌桶与键值对名称的映射关系，
//并定义 LimiterBucketRule 结构体用于存储令牌桶的一些相应规则属性，如下：
//
//Key：自定义键值对名称。
//FillInterval：间隔多久时间放 N 个令牌。
//Capacity：令牌桶的容量。
//Quantum：每次到达间隔时间后所放的具体令牌数量。
//至此我们就完成了一个 Limter 最基本的属性定义了，
//接下来我们将针对不同的情况实现我们这个项目中的限流器。

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
