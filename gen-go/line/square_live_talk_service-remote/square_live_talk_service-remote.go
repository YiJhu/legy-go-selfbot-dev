// Code generated by Thrift Compiler (0.20.0). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"line"
)

var _ = line.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  AcceptSpeakersResponse acceptSpeakers()")
  fmt.Fprintln(os.Stderr, "  AcceptToChangeRoleResponse acceptToChangeRole()")
  fmt.Fprintln(os.Stderr, "  AcceptToListenResponse acceptToListen()")
  fmt.Fprintln(os.Stderr, "  AcceptToSpeakResponse acceptToSpeak()")
  fmt.Fprintln(os.Stderr, "  CancelToSpeakResponse cancelToSpeak()")
  fmt.Fprintln(os.Stderr, "  EndLiveTalkResponse endLiveTalk()")
  fmt.Fprintln(os.Stderr, "  FetchLiveTalkEventsResponse fetchLiveTalkEvents()")
  fmt.Fprintln(os.Stderr, "  FindLiveTalkByInvitationTicketResponse findLiveTalkByInvitationTicket()")
  fmt.Fprintln(os.Stderr, "  ForceEndLiveTalkResponse forceEndLiveTalk()")
  fmt.Fprintln(os.Stderr, "  GetLiveTalkInfoForNonMemberResponse getLiveTalkInfoForNonMember()")
  fmt.Fprintln(os.Stderr, "  GetLiveTalkInvitationUrlResponse getLiveTalkInvitationUrl()")
  fmt.Fprintln(os.Stderr, "  GetLiveTalkSpeakersForNonMemberResponse getLiveTalkSpeakersForNonMember()")
  fmt.Fprintln(os.Stderr, "  GetSquareInfoByChatMidResponse getSquareInfoByChatMid()")
  fmt.Fprintln(os.Stderr, "  InviteToChangeRoleResponse inviteToChangeRole()")
  fmt.Fprintln(os.Stderr, "  InviteToListenResponse inviteToListen()")
  fmt.Fprintln(os.Stderr, "  InviteToLiveTalkResponse inviteToLiveTalk()")
  fmt.Fprintln(os.Stderr, "  InviteToSpeakResponse inviteToSpeak()")
  fmt.Fprintln(os.Stderr, "  JoinLiveTalkResponse joinLiveTalk()")
  fmt.Fprintln(os.Stderr, "  KickOutLiveTalkParticipantsResponse kickOutLiveTalkParticipants()")
  fmt.Fprintln(os.Stderr, "  RejectSpeakersResponse rejectSpeakers()")
  fmt.Fprintln(os.Stderr, "  RejectSpeakersResponse rejectToSpeak()")
  fmt.Fprintln(os.Stderr, "  ReportLiveTalkResponse reportLiveTalk()")
  fmt.Fprintln(os.Stderr, "  ReportLiveTalkSpeakerResponse reportLiveTalkSpeaker()")
  fmt.Fprintln(os.Stderr, "  RequestToListenResponse requestToListen()")
  fmt.Fprintln(os.Stderr, "  RequestToSpeakResponse requestToSpeak()")
  fmt.Fprintln(os.Stderr, "  StartLiveTalkResponse startLiveTalk()")
  fmt.Fprintln(os.Stderr, "  UpdateLiveTalkAttrsResponse updateLiveTalkAttrs()")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  var cfg *thrift.TConfiguration = nil
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans = thrift.NewTSocketConf(net.JoinHostPort(host, portStr), cfg)
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransportConf(trans, cfg)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := line.NewSquareLiveTalkServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "acceptSpeakers":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "AcceptSpeakers requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.AcceptSpeakers(context.Background()))
    fmt.Print("\n")
    break
  case "acceptToChangeRole":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "AcceptToChangeRole requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.AcceptToChangeRole(context.Background()))
    fmt.Print("\n")
    break
  case "acceptToListen":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "AcceptToListen requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.AcceptToListen(context.Background()))
    fmt.Print("\n")
    break
  case "acceptToSpeak":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "AcceptToSpeak requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.AcceptToSpeak(context.Background()))
    fmt.Print("\n")
    break
  case "cancelToSpeak":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "CancelToSpeak requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.CancelToSpeak(context.Background()))
    fmt.Print("\n")
    break
  case "endLiveTalk":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "EndLiveTalk requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.EndLiveTalk(context.Background()))
    fmt.Print("\n")
    break
  case "fetchLiveTalkEvents":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "FetchLiveTalkEvents requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.FetchLiveTalkEvents(context.Background()))
    fmt.Print("\n")
    break
  case "findLiveTalkByInvitationTicket":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "FindLiveTalkByInvitationTicket requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.FindLiveTalkByInvitationTicket(context.Background()))
    fmt.Print("\n")
    break
  case "forceEndLiveTalk":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "ForceEndLiveTalk requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.ForceEndLiveTalk(context.Background()))
    fmt.Print("\n")
    break
  case "getLiveTalkInfoForNonMember":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetLiveTalkInfoForNonMember requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetLiveTalkInfoForNonMember(context.Background()))
    fmt.Print("\n")
    break
  case "getLiveTalkInvitationUrl":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetLiveTalkInvitationUrl requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetLiveTalkInvitationUrl(context.Background()))
    fmt.Print("\n")
    break
  case "getLiveTalkSpeakersForNonMember":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetLiveTalkSpeakersForNonMember requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetLiveTalkSpeakersForNonMember(context.Background()))
    fmt.Print("\n")
    break
  case "getSquareInfoByChatMid":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetSquareInfoByChatMid requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetSquareInfoByChatMid(context.Background()))
    fmt.Print("\n")
    break
  case "inviteToChangeRole":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "InviteToChangeRole requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.InviteToChangeRole(context.Background()))
    fmt.Print("\n")
    break
  case "inviteToListen":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "InviteToListen requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.InviteToListen(context.Background()))
    fmt.Print("\n")
    break
  case "inviteToLiveTalk":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "InviteToLiveTalk requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.InviteToLiveTalk(context.Background()))
    fmt.Print("\n")
    break
  case "inviteToSpeak":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "InviteToSpeak requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.InviteToSpeak(context.Background()))
    fmt.Print("\n")
    break
  case "joinLiveTalk":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "JoinLiveTalk requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.JoinLiveTalk(context.Background()))
    fmt.Print("\n")
    break
  case "kickOutLiveTalkParticipants":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "KickOutLiveTalkParticipants requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.KickOutLiveTalkParticipants(context.Background()))
    fmt.Print("\n")
    break
  case "rejectSpeakers":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "RejectSpeakers requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.RejectSpeakers(context.Background()))
    fmt.Print("\n")
    break
  case "rejectToSpeak":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "RejectToSpeak requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.RejectToSpeak(context.Background()))
    fmt.Print("\n")
    break
  case "reportLiveTalk":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "ReportLiveTalk requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.ReportLiveTalk(context.Background()))
    fmt.Print("\n")
    break
  case "reportLiveTalkSpeaker":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "ReportLiveTalkSpeaker requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.ReportLiveTalkSpeaker(context.Background()))
    fmt.Print("\n")
    break
  case "requestToListen":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "RequestToListen requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.RequestToListen(context.Background()))
    fmt.Print("\n")
    break
  case "requestToSpeak":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "RequestToSpeak requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.RequestToSpeak(context.Background()))
    fmt.Print("\n")
    break
  case "startLiveTalk":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "StartLiveTalk requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.StartLiveTalk(context.Background()))
    fmt.Print("\n")
    break
  case "updateLiveTalkAttrs":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "UpdateLiveTalkAttrs requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.UpdateLiveTalkAttrs(context.Background()))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
