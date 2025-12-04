import os

project_root = "concurrent_kv_store"

folders = [
    f"{project_root}/cmd/benchmain",
    f"{project_root}/pkg/kv",
    f"{project_root}/pkg/bench"
]

files = {
    f"{project_root}/cmd/benchmain/main.go": "",
    f"{project_root}/pkg/kv/concurrent.go": "",
    f"{project_root}/pkg/kv/baseline.go": "",
    f"{project_root}/pkg/bench/bench.go": "",
    f"{project_root}/go.mod": "",
    f"{project_root}/README.md": "",
    f"{project_root}/RESULTS.txt": ""
}

# Create directories
for folder in folders:
    os.makedirs(folder, exist_ok=True)

# Create files
for file_path, content in files.items():
    # Ensure parent folder exists (good safety check)
    os.makedirs(os.path.dirname(file_path), exist_ok=True)
    with open(file_path, "w") as f:
        f.write(content)

print("Go project structure created successfully!")
