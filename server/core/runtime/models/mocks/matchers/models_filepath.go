// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"

	models "github.com/runatlantis/atlantis/server/core/runtime/models"
)

func AnyModelsFilePath() models.FilePath {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.FilePath))(nil)).Elem()))
	var nullValue models.FilePath
	return nullValue
}

func EqModelsFilePath(value models.FilePath) models.FilePath {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.FilePath
	return nullValue
}

func NotEqModelsFilePath(value models.FilePath) models.FilePath {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue models.FilePath
	return nullValue
}

func ModelsFilePathThat(matcher pegomock.ArgumentMatcher) models.FilePath {
	pegomock.RegisterMatcher(matcher)
	var nullValue models.FilePath
	return nullValue
}
