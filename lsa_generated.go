// This file was autogenerated using go run mkcode.go -- lsa.go
// DO NOT EDIT.

package ntdll

import "unsafe"

var (
	procLsaOpenPolicy             = modntdll.NewProc("LsaOpenPolicy")
	procLsaAddAccountRights       = modntdll.NewProc("LsaAddAccountRights")
	procLsaEnumerateAccountRights = modntdll.NewProc("LsaEnumerateAccountRights")
	procLsaRemoveAccountRights    = modntdll.NewProc("LsaRemoveAccountRights")
	procLsaClose                  = modntdll.NewProc("LsaClose")
)

// INOUT-parameter: PolicyHandle.
func LsaOpenPolicy(
	SystemName *LsaUnicodeString,
	ObjectAttributes *LsaObjectAttributes,
	DesiredAccess AccessMask,
	PolicyHandle *LsaHandle,
) NtStatus {
	r0, _, _ := procLsaOpenPolicy.Call(uintptr(unsafe.Pointer(SystemName)),
		uintptr(unsafe.Pointer(ObjectAttributes)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(PolicyHandle)))
	return NtStatus(r0)
}

func LsaAddAccountRights(
	PolicyHandle LsaHandle,
	AccountSid *Sid,
	UserRights *LsaUnicodeString,
	CountOfRights uint32,
) NtStatus {
	r0, _, _ := procLsaAddAccountRights.Call(uintptr(PolicyHandle),
		uintptr(unsafe.Pointer(AccountSid)),
		uintptr(unsafe.Pointer(UserRights)),
		uintptr(CountOfRights))
	return NtStatus(r0)
}

// OUT-parameter: UserRights, CountOfRights.
func LsaEnumerateAccountRights(
	PolicyHandle LsaHandle,
	AccountSid *Sid,
	UserRights *LsaUnicodeString,
	CountOfRights *uint32,
) NtStatus {
	r0, _, _ := procLsaEnumerateAccountRights.Call(uintptr(PolicyHandle),
		uintptr(unsafe.Pointer(AccountSid)),
		uintptr(unsafe.Pointer(UserRights)),
		uintptr(unsafe.Pointer(CountOfRights)))
	return NtStatus(r0)
}

func LsaRemoveAccountRights(
	PolicyHandle LsaHandle,
	AccountSid *Sid,
	AllRights bool,
	UserRights *LsaUnicodeString,
	CountOfRights uint32,
) NtStatus {
	r0, _, _ := procLsaRemoveAccountRights.Call(uintptr(PolicyHandle),
		uintptr(unsafe.Pointer(AccountSid)),
		fromBool(AllRights),
		uintptr(unsafe.Pointer(UserRights)),
		uintptr(CountOfRights))
	return NtStatus(r0)
}

func LsaClose(
	ObjectHandle LsaHandle,
) NtStatus {
	r0, _, _ := procLsaClose.Call(uintptr(ObjectHandle))
	return NtStatus(r0)
}
