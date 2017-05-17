package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestXMLJSONMap_Ok(t *testing.T) {
	m := defaultImageSetToJSONMapper{}
	source := []XMLImageSet{
		XMLImageSet{
			ID: "U11603547146784PeC",
			ImageSmall: XMLImage{
				FileRef: "/FT/Graphics/Online/Z_Undefined/2017/03/timeline-artboards-s.png?uuid=4258f26a-13c5-11e7-9469-afea892e4de3",
			},
			ImageMedium: XMLImage{
				FileRef: "/FT/Graphics/Online/Z_Undefined/2017/03/timeline-artboards-m.png?uuid=41614f4c-13c5-11e7-9469-afea892e4de3",
			},
			ImageLarge: XMLImage{
				FileRef: "/FT/Graphics/Online/Z_Undefined/2017/03/timeline-artboards-l.png?uuid=3ff3b7a8-13c5-11e7-9469-afea892e4de3",
			},
		},
		XMLImageSet{
			ID: "U12345547146784RfD",
			ImageSmall: XMLImage{
				FileRef: "/FT/Graphics/Online/Z_Undefined/2017/03/other-s.png?uuid=404cf8d9-1b88-4883-8afe-580e5174830d",
			},
			ImageMedium: XMLImage{
				FileRef: "/FT/Graphics/Online/Z_Undefined/2017/03/other-m.png?uuid=2fe0b459-a23e-452d-a2aa-2e0503982ed2",
			},
			ImageLarge: XMLImage{
				FileRef: "/FT/Graphics/Online/Z_Undefined/2017/03/other-l.png?uuid=2acf1caa-8014-48ec-b070-a0ffbc45d1d5",
			},
		},
	}
	actualImageSets, err := m.Map(source)
	if err != nil {
		assert.Error(t, err, "error mapping set")
	}
	expectedImageSets := []JSONImageSet{
		JSONImageSet{
			UUID: "1376ed33-0d81-3f62-ad62-a9b87b473556",
			Members: []JSONMember{
				JSONMember{
					UUID: "41614f4c-13c5-11e7-9469-afea892e4de3",
				},
				JSONMember{
					UUID: "4258f26a-13c5-11e7-9469-afea892e4de3",
				},
				JSONMember{
					UUID: "3ff3b7a8-13c5-11e7-9469-afea892e4de3",
				},
			},
		},
		JSONImageSet{
			UUID: "89e79a93-1bcc-39d6-bcc4-e77b82d3712f",
			Members: []JSONMember{
				JSONMember{
					UUID: "2fe0b459-a23e-452d-a2aa-2e0503982ed2",
				},
				JSONMember{
					UUID: "404cf8d9-1b88-4883-8afe-580e5174830d",
				},
				JSONMember{
					UUID: "2acf1caa-8014-48ec-b070-a0ffbc45d1d5",
				},
			},
		},
	}
	assert.Equal(t, expectedImageSets, actualImageSets)
}