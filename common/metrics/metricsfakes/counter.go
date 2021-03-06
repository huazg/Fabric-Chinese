
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package metricsfakes

import (
	"sync"

	"github.com/hyperledger/fabric/common/metrics"
)

type Counter struct {
	WithStub        func(labelValues ...string) metrics.Counter
	withMutex       sync.RWMutex
	withArgsForCall []struct {
		labelValues []string
	}
	withReturns struct {
		result1 metrics.Counter
	}
	withReturnsOnCall map[int]struct {
		result1 metrics.Counter
	}
	AddStub        func(delta float64)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		delta float64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Counter) With(labelValues ...string) metrics.Counter {
	fake.withMutex.Lock()
	ret, specificReturn := fake.withReturnsOnCall[len(fake.withArgsForCall)]
	fake.withArgsForCall = append(fake.withArgsForCall, struct {
		labelValues []string
	}{labelValues})
	fake.recordInvocation("With", []interface{}{labelValues})
	fake.withMutex.Unlock()
	if fake.WithStub != nil {
		return fake.WithStub(labelValues...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.withReturns.result1
}

func (fake *Counter) WithCallCount() int {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return len(fake.withArgsForCall)
}

func (fake *Counter) WithArgsForCall(i int) []string {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return fake.withArgsForCall[i].labelValues
}

func (fake *Counter) WithReturns(result1 metrics.Counter) {
	fake.WithStub = nil
	fake.withReturns = struct {
		result1 metrics.Counter
	}{result1}
}

func (fake *Counter) WithReturnsOnCall(i int, result1 metrics.Counter) {
	fake.WithStub = nil
	if fake.withReturnsOnCall == nil {
		fake.withReturnsOnCall = make(map[int]struct {
			result1 metrics.Counter
		})
	}
	fake.withReturnsOnCall[i] = struct {
		result1 metrics.Counter
	}{result1}
}

func (fake *Counter) Add(delta float64) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		delta float64
	}{delta})
	fake.recordInvocation("Add", []interface{}{delta})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(delta)
	}
}

func (fake *Counter) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *Counter) AddArgsForCall(i int) float64 {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].delta
}

func (fake *Counter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Counter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metrics.Counter = new(Counter)
