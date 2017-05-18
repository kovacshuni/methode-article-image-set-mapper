package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Sirupsen/logrus"
)

const compoundStory = "EOM::CompoundStory"

type ImageSetMapper interface {
	Map(source NativeContent, lastModified string, publishReference string) ([]JSONImageSet, error)
}

type defaultImageSetMapper struct {
	articleToImageSetMapper ArticleToImageSetMapper
	attributesMapper        AttributesMapper
	xmlImageSetToJSONMapper XMLImageSetToJSONMapper
}

func newImageSetMapper(articleToImageSetMApper ArticleToImageSetMapper, attributesMapper AttributesMapper,
	xmlImageSetToJSONMapper XMLImageSetToJSONMapper) ImageSetMapper {
	return defaultImageSetMapper{
		articleToImageSetMapper: articleToImageSetMApper,
		attributesMapper: attributesMapper,
		xmlImageSetToJSONMapper: xmlImageSetToJSONMapper,
	}
}

func (m defaultImageSetMapper) Map(source NativeContent, lastModified string, publishReference string) ([]JSONImageSet, error) {
	valueXml, err := base64.StdEncoding.DecodeString(source.Value)
	if err != nil {
		msg := fmt.Errorf("Cound't decode string as base64. %v\n", err)
		logrus.Warn(msg)
		return nil, msg
	}

	xmlImageSets, err := m.articleToImageSetMapper.Map(valueXml)
	if err != nil {
		msg := fmt.Errorf("Couldn't parse XML document. %v\n", err)
		logrus.Warn(msg)
		return nil, msg
	}

	attributes, err  := m.attributesMapper.Map(source.Attributes)
	if err != nil {
		msg := fmt.Errorf("Couldn't parse attributes XML. %v\n", err)
		logrus.Warn(msg)
		return nil, msg
	}

	jsonImageSets, err := m.xmlImageSetToJSONMapper.Map(xmlImageSets, attributes, lastModified, publishReference)
	if err != nil {
		msg := fmt.Errorf("Couldn't map ImageSets from model soruced from XML to model targeted for JSON. %v\n", err)
		logrus.Warn(msg)
		return nil, msg
	}
	return jsonImageSets, nil
}
