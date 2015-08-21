# GoIsilon

## Overview
GoIsilon represents API bindings for Go that allow you to manage Isilon storage platforms.  In the true nature of API bindings, it is intended that the functions available are basically a direct implementation of what is available through the API.

## Functions
* Namespace RAN 
   * CreateDirectory - creates a new directory under the /ifs/
   * DirectoryExists - checks the existance of a directory on the filesystem
   * DeleteDirectory - deletes the given directory path
   * UpdateDirAcl - updates directory's ACL configuration with given params
* Snapshots
   * GetSnapshots - get a list of snapshots
   * CreateSnapshot - start a snapshot on a filesystem path
   * DeleteSnapshot - delete the snapshot

## Examples
Note: You can use each api handler's corresponding test files for more reference.

Initialize Isilon package
```Go
isi, err := goisilon.New()
if err != nil {
   return nil, err.Error()
}
```

Create a directory
```Go
headers := make(map[string]string)
headers["x-isi-ifs-access-control"] = "public_read_write"
err := isi.CreateDirectory("/ifs/data/test", headers, false)
```

## Environment Variables
    GOISILON_ENDPOINT - the API endpoint, ie. https://10.5.132.140:8080/
    GOISILON_USERNAME - the username
    GOISILON_PASSWORD - the password
    GOISILON_INSECURE - whether to skip SSL validation

## Contributions
If you are familiar with Isilon cluster APIs please help. It will be greatly appreciated.

Licensing
---------
Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

Support
-------
Please file bugs and issues at the Github issues page. For more general discussions you can contact the EMC Code team at <a href="https://groups.google.com/forum/#!forum/emccode-users">Google Groups</a> or tagged with **EMC** on <a href="https://stackoverflow.com">Stackoverflow.com</a>. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
