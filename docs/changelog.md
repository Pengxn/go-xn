# Changelog

## [Unreleased]

## [0.0.3] - 2025-02-07

### Added

- Add Keybase proof ownership support. #352
- Add OAuth2 authentication route structure. #353
- Add bcrypt password hashing utility functions. #356
- Add health check endpoint. #361
- Add `NoIndex` middleware for search engine control. #365

### Changed

- Improve API routes with `JWT` middleware. #351
- Move option routes into api route group. 5108a65
- Use built-in gRPC status codes instead of custom constants. #359
- Improve static file serving logic for multiple filesystems. #358
- Use std `log` library to log startup errors. #360
- Simplify Makefile BIN variable and add generate target. 5ad823e
- Rename and standardize command variable names. #363
- Update version variable and command flag usage text. #368

### Fixed

- Use multiple filesystems for static files. f86c797
- Fix `.git-blame-ignore-revs` file is included in archive exports. #354
- Enhance WebAuthn user credential handling. #367

### Removed

- Remove deprecated `url` field and replace with `slug`. #366

## [0.0.2] - 2025-01-25

### Added

- Add gRPC implementation. #332
- Add bitcoin BIP15 alias partial support. #336
- Sync Chinese translation language file for inno-setup. #342
- Add cron library for managing scheduled jobs. #345

### Changed

- Bump actions version for GitHub Actions workflows. #303 #304 #324
- Upgrade go toolchain version to 1.23. #299 #321
- Move template debug control to the route layer. #339
- Refactor static files handling and web embed FS constants. #341

### Deprecated

- Use slug instead of url field for the Article model. be2dff6

### Fixed

- Fix non-ASCII domain whois query. #338

### Removed

- Remove log warning prompt when default config file doesn't exist. #249

### Security

- Add security.txt RFC 9116 standard support. #346 #349

## [0.0.1] - 2024-01-13

- Release first version `0.0.1`.

[Unreleased]: https://github.com/Pengxn/go-xn/compare/0.0.3...HEAD
[0.0.3]: https://github.com/Pengxn/go-xn/compare/0.0.2...0.0.3
[0.0.2]: https://github.com/Pengxn/go-xn/compare/0.0.1...0.0.2
[0.0.1]: https://github.com/Pengxn/go-xn/releases/tag/0.0.1
