/*
Copyright 2021 Flant CJSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vector

import (
	"encoding/base64"
	"strings"

	"github.com/deckhouse/deckhouse/modules/460-log-shipper/hooks/internal/impl"
	"github.com/deckhouse/deckhouse/modules/460-log-shipper/hooks/internal/v1alpha1"
)

type commonDestinationSettings struct {
	Name        string   `json:"-"`
	Type        string   `json:"type"`
	Inputs      []string `json:"inputs,omitempty"`
	Healthcheck struct {
		Enabled bool `json:"enabled"`
	} `json:"healthcheck"`
	Buffer buffer `json:"buffer,omitempty"`
}

type buffer struct {
	Size uint32 `json:"max_size,omitempty"`
	Type string `json:"type,omitempty"`
}

type region struct {
	Region string `json:"region,omitempty"`
}

// AppendInputs append inputs to destination. If input is already exists - skip it (dedup)
func (cs *commonDestinationSettings) AppendInputs(inp []string) {
	if len(cs.Inputs) == 0 {
		cs.Inputs = inp
		return
	}

	m := make(map[string]bool, len(cs.Inputs))
	for _, d := range cs.Inputs {
		m[d] = true
	}

	for _, newinp := range inp {
		if _, ok := m[newinp]; !ok {
			cs.Inputs = append(cs.Inputs, newinp)
		}
	}
}

func (cs *commonDestinationSettings) GetName() string {
	return cs.Name
}

func NewLokiDestination(name string, cspec v1alpha1.ClusterLogDestinationSpec) impl.LogDestination {
	spec := cspec.Loki
	common := commonDestinationSettings{
		Name: "d8_cluster_" + name,
		Type: "loki",
	}

	LokiENC := LokiEncoding{
		Codec: "text",
	}

	common.Buffer = buffer{
		Size: 100 * 1024 * 1024, // 100MiB in bytes for vector persistent queue
		Type: "disk",
	}

	if spec.Auth.Password != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.Auth.Password)
		spec.Auth.Password = string(res)
	}

	if spec.TLS.CAFile != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.CAFile)
		spec.TLS.CAFile = string(res)
	}

	if spec.TLS.CertFile != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.CertFile)
		spec.TLS.CertFile = string(res)
	}

	if spec.TLS.KeyFile != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.KeyFile)
		spec.TLS.KeyFile = string(res)
	}

	if spec.TLS.KeyPass != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.KeyPass)
		spec.TLS.KeyPass = string(res)
	}

	if spec.Auth.Strategy != "" {
		spec.Auth.Strategy = strings.ToLower(spec.Auth.Strategy)
	}

	// default labels
	labels := map[string]string{
		"namespace":  "{{ namespace }}",
		"container":  "{{ container }}",
		"image":      "{{ image }}",
		"pod":        "{{ pod }}",
		"node":       "{{ node_name }}",
		"pod_ip":     "{{ pod_ip }}",
		"stream":     "{{ stream }}",
		"pod_labels": "{{ pod_labels }}",
		"pod_owner":  "{{ pod_owner }}",
	}

	for k, v := range cspec.ExtraLabels {
		labels[k] = v
	}

	return &lokiDestination{
		commonDestinationSettings: common,
		Auth:                      spec.Auth,
		TLS:                       CommonTLS(spec.TLS),
		Labels:                    labels,
		Endpoint:                  spec.Endpoint,
		Encoding:                  LokiENC,
		RemoveLabelFields:         true,
		OutOfOrderAction:          "rewrite_timestamp",
	}
}

func NewElasticsearchDestination(name string, cspec v1alpha1.ClusterLogDestinationSpec) impl.LogDestination {
	spec := cspec.Elasticsearch
	var ESBatch batch

	common := commonDestinationSettings{
		Name: "d8_cluster_" + name,
		Type: "elasticsearch",
	}

	common.Buffer = buffer{
		Size: 100 * 1024 * 1024, // 100MiB in bytes for vector persistent queue
		Type: "disk",
	}

	ESBatch = batch{
		MaxSize:     10 * 1024 * 1024, // 10MiB in bytes for elasticsearch bulk api
		TimeoutSecs: 1,
	}

	if spec.Auth.Password != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.Auth.Password)
		spec.Auth.Password = string(res)
	}

	if spec.Auth.AwsAccessKey != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.Auth.AwsAccessKey)
		spec.Auth.AwsAccessKey = string(res)
	}

	if spec.Auth.AwsSecretKey != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.Auth.AwsSecretKey)
		spec.Auth.AwsSecretKey = string(res)
	}

	if spec.TLS.CAFile != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.CAFile)
		spec.TLS.CAFile = string(res)
	}

	if spec.TLS.CertFile != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.CertFile)
		spec.TLS.CertFile = string(res)
	}

	if spec.TLS.KeyFile != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.KeyFile)
		spec.TLS.KeyFile = string(res)
	}

	if spec.TLS.KeyPass != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.KeyPass)
		spec.TLS.KeyPass = string(res)
	}

	EsEnc := ElasticsearchEncoding{
		TimestampFormat: "rfc3339",
	}

	EsAuth := ElasticsearchAuth{
		AwsAccessKey:  spec.Auth.AwsAccessKey,
		AwsSecretKey:  spec.Auth.AwsSecretKey,
		AwsAssumeRole: spec.Auth.AwsAssumeRole,
		User:          spec.Auth.User,
		Strategy:      spec.Auth.Strategy,
		Password:      spec.Auth.Password,
	}

	if EsAuth.Strategy != "" {
		EsAuth.Strategy = strings.ToLower(EsAuth.Strategy)
	}

	AwsRegion := region{
		Region: spec.Auth.AwsRegion,
	}

	return &elasticsearchDestination{
		commonDestinationSettings: common,
		Auth:                      EsAuth,
		Encoding:                  EsEnc,
		TLS:                       CommonTLS(spec.TLS),
		AWS:                       AwsRegion,
		Batch:                     ESBatch,
		Endpoint:                  spec.Endpoint,
		Compression:               "gzip",
		Index:                     spec.Index,
		BulkAction:                "index",
	}
}

func NewLogstashDestination(name string, cspec v1alpha1.ClusterLogDestinationSpec) impl.LogDestination {
	spec := cspec.Logstash
	var enabledTLS bool

	common := commonDestinationSettings{
		Name: "d8_cluster_" + name,
		Type: "socket",
	}

	common.Buffer = buffer{
		Size: 100 * 1024 * 1024, // 100MiB in bytes for vector persistent queue
		Type: "disk",
	}

	if spec.TLS.CAFile != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.CAFile)
		spec.TLS.CAFile = string(res)
	}

	if spec.TLS.CertFile != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.CertFile)
		spec.TLS.CertFile = string(res)
	}

	if spec.TLS.KeyFile != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.KeyFile)
		spec.TLS.KeyFile = string(res)
	}

	if spec.TLS.KeyPass != "" {
		res, _ := base64.StdEncoding.DecodeString(spec.TLS.KeyPass)
		spec.TLS.KeyPass = string(res)
	}

	if spec.TLS.KeyFile != "" || spec.TLS.CertFile != "" || spec.TLS.CAFile != "" {
		enabledTLS = true
	} else {
		enabledTLS = false
	}
	CTLS := CommonTLS{
		CAFile:         spec.TLS.CAFile,
		CertFile:       spec.TLS.CertFile,
		KeyFile:        spec.TLS.KeyFile,
		KeyPass:        spec.TLS.KeyPass,
		VerifyHostname: spec.TLS.VerifyHostname,
	}
	LSTLS := LogstashTLS{
		CommonTLS:         CTLS,
		VerifyCertificate: spec.TLS.VerifyCertificate,
		Enabled:           enabledTLS,
	}

	LogstashEnc := LogstashEncoding{
		Codec:           "json",
		TimestampFormat: "rfc3339",
	}

	return &logstashDestination{
		commonDestinationSettings: common,
		Encoding:                  LogstashEnc,
		TLS:                       LSTLS,
		Mode:                      "tcp",
		Address:                   spec.Endpoint,
	}
}
