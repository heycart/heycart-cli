package extension

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/shyim/go-version"

	"github.com/heycart/heycart-cli/internal/validation"
)

type HeyCartBundle struct {
	path     string
	Composer heycartBundleComposerJson
	config   *Config
}

func newHeyCartBundle(path string) (*HeyCartBundle, error) {
	composerJsonFile := fmt.Sprintf("%s/composer.json", path)
	if _, err := os.Stat(composerJsonFile); err != nil {
		return nil, err
	}

	jsonFile, err := os.ReadFile(composerJsonFile)
	if err != nil {
		return nil, fmt.Errorf("newHeyCartBundle: %v", err)
	}

	var composerJson heycartBundleComposerJson
	err = json.Unmarshal(jsonFile, &composerJson)
	if err != nil {
		return nil, fmt.Errorf("newHeyCartBundle: %v", err)
	}

	if composerJson.Type != "heycart-bundle" {
		return nil, fmt.Errorf("newHeyCartBundle: composer.json type is not heycart-bundle")
	}

	if composerJson.Extra.BundleName == "" {
		return nil, fmt.Errorf("composer.json does not contain heycart-bundle-name in extra")
	}

	cfg, err := readExtensionConfig(path)
	if err != nil {
		return nil, fmt.Errorf("newHeyCartBundle: %v", err)
	}

	extension := HeyCartBundle{
		Composer: composerJson,
		path:     path,
		config:   cfg,
	}

	return &extension, nil
}

type composerAutoload struct {
	Psr4 map[string]string `json:"psr-4"`
}

type heycartBundleComposerJson struct {
	Name     string                         `json:"name"`
	Type     string                         `json:"type"`
	License  string                         `json:"license"`
	Version  string                         `json:"version"`
	Require  map[string]string              `json:"require"`
	Extra    heycartBundleComposerJsonExtra `json:"extra"`
	Suggest  map[string]string              `json:"suggest"`
	Autoload composerAutoload               `json:"autoload"`
}

type heycartBundleComposerJsonExtra struct {
	BundleName string `json:"heycart-bundle-name"`
}

func (p HeyCartBundle) GetComposerName() (string, error) {
	return p.Composer.Name, nil
}

// GetRootDir returns the src directory of the bundle.
func (p HeyCartBundle) GetRootDir() string {
	return path.Join(p.path, "src")
}

func (p HeyCartBundle) GetSourceDirs() []string {
	var result []string

	for _, val := range p.Composer.Autoload.Psr4 {
		result = append(result, path.Join(p.path, val))
	}

	return result
}

// GetResourcesDir returns the resources directory of the heycart bundle.
func (p HeyCartBundle) GetResourcesDir() string {
	return path.Join(p.GetRootDir(), "Resources")
}

func (p HeyCartBundle) GetResourcesDirs() []string {
	var result []string

	for _, val := range p.GetSourceDirs() {
		result = append(result, path.Join(val, "Resources"))
	}

	return result
}

func (p HeyCartBundle) GetName() (string, error) {
	return p.Composer.Extra.BundleName, nil
}

func (p HeyCartBundle) GetExtensionConfig() *Config {
	return p.config
}

func (p HeyCartBundle) GetHeyCartVersionConstraint() (*version.Constraints, error) {
	if p.config != nil && p.config.Build.HeyCartVersionConstraint != "" {
		constraint, err := version.NewConstraint(p.config.Build.HeyCartVersionConstraint)
		if err != nil {
			return nil, err
		}

		return &constraint, nil
	}

	heycartConstraintString, ok := p.Composer.Require["heycart/core"]

	if !ok {
		return nil, fmt.Errorf("require.heycart/core is required")
	}

	heycartConstraint, err := version.NewConstraint(heycartConstraintString)
	if err != nil {
		return nil, err
	}

	return &heycartConstraint, err
}

func (HeyCartBundle) GetType() string {
	return TypeHeyCartBundle
}

func (p HeyCartBundle) GetVersion() (*version.Version, error) {
	return version.NewVersion(p.Composer.Version)
}

func (p HeyCartBundle) GetChangelog() (*ExtensionChangelog, error) {
	return parseExtensionMarkdownChangelog(p)
}

func (p HeyCartBundle) GetLicense() (string, error) {
	return p.Composer.License, nil
}

func (p HeyCartBundle) GetPath() string {
	return p.path
}

func (p HeyCartBundle) GetIconPath() string {
	return ""
}

func (p HeyCartBundle) GetMetaData() *extensionMetadata {
	return &extensionMetadata{
		Label: extensionTranslated{
			Chinese: "FALLBACK",
			English: "FALLBACK",
		},
		Description: extensionTranslated{
			Chinese: "FALLBACK",
			English: "FALLBACK",
		},
	}
}

func (p HeyCartBundle) Validate(c context.Context, check validation.Check) {
	// HeyCartBundle validation is currently empty but signature updated to match interface
}
