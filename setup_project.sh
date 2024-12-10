#!/bin/bash

# Root project folder
ROOT_DIR="."

# Create folders
mkdir -p $ROOT_DIR/{cmd/dummydata,internal/{generator,api,storage,utils},pkg/dummydata,web/{templates,static},config,scripts,examples/{basic,advanced,api},docs,test/{integration,performance}}

# Create files with placeholders or empty content
touch $ROOT_DIR/cmd/dummydata/main.go
touch $ROOT_DIR/internal/generator/{generator.go,generator_test.go,config.go,primitive.go,structured.go,domain.go,media.go,timeseries.go,geospatial.go,validation.go}
touch $ROOT_DIR/internal/api/{handler.go,middleware.go,router.go,server.go}
touch $ROOT_DIR/internal/storage/{file.go,database.go,cache.go}
touch $ROOT_DIR/internal/utils/{random.go,validation.go,converter.go}
touch $ROOT_DIR/pkg/dummydata/{types.go,options.go,client.go}
touch $ROOT_DIR/web/{handler.go}
touch $ROOT_DIR/config/{config.go,defaults.go}
touch $ROOT_DIR/scripts/{build.sh,test.sh}
touch $ROOT_DIR/examples/{basic/.keep,advanced/.keep,api/.keep} # .keep to maintain empty dirs
touch $ROOT_DIR/docs/{api.md,cli.md,examples.md}
touch $ROOT_DIR/test/{integration/.keep,performance/.keep} # .keep to maintain empty dirs
touch $ROOT_DIR/{go.mod,go.sum,Makefile,Dockerfile,docker-compose.yml,.gitignore,README.md,LICENSE}

# Populate .gitignore with common Go exclusions
cat <<EOL > $ROOT_DIR/.gitignore
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test

# Build output
bin/

# Dependency directories
vendor/

# Go build cache
*.out
EOL

# Make README.md a placeholder
# echo "# Super Dummy Data Generator" > $ROOT_DIR/README.md
# echo "A versatile dummy data generation library for various use cases." >> $ROOT_DIR/README.md

echo "Project structure created successfully in $ROOT_DIR"
