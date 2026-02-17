# gose - Go Swiss Ephemeris

`gose` is a Go interface binding for the [Swiss Ephemeris](https://www.astro.com/swisseph/) C library.

## Installation

Because this package uses CGO to link against the Swiss Ephemeris C library, you must compile and install the C library (`libswe`) first.

### 1. Build the Swiss Ephemeris C Library

The source code for the C library is located at `https://github.com/aloistr/swisseph`.

#### Linux (x86_64, ARM64)
1. Navigate to the source directory:
   ```bash
   cd swisseph
   ```
2. Build the shared library:
   ```bash
   make libswe.so
   ```
   *Note: If the Makefile target is missing, run:*
   ```bash
   cc -g -Wall -fPIC -c *.c
   cc -shared -o libswe.so *.o -lm -ldl
   ```
3. Install the library:
   ```bash
   sudo cp libswe.so /usr/local/lib/
   sudo ldconfig
   ```

#### macOS (Intel, Apple Silicon)
1. Navigate to the source directory:
   ```bash
   cd swisseph
   ```
2. Build the dynamic library:
   ```bash
   cc -g -Wall -fPIC -c *.c
   cc -dynamiclib -o libswe.dylib *.o -lm
   ```
3. Install the library:
   ```bash
   sudo cp libswe.dylib /usr/local/lib/
   ```

#### Windows (x86, x64)
1. Navigate to `swisseph`.
2. Using MinGW/GCC:
   ```bash
   cc -g -Wall -c *.c
   cc -shared -o libswe.dll *.o -lm
   ```
3. Copy `libswe.dll` to your Windows System32 folder or keep it in the same directory as your Go executable.
4. You may need to create an import library (`libswe.a`) for linking.

### 2. Use the Go Package

After the C library is installed, you can use `gose` in your project.

```go
import "github.com/bharathak/gose"
```

## Testing the Installation

To run the provided example, ensure `libswe.so` (or `.dylib`/`.dll`) is in your library path.

```bash
# 1. Build the C library (if not already done)
cd source/swisseph
make libswe.so

# 2. Run the example from the gose/example directory
cd ../../gose/example
LD_LIBRARY_PATH=../../source/swisseph LIBRARY_PATH=../../source/swisseph CGO_LDFLAGS="-L../../source/swisseph" go run main.go
```

## Features
- Full support for Swiss Ephemeris v2.10.03.
- Thread-safe access via a global mutex.
- All planetary, house, and eclipse calculations.
- New crossing functions: `SolcrossUt`, `MooncrossUt`, etc.

## License
This project is licensed under the GNU Affero General Public License v3 (AGPLv3). See `LICENSE` for details.
