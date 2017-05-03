package main

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"encoding/base64"
	"fmt"
)

type ImageSetMapper interface {
	Map(source []byte) ([]byte, error)
}

type defaultImageSetMapper struct {
	xmlMapper ArticleToImageSetMapper
	xmlToJSON XMLImageSetToJSONMapper
}

func newImageSetMapper() ImageSetMapper {
	return defaultImageSetMapper {
		xmlMapper: defaultArticleToImageSetMapper{},
		xmlToJSON: defaultImageSetToJSONMapper{},
	}
}

type NativeContent struct {
	Value string `json:"value"`
}

func (m defaultImageSetMapper) Map(source []byte) ([]byte, error) {
	var native NativeContent
	err := json.Unmarshal(source, &native)
	if err != nil {
		msg := fmt.Errorf("Cound't decode native content as JSON doucment. %v\n", err)
		logrus.Warn(msg)
		return nil, msg
	}

	xmlDocument, err := base64.StdEncoding.DecodeString(native.Value)
	if err != nil {
		msg := fmt.Errorf("Cound't decode string as base64. %v\n", err)
		logrus.Warn(msg)
		return nil, msg
	}

	xmlImageSets, err := m.xmlMapper.Map(xmlDocument)
	if err != nil {
		msg := fmt.Errorf("Couldn't parse XML document. %v\n", err)
		logrus.Warn(msg)
		return nil, msg
	}

	jsonImageSets, err := m.xmlToJSON.Map(xmlImageSets)
	if err != nil {
		msg := fmt.Errorf("Couldn't map ImageSets from model soruced from XML to model targeted for JSON. %v\n", err)
		logrus.Warn(msg)
		return nil, msg
	}

	marshaledJSONImageSets, err := json.Marshal(jsonImageSets)
	if err != nil {
		msg := fmt.Errorf("Couldn't marshall built-up image-sets to JSON. %v\n", err)
		logrus.Warn(msg)
		return nil, msg
	}
	return marshaledJSONImageSets, nil
}