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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/sapcc/kubernetes-operators/disco/pkg/apis/disco/v1"
	scheme "github.com/sapcc/kubernetes-operators/disco/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RecordsGetter has a method to return a RecordInterface.
// A group's client should implement this interface.
type RecordsGetter interface {
	Records(namespace string) RecordInterface
}

// RecordInterface has methods to work with Record resources.
type RecordInterface interface {
	Create(*v1.Record) (*v1.Record, error)
	Update(*v1.Record) (*v1.Record, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Record, error)
	List(opts metav1.ListOptions) (*v1.RecordList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Record, err error)
	RecordExpansion
}

// records implements RecordInterface
type records struct {
	client rest.Interface
	ns     string
}

// newRecords returns a Records
func newRecords(c *DiscoV1Client, namespace string) *records {
	return &records{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the record, and returns the corresponding record object, and an error if there is any.
func (c *records) Get(name string, options metav1.GetOptions) (result *v1.Record, err error) {
	result = &v1.Record{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("records").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Records that match those selectors.
func (c *records) List(opts metav1.ListOptions) (result *v1.RecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RecordList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("records").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested records.
func (c *records) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("records").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a record and creates it.  Returns the server's representation of the record, and an error, if there is any.
func (c *records) Create(record *v1.Record) (result *v1.Record, err error) {
	result = &v1.Record{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("records").
		Body(record).
		Do().
		Into(result)
	return
}

// Update takes the representation of a record and updates it. Returns the server's representation of the record, and an error, if there is any.
func (c *records) Update(record *v1.Record) (result *v1.Record, err error) {
	result = &v1.Record{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("records").
		Name(record.Name).
		Body(record).
		Do().
		Into(result)
	return
}

// Delete takes name of the record and deletes it. Returns an error if one occurs.
func (c *records) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("records").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *records) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("records").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched record.
func (c *records) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Record, err error) {
	result = &v1.Record{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("records").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}