package recipe

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// recipesDir returns the directory where recipe YAML files are stored.
// It uses os.UserConfigDir() (e.g. ~/.config on Linux, ~/Library/Application Support on macOS)
// and appends /monet/recipes.
func recipesDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("user config dir: %w", err)
	}
	return filepath.Join(configDir, "monet", "recipes"), nil
}

func DebugRecipesDir() {
	dir, err := recipesDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Recipes directory: %s\n", dir)
}

// Load reads a single recipe by name (without the .yaml extension).
// It returns an error wrapping os.ErrNotExist if the file does not exist.
func Load(name string) (Recipe, error) {
	dir, err := recipesDir()
	if err != nil {
		return Recipe{}, err
	}
	path := filepath.Join(dir, name+".yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return Recipe{}, fmt.Errorf("recipe %q: %w", name, err)
	}

	var currentRecipe Recipe
	if err := yaml.Unmarshal(data, &currentRecipe); err != nil {
		return Recipe{}, fmt.Errorf("parsing recipe %q: %w", name, err)
	}
	return currentRecipe, nil
}

// LoadAll reads all recipe YAML files from the recipes directory.
// It returns them sorted by Name. If the directory does not exist,
// it returns an empty slice with no error.
func LoadAll() ([]Recipe, error) {
	dir, err := recipesDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Recipe{}, nil
		}
		return nil, fmt.Errorf("reading recipes dir: %w", err)
	}

	var recipes []Recipe
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".yaml") {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ".yaml")
		singleRecipe, err := Load(name)
		if err != nil {
			// Skip unreadable recipes rather than failing entirely
			continue
		}
		recipes = append(recipes, singleRecipe)
	}

	sort.Slice(recipes, func(i, j int) bool {
		return recipes[i].Name < recipes[j].Name
	})

	return recipes, nil
}
