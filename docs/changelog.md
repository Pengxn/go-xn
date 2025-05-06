# Changelog

## [Unreleased]

## [0.0.7] - 2025-04-26

### Added

- Add `NewRelicWriter` to send logs to NewRelic API. #440
- Update `LoggerConfig` support multiple log writers. #453
- Add `OpenTelemetry` integration for distributed tracing. #454

### Changed

- Refactor `BarkHook` logrus hook to `BarkWriter`. #438 #445
- Refactor `TelegramHook` logrus hook to `TelegramWriter`. #439 #445
- Regenerate proto files with latest protoc. #446
- Migrate `github.com/urfave/cli` to **`v3`**. #449
- Enhance logger to include system and version information. e2b0eff caddbab
- Refactor log writer selection by named return parameter. 7241385
- Refactor database `dsn` string building. #455

### Removed

- Remove deprecated `github.com/Pengxn/go-xn/src/util/log` packages. #441
- Remove deprecated `github.com/Pengxn/go-xn/src/util/log/hook` package. #441

### Fixed

- Fix typo in `agent` subcommand description. 10a19d2
- Fix issue with getting empty version string from git tag in ci workflow. #430

## [0.0.6] - 2025-04-01

### Added

- Add `ping` subcommand to check agent server status. #411
- Add support for (`FLOSS/fund`) funding manifest and discovery. #413
- Add tool dependencies to `tool` directive for `go 1.24+`. #414
- Add `SetLogger` function to configure logger. 9dbb934

### Changed

- Update `golang` image tag to `1.24-alpine` in docker test workflow. df9eb33
- Regenerate proto file with latest `protoc` and `protoc-gen-go`. #412 #424
- Update Makefile to add `test` dependency for `cover` target. b378030
- Disable `LoggerToFile()` middleware from routes initialization. 31b6b69
- Switch to std `log/slog` package for structured logging. #426
- Update the branch name in the nightly URL. 2a5a5a2
- Move `slog` package out of `log` directory structure. #431

### Deprecated

- Deprecated `LogFilePath` function and remove calling it. #425
- Deprecated `github.com/Pengxn/go-xn/src/util/log` package and add doc comments. #429
- Deprecated logrus hooks in `github.com/Pengxn/go-xn/src/util/log/hook` package. #432

### Removed

- Remove deprecated `dns` package and related config. #410 #422
- Remove deprecated azure-pipelines ci config files. #421

### Fixed

- Skip dependabot workflow runs when retrieving artifacts. #409
- Fix whois server extraction regex. #420

### Security

- Fix `CVE-2025-29923`/`GHSA-92cp-5422-2mw7`: bump `github.com/redis/go-redis/v9` to 9.7.3. #416
- Fix `CVE-2025-30204`/`GHSA-mh63-6h87-95cp`: bump `github.com/golang-jwt/jwt/v5` to 5.2.2. #418

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

[Unreleased]: https://github.com/Pengxn/go-xn/compare/0.0.7...HEAD
[0.0.7]: https://github.com/Pengxn/go-xn/compare/0.0.6...0.0.7
[0.0.6]: https://github.com/Pengxn/go-xn/compare/0.0.5...0.0.6
[0.0.5]: https://github.com/Pengxn/go-xn/compare/0.0.4...0.0.5
[0.0.4]: https://github.com/Pengxn/go-xn/compare/0.0.3...0.0.4
[0.0.3]: https://github.com/Pengxn/go-xn/compare/0.0.2...0.0.3
[0.0.2]: https://github.com/Pengxn/go-xn/compare/0.0.1...0.0.2
[0.0.1]: https://github.com/Pengxn/go-xn/releases/tag/0.0.1
