package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

//在上述代码中，我们针对 LimiterIface 接口实现了我们的 MethodLimiter 限流器，
//主要逻辑是在 Key 方法中根据 RequestURI 切割出核心路由作为键值对名称，
//并在 GetBucket 和 AddBuckets 进行获取和设置 Bucket 的对应逻辑。
type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	return MethodLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}

	return l
}
