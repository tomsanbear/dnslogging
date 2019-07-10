@0x9f556eafdad48795;
using Go = import "/go.capnp";
$Go.package("dnsproto");
$Go.import("zombiezen.com/go/capnproto2/example");

struct Data {
  request  @0   :Request;
  response @1   :Response;

  struct Request {
    question @0 :List(Text);
  }
  struct Response {
    answers @0 :List(Text);
  }
}