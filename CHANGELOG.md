## Version 1.1.7 - latest

#### New Features
- **ScaleFork Function & Task Engine**:  
  Introduced the `ScaleFork` function alongside a new Engine for task management and execution. The Engine now supports dynamic adjustment of worker threads based on available CPU cores and actively monitors running tasks.
  
- **File Compression Utilities**:  
  Added functions to create ZIP and TAR.GZ archives from folders, with support for multiple files. These functions include a file size check for the TAR format to ensure robust compression capabilities.
  
- **Document Translation**:  
  translation of documents into English, with the aim of improving internationalization and facilitating use by all users.
  
- **Import Path Updates**:  
  Updated import paths to reflect the new directory structure for system packages, ensuring a cleaner and more maintainable codebase.

---

## Version 1.1.6

#### New Features
- **Network Monitoring**: Added functionality to calculate download and upload rates.
- **GetFileInfo Field**: Added `FileExtension`, `FileData`, and `FileDataBuffer` fields to the `FileInfo` structure and updated `GetFileInfo` to include file extensions.

#### Refactoring
- **calculateCPUUsage**: Refactored the `calculateCPUUsage` function to average multiple samples, improving accuracy.

#### Improvements
- **Documentation**: Updated `README.md` to include the detailed structure of `FileInfo` and field descriptions.

---

## Version 1.1.5

#### New Features
- **GetFolderSize Function**: Added the 

GetFolderSize

 function to calculate the size of a folder in kilobytes.
- **GetFileInfo Function**: Added the 

GetFileInfo

 function to obtain detailed information about specific files in a directory. The function now accepts variadic file names.

#### Changed
- **Updated Documentation**: Added documentation for the function 

---

## Version 1.1.4.1

### `configs.go` Refactoring
- Improved configuration handling to properly manage empty values in the configuration file.

### `network.go` Refactoring
- Enhanced network handling to support IPv4 and IPv6 addresses seamlessly, improving the capture of ports and IPs for both IPv4 and IPv6, providing greater support for other distributions. The use of native Linux and Go tools helps with compatibility across various distributions.

### `cpu.go` Refactoring
- Updated the `CPUInfo` struct and refined CPU usage calculation for more accurate performance metrics.

### Version Update
- Updated the project version to `1.1.4.1`.

### Checklist
- [x] Refactored code for `configs.go`
- [x] Refactored code for `network.go`
- [x] Refactored code for `cpu.go`
- [x] Version updated to `1.1.4.1`
- [x] Tests updated and passed

### Release Notes

#### Added
- Continuous support for IPv4 and IPv6 addresses in `network.go`.

#### Changed
- Improved configuration handling in `configs.go` to manage empty values.
- Refined CPU usage calculation in `cpu.go` for more accurate performance metrics.

#### Fixed
- Improved overall project stability and compatibility.

#### Removed
- No functionality removed in this version.

---

### Version 1.1.4

#### New Features
- **Linux Distribution Capture**: Implemented detailed capture of information about the Linux distribution.
- **Updated Documentation**: Added references to the framework and project logo in the documentation.
- **Code Comments**: Inserted explanatory comments in all functions to facilitate code understanding.
- **Network Port Capture**: Added support for capturing system TCP and UDP ports.

#### Fixes
- **Logic Correction**: Fixed processing logic errors, resulting in greater system stability.
