/*
Copyright 2021 The cert-manager Authors.

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

package util

import (
	"sort"
	"strings"

	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"

	csiapi "github.com/cert-manager/csi-driver/pkg/apis/v1alpha1"
)

func CertificateRequestReady(cr *cmapi.CertificateRequest) bool {
	readyType := cmapi.CertificateRequestConditionReady
	readyStatus := cmmeta.ConditionTrue

	existingConditions := cr.Status.Conditions
	for _, cond := range existingConditions {
		if readyType == cond.Type && readyStatus == cond.Status {
			return true
		}
	}

	return false
}

func KeyUsagesFromAttributes(attr map[string]string) []cmapi.KeyUsage {
	usageCSV := attr[csiapi.KeyUsagesKey]

	if len(usageCSV) == 0 {
		return nil
	}

	var keyUsages []cmapi.KeyUsage
	for _, usage := range strings.Split(usageCSV, ",") {
		keyUsages = append(keyUsages, cmapi.KeyUsage(strings.TrimSpace(usage)))
	}

	return keyUsages
}

func keyUsagesMatch(a, b []cmapi.KeyUsage) bool {
	if len(a) != len(b) {
		return false
	}

	aa, bb := make([]cmapi.KeyUsage, len(a)), make([]cmapi.KeyUsage, len(b))
	copy(aa, a)
	copy(bb, b)

	sort.SliceStable(aa, func(i, j int) bool {
		return aa[i] < aa[j]
	})
	sort.SliceStable(bb, func(i, j int) bool {
		return bb[i] < bb[j]
	})

	for i, s := range aa {
		if s != bb[i] {
			return false
		}
	}

	return true
}
