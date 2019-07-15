# A comparison between different types of C# projects

This folder contains 2 types of C# project and aims to show the possible differences in scan results.

The first - `match-dlls` - is an example that bundles all the DLLs it requires including third-party dependencies and their dependencies. Run `nuget restore` which will download the required DLLs into `packages` folder prior to scanning.

The second - `match-nuget` - is an example that uses `PackageReference` format in its `.csproj` file to declare a dependency on a third-party library on Nuget.

Both projects depend on `WebAssetBundler` 0.9.0. `WebAssetBundler` also includes `dotless.ClientOnly.dll` among other dlls in its Nuget package.

When scanning `match-dlls` with a scanner that scans for DLLs, it is possible for the scan results to show more than `WebAssetBundler` as a dependency. The DLL scanner may also pick up `dotless.ClientOnly.dll` as a third-party library.

On the other hand, `match-nuget` does not bundle any DLLs. It declares a single library, `WebAssetBundler`, as a dependency. When using a scanner that conforms to MSBuild's process that generates a `project.assets.json`, the only library that will be present in the resulting `obj/project.assets.json` is `WebAssetBundler`. The scan results may show only `WebAssetBundler` as a dependency.
