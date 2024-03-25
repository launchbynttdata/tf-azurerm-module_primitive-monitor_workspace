// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

variable "name" {
  description = "The name of the resource"
  type        = string
}

variable "resource_group_name" {
  description = "The name of the resource group in which the resource exists"
  type        = string
}

variable "location" {
  description = "The location/region of the resource"
  type        = string
  default     = "eastus"
}

variable "public_network_access_enabled" {
  description = "Whether or not public network access is allowed for the storage account"
  type        = bool
  default     = true
}

variable "tags" {
  description = "Tags for the resources"
  type        = map(string)
  default     = {}
}
