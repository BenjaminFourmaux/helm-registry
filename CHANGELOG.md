# Changelog
[![](https://badgen.net/github/tag/BenjaminFourmaux/Helm-Registry?cache=600)]() [![](https://badgen.net/github/release/BenjaminFourmaux/Helm-Registry?cache=600)]() [![](https://badgen.net/github/tags/BenjaminFourmaux/Helm-Registry)]()

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## V.1.0 - Basics features requirement 
### Added
- **Charts Discovery**
	- The Charts Directory is the Single Source of Truth (SSoT) !
	- Implement Charts Directory Watcher, that listen when a action (create a file, remove file, modified file...) is dispense.
- **API**
	- Endpoint `/`, the home page to get this registry informations (in YAML output format).
	- Endpoint `/index.yaml`, the index of charts in this registry. Is a Helm requirement.
	- Endpoint `/charts/`, bind with the Charts Directory. Able to get charts archives.
- **Database**
	- SQLite Database to store registered Helm Charts and registry about the registry.
	- Database management class `Class/Database/Database/Database.go`.
	- Classes for `SELECT`, `INSERT`, `UPDATE` and `DELETE` queries.
	- Table `registry` which contains all informations about this registry (name, version, maintainer, tags...).
	- Table `charts` which contains all informations about all registered charts (in charts directory).
- **Docker**
	- Dockerfile : For build a Docker image of the app.
	- Docker compose file :  For deploying a registry container with compose.
- **Env**
	- Manage env var 
- **Misc**
	- Get OS environment that run this app (Windows, Linux or Docker) and adapte the comportement