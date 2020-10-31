package registryquery

import (
	"path"
	"runtime"
	"sort"
	"sync"

	"github.com/Masterminds/semver"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rai-project/config"
	"github.com/c3sr/dlframework"
	dl "github.com/c3sr/dlframework"
	webmodels "github.com/c3sr/dlframework/httpapi/models"
	"github.com/rai-project/parallel/tunny"
	kv "github.com/rai-project/registry"
	"github.com/rai-project/serializer"
)

type modelsTy struct {
	serializer serializer.Serializer
}

var Models modelsTy

func (m modelsTy) Manifests(frameworkName, frameworkVersion string) ([]*webmodels.DlframeworkModelManifest, error) {
	frameworkName = dl.CleanString(frameworkName)
	frameworkVersion = dl.CleanString(frameworkVersion)

	fs, err := Frameworks.Manifests()
	if err != nil {
		return nil, err
	}
	fs, err = Frameworks.FilterManifests(fs, frameworkName, frameworkVersion)
	rgs, err := kv.New()
	if err != nil {
		return nil, err
	}
	defer rgs.Close()

	var manifestsLock sync.Mutex
	var wg sync.WaitGroup
	manifests := []*webmodels.DlframeworkModelManifest{}

	poolSize := runtime.NumCPU()
	pool, err := tunny.CreatePool(poolSize, func(object interface{}) interface{} {
		key, ok := object.(string)
		if !ok {
			return errors.New("invalid key type. expecting a string type")
		}

		if path.Base(key) != "manifest.json" {
			return errors.New("skipping non manifest files")
		}

		e, err := rgs.Get(key)
		if err != nil {
			return err
		}
		registryValue := e.Value
		if registryValue == nil || len(registryValue) == 0 {
			return nil
		}

		model := new(dlframework.ModelManifest)
		if err := m.serializer.Unmarshal(registryValue, model); err != nil {
			return err
		}
		res := new(webmodels.DlframeworkModelManifest)
		if err := copier.Copy(res, model); err != nil {
			return err
		}

		manifestsLock.Lock()
		defer manifestsLock.Unlock()

		manifests = append(manifests, res)
		return nil
	}).Open()
	if err != nil {
		return nil, err
	}

	defer pool.Close()

	prefixKey := path.Join(config.App.Name, "registry")
	for _, framework := range fs {
		frameworkName, frameworkVersion := dl.CleanString(framework.Name), dl.CleanString(framework.Version)
		key := path.Join(prefixKey, frameworkName, frameworkVersion)
		kvs, err := rgs.List(key)
		if err != nil {
			continue
		}
		for _, kv := range kvs {
			if path.Dir(kv.Key) == key {
				continue
			}
			wg.Add(1)
			pool.SendWorkAsync(kv.Key, func(interface{}, error) {
				wg.Done()
			})
		}
	}

	wg.Wait()

	return manifests, nil
}

func (m modelsTy) AllManifests() ([]*webmodels.DlframeworkModelManifest, error) {
	return m.Manifests("*", "*")
}

func (modelsTy) FilterManifests(
	manifests []*webmodels.DlframeworkModelManifest,
	modelName,
	modelVersionString string,
) ([]*webmodels.DlframeworkModelManifest, error) {
	modelName = dl.CleanString(modelName)
	modelVersionString = dl.CleanString(modelVersionString)

	candidates := []*webmodels.DlframeworkModelManifest{}
	for _, manifest := range manifests {
		if modelName == "*" || dl.CleanString(manifest.Name) == modelName {
			candidates = append(candidates, manifest)
		}
	}

	if len(candidates) == 0 {
		return nil, errors.Errorf("model %s not found", modelName)
	}

	if modelVersionString == "" || modelVersionString == "*" {
		return candidates, nil
	}

	sortByVersion := func(elems []*webmodels.DlframeworkModelManifest) func(ii, jj int) bool {
		return func(ii, jj int) bool {
			f1, e1 := semver.NewVersion(elems[ii].Version)
			if e1 != nil {
				return false
			}
			f2, e2 := semver.NewVersion(elems[jj].Version)
			if e2 != nil {
				return false
			}
			return f1.LessThan(f2)
		}
	}

	if modelVersionString == "latest" {
		sort.Slice(candidates, sortByVersion(candidates))
		return []*webmodels.DlframeworkModelManifest{candidates[0]}, nil
	}

	modelVersion, err := semver.NewConstraint(modelVersionString)
	if err != nil {
		return nil, err
	}

	res := []*webmodels.DlframeworkModelManifest{}
	for _, manifest := range manifests {

		c, err := semver.NewVersion(manifest.Version)
		if err != nil {
			continue
		}
		if !modelVersion.Check(c) {
			continue
		}
		res = append(res, manifest)
	}
	if len(res) == 0 {
		return nil, errors.Errorf("model %s=%s not found", modelName, modelVersionString)
	}
	sort.Slice(res, sortByVersion(res))

	return []*webmodels.DlframeworkModelManifest{res[0]}, nil
}

func init() {
	config.AfterInit(func() {
		Models = modelsTy{
			serializer: kv.Config.Serializer,
		}
	})
}
