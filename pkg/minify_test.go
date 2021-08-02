package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinifySimpleStructure(t *testing.T) {
	input := `nodeA:
  nodeB: value1
nodeC: value2`
	expected := `{nodeA:{nodeB:value1},nodeC:value2}`

	got, err := Minify([]byte(input))
	assert.NoError(t, err)
	assert.Equal(t, expected, string(got))
}

func TestMinifyComplexStructure(t *testing.T) {
	input := `spec:
	description: Depending artifacts
	artifacts:
	  - name: artifact-a
		kustomize:
		  path: crds/rbacs
	  - name: artifact-b
		chart:
		  url: https://charts.bitnami.com/bitnami
		  name: nginx`
	expected := `{spec:{description:Depending artifacts,artifacts:[{name:artifact-a,kustomize:{path: crds/rbacs}},{name:artifact-b,chart:{url:'https://charts.bitnami.com/bitnami',name:nginx}}]}}`

	got, err := Minify([]byte(input))
	assert.NoError(t, err)
	assert.Equal(t, expected, string(got))
}
