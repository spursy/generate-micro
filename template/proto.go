package template

var (
	ProtoSRV = `syntax = "proto3";
package {{dehyphen .Alias}};

service {{title .Alias}}Dispatcher {
  // Demo
  rpc GetDemo(DemoRequest) returns (DemoListReply) {}
}

message DemoRequest {
  string k = 1;
}

message DemoListReply {
  int32 status = 1;
  string message = 2;
  repeated DemoObject demo_list = 3;
}

message DemoObject {
  uint64 id = 1;
  string title = 2;
}
`
)



