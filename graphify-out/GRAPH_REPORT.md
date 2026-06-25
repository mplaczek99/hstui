# Graph Report - .  (2026-06-25)

## Corpus Check
- Corpus is ~7,042 words - fits in a single context window. You may not need a graph.

## Summary
- 88 nodes · 185 edges · 10 communities (8 shown, 2 thin omitted)
- Extraction: 85% EXTRACTED · 15% INFERRED · 0% AMBIGUOUS · INFERRED: 27 edges (avg confidence: 0.8)
- Token cost: 22,000 input · 2,859 output

## Community Hubs (Navigation)
- [[_COMMUNITY_TUI Model & Update Loop|TUI Model & Update Loop]]
- [[_COMMUNITY_Profile Config IO|Profile Config I/O]]
- [[_COMMUNITY_Daemon Control & Project Docs|Daemon Control & Project Docs]]
- [[_COMMUNITY_View & Parsing Tests|View & Parsing Tests]]
- [[_COMMUNITY_View Rendering|View Rendering]]
- [[_COMMUNITY_Dependency Checks & Entrypoint|Dependency Checks & Entrypoint]]
- [[_COMMUNITY_Clamp Helper|Clamp Helper]]
- [[_COMMUNITY_Simple Panel Tests|Simple Panel Tests]]
- [[_COMMUNITY_Time Adjustment|Time Adjustment]]
- [[_COMMUNITY_Package Root|Package Root]]

## God Nodes (most connected - your core abstractions)
1. `T` - 21 edges
2. `model` - 12 edges
3. `hyprsunsetProfile` - 9 edges
4. `saveHyprsunsetProfiles()` - 8 edges
5. `parseProfiles()` - 8 edges
6. `loadHyprsunsetProfiles()` - 7 edges
7. `writeExecutable()` - 7 edges
8. `initialModel()` - 7 edges
9. `hstui (hyprsunset TUI)` - 7 edges
10. `defaultHyprsunsetProfile()` - 6 edges

## Surprising Connections (you probably didn't know these)
- `initialModel()` --calls--> `defaultHyprsunsetProfile()`  [INFERRED]
  tui.go → config.go
- `initialModel()` --calls--> `loadHyprsunsetProfiles()`  [INFERRED]
  tui.go → config.go
- `saveConfigCmd()` --calls--> `saveHyprsunsetProfiles()`  [INFERRED]
  tui.go → config.go
- `TestUnsetFieldOmittedFromConfig()` --calls--> `formatHyprsunsetProfiles()`  [INFERRED]
  main_test.go → config.go
- `TestParseProfiles()` --calls--> `parseProfiles()`  [INFERRED]
  main_test.go → config.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **hstui application layout (tui/view/config/hyprsunset)** — readme_hstui, tui, view, config, hyprsunset [EXTRACTED 1.00]

## Communities (10 total, 2 thin omitted)

### Community 0 - "TUI Model & Update Loop"
Cohesion: 0.22
Nodes (13): Cmd, field, statusMsg, hyprsunsetProfile, TestSKeySavesCurrentConfiguration(), Msg, edit(), fieldBit (+5 more)

### Community 1 - "Profile Config I/O"
Cohesion: 0.29
Nodes (15): configLine(), defaultHyprsunsetProfile(), formatHyprsunsetProfiles(), fieldBit, hyprsunsetConfigFile(), loadHyprsunsetProfiles(), parseMaxGamma(), parseProfiles() (+7 more)

### Community 2 - "Daemon Control & Project Docs"
Cohesion: 0.17
Nodes (14): golangci-lint v2 config, gofmt formatter, misspell linter, revive linter, IsHyprsunsetRunning(), SetHyprsunsetRunning(), startHyprsunset(), wrapOutput() (+6 more)

### Community 3 - "View & Parsing Tests"
Cohesion: 0.30
Nodes (11): T, TestBackspaceClearsAndReaddsField(), TestConfigurationUsesDayNightLabels(), TestParseProfiles(), TestProfileAddDeleteKeys(), TestSimpleDayNightControl(), TestTabSwitchesBetweenPanels(), TestUnsetFieldOmittedFromConfig() (+3 more)

### Community 4 - "View Rendering"
Cohesion: 0.31
Nodes (6): model, simpleCell(), model, profileLabel(), renderBox(), simpleBody()

### Community 5 - "Dependency Checks & Entrypoint"
Cohesion: 0.33
Nodes (5): CheckDependencies(), Notify(), main(), TestCheckDependencies(), TestNotify()

### Community 6 - "Clamp Helper"
Cohesion: 0.67
Nodes (3): TestClamp(), clamp(), T

### Community 7 - "Simple Panel Tests"
Cohesion: 0.67
Nodes (3): TestInitialModelStartsOnSimplePanel(), TestSpaceTogglesSimpleCheckbox(), writeExecutable()

## Knowledge Gaps
- **10 isolated node(s):** `fieldBit`, `hstui`, `T`, `Msg`, `model` (+5 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **2 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `hstui (hyprsunset TUI)` connect `Daemon Control & Project Docs` to `TUI Model & Update Loop`, `Profile Config I/O`, `View Rendering`?**
  _High betweenness centrality (0.244) - this node is a cross-community bridge._
- **Why does `initialModel()` connect `TUI Model & Update Loop` to `Profile Config I/O`, `Daemon Control & Project Docs`, `Dependency Checks & Entrypoint`, `Simple Panel Tests`?**
  _High betweenness centrality (0.166) - this node is a cross-community bridge._
- **Why does `T` connect `View & Parsing Tests` to `TUI Model & Update Loop`, `Profile Config I/O`, `Daemon Control & Project Docs`, `Dependency Checks & Entrypoint`, `Clamp Helper`, `Simple Panel Tests`, `Time Adjustment`?**
  _High betweenness centrality (0.143) - this node is a cross-community bridge._
- **Are the 2 inferred relationships involving `saveHyprsunsetProfiles()` (e.g. with `TestSaveHyprsunsetProfilesWritesConfig()` and `saveConfigCmd()`) actually correct?**
  _`saveHyprsunsetProfiles()` has 2 INFERRED edges - model-reasoned connections that need verification._
- **Are the 3 inferred relationships involving `parseProfiles()` (e.g. with `TestParseProfiles()` and `TestSaveHyprsunsetProfilesWritesConfig()`) actually correct?**
  _`parseProfiles()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `fieldBit`, `hstui`, `T` to the rest of the system?**
  _10 weakly-connected nodes found - possible documentation gaps or missing edges._