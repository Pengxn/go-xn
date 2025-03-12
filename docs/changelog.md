# Changelog

## [Unreleased]

## [0.0.5] - 2025-03-12

### Added

- Add markdown render support by GitHub API. #380
- Implement a function to retrieve the latest artifact link. 93cdccd
- Initialize slog with tint handler. #397
- Add `--nightly`/`-n` flag to support switching to nightly build. #398
- Update dependabot configuration to add `github-actions` package. #391

### Changed

- Normalize artifact names using `$GOOS` and `$GOARCH` variables. #388 #390
- Separate build artifacts for binary and installer in windows workflow. #389
- Update build installer output path. #394
- Use github latest asset link as default update source. #396
- Regenerate the protobuf files and gRPC code. 928bb31
- Bump go toolchain version to `1.24`. #404
- Update build docker image tag to specific pinned version. 68bbea6

### Deprecated

- Mark `github.com/Pengxn/go-xn/src/lib/dns` package as deprecated. #395

### Removed

- Remove test files from the `github.com/Pengxn/go-xn/src/lib/dns` package. #395
- Remove unused `level()` function from the `github.com/Pengxn/go-xn/src/util/log` package. cba1f7b

### Fixed

- Fix artifact name with correct arch for macOS runner image. 3fc2cfc
- Fix workflow artifact retrieval logic to skip incomplete runs. #405
- Fix unzip directory to locate the executable path. 09d29b4

## [0.0.4] - 2025-02-20

### Added

- Implement cloudflare Turnstile server-side validation integration. #348
- Add `agent` command to run the agent gRPC server. #377
- Add `changelog.md` document to track version history. #372
- Add `CODE_OF_CONDUCT.md` file as contributor code of conduct document file. #382
- Add macOS workflow to improve multi-platform testing and build. #385

### Changed

- Adjust default log level to info. f19f9ab
- Render markdown content as HTML for article view. f14a531
- Move `SECURITY.md` file to `docs` directory. #381
- Support multi-platform binary self-updates. #379
- Update outdated comments in `version` command variable for early changes. 30bbb4a
- Update `LICENSE` copyright notice to include other contributors. 6874a46

### Fixed

- Missing version information for command. #371
- Match `.yaml` file extension with 2-space indentation. #375

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

[Unreleased]: https://github.com/Pengxn/go-xn/compare/0.0.5...HEAD
[0.0.5]: https://github.com/Pengxn/go-xn/compare/0.0.4...0.0.5
[0.0.4]: https://github.com/Pengxn/go-xn/compare/0.0.3...0.0.4
[0.0.3]: https://github.com/Pengxn/go-xn/compare/0.0.2...0.0.3
[0.0.2]: https://github.com/Pengxn/go-xn/compare/0.0.1...0.0.2
[0.0.1]: https://github.com/Pengxn/go-xn/releases/tag/0.0.1
