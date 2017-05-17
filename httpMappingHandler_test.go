package main

import (
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/mock"
	"bytes"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/pkg/errors"
	"strings"
)

func TestHttpHandler_Ok(t *testing.T) {
	request, err := http.NewRequest("POST", "/map", bytes.NewReader([]byte("")))
	if err != nil {
		assert.Fail(t, "Error wasn't expected during request creation")
	}
	mockedMessageToNativeMapper := new(mockMessageToNativeMapper)
	mockedMessageToNativeMapper.On("Map", mock.MatchedBy(func(source []byte) bool { return true })).Return(NativeContent{}, nil)
	mockedImageSetMapper := new(mockImageSetMapper)
	mockedImageSetMapper.On("Map", mock.MatchedBy(func(source NativeContent) bool { return true })).Return([]JSONImageSet{}, nil)
	httpHandler := newHTTPMappingHandler(mockedMessageToNativeMapper, mockedImageSetMapper)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandler.handle)
	handler.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, []byte("[]"), recorder.Body.Bytes())
}

func TestHttpHandler_ErrorOnNativeMapper(t *testing.T) {
	request, err := http.NewRequest("POST", "/map", bytes.NewReader([]byte("")))
	if err != nil {
		assert.Fail(t, "Error wasn't expected during request creation")
	}
	mockedMessageToNativeMapper := new(mockMessageToNativeMapper)
	mockedMessageToNativeMapper.On("Map", mock.MatchedBy(func(source []byte) bool { return true })).Return(NativeContent{}, errors.New("error on native mapper"))
	httpHandler := newHTTPMappingHandler(mockedMessageToNativeMapper, nil)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandler.handle)
	handler.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.True(t, strings.Contains(string(recorder.Body.Bytes()), `{"message":"Error mapping native message.`))
}

func TestHttpHandler_ErrorOnImageSetMapper(t *testing.T) {
	request, err := http.NewRequest("POST", "/map", bytes.NewReader([]byte("")))
	if err != nil {
		assert.Fail(t, "Error wasn't expected during request creation")
	}
	mockedMessageToNativeMapper := new(mockMessageToNativeMapper)
	mockedMessageToNativeMapper.On("Map", mock.MatchedBy(func(source []byte) bool { return true })).Return(NativeContent{}, nil)
	mockedImageSetMapper := new(mockImageSetMapper)
	mockedImageSetMapper.On("Map", mock.MatchedBy(func(source NativeContent) bool { return true })).Return([]JSONImageSet{}, errors.New("error on image set mapper"))
	httpHandler := newHTTPMappingHandler(mockedMessageToNativeMapper, mockedImageSetMapper)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandler.handle)
	handler.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.True(t, strings.Contains(string(recorder.Body.Bytes()), `{"message":"Error mapping the given content.`))
}
