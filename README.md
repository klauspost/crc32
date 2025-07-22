# 2025 revival

For IEEE checksums AVX512 can be used to speed up CRC32 checksums by approximately 2x.

Castagnoli checksums (CRC32C) can also be computer with AVX512, 
but the performance gain is not as significant enough for the downsides of using it at this point.

# crc32

This package is a drop-in replacement for the standard library `hash/crc32` package, 
that features AVX 512 optimizations on x64 platforms, for a 2x speedup for IEEE CRC32 checksums.

# usage

Install using `go get github.com/klauspost/crc32`. This library is based on Go 1.24

Replace `import "hash/crc32"` with `import "github.com/klauspost/crc32"` and you are good to go.

# changes
* 2025: Revived and updated to Go 1.24, with AVX 512 optimizations.

# performance


# license

Standard Go license. See [LICENSE](LICENSE) for details.
