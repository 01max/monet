# monet

monet is a CLI and TUI tool for authoring and reusing ImageMagick recipes

Named image-processing pipelines like "resize to 200×200, quality 85%, output as JPEG". 

Recipes are stored as YAML files and can be run via `monet apply` or installed as shell wrappers (`monet bin install`) for direct use. 

Built as a Go learning project, monet uses the stdlib for core logic with Bubble Tea for the terminal UI and YAML for recipe storage.
