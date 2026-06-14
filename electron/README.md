# New API Electron Desktop App

This directory contains the Electron wrapper for New API, providing a native desktop application with system tray support for Windows, macOS, and Linux.

## Prerequisites

### 1. Go Binary (Required)
The Electron app requires the compiled Go binary to function. You have two options:

**Option A: Use existing binary (without Go installed)**
```bash
# If you have a pre-built binary (e.g., new-api-macos)
cp ../new-api-macos ../new-api
```

**Option B: Build from source (requires Go)**

Requires **Go 1.25+** and a **C compiler** — the backend is built with `CGO_ENABLED=1`.

`main.go` embeds the web assets via `//go:embed web/default/dist` and `//go:embed web/classic/dist`, so the **frontend must be built first** or the Go build will fail.

```bash
# From the repo root

# 1. Build both frontends (requires bun)
make build-all-frontends

# 2. Build the Go backend into ./new-api (./new-api.exe on Windows)
CGO_ENABLED=1 go build -ldflags="-s -w" -o new-api
```

The resulting `./new-api` at the repo root is the same binary Option A copies into place, so the Electron app (`cd electron && npm start`) picks it up automatically. To build the packaged desktop app for your current platform in one step, use the helper script instead:

```bash
cd electron
./build.sh
```

### 3. Electron Dependencies
```bash
cd electron
npm install
```

## Development

Run the app in development mode:
```bash
npm start
```

This will:
- Start the Go backend on port 3000
- Open an Electron window with DevTools enabled
- Create a system tray icon (menu bar on macOS)
- Store database in `../data/new-api.db`

## Building for Production

### Quick Build
```bash
# Ensure Go binary exists in parent directory
ls ../new-api  # Should exist

# Build for current platform
npm run build

# Platform-specific builds
npm run build:mac    # Creates .dmg and .zip
npm run build:win    # Creates .exe installer
npm run build:linux  # Creates .AppImage and .deb
```

### Build Output
- Built applications are in `electron/dist/`
- macOS: `.dmg` (installer) and `.zip` (portable)
- Windows: `.exe` (installer) and portable exe
- Linux: `.AppImage` and `.deb`

## Configuration

### Port
Default port is 3000. To change, edit `main.js`:
```javascript
const PORT = 3000; // Change to desired port
```

### Database Location
- **Development**: `../data/new-api.db` (project directory)
- **Production**:
  - macOS: `~/Library/Application Support/New API/data/`
  - Windows: `%APPDATA%/New API/data/`
  - Linux: `~/.config/New API/data/`
