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
- **100% Swiss Ephemeris v2.10.03 Parity**: All 106 exported C functions from `swephexp.h` are bound.
- **Zero-Allocation Core**: Calculations run at native C speed with 0 B/op and 0 allocs/op (over 5x faster than legacy bindings with no GC finalizer pressure).
- **Memory Safe & Robust**: Automatic buffer bounds protection, null-termination guarantees, and 64-bit/32-bit integer safety across all platforms.
- **Optimized Concurrency**: Mutex protection where required for Swiss Ephemeris internal state, while eliminating contention on pure mathematical routines.
- **Idiomatic Go Helpers**: Non-breaking ergonomic helpers for `time.Time` integration, structured coordinates (`CalcUtSimple`), house results (`HousesSimple`), and string APIs (`VersionStr`, `SetEphePathStr`).
- **Comprehensive Test Suite**: Unit tests, memory safety tests, and race detector verified (`go test -race`).

## License
This project is licensed under the GNU Affero General Public License v3 (AGPLv3). See `LICENSE` for details.
