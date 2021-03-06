package lifecycle

import (
	"github.com/RSE-Cambridge/data-acc/internal/pkg/mocks"
	"github.com/RSE-Cambridge/data-acc/internal/pkg/registry"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVolumeLifecycleManager_Mount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockVolReg := mocks.NewMockVolumeRegistry(mockCtrl)

	volume := registry.Volume{Name: "vol1", SizeBricks: 3}
	vlm := NewVolumeLifecycleManager(mockVolReg, nil, volume)
	hosts := []string{"host1", "host2"}

	mockVolReg.EXPECT().UpdateVolumeAttachments(volume.Name, map[string]registry.Attachment{
		"host1": {Hostname: "host1", State: registry.RequestAttach},
		"host2": {Hostname: "host2", State: registry.RequestAttach},
	})
	mockVolReg.EXPECT().WaitForCondition(volume.Name, gomock.Any())

	err := vlm.Mount(hosts)
	assert.Nil(t, err)
}
