// Copyright 2020 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package steps

import (
	hyperfoilv1alpha2 "github.com/Hyperfoil/hyperfoil-operator/pkg/apis/hyperfoil/v1alpha2"
	"github.com/cucumber/godog"
	"github.com/kiegroup/kogito-operator/test/pkg/framework"
	"github.com/kiegroup/kogito-operator/test/pkg/installers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func registerHyperfoilSteps(ctx *godog.ScenarioContext, data *Data) {
	ctx.Step(`^Hyperfoil Operator is deployed$`, data.hyperfoilOperatorIsDeployed)
	ctx.Step(`^Hyperfoil instance "([^"]*)" is deployed within (\d+) (?:minute|minutes)$`, data.hyperfoilInstanceIsDeployedWithinMinutes)
}

func (data *Data) hyperfoilOperatorIsDeployed() error {
	return installers.GetHyperfoilInstaller().Install(data.Namespace)
}

func (data *Data) hyperfoilInstanceIsDeployedWithinMinutes(name string, timeOutInMin int) error {
	hyperfoil := getHyperfoilDefaultResource(name, data.Namespace)

	framework.GetLogger(data.Namespace).Info("Creating Hyperfoil instance", "name", hyperfoil.Name)
	if err := framework.CreateObject(hyperfoil); err != nil {
		return err
	}

	return framework.WaitForPodsWithLabel(data.Namespace, "role", "controller", 1, 5)
}

func getHyperfoilDefaultResource(name, namespace string) *hyperfoilv1alpha2.Hyperfoil {
	return &hyperfoilv1alpha2.Hyperfoil{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
}
