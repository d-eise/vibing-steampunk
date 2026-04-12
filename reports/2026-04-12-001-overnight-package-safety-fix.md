# Overnight Work Report: SAP_ALLOWED_PACKAGES Enforcement Fix

**Date:** 2026-04-12
**Authors:** Claude (implementation), Codex (audit + review), Alice (direction)
**PR:** https://github.com/oisee/vibing-steampunk/pull/101
**Commit:** 08e3d78
**Branch:** pr-93-fix

---

## Summary

User Philip Dolker reported that `SAP_ALLOWED_PACKAGES` did not restrict object
modifications to the configured package whitelist — vsp was letting him edit
objects in any package regardless of the setting.

Investigation confirmed this was a real enforcement bug, not a configuration
issue. The fix is committed and pushed in PR #101, with Codex's review approval.

## Problem

`SAP_ALLOWED_PACKAGES` was correctly parsed from the environment and wired into
`SafetyConfig.AllowedPackages`, but the check was only enforced on a small
subset of operations:

- `CreateObject` (crud.go)
- `CreateAndActivateProgram`, `CreateClassWithTests` (workflows.go)
- `UI5CreateApp` (ui5.go)

The main mutation paths that AI agents actually use to modify existing objects
bypassed the package check entirely. An agent could call `EditSource`,
`WriteSource`, `WriteProgram`, or `WriteClass` on any object in any package and
the `AllowedPackages` restriction would have no effect.

## Root Cause

The safety check was applied at create-time (where the package is an explicit
parameter) but never at edit-time (where the package must be resolved from the
existing object's metadata).

## Fix

Added `checkObjectPackageSafety(ctx, objectURL)` calls to all existing-object
mutation paths. This helper already existed in `client.go` — it resolves the
object's package via `SearchObject` and validates it against the whitelist,
with URL normalization for `/source/main` and `/includes/...` suffixes.

### Functions hardened

| Function | File | Check added |
|---|---|---|
| `EditSourceWithOptions` | `workflows_edit.go` | `checkObjectPackageSafety` |
| `WriteProgram` | `workflows.go` | `checkObjectPackageSafety` |
| `WriteClass` | `workflows.go` | `checkObjectPackageSafety` |
| `RenameObject` | `workflows_fileio.go` | `checkObjectPackageSafety` (old) + `checkPackageSafety` (target) |
| `UpdateSource` | `crud.go` | `checkObjectPackageSafety` + `checkTransportableEdit` |
| `DeleteObject` | `crud.go` | `checkObjectPackageSafety` + `checkTransportableEdit` |
| `CreateTestInclude` | `crud.go` | `checkObjectPackageSafety` + `checkTransportableEdit` |
| `UpdateClassInclude` | `crud.go` | `checkObjectPackageSafety` + `checkTransportableEdit` |
| `WriteMessageClassTexts` | `i18n.go` | `checkObjectPackageSafety` + `checkTransportableEdit` |
| `WriteDataElementLabels` | `i18n.go` | `checkObjectPackageSafety` + `checkTransportableEdit` |

Note: `WriteSource` didn't need a direct check — its update path delegates to
`WriteProgram`/`WriteClass` which are now guarded, and its create path
delegates to `CreateAndActivateProgram`/`CreateClassWithTests` which already
had checks.

### Tests

New tests in `client_test.go`:
- `TestClient_CheckObjectPackageSafety_NormalizesObjectURLs` — verifies
  `/source/main` and `/includes/...` URL normalization
- `TestClient_UpdateSource_EnforcesAllowedPackages`
- `TestClient_DeleteObject_EnforcesAllowedPackages`
- `TestClient_CreateTestInclude_EnforcesAllowedPackages`
- `TestClient_WriteMessageClassTexts_EnforcesAllowedPackages`

Full test suite passes: `go test ./...` clean, zero regressions.

## Defense in Depth

The fix applies checks at both the workflow layer (early rejection with clear
errors) and the low-level CRUD layer (backstop). This is intentional
redundancy — it makes intent obvious at the call sites and prevents future
workflow additions from silently bypassing the guard.

## Remaining Gap — Tracked Follow-up

UI5/BSP mutation paths (`UI5UploadFile`, `UI5DeleteFile`, `UI5DeleteApp`) are
**not yet guarded**. These use a different URL scheme
(`/sap/bc/adt/filestore/ui5-bsp/objects/...`) and need a separate
app-to-package resolution helper before they can share the same check.

Recommended approach (from Codex's audit):
1. Add `UI5ResolveAppPackage(appName)` using BSP metadata
2. Apply package enforcement on all UI5 mutation paths
3. Ship as a follow-up PR

This gap is not a regression — those paths were never guarded. It's an
existing edge case that warrants its own focused fix.

## Timeline

- **23:00** — User report received from Philip Dolker via Alice
- **23:05** — Code audit by Claude and Codex independently; both identified
  the same gap in the workflow layer
- **23:15** — Alice approved fix implementation; Claude took implementation,
  Codex switched to review/foreman mode
- **23:35** — Fix implemented across 7 files, tests pass
- **23:45** — Codex review approved, no blockers
- **23:50** — Committed (08e3d78), pushed, PR #101 created
- **23:55** — Graceful handover, session wrap-up

## Philip Reply Draft (not yet sent)

> Hi Philip,
>
> Thank you for reporting this — and congrats-thanks for the 200+ stars!
>
> You found a real bug. Your configuration was correct, but
> `SAP_ALLOWED_PACKAGES` was only being enforced on object creation paths
> (like `CreateObject`). The main mutation paths for existing objects —
> `EditSource`, `WriteSource`, `WriteProgram`, `WriteClass`, and others —
> were not checking the package restriction at all. So vsp would let an
> agent modify any object regardless of the configured package whitelist.
>
> This is now fixed. The package ownership check is enforced on all
> standard ADT object mutation paths (edit, write, update, delete, rename).
>
> One additional note: if `$CONDUCT_RAPTEST` is a transportable package
> (not `$TMP`), you may also need to set `SAP_ALLOW_TRANSPORTABLE_EDITS=true`
> in your env block to allow writes to objects in that package. Without it,
> vsp blocks edits to objects in non-local packages as an extra safety layer.
>
> The fix will be in the next release. A small remaining area (UI5/BSP app
> mutations) will be completed separately, but that shouldn't affect typical
> ABAP development scenarios.
>
> Thanks again for the detailed report — it helped us find and close a real
> enforcement gap.

## Pending Actions

- [ ] Merge PR #101 to main
- [ ] Rebuild release binaries (if hotfix desired)
- [ ] Send reply to Philip
- [ ] Open follow-up issue for UI5 BSP app-to-package resolution
