/*******************************************************************************
*
* Copyright 2019 SAP SE
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You should have received a copy of the License along with this
* program. If not, you may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
*******************************************************************************/

package disco

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/pkg/errors"
	"github.com/sapcc/kubernetes-operators/disco/pkg/config"
)

func newAuthenticatedProviderClient(ao config.AuthOpts) (provider *gophercloud.ProviderClient, err error) {
	defer func() {
		if err != nil {
			ao.Password = "<password hidden>"
		}
	}()

	opts := &tokens.AuthOptions{
		IdentityEndpoint: ao.AuthURL,
		Username:         ao.Username,
		Password:         ao.Password,
		DomainName:       ao.UserDomainName,
		AllowReauth:      true,
		Scope: tokens.Scope{
			ProjectName: ao.ProjectName,
			DomainName:  ao.ProjectDomainName,
		},
	}

	provider, err = openstack.NewClient(ao.AuthURL)
	if err != nil {
		return nil, errors.Wrapf(err, "could not initialize openstack client")
	}
	provider.UseTokenLock()

	err = openstack.AuthenticateV3(provider, opts, gophercloud.EndpointOpts{})
	if err != nil {
		return nil, errors.Wrapf(err, "openstack authentication failed")
	}

	if provider.TokenID == "" {
		return nil, errors.New("token is empty. openstack authentication failed")
	}
	return provider, err
}

// NewOpenStackDesignateClient returns an authenticated Designate v2 ServiceClient.
func NewOpenStackDesignateClient(ao config.AuthOpts) (*gophercloud.ServiceClient, error) {
	provider, err := newAuthenticatedProviderClient(ao)
	if err != nil {
		return nil, err
	}

	return openstack.NewDNSV2(
		provider,
		gophercloud.EndpointOpts{Region: ao.RegionName, Availability: gophercloud.AvailabilityPublic},
	)
}
