# Protocol Documentation
<a name="top"/>

## Table of Contents

- [task.proto](#task.proto)
    - [CompleteTaskRequest](#pb.CompleteTaskRequest)
    - [CreateTaskRequest](#pb.CreateTaskRequest)
    - [GetTaskRequest](#pb.GetTaskRequest)
    - [Task](#pb.Task)
    - [UpdateTaskRequest](#pb.UpdateTaskRequest)
  
    - [Status](#pb.Status)
  
  
    - [TaskService](#pb.TaskService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="task.proto"/>
<p align="right"><a href="#top">Top</a></p>

## task.proto



<a name="pb.CompleteTaskRequest"/>

### CompleteTaskRequest
タスク完了のリクエストメッセージ


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="pb.CreateTaskRequest"/>

### CreateTaskRequest
タスク登録のリクエストメッセージ


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| Description | [string](#string) |  |  |






<a name="pb.GetTaskRequest"/>

### GetTaskRequest
タスク取得のリクエストメッセージ


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="pb.Task"/>

### Task
タスクの共通メッセージ


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Description | [string](#string) |  |  |
| Status | [Status](#pb.Status) |  |  |






<a name="pb.UpdateTaskRequest"/>

### UpdateTaskRequest
タスク更新のリクエストメッセージ


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Description | [string](#string) |  |  |





 


<a name="pb.Status"/>

### Status
ステータス

| Name | Number | Description |
| ---- | ------ | ----------- |
| DOING | 0 |  |
| COMPLETE | 1 |  |


 

 


<a name="pb.TaskService"/>

### TaskService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTask | [GetTaskRequest](#pb.GetTaskRequest) | [Task](#pb.GetTaskRequest) | タスク取得 |
| CreateTask | [CreateTaskRequest](#pb.CreateTaskRequest) | [Task](#pb.CreateTaskRequest) | タスク登録 |
| UpdateTask | [UpdateTaskRequest](#pb.UpdateTaskRequest) | [.google.protobuf.Empty](#pb.UpdateTaskRequest) | タスク更新 |
| CompleteTask | [CompleteTaskRequest](#pb.CompleteTaskRequest) | [.google.protobuf.Empty](#pb.CompleteTaskRequest) | タスク完了 |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

