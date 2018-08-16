/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */
package gseprocserver

import (
	"context"
	"net/http"

	"configcenter/src/common/metadata"
)

func (p *gseproc) OperateProcess(ctx context.Context, h http.Header, namespace string, data interface{}) (resp *metadata.GseProcRespone, err error) {
	resp = new(metadata.GseProcRespone)
	subPath := "/proc/" + namespace + "/operate"

	err = p.client.Post().
		WithContext(ctx).
		Body(data).
		SubResource(subPath).
		WithHeaders(h).
		Do().
		Into(resp)

	return
}

func (p *gseproc) QueryProcOperateResult(ctx context.Context, h http.Header, namespace, taskid string) (resp *metadata.GseProcessOperateTaskResult, err error) {
	resp = new(metadata.GseProcessOperateTaskResult)
	subPath := "/proc/" + namespace + "/operate/" + taskid

	err = p.client.Get().
		WithContext(ctx).
		Body(nil).
		SubResource(subPath).
		WithHeaders(h).
		Do().
		Into(resp)

	return
}

func (p *gseproc) QueryProcStatus(ctx context.Context, h http.Header, namespace string, data interface{}) (resp *metadata.GseProcRespone, err error) {
	resp = new(metadata.GseProcRespone)
	subPath := "/proc/" + namespace + "/status"

	err = p.client.Post().
		WithContext(ctx).
		Body(data).
		SubResource(subPath).
		WithHeaders(h).
		Do().
		Into(resp)

	return
}

func (p *gseproc) RegisterProcInfo(ctx context.Context, h http.Header, namespace string, data interface{}) (resp *metadata.GseProcRespone, err error) {
	resp = new(metadata.GseProcRespone)
	subPath := "/proc/" + namespace + "/procinfo"

	err = p.client.Post().
		WithContext(ctx).
		Body(data).
		SubResource(subPath).
		WithHeaders(h).
		Do().
		Into(resp)

	return
}

func (p *gseproc) UnRegisterProcInfo(ctx context.Context, h http.Header, namespace string, data interface{}) (resp *metadata.GseProcRespone, err error) {
	resp = new(metadata.GseProcRespone)
	subPath := "/proc/" + namespace + "/procinfo"

	err = p.client.Delete().
		WithContext(ctx).
		Body(data).
		SubResource(subPath).
		WithHeaders(h).
		Do().
		Into(resp)

	return
}