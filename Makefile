# Variabel
FileNameBuild=main
FileBuildPath=./build/
FullPathBuildFile=$(FileBuildPath)$(FileNameBuild)


# Function MakeFile
clean:
	@rm -rf $(FileBuildPath)
build:
	@go build -o $(FullPathBuildFile) ./cmd/api/main.go 
run: clean build
	@$(FullPathBuildFile)