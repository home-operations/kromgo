# Changelog

## [0.16.1](https://github.com/home-operations/kromgo/compare/0.16.0...0.16.1) (2026-08-22)


### Miscellaneous Chores

* **github-action:** update action docker/setup-buildx-action (v4.2.0 → v4.3.0) ([#376](https://github.com/home-operations/kromgo/issues/376)) ([81173a6](https://github.com/home-operations/kromgo/commit/81173a6eebd0173d112971fb5b82ee5e9d8862e2))
* **mise:** update tool oxfmt (0.63.0 → 0.64.0) ([#373](https://github.com/home-operations/kromgo/issues/373)) ([fb575c5](https://github.com/home-operations/kromgo/commit/fb575c520ca33883d7f6314e39e47c839c8863c8))
* **mise:** update tool yq (4.53.3 → 4.53.4) ([#375](https://github.com/home-operations/kromgo/issues/375)) ([29ee153](https://github.com/home-operations/kromgo/commit/29ee153fa1b780d1a7d52fab16eb82dc309bd09c))

## [0.16.0](https://github.com/home-operations/kromgo/compare/0.15.2...0.16.0) (2026-08-21)


### ⚠ BREAKING CHANGES

* **go:** Update module golang.org/x/image (v0.44.0 → v0.45.0) ([#357](https://github.com/home-operations/kromgo/issues/357))
* **go:** Update module github.com/google/cel-go (v0.30.0 → v0.31.0) ([#354](https://github.com/home-operations/kromgo/issues/354))

### Features

* **go:** update module github.com/google/cel-go (v0.29.2 → v0.30.0) ([#326](https://github.com/home-operations/kromgo/issues/326)) ([6877542](https://github.com/home-operations/kromgo/commit/687754256f87f952c3454b22b5009778ad32b73c))
* **go:** Update module github.com/google/cel-go (v0.30.0 → v0.31.0) ([#354](https://github.com/home-operations/kromgo/issues/354)) ([d0f45d5](https://github.com/home-operations/kromgo/commit/d0f45d5520522a528fdd5e954fca20a78d2d593e))
* **go:** update module github.com/google/cel-go (v0.31.0 → v0.32.0) ([#371](https://github.com/home-operations/kromgo/issues/371)) ([512b6fe](https://github.com/home-operations/kromgo/commit/512b6fe1d5fda0a7adb51926aa45a9f93f31a69d))
* **go:** update module github.com/stretchr/testify (v1.11.1 → v1.12.0) ([#367](https://github.com/home-operations/kromgo/issues/367)) ([688a29f](https://github.com/home-operations/kromgo/commit/688a29fd48479582f8a77f115fedbd6c9275e8eb))
* **go:** Update module golang.org/x/image (v0.44.0 → v0.45.0) ([#357](https://github.com/home-operations/kromgo/issues/357)) ([a9e7fc6](https://github.com/home-operations/kromgo/commit/a9e7fc6bb7c731babae8265759d72892115883fb))
* **npm:** update dependency simple-icons (16.27.1 → 16.28.0) ([#344](https://github.com/home-operations/kromgo/issues/344)) ([d7037f8](https://github.com/home-operations/kromgo/commit/d7037f836767978f2192d0b06144425396c9200f))


### Bug Fixes

* **chart:** keep example comments out of unrelated schema descriptions ([#369](https://github.com/home-operations/kromgo/issues/369)) ([0c23582](https://github.com/home-operations/kromgo/commit/0c2358225136e9a20719c0bf22c01029408946f7))
* **ci:** fail the merge gate on cancelled jobs, and key the lint cache on the toolchain ([#337](https://github.com/home-operations/kromgo/issues/337)) ([b48baa0](https://github.com/home-operations/kromgo/commit/b48baa00c56a69d18cfd2fb6e446227f8b567c3c))
* **go:** update module github.com/go-analyze/charts (v0.6.0 → v0.6.1) ([#361](https://github.com/home-operations/kromgo/issues/361)) ([796a7e5](https://github.com/home-operations/kromgo/commit/796a7e57c392c45155ce41a58f40e3e47c02abc7))
* **go:** update module github.com/stretchr/testify (v1.12.0 → v1.12.1) ([#370](https://github.com/home-operations/kromgo/issues/370)) ([b08aa8f](https://github.com/home-operations/kromgo/commit/b08aa8fbe4f569f7197a5a444df29dc699f8da35))
* **go:** update module go (1.26.0 → 1.26.5) ([#349](https://github.com/home-operations/kromgo/issues/349)) ([9c944f6](https://github.com/home-operations/kromgo/commit/9c944f64cd113bb83c9a877ae46936f3fbbe2de0))
* **go:** update to go 1.27.0 ([#372](https://github.com/home-operations/kromgo/issues/372)) ([f4dc8e7](https://github.com/home-operations/kromgo/commit/f4dc8e7791aceff3fa153222c66d7b0d4f1f10e9))
* **npm:** update dependency marked (18.0.7 → 18.0.9) ([#350](https://github.com/home-operations/kromgo/issues/350)) ([a7a1df8](https://github.com/home-operations/kromgo/commit/a7a1df86be7de139b1d2dabe26b9db09a28e20c8))
* **npm:** update dependency marked (18.0.9 → 18.0.10) ([#368](https://github.com/home-operations/kromgo/issues/368)) ([103e1dd](https://github.com/home-operations/kromgo/commit/103e1dd046f14b4bff6b3678d30900f6addb49c7))
* **npm:** update dependency simple-icons (16.27.0 → 16.27.1) ([#325](https://github.com/home-operations/kromgo/issues/325)) ([b29f5c3](https://github.com/home-operations/kromgo/commit/b29f5c38e2fcd7407ef47a39815c08a1c7816cad))
* tidy ci.yaml after the Build Success consolidation ([#322](https://github.com/home-operations/kromgo/issues/322)) ([a8f4b6f](https://github.com/home-operations/kromgo/commit/a8f4b6f234975a7d9d444df82df1b5967e020edc))


### Documentation

* add AGENTS.md with Go conventions ([#341](https://github.com/home-operations/kromgo/issues/341)) ([4edef5e](https://github.com/home-operations/kromgo/commit/4edef5ec32175863c48da154b154be9c266d0979))
* point the CI badge at ci.yaml ([#324](https://github.com/home-operations/kromgo/issues/324)) ([ebaf5ce](https://github.com/home-operations/kromgo/commit/ebaf5ced3ace973dbe481ec33bbc4f1dabfdf6d2))


### Build System

* **mise:** add actionlint and refresh the lockfile ([#327](https://github.com/home-operations/kromgo/issues/327)) ([a2a15cb](https://github.com/home-operations/kromgo/commit/a2a15cbb244dab1fd8d7da981fd0e0e2ab56ccdc))


### Continuous Integration

* consolidate pull request checks behind Build Success ([#321](https://github.com/home-operations/kromgo/issues/321)) ([c389060](https://github.com/home-operations/kromgo/commit/c389060ee58a5752183e46a24fe0134b96621f62))
* **github-action:** Update action actions/stale (v10.4.0 → v11.0.0) ([#338](https://github.com/home-operations/kromgo/issues/338)) ([13afb19](https://github.com/home-operations/kromgo/commit/13afb1983365793ae9081be7a81a7aa77c9418a0))
* **github-action:** Update action docker/github-builder (v1.15.0 → v1.16.0) ([#355](https://github.com/home-operations/kromgo/issues/355)) ([fa33fa0](https://github.com/home-operations/kromgo/commit/fa33fa09048f50ecb8ba0d77ec2e7d378ff7a5f9))
* **github-action:** Update action docker/login-action (v4.5.0 → v4.5.1) ([#330](https://github.com/home-operations/kromgo/issues/330)) ([7fc44ec](https://github.com/home-operations/kromgo/commit/7fc44ec0bb7afbe6d40c9d2cf6bb6590e15b0488))
* **github-action:** Update action docker/login-action (v4.5.2 → v4.6.0) ([#342](https://github.com/home-operations/kromgo/issues/342)) ([b64dabc](https://github.com/home-operations/kromgo/commit/b64dabc291ddfae247339444eca55e66b1ef7b59))
* **github-action:** Update action home-operations/.github/actions/workflow-lint (v1.0.2 → v1.0.3) ([#348](https://github.com/home-operations/kromgo/issues/348)) ([ccaa0be](https://github.com/home-operations/kromgo/commit/ccaa0be5d92cb457dbe1edac99067c15710dd338))
* **github-action:** Update action jdx/mise-action (v4.2.1 → v4.2.2) ([#329](https://github.com/home-operations/kromgo/issues/329)) ([ecb0435](https://github.com/home-operations/kromgo/commit/ecb04359690ab992cefdde72acfdedc8add065ee))
* **github-action:** Update action jdx/mise-action (v4.2.2 → v4.2.3) ([#332](https://github.com/home-operations/kromgo/issues/332)) ([242a7d8](https://github.com/home-operations/kromgo/commit/242a7d8b408be1fbe81109eafff670ec33b4606d))
* **github-action:** Update action jdx/mise-action (v4.2.3 → v4.2.4) ([#351](https://github.com/home-operations/kromgo/issues/351)) ([fef8f06](https://github.com/home-operations/kromgo/commit/fef8f06d160ee8bf75b793db1d82ef8e97c7c6a8))
* **github-action:** Update github-actions ([#336](https://github.com/home-operations/kromgo/issues/336)) ([4dcf08f](https://github.com/home-operations/kromgo/commit/4dcf08ff055645f60b9dea631203059f49ae2479))
* **github-action:** update workflow-lint action (1.0.0 → v1.0.2) ([#345](https://github.com/home-operations/kromgo/issues/345)) ([fec8eeb](https://github.com/home-operations/kromgo/commit/fec8eeb17921f90c92bda97a8e74ac60f97fa858))
* lint workflows with the shared composite action ([#328](https://github.com/home-operations/kromgo/issues/328)) ([ba79817](https://github.com/home-operations/kromgo/commit/ba798170fcfdbd579df2225f4ce0744a0e0d734e))
* skip release-please version-bump PRs in checks ([#320](https://github.com/home-operations/kromgo/issues/320)) ([d5d0bb2](https://github.com/home-operations/kromgo/commit/d5d0bb24be7958cd4e315a471543b746b7856a2e))
* wire govulncheck into mise and CI ([#347](https://github.com/home-operations/kromgo/issues/347)) ([0624864](https://github.com/home-operations/kromgo/commit/06248646e6a3c0dfac580364772e3d29e433cd3e))


### Miscellaneous Chores

* **github-action:** update action jdx/mise-action (v4.2.4 → v4.2.5) ([#362](https://github.com/home-operations/kromgo/issues/362)) ([e3d0dc1](https://github.com/home-operations/kromgo/commit/e3d0dc1d19bf9391c4a89fc73f298f88a118639f))
* **go:** pin go directive to 1.26.0 ([#363](https://github.com/home-operations/kromgo/issues/363)) ([eac4699](https://github.com/home-operations/kromgo/commit/eac4699b923ae88b7ef80f541c5f25eb54e72e4d))
* **mise:** Lock file maintenance tool (mise) ([#339](https://github.com/home-operations/kromgo/issues/339)) ([5d8fb93](https://github.com/home-operations/kromgo/commit/5d8fb93e97718dd46d657473b0fb2701af5c241f))
* **mise:** prune lockfile to used platforms ([#346](https://github.com/home-operations/kromgo/issues/346)) ([0eef657](https://github.com/home-operations/kromgo/commit/0eef657c1ec6c7295673c5bdc2769234a6abfe39))
* **mise:** Update tool cosign (3.1.2 → 3.1.3) ([#356](https://github.com/home-operations/kromgo/issues/356)) ([481c7c7](https://github.com/home-operations/kromgo/commit/481c7c7de6dd863311b15a75cb9eea50397954b4))
* **mise:** update tool go (1.26.5 → 1.26.6) ([#366](https://github.com/home-operations/kromgo/issues/366)) ([39f1921](https://github.com/home-operations/kromgo/commit/39f19217ff90ccb53d400480eea19c08c4ca354c))
* **mise:** update tool go:golang.org/x/vuln/cmd/govulncheck (1.6.0 → v1.7.0) ([#360](https://github.com/home-operations/kromgo/issues/360)) ([70af0e8](https://github.com/home-operations/kromgo/commit/70af0e8fb0c16a3d86e81573ae68e5003e4e3d17))
* **mise:** update tool helm (4.2.3 → 4.2.4) ([#365](https://github.com/home-operations/kromgo/issues/365)) ([0876800](https://github.com/home-operations/kromgo/commit/08768004a98ddaa60e1e29227feb8dc92fda9535))
* **mise:** Update tool node (24.18.0 → v24.18.1) ([#334](https://github.com/home-operations/kromgo/issues/334)) ([2967139](https://github.com/home-operations/kromgo/commit/2967139bcd4958abe97e0935638020592fd5b479))
* **mise:** Update tool node (24.18.1 → v24.19.0) ([#352](https://github.com/home-operations/kromgo/issues/352)) ([c1a1fc0](https://github.com/home-operations/kromgo/commit/c1a1fc045d8440c001913a4fd19052d6748d144c))
* **mise:** Update tool npm (12.0.1 → 12.0.2) ([#335](https://github.com/home-operations/kromgo/issues/335)) ([1dd88da](https://github.com/home-operations/kromgo/commit/1dd88dad16923b31fb0d6b6e2ab9e0f187499b05))
* **mise:** Update tool oxfmt (0.60.0 → 0.61.0) ([#331](https://github.com/home-operations/kromgo/issues/331)) ([b2d39ec](https://github.com/home-operations/kromgo/commit/b2d39ecfc82d447590ee30601c7d716301efbdb5))
* **mise:** Update tool oxfmt (0.61.0 → 0.62.0) ([#353](https://github.com/home-operations/kromgo/issues/353)) ([3a540b7](https://github.com/home-operations/kromgo/commit/3a540b7fe6471fcf4cb7bcab611bf412ba40ede0))
* **mise:** Update tool oxfmt (0.62.0 → 0.63.0) ([#358](https://github.com/home-operations/kromgo/issues/358)) ([bcf2fed](https://github.com/home-operations/kromgo/commit/bcf2fed78ed4ba96a4529f3045c7f8f9f7ab117a))
* **mise:** Update tool zizmor (1.28.0 → 1.29.0) ([#343](https://github.com/home-operations/kromgo/issues/343)) ([cc99c92](https://github.com/home-operations/kromgo/commit/cc99c9212319d903539399dab00a00b529d5fe48))
* **release-please:** standardize the release pull request title pattern ([#340](https://github.com/home-operations/kromgo/issues/340)) ([a8876db](https://github.com/home-operations/kromgo/commit/a8876db9e7e785a985494c8bcb20fb9cb5cea724))
* standardize release-please changelog sections ([#333](https://github.com/home-operations/kromgo/issues/333)) ([9936a24](https://github.com/home-operations/kromgo/commit/9936a24df1dd66e089712165c4e055c31e01890c))

## [0.15.2](https://github.com/home-operations/kromgo/compare/0.15.1...0.15.2) (2026-07-24)


### Features

* **deps:** update dependency simple-icons (16.26.0 → 16.27.0) ([#308](https://github.com/home-operations/kromgo/issues/308)) ([92fe7cb](https://github.com/home-operations/kromgo/commit/92fe7cbbb8dfd486195c5468d931b5bd5574647a))
* **deps:** update module github.com/prometheus/client_golang (v1.23.2 → v1.24.0) ([#309](https://github.com/home-operations/kromgo/issues/309)) ([3632476](https://github.com/home-operations/kromgo/commit/3632476fa7c7de00c5f83ca6bb2988be25a5939d))


### Bug Fixes

* **deps:** update dependency marked (18.0.6 → 18.0.7) ([#311](https://github.com/home-operations/kromgo/issues/311)) ([0ff381b](https://github.com/home-operations/kromgo/commit/0ff381bbd353841882bd0cd4fd2cf63d36361aa5))
* **deps:** update module github.com/prometheus/client_golang (v1.24.0 → v1.24.1) ([#316](https://github.com/home-operations/kromgo/issues/316)) ([5630b18](https://github.com/home-operations/kromgo/commit/5630b18d78bb070d34adb2bd13661903752d09d0))
* **deps:** update module github.com/prometheus/common (v0.70.0 → v0.70.1) ([#314](https://github.com/home-operations/kromgo/issues/314)) ([ae06e27](https://github.com/home-operations/kromgo/commit/ae06e27ae9525ae35cf5334a7e3cb2a0983beaf7))
* **helm:** stamp Chart.yaml version on release ([#319](https://github.com/home-operations/kromgo/issues/319)) ([98a4140](https://github.com/home-operations/kromgo/commit/98a4140c804c01dc48ebdd611bc2489a35222c48))


### Styles

* indent markdown at 2 to match embedded yaml ([#310](https://github.com/home-operations/kromgo/issues/310)) ([d669ff5](https://github.com/home-operations/kromgo/commit/d669ff5471ece9d97ef8bbaef9994521e401b2d2))


### Miscellaneous Chores

* **deps:** lock file maintenance ([#315](https://github.com/home-operations/kromgo/issues/315)) ([0d69b9c](https://github.com/home-operations/kromgo/commit/0d69b9c77dfcc205ffad9d26ae5f0c3a4b5894a0))
* **github-release:** Update release helm-unittest/helm-unittest (v1.1.1 → v1.1.2) ([#318](https://github.com/home-operations/kromgo/issues/318)) ([7130e66](https://github.com/home-operations/kromgo/commit/7130e66ecc5095ff96a05fde13c31b8e93399b50))
* **mise:** Update tool cosign (3.1.1 → 3.1.2) ([#306](https://github.com/home-operations/kromgo/issues/306)) ([581806f](https://github.com/home-operations/kromgo/commit/581806fe3d6513607fd7c2f6ede5cd372cc481fd))
* **mise:** Update tool oxfmt (0.59.0 → 0.60.0) ([#313](https://github.com/home-operations/kromgo/issues/313)) ([f59ee43](https://github.com/home-operations/kromgo/commit/f59ee43384ad0a3cc841a72686ceabdb60b22eec))
* **mise:** Update tool zizmor (1.27.0 → 1.28.0) ([#312](https://github.com/home-operations/kromgo/issues/312)) ([4fe43a3](https://github.com/home-operations/kromgo/commit/4fe43a38a5a2776e39ac0b1857bcd2e7c9508d21))

## [0.15.1](https://github.com/home-operations/kromgo/compare/0.15.0...0.15.1) (2026-07-14)


### Features

* add configurable favicon (GET /favicon.ico) ([#302](https://github.com/home-operations/kromgo/issues/302)) ([bd2216a](https://github.com/home-operations/kromgo/commit/bd2216a9621b16cedaa65b4a8e42301e76f4b096))
* **deps:** update dependency simple-icons (16.24.1 → 16.25.0) ([#289](https://github.com/home-operations/kromgo/issues/289)) ([c5d6b67](https://github.com/home-operations/kromgo/commit/c5d6b672c1421c5643bc9b4bd49f07eedf625117))
* **deps:** update dependency simple-icons (16.25.0 → 16.26.0) ([#301](https://github.com/home-operations/kromgo/issues/301)) ([c25bb87](https://github.com/home-operations/kromgo/commit/c25bb87de9ca9f47b81394077cd93dff80979e9f))
* **deps:** update module github.com/prometheus/common (v0.69.0 → v0.70.0) ([#299](https://github.com/home-operations/kromgo/issues/299)) ([ac96540](https://github.com/home-operations/kromgo/commit/ac9654035c02c59b37882d317ddc387ef588086c))
* **deps:** update module golang.org/x/image (v0.43.0 → v0.44.0) ([#295](https://github.com/home-operations/kromgo/issues/295)) ([95356ed](https://github.com/home-operations/kromgo/commit/95356ed8e438904ad353182001a1bf91ec390fe4))


### Bug Fixes

* **deps:** update dependency marked (18.0.5 → 18.0.6) ([#297](https://github.com/home-operations/kromgo/issues/297)) ([64527f0](https://github.com/home-operations/kromgo/commit/64527f04cb4c2b663592d4ae8f93746a4362a084))
* **deps:** update module github.com/google/cel-go (v0.29.1 → v0.29.2) ([#292](https://github.com/home-operations/kromgo/issues/292)) ([3583ca9](https://github.com/home-operations/kromgo/commit/3583ca9ae91c73ca3bd57866523b253e180ccf30))


### Miscellaneous Chores

* **mise:** Update tool go (1.26.4 → 1.26.5) ([#294](https://github.com/home-operations/kromgo/issues/294)) ([70ac4cb](https://github.com/home-operations/kromgo/commit/70ac4cbafa674f3754685250c1caf494258873da))
* **mise:** Update tool helm (4.2.2 → 4.2.3) ([#298](https://github.com/home-operations/kromgo/issues/298)) ([14fcf01](https://github.com/home-operations/kromgo/commit/14fcf0181b6104dd27b9f859663df87f2a1c39f1))
* **mise:** Update tool lefthook (2.1.9 → 2.1.10) ([#293](https://github.com/home-operations/kromgo/issues/293)) ([3adb5fb](https://github.com/home-operations/kromgo/commit/3adb5fb1c02e0c3f397bb7ab75b820b4520846dd))
* **mise:** Update tool npm (11.18.0 → 12.0.0) ([#296](https://github.com/home-operations/kromgo/issues/296)) ([56bed21](https://github.com/home-operations/kromgo/commit/56bed212fc7ac8651263dbe01387ae276a21b45a))
* **mise:** Update tool npm (12.0.0 → 12.0.1) ([#300](https://github.com/home-operations/kromgo/issues/300)) ([f8d27a5](https://github.com/home-operations/kromgo/commit/f8d27a5304d1d52f0d46e7b859c47ef19c2fc850))
* **mise:** Update tool oxfmt (0.57.0 → 0.58.0) ([#291](https://github.com/home-operations/kromgo/issues/291)) ([fa2c0e0](https://github.com/home-operations/kromgo/commit/fa2c0e0d09749285a7aa0dde72f22bbde50b48ec))
* **mise:** Update tool oxfmt (0.58.0 → 0.59.0) ([#303](https://github.com/home-operations/kromgo/issues/303)) ([0dbfcb8](https://github.com/home-operations/kromgo/commit/0dbfcb87030f12152c2d9498c2f0077ecfcdb497))
* **mise:** Update tool zizmor (1.26.1 → 1.27.0) ([#304](https://github.com/home-operations/kromgo/issues/304)) ([9a85418](https://github.com/home-operations/kromgo/commit/9a85418794a6d63a820e90392d05d17d858aa56b))

## [0.15.0](https://github.com/home-operations/kromgo/compare/0.14.12...0.15.0) (2026-07-04)


### ⚠ BREAKING CHANGES

* drop the legacy /-/health and /-/ready aliases ([#288](https://github.com/home-operations/kromgo/issues/288))
* prefix all environment variables with KROMGO_ ([#286](https://github.com/home-operations/kromgo/issues/286))
* serve health on the main port; metrics port becomes fully optional ([#285](https://github.com/home-operations/kromgo/issues/285))

### Features

* drop the legacy /-/health and /-/ready aliases ([#288](https://github.com/home-operations/kromgo/issues/288)) ([c6fe382](https://github.com/home-operations/kromgo/commit/c6fe382c911df12eddcdd7b34828a907e20494f1))
* prefix all environment variables with KROMGO_ ([#286](https://github.com/home-operations/kromgo/issues/286)) ([340c62d](https://github.com/home-operations/kromgo/commit/340c62dae2261599f98e312d877df18fbfa31330))
* serve health on the main port; metrics port becomes fully optional ([#285](https://github.com/home-operations/kromgo/issues/285)) ([948cab4](https://github.com/home-operations/kromgo/commit/948cab4e1cab29b9acafb7e4697dfe8911a8a54b))

## [0.14.12](https://github.com/home-operations/kromgo/compare/0.14.11...0.14.12) (2026-07-03)


### Features

* **deps:** update module github.com/google/cel-go (v0.28.1 → v0.29.1) ([#284](https://github.com/home-operations/kromgo/issues/284)) ([683a3c6](https://github.com/home-operations/kromgo/commit/683a3c616be1b15b543b073d51ba2fea68c97c74))


### Bug Fixes

* **deps:** update dependency simple-icons (16.24.0 → 16.24.1) ([#280](https://github.com/home-operations/kromgo/issues/280)) ([cd07474](https://github.com/home-operations/kromgo/commit/cd07474d4e5f030576448c1e45024cefa4773c47))


### Miscellaneous Chores

* **mise:** Update tool node (24.17.0 → v24.18.0) ([#279](https://github.com/home-operations/kromgo/issues/279)) ([7722f4a](https://github.com/home-operations/kromgo/commit/7722f4a65706b325f35544e347459559a0cc0fe0))
* **mise:** Update tool npm (11.17.0 → 11.18.0) ([#282](https://github.com/home-operations/kromgo/issues/282)) ([1c7d7e9](https://github.com/home-operations/kromgo/commit/1c7d7e997c6ccf33213875d0e4ad290bc16c6f32))
* **mise:** Update tool oxfmt (0.56.0 → 0.57.0) ([#283](https://github.com/home-operations/kromgo/issues/283)) ([f819988](https://github.com/home-operations/kromgo/commit/f81998879a92e952b21060c7bf5fd9cdddf1bdbc))
* **renovate:** inherit shared toolchain + chart-docs presets ([#277](https://github.com/home-operations/kromgo/issues/277)) ([56e0934](https://github.com/home-operations/kromgo/commit/56e0934730a489b5e87c792f7e75fc71bc3ffaf0))

## [0.14.11](https://github.com/home-operations/kromgo/compare/0.14.10...0.14.11) (2026-06-26)


### Features

* add humanize value helper with SI metric prefixes ([#276](https://github.com/home-operations/kromgo/issues/276)) ([a55cbd1](https://github.com/home-operations/kromgo/commit/a55cbd1be1d8d39b6cf1e6c3a085f75d406d34f5))
* **container:** update image mirror.gcr.io/curlimages/curl (8.20.0 → 8.21.0) ([#275](https://github.com/home-operations/kromgo/issues/275)) ([cb7863e](https://github.com/home-operations/kromgo/commit/cb7863ed66ebfafd93265eebb0a5c4b91f12bd5f))
* **deps:** update dependency simple-icons (16.23.0 → 16.24.0) ([#270](https://github.com/home-operations/kromgo/issues/270)) ([9b49e38](https://github.com/home-operations/kromgo/commit/9b49e381523df6d7767dd6edf0e0e610367491fb))
* **deps:** update module github.com/go-analyze/charts (v0.5.27 → v0.6.0) ([#274](https://github.com/home-operations/kromgo/issues/274)) ([bd53d31](https://github.com/home-operations/kromgo/commit/bd53d315220c4321528fc1e6dc4a7fc3a5c7d9b6))
* **deps:** update module golang.org/x/image (v0.42.0 → v0.43.0) ([#266](https://github.com/home-operations/kromgo/issues/266)) ([96e4774](https://github.com/home-operations/kromgo/commit/96e4774ad67729d473d12d121a1e41e3316b4ac2))


### Bug Fixes

* **deps:** update module go.yaml.in/yaml/v4 (v4.0.0-rc.5 → v4.0.0-rc.6) ([#268](https://github.com/home-operations/kromgo/issues/268)) ([64a8a99](https://github.com/home-operations/kromgo/commit/64a8a9952a089acfd4f5b1dc5cb4647e3bb88808))


### Miscellaneous Chores

* add minimumGroupSize to Go and Node toolchain groups ([c90f50d](https://github.com/home-operations/kromgo/commit/c90f50dffac197786e8a5c3083ff19a64c78c5b7))
* **mise:** Update tool oxfmt (0.55.0 → 0.56.0) ([#271](https://github.com/home-operations/kromgo/issues/271)) ([2597efd](https://github.com/home-operations/kromgo/commit/2597efd6376dc163f8d031b0d4ff4218ef3bd81e))
* **mise:** Update tool zizmor (1.25.2 → 1.26.1) ([#269](https://github.com/home-operations/kromgo/issues/269)) ([f3daed1](https://github.com/home-operations/kromgo/commit/f3daed16bef5622cfcad52e464a7394cdee17b44))

## [0.14.10](https://github.com/home-operations/kromgo/compare/0.14.9...0.14.10) (2026-06-18)


### Features

* move the monitoring port (health + metrics) to 8081 ([#265](https://github.com/home-operations/kromgo/issues/265)) ([36a9122](https://github.com/home-operations/kromgo/commit/36a9122c1935da1463ffb85a74e751e4d235f758))


### Miscellaneous Chores

* **mise:** update node toolchain (24.16.0 → v24.17.0) ([#263](https://github.com/home-operations/kromgo/issues/263)) ([0637abc](https://github.com/home-operations/kromgo/commit/0637abc8b7aa61e7de822416ea2054f44ddba47e))
* **mise:** update tool helm (4.2.1 → 4.2.2) ([#261](https://github.com/home-operations/kromgo/issues/261)) ([e9c06d1](https://github.com/home-operations/kromgo/commit/e9c06d147497c9cfccbc1292f33c88c7e522c2ee))

## [0.14.9](https://github.com/home-operations/kromgo/compare/0.14.8...0.14.9) (2026-06-17)


### Features

* **badge:** add for-the-badge style ([#259](https://github.com/home-operations/kromgo/issues/259)) ([d15fe1f](https://github.com/home-operations/kromgo/commit/d15fe1f623c6084f93cd55b56b5e5ac782da78a7))
* **deps:** update module github.com/prometheus/common (v0.68.1 → v0.69.0) ([#258](https://github.com/home-operations/kromgo/issues/258)) ([0472460](https://github.com/home-operations/kromgo/commit/047246024fa4e6c33487c7041dccea2f5ea554ca))


### Miscellaneous Chores

* **mise:** update tool oxfmt (0.54.0 → 0.55.0) ([#257](https://github.com/home-operations/kromgo/issues/257)) ([3a90072](https://github.com/home-operations/kromgo/commit/3a90072a2e032ed08b8fd402b7a10d508d08d6f6))

## [0.14.8](https://github.com/home-operations/kromgo/compare/0.14.7...0.14.8) (2026-06-15)


### Miscellaneous Chores

* **chart:** use the shared curl image for the helm test pod ([#253](https://github.com/home-operations/kromgo/issues/253)) ([e4601be](https://github.com/home-operations/kromgo/commit/e4601befef655f454982ff158b90089756072152))

## [0.14.7](https://github.com/home-operations/kromgo/compare/0.14.6...0.14.7) (2026-06-13)


### Miscellaneous Chores

* **helm:** add deploymentAnnotations to values.yaml ([#247](https://github.com/home-operations/kromgo/issues/247)) ([5e482ab](https://github.com/home-operations/kromgo/commit/5e482ab47a91cc797a0051b76742603b460fe99d))
* **mise:** update tool helm (4.2.0 → 4.2.1) ([#248](https://github.com/home-operations/kromgo/issues/248)) ([0811240](https://github.com/home-operations/kromgo/commit/0811240e808f423b6e16a67cc8c4dc6f82502256))

## [0.14.6](https://github.com/home-operations/kromgo/compare/0.14.5...0.14.6) (2026-06-11)


### Features

* **mise:** update tool cosign (3.0.6 → 3.1.1) ([#238](https://github.com/home-operations/kromgo/issues/238)) ([c7ef608](https://github.com/home-operations/kromgo/commit/c7ef608f080be4154fc513f9e9fcd1d2bc8339a2))


### Bug Fixes

* bound HTTP timeouts, cap graph series, guard non-finite values ([#244](https://github.com/home-operations/kromgo/issues/244)) ([837dc16](https://github.com/home-operations/kromgo/commit/837dc16d2d872edba515b964fd96c27cd11547c0))


### Miscellaneous Chores

* **mise:** update tool npm (11.16.0 → 11.17.0) ([#246](https://github.com/home-operations/kromgo/issues/246)) ([13522b0](https://github.com/home-operations/kromgo/commit/13522b0d6227f52a27a382195e38006a0d83084f))

## [0.14.5](https://github.com/home-operations/kromgo/compare/0.14.4...0.14.5) (2026-06-09)


### Features

* **deps:** update module golang.org/x/image (v0.41.0 → v0.42.0) ([#233](https://github.com/home-operations/kromgo/issues/233)) ([c9e267f](https://github.com/home-operations/kromgo/commit/c9e267fd5ccf6d8b5283780d316aa214311cff43))
* **github-release:** update release helm-unittest/helm-unittest (v1.0.3 → v1.1.1) ([#232](https://github.com/home-operations/kromgo/issues/232)) ([a7d09f0](https://github.com/home-operations/kromgo/commit/a7d09f0460b11844bc266431ff9e41994594e22a))
* **mise:** update tool oxfmt (0.53.0 → 0.54.0) ([#234](https://github.com/home-operations/kromgo/issues/234)) ([f4e93a8](https://github.com/home-operations/kromgo/commit/f4e93a8912d6f0573b103762692c5c81475702ed))


### Bug Fixes

* **deps:** update module go.yaml.in/yaml/v4 (v4.0.0-rc.4 → v4.0.0-rc.5) ([#235](https://github.com/home-operations/kromgo/issues/235)) ([eb4b9ef](https://github.com/home-operations/kromgo/commit/eb4b9efece25a82dab941fbb1483d52389564475))

## [0.14.4](https://github.com/home-operations/kromgo/compare/0.14.3...0.14.4) (2026-06-07)


### Features

* **chart:** digest pinning, generated README + values schema, and helm tests ([#227](https://github.com/home-operations/kromgo/issues/227)) ([d0aa57a](https://github.com/home-operations/kromgo/commit/d0aa57ae5f01314bd695bfb3cb135fda721e0e4c))
* **container:** update image mirror.gcr.io/busybox (1.37.0 → 1.38.0) ([#228](https://github.com/home-operations/kromgo/issues/228)) ([03cf4d7](https://github.com/home-operations/kromgo/commit/03cf4d7a83ad914bf1c627f02f3fa9c84ebbf647))
* **deps:** update dependency simple-icons (16.22.0 → 16.23.0) ([#226](https://github.com/home-operations/kromgo/issues/226)) ([11cd211](https://github.com/home-operations/kromgo/commit/11cd21134674d152c35bd9f6c0241e927473be21))


### Bug Fixes

* **chart:** pin the helm-test image as tag@digest so renovate updates both ([#230](https://github.com/home-operations/kromgo/issues/230)) ([ba91657](https://github.com/home-operations/kromgo/commit/ba916575d839d149bd4b5271b259e441e787e6c4))


### Miscellaneous Chores

* remove automerge setting from toolchain groups ([be05683](https://github.com/home-operations/kromgo/commit/be05683d104870c3433669f4b9ef219a7e64cfe2))
* Update release-please-config.json to remove paths ([cce6ef8](https://github.com/home-operations/kromgo/commit/cce6ef8f3c4a0cc1096b38f7ec53defdc50de9ef))

## [0.14.3](https://github.com/home-operations/kromgo/compare/0.14.2...0.14.3) (2026-06-05)


### Features

* **chart:** disable ServiceAccount token automount by default ([#220](https://github.com/home-operations/kromgo/issues/220)) ([2fcdd05](https://github.com/home-operations/kromgo/commit/2fcdd054ddff173b77136c4614266a88afb9df6e))
* **chart:** generate values.schema.json ([#224](https://github.com/home-operations/kromgo/issues/224)) ([f2b1ac9](https://github.com/home-operations/kromgo/commit/f2b1ac9da00ab68bee4fefc34dcdace4ab2e9d35))
* **chart:** render values through tpl ([#218](https://github.com/home-operations/kromgo/issues/218)) ([c04ee32](https://github.com/home-operations/kromgo/commit/c04ee3279f32ea1d4ecf15d144e07172274ddd1d))
* **graph:** optional area fill (fill: true) ([#222](https://github.com/home-operations/kromgo/issues/222)) ([c7045f2](https://github.com/home-operations/kromgo/commit/c7045f2b8e0a01c5f7b945953c9cc61b2fa503af))
* **graph:** y-axis min/max + reference mark lines ([#223](https://github.com/home-operations/kromgo/issues/223)) ([b304342](https://github.com/home-operations/kromgo/commit/b304342603941583f4b64635cdccc7d1372074ca))


### Bug Fixes

* **graph:** round y-axis tick values (PreferNiceIntervals) ([#221](https://github.com/home-operations/kromgo/issues/221)) ([67fcecc](https://github.com/home-operations/kromgo/commit/67fceccf5d59a6ae42108d2d31007fc8ffd638fc))

## [0.14.2](https://github.com/home-operations/kromgo/compare/0.14.1...0.14.2) (2026-06-05)


### Features

* graph valueExpr for y-axis label formatting ([#217](https://github.com/home-operations/kromgo/issues/217)) ([e31db54](https://github.com/home-operations/kromgo/commit/e31db54cb93352fe7fd9941f2dac5654e0d3ff4d))


### Bug Fixes

* no title badge rendering ([#215](https://github.com/home-operations/kromgo/issues/215)) ([e8997f3](https://github.com/home-operations/kromgo/commit/e8997f33f417f9449b473ee1d6a10a5d845e27bf))

## [0.14.1](https://github.com/home-operations/kromgo/compare/0.14.0...0.14.1) (2026-06-05)


### Features

* add kromgo helm chart ([#214](https://github.com/home-operations/kromgo/issues/214)) ([168e479](https://github.com/home-operations/kromgo/commit/168e4791702371cfb37dcdbc0bff8157b34c3641))


### Bug Fixes

* **deps:** update dependency marked (18.0.4 → 18.0.5) ([#213](https://github.com/home-operations/kromgo/issues/213)) ([f1e127d](https://github.com/home-operations/kromgo/commit/f1e127d394925f9f407418a6ae237e98d2d298de))
* **deps:** update module github.com/prometheus/common (v0.68.0 → v0.68.1) ([#210](https://github.com/home-operations/kromgo/issues/210)) ([5f7a860](https://github.com/home-operations/kromgo/commit/5f7a8606283cc4d1074fa1d3e85dff36cbac9538))
* **mise:** update tool go (1.26.3 → 1.26.4) ([cde60a0](https://github.com/home-operations/kromgo/commit/cde60a017bf76153dfef04525520fd851924a38f))

## [0.14.0](https://github.com/home-operations/kromgo/compare/0.13.1...0.14.0) (2026-06-02)


### ⚠ BREAKING CHANGES

* **badge:** unique ids, error badges, labelColor, hand-rolled formatters ([#207](https://github.com/home-operations/kromgo/issues/207))

### Features

* **badge:** unique ids, error badges, labelColor, hand-rolled formatters ([#207](https://github.com/home-operations/kromgo/issues/207)) ([6b69e8c](https://github.com/home-operations/kromgo/commit/6b69e8ccc1bd0f5aa24c51e6800d179d0cc51b4c))

## [0.13.1](https://github.com/home-operations/kromgo/compare/0.13.0...0.13.1) (2026-06-02)


### Features

* **badge:** adapt text color to background + add aria-label/title ([#205](https://github.com/home-operations/kromgo/issues/205)) ([e47ba5a](https://github.com/home-operations/kromgo/commit/e47ba5aaa8dcbacbaff7c0eb8c9259f2be55dc44))

## [0.13.0](https://github.com/home-operations/kromgo/compare/0.12.2...0.13.0) (2026-06-02)


### ⚠ BREAKING CHANGES

* **cache:** add a global cache config block for Cache-Control headers ([#203](https://github.com/home-operations/kromgo/issues/203))

### Features

* **cache:** add a global cache config block for Cache-Control headers ([#203](https://github.com/home-operations/kromgo/issues/203)) ([3175d6f](https://github.com/home-operations/kromgo/commit/3175d6f80f6df643de396b3422f2c771432e3473))
* embed DejaVu Sans + Comic Neue via npm (DejaVu default, shields.io look) ([#202](https://github.com/home-operations/kromgo/issues/202)) ([e52358e](https://github.com/home-operations/kromgo/commit/e52358e5a301224846b9b7f81b801a7becb82c91))

## [0.12.2](https://github.com/home-operations/kromgo/compare/0.12.1...0.12.2) (2026-06-01)


### Features

* stamp build version into the container image ([#201](https://github.com/home-operations/kromgo/issues/201)) ([e1d5bd5](https://github.com/home-operations/kromgo/commit/e1d5bd5a90611f4115a78f47c7033e92f5b06dea))
* support Simple Icons for badge icons alongside MDI ([#199](https://github.com/home-operations/kromgo/issues/199)) ([dc15e13](https://github.com/home-operations/kromgo/commit/dc15e138911b2e188dee14cc21906267313fa442))

## [0.12.1](https://github.com/home-operations/kromgo/compare/0.12.0...0.12.1) (2026-06-01)


### Features

* **mise:** update tool oxfmt (0.52.0 → 0.53.0) ([6cf9b97](https://github.com/home-operations/kromgo/commit/6cf9b978bd742ca97bad35278c3cf7aae89368bb))


### Miscellaneous Chores

* update mise lockfile ([efe5420](https://github.com/home-operations/kromgo/commit/efe542010b36222105fd0edc66d04650d12f6ec8))


### Code Refactoring

* request-scoped logging and comprehensive test cleanup ([#197](https://github.com/home-operations/kromgo/issues/197)) ([eb12795](https://github.com/home-operations/kromgo/commit/eb12795cc1d2b998aa4f13d83188c6eb33c87b18))

## [0.12.0](https://github.com/home-operations/kromgo/compare/0.11.1...0.12.0) (2026-06-01)


### ⚠ BREAKING CHANGES

* typed badge/graph endpoints with themed SVG/PNG graphs (0.12) ([#194](https://github.com/home-operations/kromgo/issues/194))

### Features

* typed badge/graph endpoints with themed SVG/PNG graphs (0.12) ([#194](https://github.com/home-operations/kromgo/issues/194)) ([3c15810](https://github.com/home-operations/kromgo/commit/3c15810905432056bad7abb530246a931e0dfae1))

## [0.11.1](https://github.com/home-operations/kromgo/compare/0.11.0...0.11.1) (2026-06-01)


### Code Refactoring

* reorganize kromgo package and clean up tests ([#192](https://github.com/home-operations/kromgo/issues/192)) ([1e16d0a](https://github.com/home-operations/kromgo/commit/1e16d0a2fdae4ef271e7070fe1cbbdba55bab56d))

## [0.11.0](https://github.com/home-operations/kromgo/compare/v0.10.0...0.11.0) (2026-06-01)


### ⚠ BREAKING CHANGES

* modernize kromgo — CEL config, range queries, caching, lighter deps ([#189](https://github.com/home-operations/kromgo/issues/189))

### Features

* **deps:** update module golang.org/x/image (v0.38.0 → v0.41.0) ([#190](https://github.com/home-operations/kromgo/issues/190)) ([9dca5d4](https://github.com/home-operations/kromgo/commit/9dca5d467256cb87abb45b5075f32ea85daa7adf))
* migrate kromgo to home-operations ([#187](https://github.com/home-operations/kromgo/issues/187)) ([781d872](https://github.com/home-operations/kromgo/commit/781d8724e89681747e355133b7efa3a4dfc8bf70))
* modernize kromgo — CEL config, range queries, caching, lighter deps ([#189](https://github.com/home-operations/kromgo/issues/189)) ([4a227b1](https://github.com/home-operations/kromgo/commit/4a227b1dd419e2db22ec5dc001fdf1b3408e806e))


### Bug Fixes

* **deps:** update module github.com/go-chi/chi/v5 to v5.3.0 ([#184](https://github.com/home-operations/kromgo/issues/184)) ([282a24e](https://github.com/home-operations/kromgo/commit/282a24e5b9235d5d93d59e6a23b3ff58d305b99c))
* **deps:** update module github.com/prometheus/common to v0.68.0 ([#185](https://github.com/home-operations/kromgo/issues/185)) ([3f95ab2](https://github.com/home-operations/kromgo/commit/3f95ab224c8a8eb0710ec24b87691dd3dba3c1ee))

## Changelog
