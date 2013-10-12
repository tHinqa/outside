// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

/*
Register all entry-points in wsock32.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package wsock32

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("wsock32.dll", false, EntryPoints)
	outside.AddEPs("wsock32.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"AcceptEx",
	"EnumProtocolsA",
	"GetAcceptExSockaddrs",
	"GetAddressByNameA",
	"GetNameByTypeA",
	"GetServiceA",
	"GetTypeByNameA",
	"MigrateWinsockConfiguration",
	"NPLoadNameSpaces",
	"SetServiceA",
	"TransmitFile",
	"WEP",
	"WSAAsyncGetHostByAddr",
	"WSAAsyncGetHostByName",
	"WSAAsyncGetProtoByName",
	"WSAAsyncGetProtoByNumber",
	"WSAAsyncGetServByName",
	"WSAAsyncGetServByPort",
	"WSAAsyncSelect",
	"WSACancelAsyncRequest",
	"WSACancelBlockingCall",
	"WSACleanup",
	"WSAGetLastError",
	"WSAIsBlocking",
	"WSARecvEx",
	"WSASetBlockingHook",
	"WSASetLastError",
	"WSAStartup",
	"WSAUnhookBlockingHook",
	"WSApSetPostRoutine",
	"__WSAFDIsSet",
	"accept",
	"bind",
	"closesocket",
	"connect",
	"dn_expand",
	"gethostbyaddr",
	"gethostbyname",
	"gethostname",
	"getnetbyname",
	"getpeername",
	"getprotobyname",
	"getprotobynumber",
	"getservbyname",
	"getservbyport",
	"getsockname",
	"getsockopt",
	"htonl",
	"htons",
	"inet_addr",
	"inet_network",
	"inet_ntoa",
	"ioctlsocket",
	"listen",
	"ntohl",
	"ntohs",
	"rcmd",
	"recv",
	"recvfrom",
	"rexec",
	"rresvport",
	"s_perror",
	"select",
	"send",
	"sendto",
	"sethostname",
	"setsockopt",
	"shutdown",
	"socket",
}

var UnicodeEntryPoints = outside.EPs{
	"EnumProtocolsW",
	"GetAddressByNameW",
	"GetNameByTypeW",
	"GetServiceW",
	"GetTypeByNameW",
	"SetServiceW",
}
