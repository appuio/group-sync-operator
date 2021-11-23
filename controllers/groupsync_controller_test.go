/*


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

package controllers

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	userv1 "github.com/openshift/api/user/v1"
	redhatcopv1alpha1 "github.com/redhat-cop/group-sync-operator/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var _ = Describe("GroupSyncController controller", func() {
	const (
		preexistingGroup = "preexisting-group"
		namespace        = "default"
	)

	BeforeEach(func() {
		ctx := context.Background()

		By("By having pre-existing openshift groups")
		ocpGroup := &userv1.Group{
			ObjectMeta: metav1.ObjectMeta{
				Name: preexistingGroup,
			},
		}
		Expect(k8sClient.Create(ctx, ocpGroup)).Should(Succeed())

	})

	AfterEach(func() {
		ctx := context.Background()

		k8sClient.DeleteAllOf(ctx, &redhatcopv1alpha1.GroupSync{}, client.InNamespace(namespace))
		k8sClient.DeleteAllOf(ctx, &userv1.User{})
		k8sClient.DeleteAllOf(ctx, &userv1.Group{})
	})

	It("It should sync groups from the provider", func() {
		ctx := context.Background()

		By("By creating a sync config")
		syncConfig := &redhatcopv1alpha1.GroupSync{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "sync-groups",
				Namespace: namespace,
			},
			Spec: redhatcopv1alpha1.GroupSyncSpec{
				DeleteDisappearedGroups: true,
				Providers: []redhatcopv1alpha1.Provider{
					{
						Name: "sync-groups",
						ProviderType: &redhatcopv1alpha1.ProviderType{
							Static: &redhatcopv1alpha1.StaticProvider{
								Groups: []redhatcopv1alpha1.StaticProviderGroup{
									{Name: "staff", Users: []string{"xerxes.xoe", "mermes.moe"}},
									{Name: "developers", Users: []string{"david.devel"}},
								},
							},
						},
					},
				},
			},
		}
		Expect(k8sClient.Create(ctx, syncConfig)).Should(Succeed())
		By("By checking for synced groups in cluster")
		Eventually(assertGroup(ctx, "staff"), "10s", "250ms").Should(Succeed())
		Eventually(assertGroup(ctx, "developers"), "10s", "250ms").Should(Succeed())

		By("By deleting group")
		syncConfig = &redhatcopv1alpha1.GroupSync{}
		Expect(k8sClient.Get(ctx, types.NamespacedName{Namespace: namespace, Name: "sync-groups"}, syncConfig)).Should(Succeed())
		syncConfig.Spec.Providers[0].Static.Groups = syncConfig.Spec.Providers[0].Static.Groups[0:1]
		Expect(k8sClient.Update(ctx, syncConfig)).Should(Succeed())

		By("By checking for deleted group in cluster")
		Eventually(assertGroup(ctx, "staff"), "10s", "250ms").Should(Succeed())
		Eventually(assertGroupNotFound(ctx, "developers"), "10s", "250ms").Should(Succeed())
	})
})

func assertGroup(ctx context.Context, name string) func() error {
	return func() error {
		ocpGroup := &userv1.Group{}
		err := k8sClient.Get(ctx, types.NamespacedName{Namespace: "", Name: name}, ocpGroup)
		return err
	}
}

func assertGroupNotFound(ctx context.Context, name string) func() error {
	return func() error {
		ocpGroup := &userv1.Group{}
		err := k8sClient.Get(ctx, types.NamespacedName{Namespace: "", Name: name}, ocpGroup)
		if err == nil {
			return fmt.Errorf("group '%s' exists", name)
		}
		if apierrors.IsNotFound(err) {
			return nil
		}
		return err
	}
}
