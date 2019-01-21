
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import common "github.com/hyperledger/fabric/protos/common"
import mock "github.com/stretchr/testify/mock"

//chainpuller是为chainpuller类型自动生成的模拟类型
type ChainPuller struct {
	mock.Mock
}

//close为给定字段提供模拟函数：
func (_m *ChainPuller) Close() {
	_m.Called()
}

//Heightsbyendpoints为给定字段提供模拟函数：
func (_m *ChainPuller) HeightsByEndpoints() map[string]uint64 {
	ret := _m.Called()

	var r0 map[string]uint64
	if rf, ok := ret.Get(0).(func() map[string]uint64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]uint64)
		}
	}

	return r0
}

//pullblock提供了一个具有给定字段的模拟函数：seq
func (_m *ChainPuller) PullBlock(seq uint64) *common.Block {
	ret := _m.Called(seq)

	var r0 *common.Block
	if rf, ok := ret.Get(0).(func(uint64) *common.Block); ok {
		r0 = rf(seq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Block)
		}
	}

	return r0
}