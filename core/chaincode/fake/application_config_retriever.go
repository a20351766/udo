// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/hyperledger/udo/common/channelconfig"
)

type ApplicationConfigRetriever struct {
	GetApplicationConfigStub        func(cid string) (channelconfig.Application, bool)
	getApplicationConfigMutex       sync.RWMutex
	getApplicationConfigArgsForCall []struct {
		cid string
	}
	getApplicationConfigReturns struct {
		result1 channelconfig.Application
		result2 bool
	}
	getApplicationConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Application
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ApplicationConfigRetriever) GetApplicationConfig(cid string) (channelconfig.Application, bool) {
	fake.getApplicationConfigMutex.Lock()
	ret, specificReturn := fake.getApplicationConfigReturnsOnCall[len(fake.getApplicationConfigArgsForCall)]
	fake.getApplicationConfigArgsForCall = append(fake.getApplicationConfigArgsForCall, struct {
		cid string
	}{cid})
	fake.recordInvocation("GetApplicationConfig", []interface{}{cid})
	fake.getApplicationConfigMutex.Unlock()
	if fake.GetApplicationConfigStub != nil {
		return fake.GetApplicationConfigStub(cid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getApplicationConfigReturns.result1, fake.getApplicationConfigReturns.result2
}

func (fake *ApplicationConfigRetriever) GetApplicationConfigCallCount() int {
	fake.getApplicationConfigMutex.RLock()
	defer fake.getApplicationConfigMutex.RUnlock()
	return len(fake.getApplicationConfigArgsForCall)
}

func (fake *ApplicationConfigRetriever) GetApplicationConfigArgsForCall(i int) string {
	fake.getApplicationConfigMutex.RLock()
	defer fake.getApplicationConfigMutex.RUnlock()
	return fake.getApplicationConfigArgsForCall[i].cid
}

func (fake *ApplicationConfigRetriever) GetApplicationConfigReturns(result1 channelconfig.Application, result2 bool) {
	fake.GetApplicationConfigStub = nil
	fake.getApplicationConfigReturns = struct {
		result1 channelconfig.Application
		result2 bool
	}{result1, result2}
}

func (fake *ApplicationConfigRetriever) GetApplicationConfigReturnsOnCall(i int, result1 channelconfig.Application, result2 bool) {
	fake.GetApplicationConfigStub = nil
	if fake.getApplicationConfigReturnsOnCall == nil {
		fake.getApplicationConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Application
			result2 bool
		})
	}
	fake.getApplicationConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Application
		result2 bool
	}{result1, result2}
}

func (fake *ApplicationConfigRetriever) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationConfigMutex.RLock()
	defer fake.getApplicationConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ApplicationConfigRetriever) recordInvocation(key string, args []interface{}) {
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
