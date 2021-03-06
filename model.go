package dlframework

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
	"github.com/unknwon/com"
	"golang.org/x/sync/syncmap"

	"github.com/c3sr/config"
	yaml "gopkg.in/yaml.v2"
)

var modelRegistry = syncmap.Map{}
var DefaultModelElementType string = "float32"

func (model ModelManifest) GetElementType() string {
	modelInputs := model.GetInputs()
	typeParameters := modelInputs[0].GetParameters()
	if typeParameters == nil {
		return DefaultModelElementType
	}
	pet, ok := typeParameters["element_type"]
	if !ok {
		return DefaultModelElementType
	}
	petVal := pet.Value
	if petVal == "" {
		return DefaultModelElementType
	}

	var val string
	if err := yaml.Unmarshal([]byte(petVal), &val); err != nil {
		log.Errorf("unable to get element type %v as a string", petVal)
		return DefaultModelElementType
	}

	return strings.ToLower(val)
}

func (model ModelManifest) GetInputDimensions() ([]int, error) {
	modelInputs := model.GetInputs()

	typeParameters := modelInputs[0].GetParameters()

	if typeParameters == nil {
		return nil, errors.New("invalid type parameters")
	}
	pdims, ok := typeParameters["dimensions"]
	if !ok {
		log.Debug("arbitrary input dimensions")
		return nil, nil
	}
	pdimsVal := pdims.GetValue()
	if pdimsVal == "" {
		return nil, errors.New("invalid input dimensions")
	}

	var dims []int
	if err := yaml.Unmarshal([]byte(pdimsVal), &dims); err != nil {
		return nil, errors.Errorf("unable to get input dimensions %v as an integer slice", pdimsVal)
	}

	return dims, nil
}

func (m *ModelManifest) Modality() (Modality, error) {
	inputs := m.GetInputs()
	output := m.GetOutput()

	if len(inputs) == 0 {
		return "", errors.New("no input is specified")
	}
	input := inputs[0]
	switch strings.ToLower(input.Type) {
	case "image":
		switch strings.ToLower(output.Type) {
		case "image":
			return ImageEnhancementModality, nil
		case "classification":
			return ImageClassificationModality, nil
		case "boundingbox":
			return ImageObjectDetectionModality, nil
		case "semanticsegment":
			return ImageSemanticSegmentationModality, nil
		case "instancesegment":
			return ImageInstanceSegmentationModality, nil
		}
	case "raw":
		return RawModality, nil
	case "general":
		return GeneralModality, nil
	}

	panic("unhandled modality")
	return UnknownModality, nil
}

func (m *ModelManifest) Validate() error {
	name := m.GetName()
	if name == "" {
		return errors.New("model name cannot be empty")
	}
	version := m.GetVersion()
	if version == "" {
		version = "latest"
	}
	if version != "latest" {
		if _, err := semver.NewVersion(version); err != nil {
			return errors.Wrapf(err,
				"the version %s for the model %s is not in semantic version format",
				m.GetVersion(),
				m.GetName(),
			)
		}
	}
	return nil
}

func (m ModelManifest) MustCanonicalName() string {
	s, err := m.CanonicalName()
	if err != nil {
		log.WithField("model_name", m.GetName()).Fatal("unable to get model canonical name")
		return ""
	}
	return s
}

func (m ModelManifest) CanonicalName() (string, error) {
	if m.GetFramework() == nil {
		return "", errors.Errorf("the model %s does not have a valid framework", m.GetName())
	}
	framework, err := m.ResolveFramework()
	if err != nil {
		return "", err
	}
	frameworkName, err := framework.CanonicalName()
	if err != nil {
		return "", errors.Wrapf(err, "cannot get canonical name for the framework %s and model %s in the registry", m.GetFramework().GetName(), m.GetName())
	}
	fm, ok := frameworkRegistry.Load(frameworkName)
	if !ok {
		return "", errors.Wrapf(err, "cannot get framework %s for model %s in the registry", frameworkName, m.GetName())
	}
	if _, ok := fm.(FrameworkManifest); !ok {
		return "", errors.Errorf("invalid framework %s registered for model %s in the registry", frameworkName, m.GetName())
	}
	modelName := CleanString(m.GetName())
	if modelName == "" {
		return "", errors.New("model name must not be empty")
	}
	modelVersion := CleanString(m.GetVersion())
	if modelVersion == "" {
		modelVersion = "latest"
	}
	return frameworkName + "/" + modelName + ":" + modelVersion, nil
}

func (m ModelManifest) FrameworkConstraint() (*semver.Constraints, error) {
	return semver.NewConstraint(m.GetFramework().GetVersion())
}

func (m ModelManifest) MustResolveFramework() FrameworkManifest {
	f, err := m.ResolveFramework()
	if err != nil {
		panic(err)
	}
	return f
}

func (m ModelManifest) ResolveFramework() (FrameworkManifest, error) {
	frameworks, err := Frameworks()
	if err != nil {
		return FrameworkManifest{}, err
	}
	if len(frameworks) == 0 {
		return FrameworkManifest{}, errors.New("framework not found")
	}
	if m.GetFramework().GetVersion() == "lastest" {
		sort.Slice(frameworks, func(ii, jj int) bool {
			f1, _ := semver.NewVersion(frameworks[ii].GetVersion())
			f2, _ := semver.NewVersion(frameworks[jj].GetVersion())
			return f1.LessThan(f2)
		})
		return frameworks[len(frameworks)-1], nil
	}

	modelFrameworkConstraint, err := m.FrameworkConstraint()
	if err != nil {
		return FrameworkManifest{},
			errors.Wrapf(err, "cannot get framework constraints for model %v", m.GetName())
	}

	filtered := []FrameworkManifest{}
	for _, framework := range frameworks {
		frameworkVersion, err := semver.NewVersion(framework.GetVersion())
		if err != nil {
			continue
		}
		if !modelFrameworkConstraint.Check(frameworkVersion) {
			continue
		}
		filtered = append(filtered, framework)
	}
	if len(frameworks) == 0 {
		return FrameworkManifest{}, errors.New("framework not found")
	}
	sort.Slice(filtered, func(ii, jj int) bool {
		f1, _ := semver.NewVersion(filtered[ii].GetVersion())
		f2, _ := semver.NewVersion(filtered[jj].GetVersion())
		return f1.LessThan(f2)
	})
	return filtered[len(filtered)-1], nil
}

func (m ModelManifest) Register() error {
	n, err := m.CanonicalName()
	if err != nil {
		return err
	}
	return m.RegisterNamed(n)
}

func (m ModelManifest) RegisterNamed(s string) error {
	name := strings.ToLower(s)
	if _, ok := modelRegistry.Load(name); ok {
		return errors.Errorf("the %s model has already been registered", s)
	}
	modelRegistry.Store(s, m)
	return nil
}

func RegisteredModelNames() []string {
	return syncMapKeys(&modelRegistry)
}

func Models() ([]ModelManifest, error) {
	names := RegisteredModelNames()
	models := make([]ModelManifest, len(names))
	for ii, name := range names {
		m, ok := modelRegistry.Load(name)
		if !ok {
			return nil, errors.Errorf("model %s was not found", name)
		}
		model, ok := m.(ModelManifest)
		if !ok {
			return nil, errors.Errorf("model %s was found but not of type ModelManifest", name)
		}
		models[ii] = model
	}
	sort.Slice(models, func(ii, jj int) bool {
		return models[ii].GetName() < models[jj].GetName()
	})
	return models, nil
}

func FindModel(name string) (*ModelManifest, error) {
	var model *ModelManifest
	modelRegistry.Range(func(key0 interface{}, value interface{}) bool {
		key, ok := key0.(string)
		if !ok {
			return true
		}
		if key != name {
			return true
		}
		m, ok := value.(ModelManifest)
		if !ok {
			return true
		}
		model = &m
		return false
	})
	if model == nil {
		return nil, errors.Errorf("model %s not found in registry", name)
	}
	return model, nil
}

func (m *ModelManifest) WorkDir() (string, error) {
	cannonicalName, err := m.CanonicalName()
	if err != nil {
		return "", err
	}
	cannonicalName = strings.Replace(cannonicalName, ":", "_", -1)
	cannonicalName = strings.Replace(cannonicalName, " ", "_", -1)
	cannonicalName = strings.Replace(cannonicalName, "-", "_", -1)

	workDir := filepath.Join(config.App.TempDir, "dlframework", cannonicalName)
	if !com.IsDir(workDir) {
		if err := os.MkdirAll(workDir, 0700); err != nil {
			return "", errors.Wrapf(err, "failed to create model work directory %v", workDir)
		}
	}
	return workDir, nil
}
