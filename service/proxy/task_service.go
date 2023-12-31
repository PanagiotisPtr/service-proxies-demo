// Code generated by proxygen. DO NOT EDIT.
package proxy

import (
	proxygenCaster "github.com/panagiotisptr/proxygen/caster"
	proxygenInterceptors "github.com/panagiotisptr/proxygen/interceptor"

	importiserviceTaskService1 "context"
	importiserviceTaskService2 "github.com/panagiotisptr/service-proxies-demo/models"
	importiserviceTaskService0 "github.com/panagiotisptr/service-proxies-demo/service/iservice"
)

type TaskService struct {
	Implementation importiserviceTaskService0.TaskService
	Interceptors   proxygenInterceptors.InterceptorChain
}

var _ importiserviceTaskService0.TaskService = (*TaskService)(nil)

func (this *TaskService) List(
	arg0 importiserviceTaskService1.Context,
) (
	[]*importiserviceTaskService2.Task,
	error,
) {
	rets := this.Interceptors.Apply(
		[]interface{}{
			arg0,
		},
		"List",
		func(args []interface{}) []interface{} {
			res0,
				res1 := this.Implementation.List(
				args[0].(importiserviceTaskService1.Context),
			)

			return []interface{}{
				res0,
				res1,
			}
		},
	)

	return proxygenCaster.Cast[[]*importiserviceTaskService2.Task](rets[0]),
		proxygenCaster.Cast[error](rets[1])
}

func (this *TaskService) Get(
	arg0 importiserviceTaskService1.Context,
	arg1 int64,
) (
	*importiserviceTaskService2.Task,
	error,
) {
	rets := this.Interceptors.Apply(
		[]interface{}{
			arg0,
			arg1,
		},
		"Get",
		func(args []interface{}) []interface{} {
			res0,
				res1 := this.Implementation.Get(
				args[0].(importiserviceTaskService1.Context),
				args[1].(int64),
			)

			return []interface{}{
				res0,
				res1,
			}
		},
	)

	return proxygenCaster.Cast[*importiserviceTaskService2.Task](rets[0]),
		proxygenCaster.Cast[error](rets[1])
}

func (this *TaskService) Create(
	arg0 importiserviceTaskService1.Context,
	arg1 *importiserviceTaskService2.Task,
) error {
	rets := this.Interceptors.Apply(
		[]interface{}{
			arg0,
			arg1,
		},
		"Create",
		func(args []interface{}) []interface{} {
			res0 := this.Implementation.Create(
				args[0].(importiserviceTaskService1.Context),
				args[1].(*importiserviceTaskService2.Task),
			)

			return []interface{}{
				res0,
			}
		},
	)

	return proxygenCaster.Cast[error](rets[0])
}

func (this *TaskService) Update(
	arg0 importiserviceTaskService1.Context,
	arg1 *importiserviceTaskService2.Task,
) error {
	rets := this.Interceptors.Apply(
		[]interface{}{
			arg0,
			arg1,
		},
		"Update",
		func(args []interface{}) []interface{} {
			res0 := this.Implementation.Update(
				args[0].(importiserviceTaskService1.Context),
				args[1].(*importiserviceTaskService2.Task),
			)

			return []interface{}{
				res0,
			}
		},
	)

	return proxygenCaster.Cast[error](rets[0])
}

func (this *TaskService) Delete(
	arg0 importiserviceTaskService1.Context,
	arg1 int64,
) error {
	rets := this.Interceptors.Apply(
		[]interface{}{
			arg0,
			arg1,
		},
		"Delete",
		func(args []interface{}) []interface{} {
			res0 := this.Implementation.Delete(
				args[0].(importiserviceTaskService1.Context),
				args[1].(int64),
			)

			return []interface{}{
				res0,
			}
		},
	)

	return proxygenCaster.Cast[error](rets[0])
}
