<a name="v1.6.14"></a>

## [v1.6.14](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.6.6...v1.6.14)

> 2024-06-23

### Chore

* upgrade deps
* merge latest codes
* upgrade deps
* merge latest codes
* upgrade deps and fix agent test
* merge latest codes
* upgrade deps
* merge latest codes
* upgrade go version
* upgrade deps
* merge latest codes
* upgrade jwt
* upgrade deps
* upgrade google.golang.org/protobuf to v1.33.0
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps and merge latest codes
* upgrade deps
* upgrade deps
* update deps
* update deps
* upgrade grpc to v1.59.0
* upgrade deps
* upgrade deps
* upgrade otlp
* upgrade deps
* upgrade jaeger
* upgrade deps
* upgrade deps
* merge latest codes
* update deps
* merge latest codes
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* move original license to .originallicense dir
* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* client-ip middle ware
* add env config for redis
* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* add log for client ip handler
* use 1 as default tenant id
* use 0 as default tenant id
* optimize rest middleware
* wrong handler in middleware
* optimize rest middleware
* optimize comments and add new struct
* merge latest codes
* upgrade go redis to v9.5.1 to solve the incompatible bug for older version
* disabled identify for redis because in v9.5.0 will break apps
* use new redis init function in test
* request test
* embedded span
* use new version name
* drop tools file
* test files which use jaeger
* add comment for errors in errorx
* many bugs in library
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* optimize client ip name
* use new key for middleware
* remove tools dir
* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#189](https://github.com/suyuan32/simple-admin-tools/issues/189) from suyuan32/dev
* Merge pull request [#182](https://github.com/suyuan32/simple-admin-tools/issues/182) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.15.1
* Merge pull request [#185](https://github.com/suyuan32/simple-admin-tools/issues/185) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/jedib0t/go-pretty/v6-6.4.7
* Merge pull request [#184](https://github.com/suyuan32/simple-admin-tools/issues/184) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.5
* Merge pull request [#186](https://github.com/suyuan32/simple-admin-tools/issues/186) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.5
* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5
* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.6.6"></a>

## [tools/goctl/v1.6.6](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.13...tools/goctl/v1.6.6)

> 2024-06-21

### Chore

* add trending badge ([#4194](https://github.com/suyuan32/simple-admin-tools/issues/4194))

<a name="v1.6.13"></a>

## [v1.6.13](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.12...v1.6.13)

> 2024-06-05

### Chore

* upgrade deps
* merge latest codes
* fix the confused log level in comment ([#4175](https://github.com/suyuan32/simple-admin-tools/issues/4175))

### Fix

* fix readme ([#4174](https://github.com/suyuan32/simple-admin-tools/issues/4174))
* log concurrency problems after calling WithXXX
  methods ([#4164](https://github.com/suyuan32/simple-admin-tools/issues/4164))
* use 1 as default tenant id

### Refactor

* optimize client ip name

<a name="v1.6.12"></a>

## [v1.6.12](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.12-beta...v1.6.12)

> 2024-05-24

### Fix

* use 0 as default tenant id

<a name="v1.6.12-beta"></a>

## [v1.6.12-beta](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.11...v1.6.12-beta)

> 2024-05-24

### Feat

* client-ip middle ware

<a name="v1.6.11"></a>

## [v1.6.11](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.11-beta...v1.6.11)

> 2024-05-15

### Fix

* optimize rest middleware
* wrong handler in middleware

### Refactor

* use new key for middleware

<a name="v1.6.11-beta"></a>

## [v1.6.11-beta](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.6.5...v1.6.11-beta)

> 2024-05-15

### Chore

* upgrade deps and fix agent test
* merge latest codes
* upgrade deps
* merge latest codes
* upgrade go version
* upgrade deps
* merge latest codes
* upgrade jwt
* upgrade deps
* upgrade google.golang.org/protobuf to v1.33.0
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps and merge latest codes
* upgrade deps
* upgrade deps
* update deps
* update deps
* upgrade grpc to v1.59.0
* upgrade deps
* upgrade deps
* upgrade otlp
* upgrade deps
* upgrade jaeger
* upgrade deps
* upgrade deps
* merge latest codes
* update deps
* merge latest codes
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* move original license to .originallicense dir
* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* add secure option for sending traces via
  otlphttp ([#3973](https://github.com/suyuan32/simple-admin-tools/issues/3973))
* add env config for redis
* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* optimize rest middleware
* optimize comments and add new struct
* merge latest codes
* upgrade go redis to v9.5.1 to solve the incompatible bug for older version
* disabled identify for redis because in v9.5.0 will break apps
* use new redis init function in test
* request test
* embedded span
* use new version name
* drop tools file
* test files which use jaeger
* add comment for errors in errorx
* many bugs in library
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Optimize

* simplify breaker algorithm ([#4151](https://github.com/suyuan32/simple-admin-tools/issues/4151))

### Perf

* optimize route swagger generation

### Refactor

* remove tools dir
* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#189](https://github.com/suyuan32/simple-admin-tools/issues/189) from suyuan32/dev
* Merge pull request [#182](https://github.com/suyuan32/simple-admin-tools/issues/182) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.15.1
* Merge pull request [#185](https://github.com/suyuan32/simple-admin-tools/issues/185) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/jedib0t/go-pretty/v6-6.4.7
* Merge pull request [#184](https://github.com/suyuan32/simple-admin-tools/issues/184) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.5
* Merge pull request [#186](https://github.com/suyuan32/simple-admin-tools/issues/186) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.5
* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5
* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.6.5"></a>

## [tools/goctl/v1.6.5](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.10...tools/goctl/v1.6.5)

> 2024-05-12

### Chore

* refactor mapping unmarshaler ([#4145](https://github.com/suyuan32/simple-admin-tools/issues/4145))
* refactor and coding style ([#4144](https://github.com/suyuan32/simple-admin-tools/issues/4144))
* refactor and coding style ([#4120](https://github.com/suyuan32/simple-admin-tools/issues/4120))
* add more tests ([#4094](https://github.com/suyuan32/simple-admin-tools/issues/4094))
* add code coverage ([#4090](https://github.com/suyuan32/simple-admin-tools/issues/4090))
* coding style ([#4082](https://github.com/suyuan32/simple-admin-tools/issues/4082))

### Feat

* httpx.Parse supports parsing structures that implement the Unmarshaler
  interface ([#4143](https://github.com/suyuan32/simple-admin-tools/issues/4143))
* trigger breaker on underlying service timeout ([#4112](https://github.com/suyuan32/simple-admin-tools/issues/4112))
* add fx.ParallelErr ([#4107](https://github.com/suyuan32/simple-admin-tools/issues/4107))
* add errorx.In to facility error checking ([#4105](https://github.com/suyuan32/simple-admin-tools/issues/4105))
* use breaker with ctx to prevent deadline
  exceeded ([#4091](https://github.com/suyuan32/simple-admin-tools/issues/4091))
* support context in breaker methods ([#4088](https://github.com/suyuan32/simple-admin-tools/issues/4088))
* migrate lua script to lua file ([#4069](https://github.com/suyuan32/simple-admin-tools/issues/4069))

### Fix

* make: command: Command not found ([#4132](https://github.com/suyuan32/simple-admin-tools/issues/4132))
* log panic on Error() or String() panics ([#4136](https://github.com/suyuan32/simple-admin-tools/issues/4136))
* log panic when use nil error or stringer with Field
  method ([#4130](https://github.com/suyuan32/simple-admin-tools/issues/4130))
* zrpc kube resolver builder ([#4119](https://github.com/suyuan32/simple-admin-tools/issues/4119))
* avoid duplicate in logx plain mode ([#4080](https://github.com/suyuan32/simple-admin-tools/issues/4080))
* fix ignored context.DeadlineExceeded ([#4066](https://github.com/suyuan32/simple-admin-tools/issues/4066))

### Optimize

* improve breaker algorithm on recovery time ([#4141](https://github.com/suyuan32/simple-admin-tools/issues/4141))
* improve breaker algorithm on recovery time ([#4077](https://github.com/suyuan32/simple-admin-tools/issues/4077))

<a name="v1.6.10"></a>

## [v1.6.10](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.6.4...v1.6.10)

> 2024-04-11

### Chore

* upgrade deps
* merge latest codes
* upgrade go version
* upgrade deps
* merge latest codes
* upgrade jwt
* upgrade deps
* upgrade google.golang.org/protobuf to v1.33.0
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps and merge latest codes
* upgrade deps
* upgrade deps
* update deps
* update deps
* upgrade grpc to v1.59.0
* upgrade deps
* upgrade deps
* upgrade otlp
* upgrade deps
* upgrade jaeger
* upgrade deps
* upgrade deps
* merge latest codes
* update deps
* merge latest codes
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* move original license to .originallicense dir
* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* add env config for redis
* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* fix ignored scanner.Err() ([#4063](https://github.com/suyuan32/simple-admin-tools/issues/4063))
* timeout on query should return
  context.DeadlineExceeded ([#4060](https://github.com/suyuan32/simple-admin-tools/issues/4060))
* merge latest codes
* upgrade go redis to v9.5.1 to solve the incompatible bug for older version
* disabled identify for redis because in v9.5.0 will break apps
* use new redis init function in test
* request test
* embedded span
* use new version name
* drop tools file
* test files which use jaeger
* add comment for errors in errorx
* many bugs in library
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* remove tools dir
* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#189](https://github.com/suyuan32/simple-admin-tools/issues/189) from suyuan32/dev
* Merge pull request [#182](https://github.com/suyuan32/simple-admin-tools/issues/182) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.15.1
* Merge pull request [#185](https://github.com/suyuan32/simple-admin-tools/issues/185) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/jedib0t/go-pretty/v6-6.4.7
* Merge pull request [#184](https://github.com/suyuan32/simple-admin-tools/issues/184) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.5
* Merge pull request [#186](https://github.com/suyuan32/simple-admin-tools/issues/186) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.5
* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5
* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.6.4"></a>

## [tools/goctl/v1.6.4](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.9...tools/goctl/v1.6.4)

> 2024-04-09

### Chore

* update goctl version to 1.6.4 ([#4057](https://github.com/suyuan32/simple-admin-tools/issues/4057))
* coding style ([#4054](https://github.com/suyuan32/simple-admin-tools/issues/4054))
* fix codecov ([#4044](https://github.com/suyuan32/simple-admin-tools/issues/4044))
* fix codecov ([#4043](https://github.com/suyuan32/simple-admin-tools/issues/4043))
* add more tests ([#4042](https://github.com/suyuan32/simple-admin-tools/issues/4042))
* update codecov config ([#4038](https://github.com/suyuan32/simple-admin-tools/issues/4038))
* coding style ([#4037](https://github.com/suyuan32/simple-admin-tools/issues/4037))
* fix test failure ([#4031](https://github.com/suyuan32/simple-admin-tools/issues/4031))

### Feat

* support http stream response ([#4055](https://github.com/suyuan32/simple-admin-tools/issues/4055))
* add gen api [@doc](https://github.com/doc) comment to logic handler
  routes ([#3790](https://github.com/suyuan32/simple-admin-tools/issues/3790))

### Fix

* bug on form data with slices ([#4040](https://github.com/suyuan32/simple-admin-tools/issues/4040))

<a name="v1.6.9"></a>

## [v1.6.9](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.8...v1.6.9)

> 2024-03-19

### Add

* Optimize the error log to be more specific ([#3994](https://github.com/suyuan32/simple-admin-tools/issues/3994))

### Chore

* upgrade go version
* upgrade deps
* merge latest codes
* upgrade jwt
* upgrade deps
* upgrade google.golang.org/protobuf to v1.33.0
* coding style, add code for prometheus ([#4002](https://github.com/suyuan32/simple-admin-tools/issues/4002))
* fix warnings ([#3990](https://github.com/suyuan32/simple-admin-tools/issues/3990))
* fix warnings ([#3989](https://github.com/suyuan32/simple-admin-tools/issues/3989))
* coding style ([#3987](https://github.com/suyuan32/simple-admin-tools/issues/3987))
* coding style ([#3972](https://github.com/suyuan32/simple-admin-tools/issues/3972))

### Feat

* migrate redis breaker into hook ([#3982](https://github.com/suyuan32/simple-admin-tools/issues/3982))

### Fix

* didn't count failure in allow method with breaker
  algorithm ([#4008](https://github.com/suyuan32/simple-admin-tools/issues/4008))
* StopAgent panics when trace agent disabled ([#3981](https://github.com/suyuan32/simple-admin-tools/issues/3981))
* mysql WithAcceptable bug ([#3986](https://github.com/suyuan32/simple-admin-tools/issues/3986))
* only add log middleware to not found handler when
  enabled ([#3969](https://github.com/suyuan32/simple-admin-tools/issues/3969))

### Optimize

* change err == xx to errors.Is(err, xx) ([#3991](https://github.com/suyuan32/simple-admin-tools/issues/3991))

<a name="v1.6.8"></a>

## [v1.6.8](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.6.3...v1.6.8)

> 2024-03-04

### Chore

* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps and merge latest codes
* upgrade deps
* upgrade deps
* update deps
* update deps
* upgrade grpc to v1.59.0
* upgrade deps
* upgrade deps
* upgrade otlp
* upgrade deps
* upgrade jaeger
* upgrade deps
* upgrade deps
* merge latest codes
* update deps
* merge latest codes
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* move original license to .originallicense dir
* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* merge latest codes
* upgrade go redis to v9.5.1 to solve the incompatible bug for older version
* disabled identify for redis because in v9.5.0 will break apps
* use new redis init function in test
* request test
* embedded span
* use new version name
* drop tools file
* test files which use jaeger
* add comment for errors in errorx
* many bugs in library
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* remove tools dir
* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#189](https://github.com/suyuan32/simple-admin-tools/issues/189) from suyuan32/dev
* Merge pull request [#182](https://github.com/suyuan32/simple-admin-tools/issues/182) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.15.1
* Merge pull request [#185](https://github.com/suyuan32/simple-admin-tools/issues/185) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/jedib0t/go-pretty/v6-6.4.7
* Merge pull request [#184](https://github.com/suyuan32/simple-admin-tools/issues/184) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.5
* Merge pull request [#186](https://github.com/suyuan32/simple-admin-tools/issues/186) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.5
* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5
* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.6.3"></a>

## [tools/goctl/v1.6.3](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.7...tools/goctl/v1.6.3)

> 2024-03-03

### Chore

* update goctl version to 1.6.3 ([#3965](https://github.com/suyuan32/simple-admin-tools/issues/3965))
* coding style ([#3960](https://github.com/suyuan32/simple-admin-tools/issues/3960))
* coding style ([#3959](https://github.com/suyuan32/simple-admin-tools/issues/3959))
* add more tests ([#3958](https://github.com/suyuan32/simple-admin-tools/issues/3958))
* coding style ([#3957](https://github.com/suyuan32/simple-admin-tools/issues/3957))
* add more tests ([#3954](https://github.com/suyuan32/simple-admin-tools/issues/3954))
* simplify linux nocgroup logic ([#3953](https://github.com/suyuan32/simple-admin-tools/issues/3953))
* add lock for batcherror ([#3950](https://github.com/suyuan32/simple-admin-tools/issues/3950))
* simplify linux nocgroup logic ([#3949](https://github.com/suyuan32/simple-admin-tools/issues/3949))
* add tests ([#3931](https://github.com/suyuan32/simple-admin-tools/issues/3931))
* add tests ([#3921](https://github.com/suyuan32/simple-admin-tools/issues/3921))

### Feat

* Improve Docker build ([#3682](https://github.com/suyuan32/simple-admin-tools/issues/3682))
* retry ignore specified errors ([#3808](https://github.com/suyuan32/simple-admin-tools/issues/3808))
* a concurrent runner with messages taken in pushing
  order ([#3941](https://github.com/suyuan32/simple-admin-tools/issues/3941))
* add nocgroup build tag for systems without
  cgroup ([#3948](https://github.com/suyuan32/simple-admin-tools/issues/3948))
* add break metrics for sqlx.statement ([#3947](https://github.com/suyuan32/simple-admin-tools/issues/3947))
* support breaker with sql statements ([#3936](https://github.com/suyuan32/simple-admin-tools/issues/3936))

### Feature

* add a mongo registry option to convert type
  easier. ([#3780](https://github.com/suyuan32/simple-admin-tools/issues/3780))

### Fix

* BatchError.Add() non thread safe ([#3946](https://github.com/suyuan32/simple-admin-tools/issues/3946))
* mapping FillDefault is optional! bug ([#3940](https://github.com/suyuan32/simple-admin-tools/issues/3940))

<a name="v1.6.7"></a>

## [v1.6.7](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.6...v1.6.7)

> 2024-02-21

### Fix

* upgrade go redis to v9.5.1 to solve the incompatible bug for older version

<a name="v1.6.6"></a>

## [v1.6.6](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.5...v1.6.6)

> 2024-02-20

### Chore

* upgrade deps
* optimize coding style and add unit tests ([#3917](https://github.com/suyuan32/simple-admin-tools/issues/3917))

### Feat

* support sse ignore timeout ([#2041](https://github.com/suyuan32/simple-admin-tools/issues/2041))
* optimize circuit breaker algorithm ([#3897](https://github.com/suyuan32/simple-admin-tools/issues/3897))
* support ScheduleImmediately in TaskRunner ([#3896](https://github.com/suyuan32/simple-admin-tools/issues/3896))

### Fix

* disabled identify for redis because in v9.5.0 will break apps
* use new redis init function in test

### Optimize

* improve performance on log disabled ([#3916](https://github.com/suyuan32/simple-admin-tools/issues/3916))
* shedding algorithm performance ([#3908](https://github.com/suyuan32/simple-admin-tools/issues/3908))

<a name="v1.6.5"></a>

## [v1.6.5](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.6.2...v1.6.5)

> 2024-02-04

### Chore

* upgrade deps
* upgrade deps
* upgrade deps
* upgrade deps and merge latest codes
* upgrade deps
* upgrade deps
* update deps
* update deps
* upgrade grpc to v1.59.0
* upgrade deps
* upgrade deps
* upgrade otlp
* upgrade deps
* upgrade jaeger
* upgrade deps
* upgrade deps
* merge latest codes
* update deps
* merge latest codes
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* move original license to .originallicense dir
* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* request test
* embedded span
* use new version name
* drop tools file
* test files which use jaeger
* add comment for errors in errorx
* many bugs in library
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* remove tools dir
* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#189](https://github.com/suyuan32/simple-admin-tools/issues/189) from suyuan32/dev
* Merge pull request [#182](https://github.com/suyuan32/simple-admin-tools/issues/182) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.15.1
* Merge pull request [#185](https://github.com/suyuan32/simple-admin-tools/issues/185) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/jedib0t/go-pretty/v6-6.4.7
* Merge pull request [#184](https://github.com/suyuan32/simple-admin-tools/issues/184) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.5
* Merge pull request [#186](https://github.com/suyuan32/simple-admin-tools/issues/186) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.5
* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5
* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.6.2"></a>

## [tools/goctl/v1.6.2](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.4...tools/goctl/v1.6.2)

> 2024-02-03

### Chore

* update goctl version to 1.6.2 ([#3890](https://github.com/suyuan32/simple-admin-tools/issues/3890))
* add more tests ([#3888](https://github.com/suyuan32/simple-admin-tools/issues/3888))
* fix test failure ([#3883](https://github.com/suyuan32/simple-admin-tools/issues/3883))
* refactor redis ([#3844](https://github.com/suyuan32/simple-admin-tools/issues/3844))
* refactor mapping errors ([#3843](https://github.com/suyuan32/simple-admin-tools/issues/3843))
* make cpu usage more smooth ([#3842](https://github.com/suyuan32/simple-admin-tools/issues/3842))

### Feat

* upgrade go-redis to v9 ([#3088](https://github.com/suyuan32/simple-admin-tools/issues/3088))

### Fix

* cpu stat in cgroup v2 ([#3857](https://github.com/suyuan32/simple-admin-tools/issues/3857))
* revert sqlx metric namespace ([#3847](https://github.com/suyuan32/simple-admin-tools/issues/3847))
*
issue [#3840](https://github.com/suyuan32/simple-admin-tools/issues/3840) ([#3846](https://github.com/suyuan32/simple-admin-tools/issues/3846))
* issue 3840 ([#3845](https://github.com/suyuan32/simple-admin-tools/issues/3845))

<a name="v1.6.4"></a>

## [v1.6.4](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.3...v1.6.4)

> 2024-01-03

### Chore

* upgrade deps
* upgrade deps
* upgrade deps and merge latest codes

### Fix

* embedded span
* negative float32 overflow when unmarshalling ([#3811](https://github.com/suyuan32/simple-admin-tools/issues/3811))

<a name="v1.6.3"></a>

## [v1.6.3](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.6.1...v1.6.3)

> 2023-12-19

### Chore

* upgrade deps
* upgrade deps
* update deps
* update deps
* upgrade grpc to v1.59.0
* upgrade deps
* upgrade deps
* upgrade otlp
* upgrade deps
* upgrade jaeger
* upgrade deps
* upgrade deps
* merge latest codes
* update deps
* merge latest codes
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* move original license to .originallicense dir
* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* use new version name
* drop tools file
* test files which use jaeger
* add comment for errors in errorx
* many bugs in library
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* remove tools dir
* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#189](https://github.com/suyuan32/simple-admin-tools/issues/189) from suyuan32/dev
* Merge pull request [#182](https://github.com/suyuan32/simple-admin-tools/issues/182) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.15.1
* Merge pull request [#185](https://github.com/suyuan32/simple-admin-tools/issues/185) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/jedib0t/go-pretty/v6-6.4.7
* Merge pull request [#184](https://github.com/suyuan32/simple-admin-tools/issues/184) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.5
* Merge pull request [#186](https://github.com/suyuan32/simple-admin-tools/issues/186) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.5
* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5
* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.6.1"></a>

## [tools/goctl/v1.6.1](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.2...tools/goctl/v1.6.1)

> 2023-12-17

### Chore

* update goctl deps ([#3797](https://github.com/suyuan32/simple-admin-tools/issues/3797))
* coding style ([#3796](https://github.com/suyuan32/simple-admin-tools/issues/3796))
* upgrade otel, removed ut temporarily because of otel API
  changes ([#3795](https://github.com/suyuan32/simple-admin-tools/issues/3795))

<a name="v1.6.2"></a>

## [v1.6.2](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.1...v1.6.2)

> 2023-12-15

### Fix

* use new version name

<a name="v1.6.1"></a>

## [v1.6.1](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.0...v1.6.1)

> 2023-12-15

### Chore

* refactor ring ([#3739](https://github.com/suyuan32/simple-admin-tools/issues/3739))
* update mongo driver ([#3727](https://github.com/suyuan32/simple-admin-tools/issues/3727))
* refact dart code generation ([#3726](https://github.com/suyuan32/simple-admin-tools/issues/3726))

### Feat

* add dbtest to facility db test ([#3768](https://github.com/suyuan32/simple-admin-tools/issues/3768))
* auto stop profiling after one minute ([#3742](https://github.com/suyuan32/simple-admin-tools/issues/3742))

### Fix

* drop tools file
* coredump: goctl model mysql ddl --src user_base.sql --dir .
  area… ([#3777](https://github.com/suyuan32/simple-admin-tools/issues/3777))
* primary key unique key simultaneously exist cacheIdPrefix
  duplicate ([#3763](https://github.com/suyuan32/simple-admin-tools/issues/3763))
* prevent a crash if there is a unique key constraint with a nil
  field. ([#3770](https://github.com/suyuan32/simple-admin-tools/issues/3770))
* prevent ring index overflow ([#3738](https://github.com/suyuan32/simple-admin-tools/issues/3738))
* goctl FindOne error ([#3731](https://github.com/suyuan32/simple-admin-tools/issues/3731))

<a name="v1.6.0"></a>

## [v1.6.0](https://github.com/suyuan32/simple-admin-tools/compare/v1.6.0-beta...v1.6.0)

> 2023-10-29

### Chore

* update deps

<a name="v1.6.0-beta"></a>

## [v1.6.0-beta](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.6.0...v1.6.0-beta)

> 2023-10-29

### Chore

* update deps
* upgrade grpc to v1.59.0
* upgrade deps
* upgrade deps
* upgrade otlp
* upgrade deps
* upgrade jaeger
* upgrade deps
* upgrade deps
* merge latest codes
* update deps
* merge latest codes
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* move original license to .originallicense dir
* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* test files which use jaeger
* add comment for errors in errorx
* many bugs in library
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* remove tools dir
* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#189](https://github.com/suyuan32/simple-admin-tools/issues/189) from suyuan32/dev
* Merge pull request [#182](https://github.com/suyuan32/simple-admin-tools/issues/182) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.15.1
* Merge pull request [#185](https://github.com/suyuan32/simple-admin-tools/issues/185) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/jedib0t/go-pretty/v6-6.4.7
* Merge pull request [#184](https://github.com/suyuan32/simple-admin-tools/issues/184) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.5
* Merge pull request [#186](https://github.com/suyuan32/simple-admin-tools/issues/186) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.5
* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5
* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.6.0"></a>

## [tools/goctl/v1.6.0](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.19-beta...tools/goctl/v1.6.0)

> 2023-10-28

### Chore

* update go-zero version to v1.6.0 in goctl ([#3679](https://github.com/suyuan32/simple-admin-tools/issues/3679))
* remove go build version in fuzz test ([#3678](https://github.com/suyuan32/simple-admin-tools/issues/3678))
* upgrade go to 1.19 ([#3677](https://github.com/suyuan32/simple-admin-tools/issues/3677))
* upgrade go dependencies ([#3657](https://github.com/suyuan32/simple-admin-tools/issues/3657))
* update goctl version to 1.6.0 ([#3674](https://github.com/suyuan32/simple-admin-tools/issues/3674))
* simplify prometheus check ([#3672](https://github.com/suyuan32/simple-admin-tools/issues/3672))
* refactor zrpc timeout ([#3671](https://github.com/suyuan32/simple-admin-tools/issues/3671))

### Feat

* add metrics ([#3624](https://github.com/suyuan32/simple-admin-tools/issues/3624))
* support the specified timeout of rpc methods ([#2742](https://github.com/suyuan32/simple-admin-tools/issues/2742))

### Fix

* adjust log encode output mode ([#3676](https://github.com/suyuan32/simple-admin-tools/issues/3676))

<a name="v1.5.19-beta"></a>

## [v1.5.19-beta](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.18...v1.5.19-beta)

> 2023-10-25

### Chore

* upgrade grpc to v1.59.0
* add error handling on registering event handlers to
  k8s ([#3663](https://github.com/suyuan32/simple-admin-tools/issues/3663))
* update restful/grpc servers shutdown stages ([#3662](https://github.com/suyuan32/simple-admin-tools/issues/3662))
* refactor mongo logs ([#3660](https://github.com/suyuan32/simple-admin-tools/issues/3660))
* add go-zero users ([#3659](https://github.com/suyuan32/simple-admin-tools/issues/3659))
* refactor errors to use errors.Is ([#3654](https://github.com/suyuan32/simple-admin-tools/issues/3654))
* refactor errors ([#3651](https://github.com/suyuan32/simple-admin-tools/issues/3651))
* upgrade go to 1.19 ([#3648](https://github.com/suyuan32/simple-admin-tools/issues/3648))

### Feat

* model add withSession ([#3658](https://github.com/suyuan32/simple-admin-tools/issues/3658))
* support for disable mon logs like sqlx ([#3606](https://github.com/suyuan32/simple-admin-tools/issues/3606))
* optimize logx print error ([#3649](https://github.com/suyuan32/simple-admin-tools/issues/3649))
* export devserver.Config ([#3638](https://github.com/suyuan32/simple-admin-tools/issues/3638))

### Fix

* multiple files import the same api file ([#3642](https://github.com/suyuan32/simple-admin-tools/issues/3642))
* goctl unit test ([#3636](https://github.com/suyuan32/simple-admin-tools/issues/3636))

### Refactor

* remove tools dir

<a name="v1.5.18"></a>

## [v1.5.18](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.17...v1.5.18)

> 2023-10-17

### Chore

* upgrade deps
* upgrade deps
* update readme for Mac install instructions ([#3633](https://github.com/suyuan32/simple-admin-tools/issues/3633))
* refactor signal sigterm and sigint ([#3632](https://github.com/suyuan32/simple-admin-tools/issues/3632))
* refactor error comparison ([#3629](https://github.com/suyuan32/simple-admin-tools/issues/3629))
* simplify parsing numbers with overflow ([#3610](https://github.com/suyuan32/simple-admin-tools/issues/3610))
* refactor goctl api ([#3605](https://github.com/suyuan32/simple-admin-tools/issues/3605))
* rename noOpBreaker to nopBreaker ([#3602](https://github.com/suyuan32/simple-admin-tools/issues/3602))

### Feat

* add the SIGINT signal in signals.go to subscribe the user input ctrl+c to exit the application
  operation ([#3611](https://github.com/suyuan32/simple-admin-tools/issues/3611))
* retry with ctx deadline ([#3626](https://github.com/suyuan32/simple-admin-tools/issues/3626))

### Fix

* optimize logx for less GC objects ([#3627](https://github.com/suyuan32/simple-admin-tools/issues/3627))
* unmarshal from number to string with incorrect error
  message ([#3625](https://github.com/suyuan32/simple-admin-tools/issues/3625))
* UpdateStmt doesn't update the statement correctly in
  sqlx/bulkinserter.go ([#3607](https://github.com/suyuan32/simple-admin-tools/issues/3607))

### Optimize

* fix experimental api  ([#3604](https://github.com/suyuan32/simple-admin-tools/issues/3604))

### Refactor

* update builder.go ([#3620](https://github.com/suyuan32/simple-admin-tools/issues/3620))

<a name="v1.5.17"></a>

## [v1.5.17](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.5.6...v1.5.17)

> 2023-09-29

### Chore

* upgrade otlp
* upgrade deps
* upgrade jaeger
* upgrade deps
* upgrade deps
* merge latest codes
* update deps
* merge latest codes
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* move original license to .originallicense dir
* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* test files which use jaeger
* add comment for errors in errorx
* many bugs in library
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#189](https://github.com/suyuan32/simple-admin-tools/issues/189) from suyuan32/dev
* Merge pull request [#182](https://github.com/suyuan32/simple-admin-tools/issues/182) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.15.1
* Merge pull request [#185](https://github.com/suyuan32/simple-admin-tools/issues/185) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/jedib0t/go-pretty/v6-6.4.7
* Merge pull request [#184](https://github.com/suyuan32/simple-admin-tools/issues/184) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.5
* Merge pull request [#186](https://github.com/suyuan32/simple-admin-tools/issues/186) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.5
* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5
* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.5.6"></a>

## [tools/goctl/v1.5.6](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.16...tools/goctl/v1.5.6)

> 2023-09-27

### Chore

* update goctl version to 1.5.6 ([#3593](https://github.com/suyuan32/simple-admin-tools/issues/3593))
* add more tests ([#3592](https://github.com/suyuan32/simple-admin-tools/issues/3592))
* add test for logging rotate size ([#3587](https://github.com/suyuan32/simple-admin-tools/issues/3587))
* add more tests ([#3578](https://github.com/suyuan32/simple-admin-tools/issues/3578))
* add more tests ([#3577](https://github.com/suyuan32/simple-admin-tools/issues/3577))
* add more tests ([#3575](https://github.com/suyuan32/simple-admin-tools/issues/3575))
* refactor log limit in rest ([#3572](https://github.com/suyuan32/simple-admin-tools/issues/3572))
* upgrade grpc ([#3570](https://github.com/suyuan32/simple-admin-tools/issues/3570))
* upgrade grpc ([#3558](https://github.com/suyuan32/simple-admin-tools/issues/3558))
* refactor ([#3545](https://github.com/suyuan32/simple-admin-tools/issues/3545))

### Feat

* Optimize Encoding Functions and Add Descriptive
  Comments ([#3543](https://github.com/suyuan32/simple-admin-tools/issues/3543))

### Fix

* avoid float overflow in mapping.Unmarshal ([#3590](https://github.com/suyuan32/simple-admin-tools/issues/3590))
* avoid integer overflow in mapping.Unmarshal ([#3582](https://github.com/suyuan32/simple-admin-tools/issues/3582))
* avoid losing logs before closing ([#3573](https://github.com/suyuan32/simple-admin-tools/issues/3573))
* goctl pg gen will extract all fields when the same table name exists in different
  schemas ([#3496](https://github.com/suyuan32/simple-admin-tools/issues/3496)) ([#3517](https://github.com/suyuan32/simple-admin-tools/issues/3517))

### Update

* limit logBrief http body size ([#3498](https://github.com/suyuan32/simple-admin-tools/issues/3498))

<a name="v1.5.16"></a>

## [v1.5.16](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.5.5...v1.5.16)

> 2023-09-03

### Chore

* upgrade deps
* upgrade deps
* merge latest codes
* update deps
* merge latest codes
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* move original license to .originallicense dir
* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* many bugs in library
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#189](https://github.com/suyuan32/simple-admin-tools/issues/189) from suyuan32/dev
* Merge pull request [#182](https://github.com/suyuan32/simple-admin-tools/issues/182) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.15.1
* Merge pull request [#185](https://github.com/suyuan32/simple-admin-tools/issues/185) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/jedib0t/go-pretty/v6-6.4.7
* Merge pull request [#184](https://github.com/suyuan32/simple-admin-tools/issues/184) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.5
* Merge pull request [#186](https://github.com/suyuan32/simple-admin-tools/issues/186) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.5
* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5
* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.5.5"></a>

## [tools/goctl/v1.5.5](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.15...tools/goctl/v1.5.5)

> 2023-09-02

### Chore

* upgrade go-zero, and update goctl version ([#3509](https://github.com/suyuan32/simple-admin-tools/issues/3509))

### Feat

* support json:"-" in mapping ([#3521](https://github.com/suyuan32/simple-admin-tools/issues/3521))

<a name="v1.5.15"></a>

## [v1.5.15](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.14...v1.5.15)

> 2023-08-18

### Chore

* upgrade deps
* merge latest codes
* use jsonTagKey to replace json literals ([#3479](https://github.com/suyuan32/simple-admin-tools/issues/3479))

### Fix

* many bugs in library
* [#3478](https://github.com/suyuan32/simple-admin-tools/issues/3478) ([#3493](https://github.com/suyuan32/simple-admin-tools/issues/3493))

### Pull Requests

* Merge pull request [#180](https://github.com/suyuan32/simple-admin-tools/issues/180) from suyuan32/dev
* Merge pull request [#178](https://github.com/suyuan32/simple-admin-tools/issues/178) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.12.0
* Merge pull request [#177](https://github.com/suyuan32/simple-admin-tools/issues/177) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/duke-git/lancet/v2-2.2.4
* Merge pull request [#179](https://github.com/suyuan32/simple-admin-tools/issues/179) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.5

<a name="v1.5.14"></a>

## [v1.5.14](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.5.4...v1.5.14)

> 2023-08-04

### Chore

* update deps
* merge latest codes
* fix potential nil pointer errors ([#3454](https://github.com/suyuan32/simple-admin-tools/issues/3454))
* add more tests ([#3453](https://github.com/suyuan32/simple-admin-tools/issues/3453))
* upgrade deps
* upgrade deps
* update goctls version
* merge latest codes
* update default tool version
* upgrade deps
* merge latest codes
* upgrade version
* upgrade deps
* merge latest codes
* update default tools version
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* update readme
* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* add prometheus summary metrics ([#3440](https://github.com/suyuan32/simple-admin-tools/issues/3440))
* optimize mapping error ([#3438](https://github.com/suyuan32/simple-admin-tools/issues/3438))
* proto2api command
* add mcms rpc port info
* service list in info cmd
* mcms env info
* add client trans
* api cmd support extra fields
* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl
* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* "EXPRIMENTAL" is a misspelling of "EXPERIMENTAL" ([#3462](https://github.com/suyuan32/simple-admin-tools/issues/3462))
* remove mapping redundant error ([#3439](https://github.com/suyuan32/simple-admin-tools/issues/3439))
* format error should not trigger circuit breaker in
  sqlx ([#3437](https://github.com/suyuan32/simple-admin-tools/issues/3437))
* use sub folder in api desc
* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* wrong i18n translate
* use default timeout in api etc
* optimize makefile tpl
* add default ent feature
* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl
* optimize rpc set logic
* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* recover deps
* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* localmixin
* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#169](https://github.com/suyuan32/simple-admin-tools/issues/169) from suyuan32/dev
* Merge pull request [#159](https://github.com/suyuan32/simple-admin-tools/issues/159) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.9
* Merge pull request [#161](https://github.com/suyuan32/simple-admin-tools/issues/161) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/iancoleman/strcase-0.3.0
* Merge pull request [#167](https://github.com/suyuan32/simple-admin-tools/issues/167) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.12.1
* Merge pull request [#168](https://github.com/suyuan32/simple-admin-tools/issues/168) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/gookit/color-1.5.4
* Merge pull request [#166](https://github.com/suyuan32/simple-admin-tools/issues/166) from suyuan32/dev
* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev
* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev
* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev
* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev
* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev
* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7
* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.5.4"></a>

## [tools/goctl/v1.5.4](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.13...tools/goctl/v1.5.4)

> 2023-07-14

### Chore

* update go-zero for goctl ([#3426](https://github.com/suyuan32/simple-admin-tools/issues/3426))

<a name="v1.5.13"></a>

## [v1.5.13](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.12...v1.5.13)

> 2023-07-14

### Chore

* update goctls version
* merge latest codes
* make servicegroup panic as demand ([#3422](https://github.com/suyuan32/simple-admin-tools/issues/3422))
* coding style ([#3413](https://github.com/suyuan32/simple-admin-tools/issues/3413))
* make test stable ([#3412](https://github.com/suyuan32/simple-admin-tools/issues/3412))
* avoid nested WithCodeResponseWriter ([#3406](https://github.com/suyuan32/simple-admin-tools/issues/3406))
* slice replace loop ([#3410](https://github.com/suyuan32/simple-admin-tools/issues/3410))

### Feat

* add detail type mismatch info in number fields
  check ([#3386](https://github.com/suyuan32/simple-admin-tools/issues/3386)) ([#3387](https://github.com/suyuan32/simple-admin-tools/issues/3387))
* add mcms rpc port info

### Fix

* remove unused mcms env info
* bug in frontend en locale overwrite
* optimize vben component type checker
* log format error ([#3409](https://github.com/suyuan32/simple-admin-tools/issues/3409))
* wrong i18n translate

### Pull Requests

* Merge pull request [#158](https://github.com/suyuan32/simple-admin-tools/issues/158) from suyuan32/dev
* Merge pull request [#156](https://github.com/suyuan32/simple-admin-tools/issues/156) from suyuan32/dev
* Merge pull request [#155](https://github.com/suyuan32/simple-admin-tools/issues/155) from suyuan32/dev
* Merge pull request [#154](https://github.com/suyuan32/simple-admin-tools/issues/154) from suyuan32/dev
* Merge pull request [#153](https://github.com/suyuan32/simple-admin-tools/issues/153) from suyuan32/dev
* Merge pull request [#151](https://github.com/suyuan32/simple-admin-tools/issues/151) from suyuan32/dev

<a name="v1.5.12"></a>

## [v1.5.12](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.11...v1.5.12)

> 2023-07-09

### Chore

* update default tool version
* upgrade deps
* refactor httpx.TimeoutHandler ([#3400](https://github.com/suyuan32/simple-admin-tools/issues/3400))
* merge latest codes

### Feat

* service list in info cmd
* mcms env info

### Fix

* use default timeout in api etc
* optimize makefile tpl
* add default ent feature

### Pull Requests

* Merge pull request [#150](https://github.com/suyuan32/simple-admin-tools/issues/150) from suyuan32/dev
* Merge pull request [#149](https://github.com/suyuan32/simple-admin-tools/issues/149) from suyuan32/dev
* Merge pull request [#148](https://github.com/suyuan32/simple-admin-tools/issues/148) from suyuan32/dev
* Merge pull request [#147](https://github.com/suyuan32/simple-admin-tools/issues/147) from suyuan32/dev

<a name="v1.5.11"></a>

## [v1.5.11](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.10...v1.5.11)

> 2023-07-03

### Chore

* upgrade version
* upgrade deps
* merge latest codes
* refactor httpx.SetOkHandler ([#3373](https://github.com/suyuan32/simple-admin-tools/issues/3373))
* coding style ([#3362](https://github.com/suyuan32/simple-admin-tools/issues/3362))
* add more tests ([#3359](https://github.com/suyuan32/simple-admin-tools/issues/3359))

### Feat

* add client trans
* api cmd support extra fields
* optimize must log add stack ([#3384](https://github.com/suyuan32/simple-admin-tools/issues/3384))

### Fix

* deps bug
* rest timeout test
* optimize other init code gen
* add pointer support for front end
* NewClientWithTarget miss default config ([#3358](https://github.com/suyuan32/simple-admin-tools/issues/3358))

### Opt

* improve logx performance ([#3371](https://github.com/suyuan32/simple-admin-tools/issues/3371))

### Revert

* recover deps

### Pull Requests

* Merge pull request [#143](https://github.com/suyuan32/simple-admin-tools/issues/143) from suyuan32/dev
* Merge pull request [#142](https://github.com/suyuan32/simple-admin-tools/issues/142) from suyuan32/dev
* Merge pull request [#141](https://github.com/suyuan32/simple-admin-tools/issues/141) from suyuan32/dev
* Merge pull request [#137](https://github.com/suyuan32/simple-admin-tools/issues/137) from suyuan32/dev

<a name="v1.5.10"></a>

## [v1.5.10](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.10-rc1...v1.5.10)

> 2023-06-20

### Chore

* update default tools version

### Feat

* add ent feature detect for makefile gen cmd
* add CROS info
* add CROSConf to tpl

### Fix

* remove dot when there are not normal fields
* use latest golang image in docker cmd
* add i18n val auto detect for makefile cmd
* remove space in makefile vals
* use button in components instead antd in vben
* optimize api handler and svc tpl

### Pull Requests

* Merge pull request [#136](https://github.com/suyuan32/simple-admin-tools/issues/136) from suyuan32/dev
* Merge pull request [#129](https://github.com/suyuan32/simple-admin-tools/issues/129) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.10.0
* Merge pull request [#130](https://github.com/suyuan32/simple-admin-tools/issues/130) from
  suyuan32/dependabot/go_modules/golang.org/x/sys-0.9.0
* Merge pull request [#135](https://github.com/suyuan32/simple-admin-tools/issues/135) from suyuan32/dev

<a name="v1.5.10-rc1"></a>

## [v1.5.10-rc1](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.9...v1.5.10-rc1)

> 2023-06-16

### Fix

* optimize rpc set logic

### Pull Requests

* Merge pull request [#128](https://github.com/suyuan32/simple-admin-tools/issues/128) from suyuan32/dev

<a name="v1.5.9"></a>

## [v1.5.9](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.9-beta...v1.5.9)

> 2023-06-16

### Chore

* refactor and add more tests ([#3351](https://github.com/suyuan32/simple-admin-tools/issues/3351))
* more tests ([#3340](https://github.com/suyuan32/simple-admin-tools/issues/3340))
* add more tests ([#3338](https://github.com/suyuan32/simple-admin-tools/issues/3338))
* improve logx gzip ([#3332](https://github.com/suyuan32/simple-admin-tools/issues/3332))

### Feat

* merge latest codes
* api ent cmd support pointer
* rpc gen support set not nil
* httpx add common handler ([#3269](https://github.com/suyuan32/simple-admin-tools/issues/3269))

### Fix

* update default goctls version
* default ent template when create project
* adjust api proto command to pointer
* remove unnecessary lang support
* set page params to non-optional
* fixed goctl api go --home parameter error when loading
  non-exist… ([#3319](https://github.com/suyuan32/simple-admin-tools/issues/3319))

### Pull Requests

* Merge pull request [#127](https://github.com/suyuan32/simple-admin-tools/issues/127) from suyuan32/dev
* Merge pull request [#126](https://github.com/suyuan32/simple-admin-tools/issues/126) from suyuan32/dev
* Merge pull request [#124](https://github.com/suyuan32/simple-admin-tools/issues/124) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.7

<a name="v1.5.9-beta"></a>

## [v1.5.9-beta](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.5.3...v1.5.9-beta)

> 2023-06-09

### Chore

* add more tests ([#3324](https://github.com/suyuan32/simple-admin-tools/issues/3324))
* update readme to remove upgrade parts. ([#3318](https://github.com/suyuan32/simple-admin-tools/issues/3318))
* make test stable ([#3317](https://github.com/suyuan32/simple-admin-tools/issues/3317))
* update default version of tools
* merge latest codes
* merge latest codes
* merge latest codes
* update default tool version for code gen
* update default version
* merge latest codes of go zero
* update version
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* add ent feature configuration for makefile
* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api
* add author params for docker command
* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* merge latest codes
* merge latest codes
* optimize PageInfo in base api template
* optimize PageInfo in base api template
* optimize the parse error
* optimize ent template list
* update default tools version in goctls
* error when makefile not exist
* add import prefix params for logic code gen
* add sqlx support
* add makefile style param
* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro
* remove gopakage judgement in proto parser
* group name invalid bug in single api code generation
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command
* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Fix

* logx with Compress auto delete old logs ([#3329](https://github.com/suyuan32/simple-admin-tools/issues/3329))

### Perf

* optimize route swagger generation

### Refactor

* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#123](https://github.com/suyuan32/simple-admin-tools/issues/123) from suyuan32/dev
* Merge pull request [#115](https://github.com/suyuan32/simple-admin-tools/issues/115) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.1
* Merge pull request [#112](https://github.com/suyuan32/simple-admin-tools/issues/112) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.1
* Merge pull request [#114](https://github.com/suyuan32/simple-admin-tools/issues/114) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.4
* Merge pull request [#113](https://github.com/suyuan32/simple-admin-tools/issues/113) from
  suyuan32/dependabot/go_modules/github.com/pelletier/go-toml/v2-2.0.8
* Merge pull request [#116](https://github.com/suyuan32/simple-admin-tools/issues/116) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.3
* Merge pull request [#117](https://github.com/suyuan32/simple-admin-tools/issues/117) from
  suyuan32/dependabot/go_modules/github.com/alicebob/miniredis/v2-2.30.3
* Merge pull request [#118](https://github.com/suyuan32/simple-admin-tools/issues/118) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.4
* Merge pull request [#120](https://github.com/suyuan32/simple-admin-tools/issues/120) from Jetereting/fix/page0sqlError
* Merge pull request [#119](https://github.com/suyuan32/simple-admin-tools/issues/119) from Jetereting/master
* Merge pull request [#109](https://github.com/suyuan32/simple-admin-tools/issues/109) from Jetereting/master
* Merge pull request [#108](https://github.com/suyuan32/simple-admin-tools/issues/108) from Jetereting/master
* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev
* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev
* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev
* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0
* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev
* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.5.3"></a>

## [tools/goctl/v1.5.3](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.8...tools/goctl/v1.5.3)

> 2023-06-04

### Chore

* update go-zero for goctl ([#3316](https://github.com/suyuan32/simple-admin-tools/issues/3316))
* add more tests ([#3315](https://github.com/suyuan32/simple-admin-tools/issues/3315))

<a name="v1.5.8"></a>

## [v1.5.8](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.7...v1.5.8)

> 2023-05-31

### Chore

* update default version of tools
* merge latest codes
* merge latest codes
* add more tests ([#3304](https://github.com/suyuan32/simple-admin-tools/issues/3304))
* add more tests ([#3299](https://github.com/suyuan32/simple-admin-tools/issues/3299))
* add more tests ([#3296](https://github.com/suyuan32/simple-admin-tools/issues/3296))
* refactor ([#3295](https://github.com/suyuan32/simple-admin-tools/issues/3295))
* add more tests ([#3294](https://github.com/suyuan32/simple-admin-tools/issues/3294))
* refactor retry ([#3291](https://github.com/suyuan32/simple-admin-tools/issues/3291))
* add more tests ([#3290](https://github.com/suyuan32/simple-admin-tools/issues/3290))
* add more tests ([#3288](https://github.com/suyuan32/simple-admin-tools/issues/3288))
* add more tests ([#3286](https://github.com/suyuan32/simple-admin-tools/issues/3286))
* add more tests ([#3282](https://github.com/suyuan32/simple-admin-tools/issues/3282))
* add more tests ([#3279](https://github.com/suyuan32/simple-admin-tools/issues/3279))

### Feat

* add i18n info for goctls
* add more upper case checker for ent
* add service name params for info env cmd
* support optional otel global initialization
  for [#3284](https://github.com/suyuan32/simple-admin-tools/issues/3284) ([#3292](https://github.com/suyuan32/simple-admin-tools/issues/3292))
* sqlc add SetCacheWithExpire method ([#3249](https://github.com/suyuan32/simple-admin-tools/issues/3249))
* add OtlpHttpPath config support for ZincObserve
  Telemetry ([#3271](https://github.com/suyuan32/simple-admin-tools/issues/3271))

### Fix

* optimize the parse error
* optimize ent template list
* update default tools version in goctls

### Update

* expand the retry method to support timeout and interval
  control ([#3283](https://github.com/suyuan32/simple-admin-tools/issues/3283))

### Pull Requests

* Merge pull request [#107](https://github.com/suyuan32/simple-admin-tools/issues/107) from suyuan32/dev
* Merge pull request [#104](https://github.com/suyuan32/simple-admin-tools/issues/104) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.16.0
* Merge pull request [#105](https://github.com/suyuan32/simple-admin-tools/issues/105) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/sdk-1.16.0
* Merge pull request [#106](https://github.com/suyuan32/simple-admin-tools/issues/106) from suyuan32/dev
* Merge pull request [#102](https://github.com/suyuan32/simple-admin-tools/issues/102) from suyuan32/dev
* Merge pull request [#101](https://github.com/suyuan32/simple-admin-tools/issues/101) from suyuan32/dev

<a name="v1.5.7"></a>

## [v1.5.7](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.7-beta.1...v1.5.7)

> 2023-05-23

### Chore

* merge latest codes
* add more tests ([#3265](https://github.com/suyuan32/simple-admin-tools/issues/3265))
* add more tests ([#3261](https://github.com/suyuan32/simple-admin-tools/issues/3261))
* add more tests ([#3260](https://github.com/suyuan32/simple-admin-tools/issues/3260))
* add more tests ([#3259](https://github.com/suyuan32/simple-admin-tools/issues/3259))
* add more tests ([#3258](https://github.com/suyuan32/simple-admin-tools/issues/3258))
* remove directive for tests ([#3257](https://github.com/suyuan32/simple-admin-tools/issues/3257))
* add more tests ([#3256](https://github.com/suyuan32/simple-admin-tools/issues/3256))

### Feat

* support using session to execute statements in
  transaction ([#3252](https://github.com/suyuan32/simple-admin-tools/issues/3252))

### Fix

* error when makefile not exist
* add import prefix params for logic code gen

### Pull Requests

* Merge pull request [#100](https://github.com/suyuan32/simple-admin-tools/issues/100) from suyuan32/dev
* Merge pull request [#94](https://github.com/suyuan32/simple-admin-tools/issues/94) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.1
* Merge pull request [#93](https://github.com/suyuan32/simple-admin-tools/issues/93) from
  suyuan32/dependabot/go_modules/github.com/stretchr/testify-1.8.3
* Merge pull request [#95](https://github.com/suyuan32/simple-admin-tools/issues/95) from
  suyuan32/dependabot/go_modules/github.com/go-playground/validator/v10-10.14.0
* Merge pull request [#98](https://github.com/suyuan32/simple-admin-tools/issues/98) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/stretchr/testify-1.8.3
* Merge pull request [#92](https://github.com/suyuan32/simple-admin-tools/issues/92) from suyuan32/dev
* Merge pull request [#91](https://github.com/suyuan32/simple-admin-tools/issues/91) from suyuan32/dev

<a name="v1.5.7-beta.1"></a>

## [v1.5.7-beta.1](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.6...v1.5.7-beta.1)

> 2023-05-16

### Fix

* add sqlx support
* add makefile style param

### Pull Requests

* Merge pull request [#89](https://github.com/suyuan32/simple-admin-tools/issues/89) from suyuan32/dev
* Merge pull request [#88](https://github.com/suyuan32/simple-admin-tools/issues/88) from suyuan32/dev

<a name="v1.5.6"></a>

## [v1.5.6](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.5...v1.5.6)

> 2023-05-15

### Chore

* update default tool version for code gen
* format the code ([#3243](https://github.com/suyuan32/simple-admin-tools/issues/3243))
* better comments ([#3232](https://github.com/suyuan32/simple-admin-tools/issues/3232))

### Feat

* makefile gen command
* i18n support for rpc ent logic gen
* i18n support for api ent and proto gen
* add info command
* chinese support for tools
* add environment variables support for log,rpc and api

### Fix

* makefile style param
* makefile gen rpc bug
* merge latest codes
* chinese support for ent template help message
* add goctls docker command author params intro

### Pull Requests

* Merge pull request [#87](https://github.com/suyuan32/simple-admin-tools/issues/87) from suyuan32/dev
* Merge pull request [#81](https://github.com/suyuan32/simple-admin-tools/issues/81) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.9
* Merge pull request [#82](https://github.com/suyuan32/simple-admin-tools/issues/82) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.1
* Merge pull request [#83](https://github.com/suyuan32/simple-admin-tools/issues/83) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.9
* Merge pull request [#84](https://github.com/suyuan32/simple-admin-tools/issues/84) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.1
* Merge pull request [#85](https://github.com/suyuan32/simple-admin-tools/issues/85) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.10.0

<a name="v1.5.5"></a>

## [v1.5.5](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.4...v1.5.5)

> 2023-05-09

### Chore

* update default version
* merge latest codes of go zero
* add more tests ([#3229](https://github.com/suyuan32/simple-admin-tools/issues/3229))
* remove fgprof, use pprof directly ([#3220](https://github.com/suyuan32/simple-admin-tools/issues/3220))
* optimize stat calculation ([#3213](https://github.com/suyuan32/simple-admin-tools/issues/3213))
* add more tests ([#3203](https://github.com/suyuan32/simple-admin-tools/issues/3203))

### Feat

* add otlptracegrpc otlptracehttp headers support for
  Uptrace ([#3219](https://github.com/suyuan32/simple-admin-tools/issues/3219))
* replaced color package to support Windows ([#3207](https://github.com/suyuan32/simple-admin-tools/issues/3207))
* add author params for docker command

### Fix

* remove gopakage judgement in proto parser
* generate client directory for goctl ([#3166](https://github.com/suyuan32/simple-admin-tools/issues/3166))
* timeout handler not implementing http.Flusher ([#3225](https://github.com/suyuan32/simple-admin-tools/issues/3225))
* Errorv should generate JSON Object for content field in
  log ([#3222](https://github.com/suyuan32/simple-admin-tools/issues/3222))
* group name invalid bug in single api code generation
* remove etcd pings to avoid too-many-pings error ([#3212](https://github.com/suyuan32/simple-admin-tools/issues/3212))
* optimize docker command
* remove unnecessary dockerfile gen codes
* update docker command

### Pull Requests

* Merge pull request [#80](https://github.com/suyuan32/simple-admin-tools/issues/80) from suyuan32/dev
* Merge pull request [#75](https://github.com/suyuan32/simple-admin-tools/issues/75) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp-1.15.1
* Merge pull request [#76](https://github.com/suyuan32/simple-admin-tools/issues/76) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.1
* Merge pull request [#77](https://github.com/suyuan32/simple-admin-tools/issues/77) from
  suyuan32/dependabot/go_modules/google.golang.org/grpc-1.55.0
* Merge pull request [#78](https://github.com/suyuan32/simple-admin-tools/issues/78) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/trace-1.15.1
* Merge pull request [#79](https://github.com/suyuan32/simple-admin-tools/issues/79) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.6
* Merge pull request [#74](https://github.com/suyuan32/simple-admin-tools/issues/74) from suyuan32/dev
* Merge pull request [#73](https://github.com/suyuan32/simple-admin-tools/issues/73) from suyuan32/dev
* Merge pull request [#72](https://github.com/suyuan32/simple-admin-tools/issues/72) from suyuan32/dev

<a name="v1.5.4"></a>

## [v1.5.4](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.5.2...v1.5.4)

> 2023-05-02

### Chore

* update version
* change port to 6060 by default in devserver ([#3191](https://github.com/suyuan32/simple-admin-tools/issues/3191))
* update readme ([#3190](https://github.com/suyuan32/simple-admin-tools/issues/3190))
* use logx.Must instead of log.Fatal ([#3189](https://github.com/suyuan32/simple-admin-tools/issues/3189))
* add more tests ([#3187](https://github.com/suyuan32/simple-admin-tools/issues/3187))
* simplify tests with logtest ([#3184](https://github.com/suyuan32/simple-admin-tools/issues/3184))
* update default tool version to 1.5.3
* update goctls version
* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* update changelog
* add copyright comment
* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* supported enum.
* set error code function for validator
* add drone command
* add validation register function
* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command
* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* remove unused vars
* merge latest codes
* tests bugs
* optimize imports in statusrender.tpl
* optimize imports in data.tpl
* remove redundant message alert
* frontend locale overwrite bug
* frontend locale overwrite bug
* adjust vben template to suit successMessageMode
* merge latest codes
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version
* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package
* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile
* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* optimize gen rpc comment
* optimize gen-rpc help message
* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#70](https://github.com/suyuan32/simple-admin-tools/issues/70) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/emicklei/proto-1.11.2
* Merge pull request [#71](https://github.com/suyuan32/simple-admin-tools/issues/71) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.3
* Merge pull request [#63](https://github.com/suyuan32/simple-admin-tools/issues/63) from
  suyuan32/dependabot/go_modules/k8s.io/apimachinery-0.28.0-alpha.0
* Merge pull request [#64](https://github.com/suyuan32/simple-admin-tools/issues/64) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.2
* Merge pull request [#66](https://github.com/suyuan32/simple-admin-tools/issues/66) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/zipkin-1.15.0
* Merge pull request [#68](https://github.com/suyuan32/simple-admin-tools/issues/68) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/jaeger-1.15.0
* Merge pull request [#67](https://github.com/suyuan32/simple-admin-tools/issues/67) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc-1.15.0
* Merge pull request [#69](https://github.com/suyuan32/simple-admin-tools/issues/69) from
  suyuan32/dependabot/go_modules/go.opentelemetry.io/otel-1.15.0
* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8
* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from
  suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from
  suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from
  suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0
* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from
  suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from
  suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master

<a name="tools/goctl/v1.5.2"></a>

## [tools/goctl/v1.5.2](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.3...tools/goctl/v1.5.2)

> 2023-04-24

### Chore

* update go-zero for goctl ([#3165](https://github.com/suyuan32/simple-admin-tools/issues/3165))
* refactor ([#3164](https://github.com/suyuan32/simple-admin-tools/issues/3164))
* use %q instead of %s for unmarshaling ([#3163](https://github.com/suyuan32/simple-admin-tools/issues/3163))
* remove unnecessary code ([#3161](https://github.com/suyuan32/simple-admin-tools/issues/3161))
* refactor gateway ([#3157](https://github.com/suyuan32/simple-admin-tools/issues/3157))
* update codecov to ignore mock files ([#3155](https://github.com/suyuan32/simple-admin-tools/issues/3155))
* make error message more readable ([#3154](https://github.com/suyuan32/simple-admin-tools/issues/3154))

### Feat

* refactor gateway code ([#3160](https://github.com/suyuan32/simple-admin-tools/issues/3160))
* inheritance rewrite error prompt is more
  friendly ([#3156](https://github.com/suyuan32/simple-admin-tools/issues/3156))

### Fix

*
fixed [#2945](https://github.com/suyuan32/simple-admin-tools/issues/2945) ([#2953](https://github.com/suyuan32/simple-admin-tools/issues/2953))

### Gateway

* open timeout function cabinet ([#3047](https://github.com/suyuan32/simple-admin-tools/issues/3047))

<a name="v1.5.3"></a>

## [v1.5.3](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.2...v1.5.3)

> 2023-04-22

### Chore

* update default tool version to 1.5.3
* fix go 1.18 dep problem ([#3149](https://github.com/suyuan32/simple-admin-tools/issues/3149))
* refactor max/min in fx ([#3135](https://github.com/suyuan32/simple-admin-tools/issues/3135))

### Docs

* update changelog

### Feat

* supported enum.
* set error code function for validator
* add drone command

### Feature

* two evaluate operate func addtion in package
  stream  ([#3129](https://github.com/suyuan32/simple-admin-tools/issues/3129))

### Fix

* merge latest codes
* default value for keepalive set to 0 ([#3152](https://github.com/suyuan32/simple-admin-tools/issues/3152))
* covert service name before generation
* api ent logic code service name to camel
* trans error template in authority middleware
* update ent page template
* update version

### Pull Requests

* Merge pull request [#62](https://github.com/suyuan32/simple-admin-tools/issues/62) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.2
* Merge pull request [#57](https://github.com/suyuan32/simple-admin-tools/issues/57) from
  suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.1
* Merge pull request [#56](https://github.com/suyuan32/simple-admin-tools/issues/56) from
  suyuan32/dependabot/go_modules/github.com/prometheus/client_golang-1.15.0
* Merge pull request [#61](https://github.com/suyuan32/simple-admin-tools/issues/61) from noahlann/feat-enum
* Merge pull request [#55](https://github.com/suyuan32/simple-admin-tools/issues/55) from
  suyuan32/dependabot/go_modules/go.etcd.io/etcd/api/v3-3.5.8

<a name="v1.5.2"></a>
## [v1.5.2](https://github.com/suyuan32/simple-admin-tools/compare/v0.3.3-beta...v1.5.2)

> 2023-04-16

### Docs

* add copyright comment

### Feat

* add validation register function

### Fix

* update version
* home dir in dockerfile change to app
* optimize makefile template
* optimize the kube command
* optimize the git lab template
* remove quickstart
* remove unused files
* remove redundant comment in makefile
* ent time type checker
* api proto gen create template
* lowercase for ent gen package


<a name="v0.3.3-beta"></a>

## [v0.3.3-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.3.2...v0.3.3-beta)

> 2023-04-13

### Chore

* update goctls version

### Feat

* new api support camelcase name
* rpc new support camelcase
* ent template command
* default tool version for new command

### Fix

* optimize console log
* wrong variables in rpc new gen
* tidy the deps when init ent
* empty module bug and single ent code generating bugs
* makefile tab bugs
* update the goctls version code
* remove unused migrate codes
* wrong params in makefile
* update flag info
* add version to makefile

### Refactor

* optimize gen rpc comment
* optimize gen-rpc help message

### Pull Requests

* Merge pull request [#53](https://github.com/suyuan32/simple-admin-tools/issues/53) from CtraPlatform/master
* Merge pull request [#52](https://github.com/suyuan32/simple-admin-tools/issues/52) from CtraPlatform/master
* Merge pull request [#47](https://github.com/suyuan32/simple-admin-tools/issues/47) from suyuan32/dependabot/go_modules/golang.org/x/net-0.9.0
* Merge pull request [#48](https://github.com/suyuan32/simple-admin-tools/issues/48) from suyuan32/dependabot/go_modules/go.mongodb.org/mongo-driver-1.11.4
* Merge pull request [#49](https://github.com/suyuan32/simple-admin-tools/issues/49) from suyuan32/dependabot/go_modules/golang.org/x/text-0.9.0
* Merge pull request [#50](https://github.com/suyuan32/simple-admin-tools/issues/50) from suyuan32/dependabot/go_modules/tools/goctl/github.com/zeromicro/go-zero-1.5.1
* Merge pull request [#51](https://github.com/suyuan32/simple-admin-tools/issues/51) from suyuan32/dependabot/go_modules/tools/goctl/entgo.io/ent-0.12.0


<a name="v0.3.2"></a>
## [v0.3.2](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.5.1...v0.3.2)

> 2023-04-09

### Chore

* update go version and dependencies
* update deps
* update dependencies
* update all dependencies

### Docs

* update CHANGELOG.md
* update change log
* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* shorthand for commands
* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params
* validator generation
* optimize swagger validator
* add optional supported. (proto2/3)
* API single service command
* windows build script
* add overwrite parameters for code gen
* json tag style command
* add lang to context
* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* bool component support
* status code gen in vben
* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* add gitignore
* update goctls version
* wrong usage in console function
* optimize flag intro
* remove unused compare function
* add flag intro
* antlr version
* optimize swagger
* test/test_test.go
* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl
* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support
* some ent gen bugs
* bugs in config.go generating
* bug when field need to upper and convert type
* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* makefile bug
* authority middleware
* redundant err in service context
* change page default order to desc
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* optimize dockerfile.tpl
* update makefile template
* generate pb file to types directory
* generate pb file to types directory
* uuidx use common package
* replace simple admin core pkg to  simple admin common
* optimize ent error handling
* optimize create logic
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* remove consul yaml config
* cancel the consul and use k8s in generation

### Style

* go fmt.

### Wip

* text gen
* tag validate
* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix
* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master
* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master
* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg
* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg
* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg
* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common
* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master


<a name="tools/goctl/v1.5.1"></a>
## [tools/goctl/v1.5.1](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.1...tools/goctl/v1.5.1)

> 2023-04-08

### Chore

* update go-zero version in goctl ([#3109](https://github.com/suyuan32/simple-admin-tools/issues/3109))


<a name="v1.5.1"></a>
## [v1.5.1](https://github.com/suyuan32/simple-admin-tools/compare/v0.3.2-beta...v1.5.1)

> 2023-04-08

### Chore

* coding style ([#3106](https://github.com/suyuan32/simple-admin-tools/issues/3106))

### Fix

* fillDefault,  handling of nested structs ([#3072](https://github.com/suyuan32/simple-admin-tools/issues/3072))
* should not conflict on lower members ([#3095](https://github.com/suyuan32/simple-admin-tools/issues/3095))

### Refactor

* simplify zrpc stat config ([#3107](https://github.com/suyuan32/simple-admin-tools/issues/3107))


<a name="v0.3.2-beta"></a>
## [v0.3.2-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.3.1...v0.3.2-beta)

> 2023-04-07

### Feat

* extra command
* supported proto map field, replace index+1 with original field sequence.
* supported prefix. (swagger-doc)
* proto field style params

### Fix

* optimize flag intro
* remove unused compare function
* add flag intro

### Wip

* text gen

### Pull Requests

* Merge pull request [#46](https://github.com/suyuan32/simple-admin-tools/issues/46) from zeromicro/master
* Merge pull request [#44](https://github.com/suyuan32/simple-admin-tools/issues/44) from suyuan32/dependabot/go_modules/tools/goctl/github.com/spf13/cobra-1.7.0
* Merge pull request [#45](https://github.com/suyuan32/simple-admin-tools/issues/45) from suyuan32/dependabot/go_modules/tools/goctl/golang.org/x/text-0.9.0
* Merge pull request [#42](https://github.com/suyuan32/simple-admin-tools/issues/42) from crazy6995/feat-mapfield
* Merge pull request [#43](https://github.com/suyuan32/simple-admin-tools/issues/43) from CtraPlatform/master
* Merge pull request [#41](https://github.com/suyuan32/simple-admin-tools/issues/41) from crazy6995/feat-prefix


<a name="v0.3.1"></a>
## [v0.3.1](https://github.com/suyuan32/simple-admin-tools/compare/v0.3.0...v0.3.1)

> 2023-04-03

### Chore

* add more tests ([#3084](https://github.com/suyuan32/simple-admin-tools/issues/3084))
* coding style ([#3080](https://github.com/suyuan32/simple-admin-tools/issues/3080))

### Docs

* update CHANGELOG.md

### Feat

* validator generation
* optimize swagger validator

### Fix

* antlr version

### Wip

* tag validate

### Pull Requests

* Merge pull request [#40](https://github.com/suyuan32/simple-admin-tools/issues/40) from suyuan32/mg
* Merge pull request [#39](https://github.com/suyuan32/simple-admin-tools/issues/39) from suyuan32/feat-validator-gen
* Merge pull request [#38](https://github.com/suyuan32/simple-admin-tools/issues/38) from Gowa2017/master


<a name="v0.3.0"></a>
## [v0.3.0](https://github.com/suyuan32/simple-admin-tools/compare/v0.3.0-beta...v0.3.0)

> 2023-03-30

### Chore

* coding style ([#3074](https://github.com/suyuan32/simple-admin-tools/issues/3074))

### Feat

* add optional supported. (proto2/3)

### Fix

* optimize swagger
* test/test_test.go

### Style

* go fmt.

### Pull Requests

* Merge pull request [#35](https://github.com/suyuan32/simple-admin-tools/issues/35) from suyuan32/mg
* Merge pull request [#34](https://github.com/suyuan32/simple-admin-tools/issues/34) from crazy6995/master


<a name="v0.3.0-beta"></a>
## [v0.3.0-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.9...v0.3.0-beta)

> 2023-03-28

### Chore

* update go version and dependencies

### Fix

* update dockerfile golang version
* optimize makefile
* optimize rpc makefile
* optimize enttx tpl
* optimize makefile tpl

### Refactor

* optimize dockerfile.tpl
* update makefile template


<a name="v0.2.9"></a>
## [v0.2.9](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.9-beta...v0.2.9)

> 2023-03-26

### Chore

* refactor zrpc setup ([#3064](https://github.com/suyuan32/simple-admin-tools/issues/3064))
* add more tests ([#3045](https://github.com/suyuan32/simple-admin-tools/issues/3045))

### Feat

* rest validation on http requests ([#3041](https://github.com/suyuan32/simple-admin-tools/issues/3041))

### Fix

* [#3058](https://github.com/suyuan32/simple-admin-tools/issues/3058) ([#3059](https://github.com/suyuan32/simple-admin-tools/issues/3059))

### Pull Requests

* Merge pull request [#33](https://github.com/suyuan32/simple-admin-tools/issues/33) from suyuan32/mg


<a name="v0.2.9-beta"></a>
## [v0.2.9-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.8...v0.2.9-beta)

> 2023-03-24

### Feat

* API single service command

### Fix

* gotype in vben and makefile in api service
* duplicate convert functions
* add ent int16/uint16 support

### Refactor

* generate pb file to types directory
* generate pb file to types directory


<a name="v0.2.8"></a>
## [v0.2.8](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.7...v0.2.8)

> 2023-03-19

### Chore

* update deps
* add more tests ([#3018](https://github.com/suyuan32/simple-admin-tools/issues/3018))

### Feat

* windows build script

### Fix

* some ent gen bugs
* bugs in config.go generating

### Pull Requests

* Merge pull request [#32](https://github.com/suyuan32/simple-admin-tools/issues/32) from suyuan32/feat-mg


<a name="v0.2.7"></a>
## [v0.2.7](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.3...v0.2.7)

> 2023-03-12

### Chore

* add more tests ([#3016](https://github.com/suyuan32/simple-admin-tools/issues/3016))
* refactor orm code ([#3015](https://github.com/suyuan32/simple-admin-tools/issues/3015))
* add more tests ([#3014](https://github.com/suyuan32/simple-admin-tools/issues/3014))
* update readme ([#3011](https://github.com/suyuan32/simple-admin-tools/issues/3011))
* add more tests ([#3010](https://github.com/suyuan32/simple-admin-tools/issues/3010))
* add more tests ([#3009](https://github.com/suyuan32/simple-admin-tools/issues/3009))
* add more tests ([#3006](https://github.com/suyuan32/simple-admin-tools/issues/3006))

### Docs

* update change log

### Feat

* unique redis addrs and trim spaces ([#3004](https://github.com/suyuan32/simple-admin-tools/issues/3004))
* add overwrite parameters for code gen

### Fix

* bug when field need to upper and convert type
* avoid unmarshal panic with incorrect map keys [#3002](https://github.com/suyuan32/simple-admin-tools/issues/3002) ([#3013](https://github.com/suyuan32/simple-admin-tools/issues/3013))

### Pull Requests

* Merge pull request [#31](https://github.com/suyuan32/simple-admin-tools/issues/31) from suyuan32/mg


<a name="v0.2.3"></a>
## [v0.2.3](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.3-beta...v0.2.3)

> 2023-03-08

### Chore

* remove optional in redis config ([#2979](https://github.com/suyuan32/simple-admin-tools/issues/2979))
* add comments
* clear errors on conf conflict keys ([#2972](https://github.com/suyuan32/simple-admin-tools/issues/2972))
* add tests ([#2960](https://github.com/suyuan32/simple-admin-tools/issues/2960))
* remove redundant prefix of "error: " in error creation
* add tests for logc debug
* add comments
* add more tests
* add more tests
* go mod tidy and update deps
* go mod tidy and update deps
* refine rest validator ([#2928](https://github.com/suyuan32/simple-admin-tools/issues/2928))
* reformat code ([#2925](https://github.com/suyuan32/simple-admin-tools/issues/2925))

### Feat

* json tag style command
* add lang to context
* conf add FillDefault func
* support grpc client keepalive config ([#2950](https://github.com/suyuan32/simple-admin-tools/issues/2950))
* add debug log for logc
* check key overwritten
* add configurable validator for httpx.Parse ([#2923](https://github.com/suyuan32/simple-admin-tools/issues/2923))

### Fix

* proto test
* remove redundant imports
* optimize lang in context
* optimize lang in context
* middleware trans bug
* bugs in authorization middleware
* test failure
* config map with json tag
* gateway conf doesn't work ([#2968](https://github.com/suyuan32/simple-admin-tools/issues/2968))
* security [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) ([#2949](https://github.com/suyuan32/simple-admin-tools/issues/2949))
* config map cannot handle case-insensitive keys. ([#2932](https://github.com/suyuan32/simple-admin-tools/issues/2932))
* [#2899](https://github.com/suyuan32/simple-admin-tools/issues/2899), using autoscaling/v2beta2 instead of v2beta1 ([#2900](https://github.com/suyuan32/simple-admin-tools/issues/2900))
* timeout not working if greater than global rest timeout ([#2926](https://github.com/suyuan32/simple-admin-tools/issues/2926))

### Refactor

* uuidx use common package
* replace simple admin core pkg to  simple admin common

### Pull Requests

* Merge pull request [#30](https://github.com/suyuan32/simple-admin-tools/issues/30) from suyuan32/mg
* Merge pull request [#28](https://github.com/suyuan32/simple-admin-tools/issues/28) from suyuan32/refator-common


<a name="v0.2.3-beta"></a>

## [v0.2.3-beta](https://github.com/suyuan32/simple-admin-tools/compare/v1.4.5...v0.2.3-beta)

> 2023-03-07

### Chore

* update dependencies
* reformat code ([#2903](https://github.com/suyuan32/simple-admin-tools/issues/2903))
* refactor ([#2875](https://github.com/suyuan32/simple-admin-tools/issues/2875))
* add more tests ([#2873](https://github.com/suyuan32/simple-admin-tools/issues/2873))
* remove clickhouse, added to zero-contrib ([#2848](https://github.com/suyuan32/simple-admin-tools/issues/2848))
* add more tests ([#2866](https://github.com/suyuan32/simple-admin-tools/issues/2866))
* improve codecov ([#2828](https://github.com/suyuan32/simple-admin-tools/issues/2828))
* fix missing funcs on windows ([#2825](https://github.com/suyuan32/simple-admin-tools/issues/2825))
* update goctl interface{} to any ([#2819](https://github.com/suyuan32/simple-admin-tools/issues/2819))
* change interface{} to any ([#2818](https://github.com/suyuan32/simple-admin-tools/issues/2818))
* add more tests ([#2815](https://github.com/suyuan32/simple-admin-tools/issues/2815))
* add more tests ([#2814](https://github.com/suyuan32/simple-admin-tools/issues/2814))
* add more tests ([#2812](https://github.com/suyuan32/simple-admin-tools/issues/2812))
* update all dependencies

### Docs

* update change log
* update change log
* update change log
* add license comment
* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* add status error to errorx and remove error msg
* ent error handling
* validate parameters
* Add request.ts ([#2901](https://github.com/suyuan32/simple-admin-tools/issues/2901))
* validate parameters
* add ldflags to reduce the size of binary file
* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* use dependabot for goctl ([#2869](https://github.com/suyuan32/simple-admin-tools/issues/2869))
* revert bytes.Buffer to strings.Builder
* bool component support
* status code gen in vben
* converge grpc interceptor processing ([#2830](https://github.com/suyuan32/simple-admin-tools/issues/2830))
* add MustNewRedis ([#2824](https://github.com/suyuan32/simple-admin-tools/issues/2824))
* mapreduce generic version ([#2827](https://github.com/suyuan32/simple-admin-tools/issues/2827))
* split proto files
* upgrade go to v1.18 ([#2817](https://github.com/suyuan32/simple-admin-tools/issues/2817))
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* add required tag by default in form data
* update trans code in svc
* use casbin watcher
* optimize with tx function
* bugs in parse test
* update locale
* fixed the bug that old trace instances may be fetched
* makefile bug
* test failures ([#2892](https://github.com/suyuan32/simple-admin-tools/issues/2892))
* authority middleware
* redundant err in service context
* change page default order to desc
* test failure ([#2874](https://github.com/suyuan32/simple-admin-tools/issues/2874))
* loop reset nextStart
* replace shoud replace the longest match
* optimize casbin template
* new function to init redis
* getById bug in api generation
* add url to upper check
* remove unused status req
* to lower camel case in vben gen
* status template uuid bug
* conf anonymous overlay problem ([#2847](https://github.com/suyuan32/simple-admin-tools/issues/2847))
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* improve ent generation
* problem on name overlaping in config ([#2820](https://github.com/suyuan32/simple-admin-tools/issues/2820))
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* optimize ent error handling
* optimize create logic
* simplify sqlx fail fast ping and  simplify miniredis setup in test ([#2897](https://github.com/suyuan32/simple-admin-tools/issues/2897))
* simplify stringx.Replacer, and avoid potential infinite loops ([#2877](https://github.com/suyuan32/simple-admin-tools/issues/2877))
* optimize interface
* change interface{} to any
* change api error pkg

### Revert

* remove consul yaml config
* cancel the consul and use k8s in generation

### Wip

* optimize status gen in api and rpc
* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master
* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg
* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface
* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg
* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master


<a name="v1.4.5"></a>

## [v1.4.5](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.4.5...v1.4.5)

> 2023-03-04


<a name="tools/goctl/v1.4.5"></a>

## [tools/goctl/v1.4.5](https://github.com/suyuan32/simple-admin-tools/compare/v1.5.0...tools/goctl/v1.4.5)

> 2023-03-04

### Chore

* remove any keywords
* add tests ([#2960](https://github.com/suyuan32/simple-admin-tools/issues/2960))
* remove redundant prefix of "error: " in error creation
* add tests for logc debug
* add comments
* add more tests
* add more tests
* refine rest validator ([#2928](https://github.com/suyuan32/simple-admin-tools/issues/2928))
* reformat code ([#2925](https://github.com/suyuan32/simple-admin-tools/issues/2925))
* reformat code ([#2903](https://github.com/suyuan32/simple-admin-tools/issues/2903))
* refactor ([#2875](https://github.com/suyuan32/simple-admin-tools/issues/2875))
* add more tests ([#2873](https://github.com/suyuan32/simple-admin-tools/issues/2873))
* add more tests ([#2866](https://github.com/suyuan32/simple-admin-tools/issues/2866))
* improve codecov ([#2828](https://github.com/suyuan32/simple-admin-tools/issues/2828))
* fix missing funcs on windows ([#2825](https://github.com/suyuan32/simple-admin-tools/issues/2825))
* add more tests ([#2815](https://github.com/suyuan32/simple-admin-tools/issues/2815))
* add more tests ([#2814](https://github.com/suyuan32/simple-admin-tools/issues/2814))
* add more tests ([#2812](https://github.com/suyuan32/simple-admin-tools/issues/2812))

### Feat

* conf add FillDefault func
* support grpc client keepalive config ([#2950](https://github.com/suyuan32/simple-admin-tools/issues/2950))
* add debug log for logc
* check key overwritten
* add configurable validator for httpx.Parse ([#2923](https://github.com/suyuan32/simple-admin-tools/issues/2923))
* Add request.ts ([#2901](https://github.com/suyuan32/simple-admin-tools/issues/2901))
* revert bytes.Buffer to strings.Builder
* converge grpc interceptor processing ([#2830](https://github.com/suyuan32/simple-admin-tools/issues/2830))
* add MustNewRedis ([#2824](https://github.com/suyuan32/simple-admin-tools/issues/2824))

### Fix

* config map cannot handle case-insensitive keys. ([#2932](https://github.com/suyuan32/simple-admin-tools/issues/2932))
* [#2899](https://github.com/suyuan32/simple-admin-tools/issues/2899), using autoscaling/v2beta2 instead of v2beta1 ([#2900](https://github.com/suyuan32/simple-admin-tools/issues/2900))
* timeout not working if greater than global rest timeout ([#2926](https://github.com/suyuan32/simple-admin-tools/issues/2926))
* fixed the bug that old trace instances may be fetched
* test failures ([#2892](https://github.com/suyuan32/simple-admin-tools/issues/2892))
* test failure ([#2874](https://github.com/suyuan32/simple-admin-tools/issues/2874))
* loop reset nextStart
* replace shoud replace the longest match
* conf anonymous overlay problem ([#2847](https://github.com/suyuan32/simple-admin-tools/issues/2847))
* problem on name overlaping in config ([#2820](https://github.com/suyuan32/simple-admin-tools/issues/2820))

### Refactor

* simplify sqlx fail fast ping and  simplify miniredis setup in test ([#2897](https://github.com/suyuan32/simple-admin-tools/issues/2897))
* simplify stringx.Replacer, and avoid potential infinite loops ([#2877](https://github.com/suyuan32/simple-admin-tools/issues/2877))


<a name="v1.5.0"></a>
## [v1.5.0](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.5.0...v1.5.0)

> 2023-03-04


<a name="tools/goctl/v1.5.0"></a>
## [tools/goctl/v1.5.0](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.2...tools/goctl/v1.5.0)

> 2023-03-04

### Chore

* add tests ([#2960](https://github.com/suyuan32/simple-admin-tools/issues/2960))
* remove redundant prefix of "error: " in error creation
* add tests for logc debug
* add comments
* add more tests
* add more tests
* go mod tidy and update deps
* go mod tidy and update deps
* refine rest validator ([#2928](https://github.com/suyuan32/simple-admin-tools/issues/2928))
* reformat code ([#2925](https://github.com/suyuan32/simple-admin-tools/issues/2925))

### Feat

* conf add FillDefault func
* support grpc client keepalive config ([#2950](https://github.com/suyuan32/simple-admin-tools/issues/2950))
* add debug log for logc
* check key overwritten
* add configurable validator for httpx.Parse ([#2923](https://github.com/suyuan32/simple-admin-tools/issues/2923))

### Fix

* security [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) ([#2949](https://github.com/suyuan32/simple-admin-tools/issues/2949))
* config map cannot handle case-insensitive keys. ([#2932](https://github.com/suyuan32/simple-admin-tools/issues/2932))
* [#2899](https://github.com/suyuan32/simple-admin-tools/issues/2899), using autoscaling/v2beta2 instead of v2beta1 ([#2900](https://github.com/suyuan32/simple-admin-tools/issues/2900))
* timeout not working if greater than global rest timeout ([#2926](https://github.com/suyuan32/simple-admin-tools/issues/2926))


<a name="v0.2.2"></a>
## [v0.2.2](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.2-beta...v0.2.2)

> 2023-02-28


<a name="v0.2.2-beta"></a>
## [v0.2.2-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.1...v0.2.2-beta)

> 2023-02-28

### Chore

* update dependencies

### Docs

* update change log

### Feat

* ent error handling

### Fix

* use casbin watcher
* optimize with tx function

### Pull Requests

* Merge pull request [#27](https://github.com/suyuan32/simple-admin-tools/issues/27) from vwenkk/master


<a name="v0.2.1"></a>
## [v0.2.1](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.0...v0.2.1)

> 2023-02-23

### Chore

* reformat code ([#2903](https://github.com/suyuan32/simple-admin-tools/issues/2903))

### Docs

* update change log

### Feat

* validate parameters
* Add request.ts ([#2901](https://github.com/suyuan32/simple-admin-tools/issues/2901))
* validate parameters
* add ldflags to reduce the size of binary file

### Fix

* bugs in parse test
* update locale
* fixed the bug that old trace instances may be fetched
* makefile bug
* test failures ([#2892](https://github.com/suyuan32/simple-admin-tools/issues/2892))
* authority middleware
* redundant err in service context
* change page default order to desc

### Refactor

* optimize create logic
* simplify sqlx fail fast ping and  simplify miniredis setup in test ([#2897](https://github.com/suyuan32/simple-admin-tools/issues/2897))
* simplify stringx.Replacer, and avoid potential infinite loops ([#2877](https://github.com/suyuan32/simple-admin-tools/issues/2877))


<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/suyuan32/simple-admin-tools/compare/v0.2.0-beta...v0.2.0)

> 2023-02-13

### Chore

* refactor ([#2875](https://github.com/suyuan32/simple-admin-tools/issues/2875))
* add more tests ([#2873](https://github.com/suyuan32/simple-admin-tools/issues/2873))
* remove clickhouse, added to zero-contrib ([#2848](https://github.com/suyuan32/simple-admin-tools/issues/2848))
* add more tests ([#2866](https://github.com/suyuan32/simple-admin-tools/issues/2866))

### Docs

* add license comment

### Feat

* add enable for rpc client
* add enable for rpc client
* add enable for rpc client
* use dependabot for goctl ([#2869](https://github.com/suyuan32/simple-admin-tools/issues/2869))
* revert bytes.Buffer to strings.Builder

### Fix

* test failure ([#2874](https://github.com/suyuan32/simple-admin-tools/issues/2874))
* loop reset nextStart
* replace shoud replace the longest match
* optimize casbin template
* new function to init redis
* getById bug in api generation

### Pull Requests

* Merge pull request [#26](https://github.com/suyuan32/simple-admin-tools/issues/26) from suyuan32/mg


<a name="v0.2.0-beta"></a>
## [v0.2.0-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.9...v0.2.0-beta)

> 2023-02-09

### Fix

* add url to upper check
* remove unused status req

### Refactor

* optimize interface

### Pull Requests

* Merge pull request [#25](https://github.com/suyuan32/simple-admin-tools/issues/25) from suyuan32/refactor-interface


<a name="v0.1.9"></a>
## [v0.1.9](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.8...v0.1.9)

> 2023-02-06

### Chore

* improve codecov ([#2828](https://github.com/suyuan32/simple-admin-tools/issues/2828))
* fix missing funcs on windows ([#2825](https://github.com/suyuan32/simple-admin-tools/issues/2825))
* update goctl interface{} to any ([#2819](https://github.com/suyuan32/simple-admin-tools/issues/2819))
* change interface{} to any ([#2818](https://github.com/suyuan32/simple-admin-tools/issues/2818))
* add more tests ([#2815](https://github.com/suyuan32/simple-admin-tools/issues/2815))
* add more tests ([#2814](https://github.com/suyuan32/simple-admin-tools/issues/2814))
* add more tests ([#2812](https://github.com/suyuan32/simple-admin-tools/issues/2812))
* update goctl version to 1.4.4 ([#2811](https://github.com/suyuan32/simple-admin-tools/issues/2811))
* refactor func name ([#2804](https://github.com/suyuan32/simple-admin-tools/issues/2804))
* add more tests ([#2803](https://github.com/suyuan32/simple-admin-tools/issues/2803))

### Feat

* bool component support
* status code gen in vben
* converge grpc interceptor processing ([#2830](https://github.com/suyuan32/simple-admin-tools/issues/2830))
* add MustNewRedis ([#2824](https://github.com/suyuan32/simple-admin-tools/issues/2824))
* mapreduce generic version ([#2827](https://github.com/suyuan32/simple-admin-tools/issues/2827))
* upgrade go to v1.18 ([#2817](https://github.com/suyuan32/simple-admin-tools/issues/2817))

### Fix

* to lower camel case in vben gen
* status template uuid bug
* conf anonymous overlay problem ([#2847](https://github.com/suyuan32/simple-admin-tools/issues/2847))
* notification template
* api generation file model name lowercase
* rpc proto generation list req bug
* api status code gen template
* problem on name overlaping in config ([#2820](https://github.com/suyuan32/simple-admin-tools/issues/2820))
* mapping optional dep not canonicaled ([#2807](https://github.com/suyuan32/simple-admin-tools/issues/2807))

### Wip

* optimize status gen in api and rpc

### Pull Requests

* Merge pull request [#24](https://github.com/suyuan32/simple-admin-tools/issues/24) from suyuan32/mg
* Merge pull request [#23](https://github.com/suyuan32/simple-admin-tools/issues/23) from suyuan32/mg


<a name="v0.1.8"></a>

## [v0.1.8](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.4.4...v0.1.8)

> 2023-01-29

### Chore

* update all dependencies

### Docs

* update readme
* update changelog
* update CHANGELOG.md
* update CHANGELOG.md

### Feat

* split proto files
* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* uuid code generating for vben
* api uuid code generating
* rpc uuid code generating
* merge latest code
* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* improve ent generation
* swagger env bug
* command parameters and submit template
* rpc generating space bug
* extra command for linux
* multiple group
* remove sqlx and gorm
* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* ent api proto generating bug in type
* ent rpc generating type error
* drawer generating drawer props bug
* validator error type
* gen handler
* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug
* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* change interface{} to any
* change api error pkg

### Revert

* remove consul yaml config
* cancel the consul and use k8s in generation

### Wip

* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#22](https://github.com/suyuan32/simple-admin-tools/issues/22) from suyuan32/feat-proto-split
* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent
* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql
* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen
* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg
* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg
* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg
* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master


<a name="tools/goctl/v1.4.4"></a>

## [tools/goctl/v1.4.4](https://github.com/suyuan32/simple-admin-tools/compare/v1.4.4...tools/goctl/v1.4.4)

> 2023-01-21


<a name="v1.4.4"></a>

## [v1.4.4](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.7...v1.4.4)

> 2023-01-21

### Chore

* update goctl version to 1.4.4 ([#2811](https://github.com/suyuan32/simple-admin-tools/issues/2811))
* refactor func name ([#2804](https://github.com/suyuan32/simple-admin-tools/issues/2804))
* add more tests ([#2803](https://github.com/suyuan32/simple-admin-tools/issues/2803))

### Fix

* mapping optional dep not canonicaled ([#2807](https://github.com/suyuan32/simple-admin-tools/issues/2807))


<a name="v0.1.7"></a>
## [v0.1.7](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.6...v0.1.7)

> 2023-01-17

### Chore

* add more tests ([#2801](https://github.com/suyuan32/simple-admin-tools/issues/2801))
* add more tests ([#2800](https://github.com/suyuan32/simple-admin-tools/issues/2800))
* remove mgo related packages ([#2799](https://github.com/suyuan32/simple-admin-tools/issues/2799))
* add more tests ([#2797](https://github.com/suyuan32/simple-admin-tools/issues/2797))
* add more tests ([#2795](https://github.com/suyuan32/simple-admin-tools/issues/2795))
* add more tests ([#2794](https://github.com/suyuan32/simple-admin-tools/issues/2794))
* add more tests ([#2792](https://github.com/suyuan32/simple-admin-tools/issues/2792))
* refactor ([#2785](https://github.com/suyuan32/simple-admin-tools/issues/2785))

### Docs

* update readme

### Feat

* merge latest code
* merge latest code
* add ent multiple support
* rpc ent multiple generation support
* go swagger auto install
* expose NewTimingWheelWithClock ([#2787](https://github.com/suyuan32/simple-admin-tools/issues/2787))

### Fix

* swagger env bug
* command parameters and submit template
* rpc generating space bug
* modify the generated update function and add return values for update and delete functions ([#2793](https://github.com/suyuan32/simple-admin-tools/issues/2793))
* extra command for linux
* multiple group

### Pull Requests

* Merge pull request [#21](https://github.com/suyuan32/simple-admin-tools/issues/21) from suyuan32/mg
* Merge pull request [#20](https://github.com/suyuan32/simple-admin-tools/issues/20) from suyuan32/feat-multiple-ent


<a name="v0.1.6"></a>
## [v0.1.6](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.5...v0.1.6)

> 2023-01-13

### Fix

* remove sqlx and gorm

### Pull Requests

* Merge pull request [#19](https://github.com/suyuan32/simple-admin-tools/issues/19) from suyuan32/rm-sql


<a name="v0.1.5"></a>
## [v0.1.5](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.5-beta...v0.1.5)

> 2023-01-12

### Chore

* add tests ([#2778](https://github.com/suyuan32/simple-admin-tools/issues/2778))

### Feat

* uuid code generating for vben
* support **struct in mapping ([#2784](https://github.com/suyuan32/simple-admin-tools/issues/2784))
* api uuid code generating
* rpc uuid code generating
* support ptr of ptr of ... in mapping ([#2779](https://github.com/suyuan32/simple-admin-tools/issues/2779))

### Fix

* redis
* redis
* merge latest code
* remove rpc uuid_pk parameter
* remove default sql generating code
* remove default sql generating code
* update base.api
* replace tab by space in api file
* [#2576](https://github.com/suyuan32/simple-admin-tools/issues/2576) ([#2776](https://github.com/suyuan32/simple-admin-tools/issues/2776))

### Refactor

* change interface{} to any

### Pull Requests

* Merge pull request [#18](https://github.com/suyuan32/simple-admin-tools/issues/18) from suyuan32/mg
* Merge pull request [#17](https://github.com/suyuan32/simple-admin-tools/issues/17) from suyuan32/feat-uuid-gen


<a name="v0.1.5-beta"></a>
## [v0.1.5-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.4...v0.1.5-beta)

> 2023-01-11

### Chore

* add tests ([#2774](https://github.com/suyuan32/simple-admin-tools/issues/2774))
* remove simple methods, inlined ([#2768](https://github.com/suyuan32/simple-admin-tools/issues/2768))
* refactor ([#2764](https://github.com/suyuan32/simple-admin-tools/issues/2764))
* remove unnecessary code ([#2754](https://github.com/suyuan32/simple-admin-tools/issues/2754))
* improve codecov ([#2752](https://github.com/suyuan32/simple-admin-tools/issues/2752))
* remove roadmap file, not updating ([#2749](https://github.com/suyuan32/simple-admin-tools/issues/2749))
* reorg imports ([#2745](https://github.com/suyuan32/simple-admin-tools/issues/2745))
* update tests ([#2741](https://github.com/suyuan32/simple-admin-tools/issues/2741))

### Docs

* update changelog

### Feat

* merge latest code
* add config to truncate long log content ([#2767](https://github.com/suyuan32/simple-admin-tools/issues/2767))
* replace NewBetchInserter function name ([#2769](https://github.com/suyuan32/simple-admin-tools/issues/2769))
* add middlewares config for zrpc ([#2766](https://github.com/suyuan32/simple-admin-tools/issues/2766))
* add middlewares config for rest ([#2765](https://github.com/suyuan32/simple-admin-tools/issues/2765))
* add batch inserter ([#2755](https://github.com/suyuan32/simple-admin-tools/issues/2755))
* add mongo options ([#2753](https://github.com/suyuan32/simple-admin-tools/issues/2753))
* trace http.status_code ([#2708](https://github.com/suyuan32/simple-admin-tools/issues/2708))

### Fix

* ent api proto generating bug in type
* ent rpc generating type error
* replace goctl ExactValidArgs to MatchAll ([#2759](https://github.com/suyuan32/simple-admin-tools/issues/2759))
* drawer generating drawer props bug
* validator error type
* [#2700](https://github.com/suyuan32/simple-admin-tools/issues/2700), timeout not enough for writing responses ([#2738](https://github.com/suyuan32/simple-admin-tools/issues/2738))
* [#2735](https://github.com/suyuan32/simple-admin-tools/issues/2735) ([#2736](https://github.com/suyuan32/simple-admin-tools/issues/2736))

### Refactor

* simplify the code ([#2763](https://github.com/suyuan32/simple-admin-tools/issues/2763))
* use opentelemetry's standard api to track http status code ([#2760](https://github.com/suyuan32/simple-admin-tools/issues/2760))

### Pull Requests

* Merge pull request [#16](https://github.com/suyuan32/simple-admin-tools/issues/16) from suyuan32/mg

### BREAKING CHANGE


trace Config.Batcher should use otlpgrpc instead of grpc now.


<a name="v0.1.4"></a>
## [v0.1.4](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.4-beta...v0.1.4)

> 2022-12-30

### Chore

* pass by value for config in dev server ([#2712](https://github.com/suyuan32/simple-admin-tools/issues/2712))

### Feat

* ignorecolums add sort ([#2648](https://github.com/suyuan32/simple-admin-tools/issues/2648))

### Fix

* gen handler
* key like TLSConfig not working ([#2730](https://github.com/suyuan32/simple-admin-tools/issues/2730))
* etcd publisher reconnecting problem ([#2710](https://github.com/suyuan32/simple-admin-tools/issues/2710))
* camel cased key of map item in config ([#2715](https://github.com/suyuan32/simple-admin-tools/issues/2715))

### Pull Requests

* Merge pull request [#15](https://github.com/suyuan32/simple-admin-tools/issues/15) from suyuan32/mg


<a name="v0.1.4-beta"></a>
## [v0.1.4-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.3...v0.1.4-beta)

> 2022-12-18

### Docs

* update CHANGELOG.md

### Fix

* migrate version bug and search key num bug
* only generate makefile and dockerfile when we create new api
* etc template
* makefile transErr and service context template
* service context and ent format
* pagination template bug

### Refactor

* remove duplicated code ([#2705](https://github.com/suyuan32/simple-admin-tools/issues/2705))

### Pull Requests

* Merge pull request [#13](https://github.com/suyuan32/simple-admin-tools/issues/13) from suyuan32/mg


<a name="v0.1.3"></a>
## [v0.1.3](https://github.com/suyuan32/simple-admin-tools/compare/v1.4.3...v0.1.3)

> 2022-12-15

### Chore

* update all dependencies

### Docs

* update CHANGELOG.md

### Feat

* group for rpc logic
* gitlab-ci.yml generating
* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* makefile template
* optional gen makefile and dockerfile
* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* change api error pkg

### Revert

* remove consul yaml config
* cancel the consul and use k8s in generation

### Wip

* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#12](https://github.com/suyuan32/simple-admin-tools/issues/12) from suyuan32/feat-group-logic
* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master


<a name="v1.4.3"></a>
## [v1.4.3](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.2...v1.4.3)

> 2022-12-14

### Refactor

* remove duplicated code ([#2705](https://github.com/suyuan32/simple-admin-tools/issues/2705))


<a name="v0.1.2"></a>
## [v0.1.2](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.4.3...v0.1.2)

> 2022-12-13

### Chore

* update all dependencies

### Docs

* update CHANGELOG.md

### Feat

* vben code generation via api file
* service port parameter
* service port parameter
* api crud generation by proto
* auto migrate for rpc generation
* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent
* error translation
* gorm logger
* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* cases with no lower
* merge consul mod to go zero
* all deprecated function
* validate bugs
* bugs in tests
* optimize delete button
* makefile push bug
* update ErrorCtx logic
* optimize api url
* delete ent in tools
* modify error code
* optimize go gen types
* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error
* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Perf

* optimize route swagger generation

### Refactor

* change api error pkg

### Revert

* remove consul yaml config
* cancel the consul and use k8s in generation

### Wip

* vben code generation
* api code generation
* api code generation
* ent logic generating

### Pull Requests

* Merge pull request [#11](https://github.com/suyuan32/simple-admin-tools/issues/11) from suyuan32/mg
* Merge pull request [#10](https://github.com/suyuan32/simple-admin-tools/issues/10) from suyuan32/feat-upgrade-go
* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen
* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen
* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen
* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master


<a name="tools/goctl/v1.4.3"></a>
## [tools/goctl/v1.4.3](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.1...tools/goctl/v1.4.3)

> 2022-12-13

### Feat

* add dev server and health ([#2665](https://github.com/suyuan32/simple-admin-tools/issues/2665))
* accept camelcase for config keys ([#2651](https://github.com/suyuan32/simple-admin-tools/issues/2651))

### Fix

* [#2684](https://github.com/suyuan32/simple-admin-tools/issues/2684) ([#2693](https://github.com/suyuan32/simple-admin-tools/issues/2693))
* Fix string.title ([#2687](https://github.com/suyuan32/simple-admin-tools/issues/2687))
* [#2672](https://github.com/suyuan32/simple-admin-tools/issues/2672) ([#2681](https://github.com/suyuan32/simple-admin-tools/issues/2681))


<a name="v0.1.1"></a>
## [v0.1.1](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.0...v0.1.1)

> 2022-12-08

### Chore

* tidy go.sum ([#2675](https://github.com/suyuan32/simple-admin-tools/issues/2675))
* update deps ([#2674](https://github.com/suyuan32/simple-admin-tools/issues/2674))
* upgrade dependencies ([#2658](https://github.com/suyuan32/simple-admin-tools/issues/2658))
* update deps ([#2621](https://github.com/suyuan32/simple-admin-tools/issues/2621))
* update dependencies ([#2594](https://github.com/suyuan32/simple-admin-tools/issues/2594))

### Feat

* vben code generation via api file
* add trace.SpanIDFromContext and trace.TraceIDFromContext ([#2654](https://github.com/suyuan32/simple-admin-tools/issues/2654))
* add stringx.ToCamelCase ([#2622](https://github.com/suyuan32/simple-admin-tools/issues/2622))
* validate value in options for mapping ([#2616](https://github.com/suyuan32/simple-admin-tools/issues/2616))
* support bool for env tag ([#2593](https://github.com/suyuan32/simple-admin-tools/issues/2593))
* support env tag in config ([#2577](https://github.com/suyuan32/simple-admin-tools/issues/2577))

### Fix

* update ErrorCtx logic
* fix client side in zeromicro[#2109](https://github.com/suyuan32/simple-admin-tools/issues/2109) (zeromicro[#2116](https://github.com/suyuan32/simple-admin-tools/issues/2116)) ([#2659](https://github.com/suyuan32/simple-admin-tools/issues/2659))
* log currentSize should not be 0 when file exists and size is not 0 ([#2639](https://github.com/suyuan32/simple-admin-tools/issues/2639))
* fix conflict with the import package name ([#2610](https://github.com/suyuan32/simple-admin-tools/issues/2610))

### Wip

* vben code generation

### Pull Requests

* Merge pull request [#9](https://github.com/suyuan32/simple-admin-tools/issues/9) from suyuan32/feat-crud-gen


<a name="v0.1.0"></a>
## [v0.1.0](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.0-beta...v0.1.0)

> 2022-12-03

### Feat

* service port parameter


<a name="v0.1.0-beta"></a>
## [v0.1.0-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.1.0-beta-1...v0.1.0-beta)

> 2022-12-03

### Feat

* service port parameter
* api crud generation by proto
* auto migrate for rpc generation

### Fix

* optimize api url

### Wip

* api code generation
* api code generation

### Pull Requests

* Merge pull request [#8](https://github.com/suyuan32/simple-admin-tools/issues/8) from suyuan32/feat-crud-gen
* Merge pull request [#7](https://github.com/suyuan32/simple-admin-tools/issues/7) from suyuan32/feat-crud-gen


<a name="v0.1.0-beta-1"></a>
## [v0.1.0-beta-1](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.9...v0.1.0-beta-1)

> 2022-11-27

### Feat

* generate docker file
* proto file generation and logic code generation with ent
* proto file generation and logic code generation with ent

### Fix

* delete ent in tools

### Wip

* ent logic generating

### Pull Requests

* Merge pull request [#6](https://github.com/suyuan32/simple-admin-tools/issues/6) from suyuan32/feat-crud-gen


<a name="v0.0.9"></a>
## [v0.0.9](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.8...v0.0.9)

> 2022-11-12

### Fix

* modify error code


<a name="v0.0.8"></a>
## [v0.0.8](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.7.6...v0.0.8)

> 2022-11-11

### Feat

* error translation
* conf inherit ([#2568](https://github.com/suyuan32/simple-admin-tools/issues/2568))
* gorm logger

### Fix

* optimize go gen types
* inherit issue when parent after inherits ([#2586](https://github.com/suyuan32/simple-admin-tools/issues/2586))
* potential slice append issue ([#2560](https://github.com/suyuan32/simple-admin-tools/issues/2560))


<a name="v0.0.7.6"></a>
## [v0.0.7.6](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.7.3-beta...v0.0.7.6)

> 2022-10-28

### Chore

* update "DO NOT EDIT" format ([#2559](https://github.com/suyuan32/simple-admin-tools/issues/2559))
* add more tests ([#2547](https://github.com/suyuan32/simple-admin-tools/issues/2547))
* refactor ([#2545](https://github.com/suyuan32/simple-admin-tools/issues/2545))

### Feat

* add logger.WithFields ([#2546](https://github.com/suyuan32/simple-admin-tools/issues/2546))

### Fix

* rocketmq config add optional tag
* change default file name into snake format
* producer and consumer pointer error

### Perf

* optimize route swagger generation


<a name="v0.0.7.3-beta"></a>

## [v0.0.7.3-beta](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.4.2...v0.0.7.3-beta)

> 2022-10-26

### Feat

* rocket mq plugin
* gen consul code
* consul kv store configuration
* consul support
* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* bugs in accept language parsing
* bugs in accept language parsing
* delete log message reference from simple-admin-core
* merge latest code
* etc template
* update deployment in k8s
* StackCoolDownMillis name
* yaml key name
* JSON tag in config files
* system info in swagger
* Merge latest code
* bugs when run goctls new
* rest inline bug
* inline bug
* restore field for consul config
* json field for consul conf
* add yaml tag for all configuration
* bug in load
* change interface into pointer
* load function circle implement
* update change log
* package access
* merge latest code
* add validator test
* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Refactor

* change api error pkg

### Revert

* remove consul yaml config
* cancel the consul and use k8s in generation

### Pull Requests

* Merge pull request [#4](https://github.com/suyuan32/simple-admin-tools/issues/4) from zeromicro/master
* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master
* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master


<a name="tools/goctl/v1.4.2"></a>

## [tools/goctl/v1.4.2](https://github.com/suyuan32/simple-admin-tools/compare/v1.4.2...tools/goctl/v1.4.2)

> 2022-10-22


<a name="v1.4.2"></a>

## [v1.4.2](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.7.2-beta...v1.4.2)

> 2022-10-22

### Chore

* refactor ([#2545](https://github.com/suyuan32/simple-admin-tools/issues/2545))
* adjust rpc comment format ([#2501](https://github.com/suyuan32/simple-admin-tools/issues/2501))
* add more tests ([#2536](https://github.com/suyuan32/simple-admin-tools/issues/2536))
* fix lint errors ([#2520](https://github.com/suyuan32/simple-admin-tools/issues/2520))
* add golangci-lint config file ([#2519](https://github.com/suyuan32/simple-admin-tools/issues/2519))
* fix naming problem ([#2500](https://github.com/suyuan32/simple-admin-tools/issues/2500))
* sqlx's metric name is different from redis ([#2505](https://github.com/suyuan32/simple-admin-tools/issues/2505))

### Feat

* add logger.WithFields ([#2546](https://github.com/suyuan32/simple-admin-tools/issues/2546))
* remove info log when disable log ([#2525](https://github.com/suyuan32/simple-admin-tools/issues/2525))
* support uuid.UUID in mapping ([#2537](https://github.com/suyuan32/simple-admin-tools/issues/2537))

### Fix

* redis's pipeline logs are not printed completely ([#2538](https://github.com/suyuan32/simple-admin-tools/issues/2538))


<a name="v0.0.7.2-beta"></a>
## [v0.0.7.2-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.7.1-beta...v0.0.7.2-beta)

> 2022-10-22

### Fix

* bugs in accept language parsing


<a name="v0.0.7.1-beta"></a>
## [v0.0.7.1-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.7...v0.0.7.1-beta)

> 2022-10-22

### Fix

* bugs in accept language parsing
* delete log message reference from simple-admin-core


<a name="v0.0.7"></a>
## [v0.0.7](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.7-beta-3...v0.0.7)

> 2022-10-12

### Chore

* remove unnecessary code ([#2499](https://github.com/suyuan32/simple-admin-tools/issues/2499))
* remove init if possible ([#2485](https://github.com/suyuan32/simple-admin-tools/issues/2485))

### Fix

* merge latest code
* etc template
* replace Infof() with Errorf() in DurationInterceptor ([#2495](https://github.com/suyuan32/simple-admin-tools/issues/2495)) ([#2497](https://github.com/suyuan32/simple-admin-tools/issues/2497))
* update deployment in k8s


<a name="v0.0.7-beta-3"></a>
## [v0.0.7-beta-3](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.7-beta-2...v0.0.7-beta-3)

> 2022-10-09

### Fix

* StackCoolDownMillis name


<a name="v0.0.7-beta-2"></a>
## [v0.0.7-beta-2](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.7-beta-1...v0.0.7-beta-2)

> 2022-10-09

### Revert

* remove consul yaml config


<a name="v0.0.7-beta-1"></a>
## [v0.0.7-beta-1](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.7-beta...v0.0.7-beta-1)

> 2022-10-09

### Fix

* yaml key name


<a name="v0.0.7-beta"></a>
## [v0.0.7-beta](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.6...v0.0.7-beta)

> 2022-10-09

### Chore

* refactor to reduce duplicated code ([#2477](https://github.com/suyuan32/simple-admin-tools/issues/2477))
* better shedding algorithm, make sure recover from shedding ([#2476](https://github.com/suyuan32/simple-admin-tools/issues/2476))
* sort methods ([#2470](https://github.com/suyuan32/simple-admin-tools/issues/2470))
* gofumpt ([#2439](https://github.com/suyuan32/simple-admin-tools/issues/2439))

### Feat

* add logc package, support AddGlobalFields for both logc and logx. ([#2463](https://github.com/suyuan32/simple-admin-tools/issues/2463))
* add string to map in httpx parse method ([#2459](https://github.com/suyuan32/simple-admin-tools/issues/2459))
* add color to debug ([#2433](https://github.com/suyuan32/simple-admin-tools/issues/2433))

### Fix

* JSON tag in config files
* system info in swagger
* Merge latest code
* etcd reconnecting problem ([#2478](https://github.com/suyuan32/simple-admin-tools/issues/2478))
* add more tests ([#2473](https://github.com/suyuan32/simple-admin-tools/issues/2473))
* bugs when run goctls new

### Refactor

* adjust http request slow log format ([#2440](https://github.com/suyuan32/simple-admin-tools/issues/2440))

### Revert

* cancel the consul and use k8s in generation

### Pull Requests

* Merge pull request [#3](https://github.com/suyuan32/simple-admin-tools/issues/3) from zeromicro/master


<a name="v0.0.6"></a>
## [v0.0.6](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.6-beta-7...v0.0.6)

> 2022-09-23

### Feat

* gen consul code


<a name="v0.0.6-beta-7"></a>
## [v0.0.6-beta-7](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.6-beta-6...v0.0.6-beta-7)

> 2022-09-23

### Fix

* rest inline bug


<a name="v0.0.6-beta-6"></a>
## [v0.0.6-beta-6](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.6-beta-5...v0.0.6-beta-6)

> 2022-09-23

### Fix

* inline bug
* restore field for consul config
* json field for consul conf


<a name="v0.0.6-beta-5"></a>
## [v0.0.6-beta-5](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.6-beta-4...v0.0.6-beta-5)

> 2022-09-23

### Fix

* add yaml tag for all configuration
* bug in load
* change interface into pointer


<a name="v0.0.6-beta-4"></a>
## [v0.0.6-beta-4](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.6-beta-3...v0.0.6-beta-4)

> 2022-09-23

### Fix

* load function circle implement


<a name="v0.0.6-beta-3"></a>
## [v0.0.6-beta-3](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.6-beta-1...v0.0.6-beta-3)

> 2022-09-23

### Feat

* consul kv store configuration


<a name="v0.0.6-beta-1"></a>
## [v0.0.6-beta-1](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.6-beta-2...v0.0.6-beta-1)

> 2022-09-22


<a name="v0.0.6-beta-2"></a>
## [v0.0.6-beta-2](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.5...v0.0.6-beta-2)

> 2022-09-22

### Feat

* consul support

### Fix

* update change log


<a name="v0.0.5"></a>
## [v0.0.5](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.5-beta-2...v0.0.5)

> 2022-09-21

### Chore

* replace fmt.Fprint ([#2425](https://github.com/suyuan32/simple-admin-tools/issues/2425))

### Cleanup

* deprecated field and func ([#2416](https://github.com/suyuan32/simple-admin-tools/issues/2416))

### Fix

* fix log output ([#2424](https://github.com/suyuan32/simple-admin-tools/issues/2424))

### Refactor

* redis error for prometheus metric label ([#2412](https://github.com/suyuan32/simple-admin-tools/issues/2412))

### Pull Requests

* Merge pull request [#2](https://github.com/suyuan32/simple-admin-tools/issues/2) from zeromicro/master


<a name="v0.0.5-beta-2"></a>
## [v0.0.5-beta-2](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.5-beta-1...v0.0.5-beta-2)

> 2022-09-20

### Fix

* package access


<a name="v0.0.5-beta-1"></a>
## [v0.0.5-beta-1](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.5-beta...v0.0.5-beta-1)

> 2022-09-20

### Chore

* add more tests ([#2410](https://github.com/suyuan32/simple-admin-tools/issues/2410))
* add more tests ([#2409](https://github.com/suyuan32/simple-admin-tools/issues/2409))
* update go-zero to v1.4.1
* refactor the imports ([#2406](https://github.com/suyuan32/simple-admin-tools/issues/2406))
* refactor ([#2388](https://github.com/suyuan32/simple-admin-tools/issues/2388))

### Feat

* add log debug level ([#2411](https://github.com/suyuan32/simple-admin-tools/issues/2411))
* support caller skip in logx ([#2401](https://github.com/suyuan32/simple-admin-tools/issues/2401))
* mysql and redis metric support ([#2355](https://github.com/suyuan32/simple-admin-tools/issues/2355))
* add grpc export ([#2379](https://github.com/suyuan32/simple-admin-tools/issues/2379))

### Fix

* merge latest code
* add validator test


<a name="v0.0.5-beta"></a>
## [v0.0.5-beta](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.4.1...v0.0.5-beta)

> 2022-09-19

### Feat

* merge new codes from origin fix: swagger doc gen
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* add validator
* merge upstream
* recover go gen type
* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name

### Refactor

* change api error pkg


<a name="tools/goctl/v1.4.1"></a>
## [tools/goctl/v1.4.1](https://github.com/suyuan32/simple-admin-tools/compare/v1.4.1...tools/goctl/v1.4.1)

> 2022-09-17

### Chore

* update go-zero to v1.4.1

### Feat

* support caller skip in logx ([#2401](https://github.com/suyuan32/simple-admin-tools/issues/2401))


<a name="v1.4.1"></a>
## [v1.4.1](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.4...v1.4.1)

> 2022-09-17

### Chore

* refactor the imports ([#2406](https://github.com/suyuan32/simple-admin-tools/issues/2406))
* refactor ([#2388](https://github.com/suyuan32/simple-admin-tools/issues/2388))

### Feat

* mysql and redis metric support ([#2355](https://github.com/suyuan32/simple-admin-tools/issues/2355))
* add grpc export ([#2379](https://github.com/suyuan32/simple-admin-tools/issues/2379))


<a name="v0.0.4"></a>
## [v0.0.4](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.3...v0.0.4)

> 2022-09-13

### Chore

* refactor ([#2365](https://github.com/suyuan32/simple-admin-tools/issues/2365))

### Feat

* merge new codes from origin fix: swagger doc gen
* support targetPort option in goctl kube ([#2378](https://github.com/suyuan32/simple-admin-tools/issues/2378))
* support baggage propagation in httpc ([#2375](https://github.com/suyuan32/simple-admin-tools/issues/2375))

### Fix

* issue [#2359](https://github.com/suyuan32/simple-admin-tools/issues/2359) ([#2368](https://github.com/suyuan32/simple-admin-tools/issues/2368))
* merge upstream


<a name="v0.0.3"></a>
## [v0.0.3](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.2...v0.0.3)

> 2022-09-06

### Fix

* recover go gen type


<a name="v0.0.2"></a>
## [v0.0.2](https://github.com/suyuan32/simple-admin-tools/compare/v0.0.1...v0.0.2)

> 2022-09-06

### Chore

* remove unused packages ([#2312](https://github.com/suyuan32/simple-admin-tools/issues/2312))
* refactor gateway ([#2303](https://github.com/suyuan32/simple-admin-tools/issues/2303))

### Fix

* thread-safe in getWriter of logx ([#2319](https://github.com/suyuan32/simple-admin-tools/issues/2319))
* range validation on mapping ([#2317](https://github.com/suyuan32/simple-admin-tools/issues/2317))
* handle the scenarios that content-length is invalid ([#2313](https://github.com/suyuan32/simple-admin-tools/issues/2313))
* more accurate panic message on mapreduce ([#2311](https://github.com/suyuan32/simple-admin-tools/issues/2311))
* logx disable not working in some cases ([#2306](https://github.com/suyuan32/simple-admin-tools/issues/2306))

### Improve

* number range compare left and righ value ([#2315](https://github.com/suyuan32/simple-admin-tools/issues/2315))

### Refactor

* sequential range over safemap ([#2316](https://github.com/suyuan32/simple-admin-tools/issues/2316))


<a name="v0.0.1"></a>
## [v0.0.1](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.4.0...v0.0.1)

> 2022-08-26

### Chore

* refactor stat ([#2299](https://github.com/suyuan32/simple-admin-tools/issues/2299))
* Update readme ([#2280](https://github.com/suyuan32/simple-admin-tools/issues/2280))
* refactor logx ([#2262](https://github.com/suyuan32/simple-admin-tools/issues/2262))

### Feat

* rpc add health check function configuration optional ([#2288](https://github.com/suyuan32/simple-admin-tools/issues/2288))
* add go swagger support
* casbin util
* error message
* gorm conf

### Fix

* gen system info
* gen types swagger doc
* bug in rest response
* error msg
* package name
* unsignedTypeMap type error ([#2246](https://github.com/suyuan32/simple-admin-tools/issues/2246))
* test failure, due to go 1.19 compatibility ([#2256](https://github.com/suyuan32/simple-admin-tools/issues/2256))
* time repr wrapper ([#2255](https://github.com/suyuan32/simple-admin-tools/issues/2255))

### Test

* add more tests ([#2261](https://github.com/suyuan32/simple-admin-tools/issues/2261))


<a name="tools/goctl/v1.4.0"></a>
## [tools/goctl/v1.4.0](https://github.com/suyuan32/simple-admin-tools/compare/v1.4.0...tools/goctl/v1.4.0)

> 2022-08-07

### Chore

* release action for goctl ([#2239](https://github.com/suyuan32/simple-admin-tools/issues/2239))


<a name="v1.4.0"></a>
## [v1.4.0](https://github.com/suyuan32/simple-admin-tools/compare/v1.3.5...v1.4.0)

> 2022-08-07

### Chore

* renaming configs ([#2234](https://github.com/suyuan32/simple-admin-tools/issues/2234))
* refactor redislock ([#2210](https://github.com/suyuan32/simple-admin-tools/issues/2210))
* let logx.SetWriter can be called anytime ([#2186](https://github.com/suyuan32/simple-admin-tools/issues/2186))
* refactoring ([#2182](https://github.com/suyuan32/simple-admin-tools/issues/2182))
* refactoring logx ([#2181](https://github.com/suyuan32/simple-admin-tools/issues/2181))
* refactoring mapping name ([#2168](https://github.com/suyuan32/simple-admin-tools/issues/2168))
* remove unimplemented gateway ([#2139](https://github.com/suyuan32/simple-admin-tools/issues/2139))
* refactor ([#2130](https://github.com/suyuan32/simple-admin-tools/issues/2130))
* add more tests ([#2129](https://github.com/suyuan32/simple-admin-tools/issues/2129))
* coding style ([#2120](https://github.com/suyuan32/simple-admin-tools/issues/2120))

### Docs

* update docs for gateway ([#2236](https://github.com/suyuan32/simple-admin-tools/issues/2236))
* update goctl readme ([#2136](https://github.com/suyuan32/simple-admin-tools/issues/2136))

### Feat

* more meaningful error messages, close body on httpc requests ([#2238](https://github.com/suyuan32/simple-admin-tools/issues/2238))
* logx support logs rotation based on size limitation. ([#1652](https://github.com/suyuan32/simple-admin-tools/issues/1652)) ([#2167](https://github.com/suyuan32/simple-admin-tools/issues/2167))
* Support for multiple rpc service generation and rpc grouping ([#1972](https://github.com/suyuan32/simple-admin-tools/issues/1972))
* support customized header to metadata processor ([#2162](https://github.com/suyuan32/simple-admin-tools/issues/2162))
* support google.api.http in gateway ([#2161](https://github.com/suyuan32/simple-admin-tools/issues/2161))
* set content-type to application/json ([#2160](https://github.com/suyuan32/simple-admin-tools/issues/2160))
* verify RpcPath on startup ([#2159](https://github.com/suyuan32/simple-admin-tools/issues/2159))
* support form values in gateway ([#2158](https://github.com/suyuan32/simple-admin-tools/issues/2158))
* export gateway.Server to let users add middlewares ([#2157](https://github.com/suyuan32/simple-admin-tools/issues/2157))
* restful -> grpc gateway ([#2155](https://github.com/suyuan32/simple-admin-tools/issues/2155))
* support logx.WithFields ([#2128](https://github.com/suyuan32/simple-admin-tools/issues/2128))
* add Wrap and Wrapf in errorx ([#2126](https://github.com/suyuan32/simple-admin-tools/issues/2126))

### Fix

* [#2216](https://github.com/suyuan32/simple-admin-tools/issues/2216) ([#2235](https://github.com/suyuan32/simple-admin-tools/issues/2235))
* fix comment typo ([#2220](https://github.com/suyuan32/simple-admin-tools/issues/2220))
* handling rpc error on gateway ([#2212](https://github.com/suyuan32/simple-admin-tools/issues/2212))
* only setup logx once ([#2188](https://github.com/suyuan32/simple-admin-tools/issues/2188))
* logx test foo ([#2144](https://github.com/suyuan32/simple-admin-tools/issues/2144))
* remove invalid log fields in notLoggingContentMethods ([#2187](https://github.com/suyuan32/simple-admin-tools/issues/2187))
* fix switch doesn't work bug ([#2183](https://github.com/suyuan32/simple-admin-tools/issues/2183))
* fix [#2102](https://github.com/suyuan32/simple-admin-tools/issues/2102), [#2108](https://github.com/suyuan32/simple-admin-tools/issues/2108) ([#2131](https://github.com/suyuan32/simple-admin-tools/issues/2131))
* goctl genhandler duplicate rest/httpx & goctl genhandler template support custom import httpx package ([#2152](https://github.com/suyuan32/simple-admin-tools/issues/2152))
* generated sql query fields do not match template ([#2004](https://github.com/suyuan32/simple-admin-tools/issues/2004))


<a name="v1.3.5"></a>
## [v1.3.5](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.9...v1.3.5)

> 2022-07-09


<a name="tools/goctl/v1.3.9"></a>
## [tools/goctl/v1.3.9](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.8...tools/goctl/v1.3.9)

> 2022-07-09

### Chore

* remove blank lines ([#2117](https://github.com/suyuan32/simple-admin-tools/issues/2117))
* update goctl version to 1.3.9 ([#2111](https://github.com/suyuan32/simple-admin-tools/issues/2111))
* refactor ([#2087](https://github.com/suyuan32/simple-admin-tools/issues/2087))
* refactor ([#2085](https://github.com/suyuan32/simple-admin-tools/issues/2085))
* remove lifecycle preStop because sh not exist in scratch ([#2042](https://github.com/suyuan32/simple-admin-tools/issues/2042))
* Add command desc & color commands ([#2013](https://github.com/suyuan32/simple-admin-tools/issues/2013))
* refactor to simplify disabling builtin middlewares ([#2031](https://github.com/suyuan32/simple-admin-tools/issues/2031))
* upgrade action version ([#2027](https://github.com/suyuan32/simple-admin-tools/issues/2027))
* coding style ([#2012](https://github.com/suyuan32/simple-admin-tools/issues/2012))
* rename methods ([#1998](https://github.com/suyuan32/simple-admin-tools/issues/1998))
* update dependencies ([#1985](https://github.com/suyuan32/simple-admin-tools/issues/1985))

### Feat

* add method to jsonx ([#2049](https://github.com/suyuan32/simple-admin-tools/issues/2049))
* CompareAndSwapInt32 may be better than AddInt32 ([#2077](https://github.com/suyuan32/simple-admin-tools/issues/2077))
* rest.WithChain to replace builtin middlewares ([#2033](https://github.com/suyuan32/simple-admin-tools/issues/2033))
* support build Dockerfile from current dir ([#2021](https://github.com/suyuan32/simple-admin-tools/issues/2021))
* add trace in httpc ([#2011](https://github.com/suyuan32/simple-admin-tools/issues/2011))
* Replace mongo package with monc & mon ([#2002](https://github.com/suyuan32/simple-admin-tools/issues/2002))
* convert grpc errors to http status codes ([#1997](https://github.com/suyuan32/simple-admin-tools/issues/1997))
* add 'imagePullPolicy' parameter for 'goctl kube deploy' ([#1996](https://github.com/suyuan32/simple-admin-tools/issues/1996))

### Fix

* type matching supports string to int ([#2038](https://github.com/suyuan32/simple-admin-tools/issues/2038))
* `\u003cnil\u003e` log output when http server shutdown. ([#2055](https://github.com/suyuan32/simple-admin-tools/issues/2055))
* quickstart wrong package when go.mod exists in parent dir ([#2048](https://github.com/suyuan32/simple-admin-tools/issues/2048))
* 修复 clientinterceptors/tracinginterceptor.go 显示接受消息字节为0 ([#2003](https://github.com/suyuan32/simple-admin-tools/issues/2003))

### Typo

* add type keyword ([#1992](https://github.com/suyuan32/simple-admin-tools/issues/1992))


<a name="tools/goctl/v1.3.8"></a>
## [tools/goctl/v1.3.8](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.7...tools/goctl/v1.3.8)

> 2022-06-06

### Chore

* update goctl version to 1.3.8 ([#1981](https://github.com/suyuan32/simple-admin-tools/issues/1981))

### Fix

* generate bad Dockerfile on given dir ([#1980](https://github.com/suyuan32/simple-admin-tools/issues/1980))


<a name="tools/goctl/v1.3.7"></a>
## [tools/goctl/v1.3.7](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.6...tools/goctl/v1.3.7)

> 2022-06-05

### Chore

* make methods consistent in signatures ([#1971](https://github.com/suyuan32/simple-admin-tools/issues/1971))
* make print pretty ([#1967](https://github.com/suyuan32/simple-admin-tools/issues/1967))
* better mongo logs ([#1965](https://github.com/suyuan32/simple-admin-tools/issues/1965))
* update dependencies ([#1963](https://github.com/suyuan32/simple-admin-tools/issues/1963))

### Feat

* print routes ([#1964](https://github.com/suyuan32/simple-admin-tools/issues/1964))

### Fix

* The validation of tag "options" is not working with int/uint type ([#1969](https://github.com/suyuan32/simple-admin-tools/issues/1969))

### Test

* fix fails ([#1970](https://github.com/suyuan32/simple-admin-tools/issues/1970))
* make tests stable ([#1968](https://github.com/suyuan32/simple-admin-tools/issues/1968))


<a name="tools/goctl/v1.3.6"></a>
## [tools/goctl/v1.3.6](https://github.com/suyuan32/simple-admin-tools/compare/v1.3.4...tools/goctl/v1.3.6)

> 2022-06-03

### Chore

* update version ([#1961](https://github.com/suyuan32/simple-admin-tools/issues/1961))


<a name="v1.3.4"></a>

## [v1.3.4](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.5...v1.3.4)

> 2022-06-03

### Chore

* refactoring mapping string to slice ([#1959](https://github.com/suyuan32/simple-admin-tools/issues/1959))
* update roadmap ([#1948](https://github.com/suyuan32/simple-admin-tools/issues/1948))
* refine docker for better compatible with package main ([#1944](https://github.com/suyuan32/simple-admin-tools/issues/1944))
* add release action to auto build binaries ([#1884](https://github.com/suyuan32/simple-admin-tools/issues/1884))
* update k8s.io/client-go for security reason, go is upgrade to 1.16 ([#1912](https://github.com/suyuan32/simple-admin-tools/issues/1912))
* use get for quickstart, plain logs for easy understanding ([#1905](https://github.com/suyuan32/simple-admin-tools/issues/1905))
* coding style for quickstart ([#1902](https://github.com/suyuan32/simple-admin-tools/issues/1902))
* improve codecov ([#1878](https://github.com/suyuan32/simple-admin-tools/issues/1878))
* update some logs ([#1875](https://github.com/suyuan32/simple-admin-tools/issues/1875))
* fix deprecated usages ([#1871](https://github.com/suyuan32/simple-admin-tools/issues/1871))
* refine tests ([#1864](https://github.com/suyuan32/simple-admin-tools/issues/1864))
* use time.Now() instead of timex.Time() because go optimized it ([#1860](https://github.com/suyuan32/simple-admin-tools/issues/1860))

### Docs

* add docs for logx ([#1960](https://github.com/suyuan32/simple-admin-tools/issues/1960))
* update readme ([#1849](https://github.com/suyuan32/simple-admin-tools/issues/1849))

### Feat

* update docker alpine package mirror ([#1924](https://github.com/suyuan32/simple-admin-tools/issues/1924))
* set default connection idle time for grpc servers ([#1922](https://github.com/suyuan32/simple-admin-tools/issues/1922))
* support WithStreamClientInterceptor for zrpc clients ([#1907](https://github.com/suyuan32/simple-admin-tools/issues/1907))
* add toml config ([#1899](https://github.com/suyuan32/simple-admin-tools/issues/1899))
* Add goctl quickstart ([#1889](https://github.com/suyuan32/simple-admin-tools/issues/1889))
* logx with color ([#1872](https://github.com/suyuan32/simple-admin-tools/issues/1872))
* Replace cli to cobra ([#1855](https://github.com/suyuan32/simple-admin-tools/issues/1855))
* add fields with logx methods, support using third party logging libs. ([#1847](https://github.com/suyuan32/simple-admin-tools/issues/1847))

### Fix

* panic on convert to string on fillSliceFromString() ([#1951](https://github.com/suyuan32/simple-admin-tools/issues/1951))
* Useless delete cache logic in update ([#1923](https://github.com/suyuan32/simple-admin-tools/issues/1923))

### Refactor

* refactor trace in redis & sql & mongo ([#1865](https://github.com/suyuan32/simple-admin-tools/issues/1865))

### Test

* add codecov ([#1863](https://github.com/suyuan32/simple-admin-tools/issues/1863))
* add codecov ([#1861](https://github.com/suyuan32/simple-admin-tools/issues/1861))
* add more tests ([#1856](https://github.com/suyuan32/simple-admin-tools/issues/1856))

<a name="tools/goctl/v1.3.5"></a>

## [tools/goctl/v1.3.5](https://github.com/suyuan32/simple-admin-tools/compare/v1.3.3...tools/goctl/v1.3.5)

> 2022-04-28


<a name="v1.3.3"></a>

## [v1.3.3](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.4...v1.3.3)

> 2022-04-28

### Chhore

* fix usage typo ([#1797](https://github.com/suyuan32/simple-admin-tools/issues/1797))

### Chore

* optimize code ([#1818](https://github.com/suyuan32/simple-admin-tools/issues/1818))
* remove gofumpt -s flag, default to be enabled ([#1816](https://github.com/suyuan32/simple-admin-tools/issues/1816))
* refactor ([#1814](https://github.com/suyuan32/simple-admin-tools/issues/1814))
* Embed unit test data ([#1812](https://github.com/suyuan32/simple-admin-tools/issues/1812))
* use grpc.WithTransportCredentials and insecure.NewCredentials() instead of grpc.WithInsecure ([#1798](https://github.com/suyuan32/simple-admin-tools/issues/1798))
* avoid deadlock after stopping TimingWheel ([#1768](https://github.com/suyuan32/simple-admin-tools/issues/1768))
* remove legacy code ([#1766](https://github.com/suyuan32/simple-admin-tools/issues/1766))
* add doc ([#1764](https://github.com/suyuan32/simple-admin-tools/issues/1764))

### Feat

* Support model code generation for multi tables ([#1836](https://github.com/suyuan32/simple-admin-tools/issues/1836))
* support sub domain for cors ([#1827](https://github.com/suyuan32/simple-admin-tools/issues/1827))
* upgrade grpc to 1.46, and remove the deprecated grpc.WithBalancerName ([#1820](https://github.com/suyuan32/simple-admin-tools/issues/1820))
* add trace in redis & mon & sql ([#1799](https://github.com/suyuan32/simple-admin-tools/issues/1799))
* use mongodb official driver instead of mgo ([#1782](https://github.com/suyuan32/simple-admin-tools/issues/1782))
* add httpc.Do & httpc.Service.Do ([#1775](https://github.com/suyuan32/simple-admin-tools/issues/1775))
* add goctl docker build scripts ([#1760](https://github.com/suyuan32/simple-admin-tools/issues/1760))
* support ctx in kv methods ([#1759](https://github.com/suyuan32/simple-admin-tools/issues/1759))
* use go:embed to embed templates ([#1756](https://github.com/suyuan32/simple-admin-tools/issues/1756))

### Fix

* remove deprecated dependencies ([#1837](https://github.com/suyuan32/simple-admin-tools/issues/1837))
* rest: WriteJson get 200 when Marshal failed. ([#1803](https://github.com/suyuan32/simple-admin-tools/issues/1803))
* Fix issue [#1810](https://github.com/suyuan32/simple-admin-tools/issues/1810) ([#1811](https://github.com/suyuan32/simple-admin-tools/issues/1811))
* ignore timeout on websocket ([#1802](https://github.com/suyuan32/simple-admin-tools/issues/1802))
* Hdel check result & Pfadd check result ([#1801](https://github.com/suyuan32/simple-admin-tools/issues/1801))
* model unique keys generated differently in each re-generation ([#1771](https://github.com/suyuan32/simple-admin-tools/issues/1771))

### Refactor

* move json related header vars to internal ([#1840](https://github.com/suyuan32/simple-admin-tools/issues/1840))
* simplify the code ([#1835](https://github.com/suyuan32/simple-admin-tools/issues/1835))
* move postgres to pg package ([#1781](https://github.com/suyuan32/simple-admin-tools/issues/1781))


<a name="tools/goctl/v1.3.4"></a>
## [tools/goctl/v1.3.4](https://github.com/suyuan32/simple-admin-tools/compare/v1.3.2...tools/goctl/v1.3.4)

> 2022-04-03

### Chore

* update go-zero to v1.3.2 in goctl ([#1750](https://github.com/suyuan32/simple-admin-tools/issues/1750))


<a name="v1.3.2"></a>
## [v1.3.2](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.3...v1.3.2)

> 2022-04-03

### Chore

* update goctl version to 1.3.4 ([#1742](https://github.com/suyuan32/simple-admin-tools/issues/1742))
* refactor to use const instead of var ([#1731](https://github.com/suyuan32/simple-admin-tools/issues/1731))
* refactor code ([#1708](https://github.com/suyuan32/simple-admin-tools/issues/1708))
* refactor code ([#1700](https://github.com/suyuan32/simple-admin-tools/issues/1700))
* refactor code ([#1699](https://github.com/suyuan32/simple-admin-tools/issues/1699))
* fix lint issue ([#1694](https://github.com/suyuan32/simple-admin-tools/issues/1694))

### Feat

* simplify httpc ([#1748](https://github.com/suyuan32/simple-admin-tools/issues/1748))
* return original value of setbit in redis ([#1746](https://github.com/suyuan32/simple-admin-tools/issues/1746))
* let model customizable ([#1738](https://github.com/suyuan32/simple-admin-tools/issues/1738))
* remove reentrance in redislock, timeout bug ([#1704](https://github.com/suyuan32/simple-admin-tools/issues/1704))
* add getset command in redis and kv ([#1693](https://github.com/suyuan32/simple-admin-tools/issues/1693))
* add httpc.Parse ([#1698](https://github.com/suyuan32/simple-admin-tools/issues/1698))
* support -base to specify base image for goctl docker ([#1668](https://github.com/suyuan32/simple-admin-tools/issues/1668))
* add Dockerfile for goctl ([#1666](https://github.com/suyuan32/simple-admin-tools/issues/1666))
* Remove  command `goctl rpc proto`  ([#1665](https://github.com/suyuan32/simple-admin-tools/issues/1665))

### Fix

* model generation bug on with cache ([#1743](https://github.com/suyuan32/simple-admin-tools/issues/1743))
* empty slice are set to nil ([#1702](https://github.com/suyuan32/simple-admin-tools/issues/1702))
* the new  RawFieldNames considers the tag with options. ([#1663](https://github.com/suyuan32/simple-admin-tools/issues/1663))

### Refactor

* guard timeout on API files ([#1726](https://github.com/suyuan32/simple-admin-tools/issues/1726))
* simplify the code ([#1670](https://github.com/suyuan32/simple-admin-tools/issues/1670))


<a name="tools/goctl/v1.3.3"></a>
## [tools/goctl/v1.3.3](https://github.com/suyuan32/simple-admin-tools/compare/v1.3.1...tools/goctl/v1.3.3)

> 2022-03-17

### Chore

* remove unnecessary env ([#1654](https://github.com/suyuan32/simple-admin-tools/issues/1654))
* reduce the docker image size ([#1633](https://github.com/suyuan32/simple-admin-tools/issues/1633))
* update goctl version to 1.3.3, change docker build temp dir ([#1621](https://github.com/suyuan32/simple-admin-tools/issues/1621))
* refactor code ([#1613](https://github.com/suyuan32/simple-admin-tools/issues/1613))
* add unit tests ([#1615](https://github.com/suyuan32/simple-admin-tools/issues/1615))
* update go-zero to v1.3.1 in goctl ([#1599](https://github.com/suyuan32/simple-admin-tools/issues/1599))

### Feat

* add httpc/Service for convinience ([#1641](https://github.com/suyuan32/simple-admin-tools/issues/1641))
* add httpc/Get httpc/Post ([#1640](https://github.com/suyuan32/simple-admin-tools/issues/1640))
* add rest/httpc to make http requests governacible ([#1638](https://github.com/suyuan32/simple-admin-tools/issues/1638))
* support cpu stat on cgroups v2 ([#1636](https://github.com/suyuan32/simple-admin-tools/issues/1636))
* support oracle :N dynamic parameters ([#1552](https://github.com/suyuan32/simple-admin-tools/issues/1552))
* support scratch as the base docker image ([#1634](https://github.com/suyuan32/simple-admin-tools/issues/1634))

### Fix

* typo ([#1646](https://github.com/suyuan32/simple-admin-tools/issues/1646))
* Update unix-like path regex ([#1637](https://github.com/suyuan32/simple-admin-tools/issues/1637))
* HitQuota should be returned instead of Allowed when limit is equal to 1. ([#1581](https://github.com/suyuan32/simple-admin-tools/issues/1581))
* fix(gctl): apiparser_parser auto format ([#1607](https://github.com/suyuan32/simple-admin-tools/issues/1607))

### Refactor

* httpc package for easy to use ([#1645](https://github.com/suyuan32/simple-admin-tools/issues/1645))
* httpc package for easy to use ([#1643](https://github.com/suyuan32/simple-admin-tools/issues/1643))

### Test

* add more tests ([#1604](https://github.com/suyuan32/simple-admin-tools/issues/1604))


<a name="v1.3.1"></a>
## [v1.3.1](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.2...v1.3.1)

> 2022-03-01

### Build

* update goctl dependency ddl-parser to v1.0.3 ([#1586](https://github.com/suyuan32/simple-admin-tools/issues/1586))

### Chore

* upgrade etcd ([#1597](https://github.com/suyuan32/simple-admin-tools/issues/1597))
* fix data race ([#1593](https://github.com/suyuan32/simple-admin-tools/issues/1593))
* add goctl command help ([#1578](https://github.com/suyuan32/simple-admin-tools/issues/1578))
* update help message ([#1544](https://github.com/suyuan32/simple-admin-tools/issues/1544))

### Docs

* add go-zero users ([#1546](https://github.com/suyuan32/simple-admin-tools/issues/1546))
* update roadmap ([#1537](https://github.com/suyuan32/simple-admin-tools/issues/1537))

### Feat

* supports `importValue` for more path formats ([#1569](https://github.com/suyuan32/simple-admin-tools/issues/1569))
* support pg serial type for auto_increment ([#1563](https://github.com/suyuan32/simple-admin-tools/issues/1563))
* log 404 requests with traceid ([#1554](https://github.com/suyuan32/simple-admin-tools/issues/1554))
* support ctx in sql model generation ([#1551](https://github.com/suyuan32/simple-admin-tools/issues/1551))
* support ctx in sqlx/sqlc, listed in ROADMAP ([#1535](https://github.com/suyuan32/simple-admin-tools/issues/1535))

### Feature

* Add goctl env ([#1557](https://github.com/suyuan32/simple-admin-tools/issues/1557))

### Fix

* goctl api dart support `form` tag ([#1596](https://github.com/suyuan32/simple-admin-tools/issues/1596))

### Test

* add testcase for FIFO Queue in collection module ([#1589](https://github.com/suyuan32/simple-admin-tools/issues/1589))


<a name="tools/goctl/v1.3.2"></a>
## [tools/goctl/v1.3.2](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.0...tools/goctl/v1.3.2)

> 2022-02-14

### Chore

* refactor cache ([#1532](https://github.com/suyuan32/simple-admin-tools/issues/1532))
* goctl format issue ([#1531](https://github.com/suyuan32/simple-admin-tools/issues/1531))
* update goctl version to 1.3.2 ([#1524](https://github.com/suyuan32/simple-admin-tools/issues/1524))

### Feat

* support ctx in `Cache` ([#1518](https://github.com/suyuan32/simple-admin-tools/issues/1518))

### Fix

* fix a typo ([#1522](https://github.com/suyuan32/simple-admin-tools/issues/1522))


<a name="tools/goctl/v1.3.0"></a>

## [tools/goctl/v1.3.0](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.1-alpha...tools/goctl/v1.3.0)

> 2022-02-09

### Chore

* optimize yaml unmarshaler ([#1513](https://github.com/suyuan32/simple-admin-tools/issues/1513))
* make error clearer ([#1514](https://github.com/suyuan32/simple-admin-tools/issues/1514))
* update command comment ([#1501](https://github.com/suyuan32/simple-admin-tools/issues/1501))

### Ci

* add test for win ([#1503](https://github.com/suyuan32/simple-admin-tools/issues/1503))

### Feat

* update go-redis to v8, support ctx in redis methods ([#1507](https://github.com/suyuan32/simple-admin-tools/issues/1507))

### Feature

* Add `goctl completion` ([#1505](https://github.com/suyuan32/simple-admin-tools/issues/1505))

### Refactor

* refactor yaml unmarshaler ([#1517](https://github.com/suyuan32/simple-admin-tools/issues/1517))

### Test

* change fuzz tests ([#1504](https://github.com/suyuan32/simple-admin-tools/issues/1504))

<a name="tools/goctl/v1.3.1-alpha"></a>

## [tools/goctl/v1.3.1-alpha](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.1...tools/goctl/v1.3.1-alpha)

> 2022-02-01


<a name="tools/goctl/v1.3.1"></a>

## [tools/goctl/v1.3.1](https://github.com/suyuan32/simple-admin-tools/compare/v1.3.0...tools/goctl/v1.3.1)

> 2022-02-01

### Docs

* update tal-tech to zeromico in docs ([#1498](https://github.com/suyuan32/simple-admin-tools/issues/1498))

### Fix

* goctl not compile on windows ([#1500](https://github.com/suyuan32/simple-admin-tools/issues/1500))


<a name="v1.3.0"></a>
## [v1.3.0](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.0-beta1...v1.3.0)

> 2022-02-01

### Chore

* update goctl version ([#1497](https://github.com/suyuan32/simple-admin-tools/issues/1497))
* improve migrate confirmation ([#1488](https://github.com/suyuan32/simple-admin-tools/issues/1488))

### Feat

* add runtime stats monitor ([#1496](https://github.com/suyuan32/simple-admin-tools/issues/1496))
* handling panic in mapreduce, panic in calling goroutine, not inside goroutines ([#1490](https://github.com/suyuan32/simple-admin-tools/issues/1490))

### Fix

* goroutine stuck on edge case ([#1495](https://github.com/suyuan32/simple-admin-tools/issues/1495))


<a name="tools/goctl/v1.3.0-beta1"></a>
## [tools/goctl/v1.3.0-beta1](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.3.0-alpha...tools/goctl/v1.3.0-beta1)

> 2022-01-26

### Chore

* update warning message ([#1487](https://github.com/suyuan32/simple-admin-tools/issues/1487))
* update go version for goctl ([#1484](https://github.com/suyuan32/simple-admin-tools/issues/1484))

### Patch

* goctl migrate ([#1485](https://github.com/suyuan32/simple-admin-tools/issues/1485))


<a name="tools/goctl/v1.3.0-alpha"></a>
## [tools/goctl/v1.3.0-alpha](https://github.com/suyuan32/simple-admin-tools/compare/v1.3.0-beta...tools/goctl/v1.3.0-alpha)

> 2022-01-25

### Refactor

* rename from tal-tech to zeromicro for goctl ([#1481](https://github.com/suyuan32/simple-admin-tools/issues/1481))


<a name="v1.3.0-beta"></a>
## [v1.3.0-beta](https://github.com/suyuan32/simple-admin-tools/compare/v1.3.0-alpha...v1.3.0-beta)

> 2022-01-25

### Chore

* optimize string search with Aho–Corasick algorithm ([#1476](https://github.com/suyuan32/simple-admin-tools/issues/1476))
* update unauthorized callback calling order ([#1469](https://github.com/suyuan32/simple-admin-tools/issues/1469))
* check interface satisfaction w/o allocating new variable ([#1454](https://github.com/suyuan32/simple-admin-tools/issues/1454))
* remove jwt deprecated ([#1452](https://github.com/suyuan32/simple-admin-tools/issues/1452))
* upgrade dependencies ([#1444](https://github.com/suyuan32/simple-admin-tools/issues/1444))
* fix typo ([#1437](https://github.com/suyuan32/simple-admin-tools/issues/1437))
* refactor periodlimit ([#1428](https://github.com/suyuan32/simple-admin-tools/issues/1428))

### Ci

* add translator action ([#1441](https://github.com/suyuan32/simple-admin-tools/issues/1441))

### Docs

* add go-zero users ([#1473](https://github.com/suyuan32/simple-admin-tools/issues/1473))
* add go-zero users ([#1425](https://github.com/suyuan32/simple-admin-tools/issues/1425))
* add go-zero users ([#1424](https://github.com/suyuan32/simple-admin-tools/issues/1424))
* update install readme ([#1417](https://github.com/suyuan32/simple-admin-tools/issues/1417))

### Feat

* implement console plain output for debug logs ([#1456](https://github.com/suyuan32/simple-admin-tools/issues/1456))
* 支持redis的LTrim方法 ([#1443](https://github.com/suyuan32/simple-admin-tools/issues/1443))
* Add migrate ([#1419](https://github.com/suyuan32/simple-admin-tools/issues/1419))

### Fix

* mr goroutine leak on context deadline ([#1433](https://github.com/suyuan32/simple-admin-tools/issues/1433))
* golint issue ([#1423](https://github.com/suyuan32/simple-admin-tools/issues/1423))

### Patch

* save missing templates to disk ([#1463](https://github.com/suyuan32/simple-admin-tools/issues/1463))


<a name="v1.3.0-alpha"></a>
## [v1.3.0-alpha](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.2.5...v1.3.0-alpha)

> 2022-01-05

### Chore

* refactor rest/timeouthandler ([#1415](https://github.com/suyuan32/simple-admin-tools/issues/1415))

### Feat

* rename module from tal-tech to zeromicro ([#1413](https://github.com/suyuan32/simple-admin-tools/issues/1413))


<a name="tools/goctl/v1.2.5"></a>
## [tools/goctl/v1.2.5](https://github.com/suyuan32/simple-admin-tools/compare/v1.2.5...tools/goctl/v1.2.5)

> 2022-01-03

### Chore

* update go-zero to v1.2.5 ([#1410](https://github.com/suyuan32/simple-admin-tools/issues/1410))


<a name="v1.2.5"></a>
## [v1.2.5](https://github.com/suyuan32/simple-admin-tools/compare/tools/goctl/v1.2.4...v1.2.5)

> 2022-01-02

### Chore

* simplify mapreduce ([#1401](https://github.com/suyuan32/simple-admin-tools/issues/1401))
* fix golint issues ([#1396](https://github.com/suyuan32/simple-admin-tools/issues/1396))
* update goctl version ([#1394](https://github.com/suyuan32/simple-admin-tools/issues/1394))

### Ci

* remove 386 binaries ([#1393](https://github.com/suyuan32/simple-admin-tools/issues/1393))

### Docs

* update roadmap ([#1405](https://github.com/suyuan32/simple-admin-tools/issues/1405))
* update goctl installation command ([#1403](https://github.com/suyuan32/simple-admin-tools/issues/1403))

### Feat

* support tls for etcd client ([#1390](https://github.com/suyuan32/simple-admin-tools/issues/1390))
* implement fx.NoneMatch, fx.First, fx.Last ([#1402](https://github.com/suyuan32/simple-admin-tools/issues/1402))

### Refactor

* optimize fx ([#1404](https://github.com/suyuan32/simple-admin-tools/issues/1404))


<a name="tools/goctl/v1.2.4"></a>
## [tools/goctl/v1.2.4](https://github.com/suyuan32/simple-admin-tools/compare/v1.2.4...tools/goctl/v1.2.4)

> 2021-12-30

### Chore

* update goctl version to 1.2.4 for release tools/goctl/v1.2.4 ([#1372](https://github.com/suyuan32/simple-admin-tools/issues/1372))
* add 1s for tolerance in redislock ([#1367](https://github.com/suyuan32/simple-admin-tools/issues/1367))
* coding style and comments ([#1361](https://github.com/suyuan32/simple-admin-tools/issues/1361))
* optimize `ParseJsonBody` ([#1353](https://github.com/suyuan32/simple-admin-tools/issues/1353))
* put error message in error.log for verbose mode ([#1355](https://github.com/suyuan32/simple-admin-tools/issues/1355))
* add tests & refactor ([#1346](https://github.com/suyuan32/simple-admin-tools/issues/1346))
* add comments ([#1345](https://github.com/suyuan32/simple-admin-tools/issues/1345))
* update goctl version to 1.2.5 ([#1337](https://github.com/suyuan32/simple-admin-tools/issues/1337))
* refactor ([#1331](https://github.com/suyuan32/simple-admin-tools/issues/1331))
* format code ([#1327](https://github.com/suyuan32/simple-admin-tools/issues/1327))
* rename service context from ctx to svcCtx ([#1299](https://github.com/suyuan32/simple-admin-tools/issues/1299))
* update cli version ([#1287](https://github.com/suyuan32/simple-admin-tools/issues/1287))

### Chose

* cancel the assignment and judge later ([#1359](https://github.com/suyuan32/simple-admin-tools/issues/1359))

### Ci

* remove windows 386 binary ([#1392](https://github.com/suyuan32/simple-admin-tools/issues/1392))
* add release action to auto build binaries ([#1371](https://github.com/suyuan32/simple-admin-tools/issues/1371))

### Docs

* add go-zero users ([#1381](https://github.com/suyuan32/simple-admin-tools/issues/1381))
* update slack invitation link ([#1378](https://github.com/suyuan32/simple-admin-tools/issues/1378))
* update goctl markdown ([#1370](https://github.com/suyuan32/simple-admin-tools/issues/1370))
* add go-zero users ([#1323](https://github.com/suyuan32/simple-admin-tools/issues/1323))
* add go-zero users ([#1294](https://github.com/suyuan32/simple-admin-tools/issues/1294))

### Feat

* Add --remote ([#1387](https://github.com/suyuan32/simple-admin-tools/issues/1387))
* support array in default and options tags ([#1386](https://github.com/suyuan32/simple-admin-tools/issues/1386))
* support context in MapReduce ([#1368](https://github.com/suyuan32/simple-admin-tools/issues/1368))
* treat client closed requests as code 499 ([#1350](https://github.com/suyuan32/simple-admin-tools/issues/1350))
* tidy mod, update go-zero to latest ([#1334](https://github.com/suyuan32/simple-admin-tools/issues/1334))
* tidy mod, update go-zero to latest ([#1333](https://github.com/suyuan32/simple-admin-tools/issues/1333))
* tidy mod, add go.mod for goctl ([#1328](https://github.com/suyuan32/simple-admin-tools/issues/1328))
* reduce dependencies of framework by add go.mod in goctl ([#1290](https://github.com/suyuan32/simple-admin-tools/issues/1290))

### Feature

* support adding custom cache to mongoc and sqlc ([#1313](https://github.com/suyuan32/simple-admin-tools/issues/1313))

### Fix

*  command system info missing go version ([#1377](https://github.com/suyuan32/simple-admin-tools/issues/1377))
* [#1318](https://github.com/suyuan32/simple-admin-tools/issues/1318) ([#1321](https://github.com/suyuan32/simple-admin-tools/issues/1321))
* go issue 16206 ([#1298](https://github.com/suyuan32/simple-admin-tools/issues/1298))

### Style

* format code ([#1322](https://github.com/suyuan32/simple-admin-tools/issues/1322))

### Test

* add more tests ([#1391](https://github.com/suyuan32/simple-admin-tools/issues/1391))
* add more tests ([#1352](https://github.com/suyuan32/simple-admin-tools/issues/1352))


<a name="v1.2.4"></a>
## [v1.2.4](https://github.com/suyuan32/simple-admin-tools/compare/v1.2.3...v1.2.4)

> 2021-12-01

### Chore

* cleanup zRPC retry code ([#1280](https://github.com/suyuan32/simple-admin-tools/issues/1280))
* only allow cors middleware to change headers ([#1276](https://github.com/suyuan32/simple-admin-tools/issues/1276))
* avoid superfluous WriteHeader call errors ([#1275](https://github.com/suyuan32/simple-admin-tools/issues/1275))
* update goctl version ([#1250](https://github.com/suyuan32/simple-admin-tools/issues/1250))

### Docs

* update readme to use goctl[@cli](https://github.com/cli) ([#1255](https://github.com/suyuan32/simple-admin-tools/issues/1255))

### Feat

* support third party orm to interact with go-zero ([#1286](https://github.com/suyuan32/simple-admin-tools/issues/1286))
* add etcd resolver scheme, fix discov minor issue ([#1281](https://github.com/suyuan32/simple-admin-tools/issues/1281))
* support %w in logx.Errorf ([#1278](https://github.com/suyuan32/simple-admin-tools/issues/1278))
* add rest.WithCustomCors to let caller customize the response ([#1274](https://github.com/suyuan32/simple-admin-tools/issues/1274))


<a name="v1.2.3"></a>
## [v1.2.3](https://github.com/suyuan32/simple-admin-tools/compare/v1.2.2...v1.2.3)

> 2021-11-15

### Chore

* remove conf.CheckedDuration ([#1235](https://github.com/suyuan32/simple-admin-tools/issues/1235))
* refactor, better goctl message ([#1228](https://github.com/suyuan32/simple-admin-tools/issues/1228))
* remove unused const ([#1224](https://github.com/suyuan32/simple-admin-tools/issues/1224))
* redislock use stringx.randn replace randomStr func ([#1220](https://github.com/suyuan32/simple-admin-tools/issues/1220))
* refine code ([#1215](https://github.com/suyuan32/simple-admin-tools/issues/1215))
* remove semicolon for routes of services in api files ([#1195](https://github.com/suyuan32/simple-admin-tools/issues/1195))
* update goctl version to 1.2.3, prepare for release ([#1193](https://github.com/suyuan32/simple-admin-tools/issues/1193))
* reorg imports, format code, make MaxRetires default to 0 ([#1165](https://github.com/suyuan32/simple-admin-tools/issues/1165))
* reverse the order of stopping services ([#1159](https://github.com/suyuan32/simple-admin-tools/issues/1159))

### Docs

* add go-zero users ([#1214](https://github.com/suyuan32/simple-admin-tools/issues/1214))
* update roadmap ([#1184](https://github.com/suyuan32/simple-admin-tools/issues/1184))
* update roadmap ([#1178](https://github.com/suyuan32/simple-admin-tools/issues/1178))
* add go-zero users ([#1176](https://github.com/suyuan32/simple-admin-tools/issues/1176))
* add go-zero users ([#1172](https://github.com/suyuan32/simple-admin-tools/issues/1172))
* update qr code ([#1158](https://github.com/suyuan32/simple-admin-tools/issues/1158))
* add go-zero users ([#1141](https://github.com/suyuan32/simple-admin-tools/issues/1141))
* add go-zero users ([#1135](https://github.com/suyuan32/simple-admin-tools/issues/1135))
* add go-zero users ([#1130](https://github.com/suyuan32/simple-admin-tools/issues/1130))

### Feat

* enable retry for zrpc ([#1237](https://github.com/suyuan32/simple-admin-tools/issues/1237))
* disable grpc retry, enable it in v1.2.4 ([#1233](https://github.com/suyuan32/simple-admin-tools/issues/1233))
* exit with non-zero code on errors ([#1218](https://github.com/suyuan32/simple-admin-tools/issues/1218))
* support CORS, better implementation ([#1217](https://github.com/suyuan32/simple-admin-tools/issues/1217))
* support CORS by using rest.WithCors(...) ([#1212](https://github.com/suyuan32/simple-admin-tools/issues/1212))
* ignore rest.WithPrefix on empty prefix ([#1208](https://github.com/suyuan32/simple-admin-tools/issues/1208))
* support customizing timeout for specific route ([#1203](https://github.com/suyuan32/simple-admin-tools/issues/1203))
* add NewSessionFromTx to interact with other orm ([#1202](https://github.com/suyuan32/simple-admin-tools/issues/1202))
* simplify the grpc tls authentication ([#1199](https://github.com/suyuan32/simple-admin-tools/issues/1199))
* use WithBlock() by default, NonBlock can be set in config or WithNonBlock() ([#1198](https://github.com/suyuan32/simple-admin-tools/issues/1198))
* add rest.WithPrefix to support route prefix ([#1194](https://github.com/suyuan32/simple-admin-tools/issues/1194))
* slow threshold customizable in zrpc ([#1191](https://github.com/suyuan32/simple-admin-tools/issues/1191))
* slow threshold customizable in rest ([#1189](https://github.com/suyuan32/simple-admin-tools/issues/1189))
* slow threshold customizable in sqlx ([#1188](https://github.com/suyuan32/simple-admin-tools/issues/1188))
* slow threshold customizable in redis ([#1187](https://github.com/suyuan32/simple-admin-tools/issues/1187))
* slow threshold customizable in mongo ([#1186](https://github.com/suyuan32/simple-admin-tools/issues/1186))
* slow threshold customizable in redis ([#1185](https://github.com/suyuan32/simple-admin-tools/issues/1185))
* support multiple trace agents ([#1183](https://github.com/suyuan32/simple-admin-tools/issues/1183))
* let different services start prometheus on demand ([#1182](https://github.com/suyuan32/simple-admin-tools/issues/1182))
* support auth account for etcd ([#1174](https://github.com/suyuan32/simple-admin-tools/issues/1174))
* support ssl on zrpc, simplify the config ([#1175](https://github.com/suyuan32/simple-admin-tools/issues/1175))

### Refactor

* simplify tls config in rest ([#1181](https://github.com/suyuan32/simple-admin-tools/issues/1181))

### Test

* add more tests ([#1209](https://github.com/suyuan32/simple-admin-tools/issues/1209))
* add more tests ([#1179](https://github.com/suyuan32/simple-admin-tools/issues/1179))
* add more tests ([#1166](https://github.com/suyuan32/simple-admin-tools/issues/1166))
* add more tests ([#1163](https://github.com/suyuan32/simple-admin-tools/issues/1163))
* add more tests ([#1154](https://github.com/suyuan32/simple-admin-tools/issues/1154))
* add more tests ([#1150](https://github.com/suyuan32/simple-admin-tools/issues/1150))
* add more tests ([#1149](https://github.com/suyuan32/simple-admin-tools/issues/1149))
* add more tests ([#1147](https://github.com/suyuan32/simple-admin-tools/issues/1147))
* add more tests ([#1138](https://github.com/suyuan32/simple-admin-tools/issues/1138))
* add more tests ([#1137](https://github.com/suyuan32/simple-admin-tools/issues/1137))
* add more tests ([#1134](https://github.com/suyuan32/simple-admin-tools/issues/1134))


<a name="v1.2.2"></a>
## [v1.2.2](https://github.com/suyuan32/simple-admin-tools/compare/v1.2.1...v1.2.2)

> 2021-10-12

### Chore

* refine rpc template in goctl ([#1129](https://github.com/suyuan32/simple-admin-tools/issues/1129))
* replace redis.NewRedis with redis.New ([#1103](https://github.com/suyuan32/simple-admin-tools/issues/1103))
* mark redis.NewRedis as Deprecated, use redis.New instead. ([#1100](https://github.com/suyuan32/simple-admin-tools/issues/1100))
* run unit test with go 1.14 ([#1084](https://github.com/suyuan32/simple-admin-tools/issues/1084))
* when run goctl-rpc, the order of proto message aliases should be ([#1078](https://github.com/suyuan32/simple-admin-tools/issues/1078))
* fix comment issues ([#1056](https://github.com/suyuan32/simple-admin-tools/issues/1056))
* make comment accurate ([#1055](https://github.com/suyuan32/simple-admin-tools/issues/1055))

### Ci

* add reviewdog ([#1096](https://github.com/suyuan32/simple-admin-tools/issues/1096))
* accurate error reporting on lint check ([#1089](https://github.com/suyuan32/simple-admin-tools/issues/1089))
* add Lint check on commits ([#1086](https://github.com/suyuan32/simple-admin-tools/issues/1086))

### Docs

* update roadmap ([#1110](https://github.com/suyuan32/simple-admin-tools/issues/1110))
* update roadmap ([#1094](https://github.com/suyuan32/simple-admin-tools/issues/1094))
* update roadmap ([#1093](https://github.com/suyuan32/simple-admin-tools/issues/1093))
* change organization from tal-tech to zeromicro in readme ([#1087](https://github.com/suyuan32/simple-admin-tools/issues/1087))

### Feat

* opentelemetry integration, removed self designed tracing ([#1111](https://github.com/suyuan32/simple-admin-tools/issues/1111))
* reflection grpc service ([#1107](https://github.com/suyuan32/simple-admin-tools/issues/1107))

### Fix

* opentelemetry traceid not correct ([#1108](https://github.com/suyuan32/simple-admin-tools/issues/1108))

### Test

* add more tests ([#1119](https://github.com/suyuan32/simple-admin-tools/issues/1119))
* add more tests ([#1112](https://github.com/suyuan32/simple-admin-tools/issues/1112))
* add more tests ([#1106](https://github.com/suyuan32/simple-admin-tools/issues/1106))


<a name="v1.2.1"></a>
## [v1.2.1](https://github.com/suyuan32/simple-admin-tools/compare/v1.2.0...v1.2.1)

> 2021-09-14


<a name="v1.2.0"></a>
## [v1.2.0](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.10...v1.2.0)

> 2021-09-13

### Feat

* change logger to traceLogger for getting traceId when recovering ([#374](https://github.com/suyuan32/simple-admin-tools/issues/374))

### Fix

* 解决golint 部分警告 ([#897](https://github.com/suyuan32/simple-admin-tools/issues/897))

### Pull Requests

* Merge pull request [#1023](https://github.com/suyuan32/simple-admin-tools/issues/1023) from anqiansong/1014-rollback


<a name="v1.1.10"></a>
## [v1.1.10](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.9...v1.1.10)

> 2021-08-04


<a name="v1.1.9"></a>
## [v1.1.9](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.8...v1.1.9)

> 2021-08-01


<a name="v1.1.8"></a>
## [v1.1.8](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.7...v1.1.8)

> 2021-06-23

### Fix

* Fix problems with non support for multidimensional arrays and basic type pointer arrays ([#778](https://github.com/suyuan32/simple-admin-tools/issues/778))


<a name="v1.1.7"></a>
## [v1.1.7](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.6...v1.1.7)

> 2021-05-08

### Chore

* update code format. ([#628](https://github.com/suyuan32/simple-admin-tools/issues/628))

### Doc

* fix spell mistake ([#627](https://github.com/suyuan32/simple-admin-tools/issues/627))


<a name="v1.1.6"></a>
## [v1.1.6](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.5...v1.1.6)

> 2021-03-27


<a name="v1.1.5"></a>
## [v1.1.5](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.4...v1.1.5)

> 2021-03-02


<a name="v1.1.4"></a>
## [v1.1.4](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.3-pre...v1.1.4)

> 2021-01-16


<a name="v1.1.3-pre"></a>
## [v1.1.3-pre](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.3-beta...v1.1.3-pre)

> 2021-01-13


<a name="v1.1.3-beta"></a>
## [v1.1.3-beta](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.3...v1.1.3-beta)

> 2021-01-13


<a name="v1.1.3"></a>
## [v1.1.3](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.2...v1.1.3)

> 2021-01-13

### Feature

* refactor api parse to g4 ([#365](https://github.com/suyuan32/simple-admin-tools/issues/365))


<a name="v1.1.2"></a>
## [v1.1.2](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.1...v1.1.2)

> 2021-01-05


<a name="v1.1.1"></a>
## [v1.1.1](https://github.com/suyuan32/simple-admin-tools/compare/v1.1.0...v1.1.1)

> 2020-12-12


<a name="v1.1.0"></a>
## [v1.1.0](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.29...v1.1.0)

> 2020-12-09


<a name="v1.0.29"></a>
## [v1.0.29](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.28...v1.0.29)

> 2020-11-28

### Feature

* file namestyle ([#223](https://github.com/suyuan32/simple-admin-tools/issues/223))


<a name="v1.0.28"></a>
## [v1.0.28](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.27...v1.0.28)

> 2020-11-19


<a name="v1.0.27"></a>
## [v1.0.27](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.26...v1.0.27)

> 2020-11-13


<a name="v1.0.26"></a>
## [v1.0.26](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.25...v1.0.26)

> 2020-11-08


<a name="v1.0.25"></a>
## [v1.0.25](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.24...v1.0.25)

> 2020-11-05


<a name="v1.0.24"></a>
## [v1.0.24](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.23...v1.0.24)

> 2020-11-05


<a name="v1.0.23"></a>
## [v1.0.23](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.22...v1.0.23)

> 2020-10-28

### Docs

* format markdown and add go mod in demo ([#155](https://github.com/suyuan32/simple-admin-tools/issues/155))


<a name="v1.0.22"></a>
## [v1.0.22](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.21...v1.0.22)

> 2020-10-21


<a name="v1.0.21"></a>
## [v1.0.21](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.20...v1.0.21)

> 2020-10-17

### Fix

* fx/fn.Head func will forever block when n is less than 1 ([#128](https://github.com/suyuan32/simple-admin-tools/issues/128))
* template cache key ([#121](https://github.com/suyuan32/simple-admin-tools/issues/121))


<a name="v1.0.20"></a>
## [v1.0.20](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.19...v1.0.20)

> 2020-10-10


<a name="v1.0.19"></a>
## [v1.0.19](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.18...v1.0.19)

> 2020-10-04

### Breaker

* remover useless code ([#114](https://github.com/suyuan32/simple-admin-tools/issues/114))

### Update

* fix wrong word ([#110](https://github.com/suyuan32/simple-admin-tools/issues/110))


<a name="v1.0.18"></a>
## [v1.0.18](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.17...v1.0.18)

> 2020-09-29

### Doc

* update sharedcalls.md layout ([#107](https://github.com/suyuan32/simple-admin-tools/issues/107))

### Reverts

* goreportcard not working, remove it temporarily


<a name="v1.0.17"></a>
## [v1.0.17](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.16...v1.0.17)

> 2020-09-28


<a name="v1.0.16"></a>
## [v1.0.16](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.15...v1.0.16)

> 2020-09-22

### Chore

* fix typos ([#85](https://github.com/suyuan32/simple-admin-tools/issues/85))

### Feature

* goctl jwt  ([#91](https://github.com/suyuan32/simple-admin-tools/issues/91))

### Fix

* golint: context.WithValue should should not use basic type as key ([#83](https://github.com/suyuan32/simple-admin-tools/issues/83))


<a name="v1.0.15"></a>
## [v1.0.15](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.14...v1.0.15)

> 2020-09-18


<a name="v1.0.14"></a>
## [v1.0.14](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.13...v1.0.14)

> 2020-09-16

### Optimize

* api generating for idea plugin ([#68](https://github.com/suyuan32/simple-admin-tools/issues/68))


<a name="v1.0.13"></a>
## [v1.0.13](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.12...v1.0.13)

> 2020-09-11


<a name="v1.0.12"></a>
## [v1.0.12](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.11...v1.0.12)

> 2020-09-08


<a name="v1.0.11"></a>
## [v1.0.11](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.10...v1.0.11)

> 2020-09-03


<a name="v1.0.10"></a>
## [v1.0.10](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.9...v1.0.10)

> 2020-09-03


<a name="v1.0.9"></a>
## [v1.0.9](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.8...v1.0.9)

> 2020-09-02

### Fix

* root path on windows bug ([#34](https://github.com/suyuan32/simple-admin-tools/issues/34))


<a name="v1.0.8"></a>
## [v1.0.8](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.7...v1.0.8)

> 2020-09-01


<a name="v1.0.7"></a>
## [v1.0.7](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.6...v1.0.7)

> 2020-08-29


<a name="v1.0.6"></a>
## [v1.0.6](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.5...v1.0.6)

> 2020-08-25


<a name="v1.0.5"></a>
## [v1.0.5](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.4...v1.0.5)

> 2020-08-24


<a name="v1.0.4"></a>
## [v1.0.4](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.3...v1.0.4)

> 2020-08-19

### Chore

* fix typo


<a name="v1.0.3"></a>
## [v1.0.3](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.2...v1.0.3)

> 2020-08-14


<a name="v1.0.2"></a>
## [v1.0.2](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.1...v1.0.2)

> 2020-08-14


<a name="v1.0.1"></a>
## [v1.0.1](https://github.com/suyuan32/simple-admin-tools/compare/v1.0.0...v1.0.1)

> 2020-08-13


<a name="v1.0.0"></a>
## v1.0.0

> 2020-08-11

