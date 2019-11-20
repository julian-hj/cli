// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/types"
)

type FakeSetLabelActor struct {
	UpdateApplicationLabelsByApplicationNameStub        func(string, string, map[string]types.NullString) (v7action.Warnings, error)
	updateApplicationLabelsByApplicationNameMutex       sync.RWMutex
	updateApplicationLabelsByApplicationNameArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}
	updateApplicationLabelsByApplicationNameReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateApplicationLabelsByApplicationNameReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	UpdateBuildpackLabelsByBuildpackNameAndStackStub        func(string, string, map[string]types.NullString) (v7action.Warnings, error)
	updateBuildpackLabelsByBuildpackNameAndStackMutex       sync.RWMutex
	updateBuildpackLabelsByBuildpackNameAndStackArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}
	updateBuildpackLabelsByBuildpackNameAndStackReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateBuildpackLabelsByBuildpackNameAndStackReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	UpdateDomainLabelsByDomainNameStub        func(string, map[string]types.NullString) (v7action.Warnings, error)
	updateDomainLabelsByDomainNameMutex       sync.RWMutex
	updateDomainLabelsByDomainNameArgsForCall []struct {
		arg1 string
		arg2 map[string]types.NullString
	}
	updateDomainLabelsByDomainNameReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateDomainLabelsByDomainNameReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	UpdateOrganizationLabelsByOrganizationNameStub        func(string, map[string]types.NullString) (v7action.Warnings, error)
	updateOrganizationLabelsByOrganizationNameMutex       sync.RWMutex
	updateOrganizationLabelsByOrganizationNameArgsForCall []struct {
		arg1 string
		arg2 map[string]types.NullString
	}
	updateOrganizationLabelsByOrganizationNameReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateOrganizationLabelsByOrganizationNameReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	UpdateRouteLabelsStub        func(string, string, map[string]types.NullString) (v7action.Warnings, error)
	updateRouteLabelsMutex       sync.RWMutex
	updateRouteLabelsArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}
	updateRouteLabelsReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateRouteLabelsReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	UpdateSpaceLabelsBySpaceNameStub        func(string, string, map[string]types.NullString) (v7action.Warnings, error)
	updateSpaceLabelsBySpaceNameMutex       sync.RWMutex
	updateSpaceLabelsBySpaceNameArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}
	updateSpaceLabelsBySpaceNameReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateSpaceLabelsBySpaceNameReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	UpdateStackLabelsByStackNameStub        func(string, map[string]types.NullString) (v7action.Warnings, error)
	updateStackLabelsByStackNameMutex       sync.RWMutex
	updateStackLabelsByStackNameArgsForCall []struct {
		arg1 string
		arg2 map[string]types.NullString
	}
	updateStackLabelsByStackNameReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateStackLabelsByStackNameReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSetLabelActor) UpdateApplicationLabelsByApplicationName(arg1 string, arg2 string, arg3 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateApplicationLabelsByApplicationNameMutex.Lock()
	ret, specificReturn := fake.updateApplicationLabelsByApplicationNameReturnsOnCall[len(fake.updateApplicationLabelsByApplicationNameArgsForCall)]
	fake.updateApplicationLabelsByApplicationNameArgsForCall = append(fake.updateApplicationLabelsByApplicationNameArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateApplicationLabelsByApplicationName", []interface{}{arg1, arg2, arg3})
	fake.updateApplicationLabelsByApplicationNameMutex.Unlock()
	if fake.UpdateApplicationLabelsByApplicationNameStub != nil {
		return fake.UpdateApplicationLabelsByApplicationNameStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateApplicationLabelsByApplicationNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetLabelActor) UpdateApplicationLabelsByApplicationNameCallCount() int {
	fake.updateApplicationLabelsByApplicationNameMutex.RLock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.RUnlock()
	return len(fake.updateApplicationLabelsByApplicationNameArgsForCall)
}

func (fake *FakeSetLabelActor) UpdateApplicationLabelsByApplicationNameCalls(stub func(string, string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateApplicationLabelsByApplicationNameMutex.Lock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.Unlock()
	fake.UpdateApplicationLabelsByApplicationNameStub = stub
}

func (fake *FakeSetLabelActor) UpdateApplicationLabelsByApplicationNameArgsForCall(i int) (string, string, map[string]types.NullString) {
	fake.updateApplicationLabelsByApplicationNameMutex.RLock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.RUnlock()
	argsForCall := fake.updateApplicationLabelsByApplicationNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSetLabelActor) UpdateApplicationLabelsByApplicationNameReturns(result1 v7action.Warnings, result2 error) {
	fake.updateApplicationLabelsByApplicationNameMutex.Lock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.Unlock()
	fake.UpdateApplicationLabelsByApplicationNameStub = nil
	fake.updateApplicationLabelsByApplicationNameReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateApplicationLabelsByApplicationNameReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateApplicationLabelsByApplicationNameMutex.Lock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.Unlock()
	fake.UpdateApplicationLabelsByApplicationNameStub = nil
	if fake.updateApplicationLabelsByApplicationNameReturnsOnCall == nil {
		fake.updateApplicationLabelsByApplicationNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateApplicationLabelsByApplicationNameReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateBuildpackLabelsByBuildpackNameAndStack(arg1 string, arg2 string, arg3 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.Lock()
	ret, specificReturn := fake.updateBuildpackLabelsByBuildpackNameAndStackReturnsOnCall[len(fake.updateBuildpackLabelsByBuildpackNameAndStackArgsForCall)]
	fake.updateBuildpackLabelsByBuildpackNameAndStackArgsForCall = append(fake.updateBuildpackLabelsByBuildpackNameAndStackArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateBuildpackLabelsByBuildpackNameAndStack", []interface{}{arg1, arg2, arg3})
	fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.Unlock()
	if fake.UpdateBuildpackLabelsByBuildpackNameAndStackStub != nil {
		return fake.UpdateBuildpackLabelsByBuildpackNameAndStackStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateBuildpackLabelsByBuildpackNameAndStackReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetLabelActor) UpdateBuildpackLabelsByBuildpackNameAndStackCallCount() int {
	fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.RLock()
	defer fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.RUnlock()
	return len(fake.updateBuildpackLabelsByBuildpackNameAndStackArgsForCall)
}

func (fake *FakeSetLabelActor) UpdateBuildpackLabelsByBuildpackNameAndStackCalls(stub func(string, string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.Lock()
	defer fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.Unlock()
	fake.UpdateBuildpackLabelsByBuildpackNameAndStackStub = stub
}

func (fake *FakeSetLabelActor) UpdateBuildpackLabelsByBuildpackNameAndStackArgsForCall(i int) (string, string, map[string]types.NullString) {
	fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.RLock()
	defer fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.RUnlock()
	argsForCall := fake.updateBuildpackLabelsByBuildpackNameAndStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSetLabelActor) UpdateBuildpackLabelsByBuildpackNameAndStackReturns(result1 v7action.Warnings, result2 error) {
	fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.Lock()
	defer fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.Unlock()
	fake.UpdateBuildpackLabelsByBuildpackNameAndStackStub = nil
	fake.updateBuildpackLabelsByBuildpackNameAndStackReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateBuildpackLabelsByBuildpackNameAndStackReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.Lock()
	defer fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.Unlock()
	fake.UpdateBuildpackLabelsByBuildpackNameAndStackStub = nil
	if fake.updateBuildpackLabelsByBuildpackNameAndStackReturnsOnCall == nil {
		fake.updateBuildpackLabelsByBuildpackNameAndStackReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateBuildpackLabelsByBuildpackNameAndStackReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateDomainLabelsByDomainName(arg1 string, arg2 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateDomainLabelsByDomainNameMutex.Lock()
	ret, specificReturn := fake.updateDomainLabelsByDomainNameReturnsOnCall[len(fake.updateDomainLabelsByDomainNameArgsForCall)]
	fake.updateDomainLabelsByDomainNameArgsForCall = append(fake.updateDomainLabelsByDomainNameArgsForCall, struct {
		arg1 string
		arg2 map[string]types.NullString
	}{arg1, arg2})
	fake.recordInvocation("UpdateDomainLabelsByDomainName", []interface{}{arg1, arg2})
	fake.updateDomainLabelsByDomainNameMutex.Unlock()
	if fake.UpdateDomainLabelsByDomainNameStub != nil {
		return fake.UpdateDomainLabelsByDomainNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateDomainLabelsByDomainNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetLabelActor) UpdateDomainLabelsByDomainNameCallCount() int {
	fake.updateDomainLabelsByDomainNameMutex.RLock()
	defer fake.updateDomainLabelsByDomainNameMutex.RUnlock()
	return len(fake.updateDomainLabelsByDomainNameArgsForCall)
}

func (fake *FakeSetLabelActor) UpdateDomainLabelsByDomainNameCalls(stub func(string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateDomainLabelsByDomainNameMutex.Lock()
	defer fake.updateDomainLabelsByDomainNameMutex.Unlock()
	fake.UpdateDomainLabelsByDomainNameStub = stub
}

func (fake *FakeSetLabelActor) UpdateDomainLabelsByDomainNameArgsForCall(i int) (string, map[string]types.NullString) {
	fake.updateDomainLabelsByDomainNameMutex.RLock()
	defer fake.updateDomainLabelsByDomainNameMutex.RUnlock()
	argsForCall := fake.updateDomainLabelsByDomainNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSetLabelActor) UpdateDomainLabelsByDomainNameReturns(result1 v7action.Warnings, result2 error) {
	fake.updateDomainLabelsByDomainNameMutex.Lock()
	defer fake.updateDomainLabelsByDomainNameMutex.Unlock()
	fake.UpdateDomainLabelsByDomainNameStub = nil
	fake.updateDomainLabelsByDomainNameReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateDomainLabelsByDomainNameReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateDomainLabelsByDomainNameMutex.Lock()
	defer fake.updateDomainLabelsByDomainNameMutex.Unlock()
	fake.UpdateDomainLabelsByDomainNameStub = nil
	if fake.updateDomainLabelsByDomainNameReturnsOnCall == nil {
		fake.updateDomainLabelsByDomainNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateDomainLabelsByDomainNameReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateOrganizationLabelsByOrganizationName(arg1 string, arg2 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.Lock()
	ret, specificReturn := fake.updateOrganizationLabelsByOrganizationNameReturnsOnCall[len(fake.updateOrganizationLabelsByOrganizationNameArgsForCall)]
	fake.updateOrganizationLabelsByOrganizationNameArgsForCall = append(fake.updateOrganizationLabelsByOrganizationNameArgsForCall, struct {
		arg1 string
		arg2 map[string]types.NullString
	}{arg1, arg2})
	fake.recordInvocation("UpdateOrganizationLabelsByOrganizationName", []interface{}{arg1, arg2})
	fake.updateOrganizationLabelsByOrganizationNameMutex.Unlock()
	if fake.UpdateOrganizationLabelsByOrganizationNameStub != nil {
		return fake.UpdateOrganizationLabelsByOrganizationNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateOrganizationLabelsByOrganizationNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetLabelActor) UpdateOrganizationLabelsByOrganizationNameCallCount() int {
	fake.updateOrganizationLabelsByOrganizationNameMutex.RLock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.RUnlock()
	return len(fake.updateOrganizationLabelsByOrganizationNameArgsForCall)
}

func (fake *FakeSetLabelActor) UpdateOrganizationLabelsByOrganizationNameCalls(stub func(string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.Lock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.Unlock()
	fake.UpdateOrganizationLabelsByOrganizationNameStub = stub
}

func (fake *FakeSetLabelActor) UpdateOrganizationLabelsByOrganizationNameArgsForCall(i int) (string, map[string]types.NullString) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.RLock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.RUnlock()
	argsForCall := fake.updateOrganizationLabelsByOrganizationNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSetLabelActor) UpdateOrganizationLabelsByOrganizationNameReturns(result1 v7action.Warnings, result2 error) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.Lock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.Unlock()
	fake.UpdateOrganizationLabelsByOrganizationNameStub = nil
	fake.updateOrganizationLabelsByOrganizationNameReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateOrganizationLabelsByOrganizationNameReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.Lock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.Unlock()
	fake.UpdateOrganizationLabelsByOrganizationNameStub = nil
	if fake.updateOrganizationLabelsByOrganizationNameReturnsOnCall == nil {
		fake.updateOrganizationLabelsByOrganizationNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateOrganizationLabelsByOrganizationNameReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateRouteLabels(arg1 string, arg2 string, arg3 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateRouteLabelsMutex.Lock()
	ret, specificReturn := fake.updateRouteLabelsReturnsOnCall[len(fake.updateRouteLabelsArgsForCall)]
	fake.updateRouteLabelsArgsForCall = append(fake.updateRouteLabelsArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateRouteLabels", []interface{}{arg1, arg2, arg3})
	fake.updateRouteLabelsMutex.Unlock()
	if fake.UpdateRouteLabelsStub != nil {
		return fake.UpdateRouteLabelsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateRouteLabelsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetLabelActor) UpdateRouteLabelsCallCount() int {
	fake.updateRouteLabelsMutex.RLock()
	defer fake.updateRouteLabelsMutex.RUnlock()
	return len(fake.updateRouteLabelsArgsForCall)
}

func (fake *FakeSetLabelActor) UpdateRouteLabelsCalls(stub func(string, string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateRouteLabelsMutex.Lock()
	defer fake.updateRouteLabelsMutex.Unlock()
	fake.UpdateRouteLabelsStub = stub
}

func (fake *FakeSetLabelActor) UpdateRouteLabelsArgsForCall(i int) (string, string, map[string]types.NullString) {
	fake.updateRouteLabelsMutex.RLock()
	defer fake.updateRouteLabelsMutex.RUnlock()
	argsForCall := fake.updateRouteLabelsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSetLabelActor) UpdateRouteLabelsReturns(result1 v7action.Warnings, result2 error) {
	fake.updateRouteLabelsMutex.Lock()
	defer fake.updateRouteLabelsMutex.Unlock()
	fake.UpdateRouteLabelsStub = nil
	fake.updateRouteLabelsReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateRouteLabelsReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateRouteLabelsMutex.Lock()
	defer fake.updateRouteLabelsMutex.Unlock()
	fake.UpdateRouteLabelsStub = nil
	if fake.updateRouteLabelsReturnsOnCall == nil {
		fake.updateRouteLabelsReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateRouteLabelsReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateSpaceLabelsBySpaceName(arg1 string, arg2 string, arg3 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateSpaceLabelsBySpaceNameMutex.Lock()
	ret, specificReturn := fake.updateSpaceLabelsBySpaceNameReturnsOnCall[len(fake.updateSpaceLabelsBySpaceNameArgsForCall)]
	fake.updateSpaceLabelsBySpaceNameArgsForCall = append(fake.updateSpaceLabelsBySpaceNameArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateSpaceLabelsBySpaceName", []interface{}{arg1, arg2, arg3})
	fake.updateSpaceLabelsBySpaceNameMutex.Unlock()
	if fake.UpdateSpaceLabelsBySpaceNameStub != nil {
		return fake.UpdateSpaceLabelsBySpaceNameStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateSpaceLabelsBySpaceNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetLabelActor) UpdateSpaceLabelsBySpaceNameCallCount() int {
	fake.updateSpaceLabelsBySpaceNameMutex.RLock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.RUnlock()
	return len(fake.updateSpaceLabelsBySpaceNameArgsForCall)
}

func (fake *FakeSetLabelActor) UpdateSpaceLabelsBySpaceNameCalls(stub func(string, string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateSpaceLabelsBySpaceNameMutex.Lock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.Unlock()
	fake.UpdateSpaceLabelsBySpaceNameStub = stub
}

func (fake *FakeSetLabelActor) UpdateSpaceLabelsBySpaceNameArgsForCall(i int) (string, string, map[string]types.NullString) {
	fake.updateSpaceLabelsBySpaceNameMutex.RLock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.RUnlock()
	argsForCall := fake.updateSpaceLabelsBySpaceNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSetLabelActor) UpdateSpaceLabelsBySpaceNameReturns(result1 v7action.Warnings, result2 error) {
	fake.updateSpaceLabelsBySpaceNameMutex.Lock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.Unlock()
	fake.UpdateSpaceLabelsBySpaceNameStub = nil
	fake.updateSpaceLabelsBySpaceNameReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateSpaceLabelsBySpaceNameReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateSpaceLabelsBySpaceNameMutex.Lock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.Unlock()
	fake.UpdateSpaceLabelsBySpaceNameStub = nil
	if fake.updateSpaceLabelsBySpaceNameReturnsOnCall == nil {
		fake.updateSpaceLabelsBySpaceNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateSpaceLabelsBySpaceNameReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateStackLabelsByStackName(arg1 string, arg2 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateStackLabelsByStackNameMutex.Lock()
	ret, specificReturn := fake.updateStackLabelsByStackNameReturnsOnCall[len(fake.updateStackLabelsByStackNameArgsForCall)]
	fake.updateStackLabelsByStackNameArgsForCall = append(fake.updateStackLabelsByStackNameArgsForCall, struct {
		arg1 string
		arg2 map[string]types.NullString
	}{arg1, arg2})
	fake.recordInvocation("UpdateStackLabelsByStackName", []interface{}{arg1, arg2})
	fake.updateStackLabelsByStackNameMutex.Unlock()
	if fake.UpdateStackLabelsByStackNameStub != nil {
		return fake.UpdateStackLabelsByStackNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateStackLabelsByStackNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetLabelActor) UpdateStackLabelsByStackNameCallCount() int {
	fake.updateStackLabelsByStackNameMutex.RLock()
	defer fake.updateStackLabelsByStackNameMutex.RUnlock()
	return len(fake.updateStackLabelsByStackNameArgsForCall)
}

func (fake *FakeSetLabelActor) UpdateStackLabelsByStackNameCalls(stub func(string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateStackLabelsByStackNameMutex.Lock()
	defer fake.updateStackLabelsByStackNameMutex.Unlock()
	fake.UpdateStackLabelsByStackNameStub = stub
}

func (fake *FakeSetLabelActor) UpdateStackLabelsByStackNameArgsForCall(i int) (string, map[string]types.NullString) {
	fake.updateStackLabelsByStackNameMutex.RLock()
	defer fake.updateStackLabelsByStackNameMutex.RUnlock()
	argsForCall := fake.updateStackLabelsByStackNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSetLabelActor) UpdateStackLabelsByStackNameReturns(result1 v7action.Warnings, result2 error) {
	fake.updateStackLabelsByStackNameMutex.Lock()
	defer fake.updateStackLabelsByStackNameMutex.Unlock()
	fake.UpdateStackLabelsByStackNameStub = nil
	fake.updateStackLabelsByStackNameReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) UpdateStackLabelsByStackNameReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateStackLabelsByStackNameMutex.Lock()
	defer fake.updateStackLabelsByStackNameMutex.Unlock()
	fake.UpdateStackLabelsByStackNameStub = nil
	if fake.updateStackLabelsByStackNameReturnsOnCall == nil {
		fake.updateStackLabelsByStackNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateStackLabelsByStackNameReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetLabelActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateApplicationLabelsByApplicationNameMutex.RLock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.RUnlock()
	fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.RLock()
	defer fake.updateBuildpackLabelsByBuildpackNameAndStackMutex.RUnlock()
	fake.updateDomainLabelsByDomainNameMutex.RLock()
	defer fake.updateDomainLabelsByDomainNameMutex.RUnlock()
	fake.updateOrganizationLabelsByOrganizationNameMutex.RLock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.RUnlock()
	fake.updateRouteLabelsMutex.RLock()
	defer fake.updateRouteLabelsMutex.RUnlock()
	fake.updateSpaceLabelsBySpaceNameMutex.RLock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.RUnlock()
	fake.updateStackLabelsByStackNameMutex.RLock()
	defer fake.updateStackLabelsByStackNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSetLabelActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.SetLabelActor = new(FakeSetLabelActor)
