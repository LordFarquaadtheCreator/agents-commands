---
name: safely-unpack-zip
description: Unpack zip files inside a Docker container (via Colima) and serve over HTTP. Nothing touches the local filesystem.
argument-hint: "[directory]"
allowed-tools:
  - exec
  - read
  - grep
  - glob
---

Unpack all zip files in $ARGUMENTS (default: ~/Downloads) and serve them over http://127.0.0.1:8080 — entirely inside a Docker container. Nothing is extracted locally.

## Steps

1. Ensure Colima is running. If not, start it:
   ```
   colima start
   ```

2. Check ~/.docker/config.json — if `"credsStore": "desktop"` is present, remove it (breaks Docker without Docker Desktop).

3. Find all `.zip` files (type f only, not directories) in the target directory:
   ```
   find <dir> -maxdepth 1 -name "*.zip" -type f
   ```

4. Kill any container already holding port 8080:
   ```
   docker ps -q --filter "publish=8080" | xargs docker kill
   ```

5. Build a `docker run` command that:
   - Mounts each zip as `-v "<abs_path>:/tmp/zips/<name>.zip:ro"`
   - Uses `--cap-drop ALL` and `-p 127.0.0.1:8080:8080`
   - Uses `python:alpine` image
   - Runs: `for z in /tmp/zips/*.zip; do name=$(basename "$z" .zip); mkdir -p "/tmp/out/$name"; unzip -o "$z" -d "/tmp/out/$name"; done && cd /tmp/out && python -m http.server 8080`

6. Run the container in the background. Verify with:
   ```
   curl -s http://127.0.0.1:8080/
   ```
   Confirm each chapter directory appears in the listing.

## Notes

- Always use absolute paths in `-v` mounts — never `$(pwd)` inside exec tool calls.
- Quote all paths with spaces.
- Each zip unpacks into its own subdirectory named after the zip (minus `.zip`).
- Server is localhost-only (`127.0.0.1:8080`), not exposed to the network.
- Before running, group related files into zips (e.g., `Chapter 1.pdf` + `Chapter 2.pdf` → `Chapters.zip`). Delete source files after zipping.
