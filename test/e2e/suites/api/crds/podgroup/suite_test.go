/*
Copyright 2025 NVIDIA CORPORATION
SPDX-License-Identifier: Apache-2.0
*/
package podgroup

import (
	"testing"

	"github.com/NVIDIA/KAI-scheduler/test/e2e/modules/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPodGroups(t *testing.T) {
	utils.SetLogger()
	RegisterFailHandler(Fail)
	RunSpecs(t, "PodGroup Suite")
}
